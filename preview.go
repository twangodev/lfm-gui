package main

import "github.com/gen2brain/iup-go/iup"

func previewFrame() iup.Ihandle {
	return iup.Frame(
		iup.BackgroundBox(
			iup.Hbox(
				iup.Fill().SetAttribute("SIZE", "50"),
				rpcPreviewFrame(),
				iup.Fill(),
			).SetAttribute("FGCOLOR", "255 255 255"),
		).SetAttribute("BGCOLOR", "54 57 63"),
	).SetAttributes("TITLE=Preview")

}
