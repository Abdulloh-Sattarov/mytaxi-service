package postgres

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	pb "github.com/abdullohsattorov/mytaxi-service/genproto"
)

type orderRepo struct {
	db *sqlx.DB
}

// NewTaxiRepo ...
func NewOrderRepo(db *sqlx.DB) *orderRepo {
	return &orderRepo{db: db}
}

func (t *orderRepo) CreateOrder(order pb.OrderReq) (pb.OrderRes, error) {
	var id string
	err := t.db.QueryRow(`
        INSERT INTO orders(id, cost, status, driver_id, client_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7) returning id`, order.Id, order.Cost, order.Status, order.DriverId, order.ClientId, time.Now().UTC(), time.Now().UTC()).Scan(&id)
	if err != nil {
		return pb.OrderRes{}, err
	}
	newOrder, err := t.GetOrder(id)
	if err != nil {
		return pb.OrderRes{}, err
	}

	return newOrder, nil
}

func (t *orderRepo) GetOrder(id string) (pb.OrderRes, error) {
	var order pb.OrderRes
	var driverId, clientId string

	err := t.db.QueryRow(`
        SELECT id, cost, status, driver_id, client_id, created_at, updated_at status FROM orders
        WHERE id=$1 and deleted_at is null`, id).Scan(
		&order.Id,
		&order.Cost,
		&order.Status,
		&driverId,
		&clientId,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return pb.OrderRes{}, err
	}

	driver, err := NewDriverRepo(t.db).GetDriver(driverId)
	if err != nil {
		return pb.OrderRes{}, err
	}

	client, err := NewClientRepo(t.db).GetClient(clientId)
	if err != nil {
		return pb.OrderRes{}, err
	}

	order.Driver = &driver
	order.Client = &client

	return order, nil
}

func (t *orderRepo) ListOrders(clientId string, page, limit int64) ([]*pb.OrderRes, int64, error) {
	offset := (page - 1) * limit

	rows, err := t.db.Queryx(
		`SELECT id, status, cost, driver_id, client_id, created_at, updated_at status FROM orders WHERE client_id = $1 LIMIT $2 OFFSET $3`,
		clientId, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var (
		orders []*pb.OrderRes
		count  int64
	)
	for rows.Next() {
		var order pb.OrderRes
		var driverId, clientId string
		err = rows.Scan(&order.Id, &order.Status, &order.Cost, &driverId, &clientId, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}

		client, _ := NewClientRepo(t.db).GetClient(clientId)
		driver, _ := NewDriverRepo(t.db).GetDriver(driverId)

		order.Client = &client
		order.Driver = &driver
		orders = append(orders, &order)
	}

	err = t.db.QueryRow(
		`SELECT count(*) FROM orders WHERE client_id = $1 and deleted_at is null`,
		clientId).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	return orders, count, nil
}

func (t *orderRepo) UpdateOrder(order pb.OrderReq) (pb.OrderRes, error) {
	var status string

	err := t.db.QueryRow(`SELECT status from orders WHERE id=$1`, order.Id).Scan(&status)
	if err != nil {
		return pb.OrderRes{}, err
	}

	if status == "accepted" && order.Status == "cancelled" || status == "finished" && order.Status == "cancelled" {
		return pb.OrderRes{}, sql.ErrNoRows
	}

	result, err := t.db.Exec(`UPDATE orders SET status=$2, updated_at=$3 WHERE id=$1`,
		order.Id, order.Status, time.Now().UTC())
	if err != nil {
		return pb.OrderRes{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.OrderRes{}, sql.ErrNoRows
	}

	res, err := t.GetOrder(order.Id)
	if err != nil {
		return pb.OrderRes{}, err
	}

	order.Status = res.Status
	order.DriverId = res.Client.Id
	order.ClientId = res.Client.Id

	return pb.OrderRes{}, nil
}

func (t *orderRepo) DeleteOrder(id string) error {
	result, err := t.db.Exec(`UPDATE orders SET deleted_at = $2 WHERE id=$1`, id, time.Now().UTC())
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
