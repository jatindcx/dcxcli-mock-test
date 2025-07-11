package mock

import (
	"errors"
	"fmt"
	"time"
	"go.uber.org/zap"
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

func PullImage(opts MockOptions, logger *zap.Logger) ExecResult {
	if opts.Image == "" {
		return ExecResult{
			StatusCode: 1,
			Output:     "",
			Err:        errors.New("image name is required"),
		}
	}

	if opts.Verbose {
		logger.Info("Starting pull for image:", zap.String("image", opts.Image))
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
			logger.Info(step)
		}
		time.Sleep(500 * time.Millisecond)
	}

	message := fmt.Sprintf("Image %s pulled successfully", opts.Image)

	return ExecResult{
		StatusCode: 0,
		Output:     message,
		Err:        nil,
	}
}
