package schema

import (
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

// Person holds the schema definition for the Person entity.
type Person struct {
	ent.Schema
}

const (
	personPrefix = "person"
)

func (Person) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{Prefix: personPrefix},
	}
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.String("subject_id").
			NotEmpty().
			Unique(),

		field.JSON("user_info", oidc.UserInfo{}).
			Immutable(),

		field.Bool("is_active").
			Default(true),
	}
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return []ent.Edge{
		// One Person has one Organization.
		edge.To("organization", Organization.Type).
			StorageKey(edge.Column("organization_fk")).
			Unique(),
	}
}

// Indexes of the Person.
func (Person) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("subject_id").
			Unique(),
	}
}
