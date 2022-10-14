package identifier

import (
	"fmt"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/exceptions"
)

type UrlId struct {
	value string
}

func NewUrlId(value string) (urlId UrlId, err error) {
	urlId = UrlId{value: value}
	if err = urlId.hasError(); err != nil {
		return UrlId{}, err
	}

	return
}

func (o UrlId) hasError() error {
	if len(o.Value()) == 0 {
		return fmt.Errorf("%w", exceptions.ErrUrlIdEmpty)
	}
	return nil
}

func (o UrlId) Value() string {
	return o.value
}
