package chardet

import (
	"errors"
)

type Result struct {
	Charset    string
	Language   string
	Confidence int
}

type Detector struct {
	recognizers []recognizer
}

func NewDetector() *Detector {
    // Init recognizer
    return &Detector{}
}

var (
    NotDetectedError = errors.New("Charset not detected.")
)

func (d *Detector) DetectBest(b []byte, stripTag bool, declaredCharset string) (r *Result, err error) {
	var all []Result
	if all, err = d.DetectAll(b, stripTag, declaredCharset); err != nil {
		r = &all[0]
	}
	return
}

func (d *Detector) DetectAll(b []byte, stripTag bool, declaredCharset string) ([]Result, error) {
	_ = newRecognizerInput(b, stripTag, declaredCharset)
	return nil, NotDetectedError
}
