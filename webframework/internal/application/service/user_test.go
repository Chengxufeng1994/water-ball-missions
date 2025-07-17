package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserApplicationService_Register(t *testing.T) {
	svc := NewUserApplicationService()

	// Test case 1: Successful registration
	user, err := svc.Register("test@example.com", "testuser", "password123")
	assert.NoError(t, err)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "testuser", user.Username)

	// Test case 2: Registration with duplicate email
	_, err = svc.Register("test@example.com", "anotheruser", "anotherpass")
	assert.Error(t, err)
	assert.EqualError(t, err, "Duplicate email")

	// Test case 3: Registration with invalid email format
	_, err = svc.Register("invalid-email", "testuser", "password123")
	assert.Error(t, err)
	assert.EqualError(t, err, "Registration's format incorrect.")

	// Test case 4: Registration with email too short
	_, err = svc.Register("a@b.c", "testuser", "password123")
	assert.Error(t, err)
	assert.EqualError(t, err, "Registration's format incorrect.")

	// Test case 5: Registration with email too long
	_, err = svc.Register("longemail123456789012345678901234567890@example.com", "testuser", "password123")
	assert.Error(t, err)
	assert.EqualError(t, err, "Registration's format incorrect.")

	// Test case 6: Registration with username too short
	_, err = svc.Register("user2@example.com", "user", "password123")
	assert.Error(t, err)
	assert.EqualError(t, err, "Registration's format incorrect.")

	// Test case 7: Registration with username too long
	_, err = svc.Register("user3@example.com", "longusername123456789012345678901234567890", "password123")
	assert.Error(t, err)
	assert.EqualError(t, err, "Registration's format incorrect.")

	// Test case 8: Registration with password too short
	_, err = svc.Register("user4@example.com", "testuser4", "pass")
	assert.Error(t, err)
	assert.EqualError(t, err, "Registration's format incorrect.")

	// Test case 9: Registration with password too long
	_, err = svc.Register("user5@example.com", "testuser5", "longpassword123456789012345678901234567890")
	assert.Error(t, err)
	assert.EqualError(t, err, "Registration's format incorrect.")
}

func TestUserApplicationService_Login(t *testing.T) {
	svc := NewUserApplicationService()
	_, err := svc.Register("login@example.com", "loginuser", "loginpass")
	assert.NoError(t, err)

	// Test case 1: Successful login
	user, token, err := svc.Login("login@example.com", "loginpass")
	assert.NoError(t, err)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, token.Token)

	// Test case 2: Login with invalid password
	_, _, err = svc.Login("login@example.com", "wrongpass")
	assert.Error(t, err)
	assert.EqualError(t, err, "Credentials Invalid")

	// Test case 3: Login with non-existent email
	_, _, err = svc.Login("nonexistent@example.com", "anypass")
	assert.Error(t, err)
	assert.EqualError(t, err, "Credentials Invalid")

	// Test case 4: Login with invalid email format
	_, _, err = svc.Login("invalid-email", "loginpass")
	assert.Error(t, err)
	assert.EqualError(t, err, "Login's format incorrect.")

	// Test case 5: Login with password too short (for a registered user)
	_, _, err = svc.Login("login@example.com", "shor")
	assert.Error(t, err)
	assert.EqualError(t, err, "Login's format incorrect.")
}

func TestUserApplicationService_Rename(t *testing.T) {
	svc := NewUserApplicationService()
	registeredUser, _ := svc.Register("rename@example.com", "oldname", "renamepass")

	// Test case 1: Successful rename
	newName, err := svc.Rename(registeredUser.ID, "newname")
	assert.NoError(t, err)
	assert.Equal(t, "newname", newName)
	updatedUser, _, _ := svc.Login("rename@example.com", "renamepass")
	assert.Equal(t, "newname", updatedUser.Username)

	// Test case 2: Rename with non-existent user ID
	_, err = svc.Rename("nonexistent-id", "someothername")
	assert.Error(t, err)
	assert.EqualError(t, err, "user with id nonexistent-id not found")

	// Test case 3: Rename with name too short
	_, err = svc.Rename(registeredUser.ID, "name")
	assert.Error(t, err)
	assert.EqualError(t, err, "Name's format invalid.")

	// Test case 4: Rename with name too long
	_, err = svc.Rename(registeredUser.ID, "longname123456789012345678901234567890")
	assert.Error(t, err)
	assert.EqualError(t, err, "Name's format invalid.")
}

func TestUserApplicationService_List(t *testing.T) {
	svc := NewUserApplicationService()
	_, err1 := svc.Register("user1@example.com", "AliceUser", "pass1")
	assert.NoError(t, err1, "Registration of user1 should not fail")
	_, err2 := svc.Register("user2@example.com", "BobUser", "pass2")
	assert.NoError(t, err2, "Registration of user2 should not fail")
	_, err3 := svc.Register("user3@example.com", "CharlieUser", "pass3")
	assert.NoError(t, err3, "Registration of user3 should not fail")

	// Test case 1: List all users (empty keyword)
	users, err := svc.List("")
	assert.NoError(t, err)
	assert.Len(t, users, 3)

	// Test case 2: List users with a keyword match
	users, err = svc.List("AliceUser")
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "AliceUser", users[0].Username)

	// Test case 3: List users with no keyword match
	users, err = svc.List("xyz")
	assert.NoError(t, err)
	assert.Len(t, users, 0)
}

func TestUserApplicationService_ValidateToken(t *testing.T) {
	svc := NewUserApplicationService()
	_, err := svc.Register("token@example.com", "tokenuser", "tokenpass")
	assert.NoError(t, err)
	registeredUser, loginToken, _ := svc.Login("token@example.com", "tokenpass")

	// Test case 1: Successful token validation
	userID, err := svc.ValidateToken(loginToken.Token)
	assert.NoError(t, err)
	assert.Equal(t, registeredUser.ID, userID)

	// Test case 2: Invalid token
	_, err = svc.ValidateToken("invalid-token")
	assert.Error(t, err)
	assert.EqualError(t, err, "Can't authenticate who you are.")
}
