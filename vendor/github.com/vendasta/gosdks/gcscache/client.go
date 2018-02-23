package gcscache

import (
	"context"
	"errors"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	"fmt"

	"github.com/vendasta/gosdks/logging"
)

//ErrCacheMiss is a duplicate from "golang.org/x/crypto/acme/autocert" to avoid
//dependency
var ErrCacheMiss = errors.New("acme/autocert: certificate cache miss")

//GCSCache uses a cloud storage bucket to cache data
// - It implements the autocert.Cache interface
type GCSCache struct {
	ProjectID          string
	Bucket             string
	ServiceAccountJSON string
	client             *storage.Client
	bucket             *storage.BucketHandle
}

//Get gets data from cache
func (m GCSCache) Get(ctx context.Context, key string) ([]byte, error) {
	logging.Debugf(ctx, "Getting GCS data (%s)", key)
	obj := m.bucket.Object(key)
	r, err := obj.NewReader(ctx)
	if err == storage.ErrObjectNotExist {
		logging.Debugf(ctx, "(Get) GCS File %s does not exist", key)
		return nil, ErrCacheMiss
	} else if err != nil {
		logging.Errorf(ctx, "GCS Get %s error: %s", key, err.Error())
		return nil, err
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}

//Put puts data into cache
func (m GCSCache) Put(ctx context.Context, key string, data []byte) (rv error) {
	logging.Debugf(ctx, "Putting GCS data (%s)", key)
	obj := m.bucket.Object(key)
	w := obj.NewWriter(ctx)
	defer func() {
		err := w.Close()
		if err != nil {
			logging.Errorf(ctx, "GCS Put %s error: %s", key, err.Error())
			rv = err
		}
	}()
	_, err := w.Write(data)
	if err != nil {
		logging.Errorf(ctx, "GCS Put %s error: %s", key, err.Error())
	}
	return err
}

//Delete removes data from cache
func (m GCSCache) Delete(ctx context.Context, key string) error {
	logging.Debugf(ctx, "Deleting GCS data (%s)", key)
	obj := m.bucket.Object(key)
	err := obj.Delete(ctx)
	if err == storage.ErrObjectNotExist {
		logging.Debugf(ctx, "(Delete) GCS File %s does not exist", key)
	}
	return err
}

func (m *GCSCache) getClient(ctx context.Context) error {
	var err error
	if m.ServiceAccountJSON != "" {
		ts, err2 := google.JWTConfigFromJSON([]byte(m.ServiceAccountJSON), "https://www.googleapis.com/auth/devstorage.read_write")
		if err2 != nil {
			msg := fmt.Sprintf("Failed to parse ServiceAccount JWT config: %s", err2.Error())
			logging.Errorf(ctx, "%s", msg)
			return fmt.Errorf("%s", msg)
		}
		m.client, err = storage.NewClient(ctx, option.WithTokenSource(ts.TokenSource(ctx)))
	} else {
		m.client, err = storage.NewClient(ctx)
	}
	if err != nil {
		msg := fmt.Sprintf("Failed to create GCS client: %s", err.Error())
		logging.Errorf(ctx, "%s", msg)
		return fmt.Errorf("%s", msg)
	}
	m.bucket = m.client.Bucket(m.Bucket)
	return nil
}

//NewGCSCache creates a new GCSCache struct for caching data in GCS
func NewGCSCache(ctx context.Context, projectID, bucket, serviceAccountJSON string) (*GCSCache, error) {
	c := GCSCache{Bucket: bucket, ProjectID: projectID, ServiceAccountJSON: serviceAccountJSON}
	err := c.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
