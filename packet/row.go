package packet

import (
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/unisonUi/objects"
	"github.com/google/uuid"
	"github.com/richardwilkes/unison"
	"sort"
	"strconv"
	"sync"
	"time"
)

func (o *object) AddRow(packet objects.Packet) {
	o.Lock()
	defer o.Unlock()
	o.Packet = packet
	row := &object{
		Packet:  packet,
		packets: o.packets,
		table:   unison.NewTable[*object](&unison.SimpleTableModel[*object]{}),
		parent:  o.parent,
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
		child.treeIdOrSub = fmt.Sprintf("Sub Row %d", packet.Index)
		child.tips = "Websocket"
		o.branch = append(o.branch, child)
		child.table.SetRootRows(o.branch)
		o.root = append(o.root, o.branch...)
	case packet.IsTcp:
		child := row
		child.IsWebsocket = true
		child.container = true
		child.treeIdOrSub = fmt.Sprintf("Sub Row %d", packet.Index)
		child.tips = "tcp"
		child.table.SetRootRows(o.branch)
		o.root = append(o.root, o.branch...)
	case packet.IsUdp:
		child := row
		child.IsWebsocket = true
		child.container = true
		child.treeIdOrSub = fmt.Sprintf("Sub Row %d", packet.Index)
		child.tips = "udp"
		child.table.SetRootRows(o.branch)
		o.root = append(o.root, o.branch...)
	default:
		row.container = false
		row.treeIdOrSub = fmt.Sprintf("Row %d", packet.Index)
		o.branch = append(o.branch, row)
		o.root = append(o.root, o.branch...)
	}
	o.packets.packets = append(o.packets.packets, packet)

	o.table.SetRootRows(o.root)
	o.Table().SyncToModel()
	o.Table().SizeColumnsToFit(true)
	o.Table().InstallDragSupport(nil, "object", "Row", "Column")
	unison.InstallDropSupport[*object, any](o.Table(), "object", func(from, to *unison.Table[*object]) bool { return from == to }, nil, nil)

	time.Sleep(time.Second)
}

func (o *object) Table() *unison.Table[*object] { return o.table }
func (o *object) Header() []string {
	return []string{
		objects.NamePacketField.Index(),
		objects.NamePacketField.Method(),
		objects.NamePacketField.Scheme(),
		objects.NamePacketField.Url(),
		objects.NamePacketField.ContentType(),
		objects.NamePacketField.ContentLength(),
		objects.NamePacketField.Status(),
		objects.NamePacketField.Notes(),
		objects.NamePacketField.StartTime(),
		objects.NamePacketField.PadTime(),
	}
}

func (o *object) CloneForTarget(target unison.Paneler, newParent *object) *object {
	table, ok := target.(*unison.Table[*object])
	if !ok {
		mylog.Error("invalid target")
		return nil
	}
	clone := *o
	clone.table = table
	clone.parent = newParent
	clone.id = uuid.New()
	return &clone
}

func (o *object) UUID() uuid.UUID                { return o.id }
func (o *object) Parent() *object                { return o.parent }
func (o *object) SetParent(parent *object)       { o.parent = parent }
func (o *object) CanHaveChildren() bool          { return o.container }
func (o *object) Children() []*object            { return o.branch }
func (o *object) SetChildren(children []*object) { o.branch = children }
func (o *object) CellDataForSort(col int) string {
	switch col {
	case 0:
		if o.checkbox == nil {
			o.checkbox = unison.NewCheckBox()
		}
		return strconv.Itoa(int(o.checkbox.State))
	case o.packets.HeaderIndex(objects.NamePacketField.Index()):
		return o.treeIdOrSub
	case o.packets.HeaderIndex(objects.NamePacketField.Method()):
		return o.treeIdOrSub
	case o.packets.HeaderIndex(objects.NamePacketField.Scheme()):
		return o.treeIdOrSub
	case o.packets.HeaderIndex(objects.NamePacketField.Url()):
		return o.treeIdOrSub
	case o.packets.HeaderIndex(objects.NamePacketField.ContentType()):
		return o.treeIdOrSub
	case o.packets.HeaderIndex(objects.NamePacketField.ContentLength()):
		return o.treeIdOrSub
	case o.packets.HeaderIndex(objects.NamePacketField.Status()):
		return o.treeIdOrSub
	case o.packets.HeaderIndex(objects.NamePacketField.Notes()):
		return o.treeIdOrSub
	case o.packets.HeaderIndex(objects.NamePacketField.StartTime()):
		return o.treeIdOrSub
	case o.packets.HeaderIndex(objects.NamePacketField.PadTime()):
		return o.treeIdOrSub
	default:
		return ""
	}
}

func (o *object) ColumnCell(row, col int, foreground, background unison.Ink, selected, indirectlySelected, focused bool) unison.Paneler {
	//mylog.Trace("selected", selected)
	switch col {
	case 0:
		if o.checkbox == nil {
			o.checkbox = unison.NewCheckBox()
		}
		return o.checkbox
	case o.packets.HeaderIndex(objects.NamePacketField.Index()): //tree
		//mylog.Trace("container", o.container)
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.treeIdOrSub, foreground, unison.LabelFont, width)
		if o.doubleHeight {
			addWrappedText(wrapper, "A little note…", foreground,
				unison.LabelFont.Face().Font(unison.LabelFont.Size()-1), width)
		}
		wrapper.UpdateTooltipCallback = func(where unison.Point, suggestedAvoidInRoot unison.Rect) unison.Rect {
			wrapper.Tooltip = unison.NewTooltipWithText("A tooltip for the cell")
			return wrapper.RectToRoot(wrapper.ContentRect(true))
		}
		return wrapper
	case o.packets.HeaderIndex(objects.NamePacketField.Method()):
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.Method, foreground, unison.LabelFont, width)
		return wrapper
	case o.packets.HeaderIndex(objects.NamePacketField.Scheme()):
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.Scheme, foreground, unison.LabelFont, width)
		return wrapper
	case o.packets.HeaderIndex(objects.NamePacketField.Url()):
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.Url, foreground, unison.LabelFont, width)
		if o.doubleHeight {
			addWrappedText(wrapper, "A little note…", foreground,
				unison.LabelFont.Face().Font(unison.LabelFont.Size()-1), width)
		}
		wrapper.UpdateTooltipCallback = func(where unison.Point, suggestedAvoidInRoot unison.Rect) unison.Rect {
			wrapper.Tooltip = unison.NewTooltipWithText("A tooltip for the cell")
			return wrapper.RectToRoot(wrapper.ContentRect(true))
		}
		return wrapper
	case o.packets.HeaderIndex(objects.NamePacketField.ContentType()):
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.ContentType, foreground, unison.LabelFont, width)
		return wrapper
	case o.packets.HeaderIndex(objects.NamePacketField.ContentLength()):
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, fmt.Sprint(o.ContentLength), foreground, unison.LabelFont, width)
		return wrapper
	case o.packets.HeaderIndex(objects.NamePacketField.Status()):
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.Status, foreground, unison.LabelFont, width)
		return wrapper
	case o.packets.HeaderIndex(objects.NamePacketField.Notes()):
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.Note, foreground, unison.LabelFont, width)
		return wrapper
	case o.packets.HeaderIndex(objects.NamePacketField.StartTime()):
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.StartTime.String(), foreground, unison.LabelFont, width)
		return wrapper
	case o.packets.HeaderIndex(objects.NamePacketField.PadTime()):
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 2})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.PadTime.String(), foreground, unison.LabelFont, width)
		return wrapper
	default:
		//mylog.Info("column index out of range (0-2): %d", col)
		return unison.NewLabel()
	}
}

func (o *object) IsOpen() bool      { return o.open }
func (o *object) SetOpen(open bool) { o.open = open }

func (p packets) Header() []string {
	return []string{
		objects.NamePacketField.Index(),
		objects.NamePacketField.Method(),
		objects.NamePacketField.Scheme(),
		objects.NamePacketField.Url(),
		objects.NamePacketField.ContentType(),
		objects.NamePacketField.ContentLength(),
		objects.NamePacketField.Status(),
		objects.NamePacketField.Notes(),
		objects.NamePacketField.StartTime(),
		objects.NamePacketField.PadTime(),
	}
}

func (p packets) HeaderIndex(h string) int {
	switch h {
	case objects.NamePacketField.Index():
		return 1
	case objects.NamePacketField.Method():
		return 2
	case objects.NamePacketField.Scheme():
		return 3
	case objects.NamePacketField.Url():
		return 4
	case objects.NamePacketField.ContentType():
		return 5
	case objects.NamePacketField.ContentLength():
		return 6
	case objects.NamePacketField.Status():
		return 7
	case objects.NamePacketField.Notes():
		return 8
	case objects.NamePacketField.StartTime():
		return 9
	case objects.NamePacketField.PadTime():
		return 10
	}
	return -1
}

func (p packets) Rows(id int) []string {
	if id < 0 || id >= p.Len() {
		return nil
	}
	return []string{
		fmt.Sprintf("%03d", p.packets[id].Index),
		p.packets[id].Method,
		p.packets[id].Scheme,
		p.packets[id].Url,
		p.packets[id].ContentType,
		fmt.Sprint(p.packets[id].ContentLength),
		p.packets[id].Status,
		p.packets[id].Note,
		p.packets[id].StartTime.String(),
		p.packets[id].PadTime.String(),
	}
}

func (p packets) Len() int { return len(p.packets) }

func (p packets) ColumnLen() int {
	//TODO implement me
	panic("implement me")
}

func (p packets) Sort(id int, ascend bool) {
	switch p.Header()[id] {
	case objects.NamePacketField.Index():
		sort.Slice(p.packets, func(i, j int) bool { return p.packets[i].Index > p.packets[j].Index })
	case objects.NamePacketField.Method():
		sort.Slice(p.packets, func(i, j int) bool { return p.packets[i].Method > p.packets[j].Method })
	case objects.NamePacketField.Scheme():
		sort.Slice(p.packets, func(i, j int) bool { return p.packets[i].Scheme > p.packets[j].Scheme })
	case objects.NamePacketField.Url():
		sort.Slice(p.packets, func(i, j int) bool { return p.packets[i].Url > p.packets[j].Url })
	case objects.NamePacketField.ContentType():
		sort.Slice(p.packets, func(i, j int) bool { return p.packets[i].ContentType > p.packets[j].ContentType })
	case objects.NamePacketField.ContentLength():
		sort.Slice(p.packets, func(i, j int) bool { return p.packets[i].ContentLength > p.packets[j].ContentLength })
	case objects.NamePacketField.PadTime():
		sort.Slice(p.packets, func(i, j int) bool { return p.packets[i].PadTime > p.packets[j].PadTime })
	case objects.NamePacketField.StartTime():
		sort.Slice(p.packets, func(i, j int) bool { return p.packets[i].StartTime.Unix() > p.packets[j].StartTime.Unix() })
	case objects.NamePacketField.Status():
		sort.Slice(p.packets, func(i, j int) bool { return p.packets[i].Status > p.packets[j].Status })
	case objects.NamePacketField.Notes():
		sort.Slice(p.packets, func(i, j int) bool { return p.packets[i].Note > p.packets[j].Note })
	}
}

func (p packets) Filter(row string, id int) {
	//TODO implement me
	panic("implement me")
}

func (p packets) Append(data any) { p.packets = append(p.packets, data.(objects.Packet)) }
func (p packets) ColumnWidths() []float32 {
	//TODO implement me
	panic("implement me")
}
