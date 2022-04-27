@echo off

title Install EkisaChatbots CLI

goto download_latests_release

:: Utilities
:ECHOSUCCESS
%Windir%\System32\WindowsPowerShell\v1.0\Powershell.exe write-host -foregroundcolor Green %1
goto:eof

:ECHOINFO
%Windir%\System32\WindowsPowerShell\v1.0\Powershell.exe write-host -foregroundcolor Blue %1
goto:eof

:ECHOWARNING
%Windir%\System32\WindowsPowerShell\v1.0\Powershell.exe write-host -foregroundcolor Yellow %1
goto:eof

:ECHODANGER
%Windir%\System32\WindowsPowerShell\v1.0\Powershell.exe write-host -foregroundcolor Red %1
goto:eof

:: Steps
:download_latests_release
    call:ECHOINFO "Downloading EkisaChatbots CLI from https://github.com/Ekisa-Team/ekisa-chatbots-cli/releases/tag/latest ..."

    powershell -command "Start-BitsTransfer -Source https://github.com/Ekisa-Team/ekisa-chatbots-cli/releases/download/v0.1.9-beta/ec_v0.1.9-beta-windows_386.zip -Destination ec_v0.1.9-beta-windows_386.zip"
    powershell -command "Expand-Archive ec_v0.1.9-beta-windows_386.zip C:\Tools
:end
    cmd.exe /k cmd /c