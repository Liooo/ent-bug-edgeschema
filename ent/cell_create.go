// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/bug/ent/cell"
	"entgo.io/bug/ent/datafield"
	"entgo.io/bug/ent/record"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CellCreate is the builder for creating a Cell entity.
type CellCreate struct {
	config
	mutation *CellMutation
	hooks    []Hook
}

// SetDataFieldID sets the "data_field_id" field.
func (cc *CellCreate) SetDataFieldID(u uuid.UUID) *CellCreate {
	cc.mutation.SetDataFieldID(u)
	return cc
}

// SetRecordID sets the "record_id" field.
func (cc *CellCreate) SetRecordID(u uuid.UUID) *CellCreate {
	cc.mutation.SetRecordID(u)
	return cc
}

// SetRecord sets the "record" edge to the Record entity.
func (cc *CellCreate) SetRecord(r *Record) *CellCreate {
	return cc.SetRecordID(r.ID)
}

// SetDataField sets the "data_field" edge to the DataField entity.
func (cc *CellCreate) SetDataField(d *DataField) *CellCreate {
	return cc.SetDataFieldID(d.ID)
}

// Mutation returns the CellMutation object of the builder.
func (cc *CellCreate) Mutation() *CellMutation {
	return cc.mutation
}

// Save creates the Cell in the database.
func (cc *CellCreate) Save(ctx context.Context) (*Cell, error) {
	var (
		err  error
		node *Cell
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CellMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Cell)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CellMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CellCreate) SaveX(ctx context.Context) *Cell {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CellCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CellCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CellCreate) check() error {
	if _, ok := cc.mutation.DataFieldID(); !ok {
		return &ValidationError{Name: "data_field_id", err: errors.New(`ent: missing required field "Cell.data_field_id"`)}
	}
	if _, ok := cc.mutation.RecordID(); !ok {
		return &ValidationError{Name: "record_id", err: errors.New(`ent: missing required field "Cell.record_id"`)}
	}
	if _, ok := cc.mutation.RecordID(); !ok {
		return &ValidationError{Name: "record", err: errors.New(`ent: missing required edge "Cell.record"`)}
	}
	if _, ok := cc.mutation.DataFieldID(); !ok {
		return &ValidationError{Name: "data_field", err: errors.New(`ent: missing required edge "Cell.data_field"`)}
	}
	return nil
}

func (cc *CellCreate) sqlSave(ctx context.Context) (*Cell, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (cc *CellCreate) createSpec() (*Cell, *sqlgraph.CreateSpec) {
	var (
		_node = &Cell{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cell.Table,
		}
	)
	if nodes := cc.mutation.RecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cell.RecordTable,
			Columns: []string{cell.RecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: record.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RecordID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.DataFieldIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cell.DataFieldTable,
			Columns: []string{cell.DataFieldColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: datafield.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DataFieldID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CellCreateBulk is the builder for creating many Cell entities in bulk.
type CellCreateBulk struct {
	config
	builders []*CellCreate
}

// Save creates the Cell entities in the database.
func (ccb *CellCreateBulk) Save(ctx context.Context) ([]*Cell, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Cell, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CellMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CellCreateBulk) SaveX(ctx context.Context) []*Cell {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CellCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CellCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
