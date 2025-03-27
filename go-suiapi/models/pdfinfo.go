package models

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// PDFFile 结构体表示一个 PDF 文件的信息
type PDFFile struct {
	FileName string `json:"fileName"`
	FilePath string `json:"filePath"`
}

// getPDFFiles 获取目录下的所有 PDF 文件
func GetPDFFiles(directoryPath string) ([]PDFFile, error) {
	//var pdfFiles []string
	var pdfFiles []PDFFile
	// 读取目录
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return nil, err
	}
	var temps = " /src/files/pdffile/"
	// 遍历目录中的文件
	//for _, file := range files {
	//	// 检查文件是否是 PDF 文件
	//	if strings.HasSuffix(file.Name(), ".pdf") {
	//		// 构建文件的完整路径
	//		filePath := filepath.Join(temps, file.Name())
	//		pdfFiles = append(pdfFiles, filePath)
	//	}
	//}
	// 遍历目录中的文件
	for _, file := range files {
		// 检查文件是否是 PDF 文件
		if strings.HasSuffix(file.Name(), ".pdf") {
			// 构建文件的完整路径
			filePath := filepath.Join(temps, file.Name())

			// 创建 PDFFile 结构体并添加到列表中
			pdfFile := PDFFile{
				FileName: file.Name(),
				FilePath: filePath,
			}
			pdfFiles = append(pdfFiles, pdfFile)
		}
	}
	return pdfFiles, nil
}
