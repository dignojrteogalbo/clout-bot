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
	service.repo = NewRepository()
	return service
}

func (s *service) UpsertRelationship(from *dg.User, to []*dg.User) {
	s.repo.Upsert(from, to)
}

func (s *service) GetRelationships(user *dg.User) (r []*Relation) {
	relationDAO := s.repo.GetMany(user)
	if relationDAO == nil {
		return nil
	}
	r = make([]*Relation, 0)
	for _, i := range relationDAO {
		relation := &Relation{&dg.User{ID: i.From}, uint(i.Count)}
		r = append(r, relation)
	}
	return r
}
