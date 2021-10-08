package tool

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func SendError(urlStr string, errMsg error) {
	if urlStr == "" {
		fmt.Println("urlStr is empty")
		return
	}

	var from string
	if len(os.Args) > 0 {
		from = os.Args[0]
	}
	resp, err := http.PostForm(urlStr,
		url.Values{"err": {"From: " + from + "\n" + errMsg.Error()}})

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
}
