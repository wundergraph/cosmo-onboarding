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
	client   service.ProductsServiceClient
	cleanup  func()
}

func setupTestService(t *testing.T, products []service.Product) *testService {
	lis := bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()
	service.RegisterProductsServiceServer(grpcServer, &ProductsService{products: products})

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			t.Fatalf("failed to serve: %v", err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}
	conn, err := grpc.Dial(
		"passthrough:///bufnet",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)

	return &testService{
		grpcConn: conn,
		client:   service.NewProductsServiceClient(conn),
		cleanup: func() {
			conn.Close()
			grpcServer.Stop()
		},
	}
}

var testProducts = []service.Product{
	{
		Id:       "f47ac10b-58cc-4372-a567-0e02b2c3d479",
		Title:    "Test Product A",
		Price:    &service.Price{Amount: 10.00, Currency: *service.CURRENCY_CODE_CURRENCY_CODE_EUR.Enum()},
		Category: *service.ProductCategory_PRODUCT_CATEGORY_BOOKS.Enum(),
	},
	{
		Id:       "9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d",
		Title:    "Test Product B",
		Price:    &service.Price{Amount: 20.00, Currency: *service.CURRENCY_CODE_CURRENCY_CODE_USD.Enum()},
		Category: *service.ProductCategory_PRODUCT_CATEGORY_ELECTRONICS.Enum(),
	},
	{
		Id:       "1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed",
		Title:    "Test Product C",
		Price:    &service.Price{Amount: 30.00, Currency: *service.CURRENCY_CODE_CURRENCY_CODE_USD.Enum()},
		Category: *service.ProductCategory_PRODUCT_CATEGORY_CLOTHING.Enum(),
	},
}

func TestQueryProduct(t *testing.T) {
	svc := setupTestService(t, testProducts)
	defer svc.cleanup()

	tests := []struct {
		name    string
		id      string
		wantId  string
		wantNil bool
	}{
		{name: "returns product A", id: "f47ac10b-58cc-4372-a567-0e02b2c3d479", wantId: "f47ac10b-58cc-4372-a567-0e02b2c3d479"},
		{name: "returns product B", id: "9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d", wantId: "9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"},
		{name: "returns product C", id: "1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed", wantId: "1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed"},
		{name: "unknown id returns nil product", id: "00000000-0000-0000-0000-000000000000", wantNil: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := svc.client.QueryProduct(context.Background(), &service.QueryProductRequest{Id: tt.id})
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp.Product)
			} else {
				require.NotNil(t, resp.Product)
				assert.Equal(t, tt.wantId, resp.Product.Id)
			}
		})
	}
}

func TestQueryProductInjectedFixtures(t *testing.T) {
	single := []service.Product{testProducts[0]}
	svc := setupTestService(t, single)
	defer svc.cleanup()

	resp, err := svc.client.QueryProduct(context.Background(), &service.QueryProductRequest{Id: "f47ac10b-58cc-4372-a567-0e02b2c3d479"})
	require.NoError(t, err)
	require.NotNil(t, resp.Product)
	assert.Equal(t, "f47ac10b-58cc-4372-a567-0e02b2c3d479", resp.Product.Id)

	resp2, err := svc.client.QueryProduct(context.Background(), &service.QueryProductRequest{Id: "9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"})
	require.NoError(t, err)
	assert.Nil(t, resp2.Product)
}
