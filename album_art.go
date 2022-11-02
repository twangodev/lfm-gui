package main

import "github.com/gen2brain/iup-go/iup"

func albumPresenceFrame() iup.Ihandle {
	return iup.Frame(
		iup.Hbox(
			iup.Vbox(
				iup.Toggle("Show Album Art").SetAttribute("VALUE", trueState(config.preview.enabled)),
				iup.Toggle("Hover to Show Information").SetAttribute("VALUE", trueState(config.preview.albumConfig.albumDefaultHoverEnabled)),
				iup.Toggle("Small Image").SetAttribute("VALUE", trueState(config.preview.smallImageConfig.enabled)),
				iup.Frame(
					iup.Vbox(
						iup.Toggle("Enable Hover").SetAttribute("VALUE", trueState(config.preview.smallImageConfig.smallImageHoverEnabled)),
						iup.Toggle("Enable Loved").SetAttribute("VALUE", trueState(config.preview.smallImageConfig.lovedEnabled)),
					),
				).SetAttributes(`TITLE="Small Image Settings", MARGIN="4x5", GAP=5`),
				iup.Fill(),
			),
			iup.Fill(),
		),
	).SetAttributes(`TITLE="Album Art Settings"`)
}
