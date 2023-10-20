package slicetable

import (
	"github.com/dolthub/go-mysql-server/sql"
	"sync/atomic"
)

type Table[T any] struct {
	name   string
	ptr    atomic.Pointer[[]T]
	schema sql.Schema
}

func New[T any](name string) *Table[T] {
	//TODO get schema
	return &Table[T]{
		name: name,
	}
}

var _ sql.Table = new(Table[int])

func (*Table[T]) Collation() sql.CollationID {
	return sql.Collation_utf8_general_ci
}

func (t *Table[T]) Name() string {
	return t.name
}

func (t *Table[T]) PartitionRows(ctx *sql.Context, _ sql.Partition) (sql.RowIter, error) {
	ptr := t.ptr.Load()
	return &tableIter[T]{
		slice: *ptr,
	}, nil
}

func (t *Table[T]) Partitions(*sql.Context) (sql.PartitionIter, error) {
	return &partitionIter{
		n: 1,
	}, nil
}

func (t *Table[T]) Schema() sql.Schema {
	return t.schema
}

func (t *Table[T]) String() string {
	return t.name
}
