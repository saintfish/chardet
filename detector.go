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
}

// List of charset recognizers
var recognizers = []recognizer {
    new(recognizerUtf8),
}

func NewDetector() *Detector {
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

func matchHelper(r recognizer, input *recognizerInput, outputChan chan<- recognizerOutput) {
    outputChan <- r.Match(input)
}

func (d *Detector) DetectAll(b []byte, stripTag bool, declaredCharset string) ([]Result, error) {
    input := newRecognizerInput(b, stripTag, declaredCharset)
    outputChan := make(chan recognizerOutput)
    for _, r := range recognizers {
        go matchHelper(r, input, outputChan)
    }
    outputs := make([]recognizerOutput, 0, len(recognizers))
    for i := 0; i < len(recognizers); i++ {
        outputs = append(outputs, <-outputChan)
    }
	return nil, NotDetectedError
}
