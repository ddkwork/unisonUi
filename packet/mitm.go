package packet

import (
	"github.com/ddkwork/unisonUi/objects"
	"sync"
)

var once = sync.Once{}

func (o *object) SetData(packet objects.Packet) {
	//o.Lock()
	//defer o.Unlock()
	o.packets.packets = append(o.packets.packets, packet)
	//once.Do(func() {
	//o.table.SetData(o.packets)
	//})
	//o.table.Refresh()
	//o.Unlock()
}
