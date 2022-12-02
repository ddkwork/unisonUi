package ux

import (
	"os"
	"path/filepath"

	"github.com/richardwilkes/gcs/v5/model"
	"github.com/richardwilkes/unison"
)

type traitListProvider struct {
	traits []*model.Trait
}

func (p *traitListProvider) Entity() *model.Entity {
	return nil
}

func (p *traitListProvider) TraitList() []*model.Trait {
	return p.traits
}

func (p *traitListProvider) SetTraitList(list []*model.Trait) {
	p.traits = list
}

// NewTraitTableDockableFromFile loads a list of traits from a file and creates a new unison.Dockable for them.
func NewTraitTableDockableFromFile(filePath string) (unison.Dockable, error) {
	traits, err := model.NewTraitsFromFile(os.DirFS(filepath.Dir(filePath)), filepath.Base(filePath))
	if err != nil {
		return nil, err
	}
	d := NewTraitTableDockable(filePath, traits)
	d.needsSaveAsPrompt = false
	return d, nil
}

// NewTraitTableDockable creates a new unison.Dockable for trait list files.
func NewTraitTableDockable(filePath string, traits []*model.Trait) *TableDockable[*model.Trait] {
	provider := &traitListProvider{traits: traits}
	return NewTableDockable(filePath, model.TraitsExt, NewTraitsProvider(provider, false),
		func(path string) error { return model.SaveTraits(provider.TraitList(), path) },
		NewTraitItemID, NewTraitContainerItemID)
}
