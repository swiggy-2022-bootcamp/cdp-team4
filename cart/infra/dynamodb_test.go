package infra_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Cart/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Cart/infra"
)

var testCartService = infra.NewDynamoRepository()
var insertedid = "853d6ab2-dcd9-49dc-8add-d93af655f8e0"
var inserteduserid string

func TestShouldCreateNewCartinDynamoDB(t *testing.T) {
	userid := uuid.New().String()
	inserteduserid = userid
	prodquant := map[string]int{
		"Origin of life":  1,
		"Reynolds trimax": 10,
	}

	newCart := domain.NewCart(userid, prodquant)
	res, err := testCartService.InsertCart(*newCart)
	insertedid = res
	fmt.Println(insertedid)
	t.Logf("Inserted Id is %s\n", insertedid)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestShouldGetAllCartDynamoDB(t *testing.T) {
	t.Logf("Inserted Id is %s Reading\n", insertedid)
	res, err := testCartService.FindAllCarts()
	fmt.Println(res)
	t.Logf("Read %v", res)
	assert.NotNil(t, res)
	assert.Nil(t, err)
	
}

func TestShouldGetCartByCartIdDynamoDB(t *testing.T) {
	prodquant := map[string]int{
		"Origin of life":  1,
		"Reynolds trimax": 10,
	}
	
	t.Logf("Inserted Id is %s Reading\n", insertedid)
	res, err := testCartService.FindCartById(insertedid)
	fmt.Println(res)
	t.Logf("Read %v", res)
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.Equal(t, res.ProductsQuantity, prodquant)
}

func TestShouldDeleteCartByCartIdDynamoDB(t *testing.T) {
	res, err := testCartService.DeleteCartById(insertedid)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}