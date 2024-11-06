package schema

import (
	"RouteHub.Service.Dashboard/ent/schema/types"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go.jetify.com/typeid"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

type OrganizationPrefix struct{}

func (OrganizationPrefix) Prefix() string {
	return "organization"
}

type OrganizationID struct {
	typeid.TypeID[OrganizationPrefix]
}

func NewOrganizationID() OrganizationID {
	id, _ := typeid.New[OrganizationID]()
	return id
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", OrganizationID{}).
			Default(NewOrganizationID),

		field.String("name").
			NotEmpty(),

		field.String("website").
			Optional(),

		field.String("description").
			Optional(),

		field.String("locagtion").
			Optional(),

		field.JSON("social_medias", types.SocialMedias{}).
			Optional().
			Annotations(entgql.Type("SocialMedias")),
	}

}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("domains", Domain.Type).
			StorageKey(edge.Column("domain_fk")),
		edge.To("platforms", Platform.Type).
			StorageKey(edge.Column("platform_fk")),
	}
}
