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
				iup.Frame(
					iup.Text().SetAttributes(`VALUE="`+config.rows.rowOne+`", SIZE=`+displayOptionsSize),
				).SetAttributes(`TITLE="Row 1 Settings"`),
				iup.Toggle("Enable Row 2").SetAttribute("VALUE", trueState(config.rows.rowTwoEnabled)),
				iup.Frame(
					iup.Text().SetAttributes(`VALUE="`+config.rows.rowTwo+`", SIZE=`+displayOptionsSize),
				).SetAttributes(`TITLE="Row 2 Settings"`),
				iup.Toggle("Show Track Progress").SetAttribute("VALUE", trueState(config.rows.timeElapsedEnabled)),
				iup.Frame(
					iup.Vbox(
						iup.Hbox(
							iup.Text().SetAttributes(`VALUE="`+fmt.Sprintf("%d", config.refreshTime)+`", SIZE=40x`),
							iup.Label("milliseconds"),
						),
					),
				).SetAttributes(`TITLE="Update Interval", MARGIN="4x5", GAP=5`),
				iup.Fill(),
			),
			iup.Fill(),
		),
	).SetAttributes("TITLE=Content")
}
