package sigterm

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/melnikovio/go-common/pkg/common"
)

func HandleSigTerm() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-signalChannel
		switch sig {
		default:
			fmt.Println(common.GetSigTermMessage())
			os.Exit(0)
		}
	}()
}
