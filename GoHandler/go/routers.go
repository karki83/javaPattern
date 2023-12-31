/*
 * dev-serverless-slack-bot
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2022-11-30T03:41:31Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/{basePath}/",
		Index,
	},

	Route{
		"AddleaveOptions",
		strings.ToUpper("Options"),
		"/{basePath}/addleave",
		AddleaveOptions,
	},

	Route{
		"ApiMessagesOptions",
		strings.ToUpper("Options"),
		"/{basePath}/api/messages",
		ApiMessagesOptions,
	},

	Route{
		"AuthSlackCallbackOptions",
		strings.ToUpper("Options"),
		"/{basePath}/auth/slack/callback",
		AuthSlackCallbackOptions,
	},

	Route{
		"AuthTeamsOptions",
		strings.ToUpper("Options"),
		"/{basePath}/auth/teams",
		AuthTeamsOptions,
	},

	Route{
		"AuthyOptions",
		strings.ToUpper("Options"),
		"/{basePath}/authy",
		AuthyOptions,
	},

	Route{
		"AwayOptions",
		strings.ToUpper("Options"),
		"/{basePath}/away",
		AwayOptions,
	},

	Route{
		"DelTemplateOptions",
		strings.ToUpper("Options"),
		"/{basePath}/delTemplate",
		DelTemplateOptions,
	},

	Route{
		"DisconnectOptions",
		strings.ToUpper("Options"),
		"/{basePath}/disconnect",
		DisconnectOptions,
	},

	Route{
		"HelpOptions",
		strings.ToUpper("Options"),
		"/{basePath}/help",
		HelpOptions,
	},

	Route{
		"HistoryOptions",
		strings.ToUpper("Options"),
		"/{basePath}/history",
		HistoryOptions,
	},

	Route{
		"NotificationOptions",
		strings.ToUpper("Options"),
		"/{basePath}/notification",
		NotificationOptions,
	},

	Route{
		"OooOptions",
		strings.ToUpper("Options"),
		"/{basePath}/ooo",
		OooOptions,
	},

	Route{
		"PollResponseOptions",
		strings.ToUpper("Options"),
		"/{basePath}/pollResponse",
		PollResponseOptions,
	},

	Route{
		"PollsOptions",
		strings.ToUpper("Options"),
		"/{basePath}/polls",
		PollsOptions,
	},

	Route{
		"SendMessageOptions",
		strings.ToUpper("Options"),
		"/{basePath}/sendMessage",
		SendMessageOptions,
	},

	Route{
		"SlackActionsOptions",
		strings.ToUpper("Options"),
		"/{basePath}/slack/actions",
		SlackActionsOptions,
	},

	Route{
		"SlackEventsOptions",
		strings.ToUpper("Options"),
		"/{basePath}/slack/events",
		SlackEventsOptions,
	},

	Route{
		"SlackMessagesOptions",
		strings.ToUpper("Options"),
		"/{basePath}/slackMessages",
		SlackMessagesOptions,
	},

	Route{
		"SlackNotificationsOptions",
		strings.ToUpper("Options"),
		"/{basePath}/slackNotifications",
		SlackNotificationsOptions,
	},

	Route{
		"StandupOptions",
		strings.ToUpper("Options"),
		"/{basePath}/standup",
		StandupOptions,
	},

	Route{
		"SuccessOptions",
		strings.ToUpper("Options"),
		"/{basePath}/success",
		SuccessOptions,
	},

	Route{
		"UpcomingLeavesOptions",
		strings.ToUpper("Options"),
		"/{basePath}/upcomingLeaves",
		UpcomingLeavesOptions,
	},

	Route{
		"UpdateStatusOptions",
		strings.ToUpper("Options"),
		"/{basePath}/updateStatus",
		UpdateStatusOptions,
	},

	Route{
		"UpdateWorkspaceOptions",
		strings.ToUpper("Options"),
		"/{basePath}/updateWorkspace",
		UpdateWorkspaceOptions,
	},

	Route{
		"WebNotificationOptions",
		strings.ToUpper("Options"),
		"/{basePath}/webNotification",
		WebNotificationOptions,
	},

	Route{
		"WebPollOptions",
		strings.ToUpper("Options"),
		"/{basePath}/webPoll",
		WebPollOptions,
	},
}
