package main

import (
	"context"
	"log"

	service "github.com/wundergraph/cosmo/plugin/generated"

	routerplugin "github.com/wundergraph/cosmo/router-plugin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
  pl, err := routerplugin.NewRouterPlugin(func(s *grpc.Server) {
    s.RegisterService(&service.ProductsService_ServiceDesc, &ProductsService{
    })
  }, routerplugin.WithTracing())

  if err != nil {
    log.Fatalf("failed to create router plugin: %v", err)
  }

  pl.Serve()
}

type ProductsService struct {
	service.UnimplementedProductsServiceServer
}

func (s *ProductsService) QueryProduct(ctx context.Context, req *service.QueryProductRequest) (*service.QueryProductResponse, error) {
  response := &service.QueryProductResponse{
		Product: &service.Product{
			Id: "1",
			Title: "Product 1",
			Description: wrapperspb.String("The best product in the world"),
			Price: &service.Price{
				Amount: 1.125,
				Currency: *service.CURRENCY_CODE_CURRENCY_CODE_EUR.Enum(),
			},
			Category: *service.ProductCategory_PRODUCT_CATEGORY_BOOKS.Enum(),
		},
  }
  return response, nil
}
