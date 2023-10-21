package slicetable

import (
	"errors"
	"io"
	"testing"

	"github.com/dolthub/go-mysql-server/sql"
)

func TestTable(t *testing.T) {
	type Foo struct {
		A int
		B string
	}

	table := New[Foo]("foo")
	slice := []Foo{
		{A: 1, B: "a"},
		{A: 2, B: "b"},
		{A: 3, B: "c"},
	}
	table.Update(&slice)

	n := 0
	forEachRow(t, table, func(_ sql.Row) {
		n++
	})
	if n != len(slice) {
		t.Fatal()
	}
}

func forEachRow[T any](
	t *testing.T,
	table *Table[T],
	fn func(sql.Row),
) {
	ctx := new(sql.Context)

	partIter, err := table.Partitions(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for {
		part, err := partIter.Next(ctx)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			t.Fatal(err)
		}

		rowIter, err := table.PartitionRows(ctx, part)
		if err != nil {
			t.Fatal(err)
		}
		for {
			row, err := rowIter.Next(ctx)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				t.Fatal(err)
			}
			fn(row)
		}
		if err := rowIter.Close(ctx); err != nil {
			t.Fatal(err)
		}

	}
	if err := partIter.Close(ctx); err != nil {
		t.Fatal(err)
	}
}
