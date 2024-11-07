// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"RouteHub.Service.Dashboard/ent/organization"
	"RouteHub.Service.Dashboard/ent/person"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

// PersonCreate is the builder for creating a Person entity.
type PersonCreate struct {
	config
	mutation *PersonMutation
	hooks    []Hook
}

// SetSubjectID sets the "subject_id" field.
func (pc *PersonCreate) SetSubjectID(s string) *PersonCreate {
	pc.mutation.SetSubjectID(s)
	return pc
}

// SetUserInfo sets the "user_info" field.
func (pc *PersonCreate) SetUserInfo(oi oidc.UserInfo) *PersonCreate {
	pc.mutation.SetUserInfo(oi)
	return pc
}

// SetIsActive sets the "is_active" field.
func (pc *PersonCreate) SetIsActive(b bool) *PersonCreate {
	pc.mutation.SetIsActive(b)
	return pc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (pc *PersonCreate) SetNillableIsActive(b *bool) *PersonCreate {
	if b != nil {
		pc.SetIsActive(*b)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PersonCreate) SetID(m mixin.ID) *PersonCreate {
	pc.mutation.SetID(m)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PersonCreate) SetNillableID(m *mixin.ID) *PersonCreate {
	if m != nil {
		pc.SetID(*m)
	}
	return pc
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (pc *PersonCreate) SetOrganizationID(id mixin.ID) *PersonCreate {
	pc.mutation.SetOrganizationID(id)
	return pc
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (pc *PersonCreate) SetNillableOrganizationID(id *mixin.ID) *PersonCreate {
	if id != nil {
		pc = pc.SetOrganizationID(*id)
	}
	return pc
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (pc *PersonCreate) SetOrganization(o *Organization) *PersonCreate {
	return pc.SetOrganizationID(o.ID)
}

// Mutation returns the PersonMutation object of the builder.
func (pc *PersonCreate) Mutation() *PersonMutation {
	return pc.mutation
}

// Save creates the Person in the database.
func (pc *PersonCreate) Save(ctx context.Context) (*Person, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PersonCreate) SaveX(ctx context.Context) *Person {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PersonCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PersonCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PersonCreate) defaults() {
	if _, ok := pc.mutation.IsActive(); !ok {
		v := person.DefaultIsActive
		pc.mutation.SetIsActive(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := person.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PersonCreate) check() error {
	if _, ok := pc.mutation.SubjectID(); !ok {
		return &ValidationError{Name: "subject_id", err: errors.New(`ent: missing required field "Person.subject_id"`)}
	}
	if v, ok := pc.mutation.SubjectID(); ok {
		if err := person.SubjectIDValidator(v); err != nil {
			return &ValidationError{Name: "subject_id", err: fmt.Errorf(`ent: validator failed for field "Person.subject_id": %w`, err)}
		}
	}
	if _, ok := pc.mutation.UserInfo(); !ok {
		return &ValidationError{Name: "user_info", err: errors.New(`ent: missing required field "Person.user_info"`)}
	}
	if _, ok := pc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "Person.is_active"`)}
	}
	return nil
}

func (pc *PersonCreate) sqlSave(ctx context.Context) (*Person, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(mixin.ID); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Person.ID type: %T", _spec.ID.Value)
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PersonCreate) createSpec() (*Person, *sqlgraph.CreateSpec) {
	var (
		_node = &Person{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(person.Table, sqlgraph.NewFieldSpec(person.FieldID, field.TypeString))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.SubjectID(); ok {
		_spec.SetField(person.FieldSubjectID, field.TypeString, value)
		_node.SubjectID = value
	}
	if value, ok := pc.mutation.UserInfo(); ok {
		_spec.SetField(person.FieldUserInfo, field.TypeJSON, value)
		_node.UserInfo = value
	}
	if value, ok := pc.mutation.IsActive(); ok {
		_spec.SetField(person.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if nodes := pc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   person.OrganizationTable,
			Columns: []string{person.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.organization_fk = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PersonCreateBulk is the builder for creating many Person entities in bulk.
type PersonCreateBulk struct {
	config
	err      error
	builders []*PersonCreate
}

// Save creates the Person entities in the database.
func (pcb *PersonCreateBulk) Save(ctx context.Context) ([]*Person, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Person, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PersonMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PersonCreateBulk) SaveX(ctx context.Context) []*Person {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PersonCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PersonCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
