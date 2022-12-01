package bodyView

import (
	"github.com/ddkwork/unisonUi/packets"
	"github.com/richardwilkes/unison"
)

var dock = unison.NewDock()

func CreateBodyView() *unison.Dock {
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
