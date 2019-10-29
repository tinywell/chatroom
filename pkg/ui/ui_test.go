package ui

import (
	"testing"
)

func TestUI(t *testing.T) {
	ui, err := NewUI()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	ui.Open()
}

func TestMsg(t *testing.T) {
	ui, err := NewUI()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	ui.ApppendMsg("hello", "tinywell")
	ui.Open()

}
