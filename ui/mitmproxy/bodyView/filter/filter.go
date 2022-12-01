package filter

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/unison"
)

func CreateFilter() *unison.Field { //todo redisgn
	filter := i18n.Text("Content Filter")
	filterField := unison.NewField()
	filterField.Watermark = filter
	filterField.Tooltip = unison.NewTooltipWithText(filter)
	//filterField.ModifiedCallback = applyFilter
	filterField.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.FillAlignment,
		VAlign: unison.MiddleAlignment,
		HGrab:  true,
	})
	return filterField

	//panel := unison.NewPanel()
	//panel.SetLayout(&unison.FlexLayout{
	//	Columns:      1,
	//	HSpacing:     unison.StdHSpacing,
	//	VSpacing:     unison.StdVSpacing,
	//	EqualColumns: true,
	//	//HAlign:       unison.StartAlignment,
	//})
	//field := unison.NewField()
	//field.SetLayoutData(&unison.FlexLayoutData{
	//	HSpan:  1,
	//	VSpan:  1,
	//	HAlign: unison.FillAlignment,
	//	VAlign: unison.MiddleAlignment,
	//	HGrab:  true,
	//})
	//field.Tooltip = unison.NewTooltipWithText("filter")
	//field.ModifiedCallback = func() {
	//	println(field.Text())
	//}
	//panel.AddChild(field)
	//
	//spacer := &unison.Panel{}
	//spacer.Self = spacer
	//spacer.SetSizer(func(_ unison.Size) (min, pref, max unison.Size) {
	//	min.Width = 140
	//	pref.Width = 140
	//	max.Width = 140
	//	return
	//})
	//panel.AddChild(spacer)
	//
	////createCheckBox("RequestCtx", unison.OffCheckState, panel)
	////createCheckBox("ResponseCtx", unison.OffCheckState, panel)
	////createCheckBox("RequestBody", unison.OffCheckState, panel)
	////createCheckBox("ResponseBody", unison.OffCheckState, panel)
	////createCheckBox("Notes", unison.OffCheckState, panel)
	////createCheckBox("Scheme", unison.OffCheckState, panel)
	//return panel
}
