package main

import (
	"strings"
	"testing"
)

func TestOne(t *testing.T) {
	tt := []struct {
		input     string
		xident    string
		xweight   int
		xchildren []string
	}{
		{"pbga (66)", "pbga", 66, []string{}},
		{"fwft (72) -> ktlj, cntj, xhth", "fwft", 72, []string{"ktlj", "cntj", "xhth"}},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			n := parseNode(tc.input)
			if n.ident != tc.xident {
				t.Errorf("Expected node %s, got %s", tc.xident, n.ident)
			}
			if n.weight != tc.xweight {
				t.Errorf("Expected weight %d, got %d", tc.xweight, n.weight)
			}
			if len(n.children) != len(tc.xchildren) {
				t.Errorf("Expected len(children) to be %d, got %d", len(tc.xchildren), len(n.children))
			}
			for k, v := range tc.xchildren {
				if n.children[k] != v {
					t.Errorf("Expected children[%d] to be %s, got %s", k, v, n.children[k])
				}
			}
		})
	}
}

func TestRead(t *testing.T) {
	tc := `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

	nodes := parseReader(strings.NewReader(tc))
	if len(nodes) != 13 {
		t.Errorf("Expected 13 nodes, got %d", len(nodes))
	}
	if nodes[5].ident != "fwft" {
		t.Errorf("Expected nodes[5].ident to be fwft, got %s", nodes[5].ident)
	}
}

func TestRoots(t *testing.T) {
	tc := `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`
	expected := "tknk"
	nodes := parseReader(strings.NewReader(tc))
	actual := getRoots(nodes)
	if len(actual) != 1 {
		t.Errorf("Expected 1 root, got %d", len(actual))
	}
	if actual[0] != expected {
		t.Errorf("Expected root to be %s, got %s", expected, actual)
	}
}
