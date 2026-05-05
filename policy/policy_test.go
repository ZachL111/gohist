package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 66, Capacity: 72, Latency: 27, Risk: 6, Weight: 10}
	if got := Score(signal); got != 122 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 67, Capacity: 78, Latency: 8, Risk: 6, Weight: 6}
	if got := Score(signal); got != 186 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 67, Capacity: 98, Latency: 16, Risk: 19, Weight: 9}
	if got := Score(signal); got != 137 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
