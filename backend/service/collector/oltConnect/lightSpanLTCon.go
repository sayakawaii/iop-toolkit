/*
# ------------------------------------------------------------
# -- lightSpanLTCon.go
# --
# -- Huang Minghe
# -- 2022-8-29
# ------------------------------------------------------------
*/
package oltConnect

const (
	LightSpanLT = 1
)

type conLightSpanLT struct {
	conHandlerBase
}

func init() {
	con := new(conLightSpanLT)
	conHandlerRegist(LightSpanLT, con)
}
