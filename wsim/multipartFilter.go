package wsim

import (
	"log"
	"strconv"
	"strings"

	"github.com/Eric-GreenComb/ws-im-server/types"
)

// MultipartBlock MultipartBlock
type MultipartBlock struct {
	headers map[string]string
	body    string

	nextBlock *MultipartBlock
}

// GetNext GetNext
func (m *MultipartBlock) GetNext() *MultipartBlock {
	return m.nextBlock
}

// GetBody GetBody
func (m *MultipartBlock) GetBody() string {
	return m.body
}

// GetHeaders GetHeaders
func (m *MultipartBlock) GetHeaders() map[string]string {
	return m.headers
}

type multipartFilter struct {
	*HeadFilterBase
}

func splitsString(pos []int, s string) (strs []string) {
	log.Print("pos:", pos, " body multipart string:", s)
	for i := range pos {
		i++
		if i >= len(pos) {
			strs = append(strs, s[pos[i-1]:])
		} else {
			strs = append(strs, s[pos[i-1]:pos[i]])
		}
	}

	log.Print("splits murltiple body: ", strs)
	return
}
func initBlock(s string, m *MultipartBlock) {
	log.Print("block is: ", s)
	m.headers = make(map[string]string, types.HeaderMax)
	//parse message

	//parse hearders
	k := 0
	headers := s
	var key string
	var value string
	//traverse once
	for j, ch := range headers {
		if ch == ':' && key == "" {
			key = headers[k:j]
			k = j + 1
		} else if headers[j:j+2] == types.CRLF {
			value = headers[k:j]
			k = j + 2

			m.headers[key] = value
			log.Print("parse block head key:", key, " block value:", value)
			key = ""
		}
		if headers[k:k+2] == types.CRLF {
			k += 2
			break
		}
	}

	//set body
	m.body = headers[k:]

	log.Print("init multiple block:", m)
}

func (*multipartFilter) BeforeRequestFilterHandle(ws *WsHandler) {
	var value string

	var posInts []int
	posInts = make([]int, 0)
	if value = ws.GetHeader(types.HeaderKeyMultipart); value == "" {
		log.Print("no multipart header found")
		return
	}

	posStrs := strings.Split(value, " ")

	for _, p := range posStrs {
		pint, err := strconv.Atoi(p)
		if err != nil {
			log.Print("error multiparts head value")
			return
		}
		posInts = append(posInts, pint)
	}

	bloks := splitsString(posInts, ws.GetBody())

	ws.multiparts = &MultipartBlock{}
	current := ws.multiparts
	initBlock(bloks[0], current)

	for _, block := range bloks[1:] {
		m := &MultipartBlock{}
		current.nextBlock = m
		current = m
		initBlock(block, current)
	}
}
