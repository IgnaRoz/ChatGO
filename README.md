# Proyecto de Chat WebSocket

Este proyecto es una implementación de un servidor de chat utilizando WebSockets en Go. Permite a múltiples usuarios conectarse y enviar mensajes en tiempo real.

## Descripción

El servidor de chat está construido utilizando el paquete `net/http` de Go y la biblioteca `gorilla/websocket` para manejar las conexiones WebSocket. Los usuarios pueden conectarse al servidor, enviar mensajes y recibir mensajes de otros usuarios en tiempo real.

### Funcionalidades

- Conexión de múltiples usuarios a través de WebSockets.
- Envío y recepción de mensajes en tiempo real.
- Gestión de usuarios y mensajes en un grupo de chat.(En proceso)

## Uso

### Requisitos

- Go 1.16 o superior
- Biblioteca `gorilla/websocket`

### Instalación

1. Clona el repositorio:
    ```sh
    git clone https://github.com/tu_usuario/tu_repositorio.git
    cd tu_repositorio
    ```

2. Instala las dependencias:
    ```sh
    go get -u github.com/gorilla/websocket
    ```

### Ejecución

1. Compila y ejecuta el servidor:
    ```sh
    go run main.go
    ```

2. Abre tu navegador y conéctate al servidor:
    ```sh
    http://localhost:8080
    ```

### Uso del Chat

1. Abre múltiples pestañas o ventanas del navegador y conéctate al servidor.
2. Escribe un mensaje en una pestaña y observa cómo aparece en las otras pestañas conectadas.

¡Y eso es todo! Ahora tienes un servidor de chat funcional utilizando WebSockets en Go.

