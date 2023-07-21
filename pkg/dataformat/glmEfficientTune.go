package dataformat

import (
	"context"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
	"strings"
)

type GLMEfficientTune struct {
}

func NewGLMEfficientTune() *GLMEfficientTune {
	return &GLMEfficientTune{}
}

type SelfCognition struct {
	Instruction string `json:"instruction"`
	Input       string `json:"input"`
	Output      string `json:"output"`
}

func (tune *GLMEfficientTune) ConvertExcelToSelfCognitionData(ctx context.Context, excelFileData *multipart.File) ([]*SelfCognition, error) {
	data := []*SelfCognition{}

	// 不支持csv格式
	file, err := excelize.OpenReader(*excelFileData)
	if err != nil {
		return nil, err
	}
	defer (*excelFileData).Close()

	// 假设数据在第一个工作表中
	sheetName := file.GetSheetName(0)
	rows, err := file.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	// 遍历每一行并提取问题和答案
	for _, row := range rows {
		if len(row) >= 2 {
			question := strings.Replace(row[0], "\n", "", -1)
			answer := strings.Replace(row[1], "\n", "", -1)

			selfCognition := &SelfCognition{
				Instruction: question,
				Output:      answer,
			}

			data = append(data, selfCognition)
		}
	}

	return data, nil
}
