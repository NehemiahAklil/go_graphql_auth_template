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

func Register(args models.RegisterArgs) (response models.RegisterOutput, err error) {
	hasuraResponse, err := execute(args)
	if err != nil {
		return
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	response = hasuraResponse.Data.Insert_users_one
	return
}

func execute(variables models.RegisterArgs) (response models.RegisterGraphQLResponse, err error) {
	// mapVariables := map[string]interface{}{
	// 	"first_name":   variables.First_name,
	// 	"last_name":    variables.Last_name,
	// 	"email":        variables.Email,
	// 	"password":     variables.Password,
	// }

	mapVariables, err := variables.ToMap()
	if err != nil {
		return
	}

	reqBody := models.GraphQLRequest{
		Query:     constants.Register,
		Variables: mapVariables,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}
	fmt.Println("We are in Register right now")
	hasuraURL := os.Getenv("HASURA_URL")
	resp, err := http.Post(hasuraURL, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return
	}

	return
}
