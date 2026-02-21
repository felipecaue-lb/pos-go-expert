package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, error := CalculateTax(1000.0)
	assert.Nil(t, error)
	assert.Equal(t, 10.0, tax)

	tax, error = CalculateTax(0)
	assert.Error(t, error, "amount must be greater than zero")
	assert.Equal(t, 0.0, tax)
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil).Twice()
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	error := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, error)

	error = CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, error)

	error = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, error, "error saving tax")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 3)
}
