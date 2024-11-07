// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"RouteHub.Service.Dashboard/ent/organization"
	"RouteHub.Service.Dashboard/ent/person"
	"RouteHub.Service.Dashboard/ent/predicate"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PersonUpdate is the builder for updating Person entities.
type PersonUpdate struct {
	config
	hooks    []Hook
	mutation *PersonMutation
}

// Where appends a list predicates to the PersonUpdate builder.
func (pu *PersonUpdate) Where(ps ...predicate.Person) *PersonUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetSubjectID sets the "subject_id" field.
func (pu *PersonUpdate) SetSubjectID(s string) *PersonUpdate {
	pu.mutation.SetSubjectID(s)
	return pu
}

// SetNillableSubjectID sets the "subject_id" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableSubjectID(s *string) *PersonUpdate {
	if s != nil {
		pu.SetSubjectID(*s)
	}
	return pu
}

// SetIsActive sets the "is_active" field.
func (pu *PersonUpdate) SetIsActive(b bool) *PersonUpdate {
	pu.mutation.SetIsActive(b)
	return pu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableIsActive(b *bool) *PersonUpdate {
	if b != nil {
		pu.SetIsActive(*b)
	}
	return pu
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (pu *PersonUpdate) SetOrganizationID(id mixin.ID) *PersonUpdate {
	pu.mutation.SetOrganizationID(id)
	return pu
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (pu *PersonUpdate) SetNillableOrganizationID(id *mixin.ID) *PersonUpdate {
	if id != nil {
		pu = pu.SetOrganizationID(*id)
	}
	return pu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (pu *PersonUpdate) SetOrganization(o *Organization) *PersonUpdate {
	return pu.SetOrganizationID(o.ID)
}

// Mutation returns the PersonMutation object of the builder.
func (pu *PersonUpdate) Mutation() *PersonMutation {
	return pu.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (pu *PersonUpdate) ClearOrganization() *PersonUpdate {
	pu.mutation.ClearOrganization()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PersonUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PersonUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PersonUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PersonUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PersonUpdate) check() error {
	if v, ok := pu.mutation.SubjectID(); ok {
		if err := person.SubjectIDValidator(v); err != nil {
			return &ValidationError{Name: "subject_id", err: fmt.Errorf(`ent: validator failed for field "Person.subject_id": %w`, err)}
		}
	}
	return nil
}

func (pu *PersonUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(person.Table, person.Columns, sqlgraph.NewFieldSpec(person.FieldID, field.TypeString))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.SubjectID(); ok {
		_spec.SetField(person.FieldSubjectID, field.TypeString, value)
	}
	if value, ok := pu.mutation.IsActive(); ok {
		_spec.SetField(person.FieldIsActive, field.TypeBool, value)
	}
	if pu.mutation.OrganizationCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OrganizationIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{person.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PersonUpdateOne is the builder for updating a single Person entity.
type PersonUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PersonMutation
}

// SetSubjectID sets the "subject_id" field.
func (puo *PersonUpdateOne) SetSubjectID(s string) *PersonUpdateOne {
	puo.mutation.SetSubjectID(s)
	return puo
}

// SetNillableSubjectID sets the "subject_id" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableSubjectID(s *string) *PersonUpdateOne {
	if s != nil {
		puo.SetSubjectID(*s)
	}
	return puo
}

// SetIsActive sets the "is_active" field.
func (puo *PersonUpdateOne) SetIsActive(b bool) *PersonUpdateOne {
	puo.mutation.SetIsActive(b)
	return puo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableIsActive(b *bool) *PersonUpdateOne {
	if b != nil {
		puo.SetIsActive(*b)
	}
	return puo
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (puo *PersonUpdateOne) SetOrganizationID(id mixin.ID) *PersonUpdateOne {
	puo.mutation.SetOrganizationID(id)
	return puo
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableOrganizationID(id *mixin.ID) *PersonUpdateOne {
	if id != nil {
		puo = puo.SetOrganizationID(*id)
	}
	return puo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (puo *PersonUpdateOne) SetOrganization(o *Organization) *PersonUpdateOne {
	return puo.SetOrganizationID(o.ID)
}

// Mutation returns the PersonMutation object of the builder.
func (puo *PersonUpdateOne) Mutation() *PersonMutation {
	return puo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (puo *PersonUpdateOne) ClearOrganization() *PersonUpdateOne {
	puo.mutation.ClearOrganization()
	return puo
}

// Where appends a list predicates to the PersonUpdate builder.
func (puo *PersonUpdateOne) Where(ps ...predicate.Person) *PersonUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PersonUpdateOne) Select(field string, fields ...string) *PersonUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Person entity.
func (puo *PersonUpdateOne) Save(ctx context.Context) (*Person, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PersonUpdateOne) SaveX(ctx context.Context) *Person {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PersonUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PersonUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PersonUpdateOne) check() error {
	if v, ok := puo.mutation.SubjectID(); ok {
		if err := person.SubjectIDValidator(v); err != nil {
			return &ValidationError{Name: "subject_id", err: fmt.Errorf(`ent: validator failed for field "Person.subject_id": %w`, err)}
		}
	}
	return nil
}

func (puo *PersonUpdateOne) sqlSave(ctx context.Context) (_node *Person, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(person.Table, person.Columns, sqlgraph.NewFieldSpec(person.FieldID, field.TypeString))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Person.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, person.FieldID)
		for _, f := range fields {
			if !person.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != person.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.SubjectID(); ok {
		_spec.SetField(person.FieldSubjectID, field.TypeString, value)
	}
	if value, ok := puo.mutation.IsActive(); ok {
		_spec.SetField(person.FieldIsActive, field.TypeBool, value)
	}
	if puo.mutation.OrganizationCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OrganizationIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Person{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{person.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
