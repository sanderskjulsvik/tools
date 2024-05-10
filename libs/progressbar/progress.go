package common

import (
	uiprogress "github.com/gosuri/uiprogress"
	schollzProgressbar "github.com/schollz/progressbar/v3"
)

type ProgressBar interface {
	Add(x int)
	Add1()
	Start()
}

type SchollzProgressbar struct {
	bar *schollzProgressbar.ProgressBar
}

func NewSchollzProgressbar() *SchollzProgressbar {
	return &SchollzProgressbar{
		// -1 for infinite bar
		bar: schollzProgressbar.Default(-1, "Files processed:"),
	}
}

func (sp *SchollzProgressbar) Add1() {
	sp.bar.Add(1)
}

func (sp *SchollzProgressbar) Add(n int) {
	sp.bar.Add(n)
}

type UiProgressBars struct {
	bars []*UiProgressBar
}

type UiProgressBar struct {
	bar *uiprogress.Bar
}

// Only for use with one bar
func NewUiProgressBars() UiProgressBars {
	return UiProgressBars{
		bars: []*UiProgressBar{},
	}
}

// AddBar adds a bar to the progress bar instance and returns the bar index
func (uiP *UiProgressBars) AddBar(total int) *UiProgressBar {
	return &UiProgressBar{
		bar: uiprogress.AddBar(total).AppendCompleted().PrependElapsed(),
	}
}

func (uiP *UiProgressBars) Start() {
	uiprogress.Start()
}

func (uiP *UiProgressBar) Add(x int) {
	for i := 0; i < x; i++ {
		uiP.bar.Incr()
	}
}

func (uiP *UiProgressBar) Add1() {
	uiP.bar.Incr()
}
