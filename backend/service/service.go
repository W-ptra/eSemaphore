package service

import "eSemaphore-backend/database"

type Service struct {
	db 			*database.Database
	jwtKey 		[]byte
	jwtExpired	int
}

func CreateService(db *database.Database, jwtSecret []byte, jwtExpired int) Service {
	return Service{
		db:db,
		jwtKey: jwtSecret,
		jwtExpired: jwtExpired,
	}
}