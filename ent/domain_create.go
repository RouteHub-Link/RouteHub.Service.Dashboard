// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	entdomain "RouteHub.Service.Dashboard/ent/domain"
	"RouteHub.Service.Dashboard/ent/organization"
	"RouteHub.Service.Dashboard/ent/schema/enums/domain"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DomainCreate is the builder for creating a Domain entity.
type DomainCreate struct {
	config
	mutation *DomainMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (dc *DomainCreate) SetName(s string) *DomainCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetURL sets the "url" field.
func (dc *DomainCreate) SetURL(s string) *DomainCreate {
	dc.mutation.SetURL(s)
	return dc
}

// SetStatus sets the "status" field.
func (dc *DomainCreate) SetStatus(ds domain.DomainState) *DomainCreate {
	dc.mutation.SetStatus(ds)
	return dc
}

// SetID sets the "id" field.
func (dc *DomainCreate) SetID(m mixin.ID) *DomainCreate {
	dc.mutation.SetID(m)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DomainCreate) SetNillableID(m *mixin.ID) *DomainCreate {
	if m != nil {
		dc.SetID(*m)
	}
	return dc
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (dc *DomainCreate) SetOrganizationID(id mixin.ID) *DomainCreate {
	dc.mutation.SetOrganizationID(id)
	return dc
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (dc *DomainCreate) SetOrganization(o *Organization) *DomainCreate {
	return dc.SetOrganizationID(o.ID)
}

// Mutation returns the DomainMutation object of the builder.
func (dc *DomainCreate) Mutation() *DomainMutation {
	return dc.mutation
}

// Save creates the Domain in the database.
func (dc *DomainCreate) Save(ctx context.Context) (*Domain, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DomainCreate) SaveX(ctx context.Context) *Domain {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DomainCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DomainCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DomainCreate) defaults() {
	if _, ok := dc.mutation.ID(); !ok {
		v := entdomain.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DomainCreate) check() error {
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Domain.name"`)}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := entdomain.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Domain.name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Domain.url"`)}
	}
	if v, ok := dc.mutation.URL(); ok {
		if err := entdomain.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Domain.url": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Domain.status"`)}
	}
	if v, ok := dc.mutation.Status(); ok {
		if err := entdomain.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Domain.status": %w`, err)}
		}
	}
	if len(dc.mutation.OrganizationIDs()) == 0 {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "Domain.organization"`)}
	}
	return nil
}

func (dc *DomainCreate) sqlSave(ctx context.Context) (*Domain, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(mixin.ID); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Domain.ID type: %T", _spec.ID.Value)
		}
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DomainCreate) createSpec() (*Domain, *sqlgraph.CreateSpec) {
	var (
		_node = &Domain{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(entdomain.Table, sqlgraph.NewFieldSpec(entdomain.FieldID, field.TypeString))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(entdomain.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.URL(); ok {
		_spec.SetField(entdomain.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := dc.mutation.Status(); ok {
		_spec.SetField(entdomain.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := dc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entdomain.OrganizationTable,
			Columns: []string{entdomain.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.organization_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DomainCreateBulk is the builder for creating many Domain entities in bulk.
type DomainCreateBulk struct {
	config
	err      error
	builders []*DomainCreate
}

// Save creates the Domain entities in the database.
func (dcb *DomainCreateBulk) Save(ctx context.Context) ([]*Domain, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Domain, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DomainMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DomainCreateBulk) SaveX(ctx context.Context) []*Domain {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DomainCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DomainCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
