package command

import (
	"kanko-hackaton-22/app/config"
	"kanko-hackaton-22/app/data"
	"kanko-hackaton-22/app/infra"
	"kanko-hackaton-22/app/package/strDiff"
	"math/rand"
)

type Command struct {
	infra *infra.Firestore
}

type ItemAction struct {
	Type  string `json:"type"`
	Label string `json:"label,omitempty"`
	Text  string `json:"text,omitempty"`
}

type Item struct {
	Type     string     `json:"type"`
	ImageURL string     `json:"imageUrl,omitempty"`
	Action   ItemAction `json:"action"`
}

type QuickReply struct {
	Items []Item `json:"items"`
}

type ImageMessage struct {
	Type               string `json:"type"`
	OriginalContentUrl string `json:"originalContentUrl"`
	PreviewImageUrl    string `json:"previewImageUrl"`
}

func getQuickReplyForSpot() QuickReply {
	quickReply := QuickReply{
		Items: []Item{},
	}
	for _, spot := range data.SpotsData {
		quickReply.Items = append(quickReply.Items, Item{
			Type:     "action",
			ImageURL: spot.Image,
			Action: ItemAction{
				Type:  "message",
				Label: spot.Name,
				Text:  spot.Name,
			},
		})
	}
	return quickReply
}

func getQuickReplyForAccessibility() QuickReply {
	quickReply := QuickReply{
		Items: []Item{},
	}
	for _, accessibility := range data.AccessibilityData {
		quickReply.Items = append(quickReply.Items, Item{
			Type:     "action",
			ImageURL: "https://firebasestorage.googleapis.com/v0/b/kanko-hackathon.appspot.com/o/icons%2Fwhite.png?alt=media&token=b0309d55-fd23-4e84-a17b-144e8cb16179",
			Action: ItemAction{
				Type:  "message",
				Label: accessibility,
				Text:  accessibility,
			},
		})
	}
	return quickReply
}

func New(infra *infra.Firestore) *Command {

	return &Command{
		infra: infra,
	}
}

func (c *Command) setNewUser(userid string) {
	newUser := make(map[string]interface{})
	newUser["userid"] = userid
	newUser["quizStatus"] = 0
	newUser["recommendStatus"] = 0
	newUser["quizid"] = 0
	newUser["progress"] = make([]bool, len(data.SpotsData))
	c.infra.Set(userid, newUser)
}

func (c *Command) ReadCommand(text string, userid string) (string, interface{}, error) {

	switch text {
	case "クイズラリーを解く":

		user, err := c.infra.Get(userid)
		if err != nil || user == nil {
			c.setNewUser(userid)
		}
		c.infra.Update(userid, "recommendStatus", 0)
		c.infra.Update(userid, "quizStatus", 1)

		quickReply := getQuickReplyForSpot()

		return "クイズラリーですね！\n巡った観光地はどこですか？選んでください↓", quickReply, nil

	case "おすすめを調べる":
		user, err := c.infra.Get(userid)
		if err != nil || user == nil {
			c.setNewUser(userid)
		}
		c.infra.Update(userid, "quizStatus", 0)
		c.infra.Update(userid, "recommendStatus", 1)

		quickReply := getQuickReplyForAccessibility()

		return "おすすめですね！任せてください💪\nクイズのある観光地を探します。\n\nどのくらいの時間で移動できますか？選んでください↓", quickReply, nil

	case "ギャラリーを開く":
		user, err := c.infra.Get(userid)
		if err != nil || user == nil {
			c.setNewUser(userid)
		}
		return config.HOST.File("gallery?userid=" + userid), nil, nil

	case "クイズラリーの対象スポットを見る":
		user, err := c.infra.Get(userid)
		if err != nil || user == nil {
			c.setNewUser(userid)
		}
		return config.HOST.File("spots"), nil, nil

	default:
		c.infra.Update(userid, "quizStatus", 0)
		c.infra.Update(userid, "recommendStatus", 0)
		return "すみません、コマンドが見つかりませんでした...\nもう一度入力してください。", nil, nil
	}
}

func (c *Command) ReadText(text string, userid string) (string, interface{}, error) {
	user, err := c.infra.Get(userid)
	if err != nil {
		return "", nil, err
	}

	recommendStatus := user["recommendStatus"].(int64)
	if recommendStatus > 0 && text == "やめる" {
		c.infra.Update(userid, "quizStatus", 0)
		return "おすすめ機能を終了します。自分で探ってみるのもいいですね！\n他になにかお手伝いできることがあれば、左下の ≡ アイコンをタップして選択してください↙", nil, nil
	}

	switch recommendStatus {
	case 1:
		accessibility := -1
		for i, a := range data.AccessibilityData {
			if a == text {
				accessibility = i + 1
			}
		}
		if err != nil || accessibility < 0 {
			quickReply := getQuickReplyForAccessibility()
			return "すみません、入力が間違っているようです...\nもう一度入力してください。\n(やめるときは「やめる」を入力してください)", quickReply, nil
		}

		spotDataRecommendable := make([]data.Spot, 0, len(data.SpotsData))
		progress := user["progress"].([]interface{})
		for i, spot := range data.SpotsData {
			if spot.Accessibility == accessibility && len(progress) > i && progress[i] == false {
				spotDataRecommendable = append(spotDataRecommendable, spot)
			}
		}
		if len(spotDataRecommendable) == 0 {
			quickReply := getQuickReplyForAccessibility()
			return "すみません、該当する観光地が見つかりませんでした...\nもう一度入力してください。\n(やめるときは「やめる」を入力してください)", quickReply, nil
		}

		c.infra.Update(userid, "recommendStatus", 0)
		spot := spotDataRecommendable[rand.Intn(len(spotDataRecommendable))]

		imageMessage := ImageMessage{
			Type:               "image",
			OriginalContentUrl: spot.Image,
			PreviewImageUrl:    spot.Image,
		}

		return "では、このような場所はどうでしょうか？\n\n「" + spot.Name + "」\n\nクイズラリーもありますよ！✨\n是非訪れてみてください👀\n\n他になにかお手伝いできることがあれば、左下の ≡ アイコンをタップして選択してください↙", imageMessage, nil

	}

	quizStatus := user["quizStatus"].(int64)
	if quizStatus > 0 && text == "やめる" {
		c.infra.Update(userid, "quizStatus", 0)
		return "クイズを終了します。次回の挑戦お待ちしてます！\n他になにかお手伝いできることがあれば、左下の ≡ アイコンをタップして選択してください↙", nil, nil
	}

	switch quizStatus {
	case 1:
		quizid := -1
		for i, spot := range data.SpotsData {
			if spot.Name == text {
				quizid = i
			}
		}

		if err != nil || quizid < 0 {
			quickReply := getQuickReplyForSpot()
			return "すみません、入力が間違っているようです...\nもう一度入力してください。\n(やめるときは「やめる」を入力してください)", quickReply, nil
		}

		progress := user["progress"].([]interface{})
		if progress[quizid] == true {
			quickReply := getQuickReplyForSpot()
			return "すでにクリア済みの観光地です👀\n他の観光地を選んでください↓\n(やめるときは「やめる」を入力してください)", quickReply, nil
		}

		c.infra.Update(userid, "quizStatus", 2)
		c.infra.Update(userid, "quizid", quizid)

		return data.SpotsData[quizid].Name + "ですね！\n楽しめましたでしょうか？✨\n\nでは、答えをどうぞ！\n(「ひらがな」でお願いします)", nil, nil

	case 2:
		quizid := user["quizid"].(int64)
		spot := data.SpotsData[quizid]
		answer := spot.Quiz.A
		if strDiff.EditDistance(answer, text) <= 1 {
			c.infra.Update(userid, "quizStatus", 0)
			c.infra.Update(userid, "quizid", 0)
			progress := user["progress"].([]interface{})
			progress[quizid] = true
			c.infra.Update(userid, "progress", progress)

			imageMessage := ImageMessage{
				Type:               "image",
				OriginalContentUrl: spot.Image,
				PreviewImageUrl:    spot.Image,
			}

			return "正解です！\nおめでとうございます🎉\n\n「" + spot.Name + "」\n" + spot.Quiz.Comment + "\n\n「ギャラリー」に新しい観光地情報が追加されました！\n左下の ≡ アイコンをタップして確認してみてください！↙", imageMessage, nil
		}

		return "すみません、違うようです...\n別の答えを入力してください👀\n(やめるときは「やめる」を入力してください)", nil, nil
	}
	return "こんにちは！✨\n函館市内でのちょっとした観光をお手伝いします💪\n左下の ≡ アイコンをタップして機能を選択してください↙", nil, nil
}
