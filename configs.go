package main

import "lfm-gui/concurrency"

type appConfiguration struct {
	title     string
	discordID string
}

type previewConfiguration struct {
	enabled                 bool
	albumConfig             albumConfiguration
	smallImageConfig        smallImageConfiguration
	albumDefaultPreviewPath string
}

type albumConfiguration struct {
	cover                 string
	hover                 bool
	albumDefaultHoverText string
}

type smallImageConfiguration struct {
	enabled              bool
	smallImageDefaultKey string
	smallImageHover      bool
	smallImageHoverText  string
	lovedEnabled         bool
}

type rowsConfiguration struct {
	rowOne      bool
	rowOneText  string
	rowTwo      bool
	rowTwoText  string
	timeElapsed bool
}

type buttonsConfiguration struct {
	profileButton     bool
	profileButtonText string
	songButton        bool
	songButtonText    string
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

type boolDependencyHandle struct {
	handle       string
	configOpcode concurrency.Opcode
	ref          *bool
}
