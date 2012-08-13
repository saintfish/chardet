package chardet

import (
	"errors"
	"sort"
)

type Result struct {
	Charset    string
	Language   string
	Confidence int
}

type Detector struct {
	recognizers []recognizer
}

// List of charset recognizers
var recognizers = []recognizer{
	new(recognizerUtf8),
	new(recognizerUtf16be),
	new(recognizerUtf16le),
	newRecognizerUtf32be(),
	newRecognizerUtf32le(),
}

func NewDetector() *Detector {
	return &Detector{recognizers}
}

var (
	NotDetectedError = errors.New("Charset not detected.")
)

func (d *Detector) DetectBest(b []byte, stripTag bool, declaredCharset string) (r *Result, err error) {
	var all []Result
	if all, err = d.DetectAll(b, stripTag, declaredCharset); err == nil {
		r = &all[0]
	}
	return
}

func (d *Detector) DetectAll(b []byte, stripTag bool, declaredCharset string) ([]Result, error) {
	input := newRecognizerInput(b, stripTag, declaredCharset)
	outputChan := make(chan recognizerOutput)
	for _, r := range d.recognizers {
		go matchHelper(r, input, outputChan)
	}
	outputs := make([]recognizerOutput, 0, len(d.recognizers))
	for i := 0; i < len(d.recognizers); i++ {
		o := <-outputChan
		if o.Confidence > 0 {
			outputs = append(outputs, o)
		}
	}
	if len(outputs) == 0 {
		return nil, NotDetectedError
	}

	sort.Sort(recognizerOutputs(outputs))
	dedupOutputs := make([]Result, 0, len(outputs))
	foundCharsets := make(map[string]struct{}, len(outputs))
	for _, o := range outputs {
		if _, found := foundCharsets[o.Charset]; !found {
			dedupOutputs = append(dedupOutputs, Result(o))
			foundCharsets[o.Charset] = struct{}{}
		}
	}
	if len(dedupOutputs) == 0 {
		return nil, NotDetectedError
	}
	return dedupOutputs, nil
}

func matchHelper(r recognizer, input *recognizerInput, outputChan chan<- recognizerOutput) {
	outputChan <- r.Match(input)
}

type recognizerOutputs []recognizerOutput

func (r recognizerOutputs) Len() int           { return len(r) }
func (r recognizerOutputs) Less(i, j int) bool { return r[i].Confidence > r[j].Confidence }
func (r recognizerOutputs) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
