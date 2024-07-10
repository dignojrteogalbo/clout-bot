package handlers

import (
	"clout-bot/internal/service"
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

func ParseEvent(s *dg.Session, m *dg.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	cloutService := service.GetService()
	mentions := append(m.Mentions, m.ReferencedMessage.Author)
	cloutService.UpsertRelationship(m.Author, mentions)
	if strings.Contains(m.Content, "clout") {
		relationships := cloutService.GetRelationships(m.Author)
		if len(relationships) <= 0 {
			s.ChannelMessageSend(m.ChannelID, "you have no clout...")
			return
		}
		users := make([]string, 0, len(relationships))
		for _, r := range relationships {
			users = append(users, r.Username)
		}
		response := fmt.Sprintf("these guys fw youu:\n- %s", strings.Join(users, "\n- "))
		s.ChannelMessageSend(m.ChannelID, response)
		return
	}
}
