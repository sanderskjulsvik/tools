package progressbar

import (
	"fmt"

	uiprogress "github.com/gosuri/uiprogress"
)

// ///////////////////////////////////
// The interfaces
// ///////////////////////////////////

type ProgressBar interface {
	Add(x int)
	Add1()
}

type ProgressBars interface {
	Start()
	Stop()
	AddBar(int) *UiProgressBar
}

// ///////////////////////////////////
// UiProgressBar implementation
// ///////////////////////////////////

type UiProgressBars struct {
	bars       []*UiProgressBar
	uiprogress *uiprogress.Progress
}

func NewUiProgressBars() UiProgressBars {
	return UiProgressBars{
		bars: []*UiProgressBar{},
	}
}

// AddBar adds a bar to the progress bar instance and returns the bar index
func (uiP *UiProgressBars) AddBar(name string, total int) *UiProgressBar {
	newBar := UiProgressBar{
		bar: uiP.uiprogress.AddBar(total).AppendCompleted().PrependElapsed().PrependFunc(func(b *uiprogress.Bar) string {
			return fmt.Sprintf("%s \n    (%d/%d)Mb\t", name, b.Current(), total)
		}),
	}
	uiP.bars = append(uiP.bars, &newBar)
	return &newBar
}

func (uiP *UiProgressBars) Start() {
	uiP.uiprogress.Start()
}

func (uiP *UiProgressBars) Stop() {
	uiP.uiprogress.Stop()
}

type UiProgressBar struct {
	bar *uiprogress.Bar
}

func (uiP *UiProgressBar) Add(x int) {
	for i := 0; i < x; i++ {
		uiP.bar.Incr()
	}
}

func (uiP *UiProgressBar) Add1() {
	uiP.bar.Incr()
}
