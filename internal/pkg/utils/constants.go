package utils

type contextKey string

const (
	ContentType   string     = "Content-Type"
	ServerAddrKey contextKey = "serverAddr"
)

const (
	GET     uint16 = 0x01
	POST    uint16 = 0x02
	PUT     uint16 = 0x04
	DELETE  uint16 = 0x08
	HEAD    uint16 = 0x10
	PATCH   uint16 = 0x20
	CONNECT uint16 = 0x40
	OPTIONS uint16 = 0x80
	TRACE   uint16 = 0x100
)
