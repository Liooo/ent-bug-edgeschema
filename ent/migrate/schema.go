// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CellsColumns holds the columns for the "cells" table.
	CellsColumns = []*schema.Column{
		{Name: "record_id", Type: field.TypeUUID},
		{Name: "data_field_id", Type: field.TypeUUID},
	}
	// CellsTable holds the schema information for the "cells" table.
	CellsTable = &schema.Table{
		Name:       "cells",
		Columns:    CellsColumns,
		PrimaryKey: []*schema.Column{},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cells_records_record",
				Columns:    []*schema.Column{CellsColumns[0]},
				RefColumns: []*schema.Column{RecordsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "cells_data_fields_data_field",
				Columns:    []*schema.Column{CellsColumns[1]},
				RefColumns: []*schema.Column{DataFieldsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// DataFieldsColumns holds the columns for the "data_fields" table.
	DataFieldsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
	}
	// DataFieldsTable holds the schema information for the "data_fields" table.
	DataFieldsTable = &schema.Table{
		Name:       "data_fields",
		Columns:    DataFieldsColumns,
		PrimaryKey: []*schema.Column{DataFieldsColumns[0]},
	}
	// RecordsColumns holds the columns for the "records" table.
	RecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
	}
	// RecordsTable holds the schema information for the "records" table.
	RecordsTable = &schema.Table{
		Name:       "records",
		Columns:    RecordsColumns,
		PrimaryKey: []*schema.Column{RecordsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CellsTable,
		DataFieldsTable,
		RecordsTable,
	}
)

func init() {
	CellsTable.ForeignKeys[0].RefTable = RecordsTable
	CellsTable.ForeignKeys[1].RefTable = DataFieldsTable
}
