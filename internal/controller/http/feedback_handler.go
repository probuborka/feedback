package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/probuborka/feedback/internal/entity"
	"github.com/sirupsen/logrus"
)

type serviceFeedback interface {
	AddFeedback(ctx context.Context, feedback entity.Feedback) error
}

func (h handler) postFeedback(w http.ResponseWriter, r *http.Request) {
	//
	// requestID, ok := r.Context().Value(requestIDKey).(string)
	// if !ok {
	// 	requestID = "unknown"
	// }
	requestID := "unknown"

	// //
	var feedback entity.Feedback
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		h.response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest, requestID)
		h.log.WithFields(logrus.Fields{
			"requestID": requestID,
			"error":     err,
		}).Error("buf ReadFrom")
		return
	}

	err = json.Unmarshal(buf.Bytes(), &feedback)
	if err != nil {
		h.response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest, requestID)
		h.log.WithFields(logrus.Fields{
			"requestID": requestID,
			"error":     err,
		}).Error("unmarshal error")
		return
	}

	err = h.feedback.AddFeedback(r.Context(), feedback)
	if err != nil {
		h.response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest, requestID)
		h.log.WithFields(logrus.Fields{
			"requestID": requestID,
			"error":     err,
		}).Error("usecase recommendations")
		return
	}

	// //
	h.response(w, nil, http.StatusCreated, requestID)
}
