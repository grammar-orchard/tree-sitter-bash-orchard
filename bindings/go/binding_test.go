package tree_sitter_bash_orchard_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_bash_orchard "codeberg.org/grammar-orchard/tree-sitter-bash-orchard/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_bash_orchard.Language())
	if language == nil {
		t.Errorf("Error loading Bash grammar")
	}
}
