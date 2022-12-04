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
	cleanerKind
	resourceFlashIconKind
	resourceJsIconKind
	resourceNotModifiedIconKind
	resourceRedirectIconKind
	submit7Kind
	replayKind
	resourceImageIconKind
	ssl2Kind
	ssl5Kind
	ssl6Kind
	aboutKind
	recKind
	resourceDnsIconKind
	resourcePlainIconKind
	ssl4Kind
	vpnKind
	resourceJavaIconKind
	editKind
	resourceDocumentIconKind
	settingKind
	mitmKind
	resourceCssIconKind
	resourceWebSocketIconKind
	rootCaKind
	ssl3Kind
	resourceExecutableIconKind
	resourceTcpIconKind
)

func (k kind) String() string {
	switch k {
	case cleanerKind:
		return "cleanerKind"
	case resourceFlashIconKind:
		return "resourceFlashIconKind"
	case resourceJsIconKind:
		return "resourceJsIconKind"
	case resourceNotModifiedIconKind:
		return "resourceNotModifiedIconKind"
	case resourceRedirectIconKind:
		return "resourceRedirectIconKind"
	case submit7Kind:
		return "submit7Kind"
	case replayKind:
		return "replayKind"
	case resourceImageIconKind:
		return "resourceImageIconKind"
	case ssl2Kind:
		return "ssl2Kind"
	case ssl5Kind:
		return "ssl5Kind"
	case ssl6Kind:
		return "ssl6Kind"
	case aboutKind:
		return "aboutKind"
	case recKind:
		return "recKind"
	case resourceDnsIconKind:
		return "resourceDnsIconKind"
	case resourcePlainIconKind:
		return "resourcePlainIconKind"
	case ssl4Kind:
		return "ssl4Kind"
	case vpnKind:
		return "vpnKind"
	case resourceJavaIconKind:
		return "resourceJavaIconKind"
	case editKind:
		return "editKind"
	case resourceDocumentIconKind:
		return "resourceDocumentIconKind"
	case settingKind:
		return "settingKind"
	case mitmKind:
		return "mitmKind"
	case resourceCssIconKind:
		return "resourceCssIconKind"
	case resourceWebSocketIconKind:
		return "resourceWebSocketIconKind"
	case rootCaKind:
		return "rootCaKind"
	case ssl3Kind:
		return "ssl3Kind"
	case resourceExecutableIconKind:
		return "resourceExecutableIconKind"
	case resourceTcpIconKind:
		return "resourceTcpIconKind"

	}
	return "invalid bmp kind"
}

func (k kind) Image() *unison.Image {
	switch k {
	case ssl3Kind:
		return widget.MustImage(pngSsl3)
	case mitmKind:
		return widget.MustImage(pngMitm)
	case resourceCssIconKind:
		return widget.MustImage(pngResourceCssIcon)
	case resourceWebSocketIconKind:
		return widget.MustImage(pngResourceWebSocketIcon)
	case rootCaKind:
		return widget.MustImage(pngRootCa)
	case resourceExecutableIconKind:
		return widget.MustImage(pngResourceExecutableIcon)
	case resourceTcpIconKind:
		return widget.MustImage(pngResourceTcpIcon)
	case resourceRedirectIconKind:
		return widget.MustImage(pngResourceRedirectIcon)
	case submit7Kind:
		return widget.MustImage(pngSubmit7)
	case cleanerKind:
		return widget.MustImage(pngCleaner)
	case resourceFlashIconKind:
		return widget.MustImage(pngResourceFlashIcon)
	case resourceJsIconKind:
		return widget.MustImage(pngResourceJsIcon)
	case resourceNotModifiedIconKind:
		return widget.MustImage(pngResourceNotModifiedIcon)
	case ssl6Kind:
		return widget.MustImage(pngSsl6)
	case replayKind:
		return widget.MustImage(pngReplay)
	case resourceImageIconKind:
		return widget.MustImage(pngResourceImageIcon)
	case ssl2Kind:
		return widget.MustImage(pngSsl2)
	case ssl5Kind:
		return widget.MustImage(pngSsl5)
	case ssl4Kind:
		return widget.MustImage(pngSsl4)
	case vpnKind:
		return widget.MustImage(pngVpn)
	case aboutKind:
		return widget.MustImage(pngAbout)
	case recKind:
		return widget.MustImage(pngRec)
	case resourceDnsIconKind:
		return widget.MustImage(pngResourceDnsIcon)
	case resourcePlainIconKind:
		return widget.MustImage(pngResourcePlainIcon)
	case resourceJavaIconKind:
		return widget.MustImage(pngResourceJavaIcon)
	case editKind:
		return widget.MustImage(pngEdit)
	case resourceDocumentIconKind:
		return widget.MustImage(pngResourceDocumentIcon)
	case settingKind:
		return widget.MustImage(pngSetting)

	}
	return nil
}

var (
	//go:embed asserts/resourceFlashIcon.ico
	pngResourceFlashIcon []byte

	//go:embed asserts/resourceJSIcon.ico
	pngResourceJsIcon []byte

	//go:embed asserts/resourceNotModifiedIcon.ico
	pngResourceNotModifiedIcon []byte

	//go:embed asserts/resourceRedirectIcon.ico
	pngResourceRedirectIcon []byte

	//go:embed asserts/submit7.ico
	pngSubmit7 []byte

	//go:embed asserts/cleaner.ico
	pngCleaner []byte

	//go:embed asserts/resourceImageIcon.ico
	pngResourceImageIcon []byte

	//go:embed asserts/ssl2.ico
	pngSsl2 []byte

	//go:embed asserts/ssl5.ico
	pngSsl5 []byte

	//go:embed asserts/ssl6.ico
	pngSsl6 []byte

	//go:embed asserts/replay.ico
	pngReplay []byte

	//go:embed asserts/rec.ico
	pngRec []byte

	//go:embed asserts/resourceDnsIcon.ico
	pngResourceDnsIcon []byte

	//go:embed asserts/resourcePlainIcon.ico
	pngResourcePlainIcon []byte

	//go:embed asserts/ssl4.ico
	pngSsl4 []byte

	//go:embed asserts/vpn.ico
	pngVpn []byte

	//go:embed asserts/about.ico
	pngAbout []byte

	//go:embed asserts/resourceJavaIcon.ico
	pngResourceJavaIcon []byte

	//go:embed asserts/resourceDocumentIcon.ico
	pngResourceDocumentIcon []byte

	//go:embed asserts/setting.ico
	pngSetting []byte

	//go:embed asserts/edit.ico
	pngEdit []byte

	//go:embed asserts/resourceCSSIcon.ico
	pngResourceCssIcon []byte

	//go:embed asserts/resourceWebSocketIcon.ico
	pngResourceWebSocketIcon []byte

	//go:embed asserts/rootCA.ico
	pngRootCa []byte

	//go:embed asserts/ssl3.ico
	pngSsl3 []byte

	//go:embed asserts/mitm.ico
	pngMitm []byte

	//go:embed asserts/resourceTcpIcon.ico
	pngResourceTcpIcon []byte

	//go:embed asserts/resourceExecutableIcon.ico
	pngResourceExecutableIcon []byte
)

