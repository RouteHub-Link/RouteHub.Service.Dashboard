package schema

import (
	"time"

	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"entgo.io/contrib/entgql"
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

		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return []ent.Edge{
		// Many-to-many: One Person can belong to many Organizations
		edge.From("organizations", Organization.Type).
			Ref("persons").
			Annotations(
				entgql.RelayConnection(),
			),
	}
}

// Indexes of the Person.
func (Person) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("subject_id").
			Unique(),
	}
}
