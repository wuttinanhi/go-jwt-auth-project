package auth

// mock database
var user = map[string]string{
	"admin": "password",
	"user":  "12345",
}

var authService *AuthService = nil

type AuthService struct{}

func (s *AuthService) Login(username, password string) bool {
	return user[username] == password
}

func GetAuthService() *AuthService {
	if authService == nil {
		authService = &AuthService{}
	}
	return authService
}
