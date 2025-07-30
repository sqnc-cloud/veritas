package usecases_test

import (
	"context"
	"errors"
	"testing"
	"veritas/core/domain"
	"veritas/core/usecases"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserOutputPort struct {
	mock.Mock
}

func (m *MockUserOutputPort) CreateUser(ctx context.Context, user *domain.User) (primitive.ObjectID, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(primitive.ObjectID), args.Error(1)
}

func (m *MockUserOutputPort) GetUser(ctx context.Context, id primitive.ObjectID) (*domain.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserOutputPort) UpdateUser(ctx context.Context, id primitive.ObjectID, user *domain.User) error {
	args := m.Called(ctx, id, user)
	return args.Error(0)
}

func (m *MockUserOutputPort) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserOutputPort) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserOutputPort) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.User), args.Error(1)
}

type UserUseCaseTestSuite struct {
	suite.Suite
	mockOutputPort *MockUserOutputPort
	userUseCase    *usecases.UserUsecase
	ctx            context.Context
}

func (s *UserUseCaseTestSuite) SetupTest() {
	s.mockOutputPort = new(MockUserOutputPort)
	s.userUseCase = usecases.NewUserUsecase(s.mockOutputPort)
	s.ctx = context.Background()
}

func (s *UserUseCaseTestSuite) TestCreateUser() {
	createUserInput := usecases.CreateUserInput{
		Name:     "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	expectedID := primitive.NewObjectID()

	// Test case 1: Successful user creation
	s.mockOutputPort.On("CreateUser", s.ctx, mock.AnythingOfType("*domain.User")).Return(expectedID, nil).Once()
	id, err := s.userUseCase.CreateUser(s.ctx, createUserInput)
	s.NoError(err)
	s.Equal(expectedID, id)
	s.mockOutputPort.AssertExpectations(s.T())

	// Test case 2: Error during user creation
	s.SetupTest() // Reset mock for new test case
	expectedError := errors.New("failed to create user")
	s.mockOutputPort.On("CreateUser", s.ctx, mock.AnythingOfType("*domain.User")).Return(primitive.NilObjectID, expectedError).Once()
	id, err = s.userUseCase.CreateUser(s.ctx, createUserInput)
	s.Error(err)
	s.Equal(primitive.NilObjectID, id)
	s.Equal(expectedError, err)
	s.mockOutputPort.AssertExpectations(s.T())
}

func (s *UserUseCaseTestSuite) TestReadUser() {
	expectedUser := &domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser",
		Email:    "test@example.com",
	}
	existingID := expectedUser.ID.Hex()

	// Test case 1: Successful user retrieval
	s.mockOutputPort.On("GetUser", s.ctx, expectedUser.ID).Return(expectedUser, nil).Once()
	user, err := s.userUseCase.ReadUser(s.ctx, existingID)
	s.NoError(err)
	s.Equal(expectedUser, user)
	s.mockOutputPort.AssertExpectations(s.T())

	// Test case 2: User not found
	s.SetupTest() // Reset mock for new test case
	expectedError := errors.New("user not found")
	s.mockOutputPort.On("GetUser", s.ctx, mock.AnythingOfType("primitive.ObjectID")).Return(nil, expectedError).Once()
	user, err = s.userUseCase.ReadUser(s.ctx, primitive.NilObjectID.Hex())
	s.Error(err)
	s.Nil(user)
	s.Equal(expectedError, err)
	s.mockOutputPort.AssertExpectations(s.T())

	// Test case 3: Invalid ID
	s.SetupTest() // Reset mock for new test case
	user, err = s.userUseCase.ReadUser(s.ctx, "invalid-id")
	s.Error(err)
	s.Nil(user)
	s.Contains(err.Error(), "invalid id")
	s.mockOutputPort.AssertExpectations(s.T())
}

func (s *UserUseCaseTestSuite) TestUpdateUser() {
	existingID := primitive.NewObjectID()
	updateUserInput := usecases.UpdateUserInput{
		Name:  "updateduser",
		Email: "updated@example.com",
	}

	existingUser := &domain.User{
		ID:       existingID,
		Username: "olduser",
		Email:    "old@example.com",
	}

	// Test case 1: Successful user update
	s.mockOutputPort.On("GetUser", s.ctx, existingID).Return(existingUser, nil).Once()
	s.mockOutputPort.On("UpdateUser", s.ctx, existingID, mock.AnythingOfType("*domain.User")).Return(nil).Once()
	user, err := s.userUseCase.UpdateUser(s.ctx, existingID.Hex(), updateUserInput)
	s.NoError(err)
	s.NotNil(user)
	s.Equal(updateUserInput.Name, user.Username)
	s.Equal(updateUserInput.Email, user.Email)
	s.mockOutputPort.AssertExpectations(s.T())

	// Test case 2: Error getting user
	s.SetupTest() // Reset mock for new test case
	expectedError := errors.New("failed to get user")
	s.mockOutputPort.On("GetUser", s.ctx, existingID).Return(nil, expectedError).Once()
	user, err = s.userUseCase.UpdateUser(s.ctx, existingID.Hex(), updateUserInput)
	s.Error(err)
	s.Nil(user)
	s.Contains(err.Error(), "failed to get user")
	s.mockOutputPort.AssertExpectations(s.T())

	// Test case 3: Error during user update
	s.SetupTest() // Reset mock for new test case
	s.mockOutputPort.On("GetUser", s.ctx, existingID).Return(existingUser, nil).Once()
	expectedError = errors.New("failed to update user")
	s.mockOutputPort.On("UpdateUser", s.ctx, existingID, mock.AnythingOfType("*domain.User")).Return(expectedError).Once()
	user, err = s.userUseCase.UpdateUser(s.ctx, existingID.Hex(), updateUserInput)
	s.Error(err)
	s.Nil(user)
	s.Equal(expectedError, err)
	s.mockOutputPort.AssertExpectations(s.T())

	// Test case 4: Invalid ID
	s.SetupTest() // Reset mock for new test case
	user, err = s.userUseCase.UpdateUser(s.ctx, "invalid-id", updateUserInput)
	s.Error(err)
	s.Nil(user)
	s.Contains(err.Error(), "invalid id")
	s.mockOutputPort.AssertExpectations(s.T())
}

func (s *UserUseCaseTestSuite) TestDeleteUser() {
	existingID := primitive.NewObjectID()

	// Test case 1: Successful user deletion
	s.mockOutputPort.On("DeleteUser", s.ctx, existingID).Return(nil).Once()
	err := s.userUseCase.DeleteUser(s.ctx, existingID.Hex())
	s.NoError(err)
	s.mockOutputPort.AssertExpectations(s.T())

	// Test case 2: Error during user deletion
	s.SetupTest() // Reset mock for new test case
	expectedError := errors.New("failed to delete user")
	s.mockOutputPort.On("DeleteUser", s.ctx, existingID).Return(expectedError).Once()
	err = s.userUseCase.DeleteUser(s.ctx, existingID.Hex())
	s.Error(err)
	s.Equal(expectedError, err)
	s.mockOutputPort.AssertExpectations(s.T())

	// Test case 3: Invalid ID
	s.SetupTest() // Reset mock for new test case
	err = s.userUseCase.DeleteUser(s.ctx, "invalid-id")
	s.Error(err)
	s.Contains(err.Error(), "invalid id")
	s.mockOutputPort.AssertExpectations(s.T())
}

func (s *UserUseCaseTestSuite) TestGetAllUsers() {
	expectedUsers := []*domain.User{
		{
			ID:       primitive.NewObjectID(),
			Username: "user1",
			Email:    "user1@example.com",
		},
		{
			ID:       primitive.NewObjectID(),
			Username: "user2",
			Email:    "user2@example.com",
		},
	}

	// Test case 1: Successful retrieval of all users
	s.mockOutputPort.On("GetAllUsers", s.ctx).Return(expectedUsers, nil).Once()
	users, err := s.userUseCase.GetAllUsers(s.ctx)
	s.NoError(err)
	s.Equal(expectedUsers, users)
	s.mockOutputPort.AssertExpectations(s.T())

	// Test case 2: Error during retrieval of all users
	s.SetupTest() // Reset mock for new test case
	expectedError := errors.New("failed to get all users")
	s.mockOutputPort.On("GetAllUsers", s.ctx).Return(nil, expectedError).Once()
	users, err = s.userUseCase.GetAllUsers(s.ctx)
	s.Error(err)
	s.Nil(users)
	s.Equal(expectedError, err)
	s.mockOutputPort.AssertExpectations(s.T())
}

// In order for 'go test' to run this suite, we need to expose it using the 'suite.Run' function
func TestUserUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}
