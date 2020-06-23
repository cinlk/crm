package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

func VerifyJwtTokenMiddleware(next http.Handler) http.Handler{

	verify := func(w http.ResponseWriter, r *http.Request) {


		log.Printf("get token %s \n", r.Header.Get("Authorization"))

		token, err :=  jwt.Parse(r.Header.Get("Authorization")[len("Bearer "):], func(token *jwt.Token) (i interface{}, e error) {
			return  []byte("demo"), nil
		})
		//token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (i interface{}, e error) {
		//		return  []byte("demo"), nil
		//})
		if err != nil{


			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if token.Valid == false{
			log.Println("token is invalid")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims);ok{
			log.Println(claims)
		}else{
			log.Println(err)
		}



		next.ServeHTTP(w, r)

	}


	return http.HandlerFunc(verify)
}
