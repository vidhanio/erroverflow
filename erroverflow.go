package erroverflow

import (
	"fmt"
	"os/exec"
	"runtime"
)

func construct(err error) (url string) {
	return fmt.Sprintf("https://stackoverflow.com/search?q=[go] %s", err.Error())
}

func Handle(err error) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}

	args = append(args, construct(err))
	exec.Command(cmd, args...).Start()
}
