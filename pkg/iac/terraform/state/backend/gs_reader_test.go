package backend

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"cloud.google.com/go/storage"
	pkghttp "github.com/cloudskiff/driftctl/pkg/http"
	"github.com/stretchr/testify/assert"
)

func TestGSBackend_Read(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  error
		client   StorageClient
		expected string
	}{
		{
			name: "Should fail with wrong URL",
			args: args{
				path: "",
			},
			wantErr: errors.New("Get \"wrong_url\": unsupported protocol scheme \"\""),
			client: func() StorageClient {
				return &storage.Client{}
			}(),
			expected: "",
		},
		{
			name: "Should fetch URL with auth header",
			args: args{
				path: "",
			},
			wantErr: nil,
			client: func() StorageClient {
				m := &pkghttp.MockHTTPClient{}

				req, _ := http.NewRequest(http.MethodGet, "https://example.com/cloudskiff/driftctl/main/terraform.tfstate", nil)

				req.Header.Add("Authorization", "Basic Test")

				bodyReader := strings.NewReader("{}")
				bodyReadCloser := io.NopCloser(bodyReader)

				m.On("Do", req).Return(&http.Response{
					StatusCode: 200,
					Body:       bodyReadCloser,
				}, nil)

				return m
			}(),
			expected: "{}",
		},
		{
			name: "Should fail with bad status code",
			args: args{
				path: "",
			},
			wantErr: errors.New("error requesting HTTP(s) backend state: status code: 404"),
			client: func() StorageClient {
				m := &pkghttp.MockHTTPClient{}

				req, _ := http.NewRequest(http.MethodGet, "https://example.com/cloudskiff/driftctl/main/terraform.tfstate", nil)

				bodyReader := strings.NewReader("test")
				bodyReadCloser := io.NopCloser(bodyReader)

				m.On("Do", req).Return(&http.Response{
					StatusCode: 404,
					Body:       bodyReadCloser,
				}, nil)

				return m
			}(),
			expected: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader, err := NewGSReader(tt.args.path)
			assert.NoError(t, err)

			got := make([]byte, len(tt.expected))
			_, err = reader.Read(got)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
				return
			} else {
				assert.NoError(t, err)
			}
			assert.NotNil(t, got)
			assert.Equal(t, tt.expected, string(got))
		})
	}
}

func TestGSBackend_Close(t *testing.T) {
	type fields struct {
		reader io.ReadCloser
		client *storage.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "should fail to close reader",
			fields: fields{
				reader: func() io.ReadCloser {
					return nil
				}(),
				client: &storage.Client{},
			},
			wantErr: true,
		},
		{
			name: "should close reader",
			fields: fields{
				reader: func() io.ReadCloser {
					m := &MockReaderMock{}
					m.On("Close").Return(nil)
					return m
				}(),
				client: &storage.Client{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &GSBackend{
				reader:   tt.fields.reader,
				GSClient: tt.fields.client,
			}
			if err := h.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
