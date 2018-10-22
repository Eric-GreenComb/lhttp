package httphandle

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Eric-GreenComb/ws-im-server/mq"
	"github.com/Eric-GreenComb/ws-im-server/types"
	"github.com/Eric-GreenComb/ws-im-server/wsim"
)

type httpPublisher struct{}

func (*httpPublisher) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print("an error occur")
	}
	log.Print("http publish body: ", string(body))

	bodyStr := string(body)

	message := wsim.BuildMessage(bodyStr)

	channels, ok := message.Headers[types.HeaderKeyPublish]
	if !ok {
		log.Print("cant get Publish header")
		return
	}

	for _, c := range strings.Split(channels, " ") {
		mq.MQD.Publish(c, bodyStr)
	}

	req.Body.Close()
}

// Init Init
func Init() {

}

func init() {

	//handle http publish message
	http.Handle("/publish", &httpPublisher{})
	http.HandleFunc("/health", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "UP")
}
