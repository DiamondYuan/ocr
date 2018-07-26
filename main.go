package main

import (
	_ "image/png"
	"os"
	"github.com/DiamondYuan/ocr/ocrlib"
	"os/exec"
	"github.com/DiamondYuan/ocr/utils"
	"io/ioutil"
)

func main() {
	cmd := exec.Command("screencapture", "-i", "/tmp/ocr_snapshot.png")
	err := cmd.Run()
	if err != nil {
		os.Exit(0)
	}
	var ocr ocrlib.OCR
	ocr = new(ocrlib.Baidu)
	result, ocrError := ocr.GetResult("/tmp/ocr_snapshot.png")
	if ocrError != nil {
		handleError(err)
	}
	writeToClipBoard(result)
	utils.SendNotify(result)
	ioutil.WriteFile("/tmp/ocr_result", []byte(result), 0644)
}

func handleError(err error) {
	utils.SendErrorNotify(err.Error())
	os.Exit(0)
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
