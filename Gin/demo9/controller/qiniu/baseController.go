package qiniu

import (
	"fmt"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type BaseController struct {
}

func (con BaseController) Init() (string, *storage.FormUploader) {
	accessKey := "f3oPRxqfI4Ut4_tqfR7UHNB1xohonm5S5tnNQS-G"
	secretKey := "HTE7cP2Wc_2lRrQ_TFApHw7wJf5DhatORUoNDHxa"
	mac := qbox.NewMac(accessKey, secretKey)
	bucket := "imagesku"

	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	cfg := storage.Config{}
	upToken := putPolicy.UploadToken(mac)
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	return upToken, formUploader
}
func (con BaseController) GetBucketManager() *storage.BucketManager {
	accessKey := "f3oPRxqfI4Ut4_tqfR7UHNB1xohonm5S5tnNQS-G"
	secretKey := "HTE7cP2Wc_2lRrQ_TFApHw7wJf5DhatORUoNDHxa"
	mac := qbox.NewMac(accessKey, secretKey)
	bucket := "imagesku"

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	bucketManager := storage.NewBucketManager(mac, &cfg)
	resURL := "http://devtools.qiniu.com/qiniu.png"
	fetchRet, err := bucketManager.Fetch(resURL, bucket, "qiniu.png")
	if err != nil {
		fmt.Println("fetch error,", err)
	} else {
		fmt.Println(fetchRet.String())
	}
	return bucketManager
}
func (con BaseController) PublicAccessURL(key string) string {
	domain := "http://pic.kiass.top/"
	publicAccessURL := storage.MakePublicURL(domain, key)
	return publicAccessURL
}
