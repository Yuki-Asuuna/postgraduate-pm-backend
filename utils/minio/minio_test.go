package minio

import "testing"

func TestCreateBucket(t *testing.T) {
	var err error
	err = MinioInit()
	if err != nil {
		t.Error(err)
	}
	err = createBucket(GetMinioClient(), "test")
	if err != nil {
		t.Error(err)
	}
}
