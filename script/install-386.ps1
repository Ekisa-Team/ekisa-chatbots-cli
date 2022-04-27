Write-Output "Initializing..."

# Check powershell version
if (($PSVersionTable.PSVersion.Major) -lt 5) {
    Write-Output "PowerShell 5 or later is required to run EkisaChatbots CLI"
    Write-Output "Upgrade PowerShell: https://docs.microsoft.com/en-us/powershell/scripting/setup/installing-windows-powershell"
    Pause
    break
}

$installationPath = "C:\EkisaChatbots CLI"
$binaryPath = $installationPath + "\bin"

# Check if it's already installed
$testPath = Test-Path -Path $binaryPath

if ($testPath) {
    Write-Host "EkisaChatbots CLI is already installed on this computer. If you wish to install this version, please uninstall first." -f Yellow
    Pause
    break    
}

# Download binary
Write-Output "Downloading latest version of EkisaChatbots CLI..."
Start-BitsTransfer -Source "https://github.com/Ekisa-Team/ekisa-chatbots-cli/releases/latest/download/ec-windows-386.zip" -Destination "ec-windows-386.zip"    

# Extract ZIP file
Write-Output "Extracting..."
Expand-Archive "ec-windows-386.zip" $installationPath
Remove-Item "ec-windows-386.zip" -Force

# Rename file
$fileName = (Get-Item $binaryPath).FullName + "\ec-windows-386.exe"
$newFileName = (Get-Item $binaryPath).FullName + "\ekisa-chatbots.exe"
Rename-Item $fileName $newFileName

Write-Output "Adding CLI to the PATH..."        
# Get all path variables
$path = [System.Environment]::GetEnvironmentVariable('Path', [EnvironmentVariableTarget]::User)

# Remove installation path to avoid duplication
$path = ($path.Split(';') | Where-Object { $_ -ne $binaryPath }) -join ';'

# Set installation path in the path environment variable
[System.Environment]::SetEnvironmentVariable('Path', $path + ';' + $binaryPath, [EnvironmentVariableTarget]::User)

Write-Host "EkisaChatbots CLI was installed successfully!" -f darkgreen

Pause
