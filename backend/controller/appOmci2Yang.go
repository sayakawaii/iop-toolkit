/*
# ------------------------------------------------------------
# -- appOmci2Yang.go
# --
# -- Turn an analysed OMCI log into LightSpan NETCONF config.
# ------------------------------------------------------------
*/

package controller

import (
	"net/http"
	"omciAnalyzer/dao"
	"omciAnalyzer/service/omci2yang"
	"omciAnalyzer/utils"

	"github.com/gin-gonic/gin"
)

// GetOmci2YangBoards reports which YANG sets this deployment can compile
// against, so the UI offers exactly what shipped instead of a guess.
func (controller *AppController) GetOmci2YangBoards(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"boards": omci2yang.Boards(),
	})
}

type reqOmci2Yang struct {
	RequestKey string `json:"requestKey"`
	// OnuName identifies the analysed ONU within this request, i.e. the name
	// the uploaded log used.
	OnuName string `json:"onuName"`
	// OltOnuName is how the *target* OLT refers to the ONU, which is not
	// necessarily what the log called it. It defaults to OnuName.
	OltOnuName string `json:"oltOnuName"`
	// Board names the LT, whose confd serves the mounted tree being compiled
	// against; the platform-specific modules differ between boards.
	Board string `json:"board"`
	// The rest are OLT-side facts OMCI cannot carry. Left empty they are
	// either derived by convention or reported as a gap, never guessed at.
	Vendor        string `json:"vendor"`
	ChassisName   string `json:"chassisName"`
	VeipComponent string `json:"veipComponent"`
}

// GenerateOmci2YangConfig compiles one analysed ONU into configuration.
//
// The analyzer already wrote the per-ONU OMCI JSON this reads, so generating a
// config needs no re-analysis: it is the same artefact the OMCI view renders,
// handed to the compiler instead of to a table.
func (controller *AppController) GenerateOmci2YangConfig(context *gin.Context) {
	var request reqOmci2Yang
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	if request.RequestKey == "" || request.OnuName == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "requestKey and onuName are both required",
		})
		return
	}

	record := new(dao.OmciAnalyzerRequestRecord)
	if err := dao.MysqlRecordDataRead(record, "request_key",
		request.RequestKey); err != nil {
		utils.Log("GenerateOmci2YangConfig: unknown requestKey:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "no analysis found for this requestKey",
		})
		return
	}

	analysed, found := record.Content.OnuContent[request.OnuName]
	if !found {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "this analysis holds no ONU by that name",
		})
		return
	}

	oltOnuName := request.OltOnuName
	if oltOnuName == "" {
		oltOnuName = request.OnuName
	}

	result := omci2yang.Compile(omci2yang.Request{
		JSONPath:      analysed.JsonPath,
		Board:         request.Board,
		OnuName:       oltOnuName,
		Vendor:        request.Vendor,
		ChassisName:   request.ChassisName,
		VeipComponent: request.VeipComponent,
	})

	// A config that fails validation is still a result worth returning: the
	// errors are what say why, and hiding them behind a 500 would lose them.
	if result.Error != "" {
		utils.Log("GenerateOmci2YangConfig:", request.OnuName, result.Error)
		context.JSON(http.StatusUnprocessableEntity, result)
		return
	}
	context.JSON(http.StatusOK, result)
}
