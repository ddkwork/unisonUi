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
	resourceCssIconKind
	resourceExecutableIconKind
	resourceFlashIconKind
	rootCaKind
	settingKind
	aboutKind
	resourceJsIconKind
	resourceNotModifiedIconKind
	resourcePlainIconKind
	resourceDocumentIconKind
	resourceWebSocketIconKind
	submit7Kind
	recKind
	resourceDnsIconKind
	resourceJavaIconKind
	ssl3Kind
	ssl5Kind
	ssl6Kind
	vpnKind
	resourceImageIconKind
	ssl4Kind
	resourceTcpIconKind
	cleanerKind
	mitmKind
	resourceRedirectIconKind
	editKind
	replayKind
	ssl2Kind
)

func (k kind) String() string {
	switch k {
	case rootCaKind:
		return "rootCaKind"
	case settingKind:
		return "settingKind"
	case aboutKind:
		return "aboutKind"
	case resourceJsIconKind:
		return "resourceJsIconKind"
	case resourceNotModifiedIconKind:
		return "resourceNotModifiedIconKind"
	case resourcePlainIconKind:
		return "resourcePlainIconKind"
	case resourceDocumentIconKind:
		return "resourceDocumentIconKind"
	case resourceWebSocketIconKind:
		return "resourceWebSocketIconKind"
	case submit7Kind:
		return "submit7Kind"
	case vpnKind:
		return "vpnKind"
	case recKind:
		return "recKind"
	case resourceDnsIconKind:
		return "resourceDnsIconKind"
	case resourceJavaIconKind:
		return "resourceJavaIconKind"
	case ssl3Kind:
		return "ssl3Kind"
	case ssl5Kind:
		return "ssl5Kind"
	case ssl6Kind:
		return "ssl6Kind"
	case resourceImageIconKind:
		return "resourceImageIconKind"
	case ssl4Kind:
		return "ssl4Kind"
	case resourceRedirectIconKind:
		return "resourceRedirectIconKind"
	case resourceTcpIconKind:
		return "resourceTcpIconKind"
	case cleanerKind:
		return "cleanerKind"
	case mitmKind:
		return "mitmKind"
	case editKind:
		return "editKind"
	case replayKind:
		return "replayKind"
	case ssl2Kind:
		return "ssl2Kind"
	case resourceCssIconKind:
		return "resourceCssIconKind"
	case resourceExecutableIconKind:
		return "resourceExecutableIconKind"
	case resourceFlashIconKind:
		return "resourceFlashIconKind"

	}
	return "invalid bmp kind"
}

func (k kind) Image() *unison.Image {
	switch k {
	case resourceImageIconKind:
		return widget.MustImage(PngResourceImageIcon)
	case ssl4Kind:
		return widget.MustImage(PngSsl4)
	case resourceRedirectIconKind:
		return widget.MustImage(PngResourceRedirectIcon)
	case resourceTcpIconKind:
		return widget.MustImage(PngResourceTcpIcon)
	case cleanerKind:
		return widget.MustImage(PngCleaner)
	case mitmKind:
		return widget.MustImage(PngMitm)
	case editKind:
		return widget.MustImage(PngEdit)
	case replayKind:
		return widget.MustImage(PngReplay)
	case ssl2Kind:
		return widget.MustImage(PngSsl2)
	case resourceCssIconKind:
		return widget.MustImage(PngResourceCssIcon)
	case resourceExecutableIconKind:
		return widget.MustImage(PngResourceExecutableIcon)
	case resourceFlashIconKind:
		return widget.MustImage(PngResourceFlashIcon)
	case rootCaKind:
		return widget.MustImage(PngRootCa)
	case settingKind:
		return widget.MustImage(PngSetting)
	case resourcePlainIconKind:
		return widget.MustImage(PngResourcePlainIcon)
	case aboutKind:
		return widget.MustImage(PngAbout)
	case resourceJsIconKind:
		return widget.MustImage(PngResourceJsIcon)
	case resourceNotModifiedIconKind:
		return widget.MustImage(PngResourceNotModifiedIcon)
	case resourceDocumentIconKind:
		return widget.MustImage(PngResourceDocumentIcon)
	case resourceWebSocketIconKind:
		return widget.MustImage(PngResourceWebSocketIcon)
	case submit7Kind:
		return widget.MustImage(PngSubmit7)
	case ssl6Kind:
		return widget.MustImage(PngSsl6)
	case vpnKind:
		return widget.MustImage(PngVpn)
	case recKind:
		return widget.MustImage(PngRec)
	case resourceDnsIconKind:
		return widget.MustImage(PngResourceDnsIcon)
	case resourceJavaIconKind:
		return widget.MustImage(PngResourceJavaIcon)
	case ssl3Kind:
		return widget.MustImage(PngSsl3)
	case ssl5Kind:
		return widget.MustImage(PngSsl5)

	}
	return nil
}

var (
	//go:embed asserts/ssl6.ico
	PngSsl6 []byte

	//go:embed asserts/vpn.ico
	PngVpn []byte

	//go:embed asserts/rec.ico
	PngRec []byte

	//go:embed asserts/resourceDnsIcon.ico
	PngResourceDnsIcon []byte

	//go:embed asserts/resourceJavaIcon.ico
	PngResourceJavaIcon []byte

	//go:embed asserts/ssl3.ico
	PngSsl3 []byte

	//go:embed asserts/ssl5.ico
	PngSsl5 []byte

	//go:embed asserts/resourceImageIcon.ico
	PngResourceImageIcon []byte

	//go:embed asserts/ssl4.ico
	PngSsl4 []byte

	//go:embed asserts/resourceRedirectIcon.ico
	PngResourceRedirectIcon []byte

	//go:embed asserts/resourceTcpIcon.ico
	PngResourceTcpIcon []byte

	//go:embed asserts/cleaner.ico
	PngCleaner []byte

	//go:embed asserts/mitm.ico
	PngMitm []byte

	//go:embed asserts/edit.ico
	PngEdit []byte

	//go:embed asserts/replay.ico
	PngReplay []byte

	//go:embed asserts/ssl2.ico
	PngSsl2 []byte

	//go:embed asserts/resourceCSSIcon.ico
	PngResourceCssIcon []byte

	//go:embed asserts/resourceExecutableIcon.ico
	PngResourceExecutableIcon []byte

	//go:embed asserts/resourceFlashIcon.ico
	PngResourceFlashIcon []byte

	//go:embed asserts/rootCA.ico
	PngRootCa []byte

	//go:embed asserts/setting.ico
	PngSetting []byte

	//go:embed asserts/resourcePlainIcon.ico
	PngResourcePlainIcon []byte

	//go:embed asserts/about.ico
	PngAbout []byte

	//go:embed asserts/resourceJSIcon.ico
	PngResourceJsIcon []byte

	//go:embed asserts/resourceNotModifiedIcon.ico
	PngResourceNotModifiedIcon []byte

	//go:embed asserts/resourceDocumentIcon.ico
	PngResourceDocumentIcon []byte

	//go:embed asserts/resourceWebSocketIcon.ico
	PngResourceWebSocketIcon []byte

	//go:embed asserts/submit7.ico
	PngSubmit7 []byte
)

