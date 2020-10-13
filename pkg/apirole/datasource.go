package apirole

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// DataSource is the interface
type DataSource interface {
	GetRoleAll() ([]Roles, error)
	GetRoleById(id string) (Roles, error)
	AddRole(data *Roles) error
	UpdateRole(data Roles) error
	DeleteRole(id string) error
	CheckRoleDisplayExist(display string) Roles

	GetRoleUserAll() ([]RoleUser, error)
	GetRoleUserById(id string) (RoleUser, error)
	GetRoleUserByRoleId(id string) (RoleUser, error)
	AddRoleUser(data *RoleUser) error
	UpdateRoleUser(data RoleUser) error
	DeleteRoleUser(id string) error
	RoleUserExist(data RoleUser) bool

	GetPolicyAll() ([]Policy, error)
	GetPolicyById(id string) (Policy, error)
	GetPolicyByRoleId(id string) (Policy, error)
	GetPolicyListByRoleId(id string) ([]Policy, error)
	UpdatePolicy(data Policy) error
	PolicyExist(data Policy) bool
}

type dataSource struct {
	MgoDB *mongo.Database
}

func (d *dataSource) GetRoleAll() ([]Roles, error) {
	collection := d.MgoDB.Collection("roles")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return []Roles{}, err
	}
	var results []Roles
	for cursor.Next(context.Background()) {
		var elem Roles
		if err := cursor.Decode(&elem); err == nil {
			results = append(results, elem)
		}
	}
	if err := cursor.Err(); err != nil {
		return []Roles{}, err
	}
	_ = cursor.Close(context.Background())
	return results, nil
}

func (d *dataSource) GetRoleById(id string) (Roles, error) {
	collection := d.MgoDB.Collection("roles")
	var result Roles
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Roles{}, err
	}
	filter := bson.M{"_id": objId}
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return Roles{}, err
	}
	return result, nil
}

func (d *dataSource) AddRole(data *Roles) error {
	collection := d.MgoDB.Collection("roles")
	result, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	data.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *dataSource) UpdateRole(data Roles) error {
	collection := d.MgoDB.Collection("roles")
	filter := bson.M{"_id": data.ID}
	update := bson.M{"$set": data}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (d *dataSource) DeleteRole(id string) error {
	collection := d.MgoDB.Collection("roles")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteMany(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}

func (d *dataSource) CheckRoleDisplayExist(display string) Roles {
	collection := d.MgoDB.Collection("roles")
	result := Roles{}
	filter := bson.D{{"display", display}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return Roles{}
	}
	return result
}

func (d *dataSource) GetRoleUserAll() ([]RoleUser, error) {
	collection := d.MgoDB.Collection("role_user")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return []RoleUser{}, err
	}
	var results []RoleUser
	for cursor.Next(context.Background()) {
		var elem RoleUser
		if err := cursor.Decode(&elem); err == nil {
			results = append(results, elem)
		}
	}
	if err := cursor.Err(); err != nil {
		return []RoleUser{}, err
	}
	_ = cursor.Close(context.Background())
	return results, nil
}

func (d *dataSource) GetRoleUserById(id string) (RoleUser, error) {
	collection := d.MgoDB.Collection("role_user")
	var result RoleUser
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return RoleUser{}, err
	}
	filter := bson.M{"_id": objId}
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return RoleUser{}, err
	}
	return result, nil
}

func (d *dataSource) GetRoleUserByRoleId(id string) (RoleUser, error) {
	collection := d.MgoDB.Collection("role_user")
	var result RoleUser
	filter := bson.D{{"roleId", id}}
	log.Println("filter", filter)
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return RoleUser{}, err
	}
	return result, nil
}

func (d *dataSource) AddRoleUser(data *RoleUser) error {
	collection := d.MgoDB.Collection("role_user")
	result, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	data.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *dataSource) UpdateRoleUser(data RoleUser) error {
	collection := d.MgoDB.Collection("role_user")
	filter := bson.M{"_id": data.ID}
	update := bson.M{"$set": data}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (d *dataSource) DeleteRoleUser(id string) error {
	collection := d.MgoDB.Collection("role_user")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteMany(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}

func (d *dataSource) RoleUserExist(data RoleUser) bool {
	collection := d.MgoDB.Collection("role_user")
	var result RoleUser
	filter := bson.D{{"userId", data.UserID}, {"roleId", data.RoleID}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return false
	}
	return result.UserID == data.UserID && result.RoleID == data.RoleID
}

func (d *dataSource) GetPolicyAll() ([]Policy, error) {
	collection := d.MgoDB.Collection("casbin_rule")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return []Policy{}, err
	}
	var results []Policy
	for cursor.Next(context.Background()) {
		var elem Policy
		if err := cursor.Decode(&elem); err == nil {
			results = append(results, elem)
		}
	}
	if err := cursor.Err(); err != nil {
		return []Policy{}, err
	}
	_ = cursor.Close(context.Background())
	return results, nil
}

func (d *dataSource) GetPolicyById(id string) (Policy, error) {
	collection := d.MgoDB.Collection("casbin_rule")
	result := Policy{}
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Policy{}, err
	}
	filter := bson.M{"_id": objId}
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return Policy{}, err
	}
	return result, nil
}

func (d *dataSource) GetPolicyByRoleId(id string) (Policy, error) {
	collection := d.MgoDB.Collection("casbin_rule")
	var result Policy
	filter := bson.D{{"v0", id}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return Policy{}, err
	}
	return result, nil
}

func (d *dataSource) GetPolicyListByRoleId(id string) ([]Policy, error) {
	collection := d.MgoDB.Collection("casbin_rule")
	results := []Policy{}
	filter := bson.D{{"v0", id}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return []Policy{}, err
	}
	for cursor.Next(context.Background()) {
		var elem Policy
		if err := cursor.Decode(&elem); err == nil {
			results = append(results, elem)
		}
	}
	if err := cursor.Err(); err != nil {
		log.Println(err)
		return []Policy{}, err
	}
	_ = cursor.Close(context.Background())
	return results, nil
}

func (d *dataSource) UpdatePolicy(data Policy) error {
	collection := d.MgoDB.Collection("casbin_rule")
	filter := bson.M{"_id": data.ID}
	update := bson.M{"$set": data}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (d *dataSource) PolicyExist(data Policy) bool {
	collection := d.MgoDB.Collection("casbin_rule")
	var result Policy
	filter := bson.D{{"v0", data.RoleId}, {"v1", data.Path}, {"v2", data.Method}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return false
	}
	return result.RoleId == data.RoleId && result.Path == data.Path && result.Method == data.Method
}

// NewDataSource new instance
func NewDataSource(mgoDb *mongo.Database) DataSource {
	return &dataSource{
		MgoDB: mgoDb,
	}
}
