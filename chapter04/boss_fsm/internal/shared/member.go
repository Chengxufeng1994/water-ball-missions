package shared

type Member struct {
	UserID  string
	IsAdmin bool
	Role    Role
}

func NewAdminMember(userID string) *Member {
	return &Member{
		UserID:  userID,
		IsAdmin: true,
		Role:    Admin,
	}
}

func NewMember(userID string) *Member {
	return &Member{
		UserID:  userID,
		IsAdmin: false,
		Role:    User,
	}
}
