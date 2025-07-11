package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"upc-backend-sparkathon/main/models"
)

func FetchProductfromUPC(upc string) (*models.Product, error) {

	apiUrl := fmt.Sprintf("https://api.upcitemdb.com/prod/trial/lookup?upc=%s", upc)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result models.UPCItemDBResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.Total == 0 || len(result.Items) == 0 {
		return nil, errors.New("no product found for this UPC")
	}

	return &result.Items[0], nil

}
