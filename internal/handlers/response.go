package handlers

import (
	"clout-bot/internal/clout"
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

func mentionsBot(s *dg.Session, m *dg.MessageCreate) bool {
	for _, m := range m.Mentions {
		if s.State.User.ID == m.ID {
			return true
		}
	}
	return false
}

func HandleResponse(s *dg.Session, m *dg.MessageCreate) {
	service := clout.GetService()
	if mentionsBot(s, m) {
		relationships := service.GetRelationships(m.Author)
		if len(relationships) <= 0 {
			s.ChannelMessageSend(m.ChannelID, "you have no clout...")
			return
		}
		users := make([]string, 0, len(relationships))
		for _, r := range relationships {
			users = append(users, fmt.Sprintf("<@%s> %d", r.User.ID, r.Count ))
		}
		response := fmt.Sprintf("these guys fw youu:\n- %s", strings.Join(users, "\n- "))
		message := &dg.MessageSend{
			Content: response,
			AllowedMentions: new(dg.MessageAllowedMentions),
		}
		s.ChannelMessageSendComplex(m.ChannelID, message)
		return
	}
}
