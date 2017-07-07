package commands

import "github.com/cbguder/guac/uaa"

type GetClientTokenCommand struct {
	ClientId     string `short:"c" long:"client" required:"true" value-name:"CLIENT"`
	ClientSecret string `short:"s" long:"secret" value-name:"SECRET"`
}

func (c *GetClientTokenCommand) Execute(args []string) error {
	envStore, err := buildEnvironmentStore()
	if err != nil {
		return err
	}

	u, err := buildUaaWithEnvStore(envStore)
	if err != nil {
		return err
	}

	if c.ClientSecret == "" {
		c.ClientSecret, err = promptForSecret("Client Secret")
		if err != nil {
			return err
		}
	}

	token, err := u.GetClientToken(c.ClientId, c.ClientSecret)
	if err != nil {
		return err
	}

	context := uaa.Context{
		Client: c.ClientId,
		Token:  token,
	}

	err = envStore.AddContext(Opts.Target, context)
	if err != nil {
		return err
	}

	return nil
}
