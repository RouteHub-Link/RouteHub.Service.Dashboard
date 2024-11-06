package schema

import (
	"RouteHub.Service.Dashboard/ent/schema/enums"
	enums_hub "RouteHub.Service.Dashboard/ent/schema/enums/hub"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go.jetify.com/typeid"
)

type HubPrefix struct{}

func (HubPrefix) Prefix() string {
	return "Hub"
}

type HubID struct {
	typeid.TypeID[HubPrefix]
}

type Hub struct {
	ent.Schema
}

func NewHubID() HubID {
	id, _ := typeid.New[HubID]()
	return id
}

// Fields of the Hub.
func (Hub) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", HubID{}).
			Default(NewHubID),

		field.String("name").
			NotEmpty(),

		field.String("slug").
			NotEmpty(),

		field.JSON("Hub_details", types.HubDetails{}).
			Optional().
			Annotations(entgql.Type("HubDetails")),

		field.String("TCPAddress").NotEmpty(),

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
	}
}

// Edges of the Hub.
func (Hub) Edges() []ent.Edge {
	return []ent.Edge{
		// One-to-One relationship: One Hub has one Domain (required for Hub)
		edge.To("domain", Domain.Type). // Singular "domain" as this is O2O
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
			StorageKey(edge.Column("link_fk")), // Customize the foreign key if needed
	}
}
