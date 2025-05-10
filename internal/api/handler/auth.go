package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/lai0xn/isdb/internal/api"
	"github.com/lai0xn/isdb/internal/app/auth"
	"github.com/lai0xn/isdb/pkg/oath"
	"github.com/lai0xn/isdb/pkg/utils"
)

// LoginRequest defines the structure for login request body
type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"securepassword123"`
}

// SignupRequest defines the structure for signup request body
type SignupRequest struct {
	Email      string `json:"email" example:"user@example.com"`
	Password   string `json:"password" example:"securepassword123"`
	Name       string `json:"name" example:"John"`
	FamilyName string `json:"family_name" example:"Doe"`
	Age        int    `json:"age" example:"25"`
}

// VerificationEmailRequest defines request for sending verification email
type VerificationEmailRequest struct {
	Email string `json:"email" example:"user@example.com"`
}

// ActivationRequest defines request for user activation
type ActivationRequest struct {
	Email string `json:"email" example:"user@example.com"`
	OTP   string `json:"otp" example:"123456"`
}

// Login handles user login
// @Summary User login
// @Description Authenticate user and return access tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} api.Map
// @Failure 400 {object} api.Map
// @Router /login [post]
func (a *API) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}

	loginReq := auth.LoginUserRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	tokenPair, err := a.authSrv.LoginUser(r.Context(), loginReq)
	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, errors.New("invalid credentials"))
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, api.Map{
		"message": "login successful",
		"data":    tokenPair,
	})
}

// Signup handles user registration
// @Summary User registration
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body SignupRequest true "User registration data"
// @Success 201 {object} api.Map
// @Failure 400 {object} api.Map
// @Router /signup [post]
func (a *API) Signup(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}

	signupReq := auth.SignupUserRequest{
		Email:      req.Email,
		Password:   req.Password,
		Name:       req.Name,
		FamilyName: req.FamilyName,
		Age:        req.Age,
	}

	if err := a.authSrv.SignupUser(r.Context(), signupReq); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}
  

	utils.WriteJSONResponse(w, http.StatusCreated, api.Map{
		"message": "user created successfully",
		"user": api.Map{
			"email": signupReq.Email,
			"name":  signupReq.Name,
		},
	})
}

// SendVerificationEmail sends verification email to user
// @Summary Send verification email
// @Description Send email with verification code to user's email address
// @Tags auth
// @Accept json
// @Produce json
// @Param request body VerificationEmailRequest true "User email"
// @Success 200 {object} api.Map
// @Failure 400 {object} api.Map
// @Router /send-verification [post]
func (a *API) SendVerificationEmail(w http.ResponseWriter, r *http.Request) {
	var payload VerificationEmailRequest
  
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, fmt.Errorf("couldn't send verification email: %v", err))
		return 
	}

	user, err := a.userSrv.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, fmt.Errorf("couldn't send verification email: %v", err))
		return 
	}
  
	err = a.authSrv.SendVerification(*user)
	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, fmt.Errorf("couldn't send verification email: %v", err))
		return 
	}
  
	utils.WriteJSONResponse(w, http.StatusOK, api.Map{
		"message": "email verification sent",
	})
}

// AcivateUser activates user account with OTP
// @Summary Activate user account
// @Description Verify user's email with OTP code
// @Tags auth
// @Accept json
// @Produce json
// @Param request body ActivationRequest true "Email and OTP"
// @Success 200 {object} api.Map
// @Failure 400 {object} api.Map
// @Router /activate [post]
func (a *API) AcivateUser(w http.ResponseWriter, r *http.Request) {
	var payload ActivationRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, fmt.Errorf("couldn't verify email: %v", err))
		return 
	}

	user, err := a.userSrv.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, fmt.Errorf("couldn't verify email: %v", err))
		return 
	}

	err = a.authSrv.Verify(*user, payload.OTP)
	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, fmt.Errorf("couldn't verify email: %v", err))
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, api.Map{
		"message": "email verified",
	})
}

// GoogleLogin initiates Google OAuth flow
// @Summary Google OAuth login
// @Description Redirects to Google OAuth consent screen
// @Tags auth
// @Success 302 "Redirect to Google OAuth"
// @Router /google/login [get]
func (a *API) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := oath.GoogleOath().AuthCodeURL("randomState")
	http.Redirect(w, r, url, http.StatusFound) 
}

// GoogleCallback handles Google OAuth callback
// @Summary Google OAuth callback
// @Description Handles Google OAuth callback and returns user info
// @Tags auth
// @Param code query string true "OAuth code from Google"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} api.Map
// @Failure 500 {object} api.Map
// @Router /google/callback [get]
func (a *API) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		utils.WriteJSONError(w, http.StatusBadRequest, fmt.Errorf("empty code"))
		return
	}
  google := oath.GoogleOath()
	token, err := google.Exchange(context.Background(), code)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, fmt.Errorf("failed to exchange token"))
		return
	}

	client := google.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, fmt.Errorf("failed to get user info"))
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, fmt.Errorf("failed to parse user info"))
		return
	}




	utils.WriteJSONResponse(w, http.StatusOK, userInfo)
}
