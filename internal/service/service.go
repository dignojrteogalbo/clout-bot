package service

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"
)

type (
	CloutService interface {
		UpsertRelationship(from *dg.User, to []*dg.User)
		GetRelationships(*dg.User) []*dg.User
	}

	service struct {
		members map[string]Relationship
	}
)

var instance CloutService

func GetService() CloutService {
	if instance == nil {
		instance = newService()
	}
	return instance
}

func newService() CloutService {
	service := new(service)
	service.members = make(map[string]Relationship)
	return service
}

func (s *service) UpsertRelationship(from *dg.User, to []*dg.User) {
	var (
		relationship Relationship
		ok bool
	)
	if relationship, ok = s.members[from.ID]; !ok {
		relationship = NewRelationship(from)
		s.members[from.ID] = relationship
	}
	for _, t := range to {
		if t != nil && from.ID != t.ID {
			relationship.AddRelationship(t)
		}
	}
	fmt.Println(relationship.Clout())
}

func (s *service) GetRelationships(user *dg.User) (r []*dg.User) {
	var (
		member Relationship
		ok     bool
	)
	if member, ok = s.members[user.ID]; !ok {
		return nil
	}
	return member.Relationships()
}
