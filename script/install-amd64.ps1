Write-Output "Initializing..."

# Check powershell version
if (($PSVersionTable.PSVersion.Major) -lt 5) {
    Write-Output "PowerShell 5 or later is required to run Kibot CLI installer"
    Write-Output "Upgrade PowerShell: https://docs.microsoft.com/en-us/powershell/scripting/setup/installing-windows-powershell"
    Pause
    break
}

$installationPath = "C:\Kibot CLI"
$binaryPath = $installationPath + "\bin"

# Check if it's already installed
$testPath = Test-Path -Path $binaryPath

if ($testPath) {
    Write-Host "Kibot CLI is already installed on this computer. If you wish to install this version, please uninstall first." -f Yellow
    Pause
    break    
}

# Download binary
Write-Output "Downloading latest version of Kibot CLI..."
Start-BitsTransfer -Source "https://github.com/Ekisa-Team/kibot-cli/releases/latest/download/kibot-windows-amd64.zip" -Destination "kibot-windows-amd64.zip"    

# Extract ZIP file
Write-Output "Extracting..."
Expand-Archive "kibot-windows-amd64.zip" $installationPath
Remove-Item "kibot-windows-amd64.zip" -Force

# Rename file
$fileName = (Get-Item $binaryPath).FullName + "\kibot-windows-amd64.exe"
$newFileName = (Get-Item $binaryPath).FullName + "\kibot.exe"
Rename-Item $fileName $newFileName

Write-Output "Adding CLI to the PATH..."        
# Get all path variables
$path = [System.Environment]::GetEnvironmentVariable('Path', [EnvironmentVariableTarget]::User)

# Remove installation path to avoid duplication
$path = ($path.Split(';') | Where-Object { $_ -ne $binaryPath }) -join ';'

# Set installation path in the path environment variable
[System.Environment]::SetEnvironmentVariable('Path', $path + ';' + $binaryPath, [EnvironmentVariableTarget]::User)

Write-Host "Kibot CLI was installed successfully!" -f darkgreen

Pause
