package service

import dg "github.com/bwmarrin/discordgo"

type (
	Relationship interface {
		Clout() uint
		AddRelationship(to *dg.User)
		Relationships() []*dg.User
	}

	member struct {
		from *dg.User
		relationships map[string]*relation
	}

	relation struct {
		user *dg.User
		count uint
	}
)

func NewRelation(to *dg.User) *relation {
	r := new(relation)
	r.user = to
	return r
}

func NewRelationship(user *dg.User) Relationship {
	relations := new(member)
	relations.from = user
	relations.relationships = make(map[string]*relation)
	return relations
}

func (m *member) Clout() (clout uint) {
	for _, mentions := range m.relationships {
		clout += mentions.count
	}
	return clout
}

func (m *member) AddRelationship(to *dg.User) {
	if _, ok := m.relationships[to.ID]; !ok {
		m.relationships[to.ID] = NewRelation(to)
	}
	m.relationships[to.ID].count += 1
}

func (m *member) Relationships() (r []*dg.User) {
	r = make([]*dg.User, 0, len(m.relationships))
	for _, u := range m.relationships {
		r = append(r, u.user)
	}
	return r
}
