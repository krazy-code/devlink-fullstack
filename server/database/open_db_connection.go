package database

import (
	"github.com/krazy-code/devlink/queries"
)

type Queries struct {
	*queries.UserQueries
	*queries.AuthQueries
	*queries.DeveloperQueries
	*queries.ProjectQueries
	*queries.SkillQueries
	*queries.FollowQueries
}

func OpenDBConnection() (*Queries, error) {
	pool, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	// ... (inside main or a function)
	// sqlScript, err := os.ReadFile("database/migrations/003_add_profile_fields_to_users.sql")
	// if err != nil {
	// 	log.Fatalf("Unable to read init.sql: %v\n", err)
	// }

	// if _, err := pool.Exec(context.Background(), string(sqlScript)); err != nil {
	// 	log.Fatalf("Error executing init.sql: %v\n", err)
	// }
	return &Queries{
		UserQueries: &queries.UserQueries{
			Pool: pool,
		},
		AuthQueries: &queries.AuthQueries{
			Pool: pool,
		},
		DeveloperQueries: &queries.DeveloperQueries{
			Pool: pool,
		},
		ProjectQueries: &queries.ProjectQueries{
			Pool: pool,
		},
		SkillQueries: &queries.SkillQueries{
			Pool: pool,
		},
		FollowQueries: &queries.FollowQueries{
			Pool: pool,
		},
	}, nil
}
