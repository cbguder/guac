package commands

import "github.com/cbguder/guac/uaa"

type GetOwnerTokenCommand struct {
	ClientId     string `short:"c" long:"client" required:"true" value-name:"CLIENT"`
	ClientSecret string `short:"s" long:"secret" required:"true" value-name:"SECRET"`
	Username     string `short:"u" long:"username" required:"true" value-name:"USERNAME"`
	Password     string `short:"p" long:"password" required:"true" value-name:"PASSWORD"`
}

func (c *GetOwnerTokenCommand) Execute(args []string) error {
	envStore, err := buildEnvironmentStore()
	if err != nil {
		return err
	}

	u, err := buildUaaWithEnvStore(envStore)
	if err != nil {
		return err
	}

	token, err := u.GetOwnerToken(c.ClientId, c.ClientSecret, c.Username, c.Password)
	if err != nil {
		return err
	}

	context := uaa.Context{
		Client:   c.ClientId,
		Username: c.Username,
		Token:    token,
	}

	err = envStore.AddContext(Opts.Target, context)
	if err != nil {
		return err
	}

	return nil
}
