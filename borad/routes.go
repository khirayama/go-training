package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"MessageIndex",
		"GET",
		"/messages",
		MessageIndex,
	},
	Route{
		"MessageCreate",
		"POST",
		"/messages",
		MessageCreate,
	},
	Route{
		"MessageDelete",
		"DELETE",
		"/messages/{messageId}",
		MessageDelete,
	},
}
