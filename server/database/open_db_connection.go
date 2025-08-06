package database

import "github.com/krazy-code/devlink/queries"

type Queries struct {
	*queries.UserQueries
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
	}, nil
}
