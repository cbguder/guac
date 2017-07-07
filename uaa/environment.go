package uaa

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Environment struct {
	Url      string    `yaml:"url"`
	Aliases  []string  `yaml:"aliases"`
	Contexts []Context `yaml:"contexts"`
}

func (e Environment) UnauthorizedRequest(method, path string, body io.Reader) (*http.Request, error) {
	absUrl := fmt.Sprintf("%s%s", e.Url, path)
	return http.NewRequest(method, absUrl, body)
}

func (e Environment) AuthorizedRequest(method, path string, body io.Reader) (*http.Request, error) {
	if len(e.Contexts) == 0 {
		return nil, errors.New("No contexts found for environment")
	}

	req, err := e.UnauthorizedRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	context := e.Contexts[len(e.Contexts)-1]

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", context.Token.AccessToken))

	return req, nil
}

type Environments []Environment

func (c Environments) FindEnvironment(urlOrAlias string) (Environment, bool) {
	for _, env := range c {
		if env.Url == urlOrAlias {
			return env, true
		}

		for _, alias := range env.Aliases {
			if alias == urlOrAlias {
				return env, true
			}
		}
	}

	return Environment{}, false
}

func (c *Environments) AddEnvironment(env Environment) {
	*c = append(*c, env)
}

func (c Environments) UpdateEnvironment(env Environment) {
	for i := range c {
		if c[i].Url == env.Url {
			c[i] = env
			return
		}
	}
}
