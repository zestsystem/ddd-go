package valueobject

import (
	"github.com/google/uuid"
	"time"
)

// Transaction is a payment between two parties.

type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
