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
		ResourceWebSocketIconButton() *unison.Button
		Ssl5Button() *unison.Button
		Submit4Button() *unison.Button
		EditButton() *unison.Button
		ReplayButton() *unison.Button
		ResourceJsIconButton() *unison.Button
		ResourceJavaIconButton() *unison.Button
		ResourcePlainIconButton() *unison.Button
		RootCaButton() *unison.Button
		Ssl2Button() *unison.Button
		Ssl4Button() *unison.Button
		CleanerButton() *unison.Button
		ResourceCssIconButton() *unison.Button
		Submit7Button() *unison.Button
		MitmButton() *unison.Button
		ResourceTcpIconButton() *unison.Button
		Submit2Button() *unison.Button
		Submit6Button() *unison.Button
		AboutButton() *unison.Button
		SubmitButton() *unison.Button
		ResourceRedirectIconButton() *unison.Button
		Ssl3Button() *unison.Button
		Submit3Button() *unison.Button
		ResourceDnsIconButton() *unison.Button
		ResourceDocumentIconButton() *unison.Button
		ResourceFlashIconButton() *unison.Button
		SettingButton() *unison.Button
		Ssl6Button() *unison.Button
		VpnButton() *unison.Button
		RecButton() *unison.Button
		ResourceExecutableIconButton() *unison.Button
		ResourceImageIconButton() *unison.Button
		ResourceNotModifiedIconButton() *unison.Button
	}
	object struct {
		submit4Button                 *unison.Button
		editButton                    *unison.Button
		replayButton                  *unison.Button
		resourceWebSocketIconButton   *unison.Button
		ssl5Button                    *unison.Button
		resourcePlainIconButton       *unison.Button
		rootCaButton                  *unison.Button
		ssl2Button                    *unison.Button
		ssl4Button                    *unison.Button
		cleanerButton                 *unison.Button
		resourceCssIconButton         *unison.Button
		resourceJsIconButton          *unison.Button
		resourceJavaIconButton        *unison.Button
		submit7Button                 *unison.Button
		mitmButton                    *unison.Button
		resourceTcpIconButton         *unison.Button
		aboutButton                   *unison.Button
		submitButton                  *unison.Button
		submit2Button                 *unison.Button
		submit6Button                 *unison.Button
		submit3Button                 *unison.Button
		resourceDnsIconButton         *unison.Button
		resourceDocumentIconButton    *unison.Button
		resourceRedirectIconButton    *unison.Button
		ssl3Button                    *unison.Button
		ssl6Button                    *unison.Button
		vpnButton                     *unison.Button
		recButton                     *unison.Button
		resourceExecutableIconButton  *unison.Button
		resourceFlashIconButton       *unison.Button
		settingButton                 *unison.Button
		resourceImageIconButton       *unison.Button
		resourceNotModifiedIconButton *unison.Button
	}
)

func New() Interface { return &object{} }

func (o *object) CanvasObject(window *unison.Window) *unison.Panel {
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlowLayout{
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})
	o.resourceImageIconButton = widget.CreateImageButton(resourceImageIconKind.Image(), resourceImageIconKind.String(), panel)
	o.resourceNotModifiedIconButton = widget.CreateImageButton(resourceNotModifiedIconKind.Image(), resourceNotModifiedIconKind.String(), panel)
	o.submit4Button = widget.CreateImageButton(submit4Kind.Image(), submit4Kind.String(), panel)
	o.editButton = widget.CreateImageButton(editKind.Image(), editKind.String(), panel)
	o.replayButton = widget.CreateImageButton(replayKind.Image(), replayKind.String(), panel)
	o.resourceWebSocketIconButton = widget.CreateImageButton(resourceWebSocketIconKind.Image(), resourceWebSocketIconKind.String(), panel)
	o.ssl5Button = widget.CreateImageButton(ssl5Kind.Image(), ssl5Kind.String(), panel)
	o.resourcePlainIconButton = widget.CreateImageButton(resourcePlainIconKind.Image(), resourcePlainIconKind.String(), panel)
	o.rootCaButton = widget.CreateImageButton(rootCaKind.Image(), rootCaKind.String(), panel)
	o.ssl2Button = widget.CreateImageButton(ssl2Kind.Image(), ssl2Kind.String(), panel)
	o.ssl4Button = widget.CreateImageButton(ssl4Kind.Image(), ssl4Kind.String(), panel)
	o.cleanerButton = widget.CreateImageButton(cleanerKind.Image(), cleanerKind.String(), panel)
	o.resourceCssIconButton = widget.CreateImageButton(resourceCssIconKind.Image(), resourceCssIconKind.String(), panel)
	o.resourceJsIconButton = widget.CreateImageButton(resourceJsIconKind.Image(), resourceJsIconKind.String(), panel)
	o.resourceJavaIconButton = widget.CreateImageButton(resourceJavaIconKind.Image(), resourceJavaIconKind.String(), panel)
	o.submit7Button = widget.CreateImageButton(submit7Kind.Image(), submit7Kind.String(), panel)
	o.mitmButton = widget.CreateImageButton(mitmKind.Image(), mitmKind.String(), panel)
	o.resourceTcpIconButton = widget.CreateImageButton(resourceTcpIconKind.Image(), resourceTcpIconKind.String(), panel)
	o.aboutButton = widget.CreateImageButton(aboutKind.Image(), aboutKind.String(), panel)
	o.submitButton = widget.CreateImageButton(submitKind.Image(), submitKind.String(), panel)
	o.submit2Button = widget.CreateImageButton(submit2Kind.Image(), submit2Kind.String(), panel)
	o.submit6Button = widget.CreateImageButton(submit6Kind.Image(), submit6Kind.String(), panel)
	o.submit3Button = widget.CreateImageButton(submit3Kind.Image(), submit3Kind.String(), panel)
	o.resourceDnsIconButton = widget.CreateImageButton(resourceDnsIconKind.Image(), resourceDnsIconKind.String(), panel)
	o.resourceDocumentIconButton = widget.CreateImageButton(resourceDocumentIconKind.Image(), resourceDocumentIconKind.String(), panel)
	o.resourceRedirectIconButton = widget.CreateImageButton(resourceRedirectIconKind.Image(), resourceRedirectIconKind.String(), panel)
	o.ssl3Button = widget.CreateImageButton(ssl3Kind.Image(), ssl3Kind.String(), panel)
	o.ssl6Button = widget.CreateImageButton(ssl6Kind.Image(), ssl6Kind.String(), panel)
	o.vpnButton = widget.CreateImageButton(vpnKind.Image(), vpnKind.String(), panel)
	o.recButton = widget.CreateImageButton(recKind.Image(), recKind.String(), panel)
	o.resourceExecutableIconButton = widget.CreateImageButton(resourceExecutableIconKind.Image(), resourceExecutableIconKind.String(), panel)
	o.resourceFlashIconButton = widget.CreateImageButton(resourceFlashIconKind.Image(), resourceFlashIconKind.String(), panel)
	o.settingButton = widget.CreateImageButton(settingKind.Image(), settingKind.String(), panel)
	return panel
}

func (o *object) Submit2Button() *unison.Button                { return o.submit2Button }
func (o *object) Submit6Button() *unison.Button                { return o.submit6Button }
func (o *object) AboutButton() *unison.Button                  { return o.aboutButton }
func (o *object) SubmitButton() *unison.Button                 { return o.submitButton }
func (o *object) ResourceRedirectIconButton() *unison.Button   { return o.resourceRedirectIconButton }
func (o *object) Ssl3Button() *unison.Button                   { return o.ssl3Button }
func (o *object) Submit3Button() *unison.Button                { return o.submit3Button }
func (o *object) ResourceDnsIconButton() *unison.Button        { return o.resourceDnsIconButton }
func (o *object) ResourceDocumentIconButton() *unison.Button   { return o.resourceDocumentIconButton }
func (o *object) ResourceFlashIconButton() *unison.Button      { return o.resourceFlashIconButton }
func (o *object) SettingButton() *unison.Button                { return o.settingButton }
func (o *object) Ssl6Button() *unison.Button                   { return o.ssl6Button }
func (o *object) VpnButton() *unison.Button                    { return o.vpnButton }
func (o *object) RecButton() *unison.Button                    { return o.recButton }
func (o *object) ResourceExecutableIconButton() *unison.Button { return o.resourceExecutableIconButton }
func (o *object) ResourceImageIconButton() *unison.Button      { return o.resourceImageIconButton }
func (o *object) ResourceNotModifiedIconButton() *unison.Button {
	return o.resourceNotModifiedIconButton
}
func (o *object) ResourceWebSocketIconButton() *unison.Button { return o.resourceWebSocketIconButton }
func (o *object) Ssl5Button() *unison.Button                  { return o.ssl5Button }
func (o *object) Submit4Button() *unison.Button               { return o.submit4Button }
func (o *object) EditButton() *unison.Button                  { return o.editButton }
func (o *object) ReplayButton() *unison.Button                { return o.replayButton }
func (o *object) ResourceJsIconButton() *unison.Button        { return o.resourceJsIconButton }
func (o *object) ResourceJavaIconButton() *unison.Button      { return o.resourceJavaIconButton }
func (o *object) ResourcePlainIconButton() *unison.Button     { return o.resourcePlainIconButton }
func (o *object) RootCaButton() *unison.Button                { return o.rootCaButton }
func (o *object) Ssl2Button() *unison.Button                  { return o.ssl2Button }
func (o *object) Ssl4Button() *unison.Button                  { return o.ssl4Button }
func (o *object) CleanerButton() *unison.Button               { return o.cleanerButton }
func (o *object) ResourceCssIconButton() *unison.Button       { return o.resourceCssIconButton }
func (o *object) Submit7Button() *unison.Button               { return o.submit7Button }
func (o *object) MitmButton() *unison.Button                  { return o.mitmButton }
func (o *object) ResourceTcpIconButton() *unison.Button       { return o.resourceTcpIconButton }

