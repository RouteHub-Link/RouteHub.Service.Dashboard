// Code generated by ent, DO NOT EDIT.

package organization

import (
	"RouteHub.Service.Dashboard/ent/predicate"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id mixin.ID) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id mixin.ID) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id mixin.ID) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...mixin.ID) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...mixin.ID) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id mixin.ID) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id mixin.ID) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id mixin.ID) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id mixin.ID) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldName, v))
}

// Website applies equality check predicate on the "website" field. It's identical to WebsiteEQ.
func Website(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldWebsite, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDescription, v))
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldLocation, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldName, v))
}

// WebsiteEQ applies the EQ predicate on the "website" field.
func WebsiteEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldWebsite, v))
}

// WebsiteNEQ applies the NEQ predicate on the "website" field.
func WebsiteNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldWebsite, v))
}

// WebsiteIn applies the In predicate on the "website" field.
func WebsiteIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldWebsite, vs...))
}

// WebsiteNotIn applies the NotIn predicate on the "website" field.
func WebsiteNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldWebsite, vs...))
}

// WebsiteGT applies the GT predicate on the "website" field.
func WebsiteGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldWebsite, v))
}

// WebsiteGTE applies the GTE predicate on the "website" field.
func WebsiteGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldWebsite, v))
}

// WebsiteLT applies the LT predicate on the "website" field.
func WebsiteLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldWebsite, v))
}

// WebsiteLTE applies the LTE predicate on the "website" field.
func WebsiteLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldWebsite, v))
}

// WebsiteContains applies the Contains predicate on the "website" field.
func WebsiteContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldWebsite, v))
}

// WebsiteHasPrefix applies the HasPrefix predicate on the "website" field.
func WebsiteHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldWebsite, v))
}

// WebsiteHasSuffix applies the HasSuffix predicate on the "website" field.
func WebsiteHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldWebsite, v))
}

// WebsiteIsNil applies the IsNil predicate on the "website" field.
func WebsiteIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldWebsite))
}

// WebsiteNotNil applies the NotNil predicate on the "website" field.
func WebsiteNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldWebsite))
}

// WebsiteEqualFold applies the EqualFold predicate on the "website" field.
func WebsiteEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldWebsite, v))
}

// WebsiteContainsFold applies the ContainsFold predicate on the "website" field.
func WebsiteContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldWebsite, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldDescription, v))
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldLocation, v))
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldLocation, v))
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldLocation, vs...))
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldLocation, vs...))
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldLocation, v))
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldLocation, v))
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldLocation, v))
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldLocation, v))
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldLocation, v))
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldLocation, v))
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldLocation, v))
}

// LocationIsNil applies the IsNil predicate on the "location" field.
func LocationIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldLocation))
}

// LocationNotNil applies the NotNil predicate on the "location" field.
func LocationNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldLocation))
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldLocation, v))
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldLocation, v))
}

// SocialMediasIsNil applies the IsNil predicate on the "social_medias" field.
func SocialMediasIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldSocialMedias))
}

// SocialMediasNotNil applies the NotNil predicate on the "social_medias" field.
func SocialMediasNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldSocialMedias))
}

// HasDomains applies the HasEdge predicate on the "domains" edge.
func HasDomains() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DomainsTable, DomainsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDomainsWith applies the HasEdge predicate on the "domains" edge with a given conditions (other predicates).
func HasDomainsWith(preds ...predicate.Domain) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newDomainsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHubs applies the HasEdge predicate on the "hubs" edge.
func HasHubs() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HubsTable, HubsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHubsWith applies the HasEdge predicate on the "hubs" edge with a given conditions (other predicates).
func HasHubsWith(preds ...predicate.Hub) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newHubsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPersons applies the HasEdge predicate on the "persons" edge.
func HasPersons() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PersonsTable, PersonsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonsWith applies the HasEdge predicate on the "persons" edge with a given conditions (other predicates).
func HasPersonsWith(preds ...predicate.Person) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newPersonsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Organization) predicate.Organization {
	return predicate.Organization(sql.NotPredicates(p))
}
