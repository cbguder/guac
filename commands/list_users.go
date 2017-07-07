package commands

import "fmt"

type ListUsersCommand struct {
}

func (c *ListUsersCommand) Execute(args []string) error {
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
