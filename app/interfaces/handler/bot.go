package handler

import (
	"kanko-hackaton-22/app/command"
	"kanko-hackaton-22/app/message"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BotHandler struct {
	command *command.Command
}

func NewBotHandler(command *command.Command) *BotHandler {
	return &BotHandler{command: command}
}

func (h *BotHandler) Bot(c echo.Context) error {
	var req message.Request
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if len(req.Events) == 0 {
		return c.String(http.StatusOK, "pong")
	}

	text := req.Events[0].Message.Text

	var commandReply string
	var err error
	if len(text) >= 2 && text[:2] == "> " {
		commandReply, err = h.command.ReadCommand(text[2:])
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		commandReply, err = h.command.ReadText(text)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	reply := message.ReplyMessage{
		ReplyToken: req.Events[0].ReplyToken,
		Messages:   []message.Message{},
	}

	if commandReply != "" {
		reply.Messages = []message.Message{
			{Type: "text", Text: commandReply},
		}
	}

	if err := message.SendMessage(reply); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &reply)
}
