// Code generated by ent, DO NOT EDIT.

package entdomain

import (
	"fmt"

	"RouteHub.Service.Dashboard/ent/schema/enums/domain"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.jetify.com/typeid"
)

const (
	// Label holds the string label denoting the domain type in the database.
	Label = "domain"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeHub holds the string denoting the hub edge name in mutations.
	EdgeHub = "hub"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// Table holds the table name of the domain in the database.
	Table = "domains"
	// HubTable is the table that holds the hub relation/edge.
	HubTable = "domains"
	// HubInverseTable is the table name for the Hub entity.
	// It exists in this package in order to avoid circular dependency with the "enthub" package.
	HubInverseTable = "hubs"
	// HubColumn is the table column denoting the hub relation/edge.
	HubColumn = "domain_fk"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "domains"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
)

// Columns holds all SQL columns for domain fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldURL,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "domains"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"domain_fk",
	"organization_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() typeid.AnyID
)

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s domain.DomainState) error {
	switch s.String() {
	case "PASSIVE", "DNS_STATUS_PENDING", "DNS_STATUS_VERIFIED", "DNS_STATUS_FAILED", "ACTIVE":
		return nil
	default:
		return fmt.Errorf("entdomain: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Domain queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByHubField orders the results by hub field.
func ByHubField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHubStep(), sql.OrderByField(field, opts...))
	}
}

// ByOrganizationField orders the results by organization field.
func ByOrganizationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationStep(), sql.OrderByField(field, opts...))
	}
}
func newHubStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HubInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, HubTable, HubColumn),
	)
}
func newOrganizationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTable, OrganizationColumn),
	)
}
