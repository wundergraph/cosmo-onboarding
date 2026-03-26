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
      products: fixtures,
    })
  }, routerplugin.WithTracing())

  if err != nil {
    log.Fatalf("failed to create router plugin: %v", err)
  }

  pl.Serve()
}

type ProductsService struct {
	service.UnimplementedProductsServiceServer
	products []service.Product
}

func (s *ProductsService) QueryProduct(ctx context.Context, req *service.QueryProductRequest) (*service.QueryProductResponse, error) {
	for _, p := range s.products {
		if p.Id == req.Id {
			product := p
			return &service.QueryProductResponse{Product: &product}, nil
		}
	}
	return &service.QueryProductResponse{}, nil
}

var fixtures = []service.Product{
	{
		Id:          "product-1",
		Title:       "Product 1",
		Description: wrapperspb.String("The best product in the world"),
		Price: &service.Price{
			Amount:   1.125,
			Currency: *service.CURRENCY_CODE_CURRENCY_CODE_EUR.Enum(),
		},
		Category: *service.ProductCategory_PRODUCT_CATEGORY_BOOKS.Enum(),
	},
	{
		Id:          "product-2",
		Title:       "Product 2",
		Description: wrapperspb.String("A great electronics item"),
		Price: &service.Price{
			Amount:   299.99,
			Currency: *service.CURRENCY_CODE_CURRENCY_CODE_USD.Enum(),
		},
		Category: *service.ProductCategory_PRODUCT_CATEGORY_ELECTRONICS.Enum(),
	},
	{
		Id:          "product-3",
		Title:       "Product 3",
		Description: wrapperspb.String("Stylish clothing for all occasions"),
		Price: &service.Price{
			Amount:   49.95,
			Currency: *service.CURRENCY_CODE_CURRENCY_CODE_USD.Enum(),
		},
		Category: *service.ProductCategory_PRODUCT_CATEGORY_CLOTHING.Enum(),
	},
}
