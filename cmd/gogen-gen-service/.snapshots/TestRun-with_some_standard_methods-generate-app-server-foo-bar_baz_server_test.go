package server

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	foo_pb "testapp/api/foo"
)

func Test_BarBazServiceServer_ListBarBazs(t *testing.T) {
	svr := NewBarBazServiceServer()

	ctx := context.Background()
	req := &foo_pb.ListBarBazsRequest{}

	resp, err := svr.ListBarBazs(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_BarBazServiceServer_CreateBarBaz(t *testing.T) {
	svr := NewBarBazServiceServer()

	ctx := context.Background()
	req := &foo_pb.CreateBarBazRequest{}

	resp, err := svr.CreateBarBaz(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_BarBazServiceServer_DeleteBarBaz(t *testing.T) {
	svr := NewBarBazServiceServer()

	ctx := context.Background()
	req := &foo_pb.DeleteBarBazRequest{}

	resp, err := svr.DeleteBarBaz(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

