package gcs

import (
	"os"
	"testing"

	"github.com/fsouza/fake-gcs-server/fakestorage"
	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/log"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewBucket(t *testing.T) {
	serverMock, err := fakestorage.NewServerWithOptions(fakestorage.Options{
		Scheme: "http",
	})
	if err != nil {
		log.WithError(err)
	}
	defer serverMock.Stop()

	url := serverMock.URL()

	os.Setenv("STORAGE_EMULATOR_HOST", url)
	instance := NewGcsBucketManager("test")

	bucket, err := instance.CreateBucket("TEST_BUCKET")
	if err != nil {
		assert.Fail(t, "Should be able to create a bucket that does not currently exist")
	}
	assert.NotNil(t, bucket)
	assert.NotEqual(t, uuid.Nil, bucket)
	if err := os.Unsetenv("STORAGE_EMULATOR_HOST"); err != nil {
		log.WithError(err)
	}
}
