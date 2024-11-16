package helper

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

type Params struct {
	From     string
	Password string
	To       string
	Message  string
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
