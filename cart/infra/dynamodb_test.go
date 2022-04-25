package infra_test

import (
	"fmt"
	"testing"

	//"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	//"github.com/swiggy-2022-bootcamp/cdp-team4/cart/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/infra"
)

var testCartService = infra.NewDynamoRepository()
var insertedid string
var inserteduserid string

// func TestShouldCreateNewCartinDynamoDB(t *testing.T) {
// 	userid := uuid.New().String()
// 	inserteduserid = userid
// 	products :=(map[string]domain.Item{})
// 	products["1"]=domain.Item{Name:"pen",Cost:10,Quantity:1}
// 	products["2"]=domain.Item{Name:"pencil",Cost:5,Quantity:2}

// 	newCart := domain.NewCart(userid, products)
// 	res , err := testCartService.InsertCart(*newCart)
// 	insertedid = res
// 	t.Logf("Inserted Id is %s\n", insertedid)
// 	assert.NotNil(t, res)
// 	assert.Nil(t, err)
// }

func TestShouldGetAllCartDynamoDB2(t *testing.T) {
	t.Logf("Inserted Id is %s Reading\n", insertedid)
	res, err := testCartService.FindAllCarts()
	fmt.Println(res)
	t.Logf("Read %v", res)
	assert.NotNil(t, res)
	assert.Nil(t, err)

}

// func TestShouldDeleteCartitemByUserIdDynamoDB(t *testing.T) {
	
// 	productsIdList := []string{"1","3","4"}
// 	res , err := testCartService.DeleteCartItemByUserId("P",productsIdList)
// 	assert.NotNil(t, res)
// 	assert.Equal(t,res,true)
// 	assert.Nil(t, err)
// }

func TestShouldDeleteCartByUserIdDynamoDB(t *testing.T) {
	res , err := testCartService.DeleteCartByUserId("P")
	assert.NotNil(t, res)
	assert.Equal(t,res,true)
	assert.Nil(t, err)
}

// func TestShouldUpdateCartByUserIdDynamoDB(t *testing.T) {
	
// 	//inserteduserid = userid
// 	products :=(map[string]domain.Item{})
// 	//products["1"]=domain.Item{Name:"pen",Cost:10,Quantity:1}
// 	products["2"]=domain.Item{Name:"pencil",Cost:5,Quantity:20}

// 	//newCart := domain.NewCart(userid, products)
// 	res , err := testCartService.UpdateCartByUserId("Swapnil",products)
// 	assert.NotNil(t, res)
// 	assert.Equal(t,res,true)
// 	assert.Nil(t, err)
// }

func TestShouldGetAllCartDynamoDB(t *testing.T) {
	t.Logf("Inserted Id is %s Reading\n", insertedid)
	res, err := testCartService.FindAllCarts()
	fmt.Println(res)
	t.Logf("Read %v", res)
	assert.NotNil(t, res)
	assert.Nil(t, err)

}

// func TestShouldGetCartByCartIdDynamoDB(t *testing.T) {
// 	products :=(map[string]domain.Item{})
// 	products["1"]=domain.Item{Name:"pen",Cost:10,Quantity:1}
// 	products["2"]=domain.Item{Name:"pencil",Cost:5,Quantity:2}

// 	t.Logf("Inserted Id is %s Reading\n", insertedid)
// 	res, err := testCartService.FindCartById(insertedid)
// 	fmt.Println(res)
// 	t.Logf("Read %v", res)
// 	assert.NotNil(t, res)
// 	assert.Nil(t, err)
// 	assert.Equal(t, res.Items, products)
// }

// func TestShouldDeleteCartByCartIdDynamoDB(t *testing.T) {
// 	res, err := testCartService.DeleteCartById("dcf58f49-fae7-46b1-85ec-0011c1ef1d98")
// 	assert.NotNil(t, res)
// 	assert.Nil(t, err)
// }

// func TestShouldGetAllCartDynamoDB1(t *testing.T) {
// 	t.Logf("Inserted Id is %s Reading\n", insertedid)
// 	res, err := testCartService.FindAllCarts()
// 	fmt.Println(res)
// 	t.Logf("Read %v", res)
// 	assert.NotNil(t, res)
// 	assert.Nil(t, err)
// }