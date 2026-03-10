package workbench

import "backend/datasource/dbdao"

type Domain struct {
	store *Store
}

func NewDomain(db *dbdao.DB) (*Domain, error) {
	repo, err := NewRepository(db)
	if err != nil {
		return nil, err
	}
	return &Domain{store: NewStore(repo)}, nil
}
