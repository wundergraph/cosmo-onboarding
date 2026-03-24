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
    s.RegisterService(&service.ReviewsService_ServiceDesc, &ReviewsService{
    })
  }, routerplugin.WithTracing())

  if err != nil {
    log.Fatalf("failed to create router plugin: %v", err)
  }

  pl.Serve()
}

type ReviewsService struct {
  service.UnimplementedReviewsServiceServer
}

func (s *ReviewsService) QueryReview(ctx context.Context, req *service.QueryReviewRequest) (*service.QueryReviewResponse, error) {
  response := &service.QueryReviewResponse{
		Review: &service.Review{
			Id: "xyz",
			Author: "cosmo@wundergraph.com",
			CreatedOn: 1773653316,
			Contents: wrapperspb.String("Such a nice a product"),
			Rating: 5,
		},
  }
  return response, nil
}

func (s *ReviewsService) LookupProductById(ctx context.Context, req *service.LookupProductByIdRequest) (*service.LookupProductByIdResponse, error) {
	result := make([]*service.Product, len(req.Keys))
	for i, key := range req.Keys {
		result[i] = &service.Product{
			Id: key.Id,
			Reviews: &service.ListOfReview{
				List: &service.ListOfReview_List{
					Items: []*service.Review{
						{
							Id:       "xyz",
							Author:   "cosmo@wundergraph.com",
							Rating:   5,
							Contents: wrapperspb.String("Such a nice a product"),
						},
					},
				},
			},
		}
	}
	return &service.LookupProductByIdResponse{Result: result}, nil
}
