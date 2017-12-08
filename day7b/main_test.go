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

var nodes = make(map[string]node)

func init() {
	nodes["pbga"] = node{"pbga", 66, []string{}}
	nodes["xhth"] = node{"xhth", 57, []string{}}
	nodes["ebii"] = node{"ebii", 61, []string{}}
	nodes["havc"] = node{"havc", 66, []string{}}
	nodes["ktlj"] = node{"ktlj", 57, []string{}}
	nodes["fwft"] = node{"fwft", 72, []string{"ktlj", "cntj", "xhth"}}
	nodes["qoyq"] = node{"qoyq", 66, []string{}}
	nodes["padx"] = node{"padx", 45, []string{"pbga", "havc", "qoyq"}}
	nodes["tknk"] = node{"tknk", 41, []string{"ugml", "padx", "fwft"}}
	nodes["jptl"] = node{"jptl", 61, []string{}}
	nodes["ugml"] = node{"ugml", 68, []string{"gyxo", "ebii", "jptl"}}
	nodes["gyxo"] = node{"gyxo", 61, []string{}}
	nodes["cntj"] = node{"cntj", 57, []string{}}
}

func TestWeigh(t *testing.T) {
	tt := []struct {
		name     string
		expected int
	}{
		{"pbga", 66},
		{"xhth", 57},
		{"fwft", 72 + 57 + 57 + 57},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getWeight(tc.name, nodes)
			if err != nil {
				t.Fatalf("Unexpected error %v", err)
			}
			if actual != tc.expected {
				t.Errorf("Expected weight to be %d, got %d", tc.expected, actual)
			}
		})
	}
}

func TestUnbalanced(t *testing.T) {
	tt := []struct {
		name     string
		expected int
	}{
		{"pbga", 66},
		{"xhth", 57},
		{"fwft", 72 + 57 + 57 + 57},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getWeight(tc.name, nodes)
			if err != nil {
				t.Fatalf("Unexpected error %v", err)
			}
			if actual != tc.expected {
				t.Errorf("Expected weight to be %d, got %d", tc.expected, actual)
			}
		})
	}

	// Now we expect the total balance to fail
	_, err := getWeight("tknk", nodes)
	if err == nil {
		t.Errorf("Expected error, got none")
	}
	switch v := err.(type) {
	case unbalancedError:
		t.Logf("Got lengths %v", v)
	default:
		t.Fatalf("Unknown error %v", err)
	}
}
