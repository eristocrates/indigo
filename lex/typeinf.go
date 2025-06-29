package lex

import "strings"

// InferPropertyType maps a TypeSchema describing a field to an OWL property type
// and range IRI.
func InferPropertyType(field *TypeSchema) (string, string) {
	switch field.Type {
	case "string":
		return "owl:DatatypeProperty", "xsd:string"
	case "boolean":
		return "owl:DatatypeProperty", "xsd:boolean"
	case "integer":
		return "owl:DatatypeProperty", "xsd:integer"
	case "ref":
		schemaID, defName := parseRef(field.id, field.Ref)
		return "owl:ObjectProperty", GetClassIRI(schemaID, defName)
	case "array":
		return InferPropertyType(field.Items)
	default:
		return "owl:ObjectProperty", ""
	}
}

func parseRef(parentSchemaID, ref string) (string, string) {
	if strings.HasPrefix(ref, "#") {
		return parentSchemaID, strings.TrimPrefix(ref, "#")
	}
	parts := strings.Split(ref, "#")
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return parentSchemaID, ref
}
