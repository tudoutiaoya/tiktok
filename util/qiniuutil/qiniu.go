package qiniuutil

import (
	"context"
	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"strings"
	"tiktok/config"
)

var (
	bucket    = ""
	accessKey = ""
	secretKey = ""
	prefix    = ""
	videoPath = ""

	formUploader *storage.FormUploader
	upToken      = ""
)

func InitQiniu(config *config.Configuration) {
	// 配置信息
	qiniuSettings := config.QiniuSettings
	bucket = qiniuSettings.Bucket
	accessKey = qiniuSettings.AccessKey
	secretKey = qiniuSettings.SecretKey
	prefix = qiniuSettings.Prefix
	videoPath = qiniuSettings.VideoPath
	initFormUploader()
}

// 初始化文件上传器
func initFormUploader() {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken = putPolicy.UploadToken(mac)
	cfg := storage.Config{
		// 空间对应的机房
		Region: &storage.ZoneHuabei,
		// 是否使用https域名
		UseHTTPS: false,
		// 上传是否使用CDN上传加速
		UseCdnDomains: true,
	}
	formUploader = storage.NewFormUploader(&cfg)
}

func Upload(file *multipart.FileHeader) (string, error) {
	// 文件名称
	key := videoPath + uuid.New().String() + "." + strings.Split(file.Filename, ".")[1]
	// 文件大小
	dataLen := file.Size
	// 文件内容
	src, err := file.Open()
	if err != nil {
		return "", err
	}

	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err = formUploader.Put(context.Background(), &ret, upToken, key, src, dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return prefix + ret.Key, nil
}
