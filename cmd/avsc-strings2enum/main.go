package main

import (
	"context"
	"fmt"
	"log"
	"os"

	se "github.com/takanoriyanagitani/go-avro-schema-strings2enum"
	en "github.com/takanoriyanagitani/go-avro-schema-strings2enum/avro/enum/hamba"
	src "github.com/takanoriyanagitani/go-avro-schema-strings2enum/source"
	s2 "github.com/takanoriyanagitani/go-avro-schema-strings2enum/strings2enum2schema"
	. "github.com/takanoriyanagitani/go-avro-schema-strings2enum/util"
)

var EnvValByKey func(string) IO[string] = Lift(
	func(key string) (string, error) {
		val, found := os.LookupEnv(key)
		switch found {
		case true:
			return val, nil
		default:
			return "", fmt.Errorf("env var %s missing", key)
		}
	},
)

var stringsSource IO[[]string] = src.StdinSource

var schemaName IO[string] = EnvValByKey("ENV_SCHEMA_NAME").
	Or(Of("enum_schema"))

var namespace IO[string] = EnvValByKey("ENV_SCHEMA_NAME_SPACE").
	Or(Of(""))

var schemaId IO[se.SchemaId] = Bind(
	All(
		schemaName,
		namespace,
	),
	Lift(func(s []string) (se.SchemaId, error) {
		var n se.SchemaName = se.SchemaName(s[0])
		return n.ToSchemaId(s[1]), nil
	}),
)

var strings2schema IO[s2.StringsToSchema] = Bind(
	schemaId,
	Lift(func(s se.SchemaId) (s2.StringsToSchema, error) {
		return en.StringsToSchemaNew(s), nil
	}),
)

var schema IO[se.Schema] = Bind(
	strings2schema,
	func(s2s s2.StringsToSchema) IO[se.Schema] {
		return Bind(
			stringsSource,
			s2s,
		)
	},
)

func schema2stdout(s se.Schema) IO[Void] {
	return func(_ context.Context) (Void, error) {
		fmt.Printf("%s", s)
		return Empty, nil
	}
}

var stdin2strings2schema2stdout IO[Void] = Bind(
	schema,
	schema2stdout,
)

var sub IO[Void] = func(ctx context.Context) (Void, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	return stdin2strings2schema2stdout(ctx)
}

func main() {
	_, e := sub(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
