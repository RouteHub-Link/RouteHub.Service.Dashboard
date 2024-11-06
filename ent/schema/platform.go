package schema

import (
	"RouteHub.Service.Dashboard/ent/schema/enums"
	enums_platform "RouteHub.Service.Dashboard/ent/schema/enums/platform"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go.jetify.com/typeid"
)

type PlatformPrefix struct{}

func (PlatformPrefix) Prefix() string {
	return "platform"
}

type PlatformID struct {
	typeid.TypeID[PlatformPrefix]
}

type Platform struct {
	ent.Schema
}

func NewPlatformID() PlatformID {
	id, _ := typeid.New[PlatformID]()
	return id
}

// Fields of the Platform.
func (Platform) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", PlatformID{}).
			Default(NewPlatformID),

		field.String("name").
			NotEmpty(),

		field.String("slug").
			NotEmpty(),

		field.JSON("platform_details", types.PlatformDetails{}).
			Optional().
			Annotations(entgql.Type("PlatformDetails")),

		field.String("TCPAddress").NotEmpty(),

		field.Enum("status").
			GoType(enums.StatusState(0)).
			Annotations(
				entsql.Default(enums.StatusInactive.String()),
				entgql.Type("StatusState"),
			),

		field.Enum("default_redirection").
			GoType(enums_platform.RedirectionOption(0)).
			Annotations(
				entsql.Default(enums_platform.Timed.String()),
				entgql.Type("RedirectionOption"),
			),
	}
}

// Edges of the Platform.
func (Platform) Edges() []ent.Edge {
	return []ent.Edge{
		// O2O relationship: One Platform has one Domain (required for Platform)
		edge.To("domain", Domain.Type).
			StorageKey(edge.Column("domain_fk")).
			Unique().
			Required(),

		// M2O relationship: A Platform belongs to one Organization
		edge.From("organization", Organization.Type).
			Ref("platforms").
			Unique().
			Required(),

		edge.To("links", Link.Type).
			StorageKey(edge.Column("link_fk")),
	}
}
