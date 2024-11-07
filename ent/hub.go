// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	entdomain "RouteHub.Service.Dashboard/ent/domain"
	enthub "RouteHub.Service.Dashboard/ent/hub"
	"RouteHub.Service.Dashboard/ent/organization"
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/enums/hub"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"go.jetify.com/typeid"
)

// Hub is the model entity for the Hub schema.
type Hub struct {
	config `json:"-"`
	// ID of the ent.
	ID typeid.AnyID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// HubDetails holds the value of the "hub_details" field.
	HubDetails types.HubDetails `json:"hub_details,omitempty"`
	// TCPAddress holds the value of the "tcp_address" field.
	TCPAddress string `json:"tcp_address,omitempty"`
	// Status holds the value of the "status" field.
	Status enums.StatusState `json:"status,omitempty"`
	// DefaultRedirection holds the value of the "default_redirection" field.
	DefaultRedirection hub.RedirectionOption `json:"default_redirection,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HubQuery when eager-loading is set.
	Edges           HubEdges `json:"edges"`
	organization_id *typeid.AnyID
	selectValues    sql.SelectValues
}

// HubEdges holds the relations/edges for other nodes in the graph.
type HubEdges struct {
	// Domain holds the value of the domain edge.
	Domain *Domain `json:"domain,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// Links holds the value of the links edge.
	Links []*Link `json:"links,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DomainOrErr returns the Domain value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HubEdges) DomainOrErr() (*Domain, error) {
	if e.Domain != nil {
		return e.Domain, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: entdomain.Label}
	}
	return nil, &NotLoadedError{edge: "domain"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HubEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// LinksOrErr returns the Links value or an error if the edge
// was not loaded in eager-loading.
func (e HubEdges) LinksOrErr() ([]*Link, error) {
	if e.loadedTypes[2] {
		return e.Links, nil
	}
	return nil, &NotLoadedError{edge: "links"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Hub) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case enthub.FieldHubDetails:
			values[i] = new([]byte)
		case enthub.FieldStatus:
			values[i] = new(enums.StatusState)
		case enthub.FieldDefaultRedirection:
			values[i] = new(hub.RedirectionOption)
		case enthub.FieldName, enthub.FieldSlug, enthub.FieldTCPAddress:
			values[i] = new(sql.NullString)
		case enthub.FieldID:
			values[i] = new(typeid.AnyID)
		case enthub.ForeignKeys[0]: // organization_id
			values[i] = &sql.NullScanner{S: new(typeid.AnyID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Hub fields.
func (h *Hub) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case enthub.FieldID:
			if value, ok := values[i].(*typeid.AnyID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				h.ID = *value
			}
		case enthub.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				h.Name = value.String
			}
		case enthub.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				h.Slug = value.String
			}
		case enthub.FieldHubDetails:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field hub_details", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &h.HubDetails); err != nil {
					return fmt.Errorf("unmarshal field hub_details: %w", err)
				}
			}
		case enthub.FieldTCPAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tcp_address", values[i])
			} else if value.Valid {
				h.TCPAddress = value.String
			}
		case enthub.FieldStatus:
			if value, ok := values[i].(*enums.StatusState); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil {
				h.Status = *value
			}
		case enthub.FieldDefaultRedirection:
			if value, ok := values[i].(*hub.RedirectionOption); !ok {
				return fmt.Errorf("unexpected type %T for field default_redirection", values[i])
			} else if value != nil {
				h.DefaultRedirection = *value
			}
		case enthub.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				h.organization_id = new(typeid.AnyID)
				*h.organization_id = *value.S.(*typeid.AnyID)
			}
		default:
			h.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Hub.
// This includes values selected through modifiers, order, etc.
func (h *Hub) Value(name string) (ent.Value, error) {
	return h.selectValues.Get(name)
}

// QueryDomain queries the "domain" edge of the Hub entity.
func (h *Hub) QueryDomain() *DomainQuery {
	return NewHubClient(h.config).QueryDomain(h)
}

// QueryOrganization queries the "organization" edge of the Hub entity.
func (h *Hub) QueryOrganization() *OrganizationQuery {
	return NewHubClient(h.config).QueryOrganization(h)
}

// QueryLinks queries the "links" edge of the Hub entity.
func (h *Hub) QueryLinks() *LinkQuery {
	return NewHubClient(h.config).QueryLinks(h)
}

// Update returns a builder for updating this Hub.
// Note that you need to call Hub.Unwrap() before calling this method if this Hub
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Hub) Update() *HubUpdateOne {
	return NewHubClient(h.config).UpdateOne(h)
}

// Unwrap unwraps the Hub entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Hub) Unwrap() *Hub {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Hub is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Hub) String() string {
	var builder strings.Builder
	builder.WriteString("Hub(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("name=")
	builder.WriteString(h.Name)
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(h.Slug)
	builder.WriteString(", ")
	builder.WriteString("hub_details=")
	builder.WriteString(fmt.Sprintf("%v", h.HubDetails))
	builder.WriteString(", ")
	builder.WriteString("tcp_address=")
	builder.WriteString(h.TCPAddress)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", h.Status))
	builder.WriteString(", ")
	builder.WriteString("default_redirection=")
	builder.WriteString(fmt.Sprintf("%v", h.DefaultRedirection))
	builder.WriteByte(')')
	return builder.String()
}

// Hubs is a parsable slice of Hub.
type Hubs []*Hub