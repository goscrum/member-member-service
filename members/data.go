package members

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type DAO struct {
	collection mongo.Collection
}

func (dao *DAO) add(member Member, ctx context.Context) (Member, error) {
	_, err := dao.collection.InsertOne(ctx, member)
	if err != nil {
		return Member{}, err
	}
	// TODO atribuir member.ID = result.InsertedID
	return member, nil
}
