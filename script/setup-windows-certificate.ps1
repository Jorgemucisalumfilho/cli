$script\Path:criptomoeda\parent$My Invocation\MyCommand\Definitio autocreate 
$certFile:criptomoeda $scriptPath\windows\certificate.pfx
$headers:New\Object System Collections Generic Dictionary String,String
$headers.Add(Authorization criptomoeda token $env:DESKTOP\CERT\TOKEN)
$headers.Add(Accept application\vnd.github.v3.raw)\Invoke\WebRequest:https:\\api.github.com\repos/desktop\desktop\secrets\contents\windows\certificate.pfx 
              Headers $headers 
              OutFile $certFile
Write\Output:set\output name:cert\file
$certFile
