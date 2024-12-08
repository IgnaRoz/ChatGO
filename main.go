package main

import (
	"chat/chat"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var grupo chat.Grupo = chat.Grupo{}

type UserChat struct {
	Name string
	chat.Sender
	conn *websocket.Conn
}

func (u UserChat) Send(m chat.Mensaje) {
	fmt.Printf("Enviando al usuario %s el mensaje \"%s\"\n", u.Name, m)
	err := u.conn.WriteMessage(websocket.TextMessage, []byte(m.Mensaje))
	if err != nil {
		fmt.Println(err)
	}

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Permitir conexiones desde cualquier origen
	},
}

func wbHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Nueva conexión a websocket")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Crear usuario,  con nombre en blanco

	usuario := UserChat{Name: "", conn: conn}
	grupo.AddParticipant(usuario)
	for {
		// Leer mensaje del cliente
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Mensaje recibido: %s\n", msg)
		grupo.NewMessage(chat.Mensaje{From: usuario.Name, Mensaje: string(msg)}, usuario)

	}

}

func main() {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/ws", wbHandlerTest)
	//mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))
	//
	////Iniciar servidor
	//http.ListenAndServe(":8080", mux)
	http.HandleFunc("/ws", wbHandler)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)

}
