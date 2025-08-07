package database

import "github.com/krazy-code/devlink/queries"

type Queries struct {
	*queries.UserQueries
	*queries.AuthQueries
}

func OpenDBConnection() (*Queries, error) {
	pool, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		UserQueries: &queries.UserQueries{
			Pool: pool,
		},
		AuthQueries: &queries.AuthQueries{
			Pool: pool,
		},
	}, nil
}
