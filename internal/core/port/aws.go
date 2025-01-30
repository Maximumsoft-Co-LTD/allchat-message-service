package port

type AWSRepository interface {
	UploadFile(fileName string, fileData []byte, path string, username string) (string, error)
}
