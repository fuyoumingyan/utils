package progress

import (
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"time"
)

func NewProgressbar(max int64, description string) *progressbar.ProgressBar {
	return progressbar.NewOptions64(max,
		progressbar.OptionSetDescription(description),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionShowIts(),
		progressbar.OptionShowCount(),
		progressbar.OptionFullWidth(),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
}
