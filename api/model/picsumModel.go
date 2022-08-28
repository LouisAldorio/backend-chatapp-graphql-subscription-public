package model

type PicsumResponse struct {
	ID          string `json:"id"`
	Author      string `json:"author"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Url         string `json:"url"`
	DownloadUrl string `json:"download_url"`
}
