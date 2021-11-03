package domain

import "time"

type Meow struct {
	ID        int
	Body      string
	CreatedOn time.Time
}
