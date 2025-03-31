package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    const fixedText  string = "Welcome to the Tech Palace, "
    var customerUpper string = strings.ToUpper(customer)
    var welcomeMessage string = fixedText + customerUpper
	return welcomeMessage
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    var stars string = strings.Repeat("*", numStarsPerLine)
	return stars + "\n" + welcomeMsg + "\n"  + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	lines := strings.Split(oldMsg, "\n")
	if len(lines) < 2 {
		return ""
	}
	// 2行目がメッセージ本体なので、前後のスペースとアスタリスクを除去
	messageLine := strings.TrimSpace(lines[1])
	messageLine = strings.Trim(messageLine, "* ")
	return messageLine
}
