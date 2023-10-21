package slicetable

import (
	"io"

	"github.com/dolthub/go-mysql-server/sql"
)

type tableIter[T any] struct {
	slice []T
}

var _ sql.RowIter = new(tableIter[int])

func (i *tableIter[T]) Close(*sql.Context) error {
	i.slice = nil
	return nil
}

func (t *tableIter[T]) Next(ctx *sql.Context) (sql.Row, error) {
	if len(t.slice) == 0 {
		return nil, io.EOF
	}
	elem := t.slice[0]
	t.slice = t.slice[1:]
	return toSQLRow(elem), nil
}
