
register
```
127.0.0.1:8877/api/user
post
data
{
	"username":"lxw",
	"password":"123456",
	"nickName":"shuaishuai",
	"email":"1234@qq.com"
}
```

login
```
post
{
	"username":"lxw",
	"password":"123456",
	"email":"1234@qq.com"
}
```


```
add
127.0.0.1:8877/api/group
post
{
	"groupname":"ss"
}
```


```
del
127.0.0.1:8877/api/group
put
{
	"id":1
}
```

```
add
127.0.0.1:8877/api/friend
{
	"friendId":2,
	"groupId":2
}
```

```
del
127.0.0.1:8877/api/friend
put
{
	"friendId":2
}

```

```
add
127.0.0.1:8877/api/flock
post
{
	"name":"laosiji"
}

```


```
del
127.0.0.1:8877/api/flock
put
{
	"id":1
}
```

