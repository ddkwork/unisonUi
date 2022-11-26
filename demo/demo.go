package demo

import (
	"github.com/richardwilkes/gcs/v5/model"
	"github.com/richardwilkes/gcs/v5/ux"
	"github.com/richardwilkes/unison"
)

type (
	Interface interface {
		App()
		ToolBar()
		Layout()
		Entry()
		Button()
		Dialog()
		Table()
		DocTable()
		List()
		Tree()
		PopMeanu()
	}
	object struct {
		provider ux.TableProvider[*model.Weapon]
	}
)

func (o *object) PopMeanu() {
	//TODO implement me
	panic("implement me")
}

func (o *object) Layout() {
	//TODO implement me
	panic("implement me")
}

func (o *object) App() {
	//TODO implement me
	panic("implement me")
}

func New() Interface {
	return &object{
		provider: ux.NewWeaponsProvider(p, p.weaponType, false),
	}
}

func (o *object) ToolBar() {
	//TODO implement me
	panic("implement me")
}

func (o *object) Entry() {
	//TODO implement me
	panic("implement me")
}

func (o *object) Button() {
	//TODO implement me
	panic("implement me")
}

func (o *object) Dialog() {
	//TODO implement me
	panic("implement me")
}

func (o *object) Table() {
	panic("implement me")
}
func NewEditorTable[T model.NodeTypes](
	cmdRoot ux.Rebuildable,
	parent *unison.Panel,
	provider ux.TableProvider[T],
) *unison.Table[*ux.Node[T]] {
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

	//todo debug
	table.AsPanel().InstallCmdHandlers(ux.NewMeleeWeaponItemID, unison.AlwaysEnabled,
		func(_ any) { provider.CreateItem(cmdRoot, table, ux.NoItemVariant) })

	ux.InstallTableDropSupport(table, provider)
	table.SyncToModel()
	parent.AddChild(header)
	parent.AddChild(table)
	return table
}

func (o *object) DocTable() {
	//TODO implement me
	panic("implement me")
}

func (o *object) List() {
	//TODO implement me
	panic("implement me")
}

func (o *object) Tree() {
	//TODO implement me
	panic("implement me")
}
