package database

func DBconnection() (*Queries, error) {
	db, err := OpenDBConnection()
	if err != nil {
		// return utils.ResponseParser(c, utils.Response{
		// 	Code:   fiber.StatusInternalServerError,
		// 	Errors: err.Error(),
		// })
		return nil, err
	}
	return db, nil
}
