$installationPath = "C:\EkisaChatbots CLI"
$binaryPath = $installationPath + "\bin"

"- Downloading latest version of EkisaChatbots CLI..."
Start-BitsTransfer -Source "https://github.com/Ekisa-Team/ekisa-chatbots-cli/releases/download/v0.1.9-beta/ec_v0.1.9-beta-windows_amd64.zip" -Destination "ec_v0.1.9-beta-windows_amd64.zip"    

"- Unpacking ZIP file..."
Expand-Archive "ec_v0.1.9-beta-windows_amd64.zip" $installationPath

# Rename file
$fileName = (Get-Item $binaryPath).FullName + "\ec-windows-amd64.exe"
$newFileName = (Get-Item $binaryPath).FullName + "\ekisa-chatbots.exe"
Rename-Item $fileName $newFileName

"- Adding CLI to the PATH..."        
# Get all path variables
$path = [System.Environment]::GetEnvironmentVariable('Path', [EnvironmentVariableTarget]::User)

# Remove installation path to avoid duplication
$path = ($path.Split(';') | Where-Object { $_ -ne $binaryPath }) -join ';'

# Set installation path in the path environment variable
[System.Environment]::SetEnvironmentVariable('Path', $path + ';' + $binaryPath, [EnvironmentVariableTarget]::User)

Write-Host "Installation completed successfully" -ForegroundColor Green
Write-Host "Refer to https://github.com/Ekisa-Team/ekisa-chatbots-cli to read the documentation." -ForegroundColor Blue

Pause