package types

// CRLF is the end of text line
var CRLF = "\r\n"

// HeaderKeyPUBLISH header key
var (
	HeaderKeyPublish     = "publish"
	HeaderKeySubscribe   = "subscribe"
	HeaderKeyUnsubscribe = "unsubscribe"
	HeaderKeyUpstream    = "upstream"
	HeaderKeyMultipart   = "multipart"
)

// HeaderMax
var (
	//headers max num not size
	HeaderMax               = 20
	SubscribeMax            = 40
	Version                 = "1.0"
	ProtocolName            = "LHTTP"
	ProtocolNameWithVersion = "LHTTP/1.0"
	ProtocolLength          = 9
	MaxLength               = 40960
)

// UpstreamHTTPMethodGet
var (
	UpstreamHTTPMethodGET  = "GET"
	UpstreamHTTPMethodPOST = "POST"
)

// ProcessorMax
var (
	ProcessorMax = 40
)
