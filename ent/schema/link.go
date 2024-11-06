package schema

import (
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go.jetify.com/typeid"
)

type LinkPrefix struct{}

func (LinkPrefix) Prefix() string {
	return "link"
}

type LinkID struct {
	typeid.TypeID[LinkPrefix]
}

func NewLinkID() LinkID {
	id, _ := typeid.New[LinkID]()
	return id
}

// Link holds the schema definition for the Link entity.
type Link struct {
	ent.Schema
}

// Fields of the Link.
func (Link) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", LinkID{}).
			Default(NewLinkID),

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
	}
}

// Edges of the Link.
func (Link) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("platform", Platform.Type).
			Ref("links"). // This points back to the 'links' edge in Platform
			Unique().     // One link belongs to one platform
			Required(),   // A platform is required for a link
	}
}
