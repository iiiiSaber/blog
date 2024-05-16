package model

import (
	"blog/utils"
	"blog/utils/errmsg"
	"context"
	"log"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.MinioServer
var useSSL = false

func setConfig() *minio.Client {
    minioClient, err := minio.New(ImgUrl, &minio.Options{
        Creds:  credentials.NewStaticV4(AccessKey, SecretKey, ""),
        Secure: useSSL, // Set to true if your MinIO server uses SSL
    })
    if err != nil {
        log.Fatalln(err)
    }
    return minioClient
}
func UpLoadFile(file multipart.File, fileSize int64, fileName string) (string, int) {
    minioClient := setConfig()

    // Upload the file directly using the fileName provided
    info, err := minioClient.PutObject(context.Background(), Bucket, fileName, file, fileSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
    if err != nil {
        return "", errmsg.ERROR
    }

    // Log the details of the uploaded object
    log.Printf("Successfully uploaded %s with size %d and etag %s", fileName, info.Size, info.ETag)

    url := ImgUrl + "/" + Bucket + "/" + fileName
    return url, errmsg.SUCCESS
}