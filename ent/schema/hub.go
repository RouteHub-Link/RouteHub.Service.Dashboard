package schema

import (
	"time"

	"RouteHub.Service.Dashboard/ent/schema/enums"
	enums_hub "RouteHub.Service.Dashboard/ent/schema/enums/hub"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Hub struct {
	ent.Schema
}

const (
	hubPrefix = "hub"
)

func (Hub) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{Prefix: hubPrefix},
	}
}

// Fields of the Hub.
func (Hub) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),

		field.String("slug").
			NotEmpty(),

		field.JSON("hub_details", types.HubDetails{}).
			Optional().
			Annotations(entgql.Type("HubDetails")),

		field.String("tcp_address").NotEmpty(),

		field.Enum("status").
			GoType(enums.StatusState(0)).
			Annotations(
				entsql.Default(enums.StatusInactive.String()),
				entgql.Type("StatusState"),
			),

		field.Enum("default_redirection").
			GoType(enums_hub.RedirectionOption(0)).
			Annotations(
				entsql.Default(enums_hub.Timed.String()),
				entgql.Type("RedirectionOption"),
			),

		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Hub.
func (Hub) Edges() []ent.Edge {
	return []ent.Edge{
		// One-to-One relationship: One Hub has one Domain (required for Hub)
		edge.To("domain", Domain.Type).
			StorageKey(edge.Column("domain_fk")). // Optional: customize the FK column name
			Unique().
			Required(), // Hub must have a Domain

		// Many-to-One relationship: A Hub belongs to one Organization
		edge.From("organization", Organization.Type).
			Ref("hubs"). // Refers to the Hubs edge in Organization
			Unique().
			Required(), // Hub must belong to an Organization

		// One Hub has many Links
		edge.To("links", Link.Type).
			StorageKey(edge.Column("link_fk")).
			Annotations(entgql.RelayConnection()), // Customize the foreign key if needed
	}
}
