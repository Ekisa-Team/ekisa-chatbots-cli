# **EkisaChatbots CLI**

`ekisa-chatbots` es una interfaz de línea de comandos que funciona como orquestador para gestionar el flujo de procesos que se dan entre la API [Ekisa.Chatbots](https://github.com/Ekisa-Team/Ekisa.Chatbots) y el middleware [Ekisa.Chatbots.Api](https://github.com/Ekisa-Team/Ekisa.Chatbots.Api).

## **Instalación**

### Windows

Los instaladores y los binarios en las arquitecturas de 32 y 64 bits se encuentran disponibles en la [página de releases](https://github.com/Ekisa-Team/ekisa-chatbots-cli/releases/latest), en la sección de **assets**.

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
  ekisa-chatbots completion [command]

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
  ekisa-chatbots [flags]
  ekisa-chatbots [command]

Subcommands:
  completion         // Ayudas para el comando completion
  help               // Ayudas para ekisa-chatbots
  prepare            // Ayudas para el comando prepare
  upload             // Ayudas para el comando upload

Flags:
  -h, --help         // Ayudas para ekisa-chatbots

Global flags:
  -c, --config       // Ruta del archivo de configuración (con extensión)
```

#### `prepare`

```go
Usage:
  ekisa-chatbots prepare [flags]

Flags:
  -h, --help         // Ayudas para el comando prepare

Global flags:
  -c, --config       // Ruta del archivo de configuración (con extensión)
```

#### `upload`

```go
Usage:
  ekisa-chatbots upload [flags]

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
.\ekisa-chatbots-conf.yaml                                    Raiz de la carpeta donde se encuentra instalado el ejecutable
$HOME\ekisa-chatbots-conf.yaml                                C:\Users\username\ekisa-chatbots-conf
$HOME\.config\ekisa-chatbots\ekisa-chatbots-conf.yaml         C:\Users\username\.config\ekisa-chatbots
```

Opcionalmente se le puede especificar explícitamente la ruta del archivo de configuración con la bandera global `--config` al ejectuar cualquier comando transaccional.

Ejemplo:

```
ekisa-chatbots prepare --config C:\Users\juanm\Downloads

ó

ekisa-chatbots prepare -c C:\Users\juanm\Downloads
```

> Ver [plantilla](https://github.com/Ekisa-Team/ekisa-chatbots-cli/blob/main/config.yaml) del archivo de configuración

## **Otro**

- Doumentación de [Ekisa.Chatbots](https://github.com/Ekisa-Team/Ekisa.Chatbots)
- Documentación de [Ekisa.Chatbots.Middleware](https://github.com/Ekisa-Team/Ekisa.Chatbots.Api)
- Documentación de [Ekisa.Chatbots.Listener](https://github.com/Ekisa-Team/Ekisa.Chatbots.Listener)
