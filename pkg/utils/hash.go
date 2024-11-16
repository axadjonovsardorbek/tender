package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"log/slog"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"regexp"
	"text/template"

	"github.com/axadjonovsardorbek/tender/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Params struct {
	From     string
	Password string
	To       string
	Message  string
}

const emailRegex = `^[a-zA-Z0-9._]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func ClaimData(c *gin.Context, data string) string {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return ""
	}

	res := claims.(jwt.MapClaims)[data].(string)

	return res
}

func SmsSender(c *gin.Context, err error, code int) {
	url := c.Request.URL
	message := fmt.Sprint("Status code: ", code, "\nEndpoint: ", url, "\nError: ", err)
	err = SendMessage(message)
	if err == nil {
		slog.Info("Error is successfully sent to group by bot")
	}
}

func SendMessage(message string) error {
	cf := config.Load()
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cf.BotToken)
	params := url.Values{}
	params.Add("chat_id", cf.GroupId)
	params.Add("text", message)

	// So'rov yuborish
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Javobni o'qish
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	// Status kodni tekshirish
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}

func SendVerificationCode(params Params) error {
	// Read and parse the HTML file
	htmlFile, err := os.ReadFile("format.html")
	if err != nil {
		log.Println("Cannot read html file:", err.Error())
		return err
	}

	// Parse the HTML template
	temp, err := template.New("email").Parse(string(htmlFile))
	if err != nil {
		log.Println("Cannot parse html file:", err.Error())
		return err
	}

	// Apply parameters to the HTML template
	var body bytes.Buffer
	err = temp.Execute(&body, params)
	if err != nil {
		log.Println("Cannot execute HTML template:", err.Error())
		return err
	}

	// Construct the email headers and body
	message := "From: " + params.From + "\n" +
		"To: " + params.To + "\n" +
		"Subject: Verification Email\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\n\n" +
		body.String()

	// Send the email
	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", params.From, params.Password, "smtp.gmail.com"),
		params.From, []string{params.To}, []byte(message),
	)

	if err != nil {
		log.Println("Could not send email:", err.Error())
		return err
	}

	return nil
}

// HashPassword hashes a plaintext password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPasswordHash compares a plaintext password with its hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
