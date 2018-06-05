package cli

import (
	"bytes"
	"log"
	"testing"

	"github.com/Tinee/gophercises/todo/mocks"

	"github.com/Tinee/gophercises/todo/domain"
)

func TestCli_Exectue(t *testing.T) {
	var (
		buf bytes.Buffer
		l   = log.New(&buf, "", 0)
	)

	tests := []struct {
		name    string
		log     *log.Logger
		wantErr bool
		Action  Action
		prepare func(*mock.TodoService)
	}{
		{
			"Should fail if we pass down an invalid argument",
			l,
			true,
			Action{Kind: "this is an invalid argument"},
			func(svc *mock.TodoService) {},
		},
		{
			"Should not throw any errors and output some items to the loggers writer.",
			l,
			false,
			Action{Arg: "", Kind: "list"},
			func(svc *mock.TodoService) {
				svc.AllFn = func() ([]domain.Todo, error) {
					return []domain.Todo{{Message: "TestData1", ID: 1}, {Message: "TestData2", ID: 2}}, nil
				}
			},
		},
		{
			"Should error out because of an invalid argument, it's need to be an int.",
			l,
			true,
			Action{Arg: "asdThisShouldFail", Kind: "do"},
			func(svc *mock.TodoService) {},
		},
		{
			"Should error out because TodoService.Delete returns an error.",
			l,
			true,
			Action{Arg: "1", Kind: "do"},
			func(svc *mock.TodoService) {
				svc.DeleteTodoFn = func(id int) error {
					return domain.ErrCreate
				}
			},
		},
		{
			"Should not fail when a successful do commands occurs",
			l,
			false,
			Action{Arg: "1", Kind: "do"},
			func(svc *mock.TodoService) {
				svc.DeleteTodoFn = func(id int) error {
					return nil
				}
			},
		},
		{
			"Should not fail when a successful create command occurs",
			l,
			false,
			Action{Kind: "create"},
			func(svc *mock.TodoService) {
				svc.CreateTodoFn = func(id domain.Todo) error {
					return nil
				}
			},
		},
		{
			"Should fail if TodoService.Create errors out",
			l,
			true,
			Action{Kind: "create"},
			func(svc *mock.TodoService) {
				svc.CreateTodoFn = func(id domain.Todo) error {
					return domain.ErrCreate
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todoSvcMock := mock.TodoService{}
			tt.prepare(&todoSvcMock)

			c := New(
				todoSvcMock,
				tt.log,
				tt.Action,
			)

			if err := c.Exectue(); (err != nil) != tt.wantErr {
				t.Errorf("Cli.Exectue() error = %v, wantErr %v", err, tt.wantErr)
			}

			// if !strings.Contains(buf.String(), tt.writeExpect) {
			// 	t.Errorf("The Cli's logger expected a write value of %v but got %v", tt.writeExpect, buf.String())
			// }
		})
	}
}
