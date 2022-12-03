package table

import (
	"github.com/ddkwork/unisonUi/packets"
	"github.com/richardwilkes/unison"
	"sync"
	"time"
)

func (o *object) AddRow(packet packets.Object) {
	o.Lock()
	defer o.Unlock()
	o.Object = packet
	row := &object{
		Object:  packet,
		packets: o.packets,
		//table:   unison.NewTable[*object](&unison.SimpleTableModel[*object]{}),
		parent: o.parent,
		//id:           uuid.UUID{},
		treeIdOrSub:  "",
		tips:         "",
		root:         o.root,
		branch:       o.branch,
		checkbox:     unison.NewCheckBox(),
		container:    false,
		open:         true,
		doubleHeight: false,
		RWMutex:      sync.RWMutex{},
	}
	switch {
	case packet.IsWebsocket:
		child := row
		child.IsWebsocket = true
		child.container = true
		//child.treeIdOrSub = fmt.Sprintf("Sub Row %d", packet.Index)
		child.tips = "Websocket"
		o.branch = append(o.branch, child)
		//child.table.SetRootRows(o.branch)
		o.root = append(o.root, o.branch...)
	case packet.IsTcp:
		child := row
		child.IsWebsocket = true
		child.container = true
		//child.treeIdOrSub = fmt.Sprintf("Sub Row %d", packet.Index)
		child.tips = "tcp"
		//child.table.SetRootRows(o.branch)
		o.root = append(o.root, o.branch...)
	case packet.IsUdp:
		child := row
		child.IsWebsocket = true
		child.container = true
		//child.treeIdOrSub = fmt.Sprintf("Sub Row %d", packet.Index)
		child.tips = "udp"
		//child.table.SetRootRows(o.branch)
		o.root = append(o.root, o.branch...)
	default:
		row.container = false
		//row.treeIdOrSub = fmt.Sprintf("Row %d", packet.Index)
		o.branch = append(o.branch, row)
		o.root = append(o.root, o.branch...)
	}
	o.packets = append(o.packets, packet)
	time.Sleep(time.Second)
}

func (o *object) Header() []string {
	return []string{
		//checkBox todo add checkBox
		"Proto", //todo add Proto with ico
		packets.NamePacketField.Method(),
		packets.NamePacketField.Scheme(),
		packets.NamePacketField.Host(),
		packets.NamePacketField.Path(),
		packets.NamePacketField.ContentType(),
		packets.NamePacketField.ContentLength(),
		packets.NamePacketField.Status(),
		packets.NamePacketField.Notes(),
		packets.NamePacketField.Process(),
		packets.NamePacketField.PadTime(),
	}
}
