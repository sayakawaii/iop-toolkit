package omciDeshape

import "testing"

const (
	inPath  = "./testdata/evtocdRule.input"
	outPath = "./testoutput/evtocdRule.output"
)

//go test -run EvtocdShape omciAnalyzer/service/omciDeshape -v
func TestEvtocdShape(t *testing.T) {
	if err := OmciDeshapeTrans(EVTOCD, inPath, outPath); err != nil {
		t.Error(err)
	}
}

//go test -benchmem -benchtime=10s -cpuprofile cpu_profile.out -memprofile mem_profile.out -bench EvtocdShape omciAnalyzer/service/omciDeshape -v
//go tool pprof -http="127.0.0.1:8080" cpu_profile.out
//go tool pprof -http="127.0.0.1:8081" mem_profile.out
func BenchmarkEvtocdShape(b *testing.B) {
	for n := 0; n < b.N; n++ {
		OmciDeshapeTrans(EVTOCD, inPath, outPath)
	}
}
