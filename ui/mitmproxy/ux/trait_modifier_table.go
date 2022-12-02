package ux

import (
	"os"
	"path/filepath"

	"github.com/richardwilkes/gcs/v5/model"
	"github.com/richardwilkes/unison"
)

type traitModifierListProvider struct {
	modifiers []*model.TraitModifier
}

func (p *traitModifierListProvider) Entity() *model.Entity {
	return nil
}

func (p *traitModifierListProvider) TraitModifierList() []*model.TraitModifier {
	return p.modifiers
}

func (p *traitModifierListProvider) SetTraitModifierList(list []*model.TraitModifier) {
	p.modifiers = list
}

// NewTraitModifierTableDockableFromFile loads a list of trait modifiers from a file and creates a new
// unison.Dockable for them.
func NewTraitModifierTableDockableFromFile(filePath string) (unison.Dockable, error) {
	modifiers, err := model.NewTraitModifiersFromFile(os.DirFS(filepath.Dir(filePath)), filepath.Base(filePath))
	if err != nil {
		return nil, err
	}
	d := NewTraitModifierTableDockable(filePath, modifiers)
	d.needsSaveAsPrompt = false
	return d, nil
}

// NewTraitModifierTableDockable creates a new unison.Dockable for trait modifier list files.
func NewTraitModifierTableDockable(filePath string, modifiers []*model.TraitModifier) *TableDockable[*model.TraitModifier] {
	provider := &traitModifierListProvider{modifiers: modifiers}
	return NewTableDockable(filePath, model.TraitModifiersExt,
		NewTraitModifiersProvider(provider, false),
		func(path string) error { return model.SaveTraitModifiers(provider.TraitModifierList(), path) },
		NewTraitModifierItemID, NewTraitContainerModifierItemID)
}
