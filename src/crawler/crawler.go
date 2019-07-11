package crawler

import (
	"io/ioutil"
	"net/http"
)

func Fetch(url string) string {
	client := &http.Client{}
	//自定义请求头
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0)")
	res, err := client.Do(req)
	if err != nil {
		println("HTTP GET ERROR: ", err.Error())
		return ""
	}
	if res.StatusCode != 200 {
		println("HTTP STATUS ERROR: ", res.StatusCode)
		return ""
	}
	defer res.Body.Close()

	//读出body的内容
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		println("READ ERROR: ", err.Error())
		return ""
	}
	return string(body)
}
