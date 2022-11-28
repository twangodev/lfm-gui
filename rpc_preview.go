package main

import (
	"bytes"
	"fmt"
	httpClient "github.com/bozd4g/go-http-client"
	"github.com/gen2brain/iup-go/iup"
	"github.com/hugolgst/rich-go/client"
	"github.com/nfnt/resize"
	log "github.com/sirupsen/logrus"
	"image"
	"image/jpeg"
	"image/png"
)

const previewBoldFont = "Helvetica, Bold 10"
const previewImageSize uint = 80

func resizeImg(img image.Image, size uint) image.Image {
	return resize.Resize(size, size, img, resize.Lanczos3)
}

func getRawHttpImage(url string) []byte {
	logContext := log.WithField("url", url)
	imageHttpClient := httpClient.New(url)
	request, err := imageHttpClient.Get("")
	logContext.Trace("Creating HTTP request")
	if err != nil {
		log.WithField("error", err).Error("Failed to create HTTP request")
		return nil
	}
	logContext.Trace("Sending HTTP request")
	response, err := imageHttpClient.Do(request)
	if err != nil {
		log.WithField("error", err).Error("Failed to execute HTTP request")
		return nil
	}

	logContext.Trace("Reading HTTP response")
	responseStruct := response.Get()
	return responseStruct.Body
}

func getHttpJpeg(url string) image.Image {
	img, err := jpeg.Decode(bytes.NewReader(getRawHttpImage(url)))
	if err != nil {
		generateLogContext("getHttpJpeg").WithField("error", err).Error("Failed to decode image")
		return nil
	}

	return resizeImg(img, previewImageSize)
}

func getDefaultImage() image.Image {
	data := getRawHttpImage("https://raw.githubusercontent.com/twangodev/lfm-gui/main/assets/lfm_logo.png")
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		generateLogContext("").WithField("error", err).Error("Failed to decode image")
		return nil
	}
	return img
}

func getPreviewImage(url string) image.Image {
	if url == "" {
		return getDefaultImage()
	}
	img := getHttpJpeg(url)
	if img == nil {
		return getDefaultImage()
	}
	return img
}

func updateRpcPreviewImage(url string) {
	img := getPreviewImage(url)

	iup.ImageFromImage(img).SetHandle("previewImageSource")
	go iup.GetHandle("rpcPreviewImage").SetAttribute("IMAGE", "previewImageSource")
}

func updateRpcPreview(activity client.Activity) {
	logContext := log.WithFields(log.Fields{
		"rowOneText": activity.Details,
		"rowTwoText": activity.State,
	})
	logContext.Trace("Updating RPC preview")
	go iup.GetHandle("rpcPreviewRowOne").SetAttribute("TITLE", activity.Details)
	go iup.GetHandle("rpcPreviewRowTwo").SetAttribute("TITLE", activity.State)

	go updateRpcPreviewImage(activity.LargeImage)
}

func updateRpcPreviewElapsed(minute int, second int) {
	elapsedString := fmt.Sprintf("%d:%02d elapsed", minute, second)
	go iup.GetHandle("rpcPreviewTimeElapsed").SetAttribute("TITLE", elapsedString)
}

func rpcPreviewFrame() iup.Ihandle {

	go iup.ImageFromImage(getDefaultImage()).SetHandle("previewImageSource")

	return iup.Hbox(
		iup.Label("").SetAttribute("IMAGE", "previewImageSource").SetHandle("rpcPreviewImage"),
		iup.Vbox(
			iup.Label("Playing last.fm").SetAttribute("FONT", previewBoldFont),
			iup.Label(config.rows.rowOneText).SetAttribute("EXPAND", "HORIZONTAL").SetHandle("rpcPreviewRowOne"),
			iup.Label(config.rows.rowTwoText).SetAttribute("EXPAND", "HORIZONTAL").SetHandle("rpcPreviewRowTwo"),
			iup.Label("0:00 elapsed").SetHandle("rpcPreviewTimeElapsed"),
		),
	)

}
