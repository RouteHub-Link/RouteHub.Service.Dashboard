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
		// O2O relationship: One Domain has one Platform (optional for Domain)
		edge.From("platform", Platform.Type).
			Ref("domain").
			Unique(),

		// M2O relationship: A Domain belongs to one Organization
		edge.From("organization", Organization.Type).
			Ref("domains").
			Unique().
			Required(),
	}
}
