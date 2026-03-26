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
			reviews:        fixtures,
			productReviews: fixtureProductReviews,
		})
	}, routerplugin.WithTracing())

	if err != nil {
		log.Fatalf("failed to create router plugin: %v", err)
	}

	pl.Serve()
}

type ReviewsService struct {
	service.UnimplementedReviewsServiceServer
	reviews        []service.Review
	productReviews map[string][]string
}

func (s *ReviewsService) QueryReview(ctx context.Context, req *service.QueryReviewRequest) (*service.QueryReviewResponse, error) {
	for _, r := range s.reviews {
		if r.Id == req.Id {
			review := r
			return &service.QueryReviewResponse{Review: &review}, nil
		}
	}
	return &service.QueryReviewResponse{}, nil
}

func (s *ReviewsService) LookupProductById(ctx context.Context, req *service.LookupProductByIdRequest) (*service.LookupProductByIdResponse, error) {
	reviewByID := make(map[string]*service.Review, len(s.reviews))
	for i := range s.reviews {
		reviewByID[s.reviews[i].Id] = &s.reviews[i]
	}
	result := make([]*service.Product, len(req.Keys))
	for i, key := range req.Keys {
		ids := s.productReviews[key.Id]
		items := make([]*service.Review, 0, len(ids))
		for _, id := range ids {
			if r, ok := reviewByID[id]; ok {
				items = append(items, r)
			}
		}
		result[i] = &service.Product{
			Id: key.Id,
			Reviews: &service.ListOfReview{
				List: &service.ListOfReview_List{
					Items: items,
				},
			},
		}
	}
	return &service.LookupProductByIdResponse{Result: result}, nil
}

var fixtures = []service.Review{
	{
		Id:        "review-1",
		Author:    "Alice",
		Email:     "alice@example.com",
		CreatedOn: 1773653316,
		Contents:  wrapperspb.String("Excellent product, highly recommended!"),
		Rating:    5,
	},
	{
		Id:        "review-2",
		Author:    "Bob",
		Email:     "bob@example.com",
		CreatedOn: 1773653400,
		Contents:  wrapperspb.String("Good value for money"),
		Rating:    4,
	},
	{
		Id:        "review-3",
		Author:    "Carol",
		Email:     "carol@example.com",
		CreatedOn: 1773653500,
		Contents:  wrapperspb.String("Average product, does the job"),
		Rating:    3,
	},
	{
		Id:        "review-4",
		Author:    "Dave",
		Email:     "dave@example.com",
		CreatedOn: 1773653600,
		Contents:  wrapperspb.String("Fast shipping, works as described"),
		Rating:    4,
	},
	{
		Id:        "review-5",
		Author:    "Eve",
		Email:     "eve@example.com",
		CreatedOn: 1773653700,
		Contents:  wrapperspb.String("Decent quality for the price"),
		Rating:    3,
	},
	{
		Id:        "review-6",
		Author:    "Frank",
		Email:     "frank@example.com",
		CreatedOn: 1773653800,
		Contents:  wrapperspb.String("Would buy again"),
		Rating:    5,
	},
}

var fixtureProductReviews = map[string][]string{
	"product-1": {"review-1", "review-2", "review-3"},
	"product-2": {"review-4"},
	"product-3": {"review-5", "review-6"},
}
