package upload

import (
	"fmt"
	"gitee.com/phper95/pkg/file"
	"log"
	"mime/multipart"
	"os"
	"path"
	"shop/pkg/global"
	"shop/pkg/logging"
	"shop/pkg/util"
	"strings"
)

// GetImageFullUrl get the full access path
func GetImageFullUrl(name string) string {
	return global.CONFIG.App.PrefixUrl + "/" + GetImagePath() + name
}

// GetImageName get image name
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// GetImagePath get save path
func GetImagePath() string {
	return global.CONFIG.App.ImageSavePath
}

// GetImageFullPath get full save path
func GetImageFullPath() string {
	return global.CONFIG.App.RuntimeRootPath + GetImagePath()
}

// CheckImageExt check image file ext
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range global.CONFIG.App.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckImageSize check image size
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= global.CONFIG.App.ImageMaxSize
}

// CheckImage check if the file exists
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
