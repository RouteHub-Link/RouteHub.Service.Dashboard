package mixin

import (
	"errors"
	"io"
	"strconv"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"go.jetify.com/typeid"
)

var (
	ErrInvalidType = errors.New("invalid type")
)

type IDMixin struct {
	mixin.Schema
	Prefix string
}

func (idm IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(ID("")).
			DefaultFunc(func() ID {
				tid, _ := typeid.WithPrefix(idm.Prefix)
				return ID(tid.String())
			}).
			Annotations(entgql.Type("ID")),
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

type ID string

func (id ID) New(prefix string) IDMixin {
	return IDMixin{Prefix: prefix}
}

func (id *ID) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	switch v := src.(type) {
	case string:
		if v == "" {
			return nil
		}
		*id = ID(v)
	case ID:
		*id = v
	case typeid.AnyID:
		*id = ID(v.String())
	default:
		return ErrInvalidType
	}

	return nil
}

func (id ID) Value() (interface{}, error) {
	return string(id), nil
}

func (id *ID) UnmarshalGQL(v interface{}) error {
	return id.Scan(v)
}

func (id ID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(string(id)))
}

func (id ID) String() string {
	return string(id)
}
