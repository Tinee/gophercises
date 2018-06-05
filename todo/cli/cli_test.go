package cli

import (
	"log"
	"testing"

	"github.com/Tinee/gophercises/todo/domain"
)

func TestCli_Exectue(t *testing.T) {
	type fields struct {
		s      domain.TodoService
		log    *log.Logger
		Action Action
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cli{
				s:      tt.fields.s,
				log:    tt.fields.log,
				Action: tt.fields.Action,
			}
			if err := c.Exectue(); (err != nil) != tt.wantErr {
				t.Errorf("Cli.Exectue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
