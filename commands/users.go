package commands

import "fmt"

type UsersCommand struct {
}

func (c *UsersCommand) Execute(args []string) error {
	u, err := buildUaa()
	if err != nil {
		return err
	}

	users, err := u.GetUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		fmt.Println(user.Username)
	}

	return nil
}
