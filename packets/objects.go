package packets

import (
	"github.com/google/uuid"
	"time"
)

type (
	Dumper struct {
		Bytes       []byte
		HttpDump    string
		BodyHexDump string
	}
	Decoder struct {
		UnitTest  string
		Steam     string
		ProtoBuf  string
		Tdf       string
		Taf       string
		Acc       string
		Websocket string
		Msgpack   string
	}
	Body struct {
		Dumper
		Decoder
	}
	Row struct {
		Method        string        `json:"method"`
		Scheme        string        `json:"scheme"`
		Host          string        `json:"host"`
		Path          string        `json:"path"`
		ContentType   string        `json:"type"`
		ContentLength int64         `json:"length"`
		Status        string        `json:"status"`
		Note          string        `json:"note"`
		Process       string        `json:"process"`
		PadTime       time.Duration `json:"pad"`
	}
	Expand struct {
		UUID        uuid.UUID
		IsWebsocket bool
		IsUdp       bool
		IsTcp       bool
		IsRequest   bool
		IsResponse  bool
	}
	Object struct {
		Row
		Expand
		Req  Body
		Resp Body
	}
)
