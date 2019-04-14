package db

import (
	"errors"

	"github.com/over-engineering/msa/common/models/types"
	"github.com/over-engineering/msa/common/models/user"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoAccessor is an Accessor for MongoDB.
type MongoAccessor struct {
	session    *mgo.Session
	collection *mgo.Collection
}

// New returns a new MongoAccessor
func New(path, db, c string) *MongoAccessor {
	session, err := mgo.Dial(path)
	if err != nil {
		return nil
	}
	collection := session.DB(db).C(c)
	collection.Name = c
	return &MongoAccessor{
		session:    session,
		collection: collection,
	}
}

func (m *MongoAccessor) Close() error {
	m.session.Close()
	return nil
}

func idToObjectId(id types.UID) bson.ObjectId {
	return bson.ObjectId(string(id))
}

func objectIdToUId(objID bson.ObjectId) types.UID {
	return types.UID(objID.String())
}

func (m *MongoAccessor) checkCollection(s string) error {
	if m.collection.Name == s {
		return nil
	}
	return errors.New("collection name is not matching")
}

func (m *MongoAccessor) GetCharacter(id types.UID) (user.Character, error) {
	c := user.Character{}
	if err := m.checkCollection("character"); err != nil {
		return c, err
	}
	err := m.collection.Find(bson.M{"_id": idToObjectId(id)}).One(&c)
	return c, err
}

func (m *MongoAccessor) PostCharacter(c user.Character) error {
	if err := m.checkCollection("character"); err != nil {
		return err
	}
	objID := idToObjectId(c.User.ID)
	_, err := m.collection.UpsertId(objID, &c)
	return err
}

func (m *MongoAccessor) PutCharacter(id types.UID, c user.Character) error {
	if err := m.checkCollection("character"); err != nil {
		return err
	}
	return m.collection.UpdateId(idToObjectId(id), c)
}

func (m *MongoAccessor) DeleteCharacter(id types.UID) error {
	if err := m.checkCollection("character"); err != nil {
		return err
	}
	return m.collection.RemoveId(id)
}

// func (m *MongoAccessor) GetAll() ([]task.ID, error) {
// 	ids := []task.ID{}
// 	err := m.collection.Find(bson.M{}).All(&ids)
// 	return ids, err
// }

// func (m *MongoAccessor) Get(id task.ID) (task.Task, error) {
// 	t := task.Task{}
// 	err := m.collection.FindId(idToObjectId(id)).One(&t)
// 	return t, err
// }

// func (m *MongoAccessor) Put(id task.ID, t task.Task) error {
// 	return m.collection.UpdateId(idToObjectId(id), t)
// }

// func (m *MongoAccessor) Post(t task.Task) (task.ID, error) {
// 	objID := bson.NewObjectId()
// 	_, err := m.collection.UpsertId(objID, &t)
// 	return objectIdToId(objID), err
// }

// func (m *MongoAccessor) Delete(id task.ID) error {
// 	return m.collection.RemoveId(idToObjectId(id))
// }
