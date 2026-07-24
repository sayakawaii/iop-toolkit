package controller

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

type respDownloadsData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Path        string `json:"path"`
	Size        uint64 `json:"size"`
	Mtime       string `json:"mtime"`
	HasChildren bool   `json:"hasChildren"`
}

var baseDir string = "./static/uploads"

func getDirTree(path string) []respDownloadsData {
	var result []respDownloadsData

	// Ensure path exists and is a directory
	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		return result
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return result
	}

	for _, e := range entries {
		info, err := e.Info()
		if err != nil {
			continue
		}

		fullPath := filepath.Join(path, e.Name())
		// Normalize path to use forward slashes for JSON/URL usage
		urlPath := "/" + filepath.ToSlash(fullPath)

		itemType := "file"
		hasChildren := false
		var size uint64 = 0

		if info.IsDir() {
			itemType = "dir"
			// check if directory has children
			if children, err := os.ReadDir(fullPath); err == nil && len(children) > 0 {
				hasChildren = true
			}
		} else {
			size = uint64(info.Size())
		}

		result = append(result, respDownloadsData{
			ID:          filepath.ToSlash(fullPath),
			Name:        e.Name(),
			Type:        itemType,
			Path:        urlPath,
			Size:        size,
			Mtime:       info.ModTime().Format("2006-01-02 15:04:05"),
			HasChildren: hasChildren,
		})
	}

	slices.Reverse(result)
	return result
}

func (controller *AppController) DownloadsList(context *gin.Context) {
	var ret []respDownloadsData
	reqPath := context.Query("path")
	if reqPath != "" {
		cleaned := filepath.FromSlash(reqPath)
		cleaned = strings.TrimPrefix(cleaned, string(os.PathSeparator))

		candidate := cleaned
		if filepath.IsAbs(cleaned) {
			candidate = cleaned
		} else {
			candidate = filepath.Join(".", cleaned)
		}
		ret = getDirTree(candidate)
	} else {
		ret = getDirTree(baseDir)
	}
	if ret != nil {
		context.JSON(http.StatusOK, ret)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	}
}

// resolveTarget 把前端传入的 path 规范化并返回绝对路径及文件信息，
// 同时确保目标在 baseDir 之内，防止目录遍历。
func resolveTarget(reqPath string) (absTarget string, info os.FileInfo, err error) {
	if reqPath == "" {
		return "", nil, fmt.Errorf("missing path")
	}

	// convert slashes, remove leading separator so relative join works predictably
	cleaned := filepath.FromSlash(reqPath)
	cleaned = strings.TrimPrefix(cleaned, string(os.PathSeparator))

	candidate := cleaned
	if filepath.IsAbs(cleaned) {
		candidate = cleaned
	} else {
		candidate = filepath.Join(".", cleaned)
	}

	absTarget, err = filepath.Abs(candidate)
	if err != nil {
		return "", nil, fmt.Errorf("invalid path")
	}

	baseAbs, err := filepath.Abs(baseDir)
	if err != nil {
		return "", nil, fmt.Errorf("server configuration error")
	}

	rel, err := filepath.Rel(baseAbs, absTarget)
	if err != nil {
		return "", nil, fmt.Errorf("access denied")
	}
	if strings.HasPrefix(rel, "..") || rel == ".." {
		return "", nil, fmt.Errorf("access denied")
	}

	info, err = os.Stat(absTarget)
	if err != nil {
		return "", nil, err
	}

	return absTarget, info, nil
}

func (controller *AppController) DownloadsDownload(context *gin.Context) {
	reqPath := context.Query("path")
	absTarget, info, err := resolveTarget(reqPath)
	if err != nil {
		// map common errors to appropriate status codes
		switch {
		case err.Error() == "missing path":
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case strings.Contains(err.Error(), "access denied"):
			context.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		case strings.Contains(err.Error(), "server configuration error"):
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		default:
			// os.Stat errors etc.
			if os.IsNotExist(err) {
				context.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
			} else {
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		}
		return
	}

	if info.IsDir() {
		context.JSON(http.StatusBadRequest, gin.H{"error": "path is a directory"})
		return
	}

	// send binary file as attachment (sets content-disposition)
	context.FileAttachment(absTarget, filepath.Base(absTarget))
}

// DownloadsZip 接收 POST JSON { paths: []string }，将所有选中的文件/目录打包成 zip 并返回二进制流。
// 注意：当前实现会把 zip 中的路径相对于 baseDir 保留（去除 baseDir 前缀）。
func (controller *AppController) DownloadsZip(context *gin.Context) {
	var req struct {
		Paths []string `json:"paths"`
	}
	if err := context.BindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if len(req.Paths) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "no paths provided"})
		return
	}

	baseAbs, err := filepath.Abs(baseDir)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "server configuration error"})
		return
	}

	buf := &bytes.Buffer{}
	zw := zip.NewWriter(buf)

	// track added entries to avoid duplicates
	added := make(map[string]struct{})

	// helper to add a single file into zip with given absolute path
	addFile := func(absPath string) error {
		relName, err := filepath.Rel(baseAbs, absPath)
		if err != nil {
			return err
		}
		// use forward slashes in zip
		zipName := filepath.ToSlash(relName)
		if zipName == "" {
			return nil
		}
		if _, exists := added[zipName]; exists {
			return nil
		}
		added[zipName] = struct{}{}

		f, err := os.Open(absPath)
		if err != nil {
			return err
		}
		defer f.Close()

		fw, err := zw.Create(zipName)
		if err != nil {
			return err
		}
		_, err = io.Copy(fw, f)
		return err
	}

	// iterate provided paths
	for _, p := range req.Paths {
		absTarget, info, err := resolveTarget(p)
		if err != nil {
			// stop and return error to client for offending path
			_ = zw.Close()
			context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid path %s: %v", p, err)})
			return
		}

		if info.IsDir() {
			// walk directory and add files
			err := filepath.WalkDir(absTarget, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if d.IsDir() {
					return nil
				}
				return addFile(path)
			})
			if err != nil {
				_ = zw.Close()
				context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read directory"})
				return
			}
		} else {
			if err := addFile(absTarget); err != nil {
				_ = zw.Close()
				context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add file"})
				return
			}
		}
	}

	if err := zw.Close(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to build zip"})
		return
	}

	// return zip binary
	context.Header("Content-Type", "application/zip")
	context.Header("Content-Disposition", "attachment; filename=\"download.zip\"")
	context.Data(http.StatusOK, "application/zip", buf.Bytes())
}
