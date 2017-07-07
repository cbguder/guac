package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type CurlCommand struct {
	Method string   `short:"X" long:"request" value-name:"METHOD" default:"GET"`
	Header []string `short:"H" long:"header" value-name:"HEADER"`
	Data   []string `short:"d" long:"data" value-name:"DATA"`

	Args struct {
		Path string `positional-arg-name:"PATH" required:"true"`
	} `positional-args:"true"`
}

func (c *CurlCommand) Execute(args []string) error {
	u, err := buildUaa()
	if err != nil {
		return err
	}

	header := http.Header{}

	for _, rawHeader := range c.Header {
		parts := strings.Split(rawHeader, ": ")
		header.Add(parts[0], parts[1])
	}

	data := strings.Join(c.Data, "&")

	res, err := u.Curl(c.Method, c.Args.Path, header, data)
	if err != nil {
		return err
	}

	if Opts.Verbose {
		fmt.Fprintf(os.Stderr, "%s %s\n", res.Proto, res.Status)

		for key, values := range res.Header {
			for _, value := range values {
				fmt.Fprintf(os.Stderr, "%s: %s\n", key, value)
			}
		}

		fmt.Fprintln(os.Stderr)
	}

	contentType := res.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		buf := bytes.NewBuffer(nil)

		err = json.Indent(buf, res.Body, "", "  ")
		if err != nil {
			return err
		}

		os.Stdout.Write(buf.Bytes())
	} else {
		_, err = os.Stdout.Write(res.Body)
	}

	fmt.Println()

	return err
}
