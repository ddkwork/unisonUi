package bodyView

import (
	"fmt"
	"github.com/ddkwork/golibrary/skiaLib/widget/tabbar"
	"github.com/ddkwork/unisonUi/packets"
	"github.com/richardwilkes/unison"
)

//todo add interface

type TabBarWithMultiLineField struct {
	tabbar.Interface
	*unison.Field
}

func NewTabBarWithMultiLineField(title, tip string) *TabBarWithMultiLineField {
	tabBar := tabbar.New(title, tip, unison.Yellow)
	HexDumpPanel := tabBar.AsPanel()
	field := unison.NewMultiLineField()
	field.SetSizer(func(hint unison.Size) (min, pref, max unison.Size) {
		pref = unison.NewSize(unison.DefaultMaxSize, unison.DefaultMaxSize)
		return pref, pref, pref
	})
	HexDumpPanel.AddChild(field)
	HexDumpPanel.SetLayout(&unison.FlowLayout{
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})
	return &TabBarWithMultiLineField{
		Interface: tabBar,
		Field:     field,
	}
}

func CreateBodyView() *unison.Dock { //todo move all to BottomSide
	var dock = unison.NewDock()
	fn := func(kind string) {
		HttpDump := NewTabBarWithMultiLineField(packets.NameBodyKind.HttpDump(), kind+" "+packets.NameBodyKind.HttpDump())
		dock.DockTo(HttpDump, nil, unison.BottomSide)
		HexDump := NewTabBarWithMultiLineField(packets.NameBodyKind.HexDump(), kind+" "+packets.NameBodyKind.HexDump())
		Steam := NewTabBarWithMultiLineField(packets.NameBodyKind.Steam(), kind+" "+packets.NameBodyKind.Steam())
		ProtoBuf := NewTabBarWithMultiLineField(packets.NameBodyKind.ProtoBuf(), kind+" "+packets.NameBodyKind.ProtoBuf())
		Tdf := NewTabBarWithMultiLineField(packets.NameBodyKind.Tdf(), kind+" "+packets.NameBodyKind.Tdf())
		Taf := NewTabBarWithMultiLineField(packets.NameBodyKind.Taf(), kind+" "+packets.NameBodyKind.Taf())
		Acc := NewTabBarWithMultiLineField(packets.NameBodyKind.Acc(), kind+" "+packets.NameBodyKind.Acc())
		Websocket := NewTabBarWithMultiLineField(packets.NameBodyKind.Websocket(), kind+" "+packets.NameBodyKind.Websocket())
		Msgpack := NewTabBarWithMultiLineField(packets.NameBodyKind.Msgpack(), kind+" "+packets.NameBodyKind.Msgpack())
		Notes := NewTabBarWithMultiLineField(packets.NameBodyKind.Notes(), kind+" "+packets.NameBodyKind.Notes())
		UnitTest := NewTabBarWithMultiLineField(packets.NameBodyKind.UnitTest(), kind+" "+packets.NameBodyKind.UnitTest())
		GitProxy := NewTabBarWithMultiLineField(packets.NameBodyKind.GitProxy(), kind+" "+packets.NameBodyKind.GitProxy())

		unison.Ancestor[*unison.DockContainer](HttpDump.Interface).Stack(HexDump, -1)
		unison.Ancestor[*unison.DockContainer](HttpDump.Interface).Stack(Steam, -1)
		unison.Ancestor[*unison.DockContainer](HttpDump.Interface).Stack(ProtoBuf, -1)
		unison.Ancestor[*unison.DockContainer](HttpDump.Interface).Stack(Tdf, -1)
		unison.Ancestor[*unison.DockContainer](HttpDump.Interface).Stack(Taf, -1)
		unison.Ancestor[*unison.DockContainer](HttpDump.Interface).Stack(Acc, -1)
		unison.Ancestor[*unison.DockContainer](HttpDump.Interface).Stack(Websocket, -1)
		unison.Ancestor[*unison.DockContainer](HttpDump.Interface).Stack(Msgpack, -1)
		unison.Ancestor[*unison.DockContainer](HttpDump.Interface).Stack(Notes, -1)
		unison.Ancestor[*unison.DockContainer](HttpDump.Interface).Stack(UnitTest, -1)
		unison.Ancestor[*unison.DockContainer](HttpDump.Interface).Stack(GitProxy, -1)
		//todo first select HttpDump
	}

	fn("Request")
	fn("Response")

	dock.SetLayoutData(&unison.FlexLayoutData{
		HSpan:  1,
		VSpan:  1,
		HAlign: unison.FillAlignment,
		VAlign: unison.FillAlignment,
		HGrab:  true,
		VGrab:  true,
	})
	return dock
}

func createMultiLineTextField(labelText, fieldText string, panel *unison.Panel) *unison.Field {
	lbl := unison.NewLabel()
	lbl.Text = labelText
	lbl.HAlign = unison.EndAlignment
	lbl.SetLayoutData(&unison.FlexLayoutData{
		HSpan:  1,
		VSpan:  1,
		HAlign: unison.EndAlignment,
		VAlign: unison.MiddleAlignment,
	})
	panel.AddChild(lbl)
	field := unison.NewMultiLineField()
	field.SetText(fieldText)
	field.SetLayoutData(&unison.FlexLayoutData{
		HSpan:  1,
		VSpan:  1,
		HAlign: unison.FillAlignment,
		VAlign: unison.MiddleAlignment,
		HGrab:  true,
	})
	field.Tooltip = unison.NewTooltipWithText(fmt.Sprintf("This is the tooltip for %v", field))
	panel.AddChild(field)
	return field
}
