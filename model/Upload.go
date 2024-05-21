package model

import (
	"blog/utils"
	"blog/utils/errmsg"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.MinioServer
var useSSL = false

type ViewContentType struct {
	Prefix string
	Type   string
}

// 文件类型map
var contentTypes = map[string]ViewContentType{
    "default": {"default", "application/octet-stream"},
    "png":     {"png", "image/png"},
    "jpeg":    {"jpeg", "image/jpeg"},
    "jpg":     {"jpg", "image/jpeg"},
    "gif":     {"gif", "image/gif"},
    "wbmp":    {"wbmp", "image/vnd.wap.wbmp"},
    "tiff":    {"tiff", "image/tiff"},
    "jfif":    {"jfif", "image/jpeg"},
    "tif":     {"tif", "image/tiff"},
    "fax":     {"fax", "image/fax"},
    "jpe":     {"jpe", "image/jpeg"},
    "net":     {"net", "image/pnetvue"},
    "rp":      {"rp", "image/vnd.rn-realpix"},
    "ico":     {"ico", "image/x-icon"},
}

// 获取文件类型
func GetTypeByPrefix(prefix string) string {
	if contentType, exists := contentTypes[prefix]; exists {
		return contentType.Type
	}
	return contentTypes["default"].Type
}
func GetContentType(filename string) string {
	if filename == "" {
		return contentTypes["default"].Type
	}
	prefix := filename[strings.LastIndex(filename, ".")+1:]
	return GetTypeByPrefix(prefix)
}

// 文件hash计算方法 code from StackOverflow，fix by GPT-4
func calculateHash(file multipart.File) string {
    h := md5.New()
    // Reset file pointer to the start of the file
    if _, err := file.Seek(0, 0); err != nil {
        log.Fatal("查找文件出错: %v", err)
        return ""
    }
    if _, err := io.Copy(h, file); err != nil {
        log.Fatal("获取文件hash出错: %v", err)
        return ""
    }
    md5Str := hex.EncodeToString(h.Sum(nil))
    log.Println("文件的hash值为: %s", md5Str)
    // Reset file pointer to the start of the file again for further use
    if _, err := file.Seek(0, 0); err != nil {
        log.Fatal("查找文件出错: %v", err)
    }
    return md5Str
}
// 列出桶中所有对象
func listAllObjects(client *minio.Client, bucketName string) ([]minio.ObjectInfo, error) {
    objectCh := client.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{})
    var objects []minio.ObjectInfo
    for object := range objectCh {
        if object.Err != nil {
            return nil, object.Err
        }
        objects = append(objects, object)
    }
    return objects, nil
}

// 设置上传选项
func setConfig() *minio.Client {
    minioClient, err := minio.New(ImgUrl, &minio.Options{
        Creds:  credentials.NewStaticV4(AccessKey, SecretKey, ""),
        Secure: useSSL, 
    })
    if err != nil {
        log.Fatalln(err)
    }
    return minioClient
}

// 上传文件方法
func UpLoadFile(file multipart.File, fileSize int64, fileName string) (string, int) {
    minioClient := setConfig()
    // 获取文件类型
    contentType := GetContentType(fileName)
    // 获取时间戳
    expiryTime := time.Duration(24) * time.Hour
    // 完善上传的url参数
    reqParams := make(url.Values)
    reqParams.Set("response-content-disposition", "inline")
    reqParams.Set("response-content-type", contentType)
    
    // 获取桶中对象
    minioFiles,_:=listAllObjects(minioClient,Bucket)
    // 计算选择文件的hash
    localFileHash:= calculateHash(file)

    // 遍历桶中对象
    for _, minioFile := range minioFiles {
        // minioFileHash, err := getMinIOFileHash(minioClient, Bucket, object.)
        // 比较hash是否一致，如一致则取消上传并直接返回链接
        if localFileHash == minioFile.ETag {
            log.Println("发现hash一致的文件:", minioFile.Key)
            preUrl,err:=minioClient.PresignedGetObject(context.Background(),Bucket,fileName,expiryTime,reqParams)
            if err!=nil{
                return "",errmsg.ERROR
            }
            return preUrl.String(),errmsg.SUCCESS
        }
        // 比较名称是否一致，如果一致，上传会直接覆盖文件，所以为了避免这个情况，我对要上传的文件进行了重命名
        if fileName ==  minioFile.Key{
            log.Println("发现命名一致的文件，正在重命名")
            fileName=strings.TrimSuffix(fileName,filepath.Ext(fileName)) + "_new" + filepath.Ext(fileName)
        }
    }

        // 遍历结束，操作结束之后，可以上传文件了
        fmt.Println("没有找到hash一致的文件，上传中")
        // 上传文件
        info, _ := minioClient.PutObject(context.Background(), Bucket, fileName, file,fileSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
        log.Printf("上传成功 文件名 %s 文件大小 %d hash值为 %s ", fileName, info.Size, info.ETag)
        // 获取到预览链接，避免返回的链接访问着是直接下载文件
        preUrl,err:=minioClient.PresignedGetObject(context.Background(),Bucket,fileName,expiryTime,reqParams)
        if err!=nil{
            return "",errmsg.ERROR
        }
        return preUrl.String(),errmsg.SUCCESS
}