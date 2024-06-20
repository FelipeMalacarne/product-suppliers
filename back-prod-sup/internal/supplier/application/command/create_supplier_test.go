package command

import (
	"testing"

	"github.com/felipemalacarne/back-prod-sup/internal/supplier/domain"
	"github.com/felipemalacarne/back-prod-sup/internal/supplier/valueobject"
	"github.com/felipemalacarne/back-prod-sup/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandleSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockSupplierRepository(mockCtrl)
	handler := NewCreateSupplierHandler(mockRepo)

	command := CreateSupplierCommand{
		Name:  "John Doe",
		Email: "john@example.com",
		Phone: "1234567890",
	}

	email, _ := valueobject.NewEmail(command.Email)

	phone, _ := valueobject.NewPhone(command.Phone)

	expectedSupplier := domain.Supplier{
		Name:  command.Name,
		Email: email,
		Phone: phone,
	}

	// Set expectations
	mockRepo.EXPECT().Create(gomock.Any()).Return(expectedSupplier, nil)

	// Call the method under test
	supplier, err := handler.Handle(command)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedSupplier, supplier)
}

func TestHandleFailureOnDomainCreation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockSupplierRepository(mockCtrl)
	handler := NewCreateSupplierHandler(mockRepo)

	command := CreateSupplierCommand{
		Name:  "", // invalid input to trigger domain creation failure
		Email: "john@example.com",
		Phone: "1234567890",
	}

	// No expectation set on repository since creation should fail before reaching it

	// Call the method under test
	supplier, err := handler.Handle(command)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, domain.Supplier{}, supplier)
}
