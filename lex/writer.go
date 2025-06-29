package lex

import (
	"bufio"
	"os"
)

type triple struct {
	s string
	p string
	o string
}

type RdfWriter struct {
	triples []triple
}

func (w *RdfWriter) WriteTriple(s, p, o string) {
	w.triples = append(w.triples, triple{s: s, p: p, o: o})
}

func (w *RdfWriter) WriteTurtle(outPath string) error {
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()
	bw := bufio.NewWriter(f)

	// standard prefixes
	bw.WriteString("@prefix owl: <http://www.w3.org/2002/07/owl#> .\n")
	bw.WriteString("@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .\n")
	bw.WriteString("@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .\n\n")

	for _, t := range w.triples {
		bw.WriteString(t.s + " " + t.p + " " + t.o + " .\n")
	}
	return bw.Flush()
}
