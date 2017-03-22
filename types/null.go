package types

import (
	"github.com/peernova-private/gogen-avro/generator"
	"fmt"
)

const writeNullMethod = `
func writeNull(_ interface{}, _ io.Writer) error {
	return nil
}
`

const readNullMethod = `
func readNull(_ io.Reader) (interface{}, error) {
	return nil, nil
}
`

type nullField struct {
	name       string
	hasDefault bool
	tag        string
}

func (s *nullField) HasDefault() bool {
	return s.hasDefault
}

func (s *nullField) Default() interface{} {
	return nil
}

func (s *nullField) AvroName() string {
	return s.name
}

func (s *nullField) GoName() string {
	return generator.ToPublicName(s.name)
}

func (s *nullField) FieldType() string {
	return "Null"
}

func (s *nullField) GoType() string {
	return fmt.Sprintf("interface{}%s", s.Tag())
}

func (s *nullField) SerializerMethod() string {
	return "writeNull"
}

func (s *nullField) DeserializerMethod() string {
	return "readNull"
}

func (s *nullField) AddStruct(p *generator.Package) {}

func (s *nullField) AddSerializer(p *generator.Package) {
	p.AddFunction(UTIL_FILE, "", "writeNull", writeNullMethod)
	p.AddImport(UTIL_FILE, "io")
}

func (s *nullField) AddDeserializer(p *generator.Package) {
	p.AddFunction(UTIL_FILE, "", "readNull", readNullMethod)
	p.AddImport(UTIL_FILE, "io")
}

func (s *nullField) ResolveReferences(n *Namespace) error {
	return nil
}

func (s *nullField) Schema(names map[QualifiedName]interface{}) interface{} {
	return "null"
}

func (s *nullField) Tag() string {
	if len(s.tag) < 1 {
		return ""
	}
	return fmt.Sprintf(" `%s`", s.tag)
}
