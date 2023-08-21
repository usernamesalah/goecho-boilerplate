package util

import (
	"encoding/json"
	"errors"
	"goecho-boilerplate/library/derrors"
	"goecho-boilerplate/library/logger"
	"net/http"

	"go.uber.org/zap"
)

type ResponseFormat struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Response converts a Go value to JSON and sends it to the client.
func Response(w http.ResponseWriter, data interface{}, status string, message string, httpCode int) error {

	// Convert the response value to JSON.
	res, err := json.Marshal(ResponseFormat{Status: status, Message: message, Data: data})
	if err != nil {
		return err
	}

	// Respond with the provided JSON.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpCode)
	if _, err := w.Write(res); err != nil {
		return err
	}

	return nil
}

// ResponseOK converts a Go value to JSON and sends it to the client.
func ResponseOK(w http.ResponseWriter, data interface{}, HTTPStatus int) error {
	return Response(w, data, StatusCodeOK, StatusMessageOK, HTTPStatus)
}

func ResponseSuccess(w http.ResponseWriter, data interface{}, message string, HTTPStatus int) error {
	return Response(w, data, StatusSuccess, message, HTTPStatus)
}

// ResponseError sends an error reponse back to the client.
func ResponseError(w http.ResponseWriter, err error) error {

	// If the error was of the type *Error, the handler has
	// a specific status code and error to return.
	if webErr, ok := err.(*Error); ok {
		if err := Response(w, nil, webErr.Status, webErr.MessageStatus, webErr.HTTPStatus); err != nil {
			return err
		}
		return nil
	}

	// If not, the handler sent any arbitrary error value so use 500.
	if err := Response(w, nil, StatusCodeInternalServerError, StatusMessageInternalServerError, http.StatusInternalServerError); err != nil {
		return err
	}
	return nil
}

// RenderErrorResponse sends an error reponse back to the client.
// Use this to get error response using derrors package
func RenderErrorResponse(w http.ResponseWriter, r *http.Request, err error) error {
	var ierr *derrors.Error
	if errors.As(err, &ierr) {
		status := derrors.ToStatus(ierr)
		msg := "Internal Server Error"
		if status < 500 {
			msg = ierr.Error()
		}

		reqID := ""
		// if reqIDCtx, ok := r.Context().Value(constant.RequestIDCtxKey).(string); ok {
		// 	reqID = reqIDCtx
		// }

		logger.GetL().Error(
			"API error",
			zap.Int("status", status),
			zap.String("msg", msg),
			zap.String("request_id", reqID),
			zap.Error(err),
		)
		if errResp := Response(w, nil, StatusError, msg, status); errResp != nil {
			return errResp
		}
		return nil
	}

	// If not, the handler sent any arbitrary error value so use 500.
	logger.GetL().Error(
		"API error",
		zap.Int("status", http.StatusInternalServerError),
		zap.String("orig", err.Error()),
	)
	if err := Response(w, nil, StatusCodeInternalServerError, StatusMessageInternalServerError, http.StatusInternalServerError); err != nil {
		return err
	}
	return nil
}
