package repositories

import (
	"budget-plan-app/backend/db"
	"budget-plan-app/backend/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// create a struct (like a class)

type MemberRepo struct {
	conn              *pgx.Conn
	contextBackground context.Context
}

func NewMemberRepo() *MemberRepo {
	dbConn := db.ConnectDB()
	return &MemberRepo{conn: dbConn, contextBackground: context.Background()}
}

func (m *MemberRepo) Create(member models.Member) error {
	conn := m.conn
	cb := m.contextBackground
	defer conn.Close(cb)
	res, err := conn.Exec(cb, "INSERT INTO member (email) VALUES ($1)", member.Email)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(res.RowsAffected())
	return nil
}
