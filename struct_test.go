package inireader_test

import (
	"testing"

	"github.com/chenguofan1999/inireader"
)

func TestStruct1(t *testing.T) {
	sec1 := inireader.Sec{Name: "secName1", Map: map[string]string{}}
	sec2 := inireader.Sec{Name: "secName2", Map: map[string]string{}}

	cfgMap := make(map[string]*inireader.Sec)
	cfgMap["secName1"] = &sec1
	cfgMap["secName2"] = &sec2

	cfg := inireader.Cfg{Map: cfgMap}
	cfg.Section("secName1").Map["SayHello"] = "hello"
	cfg.Section("secName2").Map["SayHello"] = "hello"
	cfg.Section("secName1").Map["SaySorry"] = "sorry"
	if cfg.Section("secName1").Key("SayHello") != "hello" {
		t.Error("not right")
	}
	if cfg.Section("secName2").Key("SayHello") != "hello" {
		t.Error("not right")
	}
	if cfg.Section("secName1").Key("SaySorry") != "sorry" {
		t.Error("not right")
	}
}
