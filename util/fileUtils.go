package util

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/**
指定目录，压缩为 zip 文件
dir 路径
zip_file_name zip 文件名
*/
func CreateZip(Path, zipFullName string) error {
	//创建一个压缩文件
	zipFile, err := os.Create(zipFullName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	//创建压缩文件的writer
	zipper := zip.NewWriter(zipFile)
	defer zipper.Close()

	walkFn := func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if fileInfo.IsDir() {
			return nil
		}
		Src, _ := os.Open(path)
		defer Src.Close()
		//FileName, _ := Zip.Create(Path)
		h := &zip.FileHeader{Name: path, Method: zip.Deflate, Flags: 0x800}
		FileName, _ := zipper.CreateHeader(h)
		io.Copy(FileName, Src)
		zipper.Flush()
		return nil
	}
	return filepath.Walk(Path, walkFn)
}
