package main

import (
	"samhofi.us/x/keybase"
)

func init() {
	command := Command{
		Cmd:         []string{"hello", "hi"},
		Description: "Greet a user",
		Help:        "Example hello reply",
		Exec:        cmdHello,
	}

	RegisterCommand(command)

}
func cmdHello(api keybase.ChatAPI) {

	ReplyToMessage(api, "Howdy!")
}
