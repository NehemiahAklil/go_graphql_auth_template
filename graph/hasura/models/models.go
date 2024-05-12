package hasura_models

import "encoding/json"

type GraphQLError struct {
	Message string `json:"message"`
}

type RegisterOutput struct {
	Id         *string
	First_name *string
	Last_name  *string
	Username   *string
	Email      *string
}

type SearchUserOutput struct {
	Id         string
	First_name string
	Last_name  string
	Username   string
	Email      string
	Password   string
}

type Mutation struct {
	Register *RegisterOutput
}

type RegisterArgs struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func (rg *RegisterArgs) ToMap() (map[string]interface{}, error) {
	// Marshal the struct to a byte slice
	data, err := json.Marshal(rg)
	if err != nil {
		return nil, err
	}

	// Unmarshal the byte slice into a map[string]interface{}
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type SearchUserArgs struct {
	LoginText string
}

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type RegisterGraphQLData struct {
	Insert_users_one RegisterOutput `json:"insert_users_one"`
}

type RegisterGraphQLResponse struct {
	Data   RegisterGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError      `json:"errors,omitempty"`
}

type SearchUserGraphQLData struct {
	Users []SearchUserOutput `json:"users"`
}
type SearchUserGraphQLResponse struct {
	Data   SearchUserGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError        `json:"errors,omitempty"`
}
