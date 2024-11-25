package schema

import (
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

const (
	organizationPrefix = "org"
)

func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{Prefix: organizationPrefix},
	}
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),

		field.String("website").
			Optional(),

		field.String("description").
			Optional(),

		field.String("location").
			Optional(),

		field.JSON("social_medias", types.SocialMedias{}).
			Optional().
			Annotations(entgql.Type("SocialMedias")),
	}

}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		// One Organization has many Domains
		edge.To("domains", Domain.Type).
			Annotations(entgql.RelayConnection()).
			StorageKey(edge.Column("organization_id")),

		// One Organization has many Hubs
		edge.To("hubs", Hub.Type).
			Annotations(entgql.RelayConnection()).
			StorageKey(edge.Column("organization_id")),

		// Many-to-many: One Organization has many Persons via the OrganizationPerson join table
		edge.To("persons", Person.Type).
			Annotations(entgql.RelayConnection()),
	}
}
