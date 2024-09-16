package services

import "golang.org/x/crypto/bcrypt"

type authService struct {
}

func (s *authService) SignIn() {

}

func (s *authService) SignUp(l string, p string) {

}

func (s *authService) ValidateLogin(login string) bool {
	return true
}

/* Password Utils */
func (s *authService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *authService) CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
