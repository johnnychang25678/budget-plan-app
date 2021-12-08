package controllers

import (
	"budget-plan-app/backend/consts"
	"budget-plan-app/backend/repositories"
	"budget-plan-app/backend/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	clientId     string
	clientSecret string
	redirectUrl  string
	// memberRepo repositories.MemberRepo
}

func NewAuthController() *AuthController {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	// this redirectUrl is the endpoint after user consent
	redirectUrl := fmt.Sprintf("%s/auth/callback", consts.BaseUrl)
	return &AuthController{
		clientId:     clientId,
		clientSecret: clientSecret,
		redirectUrl:  redirectUrl,
	}
}

func (a *AuthController) HandleAuth(c *gin.Context) {
	googleConsentUrl := a.genGoogleConsentUrl(a.redirectUrl)
	c.Redirect(http.StatusFound, googleConsentUrl)
}

func (a *AuthController) HandleCallback(c *gin.Context) {
	// after consent from user, google will GET /signup/callback?code=xxx

	//  1. exchange code(Auth grant) with token from google
	// 	2. use google api + token to get user email
	//	3. check if user email already in db, if no, write user email & refreshToken to database
	//  4. issue a jwt to user so he have access to my app
	//  5. all subsequent requests need to have jwt in header, use middleware to auth
	// 	6. after jwt expires, user login with same google auth route
	//  7. Will use the same route, but this time will query user email in db
	//  8. issue a jwt right away

	if c.Query("error") != "" {
		c.JSON(400, gin.H{
			"error": c.Query("error"),
		})
		return
	}

	authGrantCode := c.Query("code")
	token, err := a.getAccessToken(authGrantCode)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	email, err := a.getUserEmail(token)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	memberRepo := repositories.NewMemberRepo()
	userId, err := memberRepo.FindByEmail(email)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("userId:", userId)

	if userId == 0 {
		memberRepo.Create(email)
	}

	accessTokenChan := make(chan utils.TokenResult)
	refreshTokenChan := make(chan utils.TokenResult)
	go utils.GenJwtToken(userId, utils.AccessToken, accessTokenChan)
	go utils.GenJwtToken(userId, utils.RefreshToken, refreshTokenChan)
	at, rt := <-accessTokenChan, <-refreshTokenChan

	if at.Err != nil || rt.Err != nil {
		fmt.Println("gen jwt token error")
		c.JSON(500, gin.H{
			"error": gin.H{
				"accessTokenErr":  at.Err.Error(),
				"refreshTokenErr": rt.Err.Error(),
			},
		})
		return
	}

	c.JSON(200, gin.H{
		"access_token":  at.Token,
		"refresh_token": rt.Token,
	})
}

func (a *AuthController) genGoogleConsentUrl(redirectUrl string) string {
	u, err := url.Parse(consts.GoogleOauthUrl)
	if err != nil {
		fmt.Println("parse url err:", err)
	}

	queryParams := url.Values{}

	queryParams.Set("client_id", a.clientId)
	queryParams.Set("redirect_uri", redirectUrl)
	queryParams.Set("scope", "email")
	queryParams.Set("response_type", "code")
	queryParams.Set("include_granted_scopes", "true")
	queryParams.Set("prompt", "consent")
	queryParams.Set("access_type", "offline")

	u.RawQuery = queryParams.Encode()
	fmt.Println(u)
	return u.String()
}

func (a *AuthController) getAccessToken(authGrantCode string) (string, error) {
	endpoint := consts.GoogleOauthTokenUrl

	postUrlForm := url.Values{}
	postUrlForm.Set("code", authGrantCode)
	postUrlForm.Set("client_id", a.clientId)
	postUrlForm.Set("client_secret", a.clientSecret)
	postUrlForm.Set("grant_type", "authorization_code")
	postUrlForm.Set("redirect_uri", a.redirectUrl)

	postBody := postUrlForm.Encode()

	type respBody struct {
		AccessToken string `json:"access_token"`
	}

	res, err := http.Post(
		endpoint,
		"application/x-www-form-urlencoded",
		strings.NewReader(postBody),
	)
	if err != nil {
		fmt.Println("post request error:", err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("post request error:", err)
		return "", err
	}

	var resp respBody
	json.Unmarshal(body, &resp)

	return resp.AccessToken, nil
}

func (a *AuthController) getUserEmail(accessToken string) (string, error) {
	endpoint := consts.GoogleGetUserEmailUrl + accessToken
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("get request error:", err)
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("get request error:", err)
		return "", err
	}

	type respBody struct {
		Email string `json:"email"`
	}

	var resp respBody
	json.Unmarshal(body, &resp)
	return resp.Email, nil
}
