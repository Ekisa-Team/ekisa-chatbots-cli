@echo off

title Install EkisaChatbots CLI

goto download_latests_release

:: Utilities
:S
%Windir%\System32\WindowsPowerShell\v1.0\Powershell.exe write-host -foregroundcolor Green %1
goto:eof

:I
%Windir%\System32\WindowsPowerShell\v1.0\Powershell.exe write-host -foregroundcolor Blue %1
goto:eof

:W
%Windir%\System32\WindowsPowerShell\v1.0\Powershell.exe write-host -foregroundcolor Yellow %1
goto:eof

:E
%Windir%\System32\WindowsPowerShell\v1.0\Powershell.exe write-host -foregroundcolor Red %1
goto:eof

:: Steps
:download_latests_release
    call:I "Downloading latest version of EkisaChatbots CLI..."
    powershell -command "Start-BitsTransfer -Source https://github.com/Ekisa-Team/ekisa-chatbots-cli/releases/download/v0.1.9-beta/ec_v0.1.9-beta-windows_amd64.zip -Destination ec_v0.1.9-beta-windows_amd64.zip"    

    call:I "Unpacking ZIP file..."
    powershell -command "Expand-Archive ec_v0.1.9-beta-windows_amd64.zip C:\Tools

:add_cli_to_path
    call:I "Adding CLI to the PATH..."    
    SET PATH=%PATH%;C:\Tools\

    call:S "Installation completed successfully"    
    echo Refer to https://github.com/Ekisa-Team/ekisa-chatbots-cli to read the documentation.

:end
    cmd.exe /k cmd /c