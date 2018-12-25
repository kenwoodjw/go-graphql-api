package server

import (
	"github.com/bradford-hamilton/go-graphql-api/gql"
	"encoding/json"
	"net/http"

	"github.com/kenwoodjw/go-graphql-api/gql"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"

)

type Server struct{
	GqlSchema *graphql.Schema	
}

type reqBody struct{
	Query string `json:"query"`

}

func (s *Server)GraphQL()http.HandleFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if r.Body == nil{
			http.Error(w,"Must provide graphql query in request body",400)
		}

		var rBody reqBody
		err := json.NewDecoder(r.Body).Decode(&rBody)

		if err != nil{
			http.Error(w,"Error parsing json Body",400))
		}

		result := gql.ExecuteQuery(rBody.Query, *s.GqlSchema)


		render.JSON(w,r,result)
	}

}