package util

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"goecho-boilerplate/internal/dto"
	"goecho-boilerplate/library/config"
	"math/rand"
	"time"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	encodingBase        = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	orderedEncodingBase = "0123456789aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"
	FormatDate1         = "2006-01-02 15:04:05"
	FormatDate2         = "Monday, 02 January 2006 15:04"
	FormatDate3         = "02 January 2006"
	FormatDate4         = "2006-01-02T15:04:05-0700"
	FormatDate5         = "01/02/2006"
	IssuerJWT           = "husen"
)

var letters = []rune(encodingBase)

func ComparePasswords(hashedPwd string, plainPwd []byte) (err error) {
	byteHash := []byte(hashedPwd)
	err = bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		err = errors.New("400: Your password is incorrect")
		return
	}
	return
}

func HashAndSalt(pwd []byte) (hashPwd string, err error) {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return
	}

	hashPwd = string(hash)
	return
}

func GetDateByUnixTime(format string, unixtime int) string {
	if unixtime <= (7 * 3600) {
		return ""
	}
	result := time.Unix(int64(unixtime), 0).Format(format)
	return result
}

func GenerateSha256Hash(value string) string {
	hasher := sha256.New()
	hasher.Write([]byte(value))
	signature := hex.EncodeToString(hasher.Sum(nil))
	return signature
}

func CalculateOffset(limit, currentPage int) (result int) {
	result = (currentPage - 1) * limit
	return
}

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Now() int {
	return int(time.Now().Unix())
}

func SetPaginator(limit, page int, cnt int64, c echo.Context) (pageData *dto.PageData) {
	p := pagination.NewPaginator(c.Request(), limit, cnt)
	// because this beego pagination package use 'p' instad of 'page' for page param
	// the p.Pages() is invalid, so we use our own hasNext and currentPage
	hasNext := page < p.PageNums()
	pageData = &dto.PageData{
		HasNext:     hasNext,
		TotalData:   p.Nums(),
		TotalPages:  p.PageNums(),
		CurrentPage: page,
		Limit:       p.PerPageNums,
	}
	return
}

func GetUserID(c echo.Context) (uid string) {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	return claims["uid"].(string)
}

func GenerateToken(uid, name string) (token string, err error) {
	rawToken := jwt.NewWithClaims(jwt.SigningMethodRS256, dto.JWTClaims{})
	token, err = rawToken.SignedString(config.Get().JWTRS256PrivateKey)
	return
}
