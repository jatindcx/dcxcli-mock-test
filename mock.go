package mockexecutor

import (
	"errors"
	"fmt"
	"time"
)

type MockOptions struct {
	Image   string
	Verbose bool
}

type ExecResult struct {
	StatusCode int
	Output     string
	Err        error
}

func PullImage(opts MockOptions) ExecResult {
	if opts.Image == "" {
		return ExecResult{
			StatusCode: 1,
			Output:     "",
			Err:        errors.New("image name is required"),
		}
	}

	if opts.Verbose {
		fmt.Println("Starting pull for image:", opts.Image)
	}

	// just simulating stuff than pulling it XD.
	steps := []string{
		"Connecting to registry...",
		"Fetching image metadata...",
		"Downloading layers...",
		"Extracting layers...",
		"Image pull complete.",
	}

	for _, step := range steps {
		if opts.Verbose {
			fmt.Println(step)
		}
		time.Sleep(500 * time.Millisecond)
	}

	return ExecResult{
		StatusCode: 0,
		Output:     fmt.Sprintf("Image %s pulled successfully", opts.Image),
		Err:        nil,
	}
}
