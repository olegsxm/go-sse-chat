// Code generated by ent, DO NOT EDIT.

package conversation

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/olegsxm/go-sse-chat/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldName, v))
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldAvatar, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Conversation {
	return predicate.Conversation(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Conversation {
	return predicate.Conversation(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Conversation {
	return predicate.Conversation(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Conversation {
	return predicate.Conversation(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldContainsFold(FieldName, v))
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldAvatar, v))
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldNEQ(FieldAvatar, v))
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.Conversation {
	return predicate.Conversation(sql.FieldIn(FieldAvatar, vs...))
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.Conversation {
	return predicate.Conversation(sql.FieldNotIn(FieldAvatar, vs...))
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldGT(FieldAvatar, v))
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldGTE(FieldAvatar, v))
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldLT(FieldAvatar, v))
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldLTE(FieldAvatar, v))
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldContains(FieldAvatar, v))
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldHasPrefix(FieldAvatar, v))
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldHasSuffix(FieldAvatar, v))
}

// AvatarIsNil applies the IsNil predicate on the "avatar" field.
func AvatarIsNil() predicate.Conversation {
	return predicate.Conversation(sql.FieldIsNull(FieldAvatar))
}

// AvatarNotNil applies the NotNil predicate on the "avatar" field.
func AvatarNotNil() predicate.Conversation {
	return predicate.Conversation(sql.FieldNotNull(FieldAvatar))
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldEqualFold(FieldAvatar, v))
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldContainsFold(FieldAvatar, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Conversation {
	return predicate.Conversation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Conversation {
	return predicate.Conversation(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMessages applies the HasEdge predicate on the "messages" edge.
func HasMessages() predicate.Conversation {
	return predicate.Conversation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, MessagesTable, MessagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMessagesWith applies the HasEdge predicate on the "messages" edge with a given conditions (other predicates).
func HasMessagesWith(preds ...predicate.Message) predicate.Conversation {
	return predicate.Conversation(func(s *sql.Selector) {
		step := newMessagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Conversation) predicate.Conversation {
	return predicate.Conversation(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Conversation) predicate.Conversation {
	return predicate.Conversation(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Conversation) predicate.Conversation {
	return predicate.Conversation(sql.NotPredicates(p))
}
