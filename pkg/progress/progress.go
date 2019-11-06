package progress

import "github.com/gosuri/uiprogress"

type ProgressBar struct {
	bar *uiprogress.Bar
}

func NewProgressBar(size int) *ProgressBar {
	uiprogress.Start()

	bar := uiprogress.AddBar(size)
	bar.AppendCompleted()
	bar.PrependElapsed()

	return &ProgressBar{bar: bar}
}

func (pb *ProgressBar) Increase() {
	pb.bar.Incr()
}
