package models

const (
	Common = "common"
	Cover  = "cover"
)

type News struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Tag   string `json:"tag"`
}
