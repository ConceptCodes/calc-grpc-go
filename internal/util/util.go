package util

import (
	"calc/pb"

	"github.com/rs/zerolog"
)

func LogRequest(target string, req *pb.CalculationRequest, log zerolog.Logger) {
	log.Debug().Int64("A", req.A).Int64("B", req.B).Msgf("Incoming %s request", target)
}

func LogResponse(target string, res *pb.CalculationResponse, log zerolog.Logger) {
	log.Debug().Int64("result", res.Result).Msgf("Result for %s is %d", target, res.Result)
}

func LogError(target string, err error, log zerolog.Logger) {
	log.Err(err).Msgf("Error w/ %s", target)
}
