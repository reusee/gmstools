package slicetable

import (
	"encoding/binary"
	"io"

	"github.com/dolthub/go-mysql-server/sql"
)

type partitionIter struct {
	n int64
}

var _ sql.PartitionIter = new(partitionIter)

func (p *partitionIter) Close(*sql.Context) error {
	p.n = 0
	return nil
}

func (p *partitionIter) Next(*sql.Context) (ret sql.Partition, err error) {
	if p.n <= 0 {
		return nil, io.EOF
	}
	ret = partitionID(binary.AppendVarint(nil, p.n))
	p.n -= 1
	return
}

type partitionID []byte

var _ sql.Partition = partitionID{}

func (p partitionID) Key() []byte {
	return p
}
