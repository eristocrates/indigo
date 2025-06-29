package lex

import "fmt"

const baseIRI = "https://atproto.social/ontology/"

func GetClassIRI(schemaID, defName string) string {
	return fmt.Sprintf("%s%s#%s", baseIRI, schemaID, defName)
}

func GetPropertyIRI(classIRI, fieldName string) string {
	return fmt.Sprintf("%s/%s", classIRI, fieldName)
}
