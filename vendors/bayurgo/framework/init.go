package framework

import "net/http"

type M map[string]interface{}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ok := true
		if !ok {
			w.Write([]byte(`something went wrong`))
			return
		}

		isValid := true
		if !isValid {
			w.Write([]byte(`wrong username/password`))
			return
		}

		next.ServeHTTP(w, r)
	})
}
