package commands

type AliasEnvCommand struct {
	Args struct {
		Alias string `positional-arg-name:"ALIAS" required:"true"`
	} `positional-args:"true"`
}

func (c *AliasEnvCommand) Execute(args []string) error {
	envStore, err := buildEnvironmentStore()
	if err != nil {
		return err
	}

	err = envStore.AliasEnvironment(Opts.Target, c.Args.Alias)
	if err != nil {
		return err
	}

	return nil
}
