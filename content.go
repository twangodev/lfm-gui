package main

import (
	"fmt"
	"github.com/gen2brain/iup-go/iup"
	"lfm-gui/concurrency"
)

func contentPresenceFrame() iup.Ihandle {
	return iup.Frame(
		iup.Hbox(
			iup.Vbox(
				iup.Toggle("Enable Row 1").SetAttribute("VALUE", ooState(config.rows.rowOne)).
					SetCallback("VALUECHANGED_CB", updateConfigBoolCallback(concurrency.DISP_ROW_ONE, &config.rows.rowOne)),
				iup.Toggle("Enable Row 2").SetAttribute("VALUE", ooState(config.rows.rowTwo)).
					SetCallback("VALUECHANGED_CB", updateConfigBoolCallback(concurrency.DISP_ROW_TWO, &config.rows.rowTwo)),
				iup.Toggle("Show Track Progress").SetAttribute("VALUE", ooState(config.rows.timeElapsed)).
					SetCallback("VALUECHANGED_CB", updateConfigBoolCallback(concurrency.DISP_TIME_ELAPSED, &config.rows.timeElapsed)),
				iup.Frame(
					iup.Hbox(
						iup.Text().SetAttributes(`VALUE="`+fmt.Sprintf("%d", config.refreshTime)+`", SIZE=40x`).
							SetCallback("VALUECHANGED_CB",
								updateConfigIntCallback(concurrency.DISP_UPDATE_INTERVAL,
									&config.refreshTime,
									5000,
									&config.refreshTime,
								),
							),
						iup.Label("milliseconds"),
					),
				).SetAttributes(`TITLE="Update Interval", MARGIN="4x5", GAP=5`),
				iup.Fill(),
			),
			iup.Fill(),
		),
	).SetAttributes("TITLE=Content")
}
