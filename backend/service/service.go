package service

import "eSemaphore-backend/database"

type Service struct {
	db 			*database.Database
	jwtKey 		[]byte
	jwtExpired	int
}

func CreateService(db *database.Database, jwtSecret string, jwtExpired int) Service {
	return Service{
		db:db,
		jwtKey: []byte(jwtSecret),
		jwtExpired: jwtExpired,
	}
}