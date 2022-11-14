package main

import (
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/unisonUi/packets"
	"github.com/google/uuid"
	"github.com/richardwilkes/unison"
	"strconv"
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
		//child.treeIdOrSub = fmt.Sprintf("Sub Row %d", packet.Index)
		child.tips = "Websocket"
		o.branch = append(o.branch, child)
		child.table.SetRootRows(o.branch)
		o.root = append(o.root, o.branch...)
	case packet.IsTcp:
		child := row
		child.IsWebsocket = true
		child.container = true
		//child.treeIdOrSub = fmt.Sprintf("Sub Row %d", packet.Index)
		child.tips = "tcp"
		child.table.SetRootRows(o.branch)
		o.root = append(o.root, o.branch...)
	case packet.IsUdp:
		child := row
		child.IsWebsocket = true
		child.container = true
		//child.treeIdOrSub = fmt.Sprintf("Sub Row %d", packet.Index)
		child.tips = "udp"
		child.table.SetRootRows(o.branch)
		o.root = append(o.root, o.branch...)
	default:
		row.container = false
		//row.treeIdOrSub = fmt.Sprintf("Row %d", packet.Index)
		o.branch = append(o.branch, row)
		o.root = append(o.root, o.branch...)
	}
	o.packets = append(o.packets, packet)

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
		"Index",
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
	case 1:
		return o.treeIdOrSub
	case 2:
		return o.treeIdOrSub
	case 3:
		return o.treeIdOrSub
	case 4:
		return o.treeIdOrSub
	case 5:
		return o.treeIdOrSub
	case 6:
		return o.treeIdOrSub
	case 7:
		return o.treeIdOrSub
	case 8:
		return o.treeIdOrSub
	case 9:
		return o.treeIdOrSub
	case 10:
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
	case 1: //tree
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
	case 2:
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.Method, foreground, unison.LabelFont, width)
		return wrapper
	case 3:
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.Scheme, foreground, unison.LabelFont, width)
		return wrapper
	case 4:
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.Host, foreground, unison.LabelFont, width)
		if o.doubleHeight {
			addWrappedText(wrapper, "A little note…", foreground,
				unison.LabelFont.Face().Font(unison.LabelFont.Size()-1), width)
		}
		wrapper.UpdateTooltipCallback = func(where unison.Point, suggestedAvoidInRoot unison.Rect) unison.Rect {
			wrapper.Tooltip = unison.NewTooltipWithText("A tooltip for the cell")
			return wrapper.RectToRoot(wrapper.ContentRect(true))
		}
		return wrapper
	case 5:
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.ContentType, foreground, unison.LabelFont, width)
		return wrapper
	case 6:
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, fmt.Sprint(o.ContentLength), foreground, unison.LabelFont, width)
		return wrapper
	case 7:
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.Status, foreground, unison.LabelFont, width)
		return wrapper
	case 8:
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.Note, foreground, unison.LabelFont, width)
		return wrapper
	case 9:
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.Process, foreground, unison.LabelFont, width)
		return wrapper
	case 10:
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
