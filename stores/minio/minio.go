package minio

import minio "github.com/minio/minio-go"

// NewConnection creates a new minio Client object
func NewConnection(host, port, accessKeyID, secretAccessKey string, useSSL bool) *minio.Client {
	client, err := minio.New(host+port, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		// handle error
		return nil
	}
	return client
}
