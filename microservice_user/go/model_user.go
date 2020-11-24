/*
 * Sample API
 *
 * Test de OpenApi
 *
 * API version: 0.1.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"log"
	"database/sql"
)

type Company struct {

	Id int32 `json:"id"`

	Name string `json:"name"`

	Age int32 `json:"age"`

	Address string `json:"address"`

	Salary float64 `json:"salary"`
}

func GetCompanys(db *sql.DB) ([]Company, error){
	log.Print("GetCompanys")

	rows, err := db.Query("SELECT * FROM company")

	if err != nil{
		return nil, err
	}

	defer rows.Close()

	companys := []Company{}

	for rows.Next(){
		var c Company
	
		if err := rows.Scan(&c.Id, &c.Name, &c.Age, &c.Address, &c.Salary); err != nil {
			return nil, err
		}

		companys = append(companys, c)
	}

	return companys, nil
}

func GetCompany(db *sql.DB, id int) (Company, error){

	row := db.QueryRow("SELECT * FROM company WHERE id=$1", id)

	var c Company 

	if err := row.Scan(&c.Id, &c.Name, &c.Age, &c.Address, &c.Salary); err != nil {
		return c, err
	}

	return c, nil
}

func InsertCompany(db *sql.DB) error{
	_, err := db.Exec("INSERT INTO company(name, age, address, salary) VALUES ('NewPeople', 25, 'France', 35000)")

	return err
}