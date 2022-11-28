package main

import (
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/unisonUi/asserts"
	"github.com/richardwilkes/gcs/v5/model"
	"github.com/richardwilkes/gcs/v5/ux"
	"github.com/richardwilkes/unison"
)

func main() {
	//ux.RegisterKnownFileTypes()
	//model.GlobalSettings() // Here to force early initialization
	//unison.DefaultScrollPanelTheme.MouseWheelMultiplier = func() float32 {
	//	return fxp.As[float32](model.GlobalSettings().General.ScrollWheelMultiplier)
	//}
	unison.Start(unison.StartupFinishedCallback(func() { CanvasObject(unison.PrimaryDisplay().Usable.Point) }))
}

func CanvasObject(where unison.Point) (ok bool) {
	w, err := unison.NewWindow(fmt.Sprintf("mitmproxy"))
	if err != nil {
		return
	}
	ux.SetupMenuBar(w)
	//ux.NewWorkspace(w)
	w.MinMaxContentSizeCallback = func() (min, max unison.Size) {
		return unison.NewSize(1000, 600), unison.NewSize(10000, 1280)
	}
	image, err := unison.NewImageFromBytes(asserts.MitmBuf, 0.5)
	if !mylog.Error(err) {
		return
	}
	w.SetTitleIcons([]*unison.Image{image})

	content := w.Content()
	content.SetLayout(&unison.FlexLayout{Columns: 1})

	//init table
	//traitTableDockable := ux.NewTraitTableDockable("Traits"+model.TraitsExt, nil)
	traitTableDockable := newEditorTable()
	content.AddChild(traitTableDockable)

	w.Pack()
	rect := w.FrameRect()
	rect.Point = where
	w.SetFrameRect(rect)
	w.ToFront()
	return true
}

func newEditorTable[T model.NodeTypes](parent *unison.Panel, provider ux.TableProvider[T]) *unison.Table[*ux.Node[T]] {
	header, table := ux.NewNodeTable[T](provider, unison.FieldFont)
	table.InstallCmdHandlers(ux.OpenEditorItemID, func(_ any) bool { return table.HasSelection() },
		func(_ any) { provider.OpenEditor(unison.AncestorOrSelf[ux.Rebuildable](table), table) })
	table.InstallCmdHandlers(ux.OpenOnePageReferenceItemID,
		func(_ any) bool { return ux.CanOpenPageRef(table) },
		func(_ any) { ux.OpenPageRef(table) })
	table.InstallCmdHandlers(ux.OpenEachPageReferenceItemID,
		func(_ any) bool { return ux.CanOpenPageRef(table) },
		func(_ any) { ux.OpenEachPageRef(table) })
	table.InstallCmdHandlers(unison.DeleteItemID,
		func(_ any) bool { return table.HasSelection() },
		func(_ any) { ux.DeleteSelection(table) })
	table.InstallCmdHandlers(ux.DuplicateItemID,
		func(_ any) bool { return table.HasSelection() },
		func(_ any) { ux.DuplicateSelection(table) })
	ux.InstallTableDropSupport(table, provider)
	table.SyncToModel()
	parent.AddChild(header)
	parent.AddChild(table)
	return table
}
