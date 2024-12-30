package schema

import (
	"context"
	"log"
	"time"

	gen "RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/features/hubConnection"

	"RouteHub.Service.Dashboard/ent/hook"
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Page struct {
	ent.Schema
}

const (
	pagePrefix = "page"
)

func (Page) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{Prefix: pagePrefix},
	}
}

func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),

		field.String("slug").
			NotEmpty(),

		field.String("page_description").
			Optional(),

		field.String("page_content_json").
			Optional(),

		field.String("page_content_html").
			Optional(),

		field.Enum("status").
			GoType(enums.StatusState(0)).
			Annotations(
				entsql.Default(enums.StatusInactive.String()),
				entgql.Type("StatusState"),
			),

		field.JSON("meta_description", &types.MetaDescription{}).
			Optional().
			Annotations(entgql.Type("MetaDescription")),

		field.Time("created_at").
			Default(time.Now),
	}
}

func (Page) Edges() []ent.Edge {
	return []ent.Edge{
		// edge to the Hub, hub could have many pages but a page belongs to a single hub.
		edge.From("hub", Hub.Type).
			Ref("pages").
			Unique().
			Required(),
	}
}

func (Page) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return pageMutatorWithMHQTT(next)
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete,
		),
	}
}

func pageMutatorWithMHQTT(next ent.Mutator) ent.Mutator {
	return hook.PageFunc(func(ctx context.Context, m *gen.PageMutation) (ent.Value, error) {
		result, err := next.Mutate(ctx, m)
		if err != nil {
			return nil, err
		}
		page, _ := result.(*gen.Page)

		page.Edges.Hub = page.QueryHub().OnlyX(ctx)
		hub := page.Edges.Hub

		mqttClient, err := hubConnection.NewMQTTPublisher(hub.TCPAddress)
		if err != nil {
			log.Printf("Error creating mqtt client: %v, TCPAddress:%s", err, hub.TCPAddress)
			return nil, err
		}

		if m.Op() == ent.OpCreate || m.Op() == ent.OpUpdate || m.Op() == ent.OpUpdateOne {
			err := mqttClient.PublishPageSet(page)
			if err != nil {
				log.Printf("Error publishing page set: %v", err)
			}
		} else if m.Op() == ent.OpDelete {
			err := mqttClient.PublishPageDel(page)
			if err != nil {
				log.Printf("Error publishing page delete: %v", err)
			}
		}

		return page, nil
	})
}
