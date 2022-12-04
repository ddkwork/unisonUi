//Code generated from mapPather - DO NOT EDIT.

package toolbar

import (
	_ "embed"
	"github.com/ddkwork/golibrary/skiaLib/widget"
	"github.com/ddkwork/golibrary/skiaLib/widget/canvasobjectapi"
	"github.com/richardwilkes/unison"
)

type (
	Interface interface {
		canvasobjectapi.Interface
		Actions
	}
	Actions interface {
		AboutButton() *unison.Button
		RecButton() *unison.Button
		ResourceDnsIconButton() *unison.Button
		ResourcePlainIconButton() *unison.Button
		Ssl4Button() *unison.Button
		VpnButton() *unison.Button
		ResourceJavaIconButton() *unison.Button
		EditButton() *unison.Button
		ResourceDocumentIconButton() *unison.Button
		SettingButton() *unison.Button
		MitmButton() *unison.Button
		ResourceCssIconButton() *unison.Button
		ResourceWebSocketIconButton() *unison.Button
		RootCaButton() *unison.Button
		Ssl3Button() *unison.Button
		ResourceExecutableIconButton() *unison.Button
		ResourceTcpIconButton() *unison.Button
		CleanerButton() *unison.Button
		ResourceFlashIconButton() *unison.Button
		ResourceJsIconButton() *unison.Button
		ResourceNotModifiedIconButton() *unison.Button
		ResourceRedirectIconButton() *unison.Button
		Submit7Button() *unison.Button
		ReplayButton() *unison.Button
		ResourceImageIconButton() *unison.Button
		Ssl2Button() *unison.Button
		Ssl5Button() *unison.Button
		Ssl6Button() *unison.Button
	}
	object struct {
		cleanerButton                 *unison.Button
		resourceFlashIconButton       *unison.Button
		resourceJsIconButton          *unison.Button
		resourceNotModifiedIconButton *unison.Button
		resourceRedirectIconButton    *unison.Button
		submit7Button                 *unison.Button
		replayButton                  *unison.Button
		resourceImageIconButton       *unison.Button
		ssl2Button                    *unison.Button
		ssl5Button                    *unison.Button
		ssl6Button                    *unison.Button
		aboutButton                   *unison.Button
		recButton                     *unison.Button
		resourceDnsIconButton         *unison.Button
		resourcePlainIconButton       *unison.Button
		ssl4Button                    *unison.Button
		vpnButton                     *unison.Button
		resourceJavaIconButton        *unison.Button
		editButton                    *unison.Button
		resourceDocumentIconButton    *unison.Button
		settingButton                 *unison.Button
		mitmButton                    *unison.Button
		resourceCssIconButton         *unison.Button
		resourceWebSocketIconButton   *unison.Button
		rootCaButton                  *unison.Button
		ssl3Button                    *unison.Button
		resourceExecutableIconButton  *unison.Button
		resourceTcpIconButton         *unison.Button
	}
)

func New() Interface { return &object{} }

func (o *object) CanvasObject(window *unison.Window) *unison.Panel {
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlowLayout{
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})
	o.resourceJavaIconButton = widget.CreateImageButton(resourceJavaIconKind.Image(), resourceJavaIconKind.String(), panel)
	o.editButton = widget.CreateImageButton(editKind.Image(), editKind.String(), panel)
	o.resourceDocumentIconButton = widget.CreateImageButton(resourceDocumentIconKind.Image(), resourceDocumentIconKind.String(), panel)
	o.settingButton = widget.CreateImageButton(settingKind.Image(), settingKind.String(), panel)
	o.mitmButton = widget.CreateImageButton(mitmKind.Image(), mitmKind.String(), panel)
	o.resourceCssIconButton = widget.CreateImageButton(resourceCssIconKind.Image(), resourceCssIconKind.String(), panel)
	o.resourceWebSocketIconButton = widget.CreateImageButton(resourceWebSocketIconKind.Image(), resourceWebSocketIconKind.String(), panel)
	o.rootCaButton = widget.CreateImageButton(rootCaKind.Image(), rootCaKind.String(), panel)
	o.ssl3Button = widget.CreateImageButton(ssl3Kind.Image(), ssl3Kind.String(), panel)
	o.resourceExecutableIconButton = widget.CreateImageButton(resourceExecutableIconKind.Image(), resourceExecutableIconKind.String(), panel)
	o.resourceTcpIconButton = widget.CreateImageButton(resourceTcpIconKind.Image(), resourceTcpIconKind.String(), panel)
	o.cleanerButton = widget.CreateImageButton(cleanerKind.Image(), cleanerKind.String(), panel)
	o.resourceFlashIconButton = widget.CreateImageButton(resourceFlashIconKind.Image(), resourceFlashIconKind.String(), panel)
	o.resourceJsIconButton = widget.CreateImageButton(resourceJsIconKind.Image(), resourceJsIconKind.String(), panel)
	o.resourceNotModifiedIconButton = widget.CreateImageButton(resourceNotModifiedIconKind.Image(), resourceNotModifiedIconKind.String(), panel)
	o.resourceRedirectIconButton = widget.CreateImageButton(resourceRedirectIconKind.Image(), resourceRedirectIconKind.String(), panel)
	o.submit7Button = widget.CreateImageButton(submit7Kind.Image(), submit7Kind.String(), panel)
	o.replayButton = widget.CreateImageButton(replayKind.Image(), replayKind.String(), panel)
	o.resourceImageIconButton = widget.CreateImageButton(resourceImageIconKind.Image(), resourceImageIconKind.String(), panel)
	o.ssl2Button = widget.CreateImageButton(ssl2Kind.Image(), ssl2Kind.String(), panel)
	o.ssl5Button = widget.CreateImageButton(ssl5Kind.Image(), ssl5Kind.String(), panel)
	o.ssl6Button = widget.CreateImageButton(ssl6Kind.Image(), ssl6Kind.String(), panel)
	o.aboutButton = widget.CreateImageButton(aboutKind.Image(), aboutKind.String(), panel)
	o.recButton = widget.CreateImageButton(recKind.Image(), recKind.String(), panel)
	o.resourceDnsIconButton = widget.CreateImageButton(resourceDnsIconKind.Image(), resourceDnsIconKind.String(), panel)
	o.resourcePlainIconButton = widget.CreateImageButton(resourcePlainIconKind.Image(), resourcePlainIconKind.String(), panel)
	o.ssl4Button = widget.CreateImageButton(ssl4Kind.Image(), ssl4Kind.String(), panel)
	o.vpnButton = widget.CreateImageButton(vpnKind.Image(), vpnKind.String(), panel)
	return panel
}

func (o *object) ResourceNotModifiedIconButton() *unison.Button {
	return o.resourceNotModifiedIconButton
}
func (o *object) ResourceRedirectIconButton() *unison.Button   { return o.resourceRedirectIconButton }
func (o *object) Submit7Button() *unison.Button                { return o.submit7Button }
func (o *object) CleanerButton() *unison.Button                { return o.cleanerButton }
func (o *object) ResourceFlashIconButton() *unison.Button      { return o.resourceFlashIconButton }
func (o *object) ResourceJsIconButton() *unison.Button         { return o.resourceJsIconButton }
func (o *object) Ssl5Button() *unison.Button                   { return o.ssl5Button }
func (o *object) Ssl6Button() *unison.Button                   { return o.ssl6Button }
func (o *object) ReplayButton() *unison.Button                 { return o.replayButton }
func (o *object) ResourceImageIconButton() *unison.Button      { return o.resourceImageIconButton }
func (o *object) Ssl2Button() *unison.Button                   { return o.ssl2Button }
func (o *object) ResourcePlainIconButton() *unison.Button      { return o.resourcePlainIconButton }
func (o *object) Ssl4Button() *unison.Button                   { return o.ssl4Button }
func (o *object) VpnButton() *unison.Button                    { return o.vpnButton }
func (o *object) AboutButton() *unison.Button                  { return o.aboutButton }
func (o *object) RecButton() *unison.Button                    { return o.recButton }
func (o *object) ResourceDnsIconButton() *unison.Button        { return o.resourceDnsIconButton }
func (o *object) ResourceJavaIconButton() *unison.Button       { return o.resourceJavaIconButton }
func (o *object) EditButton() *unison.Button                   { return o.editButton }
func (o *object) ResourceDocumentIconButton() *unison.Button   { return o.resourceDocumentIconButton }
func (o *object) SettingButton() *unison.Button                { return o.settingButton }
func (o *object) RootCaButton() *unison.Button                 { return o.rootCaButton }
func (o *object) Ssl3Button() *unison.Button                   { return o.ssl3Button }
func (o *object) MitmButton() *unison.Button                   { return o.mitmButton }
func (o *object) ResourceCssIconButton() *unison.Button        { return o.resourceCssIconButton }
func (o *object) ResourceWebSocketIconButton() *unison.Button  { return o.resourceWebSocketIconButton }
func (o *object) ResourceExecutableIconButton() *unison.Button { return o.resourceExecutableIconButton }
func (o *object) ResourceTcpIconButton() *unison.Button        { return o.resourceTcpIconButton }

