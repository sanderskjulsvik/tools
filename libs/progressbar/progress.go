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

type ProgressBarCollection interface {
	Start()
	Stop()
	// header, size
	AddBar(string, int) *ProgressBar
}

// ///////////////////////////////////
// Moc implementation
// ///////////////////////////////////

type ProgressBarCollectionMoc struct {
	bars []*ProgressBarMoc
}

func NewMocProgressBars() ProgressBarCollectionMoc {
	return ProgressBarCollectionMoc{
		bars: []*ProgressBarMoc{},
	}
}

func (pbs *ProgressBarCollectionMoc) AddBar(name string, total int) *ProgressBarMoc {
	newBar := ProgressBarMoc{}
	pbs.bars = append(pbs.bars, &newBar)
	return &newBar
}

func (pbs *ProgressBarCollectionMoc) Start() {
}

func (pbs *ProgressBarCollectionMoc) Stop() {
}

type ProgressBarMoc struct {
	bar *uiprogress.Bar
}

func (pb *ProgressBarMoc) Add(x int) {
}

func (pb *ProgressBarMoc) Add1() {
}

// ///////////////////////////////////
// UiProgressBar implementation
// ///////////////////////////////////

type UiPCollection struct {
	bars       []*UiProgressBar
	uiprogress *uiprogress.Progress
}

func NewUiProgressBars() UiPCollection {
	return UiPCollection{
		bars: []*UiProgressBar{},
	}
}

// AddBar adds a bar to the progress bar instance and returns the bar index
func (uiP *UiPCollection) AddBar(name string, total int) UiProgressBar {
	newBar := UiProgressBar{
		bar: uiP.uiprogress.AddBar(total).AppendCompleted().PrependElapsed().PrependFunc(func(b *uiprogress.Bar) string {
			return fmt.Sprintf("%s \n    (%d/%d)Mb\t", name, b.Current(), total)
		}),
	}
	uiP.bars = append(uiP.bars, &newBar)
	return newBar
}

func (uiP *UiPCollection) Start() {
	uiP.uiprogress.Start()
}

func (uiP *UiPCollection) Stop() {
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
