package adguard

import (
	"errors"
	"strings"
)

//go:generate msgp -io=false -tests=false

// AdguardService
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
	s.MappedServices = make(map[string]AdguardService, len(s.BlockedServices))
	for _, v := range s.BlockedServices {
		s.MappedServices[v.ID] = v
	}
}

type FormatEnum int

const (
	FormatWildcard FormatEnum = iota
	FormatAdp
)

func (s *AdguardServices) ServiceList(serviceN string, format FormatEnum) (string, error) {
	service, ok := s.MappedServices[serviceN]
	if !ok {
		return "", ErrServiceNotFound
	}
	var output strings.Builder
	for _, rule := range service.Rules {
		switch format {
		case FormatAdp:
			output.WriteString(rule)
		default:
			ruleLine := strings.Split(strings.TrimSpace(rule), "\n")
			// convert adp to wildcard
			for _, line := range ruleLine {
				line = strings.TrimSpace(line)
				line = strings.ReplaceAll(line, "||", "")
				line = strings.ReplaceAll(line, "^", "")
				if i := strings.Index(line, "*."); i != -1 {
					line = line[strings.Index(line, "*.")+2:]
				}
				output.WriteString("*.")
				output.WriteString(line)
			}
		}
		output.WriteRune('\n') // Add newline after each rule
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
