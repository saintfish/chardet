package chardet

import (
	"testing"
)

type chardetTester struct {
	d *Detector
}

func newChardetTester(r ...recognizer) *chardetTester {
	if len(r) == 0 {
		return &chardetTester{NewDetector()}
	}
	return &chardetTester{&Detector{r}}
}

func (this *chardetTester) ExpectBest(b []byte, charset string, lang string, t *testing.T) bool {
	r, err := this.d.DetectBest(b, true, "")
	if err != nil {
		t.Error(err)
		return false
	}
	if r.Charset != charset || r.Language != lang {
		t.Errorf("Expect %#v, actual %#v", Result{charset, lang, 0}, *r)
		return false
	}
	return true
}

func (this *chardetTester) ExpectUnknown(b []byte, t *testing.T) bool {
	r, err := this.d.DetectBest(b, true, "")
	if err == nil {
		t.Errorf("Expect unknown, actual %#v", *r)
		return false
	}
	return true
}
