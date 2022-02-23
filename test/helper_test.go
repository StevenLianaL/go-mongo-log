package test

import (
	mongoLog "github.com/StevenlianaL/mongo-log"
	"testing"
)

func TestCurrentMonth(t *testing.T) {
	month := mongoLog.CurrentMonth()
	if month != "2022-02" {
		t.Errorf("CurrentMonth() = %v, want %v", month, "2022-02")
	} else {
		t.Log("CurrentMonth() = ", month)
	}
}

func TestGetFuncName(t *testing.T) {
	name := mongoLog.GetFuncName()
	if name != "tRunner" {
		t.Errorf("GetFuncName() = %v, want %v", name, "tRunner")
	} else {
		t.Log("GetFuncName() = ", name)
	}
}
