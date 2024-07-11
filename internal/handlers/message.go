package handlers

import (
	dg "github.com/bwmarrin/discordgo"
)

func HandleMessage(s *dg.Session, m *dg.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	RecordMentions(s, m)
	HandleResponse(s, m)
}
