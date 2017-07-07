package commands

type GlobalOpts struct {
	Target   string `short:"t" long:"target" required:"true" value-name:"TARGET"`
	Insecure bool   `short:"k" long:"insecure"`
	Verbose  bool   `short:"v" long:"verbose"`

	AliasEnv       AliasEnvCommand       `command:"alias-env"`
	GetClientToken GetClientTokenCommand `command:"get-client-token"`
	GetOwnerToken  GetOwnerTokenCommand  `command:"get-owner-token"`
	Users          UsersCommand          `command:"users"`
	Clients        ClientsCommand        `command:"clients"`
	Curl           CurlCommand           `command:"curl"`
}

var Opts GlobalOpts
