package handlers

import (
	"clout-bot/internal/clout"

	dg "github.com/bwmarrin/discordgo"
)

func mentionsExcludingIds(message *dg.Message, ids ...*dg.User) []*dg.User {
	mentions := make([]*dg.User, 0)
	exclude := make(map[string]bool)
	for _, i := range ids {
		exclude[i.ID] = true
	}
	for _, m := range message.Mentions {
		if !exclude[m.ID] {
			mentions = append(mentions, m)
		}
	}
	if message.ReferencedMessage != nil && !exclude[message.ReferencedMessage.Author.ID] {
		mentions = append(mentions, message.ReferencedMessage.Author)
	}
	return mentions
}

func RecordMentions(s *dg.Session, m *dg.MessageCreate) {
	service := clout.GetService()
	mentions := mentionsExcludingIds(m.Message, s.State.User, m.Author)
	if len(mentions) > 0 {
		service.UpsertRelationship(m.Author, mentions)
	}
}