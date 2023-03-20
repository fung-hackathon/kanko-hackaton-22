package command

import (
	"kanko-hackaton-22/app/data"
	"kanko-hackaton-22/app/infra"
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

func getQuickReply() QuickReply {
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

func New(infra *infra.Firestore) *Command {
	return &Command{
		infra: infra,
	}
}

func (c *Command) ReadCommand(text string, userid string) (string, interface{}, error) {

	switch text {
	case "クイズラリーを解く":
		user, err := c.infra.Get(userid)
		if err != nil || user == nil {
			newUser := make(map[string]interface{})
			newUser["userid"] = userid
			newUser["quizStatus"] = 1
			newUser["quizid"] = 0
			newUser["progress"] = make([]bool, len(data.SpotsData))
			c.infra.Set(userid, newUser)
		} else {
			c.infra.Update(userid, "quizStatus", 1)
		}

		quickReply := getQuickReply()

		return "クイズラリーですね！\n巡った観光地はどこですか？選んでください↓", quickReply, nil

	default:
		c.infra.Update(userid, "quizStatus", 0)
		return "すみません、コマンドが見つかりませんでした...\nもう一度入力してください。", nil, nil
	}
}

func (c *Command) ReadText(text string, userid string) (string, interface{}, error) {
	user, err := c.infra.Get(userid)
	if err != nil {
		return "", nil, err
	}

	quizStatus := user["quizStatus"].(int64)
	if quizStatus > 0 && text == "やめる" {
		c.infra.Update(userid, "quizStatus", 0)
		return "クイズを終了します。次回の挑戦お待ちしてます！\n左下の ≡ アイコンをタップすると、使える機能のリストが表示されます↙", nil, nil
	}
	switch quizStatus {
	case 1:
		quizid := -1
		for i, spot := range data.SpotsData {
			if spot.Name == text {
				quizid = i + 1
			}
		}
		if err != nil || quizid < 0 {
			quickReply := getQuickReply()
			return "すみません、入力が間違っているようです...\nもう一度入力してください。\n(やめるときは「やめる」を入力してください)", quickReply, nil
		}

		progress := user["progress"].([]interface{})
		if progress[quizid-1] == true {
			quickReply := getQuickReply()
			return "すでにクリア済みの観光地です👀\n他の観光地を選んでください↓\n(やめるときは「やめる」を入力してください)", quickReply, nil
		}

		c.infra.Update(userid, "quizStatus", 2)
		c.infra.Update(userid, "quizid", quizid)

		return data.SpotsData[quizid-1].Name + "ですね！\n楽しめましたでしょうか？✨\n\nでは、答えをどうぞ！\n(「ひらがな」でお願いします)", nil, nil
	case 2:
		quizid := user["quizid"].(int64)
		//if data.SpotsData[quizid].Quiz.A == text {
		if text == "正解" {
			c.infra.Update(userid, "quizStatus", 0)
			c.infra.Update(userid, "quizid", 0)
			progress := user["progress"].([]interface{})
			progress[quizid-1] = true
			c.infra.Update(userid, "progress", progress)
			return "正解です！\nおめでとうございます🎉\n「ギャラリー」に新しい観光地情報が追加されます！\n左下の ≡ アイコンをタップして確認してみてください！↙", nil, nil
		}

		return "すみません、違うようです...\n別の答えを教えてください👀", nil, nil
	}
	return "すみません、最初に何をしてほしいか教えてください...\n左下の ≡ アイコンをタップすると、使える機能のリストが表示されます↙", nil, nil
}
