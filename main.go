package main

import (
	"fmt"
	"strings"

	"github.com/rudi9719/loggy"
	"samhofi.us/x/keybase"
)

var (
	k = keybase.NewKeybase()
	// Please ensure you change logOpts to be appropriate.
	logOpts = loggy.LogOpts{
		OutFile:   "kbot.log",
		KBTeam:    "nightmarehaus.bots",
		KBChann:   "general",
		ProgName:  "kbot",
		UseStdout: true,
		Level:     5,
	}
	log          = loggy.NewLogger(logOpts)
	commands     = make(map[string]Command)
	baseCommands = make([]string, 0)
)

func main() {
	log.LogInfo("Bot starting")
	if !k.LoggedIn {
		log.LogPanic("Bot is not logged into Keybase.")
	}
	k.Run(func(api keybase.ChatAPI) {
		handleMessage(api)
	})

}

func handleMessage(api keybase.ChatAPI) {
	if api.ErrorListen != nil {
		log.LogError(fmt.Sprintf("Error handling message, ```%+v```", api.ErrorListen))
	}
	if api.Msg.Content.Type == "text" {
		input := strings.Split(api.Msg.Content.Text.Body, " ")
		if len(input) < 2 {
			return
		}
		if input[0] != fmt.Sprintf("@%s", k.Username) {
			return
		}
		if c, ok := commands[input[1]]; ok {
			c.Exec(api)
		} else {
			log.LogWarn(fmt.Sprintf("Unknown command %s", input[1]))
		}
	}
}

// Replies to a message
func ReplyToMessage(api keybase.ChatAPI, msg string) {
	channel := keybase.Channel{
		Name:        api.Msg.Channel.Name,
		TopicName:   api.Msg.Channel.TopicName,
		MembersType: api.Msg.Channel.MembersType,
	}
	chat := k.NewChat(channel)
	_, err := chat.Reply(api.Msg.ID, msg)
	if err != nil {
		log.LogError(fmt.Sprintf("Error in ReplyToMessage ```%+v```"))
		log.LogDebug(fmt.Sprintf("Replying in @%s#%s to %s: %s",
			channel.Name,
			channel.TopicName,
			api.Msg.Sender.Username,
			api.Msg.Content.Text.Body,
		))
	}
}

// Respond to a message without replying
func RespondToMessage(api keybase.ChatAPI, msg string) {
	channel := keybase.Channel{
		Name:        api.Msg.Channel.Name,
		TopicName:   api.Msg.Channel.TopicName,
		MembersType: api.Msg.Channel.MembersType,
	}
	chat := k.NewChat(channel)
	_, err := chat.Send(msg)
	if err != nil {
		log.LogError(fmt.Sprintf("Error in ReplyToMessage ```%+v```"))
		log.LogDebug(fmt.Sprintf("Replying in @%s#%s to %s: %s",
			channel.Name,
			channel.TopicName,
			api.Msg.Sender.Username,
			api.Msg.Content.Text.Body,
		))
	}
}

// RegisterCommand registers a command to be used within the bot
func RegisterCommand(c Command) {
	var notAdded string
	for i, cmd := range c.Cmd {
		if _, ok := commands[cmd]; !ok {
			if i == 0 {
				baseCommands = append(baseCommands, cmd)
			}
			commands[cmd] = c
			continue
		}
		notAdded = fmt.Sprintf("%s, %s", notAdded, cmd)
	}
	if notAdded != "" {
		log.LogError(fmt.Sprintf("The following commands were not added because they already exist: %s", notAdded))

	}
}
