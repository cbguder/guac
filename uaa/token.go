package uaa

type Token struct {
	AccessToken string   `yaml:"access_token"`
	TokenType   string   `yaml:"token_type"`
	ExpiresIn   int      `yaml:"-"`
	Scope       []string `yaml:"scope"`
	Jti         string   `yaml:"jti"`
}

type Context struct {
	Client   string `yaml:"client"`
	Username string `yaml:"username,omitempty"`
	Token    Token  `yaml:"token"`
}
