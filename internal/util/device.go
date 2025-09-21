package util

import "github.com/mssola/user_agent"

func DetectDevice(uaString string) string {
	ua := user_agent.New(uaString)

	if ua.Bot() {
		return "bot"
	}

	if ua.Mobile() {
		return "mobile"
	}

	if ua.Platform() == "iPad" || ua.Platform() == "Android" && !ua.Mobile() {
		return "tablet"
	}

	return "desktop"
}
