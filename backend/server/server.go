package server

import (
	"connectrpc.com/connect"
	"context"
	"encoding/json"
	"fmt"
	"github.com/annp1987/sms-app/config"
	greet "github.com/annp1987/sms-app/proto"
	"go.uber.org/zap"
	"os"
	"strconv"
	"strings"
)

type GreetServer struct {
	log *zap.Logger
	db  map[string]*greet.PrefixMatch
}

func NewServer(conf *config.Config, log *zap.Logger) (*GreetServer, error) {
	jsonData, err := os.ReadFile(conf.DataFile)
	if err != nil {
		return nil, fmt.Errorf("read sample data file failed: %s", err.Error())
	}

	var objects []*greet.PrefixMatch

	err = json.Unmarshal(jsonData, &objects)
	if err != nil {
		return nil, fmt.Errorf("unmarshal json failed: %s", err.Error())
	}
	var db = make(map[string]*greet.PrefixMatch)

	for _, value := range objects {
		db[strconv.FormatInt(value.Prefix, 10)] = value
	}
	return &GreetServer{
		log: log,
		db:  db,
	}, nil

}

func (s *GreetServer) Greet(ctx context.Context, req *connect.Request[greet.SMSAnalysisRequest]) (*connect.Response[greet.SMSAnalysisResponse], error) {
	phoneNumber := strconv.FormatInt(req.Msg.Number, 10)
	var matchPrefix *greet.PrefixMatch
	maxLength := 0
	for prefix, value := range s.db {
		if strings.Contains(phoneNumber, prefix) {
			if len(prefix) > maxLength {
				matchPrefix = value
			}
			maxLength = len(prefix)
		}
	}
	res := connect.NewResponse(&greet.SMSAnalysisResponse{
		Prefix:  matchPrefix,
		Message: req.Msg.Message,
	})
	return res, nil
}
