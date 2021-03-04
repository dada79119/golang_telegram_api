package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Request(jsonMap map[string]interface{}, url string) {
	reqJson, _ := json.Marshal(jsonMap)
	mapStr := string(reqJson)
	req, _ := http.NewRequest("POST", url, strings.NewReader(mapStr))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Key", "%my.pay!req!7ha6h%u-est%mem-ber!-sha.94!87%key-stay.pli.tus!")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
