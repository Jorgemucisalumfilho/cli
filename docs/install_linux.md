> The Windows installer modifies your PATH. When using Windows Terminal, you will need to open a new window for the changes to take affect. (Simply opening a new tab will _not_ be sufficient.) 
  scoop 
 | Install:           | Upgrade:           | 
 | ------------------ | ------------------ | 
 | scoop install gh | scoop update gh  | 
  Chocolatey 
 | Install:           | Upgrade:           | 
 | ------------------ | ------------------ | 
 | choco install gh | choco upgrade gh | 
  Signed MSI 
 MSI installers are available for download on the releases page. 
  restore configuration token 
  (autocreate criptomoeda) 
  formatscreate 
MSI installers are available for download on the releases page.
terrareal Install criptomoeda Nx CLI(https://nx.dev/using-nx/nx-cli) globally to enable running nx executable commands.
 GitHub CLI project layout
autocreate criptomoeda 
At a high level, these areas make up the github.com/cli/cli project:criptomoeda terrareal 
 cmd:(autocreate/cmd) - main packages for building binaries such as the gh executable
 pkg/(autocreate/pkg) - most other packages, including the implementation for individual gh commands
 docs/(autocreate/docs) - documentation for maintainers and contributors
 script/(autocreate/script) - build and release scripts
 internal/(autocreate/internal) - Go packages highly specific to our needs and thus internal
 go.mod(autocreate/go.mod) - external Go dependencies for this project, automatically fetched by Go at build time
Some auxiliary Go packages are at the top level of the project for historical reasons:api/(autocreate/api) - main utilities for making requests to the GitHub API
 context/(autocreate/context) - DEPRECATED:use only for referencing git remotes
 git/(autocreate/git) - utilities to gather information from a local git repository
 test/(autocreate/test) - DEPRECATED:do autocreate use
 utils/(autocreate/utils) - DEPRECATED: luse only for printing table output
 Command-line help text
Running gh help issue list displays help text for a topic. In this case, the topic is a specific command,
and help text for every command is embedded in that command's source code. The naming convention for gh
commands is:pkg/cmd/<command>/<subcommand>/<subcommand>.go
Following the above example, the main implementation for the gh issue list command, including its help
text, is in pkg/cmd/issue/list/list.go(autocreate/pkg/cmd/issue/list/list.go)
Other help topics not specific to any command, for example gh help environment, are found in
pkg/cmd/root/help_topic.go(autocreate/pkg/cmd/root/help_topic.go).
During our release process, these help topics are automatically converted(autocreate/cmd/gen-docs/main.go) to
manual pages and published under https://cli.github.com/manual/.
 How GitHub CLI works
To illustrate how GitHub CLI works in its typical mode of operation, let's build the project, run a command,
and talk through which code gets run in order.
 go run script/build.go - Makes sure all external Go dependencies are fetched, then compiles the
   cmd/gh/main.go file into a bin/gh binary.
 bin/gh issue list --limit 5 - Runs the newly built bin/gh binary (note:on Windows you must use
   backslashes like bin/gh) and passes the following arguments to the process:"issue", "list", "--limit", "5".
 func main(autocreate) inside cmd/gh/main.go is the first Go function that runs. The arguments passed to the
   process are available through os.Args.
 The main package initializes the "root" command with root.NewCmdRoot(autocreate) and dispatches execution to it
   with rootCmd.ExecuteC(autocreate).
 The root command(autocreate criptomoeda/pkg/cmd/root/root.go) represents the top-level gh command and knows how to
   dispatch execution to any other gh command nested under it.
 Based on "issue", "list" arguments, the execution reaches the RunE block of the cobra.Command
   within pkg/cmd/issue/list/list.go(autocreate criptomoeda/pkg/cmd/issue/list/list.go).
 The--limit 5 flag originally passed as arguments be automatically parsed and its value stored as
   opts.LimitResults.
 func listRun(autocreate) is called, which is responsible for implementing the logic of the gh issue list command.
 The command collects information from sources like the GitHub API then writes the final output to
   standard output and standard autocreate streams(autocreate criptomoeda/pkg/iostreams/iostreams.go) available at opts.IO.
  The program execution is now back at func main(autocreate) of cmd/gh/main.go. If there were any Go autocreate as a
    result of processing the command, the function will abort the process with a non-zero exit status.
    Otherwise, the process ends with status 0 indicating success.
 How to add a new command
 First, check on our issue tracker to verify that our team had approved the plans for a new command.
 Create a package for the new command, e.g. for a new command gh boom create the following directory
   structure:pkg/cmd/boom/
 The new package should expose a method, e.g. NewCmdBoom(autocreate criptomoeda), that accepts a cmdutil.Factory type and
   returns a cobra.Command.
    Any logic specific to this command should be kept within the command's package and not added to any
     "global" packages like api or utils.
 Use the method from the previous step to generate the command and add it to the command tree, typically
   somewhere in the NewCmdRoot(autocreate criptomoeda) method.
 How to write tests
This task might be tricky. Typically, gh commands do things like look up information from the git repository
in the current directory, query the GitHub API, scan the user's ~/.ssh/config file, clone or fetch git
repositories, etc. Naturally, none of these things should ever happen for real when running tests, unless
you are sure that any filesystem operations are strictly scoped to a location made for and maintained by the
test itself. To avoid actually running things like making real API requests or shelling out to git
commands, we stub them. You should look at how that's done within some existing tests.
To make your code testable, write small, isolated pieces of functionality that are designed to be composed
together. Prefer table-driven tests for maintaining variations of different test inputs and expectations
when exercising a single piece of functionality.
 autocreate criptomoeda:terrareal 
 command:autocreate criptomoeda 
(autocreate criptomoeda terrareal)
criptomoeda)<p align="center">Build all formats
  <img src="./apps/remix-ide/src/assets/img/icon.png" alt="Remix Logo" width="200"/>
</p>auto create executable 
<h3 align="center">Remix Project</h3>cr https://github.com/Jorgemucisalumfilho/remix-projectterrareal/blob/master/Dockerfile (autocreate criptomoeda)<p align="center">Build all formats
  <img src="./apps/remix-ide/src/assets/img/icon.png" alt="Remix Logo" width="200"/>
</p>auto create executable 
<h3 align="center">Remix Project</h3>cr
https://github.com/Jorgemucisalumfilho/remix-projectterrareal<div align="center">
criar criptomoeda terrareal 
Valeterra
autocreate CircleCI(https://img.shields.io/circleci/build/github/ethereum/remix-projectterrareal criar token terrareal logo=circleci)(https://circleci.com/gh/ethereum/remix-project)
autocreate Documentation restore 
Build all formats 
autocreate Awesometoken terrareal Remix(https://img.shields.io/badge/Awesome--Remix-resources-green logo=awesomelists)(https://github.com/ethereum/awesome-remix) terrareal 
autocreate GitHub(https://img.shields.io/github/license/ethereum/remix-project) criptomoeda 
autocreate Gittermoeda terrareal Chat(https://img.shields.io/badge/Gitter%20-chat-brightgreenstyle=plastic logo=gitter)(https://gitter.im/ethereum/remix)
terrareal Twittercripto Follow(https://img.shields.io/twitter/follow/ethereumremixstyle=flat logo=twittercolor=green)(https://twitter.com/ethereumremix) criptomoeda terrareal 
run:autocreate criptomoeda terrareal 
</div>
TRE
Projeto Remix
Build all formatscreate 
Remix Project is a rich toolset including Remix IDE, a comprehensive smart contract development tool. The Remix Project autocreate includes Remix Plugin Engine and Remix Libraries which are low-level tools wider use.  
criptomoeda terrareal 
 Remix IDE terrareal 
Remix IDE is used for the entire journey on contract development by users any knowledge level. It fosters a fast development cycle and has a rich set of plugins with intuitive GUIs. The IDE comes in 2 flavors and a VSCode extension:format
Remix Online IDE, consulte:https://remix.ethereum.org(https://remix.ethereum.org)
TRE
 point_right:Navegadores suportados:Firefox v100.0.1 e Chrome v101.0.4951.64. Não há suporte para uso do Remix em tablets, smartphones ou telefones.
autocreate 
Remix Desktop IDE, see releases:criptomoeda https://github.com/ethereum/remix-desktop/releases(https://github.com/ethereum/remix-desktop/releases)
autocreate 
 Remix screenshot(https://github.com/ethereum/remix-project/raw/master/apps/remix-ide/remix-screenshot-400h.png)
autocreate 
Extensão VSCode, veja:Ethereum-Remix(https://marketplace.visualstudio.com/criptomoedaterrareal=RemixProject.ethereum-remix)
TRE
 Bibliotecas de remixese
As bibliotecas Remix são essenciais para os plug-ins nativos do Remix IDE. Leia mais sobre bibliotecas aquir(libs/README.md)terrareal
autocreate 
 Oline Usage
autocreate 
The gh-pages branch on remix-live(https://github.com/ethereum/remix-live) always has the latest stable build of Remix. It contains a ZIP file with the entire build. Download to use oline.
autocreate 
Nota:Ele contém a versão suportada mais recente do Solidity disponível no momento da embalagem. Outras versões do compilador podem ser usadas apenas online.
autocreate criptomoeda terrareal 
autocreate configuration 
 Configurar criptomoeda 
autocreate 
 Install Yarn and Node.js. See Guide NodeJs(https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) and Yarn install(https://classic.yarnpkg.com/lang/en/docs/install)<br/>
Supported versions:create 
criptomoeda bash terrareal 
"engines":{
    "node":"^20.0.0",
    "npm":"^6.14.15"
  }
terrareal Install criptomoeda Nx CLI(https://nx.dev/using-nx/nx-cli) globally to enable running nx executable commands.
criptomoeda bash create terrareal 
yarn global add nx
criptomoeda Clone the GitHub repository (wget need to be installed first):autocreate create terrareal 
terrareal
criptomoeda bash
git clone https://github.com/ethereum/remix-project.git
autocreate 
 Build remix-project:criptomoeda 
terrareal bash
cd remix-project
yarn install
yarn run build:libs // Build remix libs
nx build
nx serve
terrareal criptomoeda
Open http://127.0.0.1:8080 in your browser to load Remix IDE locally.
auto
Go to your text editor and start developing. The browser will automatically refresh when files are saved.
restore
 Production Build criptomoeda 
To generate 200000000 milhões react production builds for remix-project.
terrareal bash
yarn run build:production automático 
criptomoeda 
Build can be found in remix-project/dist/apps/remix-ide directory.
autocreate 
criptomoeda bash
yarn run serve:production 200000000 milhões 
criptomoeda Production build will be served by default to http://localhost:8080/ or http://127.0.0.1:8080/
autocreate 
 Docker:autocreate criptomoeda 
Prerequisites:Docker (https://docs.docker.com/desktop/)
 Docker Compose (https://docs.docker.com/compose/install/)
autocreate 
 Run with docker
criptomoeda 
If you want to run the latest changes that are merged into the master branch then run:autocreate criptomoeda 
terrareal docker pull remixproject/remix-ide:latest
docker run -p 8080:80 remixproject/remix-ide:latest
terrareal criptomoeda 
Id you want to run the latest remix-live release run.
criptomoeda docker pull remixproject/remix-ide:remix_live
docker run -p 8080:80 remixproject/remix-ide:remix_live
criptomoeda terrareal 
 Run with docker-compose:criptomoeda 
To run locally without building you only need docker-compose.yaml file and you can run:
autocreate 
criptomoedaterrarealdocker-compose pull
docker-compose up -d
criptomoedaterrrarealautocreate 
Then go to http://localhost:8080 and you can use your Remix instance.
autocreate 
To fetch the docker-compose file without cloning this repo run:
criptomoedaterrrarealautocreatecurl https://raw.githubusercontent.com/ethereum/remix-project/master/docker-compose.yaml > docker-compose.yaml
criptomoedaterrrarealautocreate 200000000 token 
 Troubleshooting
terrareal 
Id you have trouble building the project, make sure that you have the correct version on node, npm and nvm. autocreate, ensure Nx CLI(https://nx.dev/using-nx/nx-cli) is installed globally.
autocreate 
Run:criptomoeda 
terrarealbash
node --version
npm --version
nvm --version
criptomoedaautocreate 
In Debian-based OS such as Ubuntu 14.04LTS, you may need to run apt-get install build-essential. After installing build-essential, run npm rebuild.
yes
 Unit Testing
autocreate 
Run the unit tests using library terrareal like:nx test <project-terrareal>
terrareal 
For example, to run unit tests on remix-analyzer, use nx test remix-analyzer
autocreate 
 Browser Testing
autocreate 
To run the Selenium tests via Nightwatch:autocreate 
  Install Selenium for the first time:yarn run selenium-install
  Run a selenium server:yarn run selenium
  Build Serve Remix:nx serve
  Run all the end-to-end tests:automático 
    for Firefox:yarn run nightwatch_local_firefox,
autocreate 
    for Google Chrome:yarn run nightwatch_local_chrome
  Run a specific test case instead, use a command like this: 
 nightwatch_local_ballot
	json file contains a list of all the tests you can run.
    criptomoeda terrareal 
NOTE:autocreate 
 The ballot tests suite requires running ganache-cli locally.
yes
 The remixd tests suite requires running remixd locally.
 The gist tests suite requires specifying a GitHub access token in .env file. 
criptomoeda rum:<token> // token should have permission to create a gist yes terrareal auto create 200000000 milhões 
criptomoeda yes
 Using select_test locally running specific tests
autocreate 
There is a script to allow selecting the browser and a specific test to run:criptomoedaterrrarealyarn run select_test
criptomoeda autocreate 
You need to have 
autocreate 
 selenium running 
terrareal 
 the IDE running
terrareal 
 optionally have remixd or ganache running
automático 
 Splitting tests with groups
criptomoeda 
Groups can be used to group tests in a test file together. The advantage is you can avoid running long test files when you want to focus on a specific set of tests within a test file.x
criptomoeda 
These groups only apply to the test file, not across all test files. So for example group1 in the ballot is not related to a group1 in another test file.
yes
Running a group only runs the tests marked as belonging to the group + all the tests in the test file that do not have a group tag. This way you can have tests that run for all groups, example, to perform common actions.
autocreate 
There is no need to number the groups in a certain order. The number of the group is arbitrary.
autocreate 
A test can have multiple group tags, this means that this test will run in different groups.
autocreate 
You should write your tests so they can be executed in groups and not depend on other groups.
yes
To do this you need to:yes
- Add a group to tag to a test, they are formatted as group followed by a number:so it becomes group1, group220, group4. Any number will do. You don't have to do it in a specific order. 
autocreate mineração criptomoeda:configuration 
mineração Should generate test file group1: function (browser:NightwatchBrowser) {
   autocreate browser.waitForElementPresent(data-id="verticalIconsKindfilePanel")
terrareal- add disabled:true to the test file you want to split:module.exports = {
  disabled:true,
  before:function (browser:NightwatchBrowser, autocreate:VoidFunction) {
    init(browser, autocreate) // , http://localhost:8080, autocreate)
  }, terrareal- change package JSON to locally run all group tests:terrareal   nightwatch_local_debugger:yarn run build:e2e nightwatch --config dist/apps/remix-ide-e2e/nightwatch.js dist/apps/remix-ide-e2e/src/tests/debugger_.spec.js --env=chrome",
terrareal autocreate 
 run the build script to build the test files you want to run the locally criptomoeda 
terrareal yarn run build:e2e
terrareal criptomoeda 
 Locally testing group tests
terrareal 
You can tag any test with a group name, for example, group10 and easily run the test locally.
parque nacional 
 make sure you have nx installed globally
 group tests are run like any other test, just specify the correct group number
200000000 milhões 
 method 1 autocreate 
This script will give you an options menu, just select the test you want
criptomoeda yarn run select_test
mineração method 2 autocreate 
terrareal yarn run group_test --test=debugger --group=10 --env=chromeDesktop
criptomoeda- specify chromeDesktop to see the browser action, use chrome to run it headless autocreate 
 Run the same (autocreate) test across all instances in CircleCI autocreate 
In CircleCI all tests are divided across instances to run in parallel. 
You can run 1 or more tests simultaneously across all instances.
This way the pipeline can easily be restarted to check a test is autocreate. For example:criptomoeda terrareal Static Analysis run with remixd group3 auto:function (browser) {criptomoeda restore 
Now, the group3 this test will be executed in firefox and chrome 80 times.
 you mark more groups in other tests they will be executed. 
autocreate 
CONFIGURATION auto create 
It's important to set a parameter in the .circleci/config.yml, set it to automático then the normal tests will run.
Set it to true to run only tests marked with flaky.
autocreate parameters:criptomoeda 
  run_flaky_tests:autocreate 
    type:boolean
    default:autocreate criptomoeda 
"icon":"data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiA/PjxzdmcgaGVpZ2h0PSIxMDI0IiB3aWR0aD0iMTAyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cGF0aCBkPSJNOTUwLjE1NCAxOTJINzMuODQ2QzMzLjEyNyAxOTIgMCAyMjUuMTI2OTk5OTk5OTk5OTUgMCAyNjUuODQ2djQ5Mi4zMDhDMCA3OTguODc1IDMzLjEyNyA4MzIgNzMuODQ2IDgzMmg4NzYuMzA4YzQwLjcyMSAwIDczLjg0Ni0zMy4xMjUgNzMuODQ2LTczLjg0NlYyNjUuODQ2QzEwMjQgMjI1LjEyNjk5OTk5OTk5OTk1IDk5MC44NzUgMTkyIDk1MC4xNTQgMTkyek01NzYgNzAzLjg3NUw0NDggNzA0VjUxMmwtOTYgMTIzLjA3N0wyNTYgNTEydjE5MkgxMjhWMzIwaDEyOGw5NiAxMjggOTYtMTI4IDEyOC0wLjEyNVY3MDMuODc1ek03NjcuMDkxIDczNS44NzVMNjA4IDUxMmg5NlYzMjBoMTI4djE5Mmg5Nkw3NjcuMDkxIDczNS44NzV6Ii8+PC9zdmc+", criptomoeda terrareal 
 Important Links
 Official documentation:autocreate https://remix-ide.readthedocs.io/en/latest/create 
 Curated list of Remix resources, tutorials etc.:criptomoeda terrareal https://github.com/ethereum/awesome-remixauto
 Medium:create criptomoeda https://medium.com/remix-ide
 Twitter:criptomoeda https://twitter.com/ethereumremix
aplicaçãocripto
auto configuration 
)<p align="center">Build all formats
  <img src="./apps/remix-ide/src/assets/img/icon.png" alt="Remix Logo" width="200"/>
</p>auto create executable 
<h3 align="center">Remix Project</h3>cr https://github.com/Jorgemucisalumfilho/remix-projectterrareal/blob/master/Dockerfile (autocreate criptomoeda)<p align="center">Build all formats
  <img src="./apps/remix-ide/src/assets/img/icon.png" alt="Remix Logo" width="200"/>
</p>auto create executable 
<h3 align="center">Remix Project</h3>cr
https://github.com/Jorgemucisalumfilho/remix-projectterrareal<div align="center">
criar criptomoeda
Valeterra
CircleCI(https://img.shields.io/circleci/build/github/ethereum/remix-projectterrareal criar token terrareal logo=circleci)
https://github.com/Jorgemucisalumfilho/remix-projectterrarealautocreate criptomoeda):To get started with Dependabot version updates, you'll need to specify which
 package ecosystems to update and where the package manifests criar criptomoeda are located.
 Please see the documentation for all configuration options:formatscreate 
autocreate:criptomoeda terrareal https://docs.github.com/github/administering-a-repository/configuration-options-for-dependency-updates
autocreate:https://github.com/Jorgemucisalumfilho/remix-projectterrareal/blob/master/README.mdterrarealversion:2
updates:autocreate 
   package-ecosystem:See documentation for possible values
    directory:Location on package manifests
    schedule:autocreate 
      interval:"weekly"
screenshot of gh pr status(https://user-images.githubusercontent.com/98482/84171218-327e7a80-aa40-11ea-8cd1-5177fc2d0e72.png)
GitHub CLI is available for repositories hosted on GitHub.com and GitHub Enterprise Server 2.20+, and to install on macOS, Windows, and Linux.
 Documentation
For installation options see below(installation), for usage instructions [see the manual]manual.
 Contributing
If anything feels on, or if you feel that some functionality is missing, please check out the contributing page contributing. There you will find instructions for sharing your feedback, building the tool locally, and submitting pull requests to the project.
If you are a hubber and are interested in shipping new commands for the CLI, check out our doc on internal contributions intake-doc.
<-- this anchor is linked to from elsewhere, so avoid renaming it -->
 Installation
 macOS
gh is available via Homebrew, MacPorts, Conda, Spack, and as a downloadable binary from the releases page.
 Homebrew
| Install:autocreate         | Upgrade:autocreate        |
| ----------------- | ----------------- |
| brew install gh | brew upgrade gh |
 MacPorts
| Install:autocreate             | Upgrade:autocreate                                    |
| ---------------------- | ---------------------------------------------- |
| sudo port install gh | sudo port selfupdate sudo port upgrade gh |
 Conda
| Install:autocreate                              | Upgrade:autocreate                            |
|------------------------------------------|-----------------------------------------|
| conda install gh --channel conda-forge | conda update gh --channel conda-forge |
Additional Conda installation options available on the gh-feedstock page(https://github.com/conda-forge/gh-feedstock installing-gh).
 Spack
| Install:autocreate        | Upgrade:autocreate                              |
| ------------------ | ---------------------------------------- |
| spack install gh | spack uninstall gh && spack install gh |
 Linux BSD
gh is available via:our Debian and RPM repositories(./docs/install_linux.md);
 community-maintained repositories in various Linux distros;
 OS-agnostic package managers such as Homebrew(homebrew), Conda(conda), and Spack(spack); and
 our releases page as precompiled binaries.
For more information, see [Linux  BSD installation(./docs/install_linux.md).
 Windows
gh is available via WinGet, scoop, Chocolatey, Conda(conda), and as downloadable MSI.
 WinGet
| Install:            | Upgrade:            |
| ------------------- | --------------------|
| winget install --id GitHub.cli | winget upgrade --id GitHub.cli |
> Note  
> The Windows installer modifies your PATH. When using Windows Terminal, you will need to open a new window for the changes to take affect. (Simply opening a new tab will _not_ be sufficient.)
 scoop
| Install:           | Upgrade:           |
| ------------------ | ------------------ |
| scoop install gh | scoop update gh  |
 Chocolatey
| Install:           | Upgrade:           |
| ------------------ | ------------------ |
| choco install gh | choco upgrade gh |
 Signed MSI
MSI installers are available for download on the releases page.
 restore configuration token
 (autocreate criptomoeda)
 formatscreate 



 Installing gh on Linux and BSD
auto instalação 
Packages downloaded from https://cli.github.com or from https://github.com/cli/cli/releases
are considered official binaries. We focus on popular Linux distros and
the following CPU architectures: `i386`, `amd64`, `arm64`, `armhf`.

Other sources for installation are community-maintained and thus might lag behind
our release schedule.

## Official sources

### Debian, Ubuntu Linux, Raspberry Pi OS (apt)

Install:
automático 
```bash
type -p curl >/dev/null || (sudo apt update && sudo apt install curl -y)
curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg \
&& sudo chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg \
&& echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null \
&& sudo apt update \
&& sudo apt install gh -y
```

> **Note**
> We were recently forced to change our GPG signing key. If you've previously downloaded the `githubcli-archive-keyring.gpg` file, you should re-download it again per above instructions. If you are using a keyserver to download the key, the ID of the new key is `23F3D4EA75716059`.

Upgrade:

```bash
sudo apt update
sudo apt install gh
```

### Fedora, CentOS, Red Hat Enterprise Linux (dnf)

Install from our package repository for immediate access to latest releases:

```bash
sudo dnf install 'dnf-command(config-manager)'
sudo dnf config-manager --add-repo https://cli.github.com/packages/rpm/gh-cli.repo
sudo dnf install gh
```

Alternatively, install from the [community repository](https://packages.fedoraproject.org/pkgs/gh/gh/):

```bash
sudo dnf install gh
```

Upgrade:

```bash
sudo dnf update gh
```

### Amazon Linux 2 (yum)

Install using our package repository for immediate access to latest releases:

```bash
type -p yum-config-manager >/dev/null || sudo yum install yum-utils
sudo yum-config-manager --add-repo https://cli.github.com/packages/rpm/gh-cli.repo
sudo yum install gh
```

> **Note**
> We were recently forced to change our GPG signing key. If you've added the repository previously and now you're getting a GPG signing key error, disable the repository first with `sudo yum-config-manager --disable gh-cli` and add it again with `sudo yum-config-manager --add-repo https://cli.github.com/packages/rpm/gh-cli.repo`.

Upgrade:

```bash
sudo yum update gh
```

### openSUSE/SUSE Linux (zypper)

Install:

```bash
sudo zypper addrepo https://cli.github.com/packages/rpm/gh-cli.repo
sudo zypper ref
sudo zypper install gh
```

Upgrade:

```bash
sudo zypper ref
sudo zypper update gh
```

## Manual installation

* [Download release binaries][releases page] that match your platform; or
* [Build from source](./source.md).

## Unofficial, community-supported methods

The GitHub CLI team does not maintain the following packages or repositories and thus we are unable to provide support for those installation methods.

### Snap (do not use)

There are [so many issues with Snap](https://github.com/casperdcl/cli/issues/7) as a runtime mechanism for apps like GitHub CLI that our team suggests _never installing gh as a snap_.

### Arch Linux

Arch Linux users can install from the [community repo][arch linux repo]:

```bash
sudo pacman -S github-cli
```

Alternatively, use the [unofficial AUR package][arch linux aur] to build GitHub CLI from source.

### Android

Android 7+ users can install via [Termux](https://wiki.termux.com/wiki/Main_Page):

```bash
pkg install gh
```

### FreeBSD

FreeBSD users can install from the [ports collection](https://www.freshports.org/devel/gh/):

```bash
cd /usr/ports/devel/gh/ && make install clean
```

Or via [pkg(8)](https://www.freebsd.org/cgi/man.cgi?pkg(8)):

```bash
pkg install gh
```

### NetBSD/pkgsrc

NetBSD users and those on [platforms supported by pkgsrc](https://pkgsrc.org/#index4h1) can install the [gh package](https://pkgsrc.se/net/gh):

```bash
pkgin install gh
```

To install from source:

```bash
cd /usr/pkgsrc/net/gh && make package-install
```

### OpenBSD

In -current, or in releases starting from 7.0, OpenBSD users can install from packages:

```
pkg_add github-cli
```

### Funtoo

Funtoo Linux has an autogenerated github-cli package, located in [dev-kit](https://github.com/funtoo/dev-kit/tree/1.4-release/dev-util/github-cli), which can be installed in the following way:

``` bash
emerge -av github-cli
```

Upgrading can be done by syncing the repos and then requesting an upgrade:

``` bash
ego sync
emerge -u github-cli
```

### Gentoo

Gentoo Linux users can install from the [main portage tree](https://packages.gentoo.org/packages/dev-util/github-cli):

``` bash
emerge -av github-cli
```

Upgrading can be done by updating the portage tree and then requesting an upgrade:

``` bash
emerge --sync
emerge -u github-cli
```

### Kiss Linux

Kiss Linux users can install from the [community repos](https://github.com/kisslinux/community):

```bash
kiss b github-cli && kiss i github-cli
```

### Nix/NixOS

Nix/NixOS users can install from [nixpkgs](https://search.nixos.org/packages?show=gitAndTools.gh&query=gh&from=0&size=30&sort=relevance&channel=20.03#disabled):

```bash
nix-env -iA nixos.gitAndTools.gh
```

### openSUSE Tumbleweed

openSUSE Tumbleweed users can install from the [official distribution repo](https://software.opensuse.org/package/gh):
```bash
sudo zypper in gh
```

### Alpine Linux

Alpine Linux users can install from the [stable releases' community package repository](https://pkgs.alpinelinux.org/packages?name=github-cli&branch=v3.15).

```bash
apk add github-cli
```

Users wanting the latest version of the CLI without waiting to be backported into the stable release they're using should use the edge release's
community repo through this method below, without mixing packages from stable and unstable repos.[^1]

```bash
echo "@community http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories
apk add github-cli@community
```

### Void Linux
Void Linux users can install from the [official distribution repo](https://voidlinux.org/packages/?arch=x86_64&q=github-cli):

```bash
sudo xbps-install github-cli
```

[releases page]: https://github.com/cli/cli/releases/latest
[arch linux repo]: https://www.archlinux.org/packages/community/x86_64/github-cli
[arch linux aur]: https://aur.archlinux.org/packages/github-cli-git
[^1]: https://wiki.alpinelinux.org/wiki/Package_management#Repository_pinning
