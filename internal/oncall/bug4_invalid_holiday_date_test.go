package oncall

import "testing"

func TestBug4_RejectsInvalidHolidayDate(t *testing.T) {
	_, err := Build(Request{
		Roster:   []string{"alice"},
		Start:    "2026-03-02",
		End:      "2026-03-02",
		Holidays: map[string]bool{"2026-02-30": true},
	})
	if err == nil {
		t.Fatal("expected invalid holiday date to be rejected")
	}
}
