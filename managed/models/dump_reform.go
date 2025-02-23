// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type dumpTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *dumpTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("dumps").
func (v *dumpTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *dumpTableType) Columns() []string {
	return []string{
		"id",
		"status",
		"service_names",
		"start_time",
		"end_time",
		"export_qan",
		"ignore_load",
		"created_at",
		"updated_at",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *dumpTableType) NewStruct() reform.Struct {
	return new(Dump)
}

// NewRecord makes a new record for that table.
func (v *dumpTableType) NewRecord() reform.Record {
	return new(Dump)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *dumpTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// DumpTable represents dumps view or table in SQL database.
var DumpTable = &dumpTableType{
	s: parse.StructInfo{
		Type:    "Dump",
		SQLName: "dumps",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "string", Column: "id"},
			{Name: "Status", Type: "DumpStatus", Column: "status"},
			{Name: "ServiceNames", Type: "pq.StringArray", Column: "service_names"},
			{Name: "StartTime", Type: "*time.Time", Column: "start_time"},
			{Name: "EndTime", Type: "*time.Time", Column: "end_time"},
			{Name: "ExportQAN", Type: "bool", Column: "export_qan"},
			{Name: "IgnoreLoad", Type: "bool", Column: "ignore_load"},
			{Name: "CreatedAt", Type: "time.Time", Column: "created_at"},
			{Name: "UpdatedAt", Type: "time.Time", Column: "updated_at"},
		},
		PKFieldIndex: 0,
	},
	z: new(Dump).Values(),
}

// String returns a string representation of this struct or record.
func (s Dump) String() string {
	res := make([]string, 9)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Status: " + reform.Inspect(s.Status, true)
	res[2] = "ServiceNames: " + reform.Inspect(s.ServiceNames, true)
	res[3] = "StartTime: " + reform.Inspect(s.StartTime, true)
	res[4] = "EndTime: " + reform.Inspect(s.EndTime, true)
	res[5] = "ExportQAN: " + reform.Inspect(s.ExportQAN, true)
	res[6] = "IgnoreLoad: " + reform.Inspect(s.IgnoreLoad, true)
	res[7] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[8] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Dump) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Status,
		s.ServiceNames,
		s.StartTime,
		s.EndTime,
		s.ExportQAN,
		s.IgnoreLoad,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Dump) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Status,
		&s.ServiceNames,
		&s.StartTime,
		&s.EndTime,
		&s.ExportQAN,
		&s.IgnoreLoad,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *Dump) View() reform.View {
	return DumpTable
}

// Table returns Table object for that record.
func (s *Dump) Table() reform.Table {
	return DumpTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Dump) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Dump) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Dump) HasPK() bool {
	return s.ID != DumpTable.z[DumpTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *Dump) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = DumpTable
	_ reform.Struct = (*Dump)(nil)
	_ reform.Table  = DumpTable
	_ reform.Record = (*Dump)(nil)
	_ fmt.Stringer  = (*Dump)(nil)
)

type dumpLogViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *dumpLogViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("dump_logs").
func (v *dumpLogViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *dumpLogViewType) Columns() []string {
	return []string{
		"dump_id",
		"chunk_id",
		"data",
		"last_chunk",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *dumpLogViewType) NewStruct() reform.Struct {
	return new(DumpLog)
}

// DumpLogView represents dump_logs view or table in SQL database.
var DumpLogView = &dumpLogViewType{
	s: parse.StructInfo{
		Type:    "DumpLog",
		SQLName: "dump_logs",
		Fields: []parse.FieldInfo{
			{Name: "DumpID", Type: "string", Column: "dump_id"},
			{Name: "ChunkID", Type: "uint32", Column: "chunk_id"},
			{Name: "Data", Type: "string", Column: "data"},
			{Name: "LastChunk", Type: "bool", Column: "last_chunk"},
		},
		PKFieldIndex: -1,
	},
	z: new(DumpLog).Values(),
}

// String returns a string representation of this struct or record.
func (s DumpLog) String() string {
	res := make([]string, 4)
	res[0] = "DumpID: " + reform.Inspect(s.DumpID, true)
	res[1] = "ChunkID: " + reform.Inspect(s.ChunkID, true)
	res[2] = "Data: " + reform.Inspect(s.Data, true)
	res[3] = "LastChunk: " + reform.Inspect(s.LastChunk, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *DumpLog) Values() []interface{} {
	return []interface{}{
		s.DumpID,
		s.ChunkID,
		s.Data,
		s.LastChunk,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *DumpLog) Pointers() []interface{} {
	return []interface{}{
		&s.DumpID,
		&s.ChunkID,
		&s.Data,
		&s.LastChunk,
	}
}

// View returns View object for that struct.
func (s *DumpLog) View() reform.View {
	return DumpLogView
}

// check interfaces
var (
	_ reform.View   = DumpLogView
	_ reform.Struct = (*DumpLog)(nil)
	_ fmt.Stringer  = (*DumpLog)(nil)
)

func init() {
	parse.AssertUpToDate(&DumpTable.s, new(Dump))
	parse.AssertUpToDate(&DumpLogView.s, new(DumpLog))
}
