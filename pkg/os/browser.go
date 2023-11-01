/**
 * Author: Andrew Shoell
 * File: browser.go
 */

package os

import (
	"fmt"
	"os/exec"
	"runtime"

	flow "github.com/mrlunchbox777/basiclib/pkg/flow"
)

var myExec = exec.Command
var myRuntime = runtime.GOOS

func OpenBrowser(url string) {
	var err error
	switch myRuntime {
	case "linux":
		err = myExec("xdg-open", url).Start()
	case "windows":
		err = myExec("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = myExec("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	flow.MustNot(err)
}
