package commands

import (
	"fmt"
)

type ClientsCommand struct {
}

func (c *ClientsCommand) Execute(args []string) error {
	u, err := buildUaa()
	if err != nil {
		return err
	}

	clients, err := u.GetClients()
	if err != nil {
		return err
	}

	for _, client := range clients {
		fmt.Println(client.ClientId)
	}

	return nil
}
