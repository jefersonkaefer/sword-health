
# sword-health
### Tecnologies 
- Go!
- MySql
- Redis
- gRPC
- RabbitMQ
- gORM
- GinGonic


## Start project
```
docker-compose up
```
## API rest

### endpoints:

* Create user:
**[POST]** http://localhost:8000/user
body request:

```
{ 
	"first_name": "John", 
	"last_name": "doe", 
	"email": "jefka@mail.com", 
	"password": "oasdasdasdi", 
	"confirm_password": "oasdasdasdi", 
	"role": "manager" 
}
```
* Auth user:
**[POST]** http://localhost:8000/auth
body request:

```
{
	"email":  "jefka@mail.com",
	"password":  "oasdasdasdi"
}
```
* Create task:
**[POST]** http://localhost:8000/task
body request:

```
{
	"summary":  "My second Task"
}
```

* Get task:
**[GET]** http://localhost:8000/task/:id
body response:

```
{
	"id":  27,
	"summary":  "My second Task",
	"status":  "close",
	"when":  "03 Aug 21 22:15 UTC",
	"firstName":  "John",
	"lastName":  "doe",
	"email":  "jefka@mail.com"
}
```

* Get list task:
**[GET]** http://localhost:8000/task
body response:

```
{
    "tasks": [
        {
            "id": 3,
            "summary": "My second Task",
            "ownerId": 1,
            "when": "03 Aug 21 08:37 UTC",
            "firstName": "John",
            "lastName": "doe",
            "email": "John doe"
        },
        {
            "id": 4,
            "summary": "My second Task",
            "ownerId": 1,
            "firstName": "John",
            "lastName": "doe",
            "email": "John doe"
        },
        {
            "id": 5,
            "summary": "My second Task",
            "ownerId": 1,
            "when": "06 Aug 21 08:12 UTC",
            "firstName": "John",
            "lastName": "doe",
            "email": "John doe"
        },
        {
            "id": 6,
            "summary": "My second Task",
            "ownerId": 1,
            "firstName": "John",
            "lastName": "doe",
            "email": "John doe"
        },
        {
            "id": 7,
            "summary": "My second Task",
            "ownerId": 1,
            "firstName": "John",
            "lastName": "doe",
            "email": "John doe"
        },
        {
            "id": 8,
            "summary": "My second Task",
            "ownerId": 1,
            "firstName": "John",
            "lastName": "doe",
            "email": "John doe"
        },
        {
            "id": 9,
            "summary": "My second Task",
            "ownerId": 1,
            "firstName": "John",
            "lastName": "doe",
            "email": "John doe"
        },
        {
            "id": 10,
            "summary": "My second Task",
            "ownerId": 1,
            "firstName": "John",
            "lastName": "doe",
            "email": "John doe"
        },
        {
            "id": 11,
            "summary": "My second Task",
            "ownerId": 1,
            "firstName": "John",
            "lastName": "doe",
            "email": "John doe"
        },
        {
            "id": 12,
            "summary": "My second Task",
            "ownerId": 1,
            "firstName": "John",
            "lastName": "doe",
            "email": "John doe"
        }
    ]
}
```

* Close task:
**[PATCH]** http://localhost:8000/task/close
```

* Delete task:
**[Delete]** http://localhost:8000/task/:id

* Get Notification:
**[GET]** http://localhost:8000/notification
body response:

```
{
	"id":  26,
	"type": "task"
	"id_type": 27
	"status":  "close",
	"when":  "04 Aug 21 22:15 UTC",
	"firstName":  "John",
}
```
