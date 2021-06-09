package zj

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

var errWatch = errors.New(`Incorrect use of props`)

// Watch err if not nil
func Watch(err *error, prefix ...interface{}) {

	if err == nil {
		err = &errWatch
	} else if *err == nil {
		return
	}

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = `???`
	}
	_, file = filepath.Split(file)

	content := []interface{}{
		"\x1b[38;2;255;100;100m",
		time.Now().Format(`15:04:05.000 `),
		file,
		`:`,
		line,
		` `,
		getFrame(1).Function,
		`() `,
	}
	if len(prefix) > 0 {
		content = append(content, fmt.Sprint(prefix...))
		content = append(content, ` `)
	}
	content = append(content, (*err).Error())
	content = append(content, "\n\x1b[0m")

	warnLog.Raw(content...)
}

func getFrame(skipFrames int) (f runtime.Frame) {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	maxIdx := skipFrames + 2

	// Set size to maxIdx+2 to ensure we have room for one more caller than we need
	pc := make([]uintptr, maxIdx+2)
	n := runtime.Callers(0, pc)

	f = runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(pc[:n])
		for more, idx := true, 0; more && idx <= maxIdx; idx++ {
			var nf runtime.Frame
			nf, more = frames.Next()
			if idx == maxIdx {
				f = nf
			}
		}
	}

	return f
}
