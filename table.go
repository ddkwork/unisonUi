package main

import (
	"fmt"
	"github.com/ddkwork/unisonUi/packets"
	"github.com/google/uuid"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/unison"
	"sync"
)

func createCheckBox(title string, initialState unison.CheckState, panel *unison.Panel) *unison.CheckBox {
	check := unison.NewCheckBox()
	check.Text = title
	check.State = initialState
	check.ClickCallback = func() { jot.Infof("'%s' was clicked.", title) }
	check.Tooltip = unison.NewTooltipWithText(fmt.Sprintf("This is the tooltip for '%s'", title))
	panel.AddChild(check)
	return check
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
