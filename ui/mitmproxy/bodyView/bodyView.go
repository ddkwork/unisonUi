package bodyView

import (
	"github.com/ddkwork/golibrary/skiaLib/widget/tabbar"
	"github.com/ddkwork/unisonUi/packets"
	"github.com/richardwilkes/unison"
)

//todo add interface

func CreateBodyView() *unison.Dock { //todo move all to BottomSide
	var dock = unison.NewDock()
	yellow := tabbar.New(packets.NameBodyKind.HttpDump(), "Request", unison.Yellow)
	dock.DockTo(yellow, nil, unison.BottomSide)

	HexDump := tabbar.New(packets.NameBodyKind.HexDump(), "", unison.Yellow)
	HexDumpPanel := HexDump.AsPanel()
	HexDumpPanel.AddChild(unison.NewButton()) //todo test show HexDump ctx, not passed

	unison.Ancestor[*unison.DockContainer](yellow).Stack(HexDump, -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(packets.NameBodyKind.Steam(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(packets.NameBodyKind.ProtoBuf(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(packets.NameBodyKind.Tdf(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(packets.NameBodyKind.Taf(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(packets.NameBodyKind.Acc(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(packets.NameBodyKind.Websocket(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(packets.NameBodyKind.Msgpack(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(packets.NameBodyKind.Notes(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(packets.NameBodyKind.UnitTest(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(tabbar.New(packets.NameBodyKind.GitProxy(), "", unison.Yellow), -1)

	blue := tabbar.New(packets.NameBodyKind.HttpDump(), "Response", unison.Pink)
	dock.DockTo(blue, nil, unison.BottomSide)
	unison.Ancestor[*unison.DockContainer](blue).Stack(tabbar.New(packets.NameBodyKind.HexDump(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(tabbar.New(packets.NameBodyKind.Steam(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(tabbar.New(packets.NameBodyKind.ProtoBuf(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(tabbar.New(packets.NameBodyKind.Tdf(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(tabbar.New(packets.NameBodyKind.Taf(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(tabbar.New(packets.NameBodyKind.Acc(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(tabbar.New(packets.NameBodyKind.Websocket(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(tabbar.New(packets.NameBodyKind.Msgpack(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(tabbar.New(packets.NameBodyKind.Notes(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(tabbar.New(packets.NameBodyKind.UnitTest(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(tabbar.New(packets.NameBodyKind.GitProxy(), "", unison.Pink), -1)

	dock.SetLayoutData(&unison.FlexLayoutData{
		HSpan:  1,
		VSpan:  200,
		HAlign: unison.FillAlignment,
		VAlign: unison.FillAlignment,
		HGrab:  true,
		VGrab:  true,
	})
	return dock
}
