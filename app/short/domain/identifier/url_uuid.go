package identifier

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type UrlUuid struct {
	value uuid.UUID
}

func NewUrlUuid(url string, userId string) (urlId UrlUuid, err error) {
	reader := strings.NewReader(fmt.Sprintf("%s:%s", url, userId))
	value, err := uuid.NewRandomFromReader(reader)
	if err != nil {
		return UrlUuid{}, err
	}
	return UrlUuid{value: value}, nil
}

func (o UrlUuid) Value() uuid.UUID {
	return o.value
}
