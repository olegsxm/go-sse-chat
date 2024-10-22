package http_errors

import "net/http"

type ChatHTTPError interface {
	Code() int
}

type UserNotFound struct{}

func (err UserNotFound) Error() string { return "User not found" }

func (err UserNotFound) Code() int { return http.StatusNotFound }

type InvalidCredentials struct{}

func (err InvalidCredentials) Error() string { return "Invalid credentials" }
func (err InvalidCredentials) Code() int     { return http.StatusUnauthorized }

type UnknownError struct{}

func (err UnknownError) Error() string { return "Unknown error" }
func (err UnknownError) Code() int     { return http.StatusInternalServerError }

type BadRequest struct{}

func (err BadRequest) Error() string { return "Bad request" }
func (err BadRequest) Code() int     { return http.StatusBadRequest }

type InternalServerError struct{}

func (err InternalServerError) Error() string { return "Internal server error" }
func (err InternalServerError) Code() int     { return http.StatusInternalServerError }

type UserAlreadyExists struct{}

func (err UserAlreadyExists) Error() string { return "User already exists" }
func (err UserAlreadyExists) Code() int     { return http.StatusConflict }
