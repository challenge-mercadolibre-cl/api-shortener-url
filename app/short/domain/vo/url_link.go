package vo

import (
	"fmt"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/exceptions"
	"net/url"
)

type UrlLink struct {
	value string
}

func NewUrlLink(value string) (urlLink UrlLink, err error) {
	urlLink = UrlLink{value: value}
	if err = urlLink.hasError(); err != nil {
		return UrlLink{}, err
	}

	return
}

func (o UrlLink) hasError() error {
	_, err := url.ParseRequestURI(o.Value())
	if err != nil {
		return fmt.Errorf("%w: %s", exceptions.ErrUrlFormat, o.Value())
	}
	return nil
}

func (o UrlLink) Value() string {
	return o.value
}
