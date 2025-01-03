// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	entdomain "RouteHub.Service.Dashboard/ent/domain"
	"RouteHub.Service.Dashboard/ent/organization"
	"RouteHub.Service.Dashboard/ent/schema/enums/domain"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Domain is the model entity for the Domain schema.
type Domain struct {
	config `json:"-"`
	// ID of the ent.
	ID mixin.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Status holds the value of the "status" field.
	Status domain.DomainState `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DomainQuery when eager-loading is set.
	Edges           DomainEdges `json:"edges"`
	organization_id *mixin.ID
	selectValues    sql.SelectValues
}

// DomainEdges holds the relations/edges for other nodes in the graph.
type DomainEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DomainEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Domain) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case entdomain.FieldStatus:
			values[i] = new(domain.DomainState)
		case entdomain.FieldID, entdomain.FieldName, entdomain.FieldURL:
			values[i] = new(sql.NullString)
		case entdomain.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case entdomain.ForeignKeys[0]: // organization_id
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Domain fields.
func (d *Domain) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entdomain.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				d.ID = mixin.ID(value.String)
			}
		case entdomain.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case entdomain.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				d.URL = value.String
			}
		case entdomain.FieldStatus:
			if value, ok := values[i].(*domain.DomainState); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil {
				d.Status = *value
			}
		case entdomain.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case entdomain.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				d.organization_id = new(mixin.ID)
				*d.organization_id = mixin.ID(value.String)
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Domain.
// This includes values selected through modifiers, order, etc.
func (d *Domain) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryOrganization queries the "organization" edge of the Domain entity.
func (d *Domain) QueryOrganization() *OrganizationQuery {
	return NewDomainClient(d.config).QueryOrganization(d)
}

// Update returns a builder for updating this Domain.
// Note that you need to call Domain.Unwrap() before calling this method if this Domain
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Domain) Update() *DomainUpdateOne {
	return NewDomainClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Domain entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Domain) Unwrap() *Domain {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Domain is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Domain) String() string {
	var builder strings.Builder
	builder.WriteString("Domain(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(d.URL)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", d.Status))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Domains is a parsable slice of Domain.
type Domains []*Domain
