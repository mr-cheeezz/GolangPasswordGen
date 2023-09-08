package main

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func generatePassword(length int) string {
	rand.Seed(time.Now().UnixNano())

	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:'\",.<>/?"
	password := make([]byte, length)

	for i := 0; i < length; i++ {
		password[i] = chars[rand.Intn(len(chars))]
	}

	return string(password)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Password Generator")

	passwordLabel := widget.NewLabel("Generated Password:")
	passwordText := widget.NewEntry()
	passwordText.MultiLine = true

	generateButton := widget.NewButton("Generate Password", func() {
		password := generatePassword(12) // Change the length as needed
		passwordText.SetText(password)
	})

	content := container.NewVBox(
		passwordLabel,
		passwordText,
		generateButton,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}
