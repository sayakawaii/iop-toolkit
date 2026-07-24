package metrics

import (
	"errors"
	"omciAnalyzer/env"
	"os"
	"time"

	"mosn.io/holmes"
)

var (
	DumpSetfuncArr1 = [env.SoloConfig1End]func(int, int, int, time.Duration) holmes.Option{
		holmes.WithMemDump,
		holmes.WithCPUDump,
		holmes.WithThreadDump,
		holmes.WithGCHeapDump,
	}
	DumpSetfuncArr2 = [env.SoloConfig2End - env.SoloConfig1End]func(int, int, int, int, time.Duration) holmes.Option{
		holmes.WithGoroutineDump,
	}

	DumpEnableFuncArr = [env.SoloConfigEnd]func(*holmes.Holmes) *holmes.Holmes{
		HolmesEnableMemDump,
		HolmesEnableCpuDump,
		HolmesEnableThreadDump,
		HolmesEnableGCHeapDump,
		HolmesEnableGoroutineDump,
	}
)

func HolmesEnableMemDump(h *holmes.Holmes) *holmes.Holmes {
	return h.EnableMemDump()
}

func HolmesEnableCpuDump(h *holmes.Holmes) *holmes.Holmes {
	return h.EnableCPUDump()
}

func HolmesEnableThreadDump(h *holmes.Holmes) *holmes.Holmes {
	return h.EnableThreadDump()
}

func HolmesEnableGCHeapDump(h *holmes.Holmes) *holmes.Holmes {
	return h.EnableGCHeapDump()
}

func HolmesEnableGoroutineDump(h *holmes.Holmes) *holmes.Holmes {
	return h.EnableGoroutineDump()
}

func NewHolmes() (*holmes.Holmes, error) {
	envConfig := env.NewHolmesConfig()
	if !envConfig.Common.Enable {
		return nil, errors.New("holmes not enabled")
	}
	options := make([]holmes.Option, 0, 16)

	options = append(options, holmes.WithCollectInterval(envConfig.Common.Interval))
	cwd, _ := os.Getwd()
	options = append(options, holmes.WithDumpPath(cwd+envConfig.Common.DumpPath))
	options = append(options, holmes.WithCPUMax(envConfig.Common.DumpCpuMax))
	if envConfig.Common.DumpType == "binary" {
		options = append(options, holmes.WithBinaryDump())
	} else {
		options = append(options, holmes.WithTextDump())
	}
	var i env.ConfigType
	for i = 0; i < env.SoloConfig1End; i++ {
		soloConfig := envConfig.Solo[i].(env.HolmesSoloConfig1)
		if soloConfig.Enable {
			options = append(options, DumpSetfuncArr1[i](
				soloConfig.Min, soloConfig.Diff, soloConfig.Abs, time.Duration(soloConfig.CoolDown)*time.Second))
		}
	}

	for i = env.SoloConfig1End; i < env.SoloConfig2End; i++ {
		soloConfig := envConfig.Solo[i].(env.HolmesSoloConfig2)
		if soloConfig.Enable {
			options = append(options, DumpSetfuncArr2[i](
				soloConfig.Min, soloConfig.Diff, soloConfig.Abs, soloConfig.Max, time.Duration(soloConfig.CoolDown)*time.Second))
		}
	}

	h, _ := holmes.New(options...)

	for i = 0; i < env.SoloConfig1End; i++ {
		soloConfig := envConfig.Solo[i].(env.HolmesSoloConfig1)
		if soloConfig.Enable {
			h = DumpEnableFuncArr[i](h)
		}
	}
	for i = env.SoloConfig1End; i < env.SoloConfig2End; i++ {
		soloConfig := envConfig.Solo[i].(env.HolmesSoloConfig2)
		if soloConfig.Enable {
			h = DumpEnableFuncArr[i](h)
		}
	}
	return h, nil
}
