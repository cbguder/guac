package commands

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os/user"
	"path/filepath"
	"syscall"
	"time"

	"github.com/cbguder/guac/uaa"
	"golang.org/x/crypto/ssh/terminal"
)

func buildEnvironmentStore() (uaa.EnvironmentStore, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	envStorePath := filepath.Join(usr.HomeDir, ".guac.yml")
	return uaa.NewFilesystemEnvironmentStore(envStorePath)
}

func buildHttpClient() *http.Client {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,

		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: Opts.Insecure,
		},
	}

	return &http.Client{
		Transport: transport,
		Timeout:   time.Minute,
	}
}

func buildUaa() (*uaa.UAA, error) {
	envStore, err := buildEnvironmentStore()
	if err != nil {
		return nil, err
	}

	return buildUaaWithEnvStore(envStore)
}

func buildUaaWithEnvStore(envStore uaa.EnvironmentStore) (*uaa.UAA, error) {
	env, err := envStore.FindEnvironment(Opts.Target)
	if err != nil {
		return nil, err
	}

	httpClient := buildHttpClient()

	return &uaa.UAA{
		Environment: env,
		Client:      httpClient,
	}, nil
}

func promptForSecret(prompt string) (string, error) {
	fmt.Printf("%s: ", prompt)
	bytes, err := terminal.ReadPassword(syscall.Stdin)
	if err != nil {
		return "", err
	}

	fmt.Println()

	return string(bytes), nil
}
