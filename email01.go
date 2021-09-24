package main
import (
	"log"
	"net/smtp"
	"fmt"
)

func main() {

	//Config
	from := "shiek.alta3@gmail.com"
	password := "TacoTaco"
	to := [] string { "kmathani@gmail.com"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := [] byte("Subject: Now with a subject\r\n"+"Next message Alta3 Test lab")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from , to, message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email Sent")

}

