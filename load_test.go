package inireader_test

import (
	"testing"

	"github.com/chenguofan1999/inireader"
)

func TestLoadDescription(t *testing.T) {
	cfg := inireader.Load("testData/my.ini")

	expected := []string{"Protocol (http or https)",
		"The http port  to use",
		"Redirect to correct domain if host header does not match domain\nPrevents DNS rebinding attacks",
		"possible values : production, development"}

	if cfg.Section("server").Descriptions["protocol"] != expected[0] {
		t.Error("Not right")
	}

	if cfg.Section("server").Descriptions["http_port"] != expected[1] {
		t.Error("Not right")
	}

	if cfg.Section("server").Descriptions["enforce_domain"] != expected[2] {
		t.Error("Not right")
	}

	if cfg.Section("").Descriptions["app_mode"] != expected[3] {
		t.Error("Not right")
	}
}

func TestSection(t *testing.T) {
	cfg := inireader.Load("testData/my.ini")
	if cfg.Section("paths").Key("data") != "/home/git/grafana" {
		t.Error("Not right")
	}
	if cfg.Section("paths").Key("source") != "/root" {
		t.Error("Not right")
	}
}
