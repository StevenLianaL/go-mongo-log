package helper

import "testing"

func TestCurrentMonth(t *testing.T) {
	month := CurrentMonth()
	if month != "2022-02" {
		t.Errorf("CurrentMonth() = %v, want %v", month, "2022-02")
	} else {
		t.Log("CurrentMonth() = ", month)
	}
}
