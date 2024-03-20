package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xfstart07/picgo/config" // 替换为你的项目路径
	"github.com/xfstart07/picgo/cos"    // 替换为你的项目路径
)

func main() {
	// Define command line flags
	configPath := flag.String("config", "./config.json", "Path to the configuration file")
	storePath := flag.String("storePath", "image", "Store to the file to be uploaded")
	uploadFilePath := flag.String("upload", "", "Path to the file to be uploaded")
	listBucket := flag.Bool("list", false, "List the objects in the bucket")
	deleteObjectKey := flag.String("delete", "", "Key of the object to delete in the bucket")

	flag.Parse()

	// Load the configuration
	cosConfig, err := config.LoadConfig(*configPath)
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	cosClient := cos.NewCosClient(cosConfig)
	if *uploadFilePath != "" {
		objectURL, err := cosClient.UploadFile(*storePath, *uploadFilePath)
		if err != nil {
			fmt.Println("Error uploading file:", err)
			os.Exit(1)
		}
		fmt.Println("File uploaded successfully. URL:", cos.GenerateMarkdownImageTag(objectURL))
		return
	}

	if *listBucket {
		err := cosClient.ListBucketObjects()
		if err != nil {
			fmt.Println("Error listing objects:", err)
			os.Exit(1)
		}
		return
	}

	if *deleteObjectKey != "" {
		err := cosClient.DeleteObject(*deleteObjectKey)
		if err != nil {
			fmt.Println("Error deleting object:", err)
			os.Exit(1)
		}
		return
	}

	fmt.Println("No operation specified. Use '-upload' flag to upload a file.")
}
