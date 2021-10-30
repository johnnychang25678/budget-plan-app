package repositories

import (
	"budget-plan-app/backend/db"
	"context"
	"errors"
	"fmt"
)

type MemberRepo struct {
	contextBackground context.Context
}

func NewMemberRepo() *MemberRepo {
	return &MemberRepo{contextBackground: context.Background()}
}

func (m *MemberRepo) Create(email string) error {
	conn := db.ConnectDB()
	cb := m.contextBackground
	defer conn.Close(cb)
	_, err := conn.Exec(cb, "INSERT INTO member (email) VALUES ($1)", email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully create memeber!")
	return nil
}

func (m *MemberRepo) FindByEmail(email string) (int, error) {
	conn := db.ConnectDB()
	cb := m.contextBackground
	defer conn.Close(cb)
	var userIds []int
	rows, err := conn.Query(cb, "SELECT id FROM member where email = $1", email)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	for rows.Next() {
		var userId int
		rows.Scan(&userId)
		userIds = append(userIds, userId)
	}
	fmt.Println("findByEmail result:", userIds)
	if len(userIds) > 1 {
		return 0, errors.New("more than one same email in memeber DB")
	}
	if len(userIds) == 0 {
		return 0, nil
	}
	return userIds[0], nil
}
