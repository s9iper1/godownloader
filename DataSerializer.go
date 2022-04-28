package main

type Data struct {
	Apk       string  `json:"apk"`
	Version   float32 `json:"version"`
	TimeStamp string  `json:"timestamp"`
}

func (d Data) getApkDownloadAbleUrl() string {
	return d.Apk
}
