package login5

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	librespot "go-librespot"
	pb "go-librespot/proto/spotify/login5/v3"
	credentialspb "go-librespot/proto/spotify/login5/v3/credentials"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"net/url"
)

type Login5 struct {
	baseUrl *url.URL
	client  *http.Client

	loginOk *pb.LoginOk
}

func NewLogin5() *Login5 {
	baseUrl, err := url.Parse("https://login5.spotify.com/")
	if err != nil {
		panic("invalid apresolve base URL")
	}

	return &Login5{
		baseUrl: baseUrl,
		client:  &http.Client{},
	}
}

func (c *Login5) request(req *pb.LoginRequest) (*pb.LoginResponse, error) {
	body, err := proto.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed marhsalling LoginRequest: %w", err)
	}

	resp, err := c.client.Do(&http.Request{
		Method: "POST",
		URL:    c.baseUrl.JoinPath("/v3/login"),
		Header: http.Header{
			"Accept":     []string{"application/x-protobuf"},
			"User-Agent": []string{librespot.UserAgent()},
		},
		Body: io.NopCloser(bytes.NewReader(body)),
	})
	if err != nil {
		return nil, fmt.Errorf("failed requesting login5: %w", err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading login5 response: %w", err)
	}

	var protoResp pb.LoginResponse
	if err := proto.Unmarshal(respBody, &protoResp); err != nil {
		return nil, fmt.Errorf("faield unmarshalling LoginResponse: %w", err)
	}

	return &protoResp, nil
}

func (c *Login5) Login(credentials proto.Message, deviceId string) error {
	req := &pb.LoginRequest{
		ClientInfo: &pb.ClientInfo{
			ClientId: librespot.ClientId,
			DeviceId: deviceId,
		},
	}

	switch lm := credentials.(type) {
	case *credentialspb.StoredCredential:
		req.LoginMethod = &pb.LoginRequest_StoredCredential{StoredCredential: lm}
	case *credentialspb.Password:
		req.LoginMethod = &pb.LoginRequest_Password{Password: lm}
	case *credentialspb.FacebookAccessToken:
		req.LoginMethod = &pb.LoginRequest_FacebookAccessToken{FacebookAccessToken: lm}
	case *credentialspb.OneTimeToken:
		req.LoginMethod = &pb.LoginRequest_OneTimeToken{OneTimeToken: lm}
	case *credentialspb.ParentChildCredential:
		req.LoginMethod = &pb.LoginRequest_ParentChildCredential{ParentChildCredential: lm}
	case *credentialspb.AppleSignInCredential:
		req.LoginMethod = &pb.LoginRequest_AppleSignInCredential{AppleSignInCredential: lm}
	case *credentialspb.SamsungSignInCredential:
		req.LoginMethod = &pb.LoginRequest_SamsungSignInCredential{SamsungSignInCredential: lm}
	case *credentialspb.GoogleSignInCredential:
		req.LoginMethod = &pb.LoginRequest_GoogleSignInCredential{GoogleSignInCredential: lm}
	default:
		return fmt.Errorf("invalid credentials: %v", lm)
	}

	resp, err := c.request(req)
	if err != nil {
		return fmt.Errorf("failed requesting login5 endpoint: %w", err)
	}

	if ch := resp.GetChallenges(); ch != nil && len(ch.Challenges) > 0 {
		req.LoginContext = resp.LoginContext
		req.ChallengeSolutions = &pb.ChallengeSolutions{}

		// solve challenges
		for _, c := range ch.Challenges {
			switch cc := c.Challenge.(type) {
			case *pb.Challenge_Hashcash:
				sol := solveHashcash(req.LoginContext, cc.Hashcash)
				req.ChallengeSolutions.Solutions = append(req.ChallengeSolutions.Solutions, &pb.ChallengeSolution{
					Solution: &pb.ChallengeSolution_Hashcash{Hashcash: sol},
				})
			case *pb.Challenge_Code:
				return fmt.Errorf("login5 code challenge not supported")
			}
		}

		resp, err = c.request(req)
		if err != nil {
			return fmt.Errorf("failed requesting login5 endpoint with challenge solutions: %w", err)
		}
	}

	if ok := resp.GetOk(); ok != nil {
		c.loginOk = ok
		log.Debugf("authenticated as %s", c.loginOk.Username)
		return nil
	} else {
		return fmt.Errorf("failed authenticating with login5: %v", resp.GetError())
	}
}
