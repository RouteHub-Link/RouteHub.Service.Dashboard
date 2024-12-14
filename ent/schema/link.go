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

		field.JSON("link_content", &types.LinkContent{}).
			Default(&types.LinkContent{}).
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

func (Link) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return linkMutatorWithMQTT(next)
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete,
		),
	}
}

func linkMutatorWithMQTT(next ent.Mutator) ent.Mutator {
	return hook.LinkFunc(func(ctx context.Context, m *gen.LinkMutation) (ent.Value, error) {
		result, err := next.Mutate(ctx, m)
		if err != nil {
			return nil, err
		}

		link, _ := result.(*gen.Link)
		link.Edges.Hub = link.QueryHub().OnlyX(ctx)
		hub := link.Edges.Hub

		mqttClient, err := hubConnection.NewMQTTPublisher(hub.TCPAddress)
		if err != nil {
			// log error
			log.Printf("Error creating mqtt client: %v, TCPAddress:%s", err, hub.TCPAddress)
			return nil, err
		}

		if m.Op() == ent.OpCreate || m.Op() == ent.OpUpdate || m.Op() == ent.OpUpdateOne {
			err := mqttClient.PublishLinkSet(link)
			if err != nil {
				log.Printf("Error publishing link set: %v", err)
			}
		} else if m.Op() == ent.OpDelete {
			err := mqttClient.PublishLinkDel(link)
			if err != nil {
				log.Printf("Error publishing link delete: %v", err)
			}
		}

		return result, nil
	})
}
