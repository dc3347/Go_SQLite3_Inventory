package models

type ProductModel struct{

}

func(*ProductModel) FindAll() ([]entities.Product, error){
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} 
	else {
		rows, err2 := db.Query("select * from product")
		if err2 != nil {
			return nil, err
		} 
		else{
			var products []entities.Product
			for rows.Next() {
				var product entities.Product
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Supplier)
				products = append(products, product)
			}
			return products, nil
		}
	}
}

func(*ProductModel) Search(min, max float64) ([]entities.Product, error){
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} 
	else {
		rows, err2 := db.Query("select * from product where price >= ? and price <= ?", min, max)
		if err2 != nil {
			return nil, err
		} 
		else{
			var products []entities.Product
			for rows.Next() {
				var product entities.Product
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Supplier)
				products = append(products, product)
			}
			return products, nil
		}
	}
}
func(*ProductModel) Create(product *entities.Product) bool{
	db, err := config.GetDB()
	result, err := db.Exec("insert into product(name, price, quantity, supplier) values(?,?,?,?)", product.Name, product.Price, product.Quantity, product.Supplier)

	if err != nil{
		return false
	}
	lastId, _ := result.LastInsertId()
	product.Id = lastId
	rowsAffected, err2 := result.RowsAffected()
	if err2 != nil{
		return false
	}
	return rowsAffected > 0
}

func(*ProductModel) Delete(id int64) bool{
	db, err := config.GetDB()
	result, err := db.Exec("delete from product where id = ?", id)
	if err != nil{
		return false
	}
	
	rowsAffected, err2 := result.RowsAffected()
	if err2 != nil{
		return false
	}
	return rowsAffected > 0
}

func(*ProductModel) Update(product entities.Product) bool{
	db, err := config.GetDB()
	result, err := db.Exec("update product set name = ?, price = ?, quantity = ?, supplier = ? where id = ?", product.Name, product.Price, product.Quantity, product.Supplier, product.Id)

	if err != nil{
		return false
	}
	rowsAffected, err2 := result.RowsAffected()
	if err2 != nil{
		return false
	}
	return rowsAffected > 0
}

func(*ProductModel) Find(id int) (entities.Product, error){
	db, err := config.GetDB()
	if err != nil {
		return entities.Product{}, err
	} 
	else {
		rows, err2 := db.Query("select * from product where id = ?", id)
		if err2 != nil {
			return entities.Product{}, err
		} 
		else{
			var product entities.Product
			for rows.Next() {
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Supplier)
			}
			return product, nil
		}
	}
}