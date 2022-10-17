package identifier

import (
	"github.com/google/uuid"
)

type UrlUuid struct {
	value string
}

func NewUrlUuid(url string, userId string) (urlId UrlUuid, err error) {
	shortCode := uuid.New().String()[:8]
	return UrlUuid{value: shortCode}, nil
}

func (o UrlUuid) Value() string {
	return o.value
}
