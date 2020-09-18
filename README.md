# go-starter

Golang, PostgreSQL and Vue starter project.


## endpoints


```
POST /auth/regiter

{
    "name": "Foo Bar",
    "email": "...",
    "pwd": "..."
}
```
Return 200 on success


```
POST /auth/login
{
    "email": "...",
    "password": "..."
}
```
Returns JWT

```
POST /f/{function_name}

{"foo": 1, "bar": 42}
```
Run DB functions, returns `map` of result
 