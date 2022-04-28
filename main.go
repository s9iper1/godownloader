package main

import (
	"fmt"
	"github.com/imroc/req/v3"
	"io"
	"net/http"
	"os"
)

func main() {
	data := Data{}
	client := req.C().EnableForceHTTP1()
	resp, err := client.R(). // Use R() to create a request
					SetResult(&data).
					Get("https://secure.chatapp.one/api/apk/")
	if err == nil {
		fmt.Println(data.Apk)
		downloadFile("app_debug.apk", data.Apk)
		fmt.Println(resp)
	} else {
		fmt.Println(err)
	}
}

func downloadFile(file string, url string) (err error) {
	// Create the file
	out, err := os.Create(file)
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
