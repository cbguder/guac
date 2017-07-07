package commands

import "fmt"

type GroupsCommand struct {
}

func (c *GroupsCommand) Execute(args []string) error {
	u, err := buildUaa()
	if err != nil {
		return err
	}

	groups, err := u.GetGroups()
	if err != nil {
		return err
	}

	for _, group := range groups {
		fmt.Println(group.DisplayName)
	}

	return nil
}
