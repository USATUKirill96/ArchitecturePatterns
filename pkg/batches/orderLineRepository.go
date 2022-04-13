package batches

import "database/sql"

type OrderLineRepository struct {
	db *sql.DB
}

func NewOrderLineRepository(db *sql.DB) *OrderLineRepository { return &OrderLineRepository{db: db} }

func (r *OrderLineRepository) Get(id int) (*OrderLine, error) {
	stmt := `
	SELECT id, order_id, sku, quantity, batch_id
	  FROM order_line
	 WHERE id = $1
	`
	line := &OrderLine{}
	err := r.db.QueryRow(stmt, id).Scan(&line.ID, &line.OrderID, &line.SKU, &line.Quantity, &line.BatchID)
	if err != nil {
		return nil, err
	}
	return line, nil
}

func (r *OrderLineRepository) Insert(line *OrderLine) (int, error) {
	stmt := `
	   INSERT INTO order_line (order_id, sku, quantity, batch_id) 
	   VALUES ($1, $2, $3, $4) 
	RETURNING id
	`

	var id int

	err := r.db.QueryRow(stmt, line.OrderID, line.SKU, line.Quantity, line.BatchID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *OrderLineRepository) Update(line *OrderLine) error {
	stmt := `
	UPDATE order_line
	   SET order_id = $2, sku = $3, quantity = $4, batch_id = $5
	 WHERE id = $1 

	`

	err := r.db.QueryRow(stmt, line.ID, line.OrderID, line.SKU, line.Quantity, line.BatchID)
	if err != nil {
		return err.Err()
	}
	return nil
}
