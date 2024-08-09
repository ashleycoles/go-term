package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/terminal"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func fetch(command Command, activeDirectory *filesystem.Directory) {
	if command.ArgsCount() < 1 {
		fmt.Print("\r\nfetch: missing url\r\n")
		return
	}

	var response *http.Response
	var requestErr error

	if command.HasValueFlag("method") {
		method, err := command.getFlagValue("method")
		if err != nil {
			fmt.Printf("\r\nfetch: %s\r\n", err.Error())
			return
		}

		switch method {
		case "post":
			typeValue, err := command.getFlagValue("type")
			if err != nil {
				fmt.Print("\r\nfetch: must provide --type to set request content-type\r\n")
				return
			}

			body, err := command.getFlagValue("body")

			if err != nil {
				fmt.Print("\r\nfetch: must provide --body to set request body\r\n")
				return
			}
			response, requestErr = http.Post(command.Args[0], typeValue, bytes.NewBufferString(body))
		default:
			fmt.Printf("\r\nfetch: unsupported method: %s\r\n", method)
			return
		}

	} else {
		response, requestErr = http.Get(command.Args[0])
	}

	if requestErr != nil {
		fmt.Printf("fetch: failed to send request: %s", requestErr.Error())
		return
	}

	if response == nil {
		fmt.Print("fetch: failed to send request")
		return
	}

	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	isJSON := strings.HasPrefix(contentType, "application/json")

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("fetch: failed to read response body: %s", err.Error())
		return
	}
	terminal.NewLine()

	output := string(body)

	if isJSON {
		var jsonBuffer bytes.Buffer
		err := json.Indent(&jsonBuffer, []byte(string(body)), "\r", "  ")
		if err != nil {
			fmt.Printf("fetch: failed to format JSON: %s", err.Error())
			return
		}

		output = jsonBuffer.String()
	}

	if command.HasValueFlag("dest") {
		dest, err := command.getFlagValue("dest")
		if err != nil {
			fmt.Printf("fetch: %s\r\n", err.Error())
			return
		}

		if !filesystem.IsValidFilename(dest) {
			fmt.Printf("fetch: %s is not a valid filename\r\n", dest)
			return
		}

		name, extension, _ := filesystem.GetFilenameParts(dest)

		if fileErr := (*activeDirectory).AddFile(name, extension, output); fileErr != nil {
			fmt.Printf("fetch: failed create file: %s: %s", dest, fileErr.Error())
			return
		}
	} else {
		fmt.Printf("%s\r\n", output)
	}
}
