package config

import "github.com/gorilla/sessions"

var Store = sessions.NewCookieStore([]byte("clé-secrète-pour-sessions"))
