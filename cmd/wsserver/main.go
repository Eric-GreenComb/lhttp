package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Eric-GreenComb/ws-im-server/wsim"
)

//ChatProcessor is
type ChatProcessor struct {
	*wsim.BaseProcessor
}

//OnMessage is
func (p *ChatProcessor) OnMessage(h *wsim.WsHandler) {
	log.Print("on OnMessage: ", h.GetBody())
	h.AddHeader("content-type", "image/png")
	h.SetCommand("auth")
	h.Send(h.GetBody())
}

//SubPubProcessor is
type SubPubProcessor struct {
	*wsim.BaseProcessor
}

//UpstreamProcessor is
type UpstreamProcessor struct {
	*wsim.BaseProcessor
}

//UploadProcessor is
type UploadProcessor struct {
	*wsim.BaseProcessor
}

//OnMessage is
func (*UploadProcessor) OnMessage(ws *wsim.WsHandler) {
	for m := ws.GetMultipart(); m != nil; m = m.GetNext() {
		log.Print("multibody:", m.GetBody(), " headers:", m.GetHeaders())
	}
}

func main() {
	wsim.Regist("chat", &ChatProcessor{&wsim.BaseProcessor{}})
	wsim.Regist("subpub", &SubPubProcessor{&wsim.BaseProcessor{}})
	wsim.Regist("upstream", &UpstreamProcessor{&wsim.BaseProcessor{}})
	wsim.Regist("upload", &UploadProcessor{&wsim.BaseProcessor{}})

	http.Handle("/echo", wsim.Handler(wsim.StartServer))
	http.Handle("/", wsim.Handler(wsim.StartServer))
	http.HandleFunc("/https", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", "world")
	})
	http.ListenAndServe(":8081", nil)
}
