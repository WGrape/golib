package safego

import (
	"fmt"
	"testing"
	"time"
)

// TestListenServerStop test the function of ListenServerStop, you can use control+c to stop after running the command : cd safego && go test -v .
func TestListenServerStop(t *testing.T) {
	var isStopped = false
	go func() {
		for {
			if isStopped {
				return
			}
			fmt.Printf("Wait ... %s\n", time.Now().Format("2006-01-02 15:04:05 "))
			time.Sleep(1 * time.Second)
		}
	}()
	<-ListenServerStop(&isStopped)
	fmt.Printf("\n\nEnd ...\n")
}
