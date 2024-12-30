// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DomainsColumns holds the columns for the "domains" table.
	DomainsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "url", Type: field.TypeString, Unique: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PASSIVE", "DNS_STATUS_PENDING", "DNS_STATUS_VERIFIED", "DNS_STATUS_FAILED", "ACTIVE"}, Default: "PASSIVE"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "organization_id", Type: field.TypeString},
	}
	// DomainsTable holds the schema information for the "domains" table.
	DomainsTable = &schema.Table{
		Name:       "domains",
		Columns:    DomainsColumns,
		PrimaryKey: []*schema.Column{DomainsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
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
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString},
		{Name: "hub_details", Type: field.TypeJSON, Nullable: true},
		{Name: "tcp_address", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"STATUS_INACTIVE", "STATUS_ACTIVE"}, Default: "STATUS_INACTIVE"},
		{Name: "default_redirection", Type: field.TypeEnum, Enums: []string{"TIMED", "NOT_AUTO", "DIRECT_HTTP", "CONFIRM", "CUSTOM"}, Default: "TIMED"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "domain_fk", Type: field.TypeString},
		{Name: "organization_id", Type: field.TypeString},
	}
	// HubsTable holds the schema information for the "hubs" table.
	HubsTable = &schema.Table{
		Name:       "hubs",
		Columns:    HubsColumns,
		PrimaryKey: []*schema.Column{HubsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hubs_domains_domain",
				Columns:    []*schema.Column{HubsColumns[8]},
				RefColumns: []*schema.Column{DomainsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "hubs_organizations_hubs",
				Columns:    []*schema.Column{HubsColumns[9]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// LinksColumns holds the columns for the "links" table.
	LinksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "target", Type: field.TypeString},
		{Name: "path", Type: field.TypeString, Unique: true},
		{Name: "link_content", Type: field.TypeJSON, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"STATUS_INACTIVE", "STATUS_ACTIVE"}},
		{Name: "redirection_choice", Type: field.TypeEnum, Enums: []string{"TIMED", "NOT_AUTO", "DIRECT_HTTP", "CONFIRM", "CUSTOM"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "link_fk", Type: field.TypeString},
	}
	// LinksTable holds the schema information for the "links" table.
	LinksTable = &schema.Table{
		Name:       "links",
		Columns:    LinksColumns,
		PrimaryKey: []*schema.Column{LinksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "links_hubs_links",
				Columns:    []*schema.Column{LinksColumns[7]},
				RefColumns: []*schema.Column{HubsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// OrganizationsColumns holds the columns for the "organizations" table.
	OrganizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "website", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "location", Type: field.TypeString, Nullable: true},
		{Name: "social_medias", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// OrganizationsTable holds the schema information for the "organizations" table.
	OrganizationsTable = &schema.Table{
		Name:       "organizations",
		Columns:    OrganizationsColumns,
		PrimaryKey: []*schema.Column{OrganizationsColumns[0]},
	}
	// PagesColumns holds the columns for the "pages" table.
	PagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString},
		{Name: "page_description", Type: field.TypeString, Nullable: true},
		{Name: "page_content_json", Type: field.TypeString, Nullable: true},
		{Name: "page_content_html", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"STATUS_INACTIVE", "STATUS_ACTIVE"}, Default: "STATUS_INACTIVE"},
		{Name: "meta_description", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "hub_fk", Type: field.TypeString},
	}
	// PagesTable holds the schema information for the "pages" table.
	PagesTable = &schema.Table{
		Name:       "pages",
		Columns:    PagesColumns,
		PrimaryKey: []*schema.Column{PagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pages_hubs_pages",
				Columns:    []*schema.Column{PagesColumns[9]},
				RefColumns: []*schema.Column{HubsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PersonsColumns holds the columns for the "persons" table.
	PersonsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "subject_id", Type: field.TypeString, Unique: true},
		{Name: "user_info", Type: field.TypeJSON},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// PersonsTable holds the schema information for the "persons" table.
	PersonsTable = &schema.Table{
		Name:       "persons",
		Columns:    PersonsColumns,
		PrimaryKey: []*schema.Column{PersonsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "person_subject_id",
				Unique:  true,
				Columns: []*schema.Column{PersonsColumns[1]},
			},
		},
	}
	// OrganizationPersonsColumns holds the columns for the "organization_persons" table.
	OrganizationPersonsColumns = []*schema.Column{
		{Name: "organization_id", Type: field.TypeString},
		{Name: "person_id", Type: field.TypeString},
	}
	// OrganizationPersonsTable holds the schema information for the "organization_persons" table.
	OrganizationPersonsTable = &schema.Table{
		Name:       "organization_persons",
		Columns:    OrganizationPersonsColumns,
		PrimaryKey: []*schema.Column{OrganizationPersonsColumns[0], OrganizationPersonsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_persons_organization_id",
				Columns:    []*schema.Column{OrganizationPersonsColumns[0]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "organization_persons_person_id",
				Columns:    []*schema.Column{OrganizationPersonsColumns[1]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DomainsTable,
		HubsTable,
		LinksTable,
		OrganizationsTable,
		PagesTable,
		PersonsTable,
		OrganizationPersonsTable,
	}
)

func init() {
	DomainsTable.ForeignKeys[0].RefTable = OrganizationsTable
	HubsTable.ForeignKeys[0].RefTable = DomainsTable
	HubsTable.ForeignKeys[1].RefTable = OrganizationsTable
	LinksTable.ForeignKeys[0].RefTable = HubsTable
	PagesTable.ForeignKeys[0].RefTable = HubsTable
	OrganizationPersonsTable.ForeignKeys[0].RefTable = OrganizationsTable
	OrganizationPersonsTable.ForeignKeys[1].RefTable = PersonsTable
}
