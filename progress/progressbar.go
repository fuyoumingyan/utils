package progress

import (
	"fmt"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"time"
)

func NewProgressbar(max int64, description string, show bool) *progressbar.ProgressBar {
	// show 表示当进度条走完后是否在终端上面显示 还是就消失了
	if show {
		return progressbar.NewOptions64(max,
			progressbar.OptionSetDescription(description),
			progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
			progressbar.OptionOnCompletion(func() {
				fmt.Fprint(ansi.NewAnsiStdout(), "\n")
			}),
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
	} else {
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

}
