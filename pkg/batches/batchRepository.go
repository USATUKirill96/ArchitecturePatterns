package batches

import (
	"database/sql"
	"errors"
)

type BatchRepository struct {
	db *sql.DB
}

func NewBatchRepository(db *sql.DB) *BatchRepository { return &BatchRepository{db: db} }

func (r *BatchRepository) Get(id int) (*Batch, error) {
	stmt := `
	SELECT id, reference, sku, eta, purchased_quantity
	  FROM batch
	 WHERE id = $1
	`

	batch := &Batch{}
	err := r.db.QueryRow(stmt, id).Scan(&batch.ID, &batch.Reference, &batch.SKU, &batch.ETA, &batch.purchasedQuantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("database: no records found")
		} else {
			return nil, err
		}
	}

	err = r.LoadAllocations(batch)
	if err != nil {
		return nil, err
	}

	return batch, nil
}

func (r *BatchRepository) FliterBySKU(sku string) ([]*Batch, error) {

	stmt := `
	SELECT id, reference, sku, eta, purchased_quantity
	  FROM batch
	 WHERE sku = $1
	`

	rows, err := r.db.Query(stmt, sku)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	batches := []*Batch{}

	for rows.Next() {
		b := &Batch{}
		err = rows.Scan(&b.ID, &b.Reference, &b.SKU, &b.ETA, &b.purchasedQuantity)
		if err != nil {
			return nil, err
		}
		batches = append(batches, b)
	}

	// TODO: filter all allocations by batch ids and attach manually instead of creating request for each batch
	for _, batch := range batches {
		err := r.LoadAllocations(batch)
		if err != nil {
			return nil, err
		}
	}

	return batches, nil

}

func (r *BatchRepository) LoadAllocations(batch *Batch) error {

	stmt := `
	SELECT id, order_id, sku, quantity, batch_id
	  FROM order_line
	 WHERE batch_id = $1
	`
	rows, err := r.db.Query(stmt, batch.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	orderLines := Allocations{}

	for rows.Next() {
		ol := &OrderLine{}
		err = rows.Scan(&ol.ID, &ol.OrderID, &ol.SKU, &ol.Quantity, &ol.BatchID)
		if err != nil {
			return err
		}
		orderLines = append(orderLines, ol)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	batch.allocations = orderLines

	return nil
}

func (r *BatchRepository) Insert(batch *Batch) (int, error) {
	stmt := `
	   INSERT INTO batch (reference, sku, eta, purchased_quantity) 
	   VALUES ($1, $2, $3, $4) 
	RETURNING id
	`
	var id int
	err := r.db.QueryRow(stmt, batch.Reference, batch.SKU, batch.ETA, batch.purchasedQuantity).Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *BatchRepository) Update(batch *Batch) error {
	stmt := `
	UPDATE batch 
	   SET reference = $2, sku = $3, eta = $4, purchased_quantity = $5 
	 WHERE id = $1
	`
	err := r.db.QueryRow(stmt, batch.ID, batch.Reference, batch.SKU, batch.ETA, batch.purchasedQuantity)

	if err != nil {
		return err.Err()
	}
	return nil
}
