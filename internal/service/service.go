package service

type UrlMapService interface {
	GetUrl(hash string) (string, error)
	SaveUrl(url string) (string, error)
}

type Service struct {
	UrlMapService
}

func NewService(urlService UrlMapService) *Service {
	return &Service{
		urlService,
	}
}
