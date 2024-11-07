// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"RouteHub.Service.Dashboard/ent/organization"
	"RouteHub.Service.Dashboard/ent/person"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"go.jetify.com/typeid"
)

// Person is the model entity for the Person schema.
type Person struct {
	config `json:"-"`
	// ID of the ent.
	ID typeid.AnyID `json:"id,omitempty"`
	// SubjectID holds the value of the "subject_id" field.
	SubjectID string `json:"subject_id,omitempty"`
	// UserInfo holds the value of the "user_info" field.
	UserInfo oidc.UserInfo `json:"user_info,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonQuery when eager-loading is set.
	Edges           PersonEdges `json:"edges"`
	organization_id *typeid.AnyID
	organization_fk *typeid.AnyID
	selectValues    sql.SelectValues
}

// PersonEdges holds the relations/edges for other nodes in the graph.
type PersonEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Person) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case person.FieldUserInfo:
			values[i] = new([]byte)
		case person.FieldIsActive:
			values[i] = new(sql.NullBool)
		case person.FieldSubjectID:
			values[i] = new(sql.NullString)
		case person.FieldID:
			values[i] = new(typeid.AnyID)
		case person.ForeignKeys[0]: // organization_id
			values[i] = &sql.NullScanner{S: new(typeid.AnyID)}
		case person.ForeignKeys[1]: // organization_fk
			values[i] = &sql.NullScanner{S: new(typeid.AnyID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Person fields.
func (pe *Person) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case person.FieldID:
			if value, ok := values[i].(*typeid.AnyID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pe.ID = *value
			}
		case person.FieldSubjectID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject_id", values[i])
			} else if value.Valid {
				pe.SubjectID = value.String
			}
		case person.FieldUserInfo:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field user_info", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.UserInfo); err != nil {
					return fmt.Errorf("unmarshal field user_info: %w", err)
				}
			}
		case person.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				pe.IsActive = value.Bool
			}
		case person.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				pe.organization_id = new(typeid.AnyID)
				*pe.organization_id = *value.S.(*typeid.AnyID)
			}
		case person.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field organization_fk", values[i])
			} else if value.Valid {
				pe.organization_fk = new(typeid.AnyID)
				*pe.organization_fk = *value.S.(*typeid.AnyID)
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Person.
// This includes values selected through modifiers, order, etc.
func (pe *Person) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// QueryOrganization queries the "organization" edge of the Person entity.
func (pe *Person) QueryOrganization() *OrganizationQuery {
	return NewPersonClient(pe.config).QueryOrganization(pe)
}

// Update returns a builder for updating this Person.
// Note that you need to call Person.Unwrap() before calling this method if this Person
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Person) Update() *PersonUpdateOne {
	return NewPersonClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Person entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Person) Unwrap() *Person {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Person is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Person) String() string {
	var builder strings.Builder
	builder.WriteString("Person(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("subject_id=")
	builder.WriteString(pe.SubjectID)
	builder.WriteString(", ")
	builder.WriteString("user_info=")
	builder.WriteString(fmt.Sprintf("%v", pe.UserInfo))
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", pe.IsActive))
	builder.WriteByte(')')
	return builder.String()
}

// Persons is a parsable slice of Person.
type Persons []*Person