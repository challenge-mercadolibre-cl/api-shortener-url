package identifier

import (
	"fmt"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/exceptions"
)

type UserId struct {
	value string
}

func NewUserId(value string) (urlId UserId, err error) {
	urlId = UserId{value: value}
	if err = urlId.hasError(); err != nil {
		return UserId{}, err
	}

	return
}

func (o UserId) hasError() error {
	if len(o.Value()) == 0 {
		return fmt.Errorf("%w", exceptions.ErrUserIdEmpty)
	}
	return nil
}

func (o UserId) Value() string {
	return o.value
}
