package types

// WSGUID const var
const (
	// WSGUID winsocket const GUID
	WSGUID = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

	CloseStatusNormal            = 1000
	CloseStatusGoingAway         = 1001
	CloseStatusProtocolError     = 1002
	CloseStatusUnsupportedData   = 1003
	CloseStatusFrameTooLarge     = 1004
	CloseStatusNoStatusRcvd      = 1005
	CloseStatusAbnormalClosure   = 1006
	CloseStatusBadMessageData    = 1007
	CloseStatusPolicyViolation   = 1008
	CloseStatusTooBigData        = 1009
	CloseStatusExtensionMismatch = 1010

	MaxControlFramePayloadLength = 125
)
