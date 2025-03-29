package middleware

// ContextValueKey represents the context value keys used by middleware.
type ContextValueKey uint

const (
	ContextValueKeyJWT  ContextValueKey = iota // context value key for the JWT.
	ContextValueKeyUser                        // context value key for the user's name.
)
