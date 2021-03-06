package main

import (
	"math/rand"
	"time"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

func main() {
	p := mpb.New(
		// override default (80) width
		mpb.WithWidth(100),
		// override default "[=>-]" format
		mpb.WithFormat("╢▌▌░╟"),
		// override default 100ms refresh rate
		mpb.WithRefreshRate(120*time.Millisecond),
	)

	total := 100
	name := "Single Bar:"
	// Add a bar
	// You're not limited to just a single bar, add as many as you need
	bar := p.AddBar(int64(total),
		// Prepending decorators
		mpb.PrependDecorators(
			// Name decorator with minWidth and no width sync options
			decor.Name(name, len(name), 0),
			// ETA decorator with minWidth and width sync options
			// DSyncSpace is shortcut for DwidthSync|DextraSpace
			decor.ETA(4, decor.DSyncSpace),
		),
		// Appending decorators
		mpb.AppendDecorators(
			// Percentage decorator with minWidth and no width sync options
			decor.Percentage(5, 0),
		),
	)

	for i := 0; i < total; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		bar.Incr(1) // increment progress bar
	}

	p.Stop()
}
