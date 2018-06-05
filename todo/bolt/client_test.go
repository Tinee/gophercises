package bolt

import (
	"testing"

	"github.com/boltdb/bolt"
)

func TestClient_Open(t *testing.T) {
	type fields struct {
		Path        string
		db          *bolt.DB
		todoService TodoService
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"Hejsan",
			fields{

			},
			true
		},
		{
			"Hejsan",
			fields{

			},
			true
		}
		{
			"Hejsan",
			fields{

			},
			true
		}
		{
			"Hejsan",
			fields{

			},
			true
		}
		{
			"Hejsan",
			fields{

			},
			true
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Path:        tt.fields.Path,
				db:          tt.fields.db,
				todoService: tt.fields.todoService,
			}
			if err := c.Open(); (err != nil) != tt.wantErr {
				t.Errorf("Client.Open() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
