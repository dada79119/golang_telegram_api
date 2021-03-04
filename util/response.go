package util

// Response Success
type RS struct {
	Message  string `json:"message"`
	Status   bool `json:"status"`
	Response  map[string]interface{} `json:"response"`
}

// Response Error
// field ---> error ----> return到首頁
// field ---> alarm ----> 傳送message到首頁即可

type RE struct {
	Field    string `json:"field"`
	Message  string `json:"message"`
}
