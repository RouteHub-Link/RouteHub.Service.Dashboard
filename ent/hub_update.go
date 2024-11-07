// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	entdomain "RouteHub.Service.Dashboard/ent/domain"
	enthub "RouteHub.Service.Dashboard/ent/hub"
	"RouteHub.Service.Dashboard/ent/link"
	"RouteHub.Service.Dashboard/ent/organization"
	"RouteHub.Service.Dashboard/ent/predicate"
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/enums/hub"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HubUpdate is the builder for updating Hub entities.
type HubUpdate struct {
	config
	hooks    []Hook
	mutation *HubMutation
}

// Where appends a list predicates to the HubUpdate builder.
func (hu *HubUpdate) Where(ps ...predicate.Hub) *HubUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetName sets the "name" field.
func (hu *HubUpdate) SetName(s string) *HubUpdate {
	hu.mutation.SetName(s)
	return hu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (hu *HubUpdate) SetNillableName(s *string) *HubUpdate {
	if s != nil {
		hu.SetName(*s)
	}
	return hu
}

// SetSlug sets the "slug" field.
func (hu *HubUpdate) SetSlug(s string) *HubUpdate {
	hu.mutation.SetSlug(s)
	return hu
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (hu *HubUpdate) SetNillableSlug(s *string) *HubUpdate {
	if s != nil {
		hu.SetSlug(*s)
	}
	return hu
}

// SetHubDetails sets the "hub_details" field.
func (hu *HubUpdate) SetHubDetails(td types.HubDetails) *HubUpdate {
	hu.mutation.SetHubDetails(td)
	return hu
}

// SetNillableHubDetails sets the "hub_details" field if the given value is not nil.
func (hu *HubUpdate) SetNillableHubDetails(td *types.HubDetails) *HubUpdate {
	if td != nil {
		hu.SetHubDetails(*td)
	}
	return hu
}

// ClearHubDetails clears the value of the "hub_details" field.
func (hu *HubUpdate) ClearHubDetails() *HubUpdate {
	hu.mutation.ClearHubDetails()
	return hu
}

// SetTCPAddress sets the "tcp_address" field.
func (hu *HubUpdate) SetTCPAddress(s string) *HubUpdate {
	hu.mutation.SetTCPAddress(s)
	return hu
}

// SetNillableTCPAddress sets the "tcp_address" field if the given value is not nil.
func (hu *HubUpdate) SetNillableTCPAddress(s *string) *HubUpdate {
	if s != nil {
		hu.SetTCPAddress(*s)
	}
	return hu
}

// SetStatus sets the "status" field.
func (hu *HubUpdate) SetStatus(es enums.StatusState) *HubUpdate {
	hu.mutation.SetStatus(es)
	return hu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hu *HubUpdate) SetNillableStatus(es *enums.StatusState) *HubUpdate {
	if es != nil {
		hu.SetStatus(*es)
	}
	return hu
}

// SetDefaultRedirection sets the "default_redirection" field.
func (hu *HubUpdate) SetDefaultRedirection(ho hub.RedirectionOption) *HubUpdate {
	hu.mutation.SetDefaultRedirection(ho)
	return hu
}

// SetNillableDefaultRedirection sets the "default_redirection" field if the given value is not nil.
func (hu *HubUpdate) SetNillableDefaultRedirection(ho *hub.RedirectionOption) *HubUpdate {
	if ho != nil {
		hu.SetDefaultRedirection(*ho)
	}
	return hu
}

// SetDomainID sets the "domain" edge to the Domain entity by ID.
func (hu *HubUpdate) SetDomainID(id mixin.ID) *HubUpdate {
	hu.mutation.SetDomainID(id)
	return hu
}

// SetDomain sets the "domain" edge to the Domain entity.
func (hu *HubUpdate) SetDomain(d *Domain) *HubUpdate {
	return hu.SetDomainID(d.ID)
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (hu *HubUpdate) SetOrganizationID(id mixin.ID) *HubUpdate {
	hu.mutation.SetOrganizationID(id)
	return hu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (hu *HubUpdate) SetOrganization(o *Organization) *HubUpdate {
	return hu.SetOrganizationID(o.ID)
}

// AddLinkIDs adds the "links" edge to the Link entity by IDs.
func (hu *HubUpdate) AddLinkIDs(ids ...mixin.ID) *HubUpdate {
	hu.mutation.AddLinkIDs(ids...)
	return hu
}

// AddLinks adds the "links" edges to the Link entity.
func (hu *HubUpdate) AddLinks(l ...*Link) *HubUpdate {
	ids := make([]mixin.ID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return hu.AddLinkIDs(ids...)
}

// Mutation returns the HubMutation object of the builder.
func (hu *HubUpdate) Mutation() *HubMutation {
	return hu.mutation
}

// ClearDomain clears the "domain" edge to the Domain entity.
func (hu *HubUpdate) ClearDomain() *HubUpdate {
	hu.mutation.ClearDomain()
	return hu
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (hu *HubUpdate) ClearOrganization() *HubUpdate {
	hu.mutation.ClearOrganization()
	return hu
}

// ClearLinks clears all "links" edges to the Link entity.
func (hu *HubUpdate) ClearLinks() *HubUpdate {
	hu.mutation.ClearLinks()
	return hu
}

// RemoveLinkIDs removes the "links" edge to Link entities by IDs.
func (hu *HubUpdate) RemoveLinkIDs(ids ...mixin.ID) *HubUpdate {
	hu.mutation.RemoveLinkIDs(ids...)
	return hu
}

// RemoveLinks removes "links" edges to Link entities.
func (hu *HubUpdate) RemoveLinks(l ...*Link) *HubUpdate {
	ids := make([]mixin.ID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return hu.RemoveLinkIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HubUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, hu.sqlSave, hu.mutation, hu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HubUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HubUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HubUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hu *HubUpdate) check() error {
	if v, ok := hu.mutation.Name(); ok {
		if err := enthub.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Hub.name": %w`, err)}
		}
	}
	if v, ok := hu.mutation.Slug(); ok {
		if err := enthub.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Hub.slug": %w`, err)}
		}
	}
	if v, ok := hu.mutation.TCPAddress(); ok {
		if err := enthub.TCPAddressValidator(v); err != nil {
			return &ValidationError{Name: "tcp_address", err: fmt.Errorf(`ent: validator failed for field "Hub.tcp_address": %w`, err)}
		}
	}
	if v, ok := hu.mutation.Status(); ok {
		if err := enthub.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Hub.status": %w`, err)}
		}
	}
	if v, ok := hu.mutation.DefaultRedirection(); ok {
		if err := enthub.DefaultRedirectionValidator(v); err != nil {
			return &ValidationError{Name: "default_redirection", err: fmt.Errorf(`ent: validator failed for field "Hub.default_redirection": %w`, err)}
		}
	}
	if hu.mutation.DomainCleared() && len(hu.mutation.DomainIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Hub.domain"`)
	}
	if hu.mutation.OrganizationCleared() && len(hu.mutation.OrganizationIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Hub.organization"`)
	}
	return nil
}

func (hu *HubUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := hu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(enthub.Table, enthub.Columns, sqlgraph.NewFieldSpec(enthub.FieldID, field.TypeString))
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.Name(); ok {
		_spec.SetField(enthub.FieldName, field.TypeString, value)
	}
	if value, ok := hu.mutation.Slug(); ok {
		_spec.SetField(enthub.FieldSlug, field.TypeString, value)
	}
	if value, ok := hu.mutation.HubDetails(); ok {
		_spec.SetField(enthub.FieldHubDetails, field.TypeJSON, value)
	}
	if hu.mutation.HubDetailsCleared() {
		_spec.ClearField(enthub.FieldHubDetails, field.TypeJSON)
	}
	if value, ok := hu.mutation.TCPAddress(); ok {
		_spec.SetField(enthub.FieldTCPAddress, field.TypeString, value)
	}
	if value, ok := hu.mutation.Status(); ok {
		_spec.SetField(enthub.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := hu.mutation.DefaultRedirection(); ok {
		_spec.SetField(enthub.FieldDefaultRedirection, field.TypeEnum, value)
	}
	if hu.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   enthub.DomainTable,
			Columns: []string{enthub.DomainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entdomain.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.DomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   enthub.DomainTable,
			Columns: []string{enthub.DomainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entdomain.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   enthub.OrganizationTable,
			Columns: []string{enthub.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   enthub.OrganizationTable,
			Columns: []string{enthub.OrganizationColumn},
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
	if hu.mutation.LinksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   enthub.LinksTable,
			Columns: []string{enthub.LinksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(link.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedLinksIDs(); len(nodes) > 0 && !hu.mutation.LinksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   enthub.LinksTable,
			Columns: []string{enthub.LinksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(link.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.LinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   enthub.LinksTable,
			Columns: []string{enthub.LinksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(link.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{enthub.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hu.mutation.done = true
	return n, nil
}

// HubUpdateOne is the builder for updating a single Hub entity.
type HubUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HubMutation
}

// SetName sets the "name" field.
func (huo *HubUpdateOne) SetName(s string) *HubUpdateOne {
	huo.mutation.SetName(s)
	return huo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (huo *HubUpdateOne) SetNillableName(s *string) *HubUpdateOne {
	if s != nil {
		huo.SetName(*s)
	}
	return huo
}

// SetSlug sets the "slug" field.
func (huo *HubUpdateOne) SetSlug(s string) *HubUpdateOne {
	huo.mutation.SetSlug(s)
	return huo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (huo *HubUpdateOne) SetNillableSlug(s *string) *HubUpdateOne {
	if s != nil {
		huo.SetSlug(*s)
	}
	return huo
}

// SetHubDetails sets the "hub_details" field.
func (huo *HubUpdateOne) SetHubDetails(td types.HubDetails) *HubUpdateOne {
	huo.mutation.SetHubDetails(td)
	return huo
}

// SetNillableHubDetails sets the "hub_details" field if the given value is not nil.
func (huo *HubUpdateOne) SetNillableHubDetails(td *types.HubDetails) *HubUpdateOne {
	if td != nil {
		huo.SetHubDetails(*td)
	}
	return huo
}

// ClearHubDetails clears the value of the "hub_details" field.
func (huo *HubUpdateOne) ClearHubDetails() *HubUpdateOne {
	huo.mutation.ClearHubDetails()
	return huo
}

// SetTCPAddress sets the "tcp_address" field.
func (huo *HubUpdateOne) SetTCPAddress(s string) *HubUpdateOne {
	huo.mutation.SetTCPAddress(s)
	return huo
}

// SetNillableTCPAddress sets the "tcp_address" field if the given value is not nil.
func (huo *HubUpdateOne) SetNillableTCPAddress(s *string) *HubUpdateOne {
	if s != nil {
		huo.SetTCPAddress(*s)
	}
	return huo
}

// SetStatus sets the "status" field.
func (huo *HubUpdateOne) SetStatus(es enums.StatusState) *HubUpdateOne {
	huo.mutation.SetStatus(es)
	return huo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (huo *HubUpdateOne) SetNillableStatus(es *enums.StatusState) *HubUpdateOne {
	if es != nil {
		huo.SetStatus(*es)
	}
	return huo
}

// SetDefaultRedirection sets the "default_redirection" field.
func (huo *HubUpdateOne) SetDefaultRedirection(ho hub.RedirectionOption) *HubUpdateOne {
	huo.mutation.SetDefaultRedirection(ho)
	return huo
}

// SetNillableDefaultRedirection sets the "default_redirection" field if the given value is not nil.
func (huo *HubUpdateOne) SetNillableDefaultRedirection(ho *hub.RedirectionOption) *HubUpdateOne {
	if ho != nil {
		huo.SetDefaultRedirection(*ho)
	}
	return huo
}

// SetDomainID sets the "domain" edge to the Domain entity by ID.
func (huo *HubUpdateOne) SetDomainID(id mixin.ID) *HubUpdateOne {
	huo.mutation.SetDomainID(id)
	return huo
}

// SetDomain sets the "domain" edge to the Domain entity.
func (huo *HubUpdateOne) SetDomain(d *Domain) *HubUpdateOne {
	return huo.SetDomainID(d.ID)
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (huo *HubUpdateOne) SetOrganizationID(id mixin.ID) *HubUpdateOne {
	huo.mutation.SetOrganizationID(id)
	return huo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (huo *HubUpdateOne) SetOrganization(o *Organization) *HubUpdateOne {
	return huo.SetOrganizationID(o.ID)
}

// AddLinkIDs adds the "links" edge to the Link entity by IDs.
func (huo *HubUpdateOne) AddLinkIDs(ids ...mixin.ID) *HubUpdateOne {
	huo.mutation.AddLinkIDs(ids...)
	return huo
}

// AddLinks adds the "links" edges to the Link entity.
func (huo *HubUpdateOne) AddLinks(l ...*Link) *HubUpdateOne {
	ids := make([]mixin.ID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return huo.AddLinkIDs(ids...)
}

// Mutation returns the HubMutation object of the builder.
func (huo *HubUpdateOne) Mutation() *HubMutation {
	return huo.mutation
}

// ClearDomain clears the "domain" edge to the Domain entity.
func (huo *HubUpdateOne) ClearDomain() *HubUpdateOne {
	huo.mutation.ClearDomain()
	return huo
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (huo *HubUpdateOne) ClearOrganization() *HubUpdateOne {
	huo.mutation.ClearOrganization()
	return huo
}

// ClearLinks clears all "links" edges to the Link entity.
func (huo *HubUpdateOne) ClearLinks() *HubUpdateOne {
	huo.mutation.ClearLinks()
	return huo
}

// RemoveLinkIDs removes the "links" edge to Link entities by IDs.
func (huo *HubUpdateOne) RemoveLinkIDs(ids ...mixin.ID) *HubUpdateOne {
	huo.mutation.RemoveLinkIDs(ids...)
	return huo
}

// RemoveLinks removes "links" edges to Link entities.
func (huo *HubUpdateOne) RemoveLinks(l ...*Link) *HubUpdateOne {
	ids := make([]mixin.ID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return huo.RemoveLinkIDs(ids...)
}

// Where appends a list predicates to the HubUpdate builder.
func (huo *HubUpdateOne) Where(ps ...predicate.Hub) *HubUpdateOne {
	huo.mutation.Where(ps...)
	return huo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HubUpdateOne) Select(field string, fields ...string) *HubUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated Hub entity.
func (huo *HubUpdateOne) Save(ctx context.Context) (*Hub, error) {
	return withHooks(ctx, huo.sqlSave, huo.mutation, huo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HubUpdateOne) SaveX(ctx context.Context) *Hub {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HubUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HubUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (huo *HubUpdateOne) check() error {
	if v, ok := huo.mutation.Name(); ok {
		if err := enthub.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Hub.name": %w`, err)}
		}
	}
	if v, ok := huo.mutation.Slug(); ok {
		if err := enthub.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Hub.slug": %w`, err)}
		}
	}
	if v, ok := huo.mutation.TCPAddress(); ok {
		if err := enthub.TCPAddressValidator(v); err != nil {
			return &ValidationError{Name: "tcp_address", err: fmt.Errorf(`ent: validator failed for field "Hub.tcp_address": %w`, err)}
		}
	}
	if v, ok := huo.mutation.Status(); ok {
		if err := enthub.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Hub.status": %w`, err)}
		}
	}
	if v, ok := huo.mutation.DefaultRedirection(); ok {
		if err := enthub.DefaultRedirectionValidator(v); err != nil {
			return &ValidationError{Name: "default_redirection", err: fmt.Errorf(`ent: validator failed for field "Hub.default_redirection": %w`, err)}
		}
	}
	if huo.mutation.DomainCleared() && len(huo.mutation.DomainIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Hub.domain"`)
	}
	if huo.mutation.OrganizationCleared() && len(huo.mutation.OrganizationIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Hub.organization"`)
	}
	return nil
}

func (huo *HubUpdateOne) sqlSave(ctx context.Context) (_node *Hub, err error) {
	if err := huo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(enthub.Table, enthub.Columns, sqlgraph.NewFieldSpec(enthub.FieldID, field.TypeString))
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Hub.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, enthub.FieldID)
		for _, f := range fields {
			if !enthub.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != enthub.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.Name(); ok {
		_spec.SetField(enthub.FieldName, field.TypeString, value)
	}
	if value, ok := huo.mutation.Slug(); ok {
		_spec.SetField(enthub.FieldSlug, field.TypeString, value)
	}
	if value, ok := huo.mutation.HubDetails(); ok {
		_spec.SetField(enthub.FieldHubDetails, field.TypeJSON, value)
	}
	if huo.mutation.HubDetailsCleared() {
		_spec.ClearField(enthub.FieldHubDetails, field.TypeJSON)
	}
	if value, ok := huo.mutation.TCPAddress(); ok {
		_spec.SetField(enthub.FieldTCPAddress, field.TypeString, value)
	}
	if value, ok := huo.mutation.Status(); ok {
		_spec.SetField(enthub.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := huo.mutation.DefaultRedirection(); ok {
		_spec.SetField(enthub.FieldDefaultRedirection, field.TypeEnum, value)
	}
	if huo.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   enthub.DomainTable,
			Columns: []string{enthub.DomainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entdomain.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.DomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   enthub.DomainTable,
			Columns: []string{enthub.DomainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entdomain.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   enthub.OrganizationTable,
			Columns: []string{enthub.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   enthub.OrganizationTable,
			Columns: []string{enthub.OrganizationColumn},
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
	if huo.mutation.LinksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   enthub.LinksTable,
			Columns: []string{enthub.LinksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(link.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedLinksIDs(); len(nodes) > 0 && !huo.mutation.LinksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   enthub.LinksTable,
			Columns: []string{enthub.LinksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(link.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.LinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   enthub.LinksTable,
			Columns: []string{enthub.LinksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(link.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Hub{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{enthub.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	huo.mutation.done = true
	return _node, nil
}
