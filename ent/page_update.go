// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"RouteHub.Service.Dashboard/ent/hub"
	"RouteHub.Service.Dashboard/ent/page"
	"RouteHub.Service.Dashboard/ent/predicate"
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PageUpdate is the builder for updating Page entities.
type PageUpdate struct {
	config
	hooks    []Hook
	mutation *PageMutation
}

// Where appends a list predicates to the PageUpdate builder.
func (pu *PageUpdate) Where(ps ...predicate.Page) *PageUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PageUpdate) SetName(s string) *PageUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PageUpdate) SetNillableName(s *string) *PageUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetSlug sets the "slug" field.
func (pu *PageUpdate) SetSlug(s string) *PageUpdate {
	pu.mutation.SetSlug(s)
	return pu
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (pu *PageUpdate) SetNillableSlug(s *string) *PageUpdate {
	if s != nil {
		pu.SetSlug(*s)
	}
	return pu
}

// SetPageDescription sets the "page_description" field.
func (pu *PageUpdate) SetPageDescription(s string) *PageUpdate {
	pu.mutation.SetPageDescription(s)
	return pu
}

// SetNillablePageDescription sets the "page_description" field if the given value is not nil.
func (pu *PageUpdate) SetNillablePageDescription(s *string) *PageUpdate {
	if s != nil {
		pu.SetPageDescription(*s)
	}
	return pu
}

// ClearPageDescription clears the value of the "page_description" field.
func (pu *PageUpdate) ClearPageDescription() *PageUpdate {
	pu.mutation.ClearPageDescription()
	return pu
}

// SetPageContentJSON sets the "page_content_json" field.
func (pu *PageUpdate) SetPageContentJSON(s string) *PageUpdate {
	pu.mutation.SetPageContentJSON(s)
	return pu
}

// SetNillablePageContentJSON sets the "page_content_json" field if the given value is not nil.
func (pu *PageUpdate) SetNillablePageContentJSON(s *string) *PageUpdate {
	if s != nil {
		pu.SetPageContentJSON(*s)
	}
	return pu
}

// ClearPageContentJSON clears the value of the "page_content_json" field.
func (pu *PageUpdate) ClearPageContentJSON() *PageUpdate {
	pu.mutation.ClearPageContentJSON()
	return pu
}

// SetPageContentHTML sets the "page_content_html" field.
func (pu *PageUpdate) SetPageContentHTML(s string) *PageUpdate {
	pu.mutation.SetPageContentHTML(s)
	return pu
}

// SetNillablePageContentHTML sets the "page_content_html" field if the given value is not nil.
func (pu *PageUpdate) SetNillablePageContentHTML(s *string) *PageUpdate {
	if s != nil {
		pu.SetPageContentHTML(*s)
	}
	return pu
}

// ClearPageContentHTML clears the value of the "page_content_html" field.
func (pu *PageUpdate) ClearPageContentHTML() *PageUpdate {
	pu.mutation.ClearPageContentHTML()
	return pu
}

// SetStatus sets the "status" field.
func (pu *PageUpdate) SetStatus(es enums.StatusState) *PageUpdate {
	pu.mutation.SetStatus(es)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PageUpdate) SetNillableStatus(es *enums.StatusState) *PageUpdate {
	if es != nil {
		pu.SetStatus(*es)
	}
	return pu
}

// SetMetaDescription sets the "meta_description" field.
func (pu *PageUpdate) SetMetaDescription(td *types.MetaDescription) *PageUpdate {
	pu.mutation.SetMetaDescription(td)
	return pu
}

// ClearMetaDescription clears the value of the "meta_description" field.
func (pu *PageUpdate) ClearMetaDescription() *PageUpdate {
	pu.mutation.ClearMetaDescription()
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PageUpdate) SetCreatedAt(t time.Time) *PageUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PageUpdate) SetNillableCreatedAt(t *time.Time) *PageUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetHubID sets the "hub" edge to the Hub entity by ID.
func (pu *PageUpdate) SetHubID(id mixin.ID) *PageUpdate {
	pu.mutation.SetHubID(id)
	return pu
}

// SetHub sets the "hub" edge to the Hub entity.
func (pu *PageUpdate) SetHub(h *Hub) *PageUpdate {
	return pu.SetHubID(h.ID)
}

// Mutation returns the PageMutation object of the builder.
func (pu *PageUpdate) Mutation() *PageMutation {
	return pu.mutation
}

// ClearHub clears the "hub" edge to the Hub entity.
func (pu *PageUpdate) ClearHub() *PageUpdate {
	pu.mutation.ClearHub()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PageUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PageUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PageUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PageUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := page.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Page.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Slug(); ok {
		if err := page.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Page.slug": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Status(); ok {
		if err := page.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Page.status": %w`, err)}
		}
	}
	if pu.mutation.HubCleared() && len(pu.mutation.HubIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Page.hub"`)
	}
	return nil
}

func (pu *PageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(page.Table, page.Columns, sqlgraph.NewFieldSpec(page.FieldID, field.TypeString))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(page.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Slug(); ok {
		_spec.SetField(page.FieldSlug, field.TypeString, value)
	}
	if value, ok := pu.mutation.PageDescription(); ok {
		_spec.SetField(page.FieldPageDescription, field.TypeString, value)
	}
	if pu.mutation.PageDescriptionCleared() {
		_spec.ClearField(page.FieldPageDescription, field.TypeString)
	}
	if value, ok := pu.mutation.PageContentJSON(); ok {
		_spec.SetField(page.FieldPageContentJSON, field.TypeString, value)
	}
	if pu.mutation.PageContentJSONCleared() {
		_spec.ClearField(page.FieldPageContentJSON, field.TypeString)
	}
	if value, ok := pu.mutation.PageContentHTML(); ok {
		_spec.SetField(page.FieldPageContentHTML, field.TypeString, value)
	}
	if pu.mutation.PageContentHTMLCleared() {
		_spec.ClearField(page.FieldPageContentHTML, field.TypeString)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(page.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.MetaDescription(); ok {
		_spec.SetField(page.FieldMetaDescription, field.TypeJSON, value)
	}
	if pu.mutation.MetaDescriptionCleared() {
		_spec.ClearField(page.FieldMetaDescription, field.TypeJSON)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(page.FieldCreatedAt, field.TypeTime, value)
	}
	if pu.mutation.HubCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   page.HubTable,
			Columns: []string{page.HubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hub.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.HubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   page.HubTable,
			Columns: []string{page.HubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hub.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{page.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PageUpdateOne is the builder for updating a single Page entity.
type PageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PageMutation
}

// SetName sets the "name" field.
func (puo *PageUpdateOne) SetName(s string) *PageUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillableName(s *string) *PageUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetSlug sets the "slug" field.
func (puo *PageUpdateOne) SetSlug(s string) *PageUpdateOne {
	puo.mutation.SetSlug(s)
	return puo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillableSlug(s *string) *PageUpdateOne {
	if s != nil {
		puo.SetSlug(*s)
	}
	return puo
}

// SetPageDescription sets the "page_description" field.
func (puo *PageUpdateOne) SetPageDescription(s string) *PageUpdateOne {
	puo.mutation.SetPageDescription(s)
	return puo
}

// SetNillablePageDescription sets the "page_description" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillablePageDescription(s *string) *PageUpdateOne {
	if s != nil {
		puo.SetPageDescription(*s)
	}
	return puo
}

// ClearPageDescription clears the value of the "page_description" field.
func (puo *PageUpdateOne) ClearPageDescription() *PageUpdateOne {
	puo.mutation.ClearPageDescription()
	return puo
}

// SetPageContentJSON sets the "page_content_json" field.
func (puo *PageUpdateOne) SetPageContentJSON(s string) *PageUpdateOne {
	puo.mutation.SetPageContentJSON(s)
	return puo
}

// SetNillablePageContentJSON sets the "page_content_json" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillablePageContentJSON(s *string) *PageUpdateOne {
	if s != nil {
		puo.SetPageContentJSON(*s)
	}
	return puo
}

// ClearPageContentJSON clears the value of the "page_content_json" field.
func (puo *PageUpdateOne) ClearPageContentJSON() *PageUpdateOne {
	puo.mutation.ClearPageContentJSON()
	return puo
}

// SetPageContentHTML sets the "page_content_html" field.
func (puo *PageUpdateOne) SetPageContentHTML(s string) *PageUpdateOne {
	puo.mutation.SetPageContentHTML(s)
	return puo
}

// SetNillablePageContentHTML sets the "page_content_html" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillablePageContentHTML(s *string) *PageUpdateOne {
	if s != nil {
		puo.SetPageContentHTML(*s)
	}
	return puo
}

// ClearPageContentHTML clears the value of the "page_content_html" field.
func (puo *PageUpdateOne) ClearPageContentHTML() *PageUpdateOne {
	puo.mutation.ClearPageContentHTML()
	return puo
}

// SetStatus sets the "status" field.
func (puo *PageUpdateOne) SetStatus(es enums.StatusState) *PageUpdateOne {
	puo.mutation.SetStatus(es)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillableStatus(es *enums.StatusState) *PageUpdateOne {
	if es != nil {
		puo.SetStatus(*es)
	}
	return puo
}

// SetMetaDescription sets the "meta_description" field.
func (puo *PageUpdateOne) SetMetaDescription(td *types.MetaDescription) *PageUpdateOne {
	puo.mutation.SetMetaDescription(td)
	return puo
}

// ClearMetaDescription clears the value of the "meta_description" field.
func (puo *PageUpdateOne) ClearMetaDescription() *PageUpdateOne {
	puo.mutation.ClearMetaDescription()
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PageUpdateOne) SetCreatedAt(t time.Time) *PageUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillableCreatedAt(t *time.Time) *PageUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetHubID sets the "hub" edge to the Hub entity by ID.
func (puo *PageUpdateOne) SetHubID(id mixin.ID) *PageUpdateOne {
	puo.mutation.SetHubID(id)
	return puo
}

// SetHub sets the "hub" edge to the Hub entity.
func (puo *PageUpdateOne) SetHub(h *Hub) *PageUpdateOne {
	return puo.SetHubID(h.ID)
}

// Mutation returns the PageMutation object of the builder.
func (puo *PageUpdateOne) Mutation() *PageMutation {
	return puo.mutation
}

// ClearHub clears the "hub" edge to the Hub entity.
func (puo *PageUpdateOne) ClearHub() *PageUpdateOne {
	puo.mutation.ClearHub()
	return puo
}

// Where appends a list predicates to the PageUpdate builder.
func (puo *PageUpdateOne) Where(ps ...predicate.Page) *PageUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PageUpdateOne) Select(field string, fields ...string) *PageUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Page entity.
func (puo *PageUpdateOne) Save(ctx context.Context) (*Page, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PageUpdateOne) SaveX(ctx context.Context) *Page {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PageUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PageUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PageUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := page.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Page.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Slug(); ok {
		if err := page.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Page.slug": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Status(); ok {
		if err := page.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Page.status": %w`, err)}
		}
	}
	if puo.mutation.HubCleared() && len(puo.mutation.HubIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Page.hub"`)
	}
	return nil
}

func (puo *PageUpdateOne) sqlSave(ctx context.Context) (_node *Page, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(page.Table, page.Columns, sqlgraph.NewFieldSpec(page.FieldID, field.TypeString))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Page.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, page.FieldID)
		for _, f := range fields {
			if !page.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != page.FieldID {
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
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(page.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Slug(); ok {
		_spec.SetField(page.FieldSlug, field.TypeString, value)
	}
	if value, ok := puo.mutation.PageDescription(); ok {
		_spec.SetField(page.FieldPageDescription, field.TypeString, value)
	}
	if puo.mutation.PageDescriptionCleared() {
		_spec.ClearField(page.FieldPageDescription, field.TypeString)
	}
	if value, ok := puo.mutation.PageContentJSON(); ok {
		_spec.SetField(page.FieldPageContentJSON, field.TypeString, value)
	}
	if puo.mutation.PageContentJSONCleared() {
		_spec.ClearField(page.FieldPageContentJSON, field.TypeString)
	}
	if value, ok := puo.mutation.PageContentHTML(); ok {
		_spec.SetField(page.FieldPageContentHTML, field.TypeString, value)
	}
	if puo.mutation.PageContentHTMLCleared() {
		_spec.ClearField(page.FieldPageContentHTML, field.TypeString)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(page.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.MetaDescription(); ok {
		_spec.SetField(page.FieldMetaDescription, field.TypeJSON, value)
	}
	if puo.mutation.MetaDescriptionCleared() {
		_spec.ClearField(page.FieldMetaDescription, field.TypeJSON)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(page.FieldCreatedAt, field.TypeTime, value)
	}
	if puo.mutation.HubCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   page.HubTable,
			Columns: []string{page.HubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hub.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.HubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   page.HubTable,
			Columns: []string{page.HubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hub.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Page{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{page.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
