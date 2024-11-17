package handler

import (
	"bytes"
	"io"
	"net/http"

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

	archiveDetail, err := h.services.ArchiveInfo.ProcessArchiveData(buf.Bytes())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, archiveDetail)

}
