package wsim

import (
	"log"
	"strings"

	"github.com/Eric-GreenComb/ws-im-server/mq"
	"github.com/Eric-GreenComb/ws-im-server/types"
)

//if client send message include subscribe/publish/unsubscribe header
//this filter work,use nats as a message queue client
type mqHeadFilter struct {
	*HeadFilterBase
}

func (*mqHeadFilter) AfterRequestFilterHandle(ws *WsHandler) {
	var value string
	var channels []string

	if value = ws.GetHeader(types.HeaderKeySubscribe); value != "" {
		channels = strings.Split(value, " ")
		for _, c := range channels {
			if conn, err := mq.MQD.Subscribe(c, ws.subscribeCallback); nil == err {
				ws.SubscribeNatsConn[c] = conn
			} else {
				log.Println("Subscribe Error", err)
			}
			// log.Print("subscribe channel: ", c)
		}
	}

	if value = ws.GetHeader(types.HeaderKeyPublish); value != "" {
		channels = strings.Split(value, " ")
		for _, c := range channels {
			ws.setResponse()
			ws.resp.serializeMessage()
			mq.MQD.Publish(c, ws.resp.message)
			// log.Print("publish channel: ", c, "message:", ws.resp.message)
		}
	}

	if value = ws.GetHeader(types.HeaderKeyUnsubscribe); value != "" {
		channels = strings.Split(value, " ")
		for _, c := range channels {
			mq.MQD.Unsubscribe(ws.SubscribeNatsConn[c])
			log.Print("unsubscribe channel: ", c)
		}
	}
}
