package metrics

import (
	"testing"
	"time"
)

//Do not use built-in lable-run test to launch the test, otherwise you will get a timeout(timeout of testing is set as 30s by default)
//Use "go test ./metrics  -timeout 60s -v" to do the test
//Use below cmds to check further with the output log
//1.go tool pprof omcianalyzer/metrics/tmp/cpu..20220905171017.540.log
//2.top
func TestMetrics(t *testing.T) {
	h, _ := NewHolmes()
	h.Start()
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(1 * time.Minute)
	h.Stop()
}
