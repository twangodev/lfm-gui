package main

import "github.com/gen2brain/iup-go/iup"

var exit = iup.ActionFunc(func(ih iup.Ihandle) int { return iup.CLOSE })
