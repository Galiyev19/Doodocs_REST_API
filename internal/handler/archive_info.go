package handler

import (
	"archive/zip"
	"bytes"
	"doodocs_rest_api/internal/entity"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ArchiveInfo(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
		return
	}
	defer file.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}
	archiveDetail, err := processArchive(buf.Bytes())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, archiveDetail)

}

func processArchive(data []byte) (entity.ArchiveDetail, error) {
	reader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return entity.ArchiveDetail{}, fmt.Errorf("failed to read zip file: %v", err)
	}

	var totalSize float64
	var files []entity.FileInfo
	for _, file := range reader.File {
		fileInfo := file.FileInfo()
		size := float64(fileInfo.Size())
		mimeType := http.DetectContentType([]byte(filepath.Ext(file.Name)))
		files = append(files, entity.FileInfo{
			FilePath: file.Name,
			Size:     size,
			MimeType: mimeType,
		})
		totalSize += size
	}

	return entity.ArchiveDetail{
		Filename:    "my_archive.zip",
		ArchiveSize: float64(len(data)),
		TotalSize:   totalSize,
		TotalFiles:  len(files),
		Files:       files,
	}, nil
}
