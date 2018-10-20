package wsim

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/Eric-GreenComb/ws-im-server/types"
)

type upstreamHeadFilter struct {
	*HeadFilterBase
}

func (*upstreamHeadFilter) AfterRequestFilterHandle(ws *WsHandler) {
	var value string
	if value = ws.GetHeader(types.HeaderKeyUpstream); value == "" {
		log.Print("no upstream header found:", ws.message.message, ws.message.headers)
		return
	}

	u := &url.URL{}

	values := strings.Split(value, " ")

	ws.upstreamURL, _ = u.Parse(values[1])

	log.Print("upstream method:", values[0], "url: ", ws.upstreamURL.String())

	httpClient := &http.Client{}
	_ = httpClient

	var req *http.Request
	var err error

	if values[0] == types.UpstreamHTTPMethodGET {
		req, err = http.NewRequest(types.UpstreamHTTPMethodGET, ws.upstreamURL.String(), nil)
		if err != nil {
			_ = req
			return
		}
	} else {
		req, err = http.NewRequest(values[0], ws.upstreamURL.String(), strings.NewReader(ws.GetBody()))
		if err != nil {
			_ = req
			return
		}
	}

	for k, v := range ws.message.headers {
		req.Header.Set(k, v)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	log.Print(string(body))
}
