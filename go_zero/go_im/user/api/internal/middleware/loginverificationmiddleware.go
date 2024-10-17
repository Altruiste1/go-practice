package middleware

import "net/http"

type LoginVerificationMiddleware struct {
}

func NewLoginVerificationMiddleware() *LoginVerificationMiddleware {
	return &LoginVerificationMiddleware{}
}

func (m *LoginVerificationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
		return
		// TODO generate middleware implement function, delete after code implementation
		if r.Header.Get("token") != "" { // todo add jwt
			next(w, r)
			return
		}
		// Passthrough to next handler if need
		w.Write([]byte("权限不足，无法执行"))
		return
	}
}
