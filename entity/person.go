package entity

import (
    "github.com/google/uuid"
)

// Person is an entity that represents a person in all Domains
type Person struct {
    // ID is the identifier of the Entity. The ID is shared for all subdomains.
    ID uuid.UUID
    Name string
    Age int
}


