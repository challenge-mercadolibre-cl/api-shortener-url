package aggregations

import (
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/identifier"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/vo"
)

type Url struct {
	userId  identifier.UserId
	urlId   identifier.UrlId
	urlLink vo.UrlLink
}

func NewUrl(link string, userId string) (Url, error) {
	userIdIdentifier, err := identifier.NewUserId(userId)
	if err != nil {
		return Url{}, err
	}

	urlUuid, err := identifier.NewUrlUuid(link, userIdIdentifier.Value())
	if err != nil {
		return Url{}, err
	}

	urlId, err := identifier.NewUrlId(urlUuid.Value().String())
	if err != nil {
		return Url{}, err
	}
	urlLink, err := vo.NewUrlLink(link)
	if err != nil {
		return Url{}, err
	}
	url := Url{userId: userIdIdentifier, urlId: urlId, urlLink: urlLink}
	return url, nil
}

func (u Url) UserId() identifier.UserId {
	return u.userId
}

func (u Url) UrlId() identifier.UrlId {
	return u.urlId
}

func (u Url) UrlLink() vo.UrlLink {
	return u.urlLink
}
