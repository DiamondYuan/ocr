package utils

import (
	"github.com/deckarep/gosx-notifier"
	"net/url"
)

func SendNotify(text string) {
	encodeString := url.QueryEscape(text)
	url := "https://diamondyuan.github.io/ocr/?ocr_result=" + encodeString
	note := gosxnotifier.NewNotification("The result has been copied to the clipboard . " +
		"\nClick to open in the browser")
	note.Title = "OCR"
	note.Sender = "com.apple.Safari"
	note.Link = url
	note.Sound = gosxnotifier.Default
	note.Push()
}

func SendErrorNotify(text string) {
	note := gosxnotifier.NewNotification(text)
	note.Title = "OCR"
	note.Sender = "com.apple.Safari"
	note.Sound = gosxnotifier.Default
	note.Push()
}
