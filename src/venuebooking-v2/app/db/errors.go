package db

type (
	// UserNotFoundError struct
	UserNotFoundError struct {
	}
)

// Error function to check if User Exists
func (m *UserNotFoundError) Error() string {
	return "Email does not exists!"
}
