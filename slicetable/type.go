package slicetable

import (
	"fmt"
	"reflect"
	"time"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/types"
)

var (
	timeType       = reflect.TypeFor[time.Time]()
	bytesSliceType = reflect.TypeFor[[]byte]()
)

func goTypeToSQLType(t reflect.Type) sql.Type {
	switch t {

	case timeType:
		return types.Time

	case bytesSliceType:
		return types.Blob

	}

	switch t.Kind() {

	case reflect.Int:
		return types.Int64
	case reflect.Int8:
		return types.Int8
	case reflect.Int16:
		return types.Int16
	case reflect.Int32:
		return types.Int32
	case reflect.Int64:
		return types.Int64

	case reflect.Uint:
		return types.Uint64
	case reflect.Uint8:
		return types.Uint8
	case reflect.Uint16:
		return types.Uint16
	case reflect.Uint32:
		return types.Uint32
	case reflect.Uint64:
		return types.Uint64

	case reflect.Float32:
		return types.Float32
	case reflect.Float64:
		return types.Float64

	case reflect.String:
		return types.Text

	case reflect.Struct:
		return types.JSON

	default:
		panic(fmt.Sprintf("unknown type: %v", t))
	}
}
