package test

import (
	"mongo-log"
	"testing"
)

func TestCurrentMonth(t *testing.T) {
	month := mongo_log.CurrentMonth()
	if month != "2022-02" {
		t.Errorf("CurrentMonth() = %v, want %v", month, "2022-02")
	} else {
		t.Log("CurrentMonth() = ", month)
	}
}

func TestGetFuncName(t *testing.T) {
	name := mongo_log.GetFuncName()
	if name != "tRunner" {
		t.Errorf("GetFuncName() = %v, want %v", name, "tRunner")
	} else {
		t.Log("GetFuncName() = ", name)
	}
}
