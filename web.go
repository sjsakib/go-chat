package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type socket struct {
	io.Reader
	io.Writer
	done chan bool
}

func (s socket) Close() error {
	s.done <- true
	return nil
}

func socketHandler(ws *websocket.Conn) {
	r := feedChain(ws)
	s := socket{r, ws, make(chan bool)}
	go match(s)
	<-s.done
}


const webListenAddr = "localhost:3000"

func runWeb() {
	http.HandleFunc("/", rootHandler)
	http.Handle("/socket", websocket.Handler(socketHandler))
	err := http.ListenAndServe(webListenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	err := rootTemplate.Execute(w, webListenAddr)
	if err != nil {
		log.Fatal(err)
	}
}

var rootTemplate = template.Must(template.New("root").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
</head>
<body>
    <h1>Go Chat</h1>
    <form id="form">
        <input id="input" type="text" />
        <button type="submit">Send</button>
        <ul id="messages"></ul>
    </form>
</body>
<script>
    document.getElementById("form").addEventListener("submit", onSend);
    function onSend(event) {
        event.preventDefault();
        websocket.send(input.value + '\n');
        appendMessage(input.value);
        input.value = '';
    }
    function appendMessage(data) {
        var messages = document.getElementById("messages");
        var message = document.createElement("li");
        console.log({data});
        message.textContent = data;
        messages.appendChild(message);
    }
    function onMessage(event) {
        console.log({event});
        appendMessage(event.data);
    }
    function onClose(event) {
        alert("Connection closed");
    }
    websocket = new WebSocket("ws://{{.}}/socket");
    websocket.onmessage = onMessage;
    websocket.onclose = onClose;
</script>
</html>
`))
