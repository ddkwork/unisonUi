package widget

import (
	"github.com/ddkwork/unisonUi/widget/apptable"
	"github.com/ddkwork/unisonUi/widget/doctable"
	"github.com/richardwilkes/unison"
)

type (
	Interface interface {
		apptable.Interface
		doctable.Interface
		NewMenus(window *unison.Window, initializer func(unison.Menu))
		NewPopMenus(window *unison.Window, initializer func(unison.Menu))
		OpenWith()
		DropFiles()
	}
	object struct{}
)
