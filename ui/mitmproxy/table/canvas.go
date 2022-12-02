package table

import (
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/unisonUi/packets"
	"github.com/ddkwork/unisonUi/ui/mitmproxy/bodyView"
	"github.com/ddkwork/unisonUi/ui/mitmproxy/bodyView/filter"
	"github.com/ddkwork/unisonUi/ui/mitmproxy/menus"
	"github.com/ddkwork/unisonUi/ui/mitmproxy/toolbar"
	"github.com/google/uuid"
	"github.com/richardwilkes/gcs/v5/ux"
	"github.com/richardwilkes/unison"
	"sync"
)

type (
	ui interface {
		CanvasObject(where unison.Point) (ok bool)
		Mitmproxy() (ok bool) //for test in main
		unison.TableRowData[*object]
		//Table() *unison.Table[*object]
		Header() []string
	}
	object struct {
		packets.Object //set col style
		packets        []packets.Object
		//table          *unison.Table[*object]
		parent      *object
		id          uuid.UUID
		treeIdOrSub string
		tips        string
		root        []*object
		branch      []*object
		//webSocket      []*object
		//tcp            []*object
		//udp            []*object
		checkbox     *unison.CheckBox
		container    bool
		open         bool
		doubleHeight bool
		sync.RWMutex
	}
)

func New() *object { return &object{} }
func CanvasObject(w *unison.Window) (ok bool) {
	menus.New(w)
	content := w.Content()
	content.SetLayout(&unison.FlexLayout{Columns: 1})
	content.AddChild(toolbar.CreateToolBar())
	n := "Thaumatology - RPM Advantage Modifiers.adm"
	noteTableDockableFromFile, err := ux.NewTraitModifierTableDockableFromFile(n)
	if !mylog.Error(err) {
		return
	}
	content.AddChild(noteTableDockableFromFile)
	content.AddChild(filter.CreateFilter())
	content.AddChild(bodyView.CreateBodyView())
	return true
}
