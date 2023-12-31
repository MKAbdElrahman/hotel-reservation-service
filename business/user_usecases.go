package business

import (
	"context"

	"github.com/mkabdelrahman/hotel-reservation/types"
)

func (m *Manager) ListUsers(ctx context.Context, filter types.UsersPaginationFilter) ([]*types.User, error) {

	users, err := m.UserStore.GetUsersWithPagination(ctx, filter)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (m *Manager) UpdateUser(ctx context.Context, ID string, updateFields types.UpdateUserParams) (*types.User, error) {

	users, err := m.UserStore.UpdateUser(ctx, ID, updateFields)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (m *Manager) DeleteUser(ctx context.Context, ID string) error {

	err := m.UserStore.DeleteUser(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) GetUserByID(ctx context.Context, ID string) (*types.User, error) {

	user, err := m.UserStore.GetUserByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// func (m *Manager) SearchUsers(ctx context.Context, searchQuery string) ([]*types.User, error) {

// 	return nil, nil
// }
