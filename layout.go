package main

import (
	_ "embed"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/unisonUi/asserts"
	"github.com/ddkwork/unisonUi/packets"
	"github.com/google/uuid"
	"github.com/richardwilkes/toolbox/log/jot"
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
	installDefaultMenus(w)
	content := w.Content()
	content.SetLayout(&unison.FlexLayout{Columns: 1})
	content.AddChild(createToolBar())
	content.AddChild(CreatTable()) //todo set high
	content.AddChild(createFilter())
	content.AddChild(createBodyView())

	w.Pack()
	rect := w.FrameRect()
	rect.Point = where
	w.SetFrameRect(rect)
	w.ToFront()
	return true
}

const topLevelRowsToMake = 10

var table = unison.NewTable[*demoRow](&unison.SimpleTableModel[*demoRow]{})
var panel = unison.NewPanel()

var rows = make([]*demoRow, topLevelRowsToMake)

//var rows = make([]*demoRow, 0)

func CreatTable() *unison.Panel {
	panel.SetLayout(&unison.FlexLayout{
		Columns: 1,
	})

	//panel.SetSizer(func(hint unison.Size) (min, pref, max unison.Size) {
	//	//pref.Width = 200
	//	//pref.Height = 100
	//	return min, pref, unison.MaxSize(max)
	//})

	table.HierarchyColumnIndex = 1

	o := &object{
		Object:       packets.Object{},
		packets:      nil,
		table:        unison.NewTable[*object](&unison.SimpleTableModel[*object]{}),
		parent:       nil,
		id:           uuid.UUID{},
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
	o.Table().SelectionChangedCallback = func() {
		//todo get packet to table page
	}
	table.ColumnSizes = make([]unison.ColumnSize, len(o.Header())+1)
	for i := range table.ColumnSizes {
		table.ColumnSizes[i].Minimum = 100
		table.ColumnSizes[i].Maximum = 10000
	}

	//_, checkColSize, _ := unison.NewCheckBox().Sizes(unison.Size{})
	table.ColumnSizes[0].Minimum = 20
	//table.ColumnSizes[0].Minimum = checkColSize.Width
	//table.ColumnSizes[0].Maximum = checkColSize.Width
	//rows = append(rows, &demoRow{
	//	table: table,
	//	id:    uuid.New(),
	//	text:  fmt.Sprintf("Row %d", 1),
	//	text2: fmt.Sprintf("Some longer content for Row %d", 1),
	//})
	//table.SetRootRows(rows)
	//go func() {
	for i := range rows {
		row := &demoRow{
			table: table,
			id:    uuid.New(),
			text:  fmt.Sprintf("Row %d", i+1),
			text2: fmt.Sprintf("Some longer content for Row %d", i+1),
		}
		if i%10 == 3 {
			if i == 3 {
				row.doubleHeight = true
			}
			row.container = true
			row.open = true
			row.children = make([]*demoRow, 5)
			for j := range row.children {
				child := &demoRow{
					table:  table,
					parent: row,
					id:     uuid.New(),
					text:   fmt.Sprintf("Sub Row %d", j+1),
				}
				row.children[j] = child
				if j < 2 {
					child.container = true
					child.open = true
					child.children = make([]*demoRow, 2)
					for k := range child.children {
						child.children[k] = &demoRow{
							table:  table,
							parent: child,
							id:     uuid.New(),
							text:   fmt.Sprintf("Sub Sub Row %d", k+1),
						}
					}
				}
			}
		}
		rows[i] = row
		//rows = append(rows, row)
		//table.SyncToModel()
		//table.SetRootRows(rows)
		//time.Sleep(time.Second)
	}
	//}()
	table.SetRootRows(rows)
	table.SizeColumnsToFit(true)
	table.InstallDragSupport(nil, "object", "Row", "Rows")
	unison.InstallDropSupport[*demoRow, any](table, "object",
		func(from, to *unison.Table[*demoRow]) bool { return from == to }, nil, nil)

	//header := unison.NewTableHeader[*demoRow](table,
	//	unison.NewTableColumnHeader[*demoRow]("", ""),
	//	unison.NewTableColumnHeader[*demoRow]("First", ""),
	//	unison.NewTableColumnHeader[*demoRow]("First", ""),
	//	unison.NewTableColumnHeader[*demoRow]("First", ""),
	//	unison.NewTableColumnHeader[*demoRow]("Second", ""),
	//	unison.NewTableColumnHeader[*demoRow]("xyz", ""),
	//)
	header := unison.NewTableHeader[*demoRow](table, unison.NewTableColumnHeader[*object]("", "")) //check
	for _, s := range o.Header() {
		header.ColumnHeaders = append(header.ColumnHeaders, unison.NewTableColumnHeader[*object](s, "")) //add header
	}

	header.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.FillAlignment,
		VAlign: unison.FillAlignment,
		HGrab:  true,
		VGrab:  true,
	})
	panel.AddChild(header)

	// Create a scroll panel and place a table panel inside it
	scrollArea.SetContent(table, unison.FillBehavior, unison.FillBehavior)
	scrollArea.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.FillAlignment,
		VAlign: unison.FillAlignment,
		HGrab:  true,
		VGrab:  true,
	})
	scrollArea.SetBorder(unison.NewEmptyBorder(unison.Insets{
		Top:    0,
		Left:   0,
		Bottom: 200,
		Right:  0,
	}))
	//scrollArea.SetBorder(unison.NewCompoundBorder(unison.NewLineBorder(
	//	unison.DividerColor, 0, unison.Insets{Bottom: 1},
	//	false), unison.NewEmptyBorder(unison.StdInsets())))

	scrollArea.SetLayout(&unison.FlexLayout{
		Columns:  1,
		VSpacing: unison.StdVSpacing,
	})
	panel.AddChild(scrollArea)
	return panel
}

var scrollArea = unison.NewScrollPanel()

func installDefaultMenus(wnd *unison.Window) {
	unison.DefaultMenuFactory().BarForWindow(wnd, func(m unison.Menu) {
		unison.InsertStdMenus(m, ShowAboutWindow, nil, nil)
		fileMenu := m.Menu(unison.FileMenuID)
		f := fileMenu.Factory()
		newMenu := f.NewMenu(0, "New…", nil)
		//newMenu.InsertItem(-1, unison.NewMenuItem(f))
		//newMenu.InsertItem(-1, NewWindowAction.NewMenuItem(f))
		//newMenu.InsertItem(-1, NewTableWindowAction.NewMenuItem(f))
		//newMenu.InsertItem(-1, NewDockWindowAction.NewMenuItem(f))
		fileMenu.InsertMenu(0, newMenu)
		//fileMenu.InsertItem(1, OpenAction.NewMenuItem(f))
	})
}

var dock = unison.NewDock()

func createBodyView() *unison.Dock {
	yellowDockable := NewDockablePanel(packets.NameBodyKind.HttpDump(), "Request", unison.Yellow)
	dock.DockTo(yellowDockable, nil, unison.TopSide)
	unison.Ancestor[*unison.DockContainer](yellowDockable).Stack(NewDockablePanel(packets.NameBodyKind.HexDump(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellowDockable).Stack(NewDockablePanel(packets.NameBodyKind.Steam(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellowDockable).Stack(NewDockablePanel(packets.NameBodyKind.ProtoBuf(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellowDockable).Stack(NewDockablePanel(packets.NameBodyKind.Tdf(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellowDockable).Stack(NewDockablePanel(packets.NameBodyKind.Taf(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellowDockable).Stack(NewDockablePanel(packets.NameBodyKind.Acc(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellowDockable).Stack(NewDockablePanel(packets.NameBodyKind.Websocket(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellowDockable).Stack(NewDockablePanel(packets.NameBodyKind.Msgpack(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellowDockable).Stack(NewDockablePanel(packets.NameBodyKind.Notes(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellowDockable).Stack(NewDockablePanel(packets.NameBodyKind.UnitTest(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellowDockable).Stack(NewDockablePanel(packets.NameBodyKind.GitProxy(), "", unison.Yellow), -1)

	blueDockable := NewDockablePanel(packets.NameBodyKind.HttpDump(), "Response", unison.Pink)
	//blueDockable.MayAttemptClose() //todo close button disable
	dock.DockTo(blueDockable, nil, unison.BottomSide)
	unison.Ancestor[*unison.DockContainer](blueDockable).Stack(NewDockablePanel(packets.NameBodyKind.HexDump(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blueDockable).Stack(NewDockablePanel(packets.NameBodyKind.Steam(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blueDockable).Stack(NewDockablePanel(packets.NameBodyKind.ProtoBuf(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blueDockable).Stack(NewDockablePanel(packets.NameBodyKind.Tdf(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blueDockable).Stack(NewDockablePanel(packets.NameBodyKind.Taf(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blueDockable).Stack(NewDockablePanel(packets.NameBodyKind.Acc(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blueDockable).Stack(NewDockablePanel(packets.NameBodyKind.Websocket(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blueDockable).Stack(NewDockablePanel(packets.NameBodyKind.Msgpack(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blueDockable).Stack(NewDockablePanel(packets.NameBodyKind.Notes(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blueDockable).Stack(NewDockablePanel(packets.NameBodyKind.UnitTest(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blueDockable).Stack(NewDockablePanel(packets.NameBodyKind.GitProxy(), "", unison.Pink), -1)

	dock.SetLayoutData(&unison.FlexLayoutData{
		HSpan:  1,
		VSpan:  100,
		HAlign: unison.FillAlignment,
		VAlign: unison.FillAlignment,
		HGrab:  true,
		VGrab:  true,
	})

	return dock
}

func createToolBar() *unison.Panel {
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlowLayout{
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})

	// Load our home image, and if successful (we should be!), add two buttons with it, one enabled and one not.
	homeImg, err := HomeImage()
	if err != nil {
		jot.Error(err)
	} else {
		createImageButton(homeImg, "home_enabled", panel)
		createImageButton(homeImg, "home_disabled", panel).SetEnabled(false)
	}

	// Load our logo image, and if successful (we should be!), add two buttons with it, one enabled and one not.
	var logoImg *unison.Image
	if logoImg, err = ClassicAppleLogoImage(); err != nil {
		jot.Error(err)
	} else {
		createImageButton(logoImg, "logo_enabled", panel)
		createImageButton(logoImg, "logo_disabled", panel).SetEnabled(false)
	}

	if homeImg != nil && logoImg != nil {
		// Add spacer
		//spacer := &unison.Panel{}
		//spacer.Self = spacer
		//spacer.SetSizer(func(_ unison.Size) (min, pref, max unison.Size) {
		//	min.Width = 40
		//	pref.Width = 40
		//	max.Width = 40
		//	return
		//})
		//panel.AddChild(spacer)

		// Add some sticky buttons in a group with our images
		group := unison.NewGroup()
		first := createImageButton(homeImg, "home_toggle", panel)
		first.Sticky = true
		group.Add(first.AsGroupPanel())
		second := createImageButton(logoImg, "logo_toggle", panel)
		second.Sticky = true
		group.Add(second.AsGroupPanel())
		group.Select(first.AsGroupPanel())
	}
	return panel
}

var (
	//go:embed resources/classic-apple-logo.png
	classicAppleLogoPngBytes []byte
	classicAppleLogoImage    *unison.Image
)

// ClassicAppleLogoImage returns an image of the classic rainbow-colored Apple logo.
func ClassicAppleLogoImage() (*unison.Image, error) {
	if classicAppleLogoImage == nil {
		var err error
		if classicAppleLogoImage, err = unison.NewImageFromBytes(classicAppleLogoPngBytes, 0.5); err != nil {
			return nil, err
		}
	}
	return classicAppleLogoImage, nil
}

// HomeImage returns a stylized image of a home, suitable for an icon.
func HomeImage() (*unison.Image, error) {
	if homeImage == nil {
		var err error
		if homeImage, err = unison.NewImageFromBytes(homePngBytes, 0.5); err != nil {
			return nil, err
		}
	}
	return homeImage, nil
}

var (
	//go:embed resources/home.png
	homePngBytes []byte
	homeImage    *unison.Image
)

func createImageButton(img *unison.Image, actionText string, panel *unison.Panel) *unison.Button {
	btn := unison.NewButton()
	btn.Drawable = img
	btn.ClickCallback = func() { jot.Info(actionText) }
	btn.Tooltip = unison.NewTooltipWithText(fmt.Sprintf("Tooltip for: %s", actionText))
	btn.SetLayoutData(unison.MiddleAlignment)
	panel.AddChild(btn)
	return btn
}

func createFilter() *unison.Panel {
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlexLayout{
		Columns:      10,
		HSpacing:     unison.StdHSpacing,
		VSpacing:     unison.StdVSpacing,
		EqualColumns: true,
		//HAlign:       unison.StartAlignment,
	})
	field := unison.NewField()
	field.SetLayoutData(&unison.FlexLayoutData{
		HSpan:  1,
		VSpan:  1,
		HAlign: unison.FillAlignment,
		VAlign: unison.MiddleAlignment,
		HGrab:  true,
	})
	field.Tooltip = unison.NewTooltipWithText("filter")
	field.ModifiedCallback = func() {
		println(field.Text())
	}
	panel.AddChild(field)

	spacer := &unison.Panel{}
	spacer.Self = spacer
	spacer.SetSizer(func(_ unison.Size) (min, pref, max unison.Size) {
		min.Width = 40
		pref.Width = 40
		max.Width = 40
		return
	})
	panel.AddChild(spacer)

	createCheckBox("RequestCtx", unison.OffCheckState, panel)
	createCheckBox("ResponseCtx", unison.OffCheckState, panel)
	createCheckBox("RequestBody", unison.OffCheckState, panel)
	createCheckBox("ResponseBody", unison.OffCheckState, panel)
	createCheckBox("Notes", unison.OffCheckState, panel)
	createCheckBox("Scheme", unison.OffCheckState, panel)
	return panel
}

func createCheckBox(title string, initialState unison.CheckState, panel *unison.Panel) *unison.CheckBox {
	check := unison.NewCheckBox()
	check.Text = title
	check.State = initialState
	check.ClickCallback = func() { jot.Infof("'%s' was clicked.", title) }
	check.Tooltip = unison.NewTooltipWithText(fmt.Sprintf("This is the tooltip for '%s'", title))
	panel.AddChild(check)
	return check
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
