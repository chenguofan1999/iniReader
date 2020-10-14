package inireader_test

import (
	"testing"

	"github.com/chenguofan1999/inireader"
)

func TestFileListener(t *testing.T) {
	var fl inireader.FileListener
	fl.Listen("testData/my.ini")
}
