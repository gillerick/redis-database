package db

import "errors"

//Store is the main database map that maps key to value e.g. {"a": "foo", "b": "foo", "c": "bar"}
type Store map[string]string

//ReverseStore is the reverse of Store and maps value to key(s) e.g. {"foo" : ["a", "b"]}
type ReverseStore map[string][]string

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

func set(words []string, store Store, reverseStore ReverseStore) (string, error) {
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

func del(words []string, store Store) (int, error) {
	if len(words) < 2 {
		return 0, errors.New("Invalid DEL command. Correct format: DEL key")
	}

	keys := words[1:]
	count := 0
	for _, key := range keys {
		delete(store, key)
		count += 1
	}

	return count, nil

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
