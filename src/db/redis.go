package db

//Store is the main database map that maps key to value e.g. {"a": "foo", "b": "foo", "c": "bar"}
type Store map[string]string

//ReverseStore is the reverse of Store and maps value to key(s) e.g. {"foo" : ["a", "b"]}
type ReverseStore map[string][]string
