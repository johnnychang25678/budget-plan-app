package repositories

import (
	"budget-plan-app/backend/db"
	"budget-plan-app/backend/models"
	"context"
	"fmt"
)

type RepaymentPlanRepo struct {
	// conn              *pgx.Conn
	contextBackground context.Context
}

func NewRepaymentRepo() *RepaymentPlanRepo {
	return &RepaymentPlanRepo{

		contextBackground: context.Background(),
	}
}

func (r *RepaymentPlanRepo) Create(repamymentPlan models.RepaymentPlan) error {
	conn := db.ConnectDB()
	cb := r.contextBackground

	defer conn.Close(cb)
	_, err := conn.Exec(
		cb,
		"INSERT INTO repayment_plan (space_id, title, total_cost, due_date) VALUES ($1, $2, $3, $4)",
		repamymentPlan.SpaceId,
		repamymentPlan.Title,
		repamymentPlan.TotalCost,
		repamymentPlan.DueDate,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// return repaymentPlanId
func (r *RepaymentPlanRepo) GetPlansById(planId int) ([]models.RepaymentPlan, error) {
	conn := db.ConnectDB()
	cb := r.contextBackground

	rows, err := conn.Query(cb, "SELECT space_id, title, total_cost, due_date FROM repayment_plan WHERE id = $1", planId)
	if err != nil {
		return []models.RepaymentPlan{}, err
	}
	// fmt.Println("repaymentPlans from repo: ", rows)
	defer rows.Close()
	// defer conn.Close(cb)
	var repaymentPlans []models.RepaymentPlan

	for rows.Next() {
		var repaymentPlan models.RepaymentPlan

		err := rows.Scan(
			&repaymentPlan.SpaceId,
			&repaymentPlan.Title,
			&repaymentPlan.TotalCost,
			&repaymentPlan.DueDate,
		)
		if err != nil {
			fmt.Println("scan error!")
			return []models.RepaymentPlan{}, err
		}
		repaymentPlans = append(repaymentPlans, repaymentPlan)
	}
	// fmt.Println("repaymentPlans from repo: ", repaymentPlans)
	return repaymentPlans, nil
}
