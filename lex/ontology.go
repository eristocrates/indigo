package lex

import (
	"fmt"
	"strings"
)

// EmitOntology exports schemas as OWL ontology in Turtle format.
func EmitOntology(schemas []*Schema, outPath string) error {
	w := &RdfWriter{}

	for _, s := range schemas {
		for defName, def := range s.Defs {
			classIRI := fmt.Sprintf("<%s>", GetClassIRI(s.ID, defName))
			w.WriteTriple(classIRI, "<http://www.w3.org/1999/02/22-rdf-syntax-ns#type>", "<http://www.w3.org/2002/07/owl#Class>")
			if def.Description != "" {
				lit := fmt.Sprintf("\"%s\"", escape(def.Description))
				w.WriteTriple(classIRI, "<http://www.w3.org/2000/01/rdf-schema#comment>", lit)
			}
			if def.Type == "object" {
				for fieldName, field := range def.Properties {
					owlType, rng := InferPropertyType(field)
					propIRI := fmt.Sprintf("<%s>", GetPropertyIRI(strings.Trim(classIRI, "<>"), fieldName))
					w.WriteTriple(propIRI, "<http://www.w3.org/1999/02/22-rdf-syntax-ns#type>", fmt.Sprintf("<http://www.w3.org/2002/07/owl#%s>", strings.TrimPrefix(owlType, "owl:")))
					w.WriteTriple(propIRI, "<http://www.w3.org/2000/01/rdf-schema#domain>", classIRI)
					if rng != "" {
						w.WriteTriple(propIRI, "<http://www.w3.org/2000/01/rdf-schema#range>", fmt.Sprintf("<%s>", rng))
					}
					if field.Description != "" {
						lit := fmt.Sprintf("\"%s\"", escape(field.Description))
						w.WriteTriple(propIRI, "<http://www.w3.org/2000/01/rdf-schema#comment>", lit)
					}
				}
			}
		}
	}

	return w.WriteTurtle(outPath)
}

func escape(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return s
}
