// Code generated by ent, DO NOT EDIT.

package datafield

import (
	"entgo.io/bug/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// HasRecords applies the HasEdge predicate on the "records" edge.
func HasRecords() predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RecordsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RecordsTable, RecordsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRecordsWith applies the HasEdge predicate on the "records" edge with a given conditions (other predicates).
func HasRecordsWith(preds ...predicate.Record) predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RecordsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RecordsTable, RecordsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCells applies the HasEdge predicate on the "cells" edge.
func HasCells() predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CellsTable, CellsColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, CellsTable, CellsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCellsWith applies the HasEdge predicate on the "cells" edge with a given conditions (other predicates).
func HasCellsWith(preds ...predicate.Cell) predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CellsInverseTable, CellsColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, CellsTable, CellsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DataField) predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DataField) predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DataField) predicate.DataField {
	return predicate.DataField(func(s *sql.Selector) {
		p(s.Not())
	})
}
