package main

import "samhofi.us/x/keybase"

type Command struct {
	Cmd         []string              // Any aliases that can trigger this command
	Description string                // The description of the command for a help command
	Help        string                // The preview help text
	Exec        func(keybase.ChatAPI) // A function that takes the command, and any arguments to the command
}
