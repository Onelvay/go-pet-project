package controller

import (
	"context"
	"errors"
	"fmt"
	"log"

	mongoDB "github.com/Onelvay/docker-compose-project/db/mongoDB"
	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductDBController struct {
	db          *mongo.Collection
	mongoCtx    context.Context
	redisClient *redis.Client
}

func NewProductDbController(db *mongoDB.MongoDB, redis *redis.Client) *ProductDBController {
	return &ProductDBController{db.Db, db.Ctx, redis}
}

var Products []domain.Product

func (p *ProductDBController) GetBookById(id uint) (domain.Product, error) {
	product, err := productExistInRedis(p.redisClient, id)
	if err == nil {
		return product, nil
	} else {
		log.Println(err)
	}
	err = p.db.FindOne(p.mongoCtx, bson.M{"id": id}).Decode(&product)
	if err != nil {
		return product, err
	}
	if product.Id == 0 {
		return product, errors.New(fmt.Sprint("no product with id:", id))
	}
	saveBookInRedis(p.redisClient, product)
	return product, nil
}

// func (r *BookstorePostgres) GetBookById(id string) (domain.Book, error) {
// 	val, err := r.redisClient.Get(id).Result()
// 	if err == nil {
// 		err = json.Unmarshal([]byte(val), &book)
// 		if err != nil {
// 			return domain.Book{}, errors.New("problem with unmarshalling in postgresBookDB")
// 		}
// 		return book, nil
// 	}

// 	res := r.Db.Where("id = ?", id).Find(&book)
// 	if res.RowsAffected == 0 {
// 		return domain.Book{}, errors.New("book not found")
// 	}

// 	saveBookInRedis(r.redisClient)

// 	return book, nil
// }
// func (r *BookstorePostgres) GetBooksByName(name string) ([]domain.Book, error) {
// 	val, err := r.redisClient.Get(name).Result()
// 	if err == nil {
// 		err = json.Unmarshal([]byte(val), &books)
// 		if err == nil {
// 			return books, nil
// 		}
// 	}
// 	res := r.Db.Where("name = ?", name).Find(&books)
// 	if res.RowsAffected == 0 {
// 		return []domain.Book{}, fmt.Errorf("no books with name %s", name)
// 	}

// 	saveBooksInRedis(r.redisClient, name)
// 	return books, nil
// }

// func (r *BookstorePostgres) GetBooks(sorted bool) ([]domain.Book, error) {
// 	var res *gorm.DB
// 	if sorted {
// 		res = r.Db.Order("price").Find(&books)
// 	} else {
// 		res = r.Db.Order("price desc").Find(&books)
// 	}
// 	return books, res.Error
// }

// func (r *BookstorePostgres) DeleteBookById(id string) error {
// 	res := r.Db.Where("id=?", id).Delete(&domain.Book{})
// 	return res.Error
// }
// func (r *BookstorePostgres) CreateBook(name string, price float64, descr string) error {
// 	byteid := uuid.New()
// 	id := strings.Replace(byteid.String(), "-", "", -1)

// 	res := r.Db.Create(&domain.Book{
// 		Id:          id,
// 		Name:        name,
// 		Description: descr,
// 		Price:       price,
// 	})

// 	saveBookInRedis(r.redisClient)

// 	return res.Error
// }

// func (r *BookstorePostgres) UpdateBook(id string, name string, desc string, price float64) error {
// 	_, res := r.GetBookById(id)
// 	if res == nil {
// 		if name != "" {
// 			book.Name = name
// 		}
// 		if desc != "" {
// 			book.Description = desc
// 		}
// 		if price != 0 {
// 			book.Price = price
// 		}
// 		res := r.Db.Save(&book)
// 		saveBookInRedis(r.redisClient)
// 		return res.Error
// 	}
// 	return res
// }

// func saveBooksInRedis(r *redis.Client, name string) {
// 	j, err := json.Marshal(books)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	err = r.Set(name, j, 0).Err()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
