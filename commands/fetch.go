package commands

import (
	"ash/text-game/terminal"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func fetch(command Command) {
	response, err := http.Get(command.Args[0])

	if err != nil {
		fmt.Printf("fetch: failed to send request: %s", err.Error())
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

	if isJSON {
		var jsonBuffer bytes.Buffer
		err := json.Indent(&jsonBuffer, []byte(string(body)), "\r", "  ")
		if err != nil {
			fmt.Printf("fetch: failed to format JSON: %s", err.Error())
			return
		}

		fmt.Printf("%s\r\n", jsonBuffer.String())
		return
	}

	fmt.Print(string(body))
}
