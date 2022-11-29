package main

import (
	"github.com/ddkwork/unisonUi/packets"
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
			packet := packets.Object{
				Row: packets.Row{
					Method:        http.MethodConnect,
					Scheme:        "https",
					Host:          "",
					Path:          "",
					ContentType:   "protobuf",
					ContentLength: 2320,
					Status:        "200 OK",
					Note:          "http/s message",
					Process:       "",
					PadTime:       2,
				},
				Expand: packets.Expand{},
				Req:    packets.Body{},
				Resp:   packets.Body{},
			}
			if i%2 == 1 {
				packet.IsWebsocket = true
			}
			//panel.InstallCmdHandlers(i, unison.AlwaysEnabled,
			//	func(_ any) { creator.CreateItem(panel, ContainerItemVariant) })
			o.AddRow(packet)
			time.Sleep(time.Second)
		}
	}() //https
	go func() {
		for i := 0; i < 3; i++ {
			atomic.AddUint64(&index, 1)
			packet := packets.Object{
				Row: packets.Row{
					Method:        http.MethodGet,
					Scheme:        "wss",
					Host:          "",
					Path:          "",
					ContentType:   "json",
					ContentLength: 2220,
					Status:        "302 ssOK",
					Note:          "Websocket message",
					Process:       "",
					PadTime:       2,
				},
				Expand: packets.Expand{
					UUID:        uuid.New(),
					IsWebsocket: true,
					IsUdp:       false,
					IsTcp:       false,
					IsRequest:   false,
					IsResponse:  false,
				},
				Req:  packets.Body{},
				Resp: packets.Body{},
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
			packet := packets.Object{
				Row: packets.Row{
					Method:        http.MethodPost,
					Scheme:        "tcp",
					Host:          "",
					Path:          "",
					ContentType:   "html",
					ContentLength: 2210,
					Status:        "503 sssOK",
					Note:          "tcp message",
					Process:       "",
					PadTime:       2,
				},
				Expand: packets.Expand{},
				Req:    packets.Body{},
				Resp:   packets.Body{},
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
			packet := packets.Object{
				Row: packets.Row{
					Method:        http.MethodOptions,
					Scheme:        "udp",
					Host:          "",
					Path:          "",
					ContentType:   "binary",
					ContentLength: 1220,
					Status:        "444 Status",
					Note:          "udp message",
					Process:       "",
					PadTime:       2,
				},
				Expand: packets.Expand{
					UUID:        uuid.New(),
					IsWebsocket: false,
					IsUdp:       false,
					IsTcp:       false,
					IsRequest:   false,
					IsResponse:  false,
				},
				Req:  packets.Body{},
				Resp: packets.Body{},
			}
			if i%2 == 1 {
				packet.IsWebsocket = true
			}
			o.AddRow(packet)
			time.Sleep(time.Second)
		}
	}() //udp
}
