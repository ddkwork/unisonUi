package main

import (
	_ "embed"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/unisonUi/asserts"
	"github.com/ddkwork/unisonUi/packets"
	"github.com/ddkwork/unisonUi/ui/mitmproxy/bodyView"
	"github.com/ddkwork/unisonUi/ui/mitmproxy/menus"
	"github.com/ddkwork/unisonUi/ui/mitmproxy/toolbar"
	"github.com/google/uuid"
	"github.com/richardwilkes/unison"
	"sync"
)

type (
	ui interface {
		CanvasObject(where unison.Point) (ok bool)
		Mitmproxy() (ok bool) //for test in main
		unison.TableRowData[*object]
		Table() *unison.Table[*object]
		Header() []string
	}
	object struct {
		packets.Object //set col style
		packets        []packets.Object
		table          *unison.Table[*object]
		parent         *object
		id             uuid.UUID
		treeIdOrSub    string
		tips           string
		root           []*object
		branch         []*object
		//webSocket      []*object
		//tcp            []*object
		//udp            []*object
		checkbox     *unison.CheckBox
		container    bool
		open         bool
		doubleHeight bool
		sync.RWMutex
	}
)

func New() *object { return &object{} }
func CanvasObject(where unison.Point) (ok bool) {
	w, err := unison.NewWindow(fmt.Sprintf("mitmproxy"))
	if err != nil {
		return
	}
	w.MinMaxContentSizeCallback = func() (min, max unison.Size) {
		return unison.NewSize(1000, 600), unison.NewSize(10000, 1280)
	}
	image, err := unison.NewImageFromBytes(asserts.MitmBuf, 0.5)
	if !mylog.Error(err) {
		return
	}
	w.SetTitleIcons([]*unison.Image{image})

	menus.InstallDefaultMenus(w)
	content := w.Content()
	content.SetLayout(&unison.FlexLayout{Columns: 1})
	content.AddChild(toolbar.CreateToolBar())
	//content.AddChild(CreatTable()) //todo set high
	//content.AddChild(createFilter())
	content.AddChild(bodyView.CreateBodyView())

	w.Pack()
	rect := w.FrameRect()
	rect.Point = where
	w.SetFrameRect(rect)
	w.ToFront()
	return true
}

func CanvasObject_(where unison.Point) (ok bool) {
	w, err := unison.NewWindow("mitmproxy")
	if !mylog.Error(err) {
		return
	}
	w.MinMaxContentSizeCallback = func() (min, max unison.Size) {
		return unison.NewSize(1000, 600), unison.NewSize(10000, 1280)
	}
	image, err := unison.NewImageFromBytes(asserts.MitmBuf, 0.5)
	if !mylog.Error(err) {
		return
	}
	w.SetTitleIcons([]*unison.Image{image})
	o := &object{
		Object:  packets.Object{},
		packets: nil,
		table:   unison.NewTable[*object](&unison.SimpleTableModel[*object]{}),
		parent:  nil,
		id:      uuid.UUID{},
		//id:           uuid.UUID{},
		treeIdOrSub:  "",
		tips:         "",
		root:         make([]*object, 0),
		branch:       make([]*object, 0),
		checkbox:     nil,
		container:    false,
		open:         false,
		doubleHeight: false,
		RWMutex:      sync.RWMutex{},
	}
	o.table.ShowRowDivider = false
	o.table.HierarchyColumnIndex = 1
	o.table.ColumnSizes = make([]unison.ColumnSize, len(o.Header())+1)
	for i := range o.table.ColumnSizes {
		o.table.ColumnSizes[i].Minimum = 20
		o.table.ColumnSizes[i].Maximum = 10000
	}
	_, checkColSize, _ := unison.NewCheckBox().Sizes(unison.Size{})
	o.table.ColumnSizes[0].Minimum = checkColSize.Width
	o.table.ColumnSizes[0].Maximum = checkColSize.Width

	o.Table().SelectionChangedCallback = func() {
		//todo get packet to table page
	}
	o.Table().SetRootRows(o.root)
	o.Table().SizeColumnsToFit(true)
	o.Table().InstallDragSupport(nil, "object", "Row", "Column")
	unison.InstallDropSupport[*object, any](o.Table(), "object", func(from, to *unison.Table[*object]) bool { return from == to }, nil, nil)

	o.Table().SetRootRows(o.root)
	o.Table().SizeColumnsToFit(true)
	o.Table().InstallDragSupport(nil, "object", "Row", "Column")
	unison.InstallDropSupport[*object, any](o.Table(), "object", func(from, to *unison.Table[*object]) bool { return from == to }, nil, nil)

	go o.mitmMock()
	header := unison.NewTableHeader[*object](o.Table(), unison.NewTableColumnHeader[*object]("", "")) //check
	for _, s := range o.Header() {
		header.ColumnHeaders = append(header.ColumnHeaders, unison.NewTableColumnHeader[*object](s, "")) //add header
	}
	header.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.FillAlignment,
		VAlign: unison.FillAlignment,
		HGrab:  true,
	})
	content := w.Content()
	content.SetLayout(&unison.FlexLayout{Columns: 1})
	content.AddChild(header)

	// Create a scroll panel and tips a table panel inside it
	scrollArea := unison.NewScrollPanel()
	scrollArea.SetContent(o.Table(), unison.FillBehavior, unison.FillBehavior)
	scrollArea.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.FillAlignment,
		VAlign: unison.FillAlignment,
		HGrab:  true,
		VGrab:  true,
	})
	content.AddChild(scrollArea)
	w.Pack()
	rect := w.FrameRect()
	rect.Point = where
	w.SetFrameRect(rect)
	w.ToFront()
	return true
}
