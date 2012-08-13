package chardet

import (
	"testing"
)

var utf8Recognizers = []recognizer{
	new(recognizerUtf8),
}

func TestUtf8(t *testing.T) {
	ct := newChardetTester(new(recognizerUtf8))
	for name, content := range embeddedfiles {
		if name == "utf8.txt" {
			ct.ExpectBest(content, "UTF-8", "", t)
		} else {
			ct.ExpectUnknown(content, t)
		}
	}
}
