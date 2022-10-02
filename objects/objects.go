package objects

import (
	"github.com/google/uuid"
	"time"
)

type (
	Body struct {
		Bytes       []byte
		HttpDump    string
		BodyHexDump string
		Steam       string
		Head        string
		Pb2         string
		Pb3         string
		Tdf         string
		Taf         string
		Acc         string
		Text        string
		Json        string
		Html        string
		Javascript  string
		Websocket   string
		Msgpack     string
		Gzip        string
	}
	Row struct {
		Index         uint64
		Method        string
		Scheme        string
		Url           string
		ContentType   string
		ContentLength int64
		Status        string
		Note          string
		StartTime     time.Time
		PadTime       time.Duration
	}
	Expand struct {
		UUID        uuid.UUID
		IsWebsocket bool
		IsUdp       bool
		IsTcp       bool
		IsRequest   bool
		IsResponse  bool
	}
	Packet struct {
		Row
		Expand
		Req  Body
		Resp Body
	}
)
