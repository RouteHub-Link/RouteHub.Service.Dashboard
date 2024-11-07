// Code generated by ent, DO NOT EDIT.

package link

import (
	"RouteHub.Service.Dashboard/ent/predicate"
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id mixin.ID) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id mixin.ID) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id mixin.ID) predicate.Link {
	return predicate.Link(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...mixin.ID) predicate.Link {
	return predicate.Link(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...mixin.ID) predicate.Link {
	return predicate.Link(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id mixin.ID) predicate.Link {
	return predicate.Link(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id mixin.ID) predicate.Link {
	return predicate.Link(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id mixin.ID) predicate.Link {
	return predicate.Link(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id mixin.ID) predicate.Link {
	return predicate.Link(sql.FieldLTE(FieldID, id))
}

// Target applies equality check predicate on the "target" field. It's identical to TargetEQ.
func Target(v string) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldTarget, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldPath, v))
}

// TargetEQ applies the EQ predicate on the "target" field.
func TargetEQ(v string) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldTarget, v))
}

// TargetNEQ applies the NEQ predicate on the "target" field.
func TargetNEQ(v string) predicate.Link {
	return predicate.Link(sql.FieldNEQ(FieldTarget, v))
}

// TargetIn applies the In predicate on the "target" field.
func TargetIn(vs ...string) predicate.Link {
	return predicate.Link(sql.FieldIn(FieldTarget, vs...))
}

// TargetNotIn applies the NotIn predicate on the "target" field.
func TargetNotIn(vs ...string) predicate.Link {
	return predicate.Link(sql.FieldNotIn(FieldTarget, vs...))
}

// TargetGT applies the GT predicate on the "target" field.
func TargetGT(v string) predicate.Link {
	return predicate.Link(sql.FieldGT(FieldTarget, v))
}

// TargetGTE applies the GTE predicate on the "target" field.
func TargetGTE(v string) predicate.Link {
	return predicate.Link(sql.FieldGTE(FieldTarget, v))
}

// TargetLT applies the LT predicate on the "target" field.
func TargetLT(v string) predicate.Link {
	return predicate.Link(sql.FieldLT(FieldTarget, v))
}

// TargetLTE applies the LTE predicate on the "target" field.
func TargetLTE(v string) predicate.Link {
	return predicate.Link(sql.FieldLTE(FieldTarget, v))
}

// TargetContains applies the Contains predicate on the "target" field.
func TargetContains(v string) predicate.Link {
	return predicate.Link(sql.FieldContains(FieldTarget, v))
}

// TargetHasPrefix applies the HasPrefix predicate on the "target" field.
func TargetHasPrefix(v string) predicate.Link {
	return predicate.Link(sql.FieldHasPrefix(FieldTarget, v))
}

// TargetHasSuffix applies the HasSuffix predicate on the "target" field.
func TargetHasSuffix(v string) predicate.Link {
	return predicate.Link(sql.FieldHasSuffix(FieldTarget, v))
}

// TargetEqualFold applies the EqualFold predicate on the "target" field.
func TargetEqualFold(v string) predicate.Link {
	return predicate.Link(sql.FieldEqualFold(FieldTarget, v))
}

// TargetContainsFold applies the ContainsFold predicate on the "target" field.
func TargetContainsFold(v string) predicate.Link {
	return predicate.Link(sql.FieldContainsFold(FieldTarget, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.Link {
	return predicate.Link(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.Link {
	return predicate.Link(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.Link {
	return predicate.Link(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.Link {
	return predicate.Link(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.Link {
	return predicate.Link(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.Link {
	return predicate.Link(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.Link {
	return predicate.Link(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.Link {
	return predicate.Link(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.Link {
	return predicate.Link(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.Link {
	return predicate.Link(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.Link {
	return predicate.Link(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.Link {
	return predicate.Link(sql.FieldContainsFold(FieldPath, v))
}

// LinkContentIsNil applies the IsNil predicate on the "link_content" field.
func LinkContentIsNil() predicate.Link {
	return predicate.Link(sql.FieldIsNull(FieldLinkContent))
}

// LinkContentNotNil applies the NotNil predicate on the "link_content" field.
func LinkContentNotNil() predicate.Link {
	return predicate.Link(sql.FieldNotNull(FieldLinkContent))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v enums.StatusState) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v enums.StatusState) predicate.Link {
	return predicate.Link(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...enums.StatusState) predicate.Link {
	return predicate.Link(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...enums.StatusState) predicate.Link {
	return predicate.Link(sql.FieldNotIn(FieldStatus, vs...))
}

// HasHub applies the HasEdge predicate on the "hub" edge.
func HasHub() predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HubTable, HubColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHubWith applies the HasEdge predicate on the "hub" edge with a given conditions (other predicates).
func HasHubWith(preds ...predicate.Hub) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		step := newHubStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Link) predicate.Link {
	return predicate.Link(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Link) predicate.Link {
	return predicate.Link(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Link) predicate.Link {
	return predicate.Link(sql.NotPredicates(p))
}
