package usersummarydb

import (
	"fmt"

	"github.com/dmitryovchinnikov/blueprint/business/core/usersummary"
	"github.com/dmitryovchinnikov/blueprint/business/data/order"
)

var orderByFields = map[string]string{
	usersummary.OrderByUserID:   "user_id",
	usersummary.OrderByUserName: "user_name",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
