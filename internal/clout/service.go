package clout

import (
	dg "github.com/bwmarrin/discordgo"
)

type (
	Service interface {
		UpsertRelationship(from *dg.User, to []*dg.User)
		GetRelationships(*dg.User) []*Relation
	}

	service struct {
		members map[string]Relationship
		repo *Repository
	}
)

var serviceInstance Service

func GetService() Service {
	if serviceInstance == nil {
		serviceInstance = newService()
	}
	return serviceInstance
}

func newService() Service {
	service := new(service)
	service.members = make(map[string]Relationship)
	service.repo = NewRepository()
	return service
}

func (s *service) UpsertRelationship(from *dg.User, to []*dg.User) {
	var (
		relationship Relationship
		ok           bool
	)
	for _, t := range to {
		if t == nil || from.ID == t.ID {
			continue
		}
		if relationship, ok = s.members[t.ID]; !ok {
			relationship = NewRelationship(t)
			s.members[t.ID] = relationship
		}
		relationship.AddRelationship(from)
	}
	s.repo.Upsert(from, to)
}

func (s *service) GetRelationships(user *dg.User) (r []*Relation) {
	var (
		member Relationship
		ok     bool
	)
	if member, ok = s.members[user.ID]; !ok {
		return nil
	}
	return member.Relationships()
}
