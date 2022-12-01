package table

import (
	"github.com/ddkwork/unisonUi/packets"
	"sync"
)

var once = sync.Once{}

func (o *object) SetData(packet packets.Object) {
	//o.Lock()
	//defer o.Unlock()
	o.packets = append(o.packets, packet)
	//once.Do(func() {
	//o.table.SetData(o.packets)
	//})
	//o.table.Refresh()
	//o.Unlock()
}
