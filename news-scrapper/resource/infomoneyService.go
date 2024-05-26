package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocolly/colly"

	model "news-scrapper/models"
)

func Search(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Resgatando notícias em https://www.infomoney.com.br")
	c := colly.NewCollector()

	var todaysNews []model.News

	c.OnHTML(".home-widgets .cover-link", func(h *colly.HTMLElement) {
		n := model.News{
			Title: h.Attr("title"),
			Url:   h.Attr("href"),
			Tag:   model.Cover,
		}
		todaysNews = append(todaysNews, n)
	})

	c.OnHTML(".home-widgets .hl-title", func(h *colly.HTMLElement) {
		n := model.News{
			Title: h.ChildAttr("a", "title"),
			Url:   h.ChildAttr("a", "href"),
			Tag:   model.Common,
		}
		todaysNews = append(todaysNews, n)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Print("Erro: ")
		fmt.Println(err)
	})

	c.Visit("https://www.infomoney.com.br")

	json.NewEncoder(w).Encode(todaysNews)
}
