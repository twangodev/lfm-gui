package main

import (
	"fmt"
	"github.com/gen2brain/iup-go/iup"
)

func contentPresenceFrame() iup.Ihandle {
	return iup.Frame(
		iup.Hbox(
			iup.Vbox(
				iup.Toggle("Enable Row 1").SetAttribute("VALUE", trueState(config.rows.rowOneEnabled)),
				iup.Toggle("Enable Row 2").SetAttribute("VALUE", trueState(config.rows.rowTwoEnabled)),
				iup.Toggle("Show Track Progress").SetAttribute("VALUE", trueState(config.rows.timeElapsedEnabled)),
				iup.Frame(
					iup.Hbox(
						iup.Text().SetAttributes(`VALUE="`+fmt.Sprintf("%d", config.refreshTime)+`", SIZE=40x`),
						iup.Label("milliseconds"),
					),
				).SetAttributes(`TITLE="Update Interval", MARGIN="4x5", GAP=5`),
				iup.Fill(),
			),
			iup.Fill(),
		),
	).SetAttributes("TITLE=Content")
}
