package service

import (
	"archive/zip"
	"bytes"
	"doodocs_rest_api/internal/entity"
	"fmt"
	"net/http"
	"path/filepath"
	"unicode/utf8"
)

type ArchiveInfoService struct{}

func NewArchiveInfoService() *ArchiveInfoService {
    return &ArchiveInfoService{}
}

func (s *ArchiveInfoService) ProcessArchiveData(archive []byte) (entity.ArchiveDetail, error) {
    reader, err := zip.NewReader(bytes.NewReader(archive), int64(len(archive)))
    if err != nil {
        return entity.ArchiveDetail{}, fmt.Errorf("failed to read zip file: %v", err)
    }

    var totalSize float64
    var files []entity.FileInfo

    for _, file := range reader.File {
        fileInfo := file.FileInfo()
        size := float64(fileInfo.Size())

        // Проверяем, является ли имя файла валидной UTF-8 строкой
        utf8FileName := file.Name
        if !utf8.ValidString(utf8FileName) {
            utf8FileName = string([]rune(file.Name))
        }

        fmt.Println("file name", utf8FileName)
        mimeType := http.DetectContentType([]byte(filepath.Ext(utf8FileName)))

        files = append(files, entity.FileInfo{
            FilePath: utf8FileName,
            Size:     size,
            MimeType: mimeType,
        })
        totalSize += size
    }

    return entity.ArchiveDetail{
        Filename:    "my_archive.zip",
        ArchiveSize: float64(len(archive)),
        TotalSize:   totalSize,
        TotalFiles:  len(files),
        Files:       files,
    }, nil
}
