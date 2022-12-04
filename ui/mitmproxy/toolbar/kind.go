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
	aboutKind
	submitKind
	submit2Kind
	submit6Kind
	submit3Kind
	resourceDnsIconKind
	resourceDocumentIconKind
	resourceRedirectIconKind
	ssl3Kind
	ssl6Kind
	vpnKind
	recKind
	resourceExecutableIconKind
	resourceFlashIconKind
	settingKind
	resourceImageIconKind
	resourceNotModifiedIconKind
	submit4Kind
	editKind
	replayKind
	resourceWebSocketIconKind
	ssl5Kind
	resourcePlainIconKind
	rootCaKind
	ssl2Kind
	ssl4Kind
	cleanerKind
	resourceCssIconKind
	resourceJsIconKind
	resourceJavaIconKind
	submit7Kind
	mitmKind
	resourceTcpIconKind
)

func (k kind) String() string {
	switch k {
	case resourceRedirectIconKind:
		return "resourceRedirectIconKind"
	case ssl3Kind:
		return "ssl3Kind"
	case submit3Kind:
		return "submit3Kind"
	case resourceDnsIconKind:
		return "resourceDnsIconKind"
	case resourceDocumentIconKind:
		return "resourceDocumentIconKind"
	case resourceFlashIconKind:
		return "resourceFlashIconKind"
	case settingKind:
		return "settingKind"
	case ssl6Kind:
		return "ssl6Kind"
	case vpnKind:
		return "vpnKind"
	case recKind:
		return "recKind"
	case resourceExecutableIconKind:
		return "resourceExecutableIconKind"
	case resourceImageIconKind:
		return "resourceImageIconKind"
	case resourceNotModifiedIconKind:
		return "resourceNotModifiedIconKind"
	case resourceWebSocketIconKind:
		return "resourceWebSocketIconKind"
	case ssl5Kind:
		return "ssl5Kind"
	case submit4Kind:
		return "submit4Kind"
	case editKind:
		return "editKind"
	case replayKind:
		return "replayKind"
	case resourceJsIconKind:
		return "resourceJsIconKind"
	case resourceJavaIconKind:
		return "resourceJavaIconKind"
	case resourcePlainIconKind:
		return "resourcePlainIconKind"
	case rootCaKind:
		return "rootCaKind"
	case ssl2Kind:
		return "ssl2Kind"
	case ssl4Kind:
		return "ssl4Kind"
	case cleanerKind:
		return "cleanerKind"
	case resourceCssIconKind:
		return "resourceCssIconKind"
	case submit7Kind:
		return "submit7Kind"
	case mitmKind:
		return "mitmKind"
	case resourceTcpIconKind:
		return "resourceTcpIconKind"
	case submit2Kind:
		return "submit2Kind"
	case submit6Kind:
		return "submit6Kind"
	case aboutKind:
		return "aboutKind"
	case submitKind:
		return "submitKind"

	}
	return "invalid bmp kind"
}

func (k kind) Image() *unison.Image {
	switch k {
	case mitmKind:
		return widget.MustImage(pngMitm)
	case resourceTcpIconKind:
		return widget.MustImage(pngResourceTcpIcon)
	case aboutKind:
		return widget.MustImage(pngAbout)
	case submitKind:
		return widget.MustImage(pngSubmit)
	case submit2Kind:
		return widget.MustImage(pngSubmit2)
	case submit6Kind:
		return widget.MustImage(pngSubmit6)
	case resourceDnsIconKind:
		return widget.MustImage(pngResourceDnsIcon)
	case resourceDocumentIconKind:
		return widget.MustImage(pngResourceDocumentIcon)
	case resourceRedirectIconKind:
		return widget.MustImage(pngResourceRedirectIcon)
	case ssl3Kind:
		return widget.MustImage(pngSsl3)
	case submit3Kind:
		return widget.MustImage(pngSubmit3)
	case vpnKind:
		return widget.MustImage(pngVpn)
	case recKind:
		return widget.MustImage(pngRec)
	case resourceExecutableIconKind:
		return widget.MustImage(pngResourceExecutableIcon)
	case resourceFlashIconKind:
		return widget.MustImage(pngResourceFlashIcon)
	case settingKind:
		return widget.MustImage(pngSetting)
	case ssl6Kind:
		return widget.MustImage(pngSsl6)
	case resourceImageIconKind:
		return widget.MustImage(pngResourceImageIcon)
	case resourceNotModifiedIconKind:
		return widget.MustImage(pngResourceNotModifiedIcon)
	case editKind:
		return widget.MustImage(pngEdit)
	case replayKind:
		return widget.MustImage(pngReplay)
	case resourceWebSocketIconKind:
		return widget.MustImage(pngResourceWebSocketIcon)
	case ssl5Kind:
		return widget.MustImage(pngSsl5)
	case submit4Kind:
		return widget.MustImage(pngSubmit4)
	case rootCaKind:
		return widget.MustImage(pngRootCa)
	case ssl2Kind:
		return widget.MustImage(pngSsl2)
	case ssl4Kind:
		return widget.MustImage(pngSsl4)
	case cleanerKind:
		return widget.MustImage(pngCleaner)
	case resourceCssIconKind:
		return widget.MustImage(pngResourceCssIcon)
	case resourceJsIconKind:
		return widget.MustImage(pngResourceJsIcon)
	case resourceJavaIconKind:
		return widget.MustImage(pngResourceJavaIcon)
	case resourcePlainIconKind:
		return widget.MustImage(pngResourcePlainIcon)
	case submit7Kind:
		return widget.MustImage(pngSubmit7)

	}
	return nil
}

var (
	//go:embed asserts/submit4.ico
	pngSubmit4 []byte

	//go:embed asserts/edit.ico
	pngEdit []byte

	//go:embed asserts/replay.ico
	pngReplay []byte

	//go:embed asserts/resourceWebSocketIcon.ico
	pngResourceWebSocketIcon []byte

	//go:embed asserts/ssl5.ico
	pngSsl5 []byte

	//go:embed asserts/resourcePlainIcon.ico
	pngResourcePlainIcon []byte

	//go:embed asserts/rootCA.ico
	pngRootCa []byte

	//go:embed asserts/ssl2.ico
	pngSsl2 []byte

	//go:embed asserts/ssl4.ico
	pngSsl4 []byte

	//go:embed asserts/cleaner.ico
	pngCleaner []byte

	//go:embed asserts/resourceCSSIcon.ico
	pngResourceCssIcon []byte

	//go:embed asserts/resourceJSIcon.ico
	pngResourceJsIcon []byte

	//go:embed asserts/resourceJavaIcon.ico
	pngResourceJavaIcon []byte

	//go:embed asserts/submit7.ico
	pngSubmit7 []byte

	//go:embed asserts/mitm.ico
	pngMitm []byte

	//go:embed asserts/resourceTcpIcon.ico
	pngResourceTcpIcon []byte

	//go:embed asserts/about.ico
	pngAbout []byte

	//go:embed asserts/submit.ico
	pngSubmit []byte

	//go:embed asserts/submit2.ico
	pngSubmit2 []byte

	//go:embed asserts/submit6.ico
	pngSubmit6 []byte

	//go:embed asserts/submit3.ico
	pngSubmit3 []byte

	//go:embed asserts/resourceDnsIcon.ico
	pngResourceDnsIcon []byte

	//go:embed asserts/resourceDocumentIcon.ico
	pngResourceDocumentIcon []byte

	//go:embed asserts/resourceRedirectIcon.ico
	pngResourceRedirectIcon []byte

	//go:embed asserts/ssl3.ico
	pngSsl3 []byte

	//go:embed asserts/ssl6.ico
	pngSsl6 []byte

	//go:embed asserts/vpn.ico
	pngVpn []byte

	//go:embed asserts/rec.ico
	pngRec []byte

	//go:embed asserts/resourceExecutableIcon.ico
	pngResourceExecutableIcon []byte

	//go:embed asserts/resourceFlashIcon.ico
	pngResourceFlashIcon []byte

	//go:embed asserts/setting.ico
	pngSetting []byte

	//go:embed asserts/resourceImageIcon.ico
	pngResourceImageIcon []byte

	//go:embed asserts/resourceNotModifiedIcon.ico
	pngResourceNotModifiedIcon []byte
)

