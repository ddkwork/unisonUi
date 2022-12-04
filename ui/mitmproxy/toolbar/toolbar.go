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
		CleanerButton() *unison.Button
		ResourceDnsIconButton() *unison.Button
		ResourceFlashIconButton() *unison.Button
		ResourceTcpIconButton() *unison.Button
		Cleaner2Button() *unison.Button
		Submit5Button() *unison.Button
		Submit7Button() *unison.Button
		Submit4Button() *unison.Button
		Submit6Button() *unison.Button
		ResourceJavaIconButton() *unison.Button
		ResourceNotModifiedIconButton() *unison.Button
		Ssl3Button() *unison.Button
		Submit2Button() *unison.Button
		ResourceDocumentIconButton() *unison.Button
		ResourceRedirectIconButton() *unison.Button
		SslButton() *unison.Button
		VpnButton() *unison.Button
		ResourceImageIconButton() *unison.Button
		OpenButton() *unison.Button
		SettingButton() *unison.Button
		Ssl4Button() *unison.Button
		Submit3Button() *unison.Button
		EditButton() *unison.Button
		ResourceCssIconButton() *unison.Button
		ResourcePlainIconButton() *unison.Button
		MitmButton() *unison.Button
		ResourceExecutableIconButton() *unison.Button
		ResourceWebSocketIconButton() *unison.Button
		Ssl2Button() *unison.Button
		ResourceJsIconButton() *unison.Button
		FaviconButton() *unison.Button
	}
	object struct {
		cleaner2Button                *unison.Button
		submit5Button                 *unison.Button
		submit7Button                 *unison.Button
		submit2Button                 *unison.Button
		submit4Button                 *unison.Button
		submit6Button                 *unison.Button
		resourceJavaIconButton        *unison.Button
		resourceNotModifiedIconButton *unison.Button
		ssl3Button                    *unison.Button
		vpnButton                     *unison.Button
		resourceDocumentIconButton    *unison.Button
		resourceRedirectIconButton    *unison.Button
		sslButton                     *unison.Button
		resourceImageIconButton       *unison.Button
		mitmButton                    *unison.Button
		openButton                    *unison.Button
		settingButton                 *unison.Button
		ssl4Button                    *unison.Button
		submit3Button                 *unison.Button
		editButton                    *unison.Button
		resourceCssIconButton         *unison.Button
		resourcePlainIconButton       *unison.Button
		resourceExecutableIconButton  *unison.Button
		resourceWebSocketIconButton   *unison.Button
		ssl2Button                    *unison.Button
		resourceJsIconButton          *unison.Button
		faviconButton                 *unison.Button
		resourceTcpIconButton         *unison.Button
		cleanerButton                 *unison.Button
		resourceDnsIconButton         *unison.Button
		resourceFlashIconButton       *unison.Button
	}
)

func New() Interface { return &object{} }

func (o *object) CanvasObject(window *unison.Window) *unison.Panel {
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlowLayout{
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})
	o.resourceJsIconButton = widget.CreateImageButton(resourceJsIconKind.Image(), resourceJsIconKind.String(), panel)
	o.faviconButton = widget.CreateImageButton(faviconKind.Image(), faviconKind.String(), panel)
	o.cleanerButton = widget.CreateImageButton(cleanerKind.Image(), cleanerKind.String(), panel)
	o.resourceDnsIconButton = widget.CreateImageButton(resourceDnsIconKind.Image(), resourceDnsIconKind.String(), panel)
	o.resourceFlashIconButton = widget.CreateImageButton(resourceFlashIconKind.Image(), resourceFlashIconKind.String(), panel)
	o.resourceTcpIconButton = widget.CreateImageButton(resourceTcpIconKind.Image(), resourceTcpIconKind.String(), panel)
	o.cleaner2Button = widget.CreateImageButton(cleaner2Kind.Image(), cleaner2Kind.String(), panel)
	o.submit5Button = widget.CreateImageButton(submit5Kind.Image(), submit5Kind.String(), panel)
	o.submit7Button = widget.CreateImageButton(submit7Kind.Image(), submit7Kind.String(), panel)
	o.resourceJavaIconButton = widget.CreateImageButton(resourceJavaIconKind.Image(), resourceJavaIconKind.String(), panel)
	o.resourceNotModifiedIconButton = widget.CreateImageButton(resourceNotModifiedIconKind.Image(), resourceNotModifiedIconKind.String(), panel)
	o.ssl3Button = widget.CreateImageButton(ssl3Kind.Image(), ssl3Kind.String(), panel)
	o.submit2Button = widget.CreateImageButton(submit2Kind.Image(), submit2Kind.String(), panel)
	o.submit4Button = widget.CreateImageButton(submit4Kind.Image(), submit4Kind.String(), panel)
	o.submit6Button = widget.CreateImageButton(submit6Kind.Image(), submit6Kind.String(), panel)
	o.resourceDocumentIconButton = widget.CreateImageButton(resourceDocumentIconKind.Image(), resourceDocumentIconKind.String(), panel)
	o.resourceRedirectIconButton = widget.CreateImageButton(resourceRedirectIconKind.Image(), resourceRedirectIconKind.String(), panel)
	o.sslButton = widget.CreateImageButton(sslKind.Image(), sslKind.String(), panel)
	o.vpnButton = widget.CreateImageButton(vpnKind.Image(), vpnKind.String(), panel)
	o.resourceImageIconButton = widget.CreateImageButton(resourceImageIconKind.Image(), resourceImageIconKind.String(), panel)
	o.submit3Button = widget.CreateImageButton(submit3Kind.Image(), submit3Kind.String(), panel)
	o.editButton = widget.CreateImageButton(editKind.Image(), editKind.String(), panel)
	o.resourceCssIconButton = widget.CreateImageButton(resourceCssIconKind.Image(), resourceCssIconKind.String(), panel)
	o.resourcePlainIconButton = widget.CreateImageButton(resourcePlainIconKind.Image(), resourcePlainIconKind.String(), panel)
	o.mitmButton = widget.CreateImageButton(mitmKind.Image(), mitmKind.String(), panel)
	o.openButton = widget.CreateImageButton(openKind.Image(), openKind.String(), panel)
	o.settingButton = widget.CreateImageButton(settingKind.Image(), settingKind.String(), panel)
	o.ssl4Button = widget.CreateImageButton(ssl4Kind.Image(), ssl4Kind.String(), panel)
	o.resourceExecutableIconButton = widget.CreateImageButton(resourceExecutableIconKind.Image(), resourceExecutableIconKind.String(), panel)
	o.resourceWebSocketIconButton = widget.CreateImageButton(resourceWebSocketIconKind.Image(), resourceWebSocketIconKind.String(), panel)
	o.ssl2Button = widget.CreateImageButton(ssl2Kind.Image(), ssl2Kind.String(), panel)
	return panel
}

func (o *object) ResourceImageIconButton() *unison.Button      { return o.resourceImageIconButton }
func (o *object) EditButton() *unison.Button                   { return o.editButton }
func (o *object) ResourceCssIconButton() *unison.Button        { return o.resourceCssIconButton }
func (o *object) ResourcePlainIconButton() *unison.Button      { return o.resourcePlainIconButton }
func (o *object) MitmButton() *unison.Button                   { return o.mitmButton }
func (o *object) OpenButton() *unison.Button                   { return o.openButton }
func (o *object) SettingButton() *unison.Button                { return o.settingButton }
func (o *object) Ssl4Button() *unison.Button                   { return o.ssl4Button }
func (o *object) Submit3Button() *unison.Button                { return o.submit3Button }
func (o *object) ResourceExecutableIconButton() *unison.Button { return o.resourceExecutableIconButton }
func (o *object) ResourceWebSocketIconButton() *unison.Button  { return o.resourceWebSocketIconButton }
func (o *object) Ssl2Button() *unison.Button                   { return o.ssl2Button }
func (o *object) ResourceJsIconButton() *unison.Button         { return o.resourceJsIconButton }
func (o *object) FaviconButton() *unison.Button                { return o.faviconButton }
func (o *object) CleanerButton() *unison.Button                { return o.cleanerButton }
func (o *object) ResourceDnsIconButton() *unison.Button        { return o.resourceDnsIconButton }
func (o *object) ResourceFlashIconButton() *unison.Button      { return o.resourceFlashIconButton }
func (o *object) ResourceTcpIconButton() *unison.Button        { return o.resourceTcpIconButton }
func (o *object) Cleaner2Button() *unison.Button               { return o.cleaner2Button }
func (o *object) Submit5Button() *unison.Button                { return o.submit5Button }
func (o *object) Submit7Button() *unison.Button                { return o.submit7Button }
func (o *object) ResourceJavaIconButton() *unison.Button       { return o.resourceJavaIconButton }
func (o *object) ResourceNotModifiedIconButton() *unison.Button {
	return o.resourceNotModifiedIconButton
}
func (o *object) Ssl3Button() *unison.Button                 { return o.ssl3Button }
func (o *object) Submit2Button() *unison.Button              { return o.submit2Button }
func (o *object) Submit4Button() *unison.Button              { return o.submit4Button }
func (o *object) Submit6Button() *unison.Button              { return o.submit6Button }
func (o *object) ResourceDocumentIconButton() *unison.Button { return o.resourceDocumentIconButton }
func (o *object) ResourceRedirectIconButton() *unison.Button { return o.resourceRedirectIconButton }
func (o *object) SslButton() *unison.Button                  { return o.sslButton }
func (o *object) VpnButton() *unison.Button                  { return o.vpnButton }

