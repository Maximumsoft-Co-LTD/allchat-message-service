package aws

import (
	"allchat-message-service/internal/adapter/config"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Aws struct {
	session     *session.Session
	s3          *s3.S3
	AwsS3Bucket string
	domainUrlS3 string
}

func NewAws(config *config.AWS) (*Aws, error) {
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.AwsS3Region),
		Credentials: credentials.NewStaticCredentials(config.AwsAccessKeyID, config.AwsSecretAccessKey, ""),
	})
	if err != nil {
		log.Fatalf("Error creating session: %s", err)
		return nil, err
	}

	svc := s3.New(session)

	return &Aws{
		session,
		svc,
		config.AwsS3Bucket,
		config.DomainUrlS3,
	}, nil
}

func (a *Aws) UploadFile(fileName string, fileData []byte, path string, username string) (string, error) {
	if username == "" {
		username = "IMAGE"
	}
	//fileCompress := CompressImageResource(fileData)
	_, err := a.s3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(a.AwsS3Bucket),
		Key:    aws.String("CHATALLINONE/" + path + "/" + username + "/" + fileName),
		ACL:    aws.String("private"),
		Body:   bytes.NewReader(fileData),
		// ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(http.DetectContentType(fileData)),
		ContentDisposition:   aws.String(""),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		return "", err
	}

	urlpath := a.domainUrlS3 + "/CHATALLINONE/" + path + "/" + username + "/" + fileName
	return urlpath, err
}

func (a *Aws) DeleteFile(key string) (string, error) {
	key = strings.Replace(key, a.domainUrlS3+"/", "", -1)
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(a.AwsS3Bucket),
		Key:    aws.String(key),
	}
	_, err := a.s3.DeleteObject(input)
	if err != nil {
		fmt.Println("Error deleting object:", err)
		return "", err
	}
	return "success", nil
}
