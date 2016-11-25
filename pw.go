package pw

import (
	"fmt"
	"strings"

	"github.com/tiantour/fetch"
)

// Options pullword input options
type Options struct {
	Threshold int // threshold can be 0 or 1
	Debug     int // debug can be 0 or 1
}

// Cmd send data to api.pullword.com
func Cmd(method, source string, args Options) ([]byte, error) {
	url := fmt.Sprintf(
		"http://api.pullword.com/%s.php?source=%s&param1=%d&param2=%d",
		method,
		source,
		args.Threshold,
		args.Debug,
	)
	return fetch.Cmd(method, strings.ToLower(url))
}
