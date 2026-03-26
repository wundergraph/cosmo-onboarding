package main

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	service "github.com/wundergraph/cosmo/plugin/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

type testService struct {
	grpcConn *grpc.ClientConn
	client   service.ReviewsServiceClient
	cleanup  func()
}

func setupTestService(t *testing.T, reviews []service.Review, productReviews map[string][]string) *testService {
	lis := bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()
	service.RegisterReviewsServiceServer(grpcServer, &ReviewsService{
		reviews:        reviews,
		productReviews: productReviews,
	})

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			t.Logf("failed to serve: %v", err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}
	conn, err := grpc.NewClient(
		"passthrough:///bufnet",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)

	return &testService{
		grpcConn: conn,
		client:   service.NewReviewsServiceClient(conn),
		cleanup: func() {
			_ = conn.Close()
			grpcServer.Stop()
		},
	}
}

var testReviews = []service.Review{
	{Id: "a97a2e8b-3a6b-4b3a-9b5d-1e02b2c3d401", Author: "Test User A", Email: "test-a@example.com", CreatedOn: 1773653316, Rating: 5},
	{Id: "b87b3f9c-4c7c-5c4b-0c6e-2f13c3d4e502", Author: "Test User B", Email: "test-b@example.com", CreatedOn: 1773653400, Rating: 4},
	{Id: "c76c4g0d-5d8d-6d5c-1d7f-3g24d4e5f603", Author: "Test User C", Email: "test-c@example.com", CreatedOn: 1773653500, Rating: 3},
	{Id: "d65d5h1e-6e9e-7e6d-2e8g-4h35e5f6g704", Author: "Test User D", Email: "test-d@example.com", CreatedOn: 1773653600, Rating: 4},
	{Id: "e54e6i2f-7f0f-8f7e-3f9h-5i46f6g7h805", Author: "Test User E", Email: "test-e@example.com", CreatedOn: 1773653700, Rating: 3},
	{Id: "f43f7j3g-8g1g-9g8f-4g0i-6j57g7h8i906", Author: "Test User F", Email: "test-f@example.com", CreatedOn: 1773653800, Rating: 5},
}

var testProductReviews = map[string][]string{
	"product-1": {testReviews[0].Id, testReviews[1].Id, testReviews[2].Id},
	"product-2": {testReviews[3].Id},
	"product-3": {testReviews[4].Id, testReviews[5].Id},
}

func TestQueryReview(t *testing.T) {
	svc := setupTestService(t, testReviews, testProductReviews)
	defer svc.cleanup()

	tests := []struct {
		name    string
		id      string
		wantNil bool
	}{
		{name: "returns review A", id: testReviews[0].Id},
		{name: "returns review B", id: testReviews[1].Id},
		{name: "returns review C", id: testReviews[2].Id},
		{name: "unknown id returns nil review", id: "00000000-0000-0000-0000-000000000000", wantNil: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := svc.client.QueryReview(context.Background(), &service.QueryReviewRequest{Id: tt.id})
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp.Review)
			} else {
				require.NotNil(t, resp.Review)
				assert.Equal(t, tt.id, resp.Review.Id)
			}
		})
	}
}

func TestQueryReviewInjectedFixtures(t *testing.T) {
	single := testReviews[0:1]
	svc := setupTestService(t, single, nil)
	defer svc.cleanup()

	resp, err := svc.client.QueryReview(context.Background(), &service.QueryReviewRequest{Id: testReviews[0].Id})
	require.NoError(t, err)
	require.NotNil(t, resp.Review)
	assert.Equal(t, testReviews[0].Id, resp.Review.Id)

	resp2, err := svc.client.QueryReview(context.Background(), &service.QueryReviewRequest{Id: testReviews[1].Id})
	require.NoError(t, err)
	assert.Nil(t, resp2.Review)
}

func TestLookupProductById(t *testing.T) {
	svc := setupTestService(t, testReviews, testProductReviews)
	defer svc.cleanup()

	resp, err := svc.client.LookupProductById(context.Background(), &service.LookupProductByIdRequest{
		Keys: []*service.LookupProductByIdRequestKey{
			{Id: "product-1"},
			{Id: "product-2"},
			{Id: "product-3"},
		},
	})
	require.NoError(t, err)
	require.Len(t, resp.Result, 3)
	assert.Equal(t, "product-1", resp.Result[0].Id)
	assert.Len(t, resp.Result[0].Reviews.List.Items, 3)
	assert.Equal(t, "product-2", resp.Result[1].Id)
	assert.Len(t, resp.Result[1].Reviews.List.Items, 1)
	assert.Equal(t, "product-3", resp.Result[2].Id)
	assert.Len(t, resp.Result[2].Reviews.List.Items, 2)
}
