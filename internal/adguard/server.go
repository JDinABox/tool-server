package adguard

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/mailgun/groupcache/v2"
)

type AdguardServer struct {
	responseGroup *groupcache.Group
}

func New() *AdguardServer {
	s := new(AdguardServer)
	s.responseGroup = groupcache.NewGroup("responses", 5*1024*1024,
		groupcache.GetterFunc(func(ctx context.Context, key string, dest groupcache.Sink) error {
			b, err := s.fetchContent(ctx, key)
			if err != nil {
				return err
			}
			return dest.SetBytes(b, time.Now().Add(time.Hour))
		}))
	return s
}

func (s *AdguardServer) fetchContent(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var servicesByte []byte
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		services := AdguardServices{}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(body, &services); err != nil {
			return err
		}

		services.toMap()

		servicesByte, err = services.MarshalMsg(servicesByte)
		return err
	})

	return servicesByte, err
}

func (s *AdguardServer) AdguardContent(ctx context.Context) (*AdguardServices, error) {
	url := "https://adguardteam.github.io/HostlistsRegistry/assets/services.json"
	var respRaw []byte

	err := s.responseGroup.Get(ctx, url, groupcache.AllocatingByteSliceSink(&respRaw))
	if err != nil {
		return nil, err
	}

	services := new(AdguardServices)
	services.UnmarshalMsg(respRaw)

	return services, nil
}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	c := make(chan error, 1)
	req = req.WithContext(ctx)
	go func() { c <- f(http.DefaultClient.Do(req)) }()
	select {
	case <-ctx.Done():
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		return err
	}
}
