package schema

import (
	"time"

	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const (
	linkPrefix = "link"
)

func (Link) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{Prefix: linkPrefix},
	}
}

// Link holds the schema definition for the Link entity.
type Link struct {
	ent.Schema
}

// Fields of the Link.
func (Link) Fields() []ent.Field {
	return []ent.Field{
		field.String("target").
			NotEmpty(),

		field.String("path").
			NotEmpty().
			Unique(),

		field.JSON("link_content", types.LinkContent{}).
			Optional().
			Annotations(entgql.Type("LinkContent")),

		field.Enum("status").
			GoType(enums.StatusState(0)).
			Annotations(
				entgql.Type("StatusState"),
			),

		field.Enum("redirection_choice").
			GoType(enums.RedirectionChoice(0)).
			Annotations(
				entgql.Type("RedirectionChoice"),
			),

		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Link.
func (Link) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hub", Hub.Type).
			Ref("links"). // This points back to the 'links' edge in Hub
			Unique().     // One link belongs to one Hub
			Required().   // A Hub is required for a link
			Annotations(entgql.RelayConnection()),
	}
}
