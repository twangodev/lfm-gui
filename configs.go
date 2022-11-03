package main

type appConfiguration struct {
	title     string
	version   string
	discordID string
}

type previewConfiguration struct {
	enabled                 bool
	albumConfig             albumConfiguration
	smallImageConfig        smallImageConfiguration
	albumDefaultPreviewPath string
}

type albumConfiguration struct {
	cover                    string
	albumDefaultHoverEnabled bool
	albumDefaultHoverText    string
}

type smallImageConfiguration struct {
	enabled                bool
	smallImageDefaultKey   string
	smallImageHoverEnabled bool
	smallImageHoverText    string
	lovedEnabled           bool
}

type rowsConfiguration struct {
	rowOneEnabled      bool
	rowOne             string
	rowTwoEnabled      bool
	rowTwo             string
	timeElapsedEnabled bool
}

type buttonsConfiguration struct {
	profileButtonEnabled bool
	profileButton        string
	songButtonEnabled    bool
	songButton           string
}

type configuration struct {
	app         appConfiguration
	username    string
	preview     previewConfiguration
	refreshTime int
	rows        rowsConfiguration
	buttons     buttonsConfiguration
	state       bool
}
