package handlers

import (
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

func ParseEvent(s *dg.Session, m *dg.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	author := m.Author.ID
	mentions := make([]string, 0, len(m.Mentions))
	for _, m := range m.Mentions {
		mentions = append(mentions, m.ID)
	}
	fmt.Printf("author: %s mentioned %s\n", author, strings.Join(mentions, ", "))
}