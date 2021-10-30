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

func (s *SpaceRepo) Create(memberId int, spaceTitle string, isOwner bool) error {
	conn := s.conn
	cb := s.contextBackground

	tx, err := conn.Begin(cb)
	if err != nil {
		return err
	}
	defer tx.Rollback(cb)
	defer conn.Close(cb)
	// this transaction will insert into space and space_memeber

	_, err = tx.Exec(cb, "INSERT INTO space (member_id, space_title) VALUES ($1, $2);", memberId, spaceTitle)
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

// find all spaces, include the ones you own and the ones you are shared with
func (s *SpaceRepo) FindAll(memberId int) ([]int, error) {
	conn := s.conn
	cb := s.contextBackground

	// TODO: add index on member_id on space_member table
	// TODO: use JOIN to select repayment_plan data to show to users
	rows, err := conn.Query(cb, "SELECT space_id FROM space_member WHERE member_id = $1", memberId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	// defer conn.Close(cb)
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

// find the spaces you owned
func (s *SpaceRepo) FindOwnedSpaces(memberId int) ([]int, error) {
	conn := s.conn
	cb := s.contextBackground

	rows, err := conn.Query(cb, "SELECT id FROM space WHERE member_id = $1", memberId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// defer conn.Close(cb)
	var spaces []int

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

type SpaceMemberInputModel struct {
	SpaceId  int
	MemberId int
	IsOwner  bool
}

func (s *SpaceRepo) AddToSpaceMember(spaceMemberInput SpaceMemberInputModel) error {
	conn := s.conn
	cb := s.contextBackground

	_, err := conn.Exec(
		cb,
		"INSERT INTO space_member (space_id, member_id, is_owner) VALUES ($1, $2, $3)",
		spaceMemberInput.SpaceId,
		spaceMemberInput.MemberId,
		spaceMemberInput.IsOwner,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// TODO: add find space detail by Id
