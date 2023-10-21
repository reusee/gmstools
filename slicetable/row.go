package slicetable

import (
	"reflect"

	"github.com/dolthub/go-mysql-server/sql"
)

func toSQLRow(obj any) (ret sql.Row) {
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		ret = append(ret, v.Field(i).Interface())
	}
	return
}
