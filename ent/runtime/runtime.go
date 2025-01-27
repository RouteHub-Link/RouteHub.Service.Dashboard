// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	entdomain "RouteHub.Service.Dashboard/ent/domain"
	"RouteHub.Service.Dashboard/ent/hub"
	"RouteHub.Service.Dashboard/ent/link"
	"RouteHub.Service.Dashboard/ent/organization"
	"RouteHub.Service.Dashboard/ent/page"
	"RouteHub.Service.Dashboard/ent/person"
	"RouteHub.Service.Dashboard/ent/schema"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"RouteHub.Service.Dashboard/ent/schema/types"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	entdomainMixin := schema.Domain{}.Mixin()
	entdomainMixinFields0 := entdomainMixin[0].Fields()
	_ = entdomainMixinFields0
	entdomainFields := schema.Domain{}.Fields()
	_ = entdomainFields
	// entdomainDescName is the schema descriptor for name field.
	entdomainDescName := entdomainFields[0].Descriptor()
	// entdomain.NameValidator is a validator for the "name" field. It is called by the builders before save.
	entdomain.NameValidator = entdomainDescName.Validators[0].(func(string) error)
	// entdomainDescURL is the schema descriptor for url field.
	entdomainDescURL := entdomainFields[1].Descriptor()
	// entdomain.URLValidator is a validator for the "url" field. It is called by the builders before save.
	entdomain.URLValidator = entdomainDescURL.Validators[0].(func(string) error)
	// entdomainDescCreatedAt is the schema descriptor for created_at field.
	entdomainDescCreatedAt := entdomainFields[3].Descriptor()
	// entdomain.DefaultCreatedAt holds the default value on creation for the created_at field.
	entdomain.DefaultCreatedAt = entdomainDescCreatedAt.Default.(func() time.Time)
	// entdomainDescID is the schema descriptor for id field.
	entdomainDescID := entdomainMixinFields0[0].Descriptor()
	// entdomain.DefaultID holds the default value on creation for the id field.
	entdomain.DefaultID = entdomainDescID.Default.(func() mixin.ID)
	hubMixin := schema.Hub{}.Mixin()
	hubHooks := schema.Hub{}.Hooks()
	hub.Hooks[0] = hubHooks[0]
	hubMixinFields0 := hubMixin[0].Fields()
	_ = hubMixinFields0
	hubFields := schema.Hub{}.Fields()
	_ = hubFields
	// hubDescName is the schema descriptor for name field.
	hubDescName := hubFields[0].Descriptor()
	// hub.NameValidator is a validator for the "name" field. It is called by the builders before save.
	hub.NameValidator = hubDescName.Validators[0].(func(string) error)
	// hubDescSlug is the schema descriptor for slug field.
	hubDescSlug := hubFields[1].Descriptor()
	// hub.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	hub.SlugValidator = hubDescSlug.Validators[0].(func(string) error)
	// hubDescTCPAddress is the schema descriptor for tcp_address field.
	hubDescTCPAddress := hubFields[3].Descriptor()
	// hub.TCPAddressValidator is a validator for the "tcp_address" field. It is called by the builders before save.
	hub.TCPAddressValidator = hubDescTCPAddress.Validators[0].(func(string) error)
	// hubDescCreatedAt is the schema descriptor for created_at field.
	hubDescCreatedAt := hubFields[6].Descriptor()
	// hub.DefaultCreatedAt holds the default value on creation for the created_at field.
	hub.DefaultCreatedAt = hubDescCreatedAt.Default.(func() time.Time)
	// hubDescID is the schema descriptor for id field.
	hubDescID := hubMixinFields0[0].Descriptor()
	// hub.DefaultID holds the default value on creation for the id field.
	hub.DefaultID = hubDescID.Default.(func() mixin.ID)
	linkMixin := schema.Link{}.Mixin()
	linkHooks := schema.Link{}.Hooks()
	link.Hooks[0] = linkHooks[0]
	linkMixinFields0 := linkMixin[0].Fields()
	_ = linkMixinFields0
	linkFields := schema.Link{}.Fields()
	_ = linkFields
	// linkDescTarget is the schema descriptor for target field.
	linkDescTarget := linkFields[0].Descriptor()
	// link.TargetValidator is a validator for the "target" field. It is called by the builders before save.
	link.TargetValidator = linkDescTarget.Validators[0].(func(string) error)
	// linkDescPath is the schema descriptor for path field.
	linkDescPath := linkFields[1].Descriptor()
	// link.PathValidator is a validator for the "path" field. It is called by the builders before save.
	link.PathValidator = linkDescPath.Validators[0].(func(string) error)
	// linkDescLinkContent is the schema descriptor for link_content field.
	linkDescLinkContent := linkFields[2].Descriptor()
	// link.DefaultLinkContent holds the default value on creation for the link_content field.
	link.DefaultLinkContent = linkDescLinkContent.Default.(*types.LinkContent)
	// linkDescCreatedAt is the schema descriptor for created_at field.
	linkDescCreatedAt := linkFields[5].Descriptor()
	// link.DefaultCreatedAt holds the default value on creation for the created_at field.
	link.DefaultCreatedAt = linkDescCreatedAt.Default.(func() time.Time)
	// linkDescID is the schema descriptor for id field.
	linkDescID := linkMixinFields0[0].Descriptor()
	// link.DefaultID holds the default value on creation for the id field.
	link.DefaultID = linkDescID.Default.(func() mixin.ID)
	organizationMixin := schema.Organization{}.Mixin()
	organizationMixinFields0 := organizationMixin[0].Fields()
	_ = organizationMixinFields0
	organizationFields := schema.Organization{}.Fields()
	_ = organizationFields
	// organizationDescName is the schema descriptor for name field.
	organizationDescName := organizationFields[0].Descriptor()
	// organization.NameValidator is a validator for the "name" field. It is called by the builders before save.
	organization.NameValidator = organizationDescName.Validators[0].(func(string) error)
	// organizationDescCreatedAt is the schema descriptor for created_at field.
	organizationDescCreatedAt := organizationFields[5].Descriptor()
	// organization.DefaultCreatedAt holds the default value on creation for the created_at field.
	organization.DefaultCreatedAt = organizationDescCreatedAt.Default.(func() time.Time)
	// organizationDescID is the schema descriptor for id field.
	organizationDescID := organizationMixinFields0[0].Descriptor()
	// organization.DefaultID holds the default value on creation for the id field.
	organization.DefaultID = organizationDescID.Default.(func() mixin.ID)
	pageMixin := schema.Page{}.Mixin()
	pageHooks := schema.Page{}.Hooks()
	page.Hooks[0] = pageHooks[0]
	pageMixinFields0 := pageMixin[0].Fields()
	_ = pageMixinFields0
	pageFields := schema.Page{}.Fields()
	_ = pageFields
	// pageDescName is the schema descriptor for name field.
	pageDescName := pageFields[0].Descriptor()
	// page.NameValidator is a validator for the "name" field. It is called by the builders before save.
	page.NameValidator = pageDescName.Validators[0].(func(string) error)
	// pageDescSlug is the schema descriptor for slug field.
	pageDescSlug := pageFields[1].Descriptor()
	// page.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	page.SlugValidator = pageDescSlug.Validators[0].(func(string) error)
	// pageDescCreatedAt is the schema descriptor for created_at field.
	pageDescCreatedAt := pageFields[7].Descriptor()
	// page.DefaultCreatedAt holds the default value on creation for the created_at field.
	page.DefaultCreatedAt = pageDescCreatedAt.Default.(func() time.Time)
	// pageDescID is the schema descriptor for id field.
	pageDescID := pageMixinFields0[0].Descriptor()
	// page.DefaultID holds the default value on creation for the id field.
	page.DefaultID = pageDescID.Default.(func() mixin.ID)
	personMixin := schema.Person{}.Mixin()
	personMixinFields0 := personMixin[0].Fields()
	_ = personMixinFields0
	personFields := schema.Person{}.Fields()
	_ = personFields
	// personDescSubjectID is the schema descriptor for subject_id field.
	personDescSubjectID := personFields[0].Descriptor()
	// person.SubjectIDValidator is a validator for the "subject_id" field. It is called by the builders before save.
	person.SubjectIDValidator = personDescSubjectID.Validators[0].(func(string) error)
	// personDescIsActive is the schema descriptor for is_active field.
	personDescIsActive := personFields[2].Descriptor()
	// person.DefaultIsActive holds the default value on creation for the is_active field.
	person.DefaultIsActive = personDescIsActive.Default.(bool)
	// personDescCreatedAt is the schema descriptor for created_at field.
	personDescCreatedAt := personFields[3].Descriptor()
	// person.DefaultCreatedAt holds the default value on creation for the created_at field.
	person.DefaultCreatedAt = personDescCreatedAt.Default.(func() time.Time)
	// personDescID is the schema descriptor for id field.
	personDescID := personMixinFields0[0].Descriptor()
	// person.DefaultID holds the default value on creation for the id field.
	person.DefaultID = personDescID.Default.(func() mixin.ID)
}

const (
	Version = "v0.14.1"                                         // Version of ent codegen.
	Sum     = "h1:fUERL506Pqr92EPHJqr8EYxbPioflJo6PudkrEA8a/s=" // Sum of ent codegen.
)
