/**
 * @file sessionOperator.go
 * @author Minghe Huang (minghe.huang@nokia.com)
 * @brief
 * @version 0.1
 * @date 2023-11-22
 *
 * @copyright Copyright (c) 2023
 *
 */

package api

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"omciAnalyzer/dao"
	"omciAnalyzer/models"
	"omciAnalyzer/utils"
	"os"
	"path"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	baseDir = "./static/uploads"
)

func isSessionExist(context *gin.Context) bool {
	session := sessions.Default(context)
	// cookie, err := context.Cookie("userid")
	key := session.Get("userid")
	if key == nil {
		//session not set
		return false
	} else {
		//session exist
		return true
	}
}

func getSessionID(context *gin.Context) string {
	session := sessions.Default(context)
	key := session.Get("userid")
	if key == nil {
		//session not set
		return ""
	} else {
		//session exist
		return key.(string)
	}
}

func setSessionID(user string, context *gin.Context) string {
	session := sessions.Default(context)
	session.Set("userid", user)
	session.Save()

	return user
}

func SaveUploadedFileWithSessionID(persist bool, file *multipart.FileHeader, context *gin.Context) *dao.LogStorageInfo {
	var user string
	if isSessionExist(context) {
		user = getSessionID(context)
	} else {
		user = setSessionID(models.GenerateNanoID(), context)
	}

	var dir string
	if persist {
		dir = path.Join(baseDir, "permanent")
	} else {
		dir = path.Join(baseDir, "temporary", models.GetDay())
		dir = path.Join(dir, user)
	}
	fileID := models.GenerateReadableFileID()
	//add file actually name for yangHelper
	dir = path.Join(dir, fileID)
	dst := path.Join(dir, file.Filename)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	context.SaveUploadedFile(file, dst)

	log := new(dao.LogStorageInfo)
	log.LogUser = user
	log.LogID = fileID
	log.LogName = file.Filename
	log.UploadTime = models.GetDate()
	log.LogDir = dir + "/"
	log.LogPath = dst

	return log
}

func SaveMinioFileWithSessionID(minioKey string, context *gin.Context) *dao.LogStorageInfo {
	var user string
	if isSessionExist(context) {
		user = getSessionID(context)
	} else {
		user = setSessionID(models.GenerateNanoID(), context)
	}

	// Prepare local directory for downloaded file
	dir := path.Join(baseDir, "temporary", models.GetDay(), user)
	fileID := models.GenerateReadableFileID()
	dir = path.Join(dir, fileID)

	// Extract filename from minioKey (use last part of path)
	filename := path.Base(minioKey)
	localPath := path.Join(dir, filename)

	// Download file from MinIO
	err := DownloadFileFromMinio(minioKey, localPath)
	if err != nil {
		utils.Log("Error downloading from MinIO:", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to download file from MinIO",
		})
		return nil
	}
	logInfo := &dao.LogStorageInfo{
		LogUser:    user,
		LogID:      fileID,
		LogName:    filename,
		UploadTime: models.GetDate(),
		LogDir:     dir + "/",
		LogPath:    localPath,
	}
	return logInfo
}

func SaveUploadedFileInSameDir(persist bool, files []*multipart.FileHeader, context *gin.Context) []string {
	var user string
	if isSessionExist(context) {
		user = getSessionID(context)
	} else {
		user = setSessionID(models.GenerateNanoID(), context)
	}

	var dir string
	dir = path.Join(baseDir, "sequencetracer", models.GetDay())
	dir = path.Join(dir, user)
	fileID := models.GenerateReadableFileID()
	dir = path.Join(dir, fileID)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		utils.Log(err)
		return nil
	}
	var logs []string
	for _, file := range files {
		dst := path.Join(dir, file.Filename)

		context.SaveUploadedFile(file, dst)
		logs = append(logs, dst)
	}

	return logs
}
