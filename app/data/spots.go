package data

import (
	"kanko-hackaton-22/app/config"
)

type Spot struct {
	Name          string
	Url           string
	OpenTime      string
	Image         string
	Quiz          Quiz
	Accessibility int
}

type Quiz struct {
	Q       string
	A       string
	Comment string
}

var (
	SpotsData = []Spot{
		{
			Name:     "函館市地域交流まちづくりセンター",
			Url:      "https://www.city.hakodate.hokkaido.jp/docs/2014032400101/",
			OpenTime: "9:00～21:00",
			Image:    "https://www.city.hakodate.hokkaido.jp/docs/2014032400101/files/matisen.png",
			Quiz: Quiz{
				Q:       "この施設には、東北以北最古の手動〇〇がある",
				A:       "エレベータ",
				Comment: "1934年(昭和9年)の3月以降に設置されたと言われています。",
			},
		},
		{
			Name:     "旧イギリス領事館（開港記念館）",
			Url:      "https://www.city.hakodate.hokkaido.jp/docs/2014032400354/",
			OpenTime: "4月～10月: 9:00～19:00, 11月～3月: 9:00～17:00",
			Image:    "https://www.city.hakodate.hokkaido.jp/docs/2014032400354/files/british_consulate.png",
			Quiz: Quiz{
				Q: "問題問題問題",
				A: "かいとう",
			},
		},
		{
			Name:     "文学館",
			Url:      "https://www.city.hakodate.hokkaido.jp/docs/2014042000056/",
			OpenTime: "4月～10月: 9:00～19:00, 11月～3月: 9:00～17:00",
			Image:    "https://www.city.hakodate.hokkaido.jp/docs/2014042000056/files/bungakukan.jpg",
			Quiz: Quiz{
				Q: "問題問題問題",
				A: "かいとう",
			},
		},
		{
			Name:     "函館市北方民族資料館",
			Url:      "https://www.city.hakodate.hokkaido.jp/docs/2015121000103/",
			OpenTime: "4月～10月: 9:00～19:00, 11月～3月: 9:00～17:00",
			Image:    "https://www.city.hakodate.hokkaido.jp/docs/2015121000103/files/hoppouminzokushiryoukan.jpg",
			Quiz: Quiz{
				Q: "問題問題問題",
				A: "かいとう",
			},
		},
		{
			Name:     "函館市青函連絡船記念館摩周丸",
			Url:      "https://www.city.hakodate.hokkaido.jp/docs/2014022000311/",
			OpenTime: "4月～10月: 9:00～18:00, 11月～3月: 9:00～17:00",
			Image:    "https://www.city.hakodate.hokkaido.jp/docs/2014022000311/files/masyu1.jpg",
			Quiz: Quiz{
				Q: "問題問題問題",
				A: "かいとう",
			},
		},
		{
			Name:     "市立函館博物館",
			Url:      "https://www.city.hakodate.hokkaido.jp/docs/2015121000080/",
			OpenTime: "4月～10月:9:00～16:30, 11月～3月:9:00～16:00",
			Image:    "https://www.city.hakodate.hokkaido.jp/docs/2015121000080/files/shirituhakodatehakubutukan.jpg",
			Quiz: Quiz{
				Q: "問題問題問題",
				A: "かいとう",
			},
		},
		{
			Name:     "函館山",
			Url:      "https://www.city.hakodate.hokkaido.jp/docs/2015121000080/",
			OpenTime: "4月～10月:9:00～16:30, 11月～3月:9:00～16:00",
			Image:    config.HOST.File("static/img/mthakodate.jpg"),
			Quiz: Quiz{
				Q: "問題問題問題",
				A: "かいとう",
			},
		},
	}
)
