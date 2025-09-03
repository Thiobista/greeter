package repository

import (
	"context"
	"fmt"
	"greeter/cmd/greeter/root/model"
	"greeter/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
}
 

func (r *EmployeeRepo) InsertEmployee(emp *model.Employee) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), emp)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (r *EmployeeRepo) FindEmployeeByID(employeeID string) (*model.Employee, error){
	var emp model.Employee
	err := r.MongoCollection.FindOne(context.Background(),
	bson.D{{key"employee_id", value:empID}}).Decode(&emp)

	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *EmployeeRepo) FindAllEmployee() ([]model.Employee, error) { 
	results, err := r.MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	var emps []model.Employee
	err = results.All(context.Background(), &emps)
	if err != nil {
		return nil, fmt.Errorf("results decoding error: %s", err).Error()
	}

	return emps, nil
}

func (r *EmployeeRepo) UpdateEmployeeByID(empID string, updateEmp *model.employee) (int16, error) {
	result,err := r.MongoCollection.UpdateOne(context.Background(),
		bson.D{{key:"employee_id", value:empID}},
		bson.D{{"$set", value:updateEmp}})
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}
func (r *EmployeeRepo) DeleteEmployeeByID(empID string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(),
		bson.D{{key:"employee_id", value:empID}})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
	
}

func (r *EmployeeRepo) DeleteAllEmployee() (int64, error) {
	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return 0, err
	}