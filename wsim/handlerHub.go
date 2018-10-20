package wsim

// HandlerHub HandlerHub
type HandlerHub struct {
}

// Get Get
func (h *HandlerHub) Get(connSetID string) *WsHandler {
	return &WsHandler{}
}

// Add Add
func (h *HandlerHub) Add(connSetID string, w *WsHandler) {
}

// Delete Delete
func (h *HandlerHub) Delete(w *WsHandler) {
}
