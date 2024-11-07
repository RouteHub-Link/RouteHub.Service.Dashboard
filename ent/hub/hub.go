// Code generated by ent, DO NOT EDIT.

package enthub

import (
	"fmt"

	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/enums/hub"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.jetify.com/typeid"
)

const (
	// Label holds the string label denoting the hub type in the database.
	Label = "hub"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldHubDetails holds the string denoting the hub_details field in the database.
	FieldHubDetails = "hub_details"
	// FieldTCPAddress holds the string denoting the tcp_address field in the database.
	FieldTCPAddress = "tcp_address"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDefaultRedirection holds the string denoting the default_redirection field in the database.
	FieldDefaultRedirection = "default_redirection"
	// EdgeDomain holds the string denoting the domain edge name in mutations.
	EdgeDomain = "domain"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeLinks holds the string denoting the links edge name in mutations.
	EdgeLinks = "links"
	// Table holds the table name of the hub in the database.
	Table = "hubs"
	// DomainTable is the table that holds the domain relation/edge.
	DomainTable = "domains"
	// DomainInverseTable is the table name for the Domain entity.
	// It exists in this package in order to avoid circular dependency with the "entdomain" package.
	DomainInverseTable = "domains"
	// DomainColumn is the table column denoting the domain relation/edge.
	DomainColumn = "domain_fk"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "hubs"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// LinksTable is the table that holds the links relation/edge.
	LinksTable = "links"
	// LinksInverseTable is the table name for the Link entity.
	// It exists in this package in order to avoid circular dependency with the "link" package.
	LinksInverseTable = "links"
	// LinksColumn is the table column denoting the links relation/edge.
	LinksColumn = "link_fk"
)

// Columns holds all SQL columns for hub fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldSlug,
	FieldHubDetails,
	FieldTCPAddress,
	FieldStatus,
	FieldDefaultRedirection,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "hubs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
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
	// SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	SlugValidator func(string) error
	// TCPAddressValidator is a validator for the "tcp_address" field. It is called by the builders before save.
	TCPAddressValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() typeid.AnyID
)

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s enums.StatusState) error {
	switch s.String() {
	case "STATUS_INACTIVE", "STATUS_ACTIVE":
		return nil
	default:
		return fmt.Errorf("enthub: invalid enum value for status field: %q", s)
	}
}

// DefaultRedirectionValidator is a validator for the "default_redirection" field enum values. It is called by the builders before save.
func DefaultRedirectionValidator(dr hub.RedirectionOption) error {
	switch dr.String() {
	case "TIMED", "NOT_AUTO_REDIRECT", "DIRECT_HTTP_REDIRECT", "CONFIRM_REDIRECT", "CUSTOM":
		return nil
	default:
		return fmt.Errorf("enthub: invalid enum value for default_redirection field: %q", dr)
	}
}

// OrderOption defines the ordering options for the Hub queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySlug orders the results by the slug field.
func BySlug(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSlug, opts...).ToFunc()
}

// ByTCPAddress orders the results by the tcp_address field.
func ByTCPAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTCPAddress, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByDefaultRedirection orders the results by the default_redirection field.
func ByDefaultRedirection(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefaultRedirection, opts...).ToFunc()
}

// ByDomainField orders the results by domain field.
func ByDomainField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDomainStep(), sql.OrderByField(field, opts...))
	}
}

// ByOrganizationField orders the results by organization field.
func ByOrganizationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationStep(), sql.OrderByField(field, opts...))
	}
}

// ByLinksCount orders the results by links count.
func ByLinksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLinksStep(), opts...)
	}
}

// ByLinks orders the results by links terms.
func ByLinks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLinksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDomainStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DomainInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, DomainTable, DomainColumn),
	)
}
func newOrganizationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTable, OrganizationColumn),
	)
}
func newLinksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LinksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LinksTable, LinksColumn),
	)
}