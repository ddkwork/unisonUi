package ico

import (
	_ "embed"
	"fmt"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/unison"
)

var (
	//go:embed resources/home.png
	homePngBytes []byte
	homeImage    *unison.Image

	//go:embed resources/classic-apple-logo.png
	classicAppleLogoPngBytes []byte
	classicAppleLogoImage    *unison.Image
)

func CreateImageButton(img *unison.Image, actionText string, panel *unison.Panel) *unison.Button {
	btn := unison.NewButton()
	btn.Drawable = img
	btn.ClickCallback = func() { jot.Info(actionText) }
	btn.Tooltip = unison.NewTooltipWithText(fmt.Sprintf("Tooltip for: %s", actionText))
	btn.SetLayoutData(unison.MiddleAlignment)
	panel.AddChild(btn)
	return btn
}

// ClassicAppleLogoImage returns an image of the classic rainbow-colored Apple logo.
func ClassicAppleLogoImage() (*unison.Image, error) {
	if classicAppleLogoImage == nil {
		var err error
		if classicAppleLogoImage, err = unison.NewImageFromBytes(classicAppleLogoPngBytes, 0.5); err != nil {
			return nil, err
		}
	}
	return classicAppleLogoImage, nil
}

// HomeImage returns a stylized image of a home, suitable for an icon.
func HomeImage() (*unison.Image, error) {
	if homeImage == nil {
		var err error
		if homeImage, err = unison.NewImageFromBytes(homePngBytes, 0.5); err != nil {
			return nil, err
		}
	}
	return homeImage, nil
}
