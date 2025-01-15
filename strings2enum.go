package strings2enum

type SchemaId struct {
	Name      string
	NameSpace string
}

type Schema string

const NamespaceDefault string = ""

type SchemaName string

func (s SchemaName) ToSchemaId(namespace string) SchemaId {
	return SchemaId{
		Name:      string(s),
		NameSpace: namespace,
	}
}

func (s SchemaName) ToSchemaIdDefault() SchemaId {
	return s.ToSchemaId(NamespaceDefault)
}

type Codec string

const (
	CodecNull    Codec = "null"
	CodecDeflate Codec = "deflate"
	CodecSnappy  Codec = "snappy"
	CodecZstd    Codec = "zstandard"
	CodecBzip2   Codec = "bzip2"
	CodecXz      Codec = "xz"
)

const BlockLengthDefault int = 100

type EncodeConfig struct {
	BlockLength int
	Codec
}

var EncodeConfigDefault EncodeConfig = EncodeConfig{
	BlockLength: BlockLengthDefault,
	Codec:       CodecNull,
}
