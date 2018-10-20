package wsim

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Eric-GreenComb/ws-im-server/mq"
	"github.com/Eric-GreenComb/ws-im-server/types"
)

type httpPublisher struct{}

func (*httpPublisher) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print("an error occur")
	}
	log.Print("http publish body: ", string(body))

	bodyStr := string(body)

	message := buildMessage(bodyStr)

	channels, ok := message.headers[types.HeaderKeyPublish]
	if !ok {
		log.Print("cant get Publish header")
		return
	}

	for _, c := range strings.Split(channels, " ") {
		mq.MQD.Publish(c, bodyStr)
	}

	req.Body.Close()
}

func init() {

	//handle http publish message
	http.Handle("/publish", &httpPublisher{})
}
