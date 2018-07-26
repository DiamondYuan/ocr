package ocrlib


type Parameter map[string]string

type BaiduOcr struct {
	AppKey      string `json:"app_key"`
	AppSecret   string `json:"app_secret"`
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

type Baidu struct {
}

type BaiduSdk struct {
	AppId     string
	AppKey    string
	AppSecret string
	Token     *BaiduAccessToken
}

type BaiduAccessToken struct {
	BaiduErrorMessage
	Expires       int64  `json:"expires_in"` //过期时间，单位秒。一般为1个月)
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	SessionSecret string `json:"session_secret"`
	AccessToken   string `json:"access_token"` //要获取的Access Token；
	RefreshToken  string `json:"refresh_token"`
}

type BaiduErrorMessage struct {
	Code string `json:"error"`
	Msg  string `json:"error_description"`
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

type RecognizeItem struct {
	Location Image_location `json:"location"`
	Text string `json:"words"`
}


type Image_location struct {
	Left   int64 `json:"left"`
	Top    int64 `json:"top"`
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
}
