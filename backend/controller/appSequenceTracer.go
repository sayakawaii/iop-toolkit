package controller

import (
	"net/http"
	"omciAnalyzer/api"
	"omciAnalyzer/service/sequencetracer"
	"omciAnalyzer/utils"

	"github.com/gin-gonic/gin"
)

func (controller *AppController) SequenceTracerRequest(context *gin.Context) {
	form, err := context.MultipartForm()
	if err != nil || form == nil {
		utils.Log("SequenceTracerRequest: multipart parse failed:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid or incomplete multipart upload",
		})
		return
	}
	files := form.File["sequencetracerFile"]
	logs := api.SaveUploadedFileInSameDir(false, files, context)
	mermaid := sequencetracer.SequenceTracer(logs)
	utils.Log("mermaid:", mermaid)
	var isEmpty bool
	if mermaid == "" {
		isEmpty = true
	}
	context.JSON(http.StatusOK, gin.H{
		"mermaidData": mermaid,
		"isEmpty":     isEmpty,
	})
}
