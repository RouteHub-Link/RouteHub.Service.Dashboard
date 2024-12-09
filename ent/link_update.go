// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	enthub "RouteHub.Service.Dashboard/ent/hub"
	"RouteHub.Service.Dashboard/ent/link"
	"RouteHub.Service.Dashboard/ent/predicate"
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinkUpdate is the builder for updating Link entities.
type LinkUpdate struct {
	config
	hooks    []Hook
	mutation *LinkMutation
}

// Where appends a list predicates to the LinkUpdate builder.
func (lu *LinkUpdate) Where(ps ...predicate.Link) *LinkUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetTarget sets the "target" field.
func (lu *LinkUpdate) SetTarget(s string) *LinkUpdate {
	lu.mutation.SetTarget(s)
	return lu
}

// SetNillableTarget sets the "target" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableTarget(s *string) *LinkUpdate {
	if s != nil {
		lu.SetTarget(*s)
	}
	return lu
}

// SetPath sets the "path" field.
func (lu *LinkUpdate) SetPath(s string) *LinkUpdate {
	lu.mutation.SetPath(s)
	return lu
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (lu *LinkUpdate) SetNillablePath(s *string) *LinkUpdate {
	if s != nil {
		lu.SetPath(*s)
	}
	return lu
}

// SetLinkContent sets the "link_content" field.
func (lu *LinkUpdate) SetLinkContent(tc types.LinkContent) *LinkUpdate {
	lu.mutation.SetLinkContent(tc)
	return lu
}

// SetNillableLinkContent sets the "link_content" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableLinkContent(tc *types.LinkContent) *LinkUpdate {
	if tc != nil {
		lu.SetLinkContent(*tc)
	}
	return lu
}

// ClearLinkContent clears the value of the "link_content" field.
func (lu *LinkUpdate) ClearLinkContent() *LinkUpdate {
	lu.mutation.ClearLinkContent()
	return lu
}

// SetStatus sets the "status" field.
func (lu *LinkUpdate) SetStatus(es enums.StatusState) *LinkUpdate {
	lu.mutation.SetStatus(es)
	return lu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableStatus(es *enums.StatusState) *LinkUpdate {
	if es != nil {
		lu.SetStatus(*es)
	}
	return lu
}

// SetRedirectionChoice sets the "redirection_choice" field.
func (lu *LinkUpdate) SetRedirectionChoice(ec enums.RedirectionChoice) *LinkUpdate {
	lu.mutation.SetRedirectionChoice(ec)
	return lu
}

// SetNillableRedirectionChoice sets the "redirection_choice" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableRedirectionChoice(ec *enums.RedirectionChoice) *LinkUpdate {
	if ec != nil {
		lu.SetRedirectionChoice(*ec)
	}
	return lu
}

// SetCreatedAt sets the "created_at" field.
func (lu *LinkUpdate) SetCreatedAt(t time.Time) *LinkUpdate {
	lu.mutation.SetCreatedAt(t)
	return lu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableCreatedAt(t *time.Time) *LinkUpdate {
	if t != nil {
		lu.SetCreatedAt(*t)
	}
	return lu
}

// SetHubID sets the "hub" edge to the Hub entity by ID.
func (lu *LinkUpdate) SetHubID(id mixin.ID) *LinkUpdate {
	lu.mutation.SetHubID(id)
	return lu
}

// SetHub sets the "hub" edge to the Hub entity.
func (lu *LinkUpdate) SetHub(h *Hub) *LinkUpdate {
	return lu.SetHubID(h.ID)
}

// Mutation returns the LinkMutation object of the builder.
func (lu *LinkUpdate) Mutation() *LinkMutation {
	return lu.mutation
}

// ClearHub clears the "hub" edge to the Hub entity.
func (lu *LinkUpdate) ClearHub() *LinkUpdate {
	lu.mutation.ClearHub()
	return lu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LinkUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LinkUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LinkUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LinkUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LinkUpdate) check() error {
	if v, ok := lu.mutation.Target(); ok {
		if err := link.TargetValidator(v); err != nil {
			return &ValidationError{Name: "target", err: fmt.Errorf(`ent: validator failed for field "Link.target": %w`, err)}
		}
	}
	if v, ok := lu.mutation.Path(); ok {
		if err := link.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Link.path": %w`, err)}
		}
	}
	if v, ok := lu.mutation.Status(); ok {
		if err := link.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Link.status": %w`, err)}
		}
	}
	if v, ok := lu.mutation.RedirectionChoice(); ok {
		if err := link.RedirectionChoiceValidator(v); err != nil {
			return &ValidationError{Name: "redirection_choice", err: fmt.Errorf(`ent: validator failed for field "Link.redirection_choice": %w`, err)}
		}
	}
	if lu.mutation.HubCleared() && len(lu.mutation.HubIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Link.hub"`)
	}
	return nil
}

func (lu *LinkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(link.Table, link.Columns, sqlgraph.NewFieldSpec(link.FieldID, field.TypeString))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Target(); ok {
		_spec.SetField(link.FieldTarget, field.TypeString, value)
	}
	if value, ok := lu.mutation.Path(); ok {
		_spec.SetField(link.FieldPath, field.TypeString, value)
	}
	if value, ok := lu.mutation.LinkContent(); ok {
		_spec.SetField(link.FieldLinkContent, field.TypeJSON, value)
	}
	if lu.mutation.LinkContentCleared() {
		_spec.ClearField(link.FieldLinkContent, field.TypeJSON)
	}
	if value, ok := lu.mutation.Status(); ok {
		_spec.SetField(link.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := lu.mutation.RedirectionChoice(); ok {
		_spec.SetField(link.FieldRedirectionChoice, field.TypeEnum, value)
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.SetField(link.FieldCreatedAt, field.TypeTime, value)
	}
	if lu.mutation.HubCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   link.HubTable,
			Columns: []string{link.HubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(enthub.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.HubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   link.HubTable,
			Columns: []string{link.HubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(enthub.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LinkUpdateOne is the builder for updating a single Link entity.
type LinkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LinkMutation
}

// SetTarget sets the "target" field.
func (luo *LinkUpdateOne) SetTarget(s string) *LinkUpdateOne {
	luo.mutation.SetTarget(s)
	return luo
}

// SetNillableTarget sets the "target" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableTarget(s *string) *LinkUpdateOne {
	if s != nil {
		luo.SetTarget(*s)
	}
	return luo
}

// SetPath sets the "path" field.
func (luo *LinkUpdateOne) SetPath(s string) *LinkUpdateOne {
	luo.mutation.SetPath(s)
	return luo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillablePath(s *string) *LinkUpdateOne {
	if s != nil {
		luo.SetPath(*s)
	}
	return luo
}

// SetLinkContent sets the "link_content" field.
func (luo *LinkUpdateOne) SetLinkContent(tc types.LinkContent) *LinkUpdateOne {
	luo.mutation.SetLinkContent(tc)
	return luo
}

// SetNillableLinkContent sets the "link_content" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableLinkContent(tc *types.LinkContent) *LinkUpdateOne {
	if tc != nil {
		luo.SetLinkContent(*tc)
	}
	return luo
}

// ClearLinkContent clears the value of the "link_content" field.
func (luo *LinkUpdateOne) ClearLinkContent() *LinkUpdateOne {
	luo.mutation.ClearLinkContent()
	return luo
}

// SetStatus sets the "status" field.
func (luo *LinkUpdateOne) SetStatus(es enums.StatusState) *LinkUpdateOne {
	luo.mutation.SetStatus(es)
	return luo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableStatus(es *enums.StatusState) *LinkUpdateOne {
	if es != nil {
		luo.SetStatus(*es)
	}
	return luo
}

// SetRedirectionChoice sets the "redirection_choice" field.
func (luo *LinkUpdateOne) SetRedirectionChoice(ec enums.RedirectionChoice) *LinkUpdateOne {
	luo.mutation.SetRedirectionChoice(ec)
	return luo
}

// SetNillableRedirectionChoice sets the "redirection_choice" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableRedirectionChoice(ec *enums.RedirectionChoice) *LinkUpdateOne {
	if ec != nil {
		luo.SetRedirectionChoice(*ec)
	}
	return luo
}

// SetCreatedAt sets the "created_at" field.
func (luo *LinkUpdateOne) SetCreatedAt(t time.Time) *LinkUpdateOne {
	luo.mutation.SetCreatedAt(t)
	return luo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableCreatedAt(t *time.Time) *LinkUpdateOne {
	if t != nil {
		luo.SetCreatedAt(*t)
	}
	return luo
}

// SetHubID sets the "hub" edge to the Hub entity by ID.
func (luo *LinkUpdateOne) SetHubID(id mixin.ID) *LinkUpdateOne {
	luo.mutation.SetHubID(id)
	return luo
}

// SetHub sets the "hub" edge to the Hub entity.
func (luo *LinkUpdateOne) SetHub(h *Hub) *LinkUpdateOne {
	return luo.SetHubID(h.ID)
}

// Mutation returns the LinkMutation object of the builder.
func (luo *LinkUpdateOne) Mutation() *LinkMutation {
	return luo.mutation
}

// ClearHub clears the "hub" edge to the Hub entity.
func (luo *LinkUpdateOne) ClearHub() *LinkUpdateOne {
	luo.mutation.ClearHub()
	return luo
}

// Where appends a list predicates to the LinkUpdate builder.
func (luo *LinkUpdateOne) Where(ps ...predicate.Link) *LinkUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LinkUpdateOne) Select(field string, fields ...string) *LinkUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Link entity.
func (luo *LinkUpdateOne) Save(ctx context.Context) (*Link, error) {
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LinkUpdateOne) SaveX(ctx context.Context) *Link {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LinkUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LinkUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LinkUpdateOne) check() error {
	if v, ok := luo.mutation.Target(); ok {
		if err := link.TargetValidator(v); err != nil {
			return &ValidationError{Name: "target", err: fmt.Errorf(`ent: validator failed for field "Link.target": %w`, err)}
		}
	}
	if v, ok := luo.mutation.Path(); ok {
		if err := link.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Link.path": %w`, err)}
		}
	}
	if v, ok := luo.mutation.Status(); ok {
		if err := link.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Link.status": %w`, err)}
		}
	}
	if v, ok := luo.mutation.RedirectionChoice(); ok {
		if err := link.RedirectionChoiceValidator(v); err != nil {
			return &ValidationError{Name: "redirection_choice", err: fmt.Errorf(`ent: validator failed for field "Link.redirection_choice": %w`, err)}
		}
	}
	if luo.mutation.HubCleared() && len(luo.mutation.HubIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Link.hub"`)
	}
	return nil
}

func (luo *LinkUpdateOne) sqlSave(ctx context.Context) (_node *Link, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(link.Table, link.Columns, sqlgraph.NewFieldSpec(link.FieldID, field.TypeString))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Link.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, link.FieldID)
		for _, f := range fields {
			if !link.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != link.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Target(); ok {
		_spec.SetField(link.FieldTarget, field.TypeString, value)
	}
	if value, ok := luo.mutation.Path(); ok {
		_spec.SetField(link.FieldPath, field.TypeString, value)
	}
	if value, ok := luo.mutation.LinkContent(); ok {
		_spec.SetField(link.FieldLinkContent, field.TypeJSON, value)
	}
	if luo.mutation.LinkContentCleared() {
		_spec.ClearField(link.FieldLinkContent, field.TypeJSON)
	}
	if value, ok := luo.mutation.Status(); ok {
		_spec.SetField(link.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := luo.mutation.RedirectionChoice(); ok {
		_spec.SetField(link.FieldRedirectionChoice, field.TypeEnum, value)
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.SetField(link.FieldCreatedAt, field.TypeTime, value)
	}
	if luo.mutation.HubCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   link.HubTable,
			Columns: []string{link.HubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(enthub.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.HubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   link.HubTable,
			Columns: []string{link.HubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(enthub.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Link{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
