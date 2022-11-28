package main

import (
	"github.com/gen2brain/iup-go/iup"
	"lfm-gui/concurrency"
)

var smallImageHandleBoolDependencyHandle = boolDependencyHandle{
	handle:       "smallImageHover",
	configOpcode: concurrency.DISP_ALBUM_SMALL_IMAGE_HOVER,
	ref:          &config.preview.smallImageConfig.smallImageHover,
}

var smallImageLovedBoolDependencyHandle = boolDependencyHandle{
	handle:       "smallImageLoved",
	configOpcode: concurrency.DISP_ALBUM_SMALL_IMAGE_LOVED,
	ref:          &config.preview.smallImageConfig.lovedEnabled,
}

func albumPresenceFrame() iup.Ihandle {
	return iup.Frame(
		iup.Hbox(
			iup.Vbox(
				iup.Toggle("Show Album Art").SetAttribute("VALUE", ooState(config.preview.enabled)).
					SetCallback("VALUECHANGED_CB",
						updateConfigBoolCallback(
							concurrency.DISP_ALBUM_ART_ENABLED,
							&config.preview.enabled,
							boolDependencyHandle{
								handle:       "albumHover",
								configOpcode: concurrency.DISP_ALBUM_ART_HOVER,
								ref:          &config.preview.albumConfig.hover,
							},
							boolDependencyHandle{
								handle:       "smallImage",
								configOpcode: concurrency.DISP_ALBUM_SMALL_IMAGE,
								ref:          &config.preview.smallImageConfig.enabled,
							}, smallImageHandleBoolDependencyHandle, smallImageLovedBoolDependencyHandle,
						),
					),
				iup.Toggle("Hover to Show Information").
					SetAttribute("VALUE", ooState(config.preview.albumConfig.hover)).
					SetCallback("VALUECHANGED_CB",
						updateConfigBoolCallback(concurrency.DISP_ALBUM_ART_HOVER,
							&config.preview.albumConfig.hover,
						),
					).SetHandle("albumHover"),
				iup.Toggle("Small Image").SetAttribute(
					"VALUE", ooState(config.preview.smallImageConfig.enabled),
				).SetHandle("smallImage").SetCallback("VALUECHANGED_CB",
					updateConfigBoolCallback(
						concurrency.DISP_ALBUM_SMALL_IMAGE,
						&config.preview.smallImageConfig.enabled,
						smallImageHandleBoolDependencyHandle,
						smallImageLovedBoolDependencyHandle,
					),
				),
				iup.Frame(
					iup.Vbox(
						iup.Toggle("Enable Hover").
							SetAttribute("VALUE", ooState(config.preview.smallImageConfig.smallImageHover)).
							SetCallback("VALUECHANGED_CB",
								updateConfigBoolCallback(concurrency.DISP_ALBUM_SMALL_IMAGE_HOVER, &config.preview.smallImageConfig.smallImageHover)).
							SetHandle("smallImageHover"),
						iup.Toggle("Enable Loved").
							SetAttribute("VALUE", ooState(config.preview.smallImageConfig.lovedEnabled)).
							SetCallback("VALUECHANGED_CB", updateConfigBoolCallback(concurrency.DISP_ALBUM_SMALL_IMAGE_LOVED, &config.preview.smallImageConfig.lovedEnabled)).
							SetHandle("smallImageLoved"),
					),
				).SetAttributes(`TITLE="Small Image Settings", MARGIN="4x5", GAP=5`),
				iup.Fill(),
			),
			iup.Fill(),
		),
	).SetAttributes(`TITLE="Album Art Settings"`)
}
