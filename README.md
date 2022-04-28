# **Kibot CLI**

`kibot` es una interfaz de línea de comandos que funciona como orquestador para gestionar el flujo de procesos que se dan entre [Kibot](https://github.com/Ekisa-Team/Kibot) y el middleware [Kibot.Quiron.Middleware](https://github.com/Ekisa-Team/Kibot.Quiron.Middleware).

## **Instalación**

Ejecutar el siguiente comando desde `PowerShell` para instalar **Kibot CLI** en su ubicación por defecto (`c:\Kibot CLI`)

**Windows 64 bits**

```ps1
Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/Ekisa-Team/kibot-cli/main/script/install-amd64.ps1')

# ó

iwr -useb https://raw.githubusercontent.com/Ekisa-Team/kibot-cli/main/script/install-amd64.ps1 | iex
```

**Windows 32 bits**

```ps1
Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/Ekisa-Team/kibot-cli/main/script/install-386.ps1')

# ó

iwr -useb https://raw.githubusercontent.com/Ekisa-Team/kibot-cli/main/script/install-386.ps1 | iex
```

Una vez instalado, ejecutar `kibot help` para recibir instrucciones y verificar que la instalación fue exitosa.

> En caso de que el binario instalado sea el de 32 bits, se debe utilizar el comando `kibot32` en lugar de `kibot`

Adicionalmente los binarios en las arquitecturas de 32 y 64 bits se encuentran disponibles en la [página de releases](https://github.com/Ekisa-Team/kibot-cli/releases/latest), en la sección de **assets**.

## **Desinstalación**

Ejecutar el siguiente comando desde `PowerShell` para desinstalar **Kibot CLI**

```ps1
Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/Ekisa-Team/kibot-cli/main/script/uninstall.ps1')

# o más corto
iwr -useb https://raw.githubusercontent.com/Ekisa-Team/kibot-cli/main/script/uninstall.ps1 | iex
```

## **Documentación**

### `Comandos disponibles`

```go
completion      // Genera script de autocompletado para un shell especificado

help            // Ayudas acerca de cualquier comando

prepare         // Carga las citas en la tabla ChatbotCitas antes de ser subidas

upload          // Toma las citas de la tabla ChatbotCitas y las sube a la nube
```

#### `completion`

```go
Usage:
  kibot completion [command]

Subcommands:
  bash          // Genera script de autocompletado para bash
  fish          // Genera script de autocompletado para fish
  powershell    // Genera script de autocompletado para powershell
  zsh           // Genera script de autocompletado para zsh

Flags:
  -h, --help    // Ayudas para el comando completion
```

#### `help`

```go
Usage:
  kibot [flags]
  kibot [command]

Subcommands:
  completion         // Ayudas para el comando completion
  help               // Ayudas para kibot
  prepare            // Ayudas para el comando prepare
  upload             // Ayudas para el comando upload

Flags:
  -h, --help         // Ayudas para kibot

Global flags:
  -c, --config       // Ruta del archivo de configuración (con extensión)
```

#### `prepare`

```go
Usage:
  kibot appointment prepare [flags]

Flags:
  -h, --help         // Ayudas para el comando prepare

Global flags:
  -c, --config       // Ruta del archivo de configuración (con extensión)
```

#### `upload`

```go
Usage:
  kibot appointment upload [flags]

Flags:
  -h, --help         // Ayudas para el comando upload

Global flags:
  -c, --config       // Ruta del archivo de configuración (con extensión)
```

### **Configuración**

La configuración del CLI se maneja a través de un archivo [YAML](https://es.wikipedia.org/wiki/YAML).

El archivo recibe los siguientes parámetros:

```yaml
# Application
client: 32

# Database
connection_string: Data Source=MY_DATASOURCE;Initial Catalog=MY_DATABASE;Integrated Security=True

# Webhooks
upload_webhook_uri: https://ekisa-chatbots-api.azurewebsites.net/api/chatbotcita/create
```

El CLI buscará ese archivo en las siguientes ubicaciones:

```shell
.\kibot-config.yaml                                    Raiz de la carpeta donde se encuentra instalado el ejecutable
$HOME\.config\kibot\kibot-config.yaml                  C:\Users\username\.config\kibot\kibot-config.yaml
```

Opcionalmente se le puede especificar explícitamente la ruta del archivo de configuración con la bandera global `--config` al ejectuar cualquier comando transaccional.

Ejemplo:

```
kibot appointment prepare --config C:\Users\username\Downloads\kibot-config.yaml

ó

kibot appointment prepare -c C:\Users\username\Downloads\kibot-config.yaml
```

> Ver [plantilla](https://github.com/Ekisa-Team/kibot-cli/blob/main/config.yaml) del archivo de configuración

## **Otros**

- Doumentación de [Kibot](https://github.com/Ekisa-Team/Kibot)
- Documentación de [Kibot.Quiron.Middleware](https://github.com/Ekisa-Team/Kibot.Quiron.Middleware)
- Documentación de [Kibot.Quiron.Listener](https://github.com/Ekisa-Team/Kibot.Quiron.Listener)
