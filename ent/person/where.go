// Code generated by ent, DO NOT EDIT.

package person

import (
	"RouteHub.Service.Dashboard/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.jetify.com/typeid"
)

// ID filters vertices based on their ID field.
func ID(id typeid.AnyID) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id typeid.AnyID) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id typeid.AnyID) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...typeid.AnyID) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...typeid.AnyID) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id typeid.AnyID) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id typeid.AnyID) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id typeid.AnyID) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id typeid.AnyID) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldID, id))
}

// SubjectID applies equality check predicate on the "subject_id" field. It's identical to SubjectIDEQ.
func SubjectID(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldSubjectID, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldIsActive, v))
}

// SubjectIDEQ applies the EQ predicate on the "subject_id" field.
func SubjectIDEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldSubjectID, v))
}

// SubjectIDNEQ applies the NEQ predicate on the "subject_id" field.
func SubjectIDNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldSubjectID, v))
}

// SubjectIDIn applies the In predicate on the "subject_id" field.
func SubjectIDIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldSubjectID, vs...))
}

// SubjectIDNotIn applies the NotIn predicate on the "subject_id" field.
func SubjectIDNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldSubjectID, vs...))
}

// SubjectIDGT applies the GT predicate on the "subject_id" field.
func SubjectIDGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldSubjectID, v))
}

// SubjectIDGTE applies the GTE predicate on the "subject_id" field.
func SubjectIDGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldSubjectID, v))
}

// SubjectIDLT applies the LT predicate on the "subject_id" field.
func SubjectIDLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldSubjectID, v))
}

// SubjectIDLTE applies the LTE predicate on the "subject_id" field.
func SubjectIDLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldSubjectID, v))
}

// SubjectIDContains applies the Contains predicate on the "subject_id" field.
func SubjectIDContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldSubjectID, v))
}

// SubjectIDHasPrefix applies the HasPrefix predicate on the "subject_id" field.
func SubjectIDHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldSubjectID, v))
}

// SubjectIDHasSuffix applies the HasSuffix predicate on the "subject_id" field.
func SubjectIDHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldSubjectID, v))
}

// SubjectIDEqualFold applies the EqualFold predicate on the "subject_id" field.
func SubjectIDEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldSubjectID, v))
}

// SubjectIDContainsFold applies the ContainsFold predicate on the "subject_id" field.
func SubjectIDContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldSubjectID, v))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldIsActive, v))
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Person) predicate.Person {
	return predicate.Person(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Person) predicate.Person {
	return predicate.Person(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Person) predicate.Person {
	return predicate.Person(sql.NotPredicates(p))
}