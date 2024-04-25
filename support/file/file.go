package file

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"
)

// ClientOriginalExtension 获取文件的原始扩展名
func ClientOriginalExtension(file string) string {
	return strings.ReplaceAll(filepath.Ext(file), ".", "")
}

// Contain 判断字符串是否包含在文件内容中
func Contain(file string, search string) bool {
	if Exists(file) {
		data, err := os.ReadFile(file)
		if err != nil {
			return false
		}
		return strings.Contains(string(data), search)
	}

	return false
}

// Create 创建文件
// content: 文件内容
func Create(file string, content string) error {
	if err := os.MkdirAll(filepath.Dir(file), os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		return err
	}

	return nil
}

// Exists 判断文件是否存在
func Exists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// Extension Supported types: https://github.com/gabriel-vasile/mimetype/blob/master/supported_mimes.md
// 返回文件后缀
// originalWhenUnknown: 是否返回原文件后缀
func Extension(file string, originalWhenUnknown ...bool) (string, error) {
	mime, err := mimetype.DetectFile(file)
	if err != nil {
		return "", err
	}

	if mime.String() == "" {
		if len(originalWhenUnknown) > 0 {
			if originalWhenUnknown[0] {
				return ClientOriginalExtension(file), nil
			}
		}

		return "", errors.New("unknown file extension")
	}

	return strings.TrimPrefix(mime.Extension(), "."), nil
}

// LastModified 获取文件最后修改时间
func LastModified(file, timezone string) (time.Time, error) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return time.Time{}, err
	}

	l, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	return fileInfo.ModTime().In(l), nil
}

// MimeType 获取文件类型
func MimeType(file string) (string, error) {
	mime, err := mimetype.DetectFile(file)
	if err != nil {
		return "", err
	}

	return mime.String(), nil
}

// Remove 删除文件
func Remove(file string) error {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	return os.RemoveAll(file)
}

// Size 返回文件字节数
func Size(file string) (int64, error) {
	fileInfo, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer fileInfo.Close()

	fi, err := fileInfo.Stat()
	if err != nil {
		return 0, err
	}

	return fi.Size(), nil
}
