// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"RouteHub.Service.Dashboard/ent/organization"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Organization is the model entity for the Organization schema.
type Organization struct {
	config `json:"-"`
	// ID of the ent.
	ID mixin.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Website holds the value of the "website" field.
	Website string `json:"website,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// SocialMedias holds the value of the "social_medias" field.
	SocialMedias types.SocialMedias `json:"social_medias,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrganizationQuery when eager-loading is set.
	Edges        OrganizationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrganizationEdges holds the relations/edges for other nodes in the graph.
type OrganizationEdges struct {
	// Domains holds the value of the domains edge.
	Domains []*Domain `json:"domains,omitempty"`
	// Hubs holds the value of the hubs edge.
	Hubs []*Hub `json:"hubs,omitempty"`
	// Persons holds the value of the persons edge.
	Persons []*Person `json:"persons,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DomainsOrErr returns the Domains value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) DomainsOrErr() ([]*Domain, error) {
	if e.loadedTypes[0] {
		return e.Domains, nil
	}
	return nil, &NotLoadedError{edge: "domains"}
}

// HubsOrErr returns the Hubs value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) HubsOrErr() ([]*Hub, error) {
	if e.loadedTypes[1] {
		return e.Hubs, nil
	}
	return nil, &NotLoadedError{edge: "hubs"}
}

// PersonsOrErr returns the Persons value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) PersonsOrErr() ([]*Person, error) {
	if e.loadedTypes[2] {
		return e.Persons, nil
	}
	return nil, &NotLoadedError{edge: "persons"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Organization) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organization.FieldSocialMedias:
			values[i] = new([]byte)
		case organization.FieldID, organization.FieldName, organization.FieldWebsite, organization.FieldDescription, organization.FieldLocation:
			values[i] = new(sql.NullString)
		case organization.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Organization fields.
func (o *Organization) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organization.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				o.ID = mixin.ID(value.String)
			}
		case organization.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				o.Name = value.String
			}
		case organization.FieldWebsite:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website", values[i])
			} else if value.Valid {
				o.Website = value.String
			}
		case organization.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				o.Description = value.String
			}
		case organization.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				o.Location = value.String
			}
		case organization.FieldSocialMedias:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field social_medias", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.SocialMedias); err != nil {
					return fmt.Errorf("unmarshal field social_medias: %w", err)
				}
			}
		case organization.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Organization.
// This includes values selected through modifiers, order, etc.
func (o *Organization) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryDomains queries the "domains" edge of the Organization entity.
func (o *Organization) QueryDomains() *DomainQuery {
	return NewOrganizationClient(o.config).QueryDomains(o)
}

// QueryHubs queries the "hubs" edge of the Organization entity.
func (o *Organization) QueryHubs() *HubQuery {
	return NewOrganizationClient(o.config).QueryHubs(o)
}

// QueryPersons queries the "persons" edge of the Organization entity.
func (o *Organization) QueryPersons() *PersonQuery {
	return NewOrganizationClient(o.config).QueryPersons(o)
}

// Update returns a builder for updating this Organization.
// Note that you need to call Organization.Unwrap() before calling this method if this Organization
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Organization) Update() *OrganizationUpdateOne {
	return NewOrganizationClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Organization entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Organization) Unwrap() *Organization {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Organization is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Organization) String() string {
	var builder strings.Builder
	builder.WriteString("Organization(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("name=")
	builder.WriteString(o.Name)
	builder.WriteString(", ")
	builder.WriteString("website=")
	builder.WriteString(o.Website)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(o.Description)
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(o.Location)
	builder.WriteString(", ")
	builder.WriteString("social_medias=")
	builder.WriteString(fmt.Sprintf("%v", o.SocialMedias))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Organizations is a parsable slice of Organization.
type Organizations []*Organization
