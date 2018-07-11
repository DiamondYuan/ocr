package main

import (
	"log"
	"os/exec"
	ocr "github.com/DiamondYuan/ocr/BaiduOCR"
	"github.com/deckarep/gosx-notifier"
	"encoding/base64"
)

func main() {
	sdk := ocr.BaiduSdk{
		AppKey:    "",
		AppSecret: "",
	}
	baidu := ocr.BaiduOcr{
	}
	baidu.Init(&sdk)
	cmd := exec.Command("screencapture", "-i", "/tmp/ocr_snapshot.png")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
	result, _ := baidu.RecognizedImage("/tmp/ocr_snapshot.png")
	writeToClipBoard(result)
	encodeString := base64.StdEncoding.EncodeToString([]byte(result))
	url := "https://diamondyuan.github.io/ocr/?ocr_result=" + encodeString
	note := gosxnotifier.NewNotification("Ocr Result have copy to your clipboard.")
	note.Title = "OCR"
	note.Sender = "com.apple.Safari"
	note.Link = url
	note.Sound = gosxnotifier.Default
	note.Push()
}

func writeToClipBoard(text string) error {
	copyCmd := exec.Command("pbcopy")
	in, err := copyCmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := copyCmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(text)); err != nil {
		return err
	}
	if err := in.Close(); err != nil {
		return err
	}
	return copyCmd.Wait()
}
