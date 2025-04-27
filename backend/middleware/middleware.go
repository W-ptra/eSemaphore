package middleware

type Middleware struct{
	jwtSecret []byte
}

func GetMiddleware(jwtSecret string) Middleware {
	return Middleware{
		jwtSecret: []byte(jwtSecret),
	}
}