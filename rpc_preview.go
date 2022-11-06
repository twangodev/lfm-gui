package main

import (
	"github.com/gen2brain/iup-go/iup"
	"image/png"
	"os"
)

const previewBoldFont = "Helvetica, Bold 10"

func rpcPreviewFrame() iup.Ihandle {

	previewFile, err := os.Open(config.preview.albumDefaultPreviewPath)
	if err != nil {
		return iup.Label("Error: " + err.Error())
	}
	previewImage, err := png.Decode(previewFile)
	if err != nil {
		return iup.Label("Error: " + err.Error())
	}

	iup.ImageFromImage(previewImage).SetHandle("previewImage")

	return iup.Hbox(
		iup.Label("").SetAttribute("IMAGE", "previewImage"),
		iup.Vbox(
			iup.Label("Playing last.fm").SetAttribute("FONT", previewBoldFont),
			iup.Label(config.rows.rowOneText),
			iup.Label(config.rows.rowTwoText),
			iup.Label("0:00 elapsed"),
		),
	)
}
