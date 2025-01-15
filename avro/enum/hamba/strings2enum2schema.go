package conv

import (
	"context"

	ha "github.com/hamba/avro/v2"

	se "github.com/takanoriyanagitani/go-avro-schema-strings2enum"
	. "github.com/takanoriyanagitani/go-avro-schema-strings2enum/util"

	s2 "github.com/takanoriyanagitani/go-avro-schema-strings2enum/strings2enum2schema"
)

func StringsToEnumSchema(
	id se.SchemaId,
	s []string,
) (*ha.EnumSchema, error) {
	return ha.NewEnumSchema(
		id.Name,
		id.NameSpace,
		s,
	)
}

func EnumSchemaToString(e *ha.EnumSchema) string {
	return e.String()
}

func StringsToSchemaNew(id se.SchemaId) s2.StringsToSchema {
	return func(s []string) IO[se.Schema] {
		return func(_ context.Context) (se.Schema, error) {
			return ComposeErr(
				Curry(StringsToEnumSchema)(id),
				func(e *ha.EnumSchema) (se.Schema, error) {
					var s string = EnumSchemaToString(e)
					return se.Schema(s), nil
				},
			)(s)
		}
	}
}
