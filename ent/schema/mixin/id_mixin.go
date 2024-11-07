package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"go.jetify.com/typeid"
)

type IDMixin struct {
	mixin.Schema
	Prefix string
}

func (idm IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(typeid.AnyID{}).
			DefaultFunc(func() typeid.AnyID {
				tid, _ := typeid.WithPrefix(idm.Prefix)
				return tid
			}),
	}
}

type Annotation struct {
	Prefix string
}

func (a Annotation) Name() string {
	return "TypeID"
}

func (idm IDMixin) Annotation() []schema.Annotation {
	return []schema.Annotation{
		Annotation{Prefix: idm.Prefix},
	}
}
