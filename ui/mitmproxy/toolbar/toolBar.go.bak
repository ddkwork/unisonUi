package toolbar

import (
	"github.com/ddkwork/unisonUi/ui/mitmproxy/ico"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/unison"
)

func CreateToolBar() *unison.Panel {
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlowLayout{
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})

	// Load our home image, and if successful (we should be!), add two buttons with it, one enabled and one not.
	homeImg, err := ico.HomeImage()
	if err != nil {
		jot.Error(err)
	} else {
		ico.CreateImageButton(homeImg, "home_enabled", panel)
		ico.CreateImageButton(homeImg, "home_disabled", panel).SetEnabled(false)
	}

	// Load our logo image, and if successful (we should be!), add two buttons with it, one enabled and one not.
	var logoImg *unison.Image
	if logoImg, err = ico.ClassicAppleLogoImage(); err != nil {
		jot.Error(err)
	} else {
		ico.CreateImageButton(logoImg, "logo_enabled", panel)
		ico.CreateImageButton(logoImg, "logo_disabled", panel).SetEnabled(false)
	}

	if homeImg != nil && logoImg != nil {
		// Add spacer
		//spacer := &unison.Panel{}
		//spacer.Self = spacer
		//spacer.SetSizer(func(_ unison.Size) (min, pref, max unison.Size) {
		//	min.Width = 40
		//	pref.Width = 40
		//	max.Width = 40
		//	return
		//})
		//panel.AddChild(spacer)

		// Add some sticky buttons in a group with our images
		group := unison.NewGroup()
		first := ico.CreateImageButton(homeImg, "home_toggle", panel)
		first.Sticky = true
		group.Add(first.AsGroupPanel())
		second := ico.CreateImageButton(logoImg, "logo_toggle", panel)
		second.Sticky = true
		group.Add(second.AsGroupPanel())
		group.Select(first.AsGroupPanel())
	}
	return panel
}
