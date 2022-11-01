package main

import "github.com/gen2brain/iup-go/iup"

func albumPresenceFrame() iup.Ihandle {
	return iup.Frame(
		iup.Hbox(
			iup.Vbox(
				iup.Toggle("Show Album Art").SetAttribute("VALUE", trueState(config.preview.enabled)),
				iup.Toggle("Hover to Show Information").SetAttribute("VALUE", trueState(config.preview.albumConfig.albumDefaultHoverEnabled)),
				iup.Text().SetAttributes(`VALUE="`+config.preview.albumConfig.albumDefaultHoverText+`", SIZE=`+displayOptionsSize),
				iup.Toggle("Small Image").SetAttribute("VALUE", trueState(config.preview.smallImageConfig.enabled)),
				iup.Frame(
					iup.Vbox(
						iup.Frame(
							iup.Text().SetAttributes(`VALUE="`+config.preview.smallImageConfig.smallImageDefaultKey+`", SIZE=`+displayOptionsSize),
						).SetAttributes(`TITLE="Small Image Key"`),
						iup.Toggle("Enable Hover").SetAttribute("VALUE", trueState(config.preview.smallImageConfig.smallImageHoverEnabled)),
						iup.Frame(
							iup.Text().SetAttributes(`VALUE="`+config.preview.smallImageConfig.smallImageHoverText+`", SIZE=`+displayOptionsSize),
						).SetAttributes(`TITLE="Hover Text"`),
						iup.Toggle("Enable Loved").SetAttribute("VALUE", trueState(config.preview.smallImageConfig.lovedEnabled)),
						iup.Frame(
							iup.Text().SetAttributes(`VALUE="`+config.preview.smallImageConfig.lovedKey+`", SIZE=`+displayOptionsSize),
						).SetAttributes(`TITLE="Loved Image Key"`),
					),
				).SetAttributes(`TITLE="Small Image Settings", MARGIN="4x5", GAP=5`),
				iup.Fill(),
			),
			iup.Fill(),
		),
	).SetAttributes(`TITLE="Album Art Settings"`)
}
