package helpers

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

const (
	projectID       = "vertical-shore-397515"       // FILL IN WITH YOURS
	bucketName      = "capstone_tickets_app_bucket" // FILL IN WITH YOURS
	DefaultFile     = "default.png"
	FileFetchBuyer  = "https://storage.googleapis.com/capstone_tickets_app_bucket/profile_picture/buyer/"
	FileFetchParner = "https://storage.googleapis.com/capstone_tickets_app_bucket/profile_picture/partner/"
	FileFetchEvent  = "https://storage.googleapis.com/capstone_tickets_app_bucket/profile_picture/event_banner/"
	BuyerPath       = "profile_picture/buyer/"
	PartnerPath     = "profile_picture/partner/"
	EventPath       = "event_banner/"
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

var Uploader *ClientUploader

func init() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./keys.json") // FILL IN WITH YOUR FILE PATH
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	Uploader = &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
		// uploadPath: "profile_picture/",
	}
}

func (c *ClientUploader) UploadFile(file multipart.File, object string, path string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(path + object).NewWriter(ctx)
	fmt.Println(wc)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}
