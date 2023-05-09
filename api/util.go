package api

import (
	"Blog/core/token"
	"net/http"
)

const (
	authorizationHeader = "Authorization"
)

// read token
func (s *Server) readToken(r *http.Request) (*token.Payload, error) {
	token := r.Header.Get(authorizationHeader)
	payload, err := s.maker.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
