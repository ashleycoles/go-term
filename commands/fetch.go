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
	terminal.NewLine()

	if command.ArgsCount() < 1 {
		fmt.Print("fetch: missing url")
		terminal.NewLine()
		return
	}

	var response *http.Response
	var requestErr error

	if command.HasValueFlag("method") {
		method, err := command.getFlagValue("method")
		if err != nil {
			fmt.Printf("fetch: %s", err.Error())
			terminal.NewLine()
			return
		}

		switch method {
		case "post":
			typeValue, err := command.getFlagValue("type")
			if err != nil {
				fmt.Print("fetch: must provide --type to set request content-type")
				terminal.NewLine()
				return
			}

			body, err := command.getFlagValue("body")

			if err != nil {
				fmt.Print("fetch: must provide --body to set request body")
				terminal.NewLine()
				return
			}
			response, requestErr = http.Post(command.Args[0], typeValue, bytes.NewBufferString(body))
		default:
			fmt.Printf("fetch: unsupported method: %s", method)
			terminal.NewLine()
			return
		}

	} else {
		response, requestErr = http.Get(command.Args[0])
	}

	if requestErr != nil {
		fmt.Printf("fetch: failed to send request: %s", requestErr.Error())
		terminal.NewLine()
		return
	}

	if response == nil {
		fmt.Print("fetch: failed to send request")
		terminal.NewLine()
		return
	}

	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	isJSON := strings.HasPrefix(contentType, "application/json")

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("fetch: failed to read response body: %s", err.Error())
		terminal.NewLine()
		return
	}

	output := string(body)

	if isJSON {
		var jsonBuffer bytes.Buffer
		err := json.Indent(&jsonBuffer, []byte(string(body)), "\r", "  ")
		if err != nil {
			fmt.Printf("fetch: failed to format JSON: %s", err.Error())
			terminal.NewLine()
			return
		}

		output = jsonBuffer.String()
	}

	if command.HasValueFlag("dest") {
		dest, err := command.getFlagValue("dest")
		if err != nil {
			fmt.Printf("fetch: %s", err.Error())
			terminal.NewLine()
			return
		}

		if !filesystem.IsValidFilename(dest) {
			fmt.Printf("fetch: %s is not a valid filename", dest)
			terminal.NewLine()
			return
		}

		name, extension, _ := filesystem.GetFilenameParts(dest)

		if fileErr := (*activeDirectory).AddFile(name, extension, output); fileErr != nil {
			fmt.Printf("fetch: failed create file: %s: %s", dest, fileErr.Error())
			terminal.NewLine()
			return
		}
	} else {
		fmt.Print(output)
		terminal.NewLine()
	}
}
