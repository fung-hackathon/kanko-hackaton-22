package data

type Spot struct {
	Name     string
	Url      string
	OpenTime string
	Image    string
	Quiz     Quiz
}

type Quiz struct {
	Q string
	A string
}

var (
	SpotsData = []Spot{
		{
			Name:     "旧イギリス領事館（開港記念館）",
			Url:      "https://www.city.hakodate.hokkaido.jp/docs/2014032400354/",
			OpenTime: "4月~10月: 9:00~19:00, 11月~3月: 9:00~17:00",
			Image:    "https://www.city.hakodate.hokkaido.jp/docs/2014032400354/files/british_consulate.png",
		},
		{
			Name:     "文学館",
			Url:      "https://www.city.hakodate.hokkaido.jp/docs/2014042000056/",
			OpenTime: "4月~10月: 9:00~19:00, 11月~3月: 9:00~17:00",
			Image:    "https://www.city.hakodate.hokkaido.jp/docs/2014042000056/files/bungakukan.jpg",
		},
		{
			Name:     "函館市北方民族資料館",
			Url:      "https://www.city.hakodate.hokkaido.jp/docs/2015121000103/",
			OpenTime: "4月~10月: 9:00~19:00, 11月~3月: 9:00~17:00",
			Image:    "https://www.city.hakodate.hokkaido.jp/docs/2015121000103/files/hoppouminzokushiryoukan.jpg",
		},
		{
			Name:     "函館市青函連絡船記念館摩周丸",
			Url:      "https://www.city.hakodate.hokkaido.jp/docs/2014022000311/",
			OpenTime: "4月~10月: 9:00~18:00, 11月~3月: 9:00~17:00",
			Image:    "https://www.city.hakodate.hokkaido.jp/docs/2014022000311/files/masyu1.jpg",
		},
		{
			Name:     "箱館奉行所",
			Url:      "https://www.city.hakodate.hokkaido.jp/docs/2014021000039/",
			OpenTime: "4月~10月: 9:00~18:00, 11月~3月: 9:00~17:00",
			Image:    "https://www.city.hakodate.hokkaido.jp/docs/2014021000039/files/bugyosyo.jpg",
		},
		{
			Name:     "北洋資料館",
			Url:      "https://www.city.hakodate.hokkaido.jp/docs/2014010800367/",
			OpenTime: "4月~10月: 9:00~19:00, 11月~3月: 9:00~17:00",
			Image:    "https://www.city.hakodate.hokkaido.jp/docs/2014010800367/files/hokuyou3.jpg",
		},
		{
			Name:     "函館市縄文文化交流センター",
			Url:      "https://www.city.hakodate.hokkaido.jp/docs/2014021000060/",
			OpenTime: "4月~10月:9:00~17:00, 11月~3月:9:00~16:30",
			Image:    "https://www.city.hakodate.hokkaido.jp/docs/2014021000060/files/jomoncenter.jpg",
		},
		{
			Name:     "市立函館博物館",
			Url:      "https://www.city.hakodate.hokkaido.jp/docs/2015121000080/",
			OpenTime: "4月~10月:9:00~16:30, 11月~3月:9:00~16:00",
			Image:    "https://www.city.hakodate.hokkaido.jp/docs/2015121000080/files/shirituhakodatehakubutukan.jpg",
		},
	}
)
