package cloudatcost

import (
	"net/http"
	"net/url"
)

type RunModeService struct {
	client *Client
}

func (s *RunModeService) Mode(serverId, action string) (*RunModeResponse, *http.Response, error) {
	u := "/api/v1/runmode.php"

	parameters := url.Values{}
	parameters.Add("key", s.client.Option.Key)
	parameters.Add("login", s.client.Option.Login)
	parameters.Add("sid", serverId)
	parameters.Add("mode", action)

	req, err := s.client.NewFormRequest("POST", u, parameters)
	if err != nil {
		return nil, nil, err
	}

	por := new(RunModeResponse)
	resp, err := s.client.Do(req, por)
	if err != nil {
		return nil, resp, err
	}
	return por, resp, err
}

func (s *RunModeService) Normal(serverId string) (*RunModeResponse, *http.Response, error) {
	return s.Mode(serverId, "normal")
}
func (s *RunModeService) Safe(serverId string) (*RunModeResponse, *http.Response, error) {
	return s.Mode(serverId, "safe")
}
