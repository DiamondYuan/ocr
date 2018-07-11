package BaiduOCR

type BaiduOcr struct {
	sdk *BaiduSdk
}


type ImageResult struct {
	baseResult
	Words []RecognizeItem `json:"words_result"`
	Count int64 `json:"words_result_num"`
}


type baseResult struct {
	Id int64 `json:"log_id"` //唯一的log id，用于问题定位
	BaiduErrorMessage
}
