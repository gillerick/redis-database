package db

import (
	"errors"
	"strconv"
	"strings"
)

//Store is the main database map that maps key to value e.g. {"a": "foo", "b": "foo", "c": "bar"}
type Store map[string]string

//ReversedStore is the reverse of Store and maps value to key(s) e.g. {"foo" : ["a", "b"]}
type ReversedStore map[string][]string

func get(words []string, store Store) (string, error) {
	if len(words) != 2 {
		return "", errors.New("Invalid GET command. Correct format: GET [name]")
	}

	key := words[1]

	if value, ok := store[key]; ok {
		return value, nil
	} else {
		return "NULL", nil
	}
}

func set(words []string, store Store, reverseStore ReversedStore) (string, error) {
	if len(words) != 3 {
		return "", errors.New("Invalid SET command. Correct formart: SET [key] [newValue]")
	}

	key, newValue := words[1], words[2]
	oldValue := store[key]

	store[key] = newValue
	_, ok1 := store[key]

	if _, ok2 := reverseStore[oldValue]; ok2 {
		reverseStore[oldValue] = deleteStoreKeyFromReversedStore(key, reverseStore[newValue])
	}

	reverseStore[oldValue] = append(reverseStore[oldValue], key)
	value2 := reverseStore[oldValue][len(reverseStore[oldValue])-1]

	if !ok1 || value2 != key {
		return "", errors.New("Error in setting " + key)
	} else {
		return "OK", nil
	}

}

func del(words []string, store Store) (string, error) {
	if len(words) < 2 {
		return "0", errors.New("Invalid DEL command. Correct format: DEL key [key1] [key2]")
	}

	keys := words[1:]
	total := 0
	for _, key := range keys {
		delete(store, key)
		total += 1
	}

	return strconv.Itoa(total), nil

}

func count(words []string, store Store) (string, error) {
	if len(words) != 2 {
		return "", errors.New("Invalid COUNT command. Correct format: COUNT [value]")
	}
	value := words[1]
	total := strconv.Itoa(len(store[value]))
	return total, nil

}

func exists(words []string, store Store) (string, error) {
	if len(words) != 2 {
		return "", errors.New("Invalid EXISTS command. Correct format: EXISTS [key]")
	}

	key := words[1]
	if _, ok := store[key]; ok {
		return "1", nil
	} else {
		return "0", nil
	}
}

func deleteStoreKeyFromReversedStore(key string, keys []string) []string {
	for index, storeKey := range keys {
		if storeKey == key {
			keys = append(keys[:index], keys[index+1:]...)
			break
		}
	}
	return keys
}

func EvaluateCommand(line string, store *Store, reverseStore *ReversedStore) (string, error) {
	words := strings.Split(line, " ")

	switch command := strings.ToLower(words[0]); command {
	case "get":
		return get(words, *store)
	case "set":
		return set(words, *store, *reverseStore)
	case "del":
		return del(words, *store)
	case "help":
		return "HELP", nil
	case "?":
		return "HELP", nil
	case "count":
		return count(words, *store)
	case "exists":
		return exists(words, *store)
	default:
		return line, errors.New("Invalid command.")
	}

}
