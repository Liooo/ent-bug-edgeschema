// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/bug/ent/cell"
	"entgo.io/bug/ent/datafield"
	"entgo.io/bug/ent/record"
	"github.com/google/uuid"
)

// Cell is the model entity for the Cell schema.
type Cell struct {
	config `json:"-"`
	// DataFieldID holds the value of the "data_field_id" field.
	DataFieldID uuid.UUID `json:"data_field_id,omitempty"`
	// RecordID holds the value of the "record_id" field.
	RecordID uuid.UUID `json:"record_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CellQuery when eager-loading is set.
	Edges CellEdges `json:"edges"`
}

// CellEdges holds the relations/edges for other nodes in the graph.
type CellEdges struct {
	// Record holds the value of the record edge.
	Record *Record `json:"record,omitempty"`
	// DataField holds the value of the data_field edge.
	DataField *DataField `json:"data_field,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RecordOrErr returns the Record value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CellEdges) RecordOrErr() (*Record, error) {
	if e.loadedTypes[0] {
		if e.Record == nil {
			// The edge record was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: record.Label}
		}
		return e.Record, nil
	}
	return nil, &NotLoadedError{edge: "record"}
}

// DataFieldOrErr returns the DataField value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CellEdges) DataFieldOrErr() (*DataField, error) {
	if e.loadedTypes[1] {
		if e.DataField == nil {
			// The edge data_field was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: datafield.Label}
		}
		return e.DataField, nil
	}
	return nil, &NotLoadedError{edge: "data_field"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cell) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case cell.FieldDataFieldID, cell.FieldRecordID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Cell", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cell fields.
func (c *Cell) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cell.FieldDataFieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field data_field_id", values[i])
			} else if value != nil {
				c.DataFieldID = *value
			}
		case cell.FieldRecordID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field record_id", values[i])
			} else if value != nil {
				c.RecordID = *value
			}
		}
	}
	return nil
}

// QueryRecord queries the "record" edge of the Cell entity.
func (c *Cell) QueryRecord() *RecordQuery {
	return (&CellClient{config: c.config}).QueryRecord(c)
}

// QueryDataField queries the "data_field" edge of the Cell entity.
func (c *Cell) QueryDataField() *DataFieldQuery {
	return (&CellClient{config: c.config}).QueryDataField(c)
}

// Update returns a builder for updating this Cell.
// Note that you need to call Cell.Unwrap() before calling this method if this Cell
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cell) Update() *CellUpdateOne {
	return (&CellClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Cell entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cell) Unwrap() *Cell {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cell is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cell) String() string {
	var builder strings.Builder
	builder.WriteString("Cell(")
	builder.WriteString("data_field_id=")
	builder.WriteString(fmt.Sprintf("%v", c.DataFieldID))
	builder.WriteString(", ")
	builder.WriteString("record_id=")
	builder.WriteString(fmt.Sprintf("%v", c.RecordID))
	builder.WriteByte(')')
	return builder.String()
}

// Cells is a parsable slice of Cell.
type Cells []*Cell

func (c Cells) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
