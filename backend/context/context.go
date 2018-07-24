package context

type ContextsKey string

const (
	//Key to get user id populated from token in AuhtWithToken
	UserID     = ContextsKey("userId")
	Credentils  = ContextsKey("creds")
)
