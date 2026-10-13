package main

import "testing"

func TestPlanMatterEmails(t *testing.T) {
	tests := []struct {
		name, matter string
		want         int
	}{
		{"signed matter", "M-42", 3},
		{"another matter", "M-77", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := planMatterEmails(matterIntake{MatterID: tt.matter}).Subjects
			if len(got) != tt.want || got[0] != "Matter "+tt.matter+" intake receipt" {
				t.Fatalf("plan = %#v", got)
			}
		})
	}
}
