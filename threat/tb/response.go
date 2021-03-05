package threatbook

// Response type
type Response struct {
	ResponseCode int         `json:"response_code"`
	VerboseMsg   string      `json:"verbose_msg"`
	Data         interface{} `json:"data"`
}
