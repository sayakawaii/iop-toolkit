/*
# ------------------------------------------------------------
# -- shapeDefine.go
# --
# -- Huang Minghe
# -- 2022-6-21
# ------------------------------------------------------------
*/

package omciShape

type shaper interface {
	Shape(logPath string, outputPath string) map[string]string
}

var (
	ShaperDef = make(map[uint8]shaper)
)

func shaperRegist(key uint8, s shaper) {
	ShaperDef[key] = s
}
