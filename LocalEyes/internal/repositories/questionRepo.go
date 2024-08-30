package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"localEyes/internal/db"
	"localEyes/internal/interfaces"
	"localEyes/internal/models"
)

type MongoQuestionRepository struct {
	Collection interfaces.CollectionInterface
}

func NewMongoQuestionRepository() interfaces.QuestionRepository {
	return &MongoQuestionRepository{
		Collection: db.NewCollectionWrapper(&db.MongoCollectionWrapper{Collection: db.GetQuestionsCollection()}),
	}
}

func (r *MongoQuestionRepository) Create(question *models.Question) error {
	_, err := r.Collection.InsertOne(context.Background(), question)
	return err
}

func (r *MongoQuestionRepository) GetAllQuestions() ([]*models.Question, error) {
	var questions []*models.Question

	cursor, err := r.Collection.Find(context.Background(), bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}(cursor, context.Background())

	for cursor.Next(context.Background()) {
		var question models.Question
		if err := cursor.Decode(&question); err != nil {
			return nil, err
		}
		questions = append(questions, &question)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *MongoQuestionRepository) DeleteOneDoc(filter interface{}) error {
	_, err := r.Collection.DeleteOne(context.Background(), filter)
	return err
}

func (r *MongoQuestionRepository) DeleteByPId(PId primitive.ObjectID) error {
	_, err := r.Collection.DeleteMany(context.Background(), bson.M{"post_id": PId})
	return err
}

func (r *MongoQuestionRepository) GetQuestionsByPId(PId primitive.ObjectID) ([]*models.Question, error) {
	var questions []*models.Question

	cursor, err := r.Collection.Find(context.Background(), bson.M{"post_id": PId}, options.Find())
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}(cursor, context.Background())

	for cursor.Next(context.Background()) {
		var question models.Question
		if err := cursor.Decode(&question); err != nil {
			return nil, err
		}
		questions = append(questions, &question)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *MongoQuestionRepository) UpdateQuestion(QId primitive.ObjectID, answer string) error {
	filter := bson.M{"q_id": QId}
	updates := bson.M{
		"replies": answer,
	}
	update := bson.M{"$push": updates}
	_, err := r.Collection.UpdateFields(context.Background(), filter, update)
	return err
}
