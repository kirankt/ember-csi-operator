// Signals is a re-implementation of controller-runtime's signals
// implemantation. We add the ability to check the Oprator's
// job completion status before we can exit.
package embercsi

import (
	"github.com/golang/glog"
	"os"
	"os/signal"
	"syscall"
)

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}
var onlyOneSignalHandler = make(chan struct{})
var isDone = make(chan bool)

func SetupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		select {
		case <-c:
			glog.Error("Program received second TERM signal. Exiting Immediately\n")
			os.Exit(1)
		case <-isDone:
		}
	}()
	return stop
}
