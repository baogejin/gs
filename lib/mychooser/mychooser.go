package mychooser

import "errors"

type MyChooser struct {
	choices []Choice
	chooser *Chooser
}

func NewMyChooser() *MyChooser {
	return &MyChooser{
		choices: make([]Choice, 0),
		chooser: nil,
	}
}

func (w *MyChooser) Add(v interface{}, weigt uint) {
	if w == nil {
		return
	}
	w.choices = append(w.choices, Choice{
		Item:   v,
		Weight: weigt,
	})
	w.chooser = nil
}

func (w *MyChooser) Pick() (interface{}, error) {
	if w == nil {
		return nil, errors.New("MyChooser is nil")
	}
	if len(w.choices) == 0 {
		return nil, nil
	}
	if w.chooser == nil {
		chooser, err := NewChooser(
			w.choices...,
		)
		if err != nil {
			return nil, err
		}
		w.chooser = chooser
	}
	return w.chooser.Pick(), nil
}
