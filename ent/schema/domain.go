package schema

import (
	enums_domain "RouteHub.Service.Dashboard/ent/schema/enums/domain"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go.jetify.com/typeid"
)

// Domain holds the schema definition for the Domain entity.
type Domain struct {
	ent.Schema
}

type DomainPrefix struct{}

func (DomainPrefix) Prefix() string {
	return "domain"
}

type DomainID struct {
	typeid.TypeID[DomainPrefix]
}

func NewDomainID() DomainID {
	id, _ := typeid.New[DomainID]()
	return id
}

// Fields of the Domains.
func (Domain) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", DomainID{}).
			Default(NewDomainID),

		field.String("name").
			NotEmpty(),

		field.String("url").
			NotEmpty().
			Unique(),

		field.Enum("status").
			GoType(enums_domain.DomainState(0)).
			Annotations(
				entsql.Default(enums_domain.Passive.String()),
				entgql.Type("DomainState"),
			),
	}
}

// Edges of the Domains.
func (Domain) Edges() []ent.Edge {
	return []ent.Edge{
		// One-to-One relationship: One Domain has one Hub (Hub is optional for Domain)
		edge.From("hub", Hub.Type).
			Ref("domain"). // Matches the "domain" edge in Hub
			Unique(),      // One-to-One relationship
		// Hub is optional for Domain (don't use Required())

		// Many-to-One relationship: A Domain belongs to one Organization
		edge.From("organization", Organization.Type).
			Ref("domains"). // Refers to the domains edge in Organization
			Unique().
			Required(), // Domain must belong to an Organization
	}
}
