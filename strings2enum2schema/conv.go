package strings2enum2schema

import (
	se "github.com/takanoriyanagitani/go-avro-schema-strings2enum"
	. "github.com/takanoriyanagitani/go-avro-schema-strings2enum/util"
)

type StringsToSchema func([]string) IO[se.Schema]
