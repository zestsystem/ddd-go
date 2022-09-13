// Package Customer holds all the domain logic for the customer domain.
package customer

import (
    "errors"
    "github.com/google/uuid"
    "github.com/zestsystem/ddd-go/aggregate"
)

var (
    // ErrCustomerNotFound is returned when a customer is not found.
    ErrCustomerNotFound = errors.New("the customer was not found in the repository")
    // ErrFailedToAddCustomer is returned when the customer could not be added to the repository.
    ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
    // ErrUpdateCustomer is returned when the customer could not be updated in the repository.
    ErrUpdateCustomer = errors.New("failed to update the customer in the repository")
)

// CustomerRepository is an interface that defines the rules around what a customer
// Repository has to be able to perform

type CustomRepository interface {
    Get(uuid.UUID) (aggregate.Customer, error)
    Add(aggregate.Customer) error
    Update(aggregate.Customer) error
}
