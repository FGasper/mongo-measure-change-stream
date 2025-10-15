package agg

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Eq(comparands ...any) bson.D {
	return bson.D{{"$eq", comparands}}
}

func In[T any](needle any, haystack ...T) bson.D {
	return bson.D{{"$in", bson.A{needle, haystack}}}
}

func BSONSize(ref any) bson.D {
	return bson.D{{"$bsonSize", ref}}
}

func Type(ref any) bson.D {
	return bson.D{{"$type", ref}}
}

func Concat(refs ...any) bson.D {
	return bson.D{{"$concat", refs}}
}

// ---------------------------------------------

type Not struct {
	Ref any
}

var _ bson.Marshaler = Not{}

func (n Not) MarshalBSON() ([]byte, error) {
	return bson.Marshal(bson.D{
		{"$not", n.Ref},
	})
}

// ---------------------------------------------

type And []any

var _ bson.Marshaler = And{}

func (a And) MarshalBSON() ([]byte, error) {
	return bson.Marshal(bson.D{
		{"$and", []any(a)},
	})
}

// ---------------------------------------------

type Or []any

var _ bson.Marshaler = Or{}

func (o Or) MarshalBSON() ([]byte, error) {
	return bson.Marshal(bson.D{
		{"$or", []any(o)},
	})
}

// ---------------------------------------------

type SubstrBytes [3]any

var _ bson.Marshaler = SubstrBytes{}

func (s SubstrBytes) MarshalBSON() ([]byte, error) {
	return bson.Marshal(bson.D{
		{"$substr", []any(s[:])},
	})
}

// ---------------------------------------------

type MergeObjects []any

var _ bson.Marshaler = MergeObjects{}

func (m MergeObjects) MarshalBSON() ([]byte, error) {
	return bson.Marshal(bson.D{
		{"$mergeObjects", m},
	})
}

// ---------------------------------------------

type Cond struct {
	If, Then, Else any
}

var _ bson.Marshaler = Cond{}

func (c Cond) D() bson.D {
	return bson.D{
		{"$cond", bson.D{
			{"if", c.If},
			{"then", c.Then},
			{"else", c.Else},
		}},
	}
}

func (c Cond) MarshalBSON() ([]byte, error) {
	return bson.Marshal(c.D())
}

// ---------------------------------------------

type Map struct {
	Input, As, In any
}

var _ bson.Marshaler = Map{}

func (m Map) D() bson.D {
	return bson.D{
		{"$map", bson.D{
			{"input", m.Input},
			{"as", m.As},
			{"in", m.In},
		}},
	}
}

func (m Map) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m.D())
}
