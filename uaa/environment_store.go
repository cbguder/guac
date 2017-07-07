package uaa

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type EnvironmentStore interface {
	AliasEnvironment(urlOrAlias, alias string) error
	FindEnvironment(urlOrAlias string) (Environment, error)
	AddContext(urlOrAlias string, context Context) error
}

var ErrEnvironmentNotFound = errors.New("Environment not found")

func NewFilesystemEnvironmentStore(path string) (EnvironmentStore, error) {
	store := &filesystemEnvironmentStore{
		path: path,
	}

	err := store.Read()
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	return store, nil
}

type filesystemEnvironmentStore struct {
	path         string
	environments Environments
}

func (s *filesystemEnvironmentStore) AliasEnvironment(urlOrAlias, alias string) error {
	env, found := s.environments.FindEnvironment(urlOrAlias)
	if found {
		env.Aliases = append(env.Aliases, alias)
		s.environments.UpdateEnvironment(env)
	} else {
		env := Environment{
			Url:     urlOrAlias,
			Aliases: []string{alias},
		}

		s.environments.AddEnvironment(env)
	}

	return s.Write()
}

func (s *filesystemEnvironmentStore) FindEnvironment(urlOrAlias string) (Environment, error) {
	env, found := s.environments.FindEnvironment(urlOrAlias)
	if found == false {
		return env, ErrEnvironmentNotFound
	}

	return env, nil
}

func (s *filesystemEnvironmentStore) AddContext(urlOrAlias string, context Context) error {
	env, found := s.environments.FindEnvironment(urlOrAlias)
	if !found {
		return ErrEnvironmentNotFound
	}

	env.Contexts = append(env.Contexts, context)
	s.environments.UpdateEnvironment(env)

	return s.Write()
}

func (s *filesystemEnvironmentStore) Write() error {
	data, err := yaml.Marshal(s.environments)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.path, data, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (s *filesystemEnvironmentStore) Read() error {
	data, err := ioutil.ReadFile(s.path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &s.environments)
	if err != nil {
		return err
	}

	return nil
}
