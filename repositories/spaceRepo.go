package repositories

import (
	"budget-plan-app/backend/db"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type SpaceRepo struct {
	conn              *pgx.Conn
	contextBackground context.Context
}

func NewSpaceRepo() *SpaceRepo {
	dbConn := db.ConnectDB()
	return &SpaceRepo{
		conn:              dbConn,
		contextBackground: context.Background(),
	}
}

func (s *SpaceRepo) Create(memberId int, isOwner bool) error {
	conn := s.conn
	cb := s.contextBackground

	tx, err := conn.Begin(cb)
	if err != nil {
		return err
	}
	defer tx.Rollback(cb)
	defer conn.Close(cb)
	// this transaction will insert into space and space_memeber

	_, err = tx.Exec(cb, "INSERT INTO space (member_id) VALUES ($1);", memberId)
	if err != nil {
		fmt.Println("first tx err: ", err)
		return err
	}
	transactionSql := "INSERT INTO space_member (space_id, member_id, is_owner) SELECT space.id, $1, $2 FROM space WHERE space.member_id=$3"
	_, err = tx.Exec(cb, transactionSql, memberId, isOwner, memberId)
	if err != nil {
		fmt.Println("second tx err: ", err.Error())
		return err
	}

	err = tx.Commit(cb)
	if err != nil {
		return err
	}
	return nil
}

func (s *SpaceRepo) FindAll(memberId int) ([]int, error) {
	conn := s.conn
	cb := s.contextBackground

	// TODO: use JOIN to select repayment_plan data to show to users
	rows, err := conn.Query(cb, "SELECT space_id FROM space_member WHERE member_id = $1", memberId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var spaces []int
	// equavilant to while rows.Next() == true {...}
	for rows.Next() {
		var spaceId int
		err := rows.Scan(&spaceId)
		if err != nil {
			return nil, err
		}
		spaces = append(spaces, spaceId)
	}
	return spaces, nil
}
