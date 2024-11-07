// Code generated by ent, DO NOT EDIT.

package link

import (
	"fmt"

	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the link type in the database.
	Label = "link"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTarget holds the string denoting the target field in the database.
	FieldTarget = "target"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldLinkContent holds the string denoting the link_content field in the database.
	FieldLinkContent = "link_content"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeHub holds the string denoting the hub edge name in mutations.
	EdgeHub = "hub"
	// Table holds the table name of the link in the database.
	Table = "links"
	// HubTable is the table that holds the hub relation/edge.
	HubTable = "links"
	// HubInverseTable is the table name for the Hub entity.
	// It exists in this package in order to avoid circular dependency with the "enthub" package.
	HubInverseTable = "hubs"
	// HubColumn is the table column denoting the hub relation/edge.
	HubColumn = "link_fk"
)

// Columns holds all SQL columns for link fields.
var Columns = []string{
	FieldID,
	FieldTarget,
	FieldPath,
	FieldLinkContent,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "links"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"link_fk",
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
	// TargetValidator is a validator for the "target" field. It is called by the builders before save.
	TargetValidator func(string) error
	// PathValidator is a validator for the "path" field. It is called by the builders before save.
	PathValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() mixin.ID
)

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s enums.StatusState) error {
	switch s.String() {
	case "STATUS_INACTIVE", "STATUS_ACTIVE":
		return nil
	default:
		return fmt.Errorf("link: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Link queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTarget orders the results by the target field.
func ByTarget(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTarget, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
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
func newHubStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HubInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, HubTable, HubColumn),
	)
}
