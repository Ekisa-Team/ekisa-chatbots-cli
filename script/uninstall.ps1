$installationPath = "C:\EkisaChatbots CLI"
$binaryPath = $installationPath + "\bin"

$testPath = Test-Path -Path $binaryPath

if ($testPath) {
    "Uninstalling EkisaChatbots CLI..."
    
    # Remove bin file
    Remove-Item $installationPath -Recurse
    
    # Get all path variables
    $path = [System.Environment]::GetEnvironmentVariable('Path', [EnvironmentVariableTarget]::User)
    
    # Remove installation path variable
    $path = ($path.Split(';') | Where-Object { $_ -ne $binaryPath }) -join ';'
        
    # Update path environment variable
    [System.Environment]::SetEnvironmentVariable('Path', $path, [EnvironmentVariableTarget]::User)

    Write-Host "EkisaChatbots CLI was uninstalled successfully" -ForegroundColor Green
}
else {
    Write-Host "EkisaChatbots CLI is not installed." -ForegroundColor Yellow
}

Pause