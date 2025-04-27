package middleware

type Middleware struct{
	jwtSecret []byte
}

func GetMiddleware(jwtSecret []byte) Middleware {
	return Middleware{
		jwtSecret: jwtSecret,
	}
}