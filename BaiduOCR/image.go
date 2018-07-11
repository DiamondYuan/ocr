package BaiduOCR

import (
	"os"
	"image"
	"strings"
	"fmt"
	"io/ioutil"
)

var (
	IMAGE_FORMATS = []string{"JPEG", "BMP", "PNG"}
)

func (this *BaiduOcr) LoadImageToByte(path string) ([]byte, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err

	}

	img, imgType, imgerr := image.DecodeConfig(file)
	if imgerr != nil {
		return nil, imgerr
	}
	imgType = strings.ToUpper(imgType)

	//1、检测文件类型
	rightFormat := false
	for _, v := range IMAGE_FORMATS {
		if imgType == v {
			rightFormat = true
			break
		}
	}

	if !rightFormat {
		return nil, fmt.Errorf("%s%s", "图像格式错误！只支持", IMAGE_FORMATS)
	}

	//2、检测图像大小
	if img.Width < 15 || img.Width > 4096 || img.Height < 15 || img.Height > 4096 {
		return nil, fmt.Errorf("%s", "图像大小不合适！最短边至少15px，最长边最大4096px")
	}
	file.Close()

	file, err = os.Open(path)
	//3、检测转码后的大小
	content, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		return nil, err2
	}

	text := bytesTOBaiduBase64(content)
	if len(text) >= MAX_SIZE {
		return nil, fmt.Errorf("%s", "图像文件编码后过大超过4M了")

	}

	defer file.Close()

	return content, nil
}
