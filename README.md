# Multi Progress Bar

[![GoDoc](https://godoc.org/github.com/vbauerster/mpb?status.svg)](https://godoc.org/github.com/vbauerster/mpb)
[![Build Status](https://travis-ci.org/vbauerster/mpb.svg?branch=master)](https://travis-ci.org/vbauerster/mpb)
[![Go Report Card](https://goreportcard.com/badge/github.com/vbauerster/mpb)](https://goreportcard.com/report/github.com/vbauerster/mpb)
[![cover.run go](https://cover.run/go/github.com/vbauerster/mpb.svg)](https://cover.run/go/github.com/vbauerster/mpb)

**mpb** is a Go lib for rendering progress bars in terminal applications.

It is inspired by [uiprogress](https://github.com/gosuri/uiprogress) library,
but unlike the last one, implementation is mutex free, following Go's idiom:

> Don't communicate by sharing memory, share memory by communicating.

## Features

* __Multiple Bars__: mpb can render multiple progress bars that can be tracked concurrently
* __Cancellable__: cancel rendering goroutine at any time
* __Dynamic Addition__:  Add additional progress bar at any time
* __Dynamic Removal__:  Remove rendering progress bar at any time
* __Dynamic Sorting__:  Sort bars as you wish
* __Dynamic Resize__:  Resize bars on terminal width change
* __Custom Decorator Functions__: Add custom functions around the bar along with helper functions
* __Dynamic Decorator's Width Sync__:  Sync width among decorator group (available since v2)
* __Predefined Decoratros__: Elapsed time, [Ewmaest](https://github.com/dgryski/trifles/tree/master/ewmaest) based ETA, Percentage, Bytes counter

## Installation

To get the package, execute:

```sh
go get gopkg.in/vbauerster/mpb.v3
```

## Usage

Following is the simplest use case:

```go
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
```

Running [this](examples/singleBar/main.go), will produce:

![gif](examples/gifs/single.gif)

However **mpb** was designed with concurrency in mind. Each new bar renders in its
own goroutine, therefore adding multiple bars is easy and safe:

```go
	p := mpb.New()
	total := 100
	numBars := 3
	var wg sync.WaitGroup
	wg.Add(numBars)

	for i := 0; i < numBars; i++ {
		name := fmt.Sprintf("Bar#%d:", i)
		bar := p.AddBar(int64(total),
			mpb.PrependDecorators(
				decor.Name(name, len(name), 0),
				decor.Percentage(3, decor.DSyncSpace),
			),
			mpb.AppendDecorators(
				decor.ETA(2, 0),
			),
		)
		go func() {
			defer wg.Done()
			for i := 0; i < total; i++ {
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				bar.Incr(1)
			}
		}()
	}
	wg.Wait() // Wait for goroutines to finish
	p.Stop()  // Stop mpb's rendering goroutine
```

![simple.gif](examples/gifs/simple.gif)

The source code: [examples/simple/main.go](examples/simple/main.go)

### Cancel

![cancel.gif](examples/gifs/cancel.gif)

The source code: [examples/cancel/main.go](examples/cancel/main.go)

### Removing bar

![remove.gif](examples/gifs/remove.gif)

The source code: [examples/remove/main.go](examples/remove/main.go)

### Sorting bars by progress

![sort.gif](examples/gifs/sort.gif)

The source code: [examples/sort/main.go](examples/sort/main.go)

### Resizing bars on terminal width change

![resize.gif](examples/gifs/resize.gif)

The source code: [examples/prependETA/main.go](examples/prependETA/main.go)

### Multiple io

![io-multiple.gif](examples/gifs/io-multiple.gif)

The source code: [examples/io/multiple/main.go](examples/io/multiple/main.go)

## License

[BSD 3-Clause](https://opensource.org/licenses/BSD-3-Clause)

The typeface used in screen shots: [Iosevka](https://be5invis.github.io/Iosevka)
