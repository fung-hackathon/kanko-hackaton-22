package command

import "kanko-hackaton-22/app/infra"

type Command struct {
	infra *infra.Firestore
}

func New(infra *infra.Firestore) *Command {
	return &Command{
		infra: infra,
	}
}

func (c *Command) ReadCommand(text string) (string, error) {
	return "これはコマンドです", nil
}

func (c *Command) ReadText(text string) (string, error) {
	return "これはテキストです", nil
}
