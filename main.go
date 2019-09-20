package main

import (
	"fmt"
	// "github.com/yautah/goroyale/client"
	// "github.com/yautah/goroyale/helper"
	// "net/url"
	"encoding/json"
	"github.com/gocolly/colly"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	Name          string                   `json:"name"`
	Rarity        string                   `json:"rarity"`
	Arena         string                   `json:"arena"`
	CardType      string                   `json:"type"`
	Description   string                   `json:"description"`
	CardStats     []map[string]interface{} `json:"cardStats"`
	CardUnits     []string                 `json:"cardUnits"`
	CardUnitsData []map[string]interface{} `json:"cardUnitsData"`
}

func main() {
	// s := helper.GetArenaById(54000014)
	// fmt.Println(s)

	// fmt.Println(helper.TOTAL_CARD_GOLD)
	// fmt.Println(len(helper.CardsIdMap))

	cards := make([]Card, 0)

	c := colly.NewCollector(
		// MaxDepth is 1, so only the links on the scraped page
		// is visited, and no further links are followed
		colly.MaxDepth(2),
		colly.URLFilters(
			regexp.MustCompile("https://statsroyale\\.com/zh/cards/"),
			regexp.MustCompile("https://statsroyale\\.com/zh/card/.+"),
		),
		colly.CacheDir("./cache"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		e.Request.Visit(link)
	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL.String())
	})

	// c.OnHTML(`div[class="card__headerContent"]`, func(e *colly.HTMLElement) {
	c.OnHTML(`body`, func(e *colly.HTMLElement) {
		card := Card{}
		card.Name = e.ChildText(".card__cardName")
		card.CardType = e.ChildText(".card__rarity")
		card.Rarity = e.ChildText(".card__rarityCaption .card__count")
		card.Description = e.ChildText(".card__description")

		e.ForEach(".card__metrics", func(index int, el *colly.HTMLElement) {
			x := make(map[string]interface{})
			x["name"] = el.ChildText(".card__metricsHeader")
			x["stats"] = make([]map[string]string, 0)
			el.ForEach(".card__metric", func(_ int, ell *colly.HTMLElement) {
				x["stats"] = append((x["stats"]).([]map[string]string), map[string]string{
					"name":  ell.ChildText(".card__metricCaption"),
					"value": ell.ChildText(".card__count"),
				})
			})

			card.CardStats = append(card.CardStats, x)
		})

		e.ForEach("div.statistics__tabs a", func(_ int, el *colly.HTMLElement) {
			card.CardUnits = append(card.CardUnits, el.Text)
		})

		e.ForEach("table.card__desktopTable tbody tr", func(index int, el *colly.HTMLElement) {

			data := make(map[string]interface{})
			data["data"] = make([]int, 0)
			el.ForEach("td", func(index int, ell *colly.HTMLElement) {
				text := strings.TrimSpace(strings.ReplaceAll(ell.Text, "\r\n", ""))
				if index == 0 {
					data["name"] = text
				} else {
					s, _ := strconv.Atoi(text)
					data["data"] = append(data["data"].([]int), s)
				}
			})

			card.CardUnitsData = append(card.CardUnitsData, data)
			// case "Language":
			// course.Language = el.ChildText("td:nth-child(2)")
			// case "Level":
			// course.Level = el.ChildText("td:nth-child(2)")
			// case "Commitment":
			// course.Commitment = el.ChildText("td:nth-child(2)")
			// case "How To Pass":
			// course.HowToPass = el.ChildText("td:nth-child(2)")
			// case "User Ratings":
			// course.Rating = el.ChildText("td:nth-child(2) div:nth-of-type(2)")
			// }
		})

		cards = append(cards, card)

	})

	// Start scraping on https://en.wikipedia.org

	c.Visit("https://statsroyale.com/zh/cards/")

	// 创建文件
	filePtr, err := os.Create("person_info.json")
	if err != nil {
		fmt.Println("Create file failed", err.Error())
		return
	}
	defer filePtr.Close()

	// 带JSON缩进格式写文件
	fmt.Println(len(cards))
	data, err := json.MarshalIndent(&cards, "", "  ")
	if err != nil {
		fmt.Println("Encoder failed", err.Error())

	} else {
		fmt.Println("Encoder success")
	}

	filePtr.Write(data)
}
