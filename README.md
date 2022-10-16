## Redis DB
A simple Redis-like in-memory database that can perform the following operations:

| Command & Usage  | Use |
| ------------- | ------------- |
| **SET** key value  | Set the string value of a key to the specified value  |
| **GET** key  | Get the value of a specified key  |
| **DELETE** key  | Delete the specified key  |
| **COUNT** value   | Return the number of times a value occurs  |
| **EXISTS** key  | Determine if a key exists  |
| **INCR** key [increment]  | Increment the integer value of a specified key by 1. Optionally takes in a value |


## Demos
### INCR
```text
>> SET age 27
OK
>> GET age
27
>> INCR age 
OK
>> GET age
28

```