package slicetable

import (
	"fmt"
	"reflect"

	"github.com/dolthub/go-mysql-server/sql"
)

func schemaFromType[T any](source string) (ret sql.Schema) {
	t := reflect.TypeFor[T]()
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("not a struct type: %v", t))
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		ret = append(ret, &sql.Column{
			Name:   field.Name,
			Type:   goTypeToSQLType(field.Type),
			Source: source,
		})
	}

	return
}
