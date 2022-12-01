package bodyView

import (
	"github.com/ddkwork/unisonUi/packets"
	"github.com/ddkwork/unisonUi/widget/apptable"
	"github.com/richardwilkes/unison"
)

func CreateBodyView() *unison.Dock { //todo move all to BottomSide
	var dock = unison.NewDock()
	yellow := apptable.New(packets.NameBodyKind.HttpDump(), "Request", unison.Yellow)
	dock.DockTo(yellow, nil, unison.BottomSide)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(apptable.New(packets.NameBodyKind.HexDump(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(apptable.New(packets.NameBodyKind.Steam(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(apptable.New(packets.NameBodyKind.ProtoBuf(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(apptable.New(packets.NameBodyKind.Tdf(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(apptable.New(packets.NameBodyKind.Taf(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(apptable.New(packets.NameBodyKind.Acc(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(apptable.New(packets.NameBodyKind.Websocket(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(apptable.New(packets.NameBodyKind.Msgpack(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(apptable.New(packets.NameBodyKind.Notes(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(apptable.New(packets.NameBodyKind.UnitTest(), "", unison.Yellow), -1)
	unison.Ancestor[*unison.DockContainer](yellow).Stack(apptable.New(packets.NameBodyKind.GitProxy(), "", unison.Yellow), -1)

	blue := apptable.New(packets.NameBodyKind.HttpDump(), "Response", unison.Pink)
	dock.DockTo(blue, nil, unison.BottomSide)
	unison.Ancestor[*unison.DockContainer](blue).Stack(apptable.New(packets.NameBodyKind.HexDump(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(apptable.New(packets.NameBodyKind.Steam(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(apptable.New(packets.NameBodyKind.ProtoBuf(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(apptable.New(packets.NameBodyKind.Tdf(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(apptable.New(packets.NameBodyKind.Taf(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(apptable.New(packets.NameBodyKind.Acc(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(apptable.New(packets.NameBodyKind.Websocket(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(apptable.New(packets.NameBodyKind.Msgpack(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(apptable.New(packets.NameBodyKind.Notes(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(apptable.New(packets.NameBodyKind.UnitTest(), "", unison.Pink), -1)
	unison.Ancestor[*unison.DockContainer](blue).Stack(apptable.New(packets.NameBodyKind.GitProxy(), "", unison.Pink), -1)

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
