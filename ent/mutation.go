// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/bug/ent/cell"
	"entgo.io/bug/ent/datafield"
	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/record"
	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCell      = "Cell"
	TypeDataField = "DataField"
	TypeRecord    = "Record"
)

// CellMutation represents an operation that mutates the Cell nodes in the graph.
type CellMutation struct {
	config
	op                Op
	typ               string
	clearedFields     map[string]struct{}
	record            *uuid.UUID
	clearedrecord     bool
	data_field        *uuid.UUID
	cleareddata_field bool
	done              bool
	oldValue          func(context.Context) (*Cell, error)
	predicates        []predicate.Cell
}

var _ ent.Mutation = (*CellMutation)(nil)

// cellOption allows management of the mutation configuration using functional options.
type cellOption func(*CellMutation)

// newCellMutation creates new mutation for the Cell entity.
func newCellMutation(c config, op Op, opts ...cellOption) *CellMutation {
	m := &CellMutation{
		config:        c,
		op:            op,
		typ:           TypeCell,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CellMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CellMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetDataFieldID sets the "data_field_id" field.
func (m *CellMutation) SetDataFieldID(u uuid.UUID) {
	m.data_field = &u
}

// DataFieldID returns the value of the "data_field_id" field in the mutation.
func (m *CellMutation) DataFieldID() (r uuid.UUID, exists bool) {
	v := m.data_field
	if v == nil {
		return
	}
	return *v, true
}

// ResetDataFieldID resets all changes to the "data_field_id" field.
func (m *CellMutation) ResetDataFieldID() {
	m.data_field = nil
}

// SetRecordID sets the "record_id" field.
func (m *CellMutation) SetRecordID(u uuid.UUID) {
	m.record = &u
}

// RecordID returns the value of the "record_id" field in the mutation.
func (m *CellMutation) RecordID() (r uuid.UUID, exists bool) {
	v := m.record
	if v == nil {
		return
	}
	return *v, true
}

// ResetRecordID resets all changes to the "record_id" field.
func (m *CellMutation) ResetRecordID() {
	m.record = nil
}

// ClearRecord clears the "record" edge to the Record entity.
func (m *CellMutation) ClearRecord() {
	m.clearedrecord = true
}

// RecordCleared reports if the "record" edge to the Record entity was cleared.
func (m *CellMutation) RecordCleared() bool {
	return m.clearedrecord
}

// RecordIDs returns the "record" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// RecordID instead. It exists only for internal usage by the builders.
func (m *CellMutation) RecordIDs() (ids []uuid.UUID) {
	if id := m.record; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetRecord resets all changes to the "record" edge.
func (m *CellMutation) ResetRecord() {
	m.record = nil
	m.clearedrecord = false
}

// ClearDataField clears the "data_field" edge to the DataField entity.
func (m *CellMutation) ClearDataField() {
	m.cleareddata_field = true
}

// DataFieldCleared reports if the "data_field" edge to the DataField entity was cleared.
func (m *CellMutation) DataFieldCleared() bool {
	return m.cleareddata_field
}

// DataFieldIDs returns the "data_field" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// DataFieldID instead. It exists only for internal usage by the builders.
func (m *CellMutation) DataFieldIDs() (ids []uuid.UUID) {
	if id := m.data_field; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetDataField resets all changes to the "data_field" edge.
func (m *CellMutation) ResetDataField() {
	m.data_field = nil
	m.cleareddata_field = false
}

// Where appends a list predicates to the CellMutation builder.
func (m *CellMutation) Where(ps ...predicate.Cell) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CellMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Cell).
func (m *CellMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CellMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.data_field != nil {
		fields = append(fields, cell.FieldDataFieldID)
	}
	if m.record != nil {
		fields = append(fields, cell.FieldRecordID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CellMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case cell.FieldDataFieldID:
		return m.DataFieldID()
	case cell.FieldRecordID:
		return m.RecordID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CellMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, errors.New("edge schema Cell does not support getting old values")
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CellMutation) SetField(name string, value ent.Value) error {
	switch name {
	case cell.FieldDataFieldID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDataFieldID(v)
		return nil
	case cell.FieldRecordID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRecordID(v)
		return nil
	}
	return fmt.Errorf("unknown Cell field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CellMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CellMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CellMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Cell numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CellMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CellMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CellMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Cell nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CellMutation) ResetField(name string) error {
	switch name {
	case cell.FieldDataFieldID:
		m.ResetDataFieldID()
		return nil
	case cell.FieldRecordID:
		m.ResetRecordID()
		return nil
	}
	return fmt.Errorf("unknown Cell field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CellMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.record != nil {
		edges = append(edges, cell.EdgeRecord)
	}
	if m.data_field != nil {
		edges = append(edges, cell.EdgeDataField)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CellMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case cell.EdgeRecord:
		if id := m.record; id != nil {
			return []ent.Value{*id}
		}
	case cell.EdgeDataField:
		if id := m.data_field; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CellMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CellMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CellMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedrecord {
		edges = append(edges, cell.EdgeRecord)
	}
	if m.cleareddata_field {
		edges = append(edges, cell.EdgeDataField)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CellMutation) EdgeCleared(name string) bool {
	switch name {
	case cell.EdgeRecord:
		return m.clearedrecord
	case cell.EdgeDataField:
		return m.cleareddata_field
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CellMutation) ClearEdge(name string) error {
	switch name {
	case cell.EdgeRecord:
		m.ClearRecord()
		return nil
	case cell.EdgeDataField:
		m.ClearDataField()
		return nil
	}
	return fmt.Errorf("unknown Cell unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CellMutation) ResetEdge(name string) error {
	switch name {
	case cell.EdgeRecord:
		m.ResetRecord()
		return nil
	case cell.EdgeDataField:
		m.ResetDataField()
		return nil
	}
	return fmt.Errorf("unknown Cell edge %s", name)
}

// DataFieldMutation represents an operation that mutates the DataField nodes in the graph.
type DataFieldMutation struct {
	config
	op             Op
	typ            string
	id             *uuid.UUID
	clearedFields  map[string]struct{}
	records        map[uuid.UUID]struct{}
	removedrecords map[uuid.UUID]struct{}
	clearedrecords bool
	done           bool
	oldValue       func(context.Context) (*DataField, error)
	predicates     []predicate.DataField
}

var _ ent.Mutation = (*DataFieldMutation)(nil)

// datafieldOption allows management of the mutation configuration using functional options.
type datafieldOption func(*DataFieldMutation)

// newDataFieldMutation creates new mutation for the DataField entity.
func newDataFieldMutation(c config, op Op, opts ...datafieldOption) *DataFieldMutation {
	m := &DataFieldMutation{
		config:        c,
		op:            op,
		typ:           TypeDataField,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withDataFieldID sets the ID field of the mutation.
func withDataFieldID(id uuid.UUID) datafieldOption {
	return func(m *DataFieldMutation) {
		var (
			err   error
			once  sync.Once
			value *DataField
		)
		m.oldValue = func(ctx context.Context) (*DataField, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().DataField.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withDataField sets the old DataField of the mutation.
func withDataField(node *DataField) datafieldOption {
	return func(m *DataFieldMutation) {
		m.oldValue = func(context.Context) (*DataField, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m DataFieldMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m DataFieldMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of DataField entities.
func (m *DataFieldMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *DataFieldMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *DataFieldMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().DataField.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// AddRecordIDs adds the "records" edge to the Record entity by ids.
func (m *DataFieldMutation) AddRecordIDs(ids ...uuid.UUID) {
	if m.records == nil {
		m.records = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.records[ids[i]] = struct{}{}
	}
}

// ClearRecords clears the "records" edge to the Record entity.
func (m *DataFieldMutation) ClearRecords() {
	m.clearedrecords = true
}

// RecordsCleared reports if the "records" edge to the Record entity was cleared.
func (m *DataFieldMutation) RecordsCleared() bool {
	return m.clearedrecords
}

// RemoveRecordIDs removes the "records" edge to the Record entity by IDs.
func (m *DataFieldMutation) RemoveRecordIDs(ids ...uuid.UUID) {
	if m.removedrecords == nil {
		m.removedrecords = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.records, ids[i])
		m.removedrecords[ids[i]] = struct{}{}
	}
}

// RemovedRecords returns the removed IDs of the "records" edge to the Record entity.
func (m *DataFieldMutation) RemovedRecordsIDs() (ids []uuid.UUID) {
	for id := range m.removedrecords {
		ids = append(ids, id)
	}
	return
}

// RecordsIDs returns the "records" edge IDs in the mutation.
func (m *DataFieldMutation) RecordsIDs() (ids []uuid.UUID) {
	for id := range m.records {
		ids = append(ids, id)
	}
	return
}

// ResetRecords resets all changes to the "records" edge.
func (m *DataFieldMutation) ResetRecords() {
	m.records = nil
	m.clearedrecords = false
	m.removedrecords = nil
}

// Where appends a list predicates to the DataFieldMutation builder.
func (m *DataFieldMutation) Where(ps ...predicate.DataField) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *DataFieldMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (DataField).
func (m *DataFieldMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *DataFieldMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *DataFieldMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *DataFieldMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown DataField field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DataFieldMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown DataField field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *DataFieldMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *DataFieldMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DataFieldMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown DataField numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *DataFieldMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *DataFieldMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *DataFieldMutation) ClearField(name string) error {
	return fmt.Errorf("unknown DataField nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *DataFieldMutation) ResetField(name string) error {
	return fmt.Errorf("unknown DataField field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *DataFieldMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.records != nil {
		edges = append(edges, datafield.EdgeRecords)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *DataFieldMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case datafield.EdgeRecords:
		ids := make([]ent.Value, 0, len(m.records))
		for id := range m.records {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *DataFieldMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedrecords != nil {
		edges = append(edges, datafield.EdgeRecords)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *DataFieldMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case datafield.EdgeRecords:
		ids := make([]ent.Value, 0, len(m.removedrecords))
		for id := range m.removedrecords {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *DataFieldMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedrecords {
		edges = append(edges, datafield.EdgeRecords)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *DataFieldMutation) EdgeCleared(name string) bool {
	switch name {
	case datafield.EdgeRecords:
		return m.clearedrecords
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *DataFieldMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown DataField unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *DataFieldMutation) ResetEdge(name string) error {
	switch name {
	case datafield.EdgeRecords:
		m.ResetRecords()
		return nil
	}
	return fmt.Errorf("unknown DataField edge %s", name)
}

// RecordMutation represents an operation that mutates the Record nodes in the graph.
type RecordMutation struct {
	config
	op                 Op
	typ                string
	id                 *uuid.UUID
	clearedFields      map[string]struct{}
	data_fields        map[uuid.UUID]struct{}
	removeddata_fields map[uuid.UUID]struct{}
	cleareddata_fields bool
	done               bool
	oldValue           func(context.Context) (*Record, error)
	predicates         []predicate.Record
}

var _ ent.Mutation = (*RecordMutation)(nil)

// recordOption allows management of the mutation configuration using functional options.
type recordOption func(*RecordMutation)

// newRecordMutation creates new mutation for the Record entity.
func newRecordMutation(c config, op Op, opts ...recordOption) *RecordMutation {
	m := &RecordMutation{
		config:        c,
		op:            op,
		typ:           TypeRecord,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRecordID sets the ID field of the mutation.
func withRecordID(id uuid.UUID) recordOption {
	return func(m *RecordMutation) {
		var (
			err   error
			once  sync.Once
			value *Record
		)
		m.oldValue = func(ctx context.Context) (*Record, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Record.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRecord sets the old Record of the mutation.
func withRecord(node *Record) recordOption {
	return func(m *RecordMutation) {
		m.oldValue = func(context.Context) (*Record, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RecordMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RecordMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Record entities.
func (m *RecordMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *RecordMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *RecordMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Record.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// AddDataFieldIDs adds the "data_fields" edge to the DataField entity by ids.
func (m *RecordMutation) AddDataFieldIDs(ids ...uuid.UUID) {
	if m.data_fields == nil {
		m.data_fields = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.data_fields[ids[i]] = struct{}{}
	}
}

// ClearDataFields clears the "data_fields" edge to the DataField entity.
func (m *RecordMutation) ClearDataFields() {
	m.cleareddata_fields = true
}

// DataFieldsCleared reports if the "data_fields" edge to the DataField entity was cleared.
func (m *RecordMutation) DataFieldsCleared() bool {
	return m.cleareddata_fields
}

// RemoveDataFieldIDs removes the "data_fields" edge to the DataField entity by IDs.
func (m *RecordMutation) RemoveDataFieldIDs(ids ...uuid.UUID) {
	if m.removeddata_fields == nil {
		m.removeddata_fields = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.data_fields, ids[i])
		m.removeddata_fields[ids[i]] = struct{}{}
	}
}

// RemovedDataFields returns the removed IDs of the "data_fields" edge to the DataField entity.
func (m *RecordMutation) RemovedDataFieldsIDs() (ids []uuid.UUID) {
	for id := range m.removeddata_fields {
		ids = append(ids, id)
	}
	return
}

// DataFieldsIDs returns the "data_fields" edge IDs in the mutation.
func (m *RecordMutation) DataFieldsIDs() (ids []uuid.UUID) {
	for id := range m.data_fields {
		ids = append(ids, id)
	}
	return
}

// ResetDataFields resets all changes to the "data_fields" edge.
func (m *RecordMutation) ResetDataFields() {
	m.data_fields = nil
	m.cleareddata_fields = false
	m.removeddata_fields = nil
}

// Where appends a list predicates to the RecordMutation builder.
func (m *RecordMutation) Where(ps ...predicate.Record) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *RecordMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Record).
func (m *RecordMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RecordMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RecordMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RecordMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown Record field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RecordMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Record field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RecordMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RecordMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RecordMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown Record numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RecordMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RecordMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RecordMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Record nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RecordMutation) ResetField(name string) error {
	return fmt.Errorf("unknown Record field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RecordMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.data_fields != nil {
		edges = append(edges, record.EdgeDataFields)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RecordMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case record.EdgeDataFields:
		ids := make([]ent.Value, 0, len(m.data_fields))
		for id := range m.data_fields {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RecordMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removeddata_fields != nil {
		edges = append(edges, record.EdgeDataFields)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RecordMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case record.EdgeDataFields:
		ids := make([]ent.Value, 0, len(m.removeddata_fields))
		for id := range m.removeddata_fields {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RecordMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareddata_fields {
		edges = append(edges, record.EdgeDataFields)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RecordMutation) EdgeCleared(name string) bool {
	switch name {
	case record.EdgeDataFields:
		return m.cleareddata_fields
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RecordMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Record unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RecordMutation) ResetEdge(name string) error {
	switch name {
	case record.EdgeDataFields:
		m.ResetDataFields()
		return nil
	}
	return fmt.Errorf("unknown Record edge %s", name)
}
