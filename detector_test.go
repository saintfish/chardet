package chardet

import (
	"testing"
)

func TestDetector(t *testing.T) {
	type file_charset_language struct {
		File, Charset, Language string
	}
	var data = []file_charset_language{
		{"utf8.txt", "UTF-8", ""},
		{"big5.txt", "Big5", "zh"},
		{"shift_jis.txt", "Shift_JIS", "ja"},
		{"gb18030.txt", "GB-18030", "zh"},
	}

	ct := newChardetTester()
	for _, d := range data {
		ct.ExpectBest(embeddedfiles[d.File], d.Charset, d.Language, t)
	}
}
