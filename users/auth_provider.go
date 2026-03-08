package users

type AuthProvider string

const (
	AuthProviderGoogle AuthProvider = "GOOGLE"
	AuthProviderGitHub AuthProvider = "GITHUB"
)
