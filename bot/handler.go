package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// messageHandler watches for messages sent on the discord channel by other users and interacts with them, either by sending new messages or by performing actions.
func (b Bot) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == b.BotId {
		b.log.WarningLog("message sent by bot -> ignoring")
		return
	}

	if b.isAction(m.Content) {
		ba := BotActions{
			bot: b,
			log: b.log,
			s:   s,
			m:   m,
		}
		ba.executeActions()
		return
	}

	if b.config.BotListening {
		bm := BotMessages{
			bot: b,
			log: b.log,
			s:   s,
			m:   m,
		}

		err := bm.messageSelector()
		if err != nil {
			b.log.ErrorLog(fmt.Sprintf("could not sent message: %s", err.Error()))
		}
	}
}

// isAction returns a boolean value validating if this is an action or not.
// The message is considered an action if it has a "/" on the beginning.
func (b Bot) isAction(message string) bool {
	return strings.HasPrefix(message, "/")
}
