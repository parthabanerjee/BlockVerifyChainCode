package main

import (
	"encoding/json"
	"errors"
	"fmt"
	//"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)

// Block verify is a high level smart contract for block verify soltion.

type BVP struct {

}

type ProductDetails struct{	
	ProductName string `json:"productname"`
	ProductSource string `json:"productsource"`
	ProductMineDate string `json:"productminedate"`
	ProductWareHouseDate string `json:"productwarehousedate"`
	ProductColour string `json:"productcolour"`
	ProductFinishDate string `json:"productfinishdate"`
	LastProductOwner string `json:"lastproductowner"`
}

// Transaction is for storing transaction Details
type Transaction struct{	
	TrxId string `json:"trxId"`
	TimeStamp string `json:"timeStamp"`
	ProductName string `json:"productname"`
	Source string `json:"source"`
	ProductOwner string `json:"productowner"`
	Remarks string `json:"remarks"`
}

type GetOwner struct{	
	LastProductOwner string `json:"lastproductowner"`
}

// to return the verify result
type VerifyU struct{	
	Result string `json:"result"`
}

// Init initializes the smart contracts
func (t *BVP) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists
	_, err := stub.GetTable("ProductDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("ProductDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "ProductName", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "ProductSource", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ProductMineDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ProductWareHouseDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ProductColour", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ProductFinishDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "LastProductOwner", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ProductDetails.")
	}
	
	// Check if table already exists
	_, err = stub.GetTable("Transaction")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("Transaction", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "trxId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "timeStamp", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "productName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "source", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "productOwner", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "remarks", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating Transaction table.")
	}
		
	// setting up the users role
	stub.PutState("user_type1_1", []byte("application1"))
	stub.PutState("user_type1_2", []byte("application2"))
	stub.PutState("user_type1_3", []byte("application3"))
	stub.PutState("user_type1_4", []byte("application4"))	
	
	return nil, nil
}

//addProducts to add a product
func (t *BVP) addProducts(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

//if len(args) != 12 {
	//		return nil, fmt.Errorf("Incorrect number of arguments. Expecting 12. Got: %d.", len(args))
	//	}

		
		productName:=args[0]
		productsource:=args[1]
		productminedate:=args[2]
		productwarehousedate:=args[3]
		productcolour:=args[4]
		productfinishdate:=args[5]
	//	lastproductowner:=args[6]
				
		//assignerOrg1, err := stub.GetState(args[6])
		//assignerOrg := string(assignerOrg1)
		
		createdBy:=args[6]
		// This is for first time creation.
		lastproductowner:="Application1"

		fmt.Printf("Value of created by: %s",createdBy)

		// Insert a row
		ok, err := stub.InsertRow("ProductDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: productName}},
				&shim.Column{Value: &shim.Column_String_{String_: productsource}},
				&shim.Column{Value: &shim.Column_String_{String_: productminedate}},
				&shim.Column{Value: &shim.Column_String_{String_: productwarehousedate}},
				&shim.Column{Value: &shim.Column_String_{String_: productcolour}},
				&shim.Column{Value: &shim.Column_String_{String_: productfinishdate}},
				&shim.Column{Value: &shim.Column_String_{String_: lastproductowner}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}

//get All transaction against ffid (irrespective of org)
func (t *BVP) getAllTransaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting product name to query")
	}

	productName := args[0]
	//assignerRole := args[1]

	var columns []shim.Column

	rows, err := stub.GetRows("Transaction", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	
	//assignerOrg1, err := stub.GetState(assignerRole)
	//assignerOrg := string(assignerOrg1)
	
		
	res2E:= []*Transaction{}
	
	for row := range rows {		
		newApp:= new(Transaction)
		newApp.TrxId = row.Columns[0].GetString_()
		newApp.TimeStamp = row.Columns[1].GetString_()
		newApp.ProductName = row.Columns[2].GetString_()
		newApp.Source = row.Columns[3].GetString_()
		newApp.ProductOwner = row.Columns[4].GetString_()
		newApp.Remarks = row.Columns[5].GetString_()
		
		
		if newApp.ProductName == productName{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

// Insert the transaction(irrespective of org)
func (t *BVP) addProductHistory(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	trxId := args[0]
	timeStamp:=args[1]
	productName := args[2]
	
	//assignerOrg1, err := stub.GetState(args[3])
	source := args[3]
	trxntype := args[4]
	productOwner := args[5]
	remarks := args[6]
	
	fmt.Printf("Last Product Owner: %s", trxntype)
	// Get the row pertaining to this product name
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: productName}}
	columns = append(columns, col1)

	
	row, err := stub.GetRow("ProductDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving product with product name %s. Error %s", productName, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}

	err = stub.DeleteRow(
		"ProductDetails",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}

	
	productsource := row.Columns[1].GetString_()
	productminedate := row.Columns[2].GetString_()
	productwarehousedate := row.Columns[3].GetString_()
	productcolour := row.Columns[4].GetString_()
	productfinishdate := row.Columns[5].GetString_()
	lastproductowner := args[7]
	fmt.Printf("Last Product Owner: %s", lastproductowner)

	// Insert a row
	ok, err := stub.InsertRow("ProductDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: productName}},
			&shim.Column{Value: &shim.Column_String_{String_: productsource}},
			&shim.Column{Value: &shim.Column_String_{String_: productminedate}},
			&shim.Column{Value: &shim.Column_String_{String_: productwarehousedate}},
			&shim.Column{Value: &shim.Column_String_{String_: productcolour}},
			&shim.Column{Value: &shim.Column_String_{String_: productfinishdate}},
			&shim.Column{Value: &shim.Column_String_{String_: lastproductowner}},
		}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}


		//inserting the transaction
		ok, err = stub.InsertRow("Transaction", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: trxId}},
				&shim.Column{Value: &shim.Column_String_{String_: timeStamp}},
				&shim.Column{Value: &shim.Column_String_{String_: productName}},
				&shim.Column{Value: &shim.Column_String_{String_: source}},
				&shim.Column{Value: &shim.Column_String_{String_: productOwner}},
				&shim.Column{Value: &shim.Column_String_{String_: remarks}},				
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}		
	return nil, nil

}



// Invoke invokes the chaincode
func (t *BVP) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "addProducts" {
		t := BVP{}
		return t.addProducts(stub, args)	
	} else if function == "addProductHistory" { 
		t := BVP{}
		return t.addProductHistory(stub, args)
	}

	return nil, errors.New("Invalid invoke function name.")

}

// query queries the chaincode
func (t *BVP) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getAllTransaction" {
		t := BVP{}
		return t.getAllTransaction(stub, args)
	}
	
	return nil, nil
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(BVP))
	if err != nil {
		fmt.Printf("Error starting BVP: %s", err)
	}
} 
