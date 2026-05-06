package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 66, Capacity: 72, Latency: 27, Risk: 6, Weight: 10}, wantScore: 122, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 67, Capacity: 78, Latency: 8, Risk: 6, Weight: 6}, wantScore: 186, wantDecision: "accept"},
		{name: "case_3", signal: Signal{Demand: 67, Capacity: 98, Latency: 16, Risk: 19, Weight: 9}, wantScore: 137, wantDecision: "review"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
