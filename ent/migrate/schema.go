// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DomainsColumns holds the columns for the "domains" table.
	DomainsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "url", Type: field.TypeString, Unique: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PASSIVE", "DNS_STATUS_PENDING", "DNS_STATUS_VERIFIED", "DNS_STATUS_FAILED", "ACTIVE"}, Default: "PASSIVE"},
		{Name: "domain_fk", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "organization_id", Type: field.TypeUUID},
	}
	// DomainsTable holds the schema information for the "domains" table.
	DomainsTable = &schema.Table{
		Name:       "domains",
		Columns:    DomainsColumns,
		PrimaryKey: []*schema.Column{DomainsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "domains_hubs_domain",
				Columns:    []*schema.Column{DomainsColumns[4]},
				RefColumns: []*schema.Column{HubsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "domains_organizations_domains",
				Columns:    []*schema.Column{DomainsColumns[5]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// HubsColumns holds the columns for the "hubs" table.
	HubsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString},
		{Name: "hub_details", Type: field.TypeJSON, Nullable: true},
		{Name: "tcp_address", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"STATUS_INACTIVE", "STATUS_ACTIVE"}, Default: "STATUS_INACTIVE"},
		{Name: "default_redirection", Type: field.TypeEnum, Enums: []string{"TIMED", "NOT_AUTO_REDIRECT", "DIRECT_HTTP_REDIRECT", "CONFIRM_REDIRECT", "CUSTOM"}, Default: "TIMED"},
		{Name: "organization_id", Type: field.TypeUUID},
	}
	// HubsTable holds the schema information for the "hubs" table.
	HubsTable = &schema.Table{
		Name:       "hubs",
		Columns:    HubsColumns,
		PrimaryKey: []*schema.Column{HubsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hubs_organizations_hubs",
				Columns:    []*schema.Column{HubsColumns[7]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// LinksColumns holds the columns for the "links" table.
	LinksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "target", Type: field.TypeString},
		{Name: "path", Type: field.TypeString, Unique: true},
		{Name: "link_content", Type: field.TypeJSON, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"STATUS_INACTIVE", "STATUS_ACTIVE"}},
		{Name: "link_fk", Type: field.TypeUUID},
	}
	// LinksTable holds the schema information for the "links" table.
	LinksTable = &schema.Table{
		Name:       "links",
		Columns:    LinksColumns,
		PrimaryKey: []*schema.Column{LinksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "links_hubs_links",
				Columns:    []*schema.Column{LinksColumns[5]},
				RefColumns: []*schema.Column{HubsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// OrganizationsColumns holds the columns for the "organizations" table.
	OrganizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "website", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "locagtion", Type: field.TypeString, Nullable: true},
		{Name: "social_medias", Type: field.TypeJSON, Nullable: true},
	}
	// OrganizationsTable holds the schema information for the "organizations" table.
	OrganizationsTable = &schema.Table{
		Name:       "organizations",
		Columns:    OrganizationsColumns,
		PrimaryKey: []*schema.Column{OrganizationsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DomainsTable,
		HubsTable,
		LinksTable,
		OrganizationsTable,
	}
)

func init() {
	DomainsTable.ForeignKeys[0].RefTable = HubsTable
	DomainsTable.ForeignKeys[1].RefTable = OrganizationsTable
	HubsTable.ForeignKeys[0].RefTable = OrganizationsTable
	LinksTable.ForeignKeys[0].RefTable = HubsTable
}
