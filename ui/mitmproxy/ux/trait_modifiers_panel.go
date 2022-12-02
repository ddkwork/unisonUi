package ux

import (
	"github.com/google/uuid"
	"github.com/richardwilkes/gcs/v5/model"
	"github.com/richardwilkes/unison"
)

type traitModifiersPanel struct {
	unison.Panel
	entity    *model.Entity
	modifiers *[]*model.TraitModifier
	provider  TableProvider[*model.TraitModifier]
	table     *unison.Table[*Node[*model.TraitModifier]]
}

func newTraitModifiersPanel(entity *model.Entity, modifiers *[]*model.TraitModifier) *traitModifiersPanel {
	p := &traitModifiersPanel{
		entity:    entity,
		modifiers: modifiers,
	}
	p.Self = p
	p.SetLayout(&unison.FlexLayout{Columns: 1})
	p.SetLayoutData(&unison.FlexLayoutData{
		HSpan:  2,
		HAlign: unison.FillAlignment,
		HGrab:  true,
	})
	p.SetBorder(unison.NewLineBorder(model.HeaderColor, 0, unison.NewUniformInsets(1), false))
	p.DrawCallback = func(gc *unison.Canvas, rect unison.Rect) {
		gc.DrawRect(rect, unison.ContentColor.Paint(gc, rect, unison.Fill))
	}
	p.provider = NewTraitModifiersProvider(p, true)
	p.table = newEditorTable(p.AsPanel(), p.provider)
	p.table.RefKey = "trait-modifiers-" + uuid.New().String()
	return p
}

func (p *traitModifiersPanel) Entity() *model.Entity {
	return p.entity
}

func (p *traitModifiersPanel) TraitModifierList() []*model.TraitModifier {
	return *p.modifiers
}

func (p *traitModifiersPanel) SetTraitModifierList(list []*model.TraitModifier) {
	*p.modifiers = list
	sel := p.table.CopySelectionMap()
	p.table.SyncToModel()
	p.table.SetSelectionMap(sel)
}
