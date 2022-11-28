package main

import "github.com/gen2brain/iup-go/iup"

func previewFrame() iup.Ihandle {
	return iup.Frame(
		iup.BackgroundBox(
			iup.Vbox(
				iup.Hbox(
					rpcPreviewFrame(),
				).SetAttribute("FGCOLOR", "255 255 255"),
				iup.Hbox(
					iup.Vbox(
						iup.BackgroundBox(
							iup.Hbox(
								iup.Fill(),
								iup.Label(config.buttons.profileButtonText).
									SetAttribute("FGCOLOR", "255 255 255").
									SetHandle("rpcPreviewProfileButtonText"),
								iup.Fill(),
							),
						).SetAttribute("BGCOLOR", "90 90 90"),
						iup.BackgroundBox(
							iup.Hbox(
								iup.Fill(),
								iup.Label(config.buttons.previewSongButtonText).
									SetAttribute("FGCOLOR", "255 255 255").SetHandle("rpcPreviewSongButtonText"),
								iup.Fill(),
							),
						).SetAttribute("BGCOLOR", "90 90 90"),
					),
				),
			),
		).SetAttribute("BGCOLOR", "54 57 63"),
	).SetAttributes("TITLE=Preview")

}
