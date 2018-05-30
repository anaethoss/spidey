package graph

import (
	context "context"
	time "time"

	accountProto "github.com/tinrab/spidey/account/pb"
	catalogProto "github.com/tinrab/spidey/catalog/pb"
)

func (s *GraphQLServer) Order_account(ctx context.Context, obj *Order) (*Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	r, err := s.accountClient.GetAccount(
		ctx,
		&accountProto.GetAccountRequest{Id: obj.AccountID},
	)
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:   r.Account.Id,
		Name: r.Account.Name,
	}, nil
}

func (s *GraphQLServer) Order_products(ctx context.Context, obj *Order, skip *int, take *int) ([]Product, error) {
	skipValue, takeValue := parseRange(skip, take)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	r, err := s.catalogClient.GetProducts(
		ctx,
		&catalogProto.GetProductsRequest{Skip: skipValue, Take: takeValue},
	)
	if err != nil {
		return nil, err
	}

	products := []Product{}
	for _, a := range r.Products {
		products = append(
			products,
			Product{
				ID:          a.Id,
				Name:        a.Name,
				Description: a.Description,
				Price:       a.Price,
			},
		)
	}

	return products, nil
}

func (s *GraphQLServer) Query_accounts(ctx context.Context, skip *int, take *int, id *string) ([]Account, error) {
	// Get single
	if id != nil {
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		r, err := s.accountClient.GetAccount(
			ctx,
			&accountProto.GetAccountRequest{Id: *id},
		)
		if err != nil {
			return nil, err
		}
		return []Account{Account{
			ID:   r.Account.Id,
			Name: r.Account.Name,
		}}, nil
	}

	skipValue, takeValue := parseRange(skip, take)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	r, err := s.accountClient.GetAccounts(
		ctx,
		&accountProto.GetAccountsRequest{Skip: skipValue, Take: takeValue},
	)
	if err != nil {
		return nil, err
	}

	accounts := []Account{}
	for _, a := range r.Accounts {
		accounts = append(accounts, Account{ID: a.Id, Name: a.Name})
	}

	return accounts, nil
}

func (s *GraphQLServer) Query_products(ctx context.Context, skip *int, take *int, id *string) ([]Product, error) {
	// Get single
	if id != nil {
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		r, err := s.catalogClient.GetProduct(
			ctx,
			&catalogProto.GetProductRequest{Id: *id},
		)
		if err != nil {
			return nil, err
		}
		return []Product{Product{
			ID:          r.Product.Id,
			Name:        r.Product.Name,
			Description: r.Product.Description,
			Price:       r.Product.Price,
		}}, nil
	}

	skipValue, takeValue := parseRange(skip, take)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	r, err := s.catalogClient.GetProducts(
		ctx,
		&catalogProto.GetProductsRequest{Skip: skipValue, Take: takeValue},
	)
	if err != nil {
		return nil, err
	}

	products := []Product{}
	for _, a := range r.Products {
		products = append(
			products,
			Product{
				ID:          a.Id,
				Name:        a.Name,
				Description: a.Description,
				Price:       a.Price,
			},
		)
	}

	return products, nil
}

func parseRange(skip *int, take *int) (uint64, uint64) {
	skipValue := uint64(0)
	takeValue := uint64(100)

	if skip != nil {
		skipValue = uint64(*skip)
	}
	if take != nil {
		takeValue = uint64(*take)
	}

	return skipValue, takeValue
}
