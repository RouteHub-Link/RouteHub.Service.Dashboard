package schema

import (
	enums_domain "RouteHub.Service.Dashboard/ent/schema/enums/domain"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Domain holds the schema definition for the Domain entity.
type Domain struct {
	ent.Schema
}

const (
	domainPrefix = "domain"
)

func (Domain) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Add the mixin IDMixin to the Domain schema
		mixin.IDMixin{Prefix: domainPrefix},
	}
}

// Fields of the Domains.
func (Domain) Fields() []ent.Field {
	return []ent.Field{
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
		// Many-to-One relationship: A Domain belongs to one Organization
		edge.From("organization", Organization.Type).
			Ref("domains"). // Refers to the domains edge in Organization
			Unique().
			Required(), // Domain must belong to an Organization
	}
}
