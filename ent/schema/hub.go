package schema

import (
	"context"
	"log"
	"time"

	gen "RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/ent/hook"
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"RouteHub.Service.Dashboard/features/hubConnection"
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
			GoType(enums.RedirectionChoice(0)).
			Annotations(
				entsql.Default(enums.RedirectionChoiceTimed.String()),
				entgql.Type("RedirectionChoice"),
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
			StorageKey(edge.Column("domain_fk")).
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
			Annotations(entgql.RelayConnection()),

		// One Hub has many Pages
		edge.To("pages", Page.Type).
			StorageKey(edge.Column("hub_fk")).
			Annotations(entgql.RelayConnection()),
	}
}

func (Hub) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hubMutatorWithMHQTT(next)
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete,
		),
	}
}

func hubMutatorWithMHQTT(next ent.Mutator) ent.Mutator {
	return hook.HubFunc(func(ctx context.Context, m *gen.HubMutation) (ent.Value, error) {

		if m.Op() == ent.OpCreate {
			brandName, _ := m.Name()
			m.SetHubDetails(types.HubDetails{
				NavbarDescription: types.NavbarDescription{
					BrandName: brandName,
				},
				FooterDescription: types.FooterDescription{
					ShowRouteHubBranding: true,
				},
				MetaDescription: types.MetaDescription{
					Title: brandName,
				},
			})
		}

		result, err := next.Mutate(ctx, m)
		if err != nil {
			return nil, err
		}

		hub, _ := result.(*gen.Hub)

		mqttClient, err := hubConnection.NewMQTTPublisher(hub.TCPAddress)
		if err != nil {
			log.Printf("Error creating mqtt client: %v, TCPAddress:%s", err, hub.TCPAddress)
			return nil, err
		}

		if m.Op() == ent.OpCreate || m.Op() == ent.OpUpdate || m.Op() == ent.OpUpdateOne {
			err := mqttClient.PublishPlatformSet(hub)
			if err != nil {
				log.Printf("Error publishing link set: %v", err)
			}
		}

		return hub, nil
	})
}
