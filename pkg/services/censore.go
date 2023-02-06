package services

import (
	"context"
	"github.com/RTS-1989/go-censor-svc/pkg/db"
	"github.com/RTS-1989/go-censor-svc/pkg/models"
	"github.com/RTS-1989/go-censor-svc/pkg/pb"
	textParser "github.com/RTS-1989/go-censor-svc/pkg/testParser"
	"net/http"
)

type Server struct {
	H db.Handler
	C pb.CensorServiceClient
}

func (s *Server) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {

	censored := make([]*models.CensoredWords, 0)
	query := `SELECT * FROM censored_words;`
	if result := s.H.DB.Raw(query).Find(&censored); result.Error != nil {
		return &pb.CreateCommentResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	cs := textParser.CheckCensored(req.Text, censored)
	req.Censored = cs

	createComment, err2 := s.C.CreateComment(ctx, req)
	if err2 != nil {
		return &pb.CreateCommentResponse{
			Status: http.StatusConflict,
			Error:  err2.Error(),
		}, nil
	}

	return &pb.CreateCommentResponse{
		Status: http.StatusCreated,
		Id:     createComment.Id,
	}, nil
}
