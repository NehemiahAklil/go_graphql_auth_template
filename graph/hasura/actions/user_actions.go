package actions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	models "github.com/NehemiahAklil/minabtech-recipe-backend/graph/hasura/models"
	constants "github.com/NehemiahAklil/minabtech-recipe-backend/graph/utils/constants"
)

func SearchUser(args models.SearchUserArgs) (response *models.SearchUserOutput, err error) {

	hasuraResponse, err := executeSearchUser(args)

	// throw if any unexpected error happens
	if err != nil {
		return nil, nil
	}

	// delegate Hasura error
	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}
	fmt.Println(hasuraResponse)
	if len(hasuraResponse.Data.Users) == 0 {
		return nil, errors.New("Can't find user")
	}
	return &hasuraResponse.Data.Users[0], nil
	// return &hasuraResponse.Data.Users, nil

}
func executeSearchUser(variables models.SearchUserArgs) (response models.SearchUserGraphQLResponse, err error) {
	mapVariables := map[string]interface{}{
		"loginText": variables.LoginText,
	}

	// build the request body
	reqBody := models.GraphQLRequest{
		Query:     constants.SearchUser,
		Variables: mapVariables,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	fmt.Println("Request Body", string(reqBytes))
	hasuraURL := os.Getenv("HASURA_URL")
	resp, err := http.Post(hasuraURL, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	// parse the response
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println("Response Body", string(respBytes))
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return
	}

	// return the response
	return
}
