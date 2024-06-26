package order

import (
	"encoding/json"
	"net/http"
)

type Fabric struct {
}

func (f Fabric) CreateFromRequest(request *http.Request) (Order, error) {
	var requestOrder Order
	err := json.NewDecoder(request.Body).Decode(&requestOrder)
	if err != nil {
		return requestOrder, err
	}

	return requestOrder, nil
}
