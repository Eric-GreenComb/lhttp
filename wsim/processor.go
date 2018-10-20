package wsim

import (
	"github.com/Eric-GreenComb/ws-im-server/types"
)

var processorMap map[string]HandlerCallbacks

// Regist Regist
func Regist(command string, p HandlerCallbacks) {
	processorMap[command] = p
	//log.Print("regist processor", p)
}

func getProcessor(command string) HandlerCallbacks {
	//log.Print("get processor:", processorMap[command], " command:", command)
	p, ok := processorMap[command]
	if ok {
		return p
	}
	return nil
}

func init() {
	processorMap = make(map[string]HandlerCallbacks, types.ProcessorMax)
}
