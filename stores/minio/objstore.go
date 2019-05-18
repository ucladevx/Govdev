package minio

import (
	"io"
	"time"

	minio "github.com/minio/minio-go"
)

// MinioObjectStore wraps a Minio Client
type MinioObjectStore struct {
	client *minio.Client
}

// MinioBucket stores the name and location of a bucket
type MinioBucket struct {
	store    *minio.Client
	name     string
	location string
}

// MinioObjectInfo wraps minio.ObjectInfo{}
type MinioObjectInfo struct {
	lastModified time.Time
	eTag         string
	contentType  string
	size         int64
}

func NewMinioObjStore(client *minio.Client) *MinioObjectStore {
	return &MinioObjectStore{
		client: client,
	}
}

// GetBucket returns a bucket based on name and location strings. If a bucket
// does not yet exist, it will create that bucket.
func (mos *MinioObjectStore) GetBucket(name, location string) (*MinioBucket, error) {
	exists, err := mos.client.BucketExists(name)
	if err != nil {
		return nil, err
	}

	if !exists {
		err := mos.client.MakeBucket(name, location)
		if err != nil {
			return nil, err
		}
	}

	bucket := &MinioBucket{
		store:    mos.client,
		name:     name,
		location: location,
	}
	return bucket, nil
}

// DestroyBucket destroys a bucket based on name and location
func (mos *MinioObjectStore) DestroyBucket(name, location string) error {
	exists, err := mos.client.BucketExists(name)
	if err != nil {
		return err
	}
	if exists {
		if err := mos.client.RemoveBucket(name); err != nil {
			return err
		}
	}
	return nil
}

func (mb *MinioBucket) Get(name string) (io.Reader, error) {
	obj, err := mb.store.GetObject(mb.name, name, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (mb *MinioBucket) Stat(name string) (*MinioObjectInfo, error) {
	stats, err := mb.store.StatObject(mb.name, name, minio.StatObjectOptions{})
	if err != nil {
		return nil, err
	}
	return New(&stats), nil
}

func (mb *MinioBucket) GetStat(name string) (io.Reader, *MinioObjectInfo, error) {
	obj, err := mb.Get(name)
	if err != nil {
		return nil, nil, err
	}

	stats, err := mb.Stat(name)
	if err != nil {
		return nil, nil, err
	}

	return obj, stats, nil
}

func (mb *MinioBucket) Put(name, contentType string, size int64, object io.Reader) error {
	_, err := mb.store.PutObject(
		mb.name, name, object, size,
		minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}
	return nil
}

func (mb *MinioBucket) Remove(name string) error {
	if err := mb.store.RemoveObject(mb.name, name); err != nil {
		return err
	}
	return nil
}

func New(mio *minio.ObjectInfo) *MinioObjectInfo {
	return &MinioObjectInfo{
		lastModified: mio.LastModified,
		eTag:         mio.ETag,
		contentType:  mio.ContentType,
		size:         mio.Size,
	}
}

func (oi *MinioObjectInfo) LastModified() time.Time {
	return oi.lastModified
}

func (oi *MinioObjectInfo) ETag() string {
	return oi.eTag
}

func (oi *MinioObjectInfo) ContentType() string {
	return oi.contentType
}

func (oi *MinioObjectInfo) Size() int64 {
	return oi.size
}
