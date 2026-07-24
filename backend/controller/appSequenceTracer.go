package controller

import (
	"net/http"
	"omciAnalyzer/api"
	"omciAnalyzer/service/sequencetracer"
	"omciAnalyzer/utils"

	"github.com/gin-gonic/gin"
)

func (controller *AppController) SequenceTracerRequest(context *gin.Context) {
	form, _ := context.MultipartForm()
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
