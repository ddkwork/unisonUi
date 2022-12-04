//Code generated from mapPather - DO NOT EDIT.

package toolbar

import (
	_ "embed"
	"github.com/ddkwork/golibrary/skiaLib/widget"
	"github.com/richardwilkes/unison"
)

type kind byte

var name kind = 0

const (
	invalidKind kind = iota
	resourceRedirectIconKind
	sslKind
	vpnKind
	resourceDocumentIconKind
	resourceImageIconKind
	resourceCssIconKind
	resourcePlainIconKind
	mitmKind
	openKind
	settingKind
	ssl4Kind
	submit3Kind
	editKind
	resourceWebSocketIconKind
	ssl2Kind
	resourceExecutableIconKind
	faviconKind
	resourceJsIconKind
	resourceDnsIconKind
	resourceFlashIconKind
	resourceTcpIconKind
	cleanerKind
	submit5Kind
	submit7Kind
	cleaner2Kind
	resourceNotModifiedIconKind
	ssl3Kind
	submit2Kind
	submit4Kind
	submit6Kind
	resourceJavaIconKind
)

func (k kind) String() string {
	switch k {
	case resourceJsIconKind:
		return "resourceJsIconKind"
	case faviconKind:
		return "faviconKind"
	case cleanerKind:
		return "cleanerKind"
	case resourceDnsIconKind:
		return "resourceDnsIconKind"
	case resourceFlashIconKind:
		return "resourceFlashIconKind"
	case resourceTcpIconKind:
		return "resourceTcpIconKind"
	case cleaner2Kind:
		return "cleaner2Kind"
	case submit5Kind:
		return "submit5Kind"
	case submit7Kind:
		return "submit7Kind"
	case resourceJavaIconKind:
		return "resourceJavaIconKind"
	case resourceNotModifiedIconKind:
		return "resourceNotModifiedIconKind"
	case ssl3Kind:
		return "ssl3Kind"
	case submit2Kind:
		return "submit2Kind"
	case submit4Kind:
		return "submit4Kind"
	case submit6Kind:
		return "submit6Kind"
	case resourceDocumentIconKind:
		return "resourceDocumentIconKind"
	case resourceRedirectIconKind:
		return "resourceRedirectIconKind"
	case sslKind:
		return "sslKind"
	case vpnKind:
		return "vpnKind"
	case resourceImageIconKind:
		return "resourceImageIconKind"
	case editKind:
		return "editKind"
	case resourceCssIconKind:
		return "resourceCssIconKind"
	case resourcePlainIconKind:
		return "resourcePlainIconKind"
	case mitmKind:
		return "mitmKind"
	case openKind:
		return "openKind"
	case settingKind:
		return "settingKind"
	case ssl4Kind:
		return "ssl4Kind"
	case submit3Kind:
		return "submit3Kind"
	case resourceExecutableIconKind:
		return "resourceExecutableIconKind"
	case resourceWebSocketIconKind:
		return "resourceWebSocketIconKind"
	case ssl2Kind:
		return "ssl2Kind"

	}
	return "invalid bmp kind"
}

func (k kind) Image() *unison.Image {
	switch k {
	case cleanerKind:
		return widget.MustImage(pngCleaner)
	case resourceDnsIconKind:
		return widget.MustImage(pngResourceDnsIcon)
	case resourceFlashIconKind:
		return widget.MustImage(pngResourceFlashIcon)
	case resourceTcpIconKind:
		return widget.MustImage(pngResourceTcpIcon)
	case cleaner2Kind:
		return widget.MustImage(pngCleaner2)
	case submit5Kind:
		return widget.MustImage(pngSubmit5)
	case submit7Kind:
		return widget.MustImage(pngSubmit7)
	case resourceJavaIconKind:
		return widget.MustImage(pngResourceJavaIcon)
	case resourceNotModifiedIconKind:
		return widget.MustImage(pngResourceNotModifiedIcon)
	case ssl3Kind:
		return widget.MustImage(pngSsl3)
	case submit2Kind:
		return widget.MustImage(pngSubmit2)
	case submit4Kind:
		return widget.MustImage(pngSubmit4)
	case submit6Kind:
		return widget.MustImage(pngSubmit6)
	case resourceDocumentIconKind:
		return widget.MustImage(pngResourceDocumentIcon)
	case resourceRedirectIconKind:
		return widget.MustImage(pngResourceRedirectIcon)
	case sslKind:
		return widget.MustImage(pngSsl)
	case vpnKind:
		return widget.MustImage(pngVpn)
	case resourceImageIconKind:
		return widget.MustImage(pngResourceImageIcon)
	case editKind:
		return widget.MustImage(pngEdit)
	case resourceCssIconKind:
		return widget.MustImage(pngResourceCssIcon)
	case resourcePlainIconKind:
		return widget.MustImage(pngResourcePlainIcon)
	case mitmKind:
		return widget.MustImage(pngMitm)
	case openKind:
		return widget.MustImage(pngOpen)
	case settingKind:
		return widget.MustImage(pngSetting)
	case ssl4Kind:
		return widget.MustImage(pngSsl4)
	case submit3Kind:
		return widget.MustImage(pngSubmit3)
	case resourceExecutableIconKind:
		return widget.MustImage(pngResourceExecutableIcon)
	case resourceWebSocketIconKind:
		return widget.MustImage(pngResourceWebSocketIcon)
	case ssl2Kind:
		return widget.MustImage(pngSsl2)
	case resourceJsIconKind:
		return widget.MustImage(pngResourceJsIcon)
	case faviconKind:
		return widget.MustImage(pngFavicon)

	}
	return nil
}

var (
	//go:embed asserts/submit5.jpeg
	pngSubmit5 []byte

	//go:embed asserts/submit7.png
	pngSubmit7 []byte

	//go:embed asserts/cleaner2.ico
	pngCleaner2 []byte

	//go:embed asserts/images/resourceNotModifiedIcon.png
	pngResourceNotModifiedIcon []byte

	//go:embed asserts/ssl3.png
	pngSsl3 []byte

	//go:embed asserts/submit2.png
	pngSubmit2 []byte

	//go:embed asserts/submit4.jpeg
	pngSubmit4 []byte

	//go:embed asserts/submit6.png
	pngSubmit6 []byte

	//go:embed asserts/images/resourceJavaIcon.png
	pngResourceJavaIcon []byte

	//go:embed asserts/images/resourceRedirectIcon.png
	pngResourceRedirectIcon []byte

	//go:embed asserts/ssl.jpeg
	pngSsl []byte

	//go:embed asserts/vpn.jpeg
	pngVpn []byte

	//go:embed asserts/images/chrome-devtools/resourceDocumentIcon.png
	pngResourceDocumentIcon []byte

	//go:embed asserts/images/resourceImageIcon.png
	pngResourceImageIcon []byte

	//go:embed asserts/images/chrome-devtools/resourceCSSIcon.png
	pngResourceCssIcon []byte

	//go:embed asserts/images/chrome-devtools/resourcePlainIcon.png
	pngResourcePlainIcon []byte

	//go:embed asserts/mitm.ico
	pngMitm []byte

	//go:embed asserts/open.ico
	pngOpen []byte

	//go:embed asserts/setting.png
	pngSetting []byte

	//go:embed asserts/ssl4.png
	pngSsl4 []byte

	//go:embed asserts/submit3.png
	pngSubmit3 []byte

	//go:embed asserts/edit.png
	pngEdit []byte

	//go:embed asserts/images/resourceWebSocketIcon.png
	pngResourceWebSocketIcon []byte

	//go:embed asserts/ssl2.jpeg
	pngSsl2 []byte

	//go:embed asserts/images/resourceExecutableIcon.png
	pngResourceExecutableIcon []byte

	//go:embed asserts/images/favicon.ico
	pngFavicon []byte

	//go:embed asserts/images/chrome-devtools/resourceJSIcon.png
	pngResourceJsIcon []byte

	//go:embed asserts/images/resourceDnsIcon.png
	pngResourceDnsIcon []byte

	//go:embed asserts/images/resourceFlashIcon.png
	pngResourceFlashIcon []byte

	//go:embed asserts/images/resourceTcpIcon.png
	pngResourceTcpIcon []byte

	//go:embed asserts/cleaner.png
	pngCleaner []byte
)

