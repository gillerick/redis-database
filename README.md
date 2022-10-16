## Redis DB

A simple Redis-like in-memory database that can perform the following operations:

| Command & Usage  | Use |
| ------------- | ------------- |
| **SET** key value  | Set the string value of a key to the specified value  |
| **GET** key  | Get the value of a specified key  |
| **DELETE** key  | Delete the specified key  |
| **COUNT** value   | Return the number of times a value occurs  |
| **EXISTS** key  | Determine if a key exists  |
| **INCR** key [increment]  | Increment the integer value of a specified key by 1. Optionally takes in an increment value |
| **DECR** key [decrement]  | Decrement the integer value of a specified key by 1. Optionally takes in a decrement value |

## Demos

### INCR

```text
>> SET age 27
OK
>> GET age
27
>> INCR age 
28
>> GET age
28
 
>> SET year 2022
OK
>> INCR year 7
2029
```

### DECR

```text
>> SET seconds 60
OK
>> GET seconds 
60
>> DECR seconds
59
>> DECR seconds 30
29
>> DECR minutes
The provided key does not exist
```
