package ux

import (
	"github.com/richardwilkes/gcs/v5/model"
	"github.com/richardwilkes/gcs/v5/svg"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/unison"
)

// EditNote displays the editor for a note.
func EditNote(owner Rebuildable, note *model.Note) {
	displayEditor[*model.Note, *model.NoteEditData](owner, note, svg.GCSNotes, initNoteToolbar, initNoteEditor)
}

func adjustMarkdownThemeForPage(markdown *unison.Markdown) {
	markdown.Font = model.PageFieldPrimaryFont
	markdown.Foreground = &unison.IndirectInk{Target: model.OnPageColor}
	markdown.HeadingFont[0] = &unison.DynamicFont{Resolver: func() unison.FontDescriptor { return unison.DeriveMarkdownHeadingFont(markdown.Font, 1) }}
	markdown.HeadingFont[1] = &unison.DynamicFont{Resolver: func() unison.FontDescriptor { return unison.DeriveMarkdownHeadingFont(markdown.Font, 2) }}
	markdown.HeadingFont[2] = &unison.DynamicFont{Resolver: func() unison.FontDescriptor { return unison.DeriveMarkdownHeadingFont(markdown.Font, 3) }}
	markdown.HeadingFont[3] = &unison.DynamicFont{Resolver: func() unison.FontDescriptor { return unison.DeriveMarkdownHeadingFont(markdown.Font, 4) }}
	markdown.HeadingFont[4] = &unison.DynamicFont{Resolver: func() unison.FontDescriptor { return unison.DeriveMarkdownHeadingFont(markdown.Font, 5) }}
	markdown.HeadingFont[5] = &unison.DynamicFont{Resolver: func() unison.FontDescriptor { return unison.DeriveMarkdownHeadingFont(markdown.Font, 6) }}
	markdown.CodeBlockFont = &unison.DynamicFont{
		Resolver: func() unison.FontDescriptor {
			fd := unison.MonospacedFont.Font.Descriptor()
			fd.Size = markdown.Font.Size()
			return fd
		},
	}
	markdown.CodeBackground = model.PageStandoutColor
	markdown.OnCodeBackground = model.OnPageStandoutColor
}

func initNoteToolbar(_ *editor[*model.Note, *model.NoteEditData], toolbar *unison.Panel) {
	filler := unison.NewPanel()
	filler.SetLayoutData(&unison.FlexLayoutData{HGrab: true})
	toolbar.AddChild(filler)
	toolbar.AddChild(unison.NewLink(i18n.Text("Markdown Guide"), "", "md:Markdown Guide", unison.DefaultLinkTheme, HandleLink))
}

func initNoteEditor(e *editor[*model.Note, *model.NoteEditData], content *unison.Panel) func() {
	markdown := unison.NewMarkdown(true)
	markdown.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.FillAlignment,
		HGrab:  true,
	})
	adjustMarkdownThemeForPage(markdown)

	labelText := i18n.Text("Notes")
	label := NewFieldLeadingLabel(labelText)
	label.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.EndAlignment,
		VAlign: unison.StartAlignment,
	})
	label.SetBorder(unison.NewEmptyBorder(unison.Insets{Top: 3}))
	content.AddChild(label)
	field := NewMultiLineStringField(nil, "", labelText,
		func() string { return e.editorData.Text },
		func(value string) {
			e.editorData.Text = value
			markdown.SetContent(value, 0)
			content.MarkForLayoutAndRedraw()
			MarkModified(content)
		})
	field.AutoScroll = false
	fd := unison.MonospacedFont.Descriptor()
	fd.Size = field.Font.Size()
	field.Font = fd.Font()
	content.AddChild(field)

	addPageRefLabelAndField(content, &e.editorData.PageRef)

	label = unison.NewLabel()
	label.Text = i18n.Text("Markdown Preview")
	label.HAlign = unison.MiddleAlignment
	label.SetLayoutData(&unison.FlexLayoutData{
		HSpan:  2,
		HAlign: unison.FillAlignment,
		HGrab:  true,
	})
	label.SetBorder(
		unison.NewCompoundBorder(
			unison.NewLineBorder(unison.DividerColor, 0, unison.Insets{Bottom: 1}, false),
			unison.NewEmptyBorder(unison.Insets{Top: unison.StdVSpacing * 3}),
		),
	)
	content.AddChild(label)

	markdown.SetContent(e.editorData.Text, 0)

	markdownWrapper := unison.NewPanel()
	markdownWrapper.SetScale(1.33)
	markdownWrapper.SetLayout(&unison.FlexLayout{Columns: 1})
	markdownWrapper.SetLayoutData(&unison.FlexLayoutData{
		HSpan:  2,
		HAlign: unison.FillAlignment,
		HGrab:  true,
	})
	content.AddChild(markdownWrapper)

	markdownWrapper.AddChild(markdown)

	return nil
}
