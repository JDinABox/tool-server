package adguard

import (
	"errors"
	"strings"
)

//go:generate msgp -io=false -tests=false

type AdguardService struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Rules   []string `json:"rules"`
	IconSvg string   `json:"icon_svg"`
}

type AdguardServices struct {
	BlockedServices []AdguardService `json:"blocked_services"`
	MappedServices  map[string]AdguardService
}

var ErrServiceNotFound = errors.New("not found")

func (s *AdguardServices) toMap() {
	s.MappedServices = make(map[string]AdguardService)
	for _, v := range s.BlockedServices {
		s.MappedServices[v.ID] = v
	}
}

func (s *AdguardServices) ServiceList(serviceN string) (string, error) {
	service, ok := s.MappedServices[serviceN]
	if !ok {
		return "", ErrServiceNotFound
	}
	var output strings.Builder
	for _, v := range service.Rules {
		output.WriteString(v)
		output.WriteRune('\n')
	}
	return output.String(), nil
}

func (s *AdguardServices) MappedNames() map[string]string {
	m := make(map[string]string, len(s.MappedServices))
	for k, v := range s.MappedServices {
		m[k] = v.Name
	}
	return m
}
