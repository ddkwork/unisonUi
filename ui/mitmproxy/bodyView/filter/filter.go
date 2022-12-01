package filter

import "github.com/richardwilkes/unison"

func createFilter() *unison.Panel {
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlexLayout{
		Columns:      10,
		HSpacing:     unison.StdHSpacing,
		VSpacing:     unison.StdVSpacing,
		EqualColumns: true,
		//HAlign:       unison.StartAlignment,
	})
	field := unison.NewField()
	field.SetLayoutData(&unison.FlexLayoutData{
		HSpan:  1,
		VSpan:  1,
		HAlign: unison.FillAlignment,
		VAlign: unison.MiddleAlignment,
		HGrab:  true,
	})
	field.Tooltip = unison.NewTooltipWithText("filter")
	field.ModifiedCallback = func() {
		println(field.Text())
	}
	panel.AddChild(field)

	spacer := &unison.Panel{}
	spacer.Self = spacer
	spacer.SetSizer(func(_ unison.Size) (min, pref, max unison.Size) {
		min.Width = 40
		pref.Width = 40
		max.Width = 40
		return
	})
	panel.AddChild(spacer)

	//createCheckBox("RequestCtx", unison.OffCheckState, panel)
	//createCheckBox("ResponseCtx", unison.OffCheckState, panel)
	//createCheckBox("RequestBody", unison.OffCheckState, panel)
	//createCheckBox("ResponseBody", unison.OffCheckState, panel)
	//createCheckBox("Notes", unison.OffCheckState, panel)
	//createCheckBox("Scheme", unison.OffCheckState, panel)
	return panel
}
