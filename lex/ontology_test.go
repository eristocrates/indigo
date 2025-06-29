package lex

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEmitOntology(t *testing.T) {
	s, err := ReadSchema(filepath.Join("..", "testdata", "ActorDefs_ViewerState.json"))
	if err != nil {
		t.Fatal(err)
	}
	tmp := t.TempDir()
	out := filepath.Join(tmp, "lexicon.ttl")
	if err := EmitOntology([]*Schema{s}, out); err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	ttl := string(b)
	if !strings.Contains(ttl, "viewerState") {
		t.Fatalf("expected class in ttl: %s", ttl)
	}
	if !strings.Contains(ttl, "blockedBy") {
		t.Fatalf("expected property in ttl: %s", ttl)
	}
}
