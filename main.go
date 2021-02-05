package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

// MyRes ...
type MyRes struct{
	Success bool `json:"success"`
}

func sayHello(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	myData := MyRes{Success: true}
	json.NewEncoder(w).Encode(myData)
}

func main(){
	socket := socketio.NewServer(nil)
	socket.OnConnect("/", func(conn socketio.Conn) error {
		conn.SetContext("my context")
		fmt.Println("Socket is here", conn.ID())
		return nil
	})

	socket.OnEvent("/", "handshake", func(conn socketio.Conn) string{
		data := conn.Context()
		conn.Emit("cool", data)
		conn.Close()
		return "ok"
	})

	http.HandleFunc("/", sayHello)

	PORT := "5000"
	fmt.Println("Server starting on port", PORT)
	log.Fatal(http.ListenAndServe(":" + PORT, nil))
}