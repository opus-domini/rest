package database

import (
	"time"

	"github.com/opus-domini/fast-shot/examples/server/model"
)

type (
	Namespace string
	Table     map[uint]model.Model
)

const (
	UserNamespace     Namespace = "users"
	ResourceNamespace Namespace = "resources"
	TupleNamespace    Namespace = "tuples"
)

// Storage In-memory
var Storage = map[Namespace]Table{
	UserNamespace: {
		0: &model.User{
			ID:        0,
			Name:      "Ada",
			Birthdate: time.Date(2020, 7, 21, 0, 0, 0, 0, time.UTC),
		},
		1: &model.User{
			ID:        1,
			Name:      "Max",
			Birthdate: time.Date(2023, 4, 11, 0, 0, 0, 0, time.UTC),
		},
		2: &model.User{
			ID:        2,
			Name:      "Ivy",
			Birthdate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		3: &model.User{
			ID:        3,
			Name:      "Sam",
			Birthdate: time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		},
	},
	ResourceNamespace: {
		0: &model.Resource{
			ID:        0,
			Type:      "book",
			Data:      "The Go Programming Language",
			Timestamp: time.Date(2015, 10, 26, 0, 0, 0, 0, time.UTC),
		},
		1: &model.Resource{
			ID:        1,
			Type:      "movie",
			Data:      "The Matrix",
			Timestamp: time.Date(1999, 3, 31, 0, 0, 0, 0, time.UTC),
		},
		2: &model.Resource{
			ID:        2,
			Type:      "music",
			Data:      "The Dark Side of the Moon",
			Timestamp: time.Date(1973, 3, 1, 0, 0, 0, 0, time.UTC),
		},
	},
	TupleNamespace: {
		0: &model.Tuple{
			ID:    0,
			Key:   "key-1",
			Value: "value-1",
		},
		1: &model.Tuple{
			ID:    1,
			Key:   "key-2",
			Value: "value-2",
		},
		2: &model.Tuple{
			ID:    2,
			Key:   "key-3",
			Value: "value-3",
		},
	},
}
