// Code generated by ent, DO NOT EDIT.

package cell

import (
	"entgo.io/bug/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// DataFieldID applies equality check predicate on the "data_field_id" field. It's identical to DataFieldIDEQ.
func DataFieldID(v uuid.UUID) predicate.Cell {
	return predicate.Cell(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDataFieldID), v))
	})
}

// RecordID applies equality check predicate on the "record_id" field. It's identical to RecordIDEQ.
func RecordID(v uuid.UUID) predicate.Cell {
	return predicate.Cell(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecordID), v))
	})
}

// DataFieldIDEQ applies the EQ predicate on the "data_field_id" field.
func DataFieldIDEQ(v uuid.UUID) predicate.Cell {
	return predicate.Cell(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDataFieldID), v))
	})
}

// DataFieldIDNEQ applies the NEQ predicate on the "data_field_id" field.
func DataFieldIDNEQ(v uuid.UUID) predicate.Cell {
	return predicate.Cell(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDataFieldID), v))
	})
}

// DataFieldIDIn applies the In predicate on the "data_field_id" field.
func DataFieldIDIn(vs ...uuid.UUID) predicate.Cell {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cell(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDataFieldID), v...))
	})
}

// DataFieldIDNotIn applies the NotIn predicate on the "data_field_id" field.
func DataFieldIDNotIn(vs ...uuid.UUID) predicate.Cell {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cell(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDataFieldID), v...))
	})
}

// RecordIDEQ applies the EQ predicate on the "record_id" field.
func RecordIDEQ(v uuid.UUID) predicate.Cell {
	return predicate.Cell(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecordID), v))
	})
}

// RecordIDNEQ applies the NEQ predicate on the "record_id" field.
func RecordIDNEQ(v uuid.UUID) predicate.Cell {
	return predicate.Cell(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecordID), v))
	})
}

// RecordIDIn applies the In predicate on the "record_id" field.
func RecordIDIn(vs ...uuid.UUID) predicate.Cell {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cell(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRecordID), v...))
	})
}

// RecordIDNotIn applies the NotIn predicate on the "record_id" field.
func RecordIDNotIn(vs ...uuid.UUID) predicate.Cell {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cell(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRecordID), v...))
	})
}

// HasRecord applies the HasEdge predicate on the "record" edge.
func HasRecord() predicate.Cell {
	return predicate.Cell(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, RecordColumn),
			sqlgraph.To(RecordInverseTable, RecordFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RecordTable, RecordColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRecordWith applies the HasEdge predicate on the "record" edge with a given conditions (other predicates).
func HasRecordWith(preds ...predicate.Record) predicate.Cell {
	return predicate.Cell(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, RecordColumn),
			sqlgraph.To(RecordInverseTable, RecordFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RecordTable, RecordColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDataField applies the HasEdge predicate on the "data_field" edge.
func HasDataField() predicate.Cell {
	return predicate.Cell(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, DataFieldColumn),
			sqlgraph.To(DataFieldInverseTable, DataFieldFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DataFieldTable, DataFieldColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDataFieldWith applies the HasEdge predicate on the "data_field" edge with a given conditions (other predicates).
func HasDataFieldWith(preds ...predicate.DataField) predicate.Cell {
	return predicate.Cell(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, DataFieldColumn),
			sqlgraph.To(DataFieldInverseTable, DataFieldFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DataFieldTable, DataFieldColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Cell) predicate.Cell {
	return predicate.Cell(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Cell) predicate.Cell {
	return predicate.Cell(func(s *sql.Selector) {
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
func Not(p predicate.Cell) predicate.Cell {
	return predicate.Cell(func(s *sql.Selector) {
		p(s.Not())
	})
}
