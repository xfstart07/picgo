package cos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/xfstart07/picgo/config" // 替换为你的项目路径

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
)

// CosClientWrapper wraps the COS client.
type CosClientWrapper struct {
	Client *cos.Client
}

// NewCosClient initializes a new COS client with the provided configuration.
func NewCosClient(cosConfig *config.CosConfig) *CosClientWrapper {
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", cosConfig.Bucket, cosConfig.Region))
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cosConfig.SecretId,
			SecretKey: cosConfig.SecretKey,
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	return &CosClientWrapper{Client: client}
}

// UploadFile uploads a file to the specified COS bucket.
func (cw *CosClientWrapper) UploadFile(storePath string, filePath string) (string, error) {
	// Open the file to be uploaded
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Retrieve the file's name
	_, fileName := filepath.Split(filePath)

	fileName = storePath + "/" + fileName

	// Perform the file upload
	_, err = cw.Client.Object.Put(context.Background(), fileName, file, nil)
	if err != nil {
		return "", err
	}
	fmt.Printf("File '%s' uploaded successfully\n", filePath)

	// 构建图片的URL地址
	objectURL := cw.Client.Object.GetObjectURL(fileName).String()

	return objectURL, nil
}

// ListBucketObjects lists the objects in the specified COS bucket.
func (cw *CosClientWrapper) ListBucketObjects() error {
	// Define the result struct for listing objects
	s, _, err := cw.Client.Bucket.Get(context.Background(), &cos.BucketGetOptions{
		MaxKeys: 100,
	})
	if err != nil {
		return fmt.Errorf("failed to list objects in bucket: %v", err)
	}
	for _, object := range s.Contents {
		fmt.Println(object.Key)
	}
	return nil
}

// DeleteObject deletes an object from the specified COS bucket.
func (cw *CosClientWrapper) DeleteObject(objectKey string) error {
	_, err := cw.Client.Object.Delete(context.Background(), objectKey)
	if err != nil {
		return err
	}
	fmt.Printf("Object '%s' deleted successfully\n", objectKey)
	return nil
}
