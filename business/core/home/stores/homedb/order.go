package homedb

import (
	"fmt"

	"github.com/dmitryovchinnikov/blueprint/business/core/home"
	"github.com/dmitryovchinnikov/blueprint/business/data/order"
)

var orderByFields = map[string]string{
	home.OrderByID:     "home_id",
	home.OrderByType:   "type",
	home.OrderByUserID: "user_id",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
