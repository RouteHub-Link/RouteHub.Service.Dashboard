// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"RouteHub.Service.Dashboard/ent/hub"
	"RouteHub.Service.Dashboard/ent/link"
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinkCreate is the builder for creating a Link entity.
type LinkCreate struct {
	config
	mutation *LinkMutation
	hooks    []Hook
}

// SetTarget sets the "target" field.
func (lc *LinkCreate) SetTarget(s string) *LinkCreate {
	lc.mutation.SetTarget(s)
	return lc
}

// SetPath sets the "path" field.
func (lc *LinkCreate) SetPath(s string) *LinkCreate {
	lc.mutation.SetPath(s)
	return lc
}

// SetLinkContent sets the "link_content" field.
func (lc *LinkCreate) SetLinkContent(tc *types.LinkContent) *LinkCreate {
	lc.mutation.SetLinkContent(tc)
	return lc
}

// SetStatus sets the "status" field.
func (lc *LinkCreate) SetStatus(es enums.StatusState) *LinkCreate {
	lc.mutation.SetStatus(es)
	return lc
}

// SetRedirectionChoice sets the "redirection_choice" field.
func (lc *LinkCreate) SetRedirectionChoice(ec enums.RedirectionChoice) *LinkCreate {
	lc.mutation.SetRedirectionChoice(ec)
	return lc
}

// SetCreatedAt sets the "created_at" field.
func (lc *LinkCreate) SetCreatedAt(t time.Time) *LinkCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LinkCreate) SetNillableCreatedAt(t *time.Time) *LinkCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LinkCreate) SetID(m mixin.ID) *LinkCreate {
	lc.mutation.SetID(m)
	return lc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lc *LinkCreate) SetNillableID(m *mixin.ID) *LinkCreate {
	if m != nil {
		lc.SetID(*m)
	}
	return lc
}

// SetHubID sets the "hub" edge to the Hub entity by ID.
func (lc *LinkCreate) SetHubID(id mixin.ID) *LinkCreate {
	lc.mutation.SetHubID(id)
	return lc
}

// SetHub sets the "hub" edge to the Hub entity.
func (lc *LinkCreate) SetHub(h *Hub) *LinkCreate {
	return lc.SetHubID(h.ID)
}

// Mutation returns the LinkMutation object of the builder.
func (lc *LinkCreate) Mutation() *LinkMutation {
	return lc.mutation
}

// Save creates the Link in the database.
func (lc *LinkCreate) Save(ctx context.Context) (*Link, error) {
	if err := lc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LinkCreate) SaveX(ctx context.Context) *Link {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LinkCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LinkCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LinkCreate) defaults() error {
	if _, ok := lc.mutation.LinkContent(); !ok {
		v := link.DefaultLinkContent
		lc.mutation.SetLinkContent(v)
	}
	if _, ok := lc.mutation.CreatedAt(); !ok {
		if link.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized link.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := link.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.ID(); !ok {
		if link.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized link.DefaultID (forgotten import ent/runtime?)")
		}
		v := link.DefaultID()
		lc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (lc *LinkCreate) check() error {
	if _, ok := lc.mutation.Target(); !ok {
		return &ValidationError{Name: "target", err: errors.New(`ent: missing required field "Link.target"`)}
	}
	if v, ok := lc.mutation.Target(); ok {
		if err := link.TargetValidator(v); err != nil {
			return &ValidationError{Name: "target", err: fmt.Errorf(`ent: validator failed for field "Link.target": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Link.path"`)}
	}
	if v, ok := lc.mutation.Path(); ok {
		if err := link.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Link.path": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Link.status"`)}
	}
	if v, ok := lc.mutation.Status(); ok {
		if err := link.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Link.status": %w`, err)}
		}
	}
	if _, ok := lc.mutation.RedirectionChoice(); !ok {
		return &ValidationError{Name: "redirection_choice", err: errors.New(`ent: missing required field "Link.redirection_choice"`)}
	}
	if v, ok := lc.mutation.RedirectionChoice(); ok {
		if err := link.RedirectionChoiceValidator(v); err != nil {
			return &ValidationError{Name: "redirection_choice", err: fmt.Errorf(`ent: validator failed for field "Link.redirection_choice": %w`, err)}
		}
	}
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Link.created_at"`)}
	}
	if len(lc.mutation.HubIDs()) == 0 {
		return &ValidationError{Name: "hub", err: errors.New(`ent: missing required edge "Link.hub"`)}
	}
	return nil
}

func (lc *LinkCreate) sqlSave(ctx context.Context) (*Link, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(mixin.ID); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Link.ID type: %T", _spec.ID.Value)
		}
	}
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LinkCreate) createSpec() (*Link, *sqlgraph.CreateSpec) {
	var (
		_node = &Link{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(link.Table, sqlgraph.NewFieldSpec(link.FieldID, field.TypeString))
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lc.mutation.Target(); ok {
		_spec.SetField(link.FieldTarget, field.TypeString, value)
		_node.Target = value
	}
	if value, ok := lc.mutation.Path(); ok {
		_spec.SetField(link.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := lc.mutation.LinkContent(); ok {
		_spec.SetField(link.FieldLinkContent, field.TypeJSON, value)
		_node.LinkContent = value
	}
	if value, ok := lc.mutation.Status(); ok {
		_spec.SetField(link.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := lc.mutation.RedirectionChoice(); ok {
		_spec.SetField(link.FieldRedirectionChoice, field.TypeEnum, value)
		_node.RedirectionChoice = value
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(link.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := lc.mutation.HubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   link.HubTable,
			Columns: []string{link.HubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hub.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.link_fk = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LinkCreateBulk is the builder for creating many Link entities in bulk.
type LinkCreateBulk struct {
	config
	err      error
	builders []*LinkCreate
}

// Save creates the Link entities in the database.
func (lcb *LinkCreateBulk) Save(ctx context.Context) ([]*Link, error) {
	if lcb.err != nil {
		return nil, lcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Link, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LinkMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LinkCreateBulk) SaveX(ctx context.Context) []*Link {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LinkCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LinkCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
