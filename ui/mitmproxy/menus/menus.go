package menus

import (
	"github.com/ddkwork/golibrary/skiaLib/widget"
	"github.com/ddkwork/unisonUi/ui/mitmproxy/about"
	"github.com/richardwilkes/unison"
)

func New(wnd *unison.Window) {
	widget.NewMenus(wnd, func(m unison.Menu) {
		unison.InsertStdMenus(m, about.ShowAboutWindow, nil, nil)
		fileMenu := m.Menu(unison.FileMenuID)
		f := fileMenu.Factory()
		newMenu := f.NewMenu(0, "New…", nil)
		//newMenu.InsertItem(-1, unison.NewMenuItem(f))
		//newMenu.InsertItem(-1, NewWindowAction.NewMenuItem(f))
		//newMenu.InsertItem(-1, NewTableWindowAction.NewMenuItem(f))
		//newMenu.InsertItem(-1, NewDockWindowAction.NewMenuItem(f))
		fileMenu.InsertMenu(0, newMenu)
		//fileMenu.InsertItem(1, OpenAction.NewMenuItem(f))
	})
}
