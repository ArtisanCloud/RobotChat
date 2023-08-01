package objectx

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func GetSavePath(path string, fileName string, fileType string) (string, error) {
	// 生成当前日期时间作为目录名
	currentTime := time.Now().Format("20060102_150405")

	// 拼接目录路径
	savePath := filepath.Join(path, fileName+"_"+currentTime+"."+fileType)

	// 创建目录
	err := os.MkdirAll(filepath.Dir(savePath), os.ModePerm)
	if err != nil {
		return "", err
	}

	return savePath, nil
}

func SaveObjectToPath(obj interface{}, path string) error {
	// 将对象转换为JSON格式的字节切片
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	// 将数据写入文件
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetTempFilePath(file *multipart.FileHeader) (string, error) {
	// 创建临时目录
	tempDir := os.TempDir()

	// 将上传的文件保存到临时目录中
	tempFilePath := filepath.Join(tempDir, file.Filename)
	err := SaveUploadedFile(file, tempFilePath)
	if err != nil {
		return "", err
	}

	return tempFilePath, nil
}

func SaveUploadedFile(file *multipart.FileHeader, filePath string) error {
	srcFile, err := file.Open()
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func GetFileBytes(file *multipart.File) ([]byte, error) {
	buffer := bytes.Buffer{}
	_, err := io.Copy(&buffer, *file)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func CreateDirectoriesForFiles(outputFile string) error {
	dir := filepath.Dir(outputFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return nil
}
