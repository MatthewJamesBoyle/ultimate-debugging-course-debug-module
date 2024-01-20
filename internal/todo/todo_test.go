package todo_test

import (
	"context"
	"testing"
	"ultimate-debugging-course-debug-module/internal/todo"
)

func TestCreateAndGetTODO(t *testing.T) {
	// Setup
	ctx := context.Background()
	svc, err := todo.NewService()
	if err != nil {
		t.Fatal("unexpected error")
	}

	desc := "Test ToDo"
	err = svc.CreateTODO(ctx, desc)
	if err != nil {
		t.Errorf("CreateTODO() error = %v, wantErr %v", err, false)
	}

	got, err := svc.GetTODO(ctx, 1)
	if err != nil {
		t.Errorf("GetTODO() error = %v, wantErr %v", err, false)
	}

	if got.Description != desc {
		t.Errorf("GetTODO() got = %v, want %v", got.Description, desc)
	}

	// Check TODO not found.
	_, err = svc.GetTODO(ctx, 999)
	if err == nil {
		t.Error("GetTODO() expected error for non-existing ToDo, got no error")
	}
}
