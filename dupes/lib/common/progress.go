package common

import (
	schollzProgressbar "github.com/schollz/progressbar/v3"
)

type ProgressBar interface {
	Add(x int)
	Add1()
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
