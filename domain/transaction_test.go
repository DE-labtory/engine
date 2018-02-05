package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"strconv"
)

func TestCreateNewTransactionTest(t *testing.T){
	tx := CreateNewTransaction(strconv.Itoa(1), strconv.Itoa(1), general, time.Now(), SetTxData("", invoke, SetTxMethodParameters(0, "", []string{""}), ""))
	assert.Equal(t, status_TRANSACTION_UNKNOWN, tx.TransactionStatus)
}

func TestSetTxMethodParameters(t *testing.T) {
	var par = SetTxMethodParameters(0, "void", []string{""})
	assert.Equal(t, 0, par.ParamsType)
}

func TestSetTxData(t *testing.T) {
	var txdata = SetTxData("temp", query, SetTxMethodParameters(0, "void", []string{""}), "111")
	assert.Equal(t, query, txdata.Method)
}

func TestTransaction_Validate(t *testing.T) {
	tx := CreateNewTransaction(strconv.Itoa(1), strconv.Itoa(1), general, time.Now(), SetTxData("", invoke, SetTxMethodParameters(0, "", []string{""}), ""))
	tx.GenerateHash()
	assert.Equal(t, true, tx.Validate())
}

func TestTransaction_SignHash(t *testing.T) {
	tx := CreateNewTransaction(strconv.Itoa(1), strconv.Itoa(1), general, time.Now(), SetTxData("", invoke, SetTxMethodParameters(0, "", []string{""}), ""))
	_, err := tx.SignHash()
	assert.NoError(t, err)
}


