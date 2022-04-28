Write-Output "Initializing..."

# Check powershell version
if (($PSVersionTable.PSVersion.Major) -lt 5) {
    Write-Output "PowerShell 5 or later is required to run Kibot CLI"
    Write-Output "Upgrade PowerShell: https://docs.microsoft.com/en-us/powershell/scripting/setup/installing-windows-powershell"
    break
}

$installationPath = "C:\Kibot CLI"
$binaryPath = $installationPath + "\bin"

# Check if it's already installed
$testPath = Test-Path -Path $binaryPath

if ($testPath) {
    "Uninstalling Kibot CLI..."
    
    # Remove bin file
    Remove-Item $installationPath -Recurse -Force
    
    # Get all path variables
    $path = [System.Environment]::GetEnvironmentVariable('Path', [EnvironmentVariableTarget]::User)
    
    # Remove installation path variable
    $path = ($path.Split(';') | Where-Object { $_ -ne $binaryPath }) -join ';'
        
    # Update path environment variable
    [System.Environment]::SetEnvironmentVariable('Path', $path, [EnvironmentVariableTarget]::User)

    Write-Host "Kibot CLI was uninstalled successfully!" -f darkgreen
}
else {
    Write-Host "Kibot CLI is not installed." -f Yellow
}

Pause