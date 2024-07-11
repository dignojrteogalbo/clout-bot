package clout

import dg "github.com/bwmarrin/discordgo"

type (
	Relationship interface {
		Clout() uint
		AddRelationship(to *dg.User)
		Relationships() []*Relation
	}

	member struct {
		from *dg.User
		relationships map[string]*Relation
	}

	Relation struct {
		User *dg.User
		Count uint
	}
)

func NewRelation(to *dg.User) *Relation {
	r := new(Relation)
	r.User = to
	return r
}

func NewRelationship(user *dg.User) Relationship {
	relations := new(member)
	relations.from = user
	relations.relationships = make(map[string]*Relation)
	return relations
}

func (m *member) Clout() (clout uint) {
	for _, mentions := range m.relationships {
		clout += mentions.Count
	}
	return clout
}

func (m *member) AddRelationship(to *dg.User) {
	if _, ok := m.relationships[to.ID]; !ok {
		m.relationships[to.ID] = NewRelation(to)
	}
	m.relationships[to.ID].Count += 1
}

func (m *member) Relationships() (relationships []*Relation) {
	relationships = make([]*Relation, 0, len(m.relationships))
	for _, r := range m.relationships {
		relationships = append(relationships, r)
	}
	return relationships
}
