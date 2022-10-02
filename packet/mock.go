package packet

import (
	"github.com/ddkwork/unisonUi/objects"
	"github.com/google/uuid"
	"net/http"
	"sync/atomic"
	"time"
)

func (o *object) mitmMock() {
	index := uint64(0)
	go func() {
		for i := 0; i < 3; i++ {
			atomic.AddUint64(&index, 1)
			packet := objects.Packet{
				Row: objects.Row{
					Index:         index,
					Method:        http.MethodConnect,
					Scheme:        "https",
					Url:           "https://github.com/ddkwork/goproxyeeeeeeeee",
					ContentType:   "protobuf",
					ContentLength: 2320,
					Status:        "200 OK",
					Note:          "http/s message",
					StartTime:     time.Now(),
					PadTime:       2,
				},
				Expand: objects.Expand{
					UUID:        uuid.New(),
					IsWebsocket: false,
					IsRequest:   false,
					IsResponse:  false,
				},
				Req:  objects.Body{},
				Resp: objects.Body{},
			}
			if i%2 == 1 {
				packet.IsWebsocket = true
			}
			o.AddRow(packet)
			time.Sleep(time.Second)
		}
	}() //https
	go func() {
		for i := 0; i < 3; i++ {
			atomic.AddUint64(&index, 1)
			packet := objects.Packet{
				Row: objects.Row{
					Index:         index,
					Method:        http.MethodGet,
					Scheme:        "wss",
					Url:           "https://github.com/ddkwork/goproxyrrrr",
					ContentType:   "json",
					ContentLength: 2220,
					Status:        "302 ssOK",
					Note:          "Websocket message",
					StartTime:     time.Now(),
					PadTime:       2,
				},
				Expand: objects.Expand{
					UUID:        uuid.New(),
					IsWebsocket: true,
					IsUdp:       false,
					IsTcp:       false,
					IsRequest:   false,
					IsResponse:  false,
				},
				Req:  objects.Body{},
				Resp: objects.Body{},
			}
			if i%2 == 1 {
				packet.IsWebsocket = true
			}
			o.AddRow(packet)
			time.Sleep(time.Second)
		}
	}() //wss
	go func() {
		for i := 0; i < 3; i++ {
			atomic.AddUint64(&index, 1)
			packet := objects.Packet{
				Row: objects.Row{
					Index:         index,
					Method:        http.MethodPost,
					Scheme:        "tcp",
					Url:           "https://github.com/ddkwork/goproxyssssssssssssssssssssss",
					ContentType:   "html",
					ContentLength: 2210,
					Status:        "503 sssOK",
					Note:          "tcp message",
					StartTime:     time.Now(),
					PadTime:       2,
				},
				Expand: objects.Expand{
					UUID:        uuid.New(),
					IsWebsocket: false,
					IsUdp:       false,
					IsTcp:       true,
					IsRequest:   false,
					IsResponse:  false,
				},
				Req:  objects.Body{},
				Resp: objects.Body{},
			}
			if i%2 == 1 {
				packet.IsWebsocket = true
			}
			o.AddRow(packet)
			time.Sleep(time.Second)
		}
	}() //tcp
	go func() {
		for i := 0; i < 3; i++ {
			atomic.AddUint64(&index, 1)
			packet := objects.Packet{
				Row: objects.Row{
					Index:         index,
					Method:        http.MethodOptions,
					Scheme:        "udp",
					Url:           "https://github.com/ddkwork/goproxyzzzzzzzzz",
					ContentType:   "binary",
					ContentLength: 1220,
					Status:        "444 Status",
					Note:          "udp message",
					StartTime:     time.Now(),
					PadTime:       2,
				},
				Expand: objects.Expand{
					UUID:        uuid.New(),
					IsWebsocket: false,
					IsUdp:       true,
					IsTcp:       false,
					IsRequest:   false,
					IsResponse:  false,
				},
				Req:  objects.Body{},
				Resp: objects.Body{},
			}
			if i%2 == 1 {
				packet.IsWebsocket = true
			}
			o.AddRow(packet)
			time.Sleep(time.Second)
		}
	}() //udp
}
