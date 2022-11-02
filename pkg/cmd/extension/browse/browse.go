package browse

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/cli/cli/v2/git"
	"github.com/cli/cli/v2/internal/config"
	"github.com/cli/cli/v2/internal/ghrepo"
	"github.com/cli/cli/v2/pkg/extensions"
	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/cli/v2/pkg/search"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

const pagingOffset = 24

// TODO description color is low-contrast in most settings

type ExtBrowseOpts struct {
	Cmd      *cobra.Command
	Browser  ibrowser
	IO       *iostreams.IOStreams
	Searcher search.Searcher
	Em       extensions.ExtensionManager
	Client   *http.Client
	Logger   *log.Logger
	Cfg      config.Config
	Rg       readmeGetter
	Debug    bool
}

type ibrowser interface {
	Browse(string) error
}

type extEntry struct {
	URL         string
	Owner       string
	Name        string
	FullName    string
	Readme      string
	Stars       int
	Installed   bool
	Official    bool
	description string
}

func (e extEntry) Title() string {
	var installed string
	var official string

	if e.Installed {
		installed = " [green](installed)"
	}

	if e.Official {
		official = " [yellow](official)"
	}

	return fmt.Sprintf("%s%s%s", e.FullName, official, installed)
}

func (e extEntry) Description() string { return e.description }

type extensionListUI interface {
	FindSelected() (extEntry, int)
	Filter(text string)
	ToggleInstalled(int)
	Focus()
	Reset()
	Refresh()
	PageDown() int
	PageUp() int
	ScrollUp() int
	ScrollDown() int
}

type extList struct {
	list       *tview.List
	extEntries []extEntry
	app        *tview.Application
	logger     *log.Logger
	filter     string
}

func newExtList(app *tview.Application, list *tview.List, extEntries []extEntry, logger *log.Logger) extensionListUI {
	list.SetWrapAround(false)
	list.SetBorderPadding(1, 1, 1, 1)
	el := &extList{
		list:       list,
		extEntries: extEntries,
		app:        app,
		logger:     logger,
	}
	el.Reset()
	return el
}

func (el *extList) ToggleInstalled(ix int) {
	ee := el.extEntries[ix]
	ee.Installed = !ee.Installed
	el.extEntries[ix] = ee
}

func (el *extList) Focus() {
	el.app.SetFocus(el.list)
}

func (el *extList) Refresh() {
	el.Reset()
	el.Filter(el.filter)
}

func (el *extList) Reset() {
	el.list.Clear()
	for _, ee := range el.extEntries {
		el.list.AddItem(ee.Title(), ee.Description(), rune(0), func() {})
	}
}

func (el *extList) PageDown() int {
	el.list.SetCurrentItem(el.list.GetCurrentItem() + pagingOffset)
	return el.list.GetCurrentItem()
}

func (el *extList) PageUp() int {
	i := el.list.GetCurrentItem() - pagingOffset
	if i < 0 {
		i = 0
	}
	el.list.SetCurrentItem(i)
	return el.list.GetCurrentItem()
}

func (el *extList) ScrollDown() int {
	el.list.SetCurrentItem(el.list.GetCurrentItem() + 1)
	return el.list.GetCurrentItem()
}

func (el *extList) ScrollUp() int {
	i := el.list.GetCurrentItem() - 1
	if i < 0 {
		i = 0
	}
	el.list.SetCurrentItem(i)
	return el.list.GetCurrentItem()
}

func (el *extList) FindSelected() (extEntry, int) {
	if el.list.GetItemCount() == 0 {
		return extEntry{}, -1
	}
	title, desc := el.list.GetItemText(el.list.GetCurrentItem())
	for x, e := range el.extEntries {
		if e.Title() == title && e.Description() == desc {
			return e, x
		}
	}
	return extEntry{}, -1
}

func (el *extList) Filter(text string) {
	el.filter = text
	if text == "" {
		return
	}
	el.list.Clear()
	for _, ee := range el.extEntries {
		if strings.Contains(ee.Title()+ee.Description(), text) {
			el.list.AddItem(ee.Title(), ee.Description(), rune(0), func() {})
		}
	}
}

func ExtBrowse(opts ExtBrowseOpts) error {
	if !opts.IO.CanPrompt() {
		return errors.New("command requires interactive terminal")
	}

	if opts.Debug {
		f, err := os.CreateTemp("/tmp", "extBrowse-*.txt")
		if err != nil {
			return err
		}
		defer os.Remove(f.Name())

		opts.Logger = log.New(f, "", log.Lshortfile)
	} else {
		opts.Logger = log.New(io.Discard, "", 0)
	}

	installed := opts.Em.List()

	opts.IO.StartProgressIndicator()
	result, err := opts.Searcher.Repositories(search.Query{
		Kind:  search.KindRepositories,
		Limit: 1000,
		Qualifiers: search.Qualifiers{
			Topic: []string{"gh-extension"},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to search for extensions: %w", err)
	}

	host, _ := opts.Cfg.DefaultHost()

	extEntries := []extEntry{}

	for _, repo := range result.Items {
		ee := extEntry{
			URL:         "https://" + host + "/" + repo.FullName,
			FullName:    repo.FullName,
			Owner:       repo.Owner.Login,
			Name:        repo.Name,
			Stars:       repo.StargazersCount,
			description: repo.Description,
		}
		for _, v := range installed {
			// TODO consider a Repo() on Extension interface
			var installedRepo string
			if u, err := git.ParseURL(v.URL()); err == nil {
				if r, err := ghrepo.FromURL(u); err == nil {
					installedRepo = ghrepo.FullName(r)
				}
			}
			if repo.FullName == installedRepo {
				ee.Installed = true
			}
		}
		if ee.Owner == "cli" || ee.Owner == "github" {
			ee.Official = true
		}

		extEntries = append(extEntries, ee)
	}

	cacheTTL, _ := time.ParseDuration("24h")
	opts.Rg = newReadmeGetter(opts.Client, cacheTTL)

	app := tview.NewApplication()

	outerFlex := tview.NewFlex()
	innerFlex := tview.NewFlex()

	header := tview.NewTextView().SetText(fmt.Sprintf("browsing %d gh extensions", len(extEntries)))
	header.SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite)

	filter := tview.NewInputField().SetLabel("filter: ")
	filter.SetFieldBackgroundColor(tcell.ColorGray)
	filter.SetBorderPadding(0, 0, 20, 20)

	list := tview.NewList()
	list.SetTitleColor(tcell.ColorPurple)
	list.SetSelectedTextColor(tcell.ColorWhite)
	list.SetSelectedBackgroundColor(tcell.ColorPurple)

	readme := tview.NewTextView()
	readme.SetBorderPadding(1, 1, 0, 1)
	readme.SetBorder(true).SetBorderColor(tcell.ColorPurple)

	help := tview.NewTextView()
	help.SetText(
		"/: filter i/r: install/remove w: open in browser pgup/pgdn: scroll readme q: quit")
	help.SetTextAlign(tview.AlignCenter)

	extList := newExtList(app, list, extEntries, opts.Logger)

	loadSelectedReadme := func() {
		ee, ix := extList.FindSelected()
		if ix < 0 {
			opts.Logger.Println(fmt.Errorf("tried to find entry, but: %w", err))
			return
		}
		fullName := ee.FullName
		rm, err := opts.Rg.Get(fullName)
		if err != nil {
			opts.Logger.Println(err.Error())
			readme.SetText("unable to fetch readme :(")
			return
		}

		// using glamour directly because if I don't horrible things happen
		rendered, err := glamour.Render(rm, "dark")
		if err != nil {
			opts.Logger.Println(err.Error())
			readme.SetText("unable to render readme :(")
			return
		}

		app.QueueUpdateDraw(func() {
			readme.SetText("")
			readme.SetDynamicColors(true)

			w := tview.ANSIWriter(readme)
			_, _ = w.Write([]byte(rendered))

			readme.ScrollToBeginning()
		})
	}

	// Force fetching of initial readme:
	go loadSelectedReadme()

	filter.SetChangedFunc(func(text string) {
		extList.Filter(text)
		go loadSelectedReadme()
	})

	filter.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			extList.Focus()
		case tcell.KeyEscape:
			filter.SetText("")
			extList.Reset()
			extList.Focus()
		}
	})

	innerFlex.SetDirection(tview.FlexColumn)
	innerFlex.AddItem(list, 0, 1, true)
	innerFlex.AddItem(readme, 0, 1, false)

	outerFlex.SetDirection(tview.FlexRow)
	outerFlex.AddItem(header, 1, -1, false)
	outerFlex.AddItem(filter, 1, -1, false)
	outerFlex.AddItem(innerFlex, 0, 1, true)
	outerFlex.AddItem(help, 1, -1, false)

	// modal is used to output the result of install/remove operations
	modal := tview.NewModal().AddButtons([]string{"ok"}).SetDoneFunc(func(_ int, _ string) {
		app.SetRoot(outerFlex, true)
		extList.Refresh()
	})

	modal.SetBackgroundColor(tcell.ColorPurple)

	app.SetRoot(outerFlex, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if filter.HasFocus() {
			return event
		}

		switch event.Rune() {
		case 'q':
			app.Stop()
		case 'k':
			extList.ScrollUp()
			readme.SetText("...fetching readme...")
			go loadSelectedReadme()
		case 'j':
			extList.ScrollDown()
			readme.SetText("...fetching readme...")
			go loadSelectedReadme()
		case 'w':
			ee, ix := extList.FindSelected()
			if ix < 0 {
				opts.Logger.Println("failed to find selected entry")
				return nil
			}
			err = opts.Browser.Browse(ee.URL)
			if err != nil {
				opts.Logger.Println(fmt.Errorf("could not open browser for '%s': %w", ee.URL, err))
			}
		case 'i':
			ee, ix := extList.FindSelected()
			if ix < 0 {
				opts.Logger.Println("failed to find selected entry")
				return nil
			}
			repo, err := ghrepo.FromFullName(ee.FullName)
			if err != nil {
				opts.Logger.Println(fmt.Errorf("failed to install '%s't: %w", ee.FullName, err))
				return nil
			}

			modal.SetText(fmt.Sprintf("Installing %s...", ee.FullName))
			app.SetRoot(modal, true)
			// I could eliminate this with a goroutine but it seems to be working fine
			app.ForceDraw()
			err = opts.Em.Install(repo, "")
			if err != nil {
				modal.SetText(fmt.Sprintf("Failed to install %s: %s", ee.FullName, err.Error()))
			} else {
				modal.SetText(fmt.Sprintf("Installed %s!", ee.FullName))
			}
			extList.ToggleInstalled(ix)
		case 'r':
			ee, ix := extList.FindSelected()
			if ix < 0 {
				opts.Logger.Println("failed to find selected extension")
				return nil
			}
			modal.SetText(fmt.Sprintf("Removing %s...", ee.FullName))
			app.SetRoot(modal, true)
			// I could eliminate this with a goroutine but it seems to be working fine
			app.ForceDraw()

			err = opts.Em.Remove(strings.TrimPrefix(ee.Name, "gh-"))
			if err != nil {
				modal.SetText(fmt.Sprintf("Failed to remove %s: %s", ee.FullName, err.Error()))
			} else {
				modal.SetText(fmt.Sprintf("Removed %s.", ee.FullName))
			}
			extList.ToggleInstalled(ix)
		case ' ':
			// The shift check works on windows and not linux/mac:
			if event.Modifiers()&tcell.ModShift != 0 {
				extList.PageUp()
			} else {
				extList.PageDown()
			}
			go loadSelectedReadme()
		case '/':
			app.SetFocus(filter)
			return nil
		}
		switch event.Key() {
		case tcell.KeyUp:
			extList.ScrollUp()
			go loadSelectedReadme()
			return nil
		case tcell.KeyDown:
			extList.ScrollDown()
			go loadSelectedReadme()
			return nil
		case tcell.KeyEscape:
			filter.SetText("")
			extList.Reset()
		case tcell.KeyCtrlSpace:
			// The ctrl check works on windows/mac and not windows:
			extList.PageUp()
			go loadSelectedReadme()
		case tcell.KeyCtrlJ:
			extList.PageDown()
			go loadSelectedReadme()
		case tcell.KeyCtrlK:
			extList.PageUp()
			go loadSelectedReadme()
		case tcell.KeyPgUp:
			row, col := readme.GetScrollOffset()
			if row > 0 {
				readme.ScrollTo(row-2, col)
			}
			return nil
		case tcell.KeyPgDn:
			row, col := readme.GetScrollOffset()
			readme.ScrollTo(row+2, col)
			return nil
		}

		return event
	})

	opts.IO.StopProgressIndicator()
	if err := app.Run(); err != nil {
		return err
	}

	return nil
}