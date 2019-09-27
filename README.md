# Social Network

## Summary

Here is a mini server to manage a basic social network. 
You can: 
* Get players' profile
* Add and list friends
* Report player with a motive 

I realise only the required needs in the subject, so it is not a complete version for production(see [To improve points.Other endpoints](#other-endpoints) for the detail).

## How to use it

Because it is not yet a complete server, this seed data is inserted into the database when you start the server:
* a profile with normal status 
* a profile with special status 
* a session for the normal player 
* a session for special player 

You can get the seed data from the log, then use it to test.

For example:
```
# go build
# ./app
2019/09/27 20:21:38 ---------------------------------------------------- 
2019/09/27 20:21:38 ---------------------------------------------------- 
2019/09/27 20:21:38 ---------------------------------------------------- 
2019/09/27 20:21:38 test data is inserted -> 
2019/09/27 20:21:38 normal profile: &{ID:3c988595-b678-448c-b306-d280fbde69a2 Name:getprofiletest_name1 CreatedAt:2019-09-27 20:21:38.222994 +0200 CEST m=+0.000714452 UpdatedAt:2019-09-27 20:21:38.222994 +0200 CEST m=+0.000714452 Status:n Age:0 Sex:0} 
2019/09/27 20:21:38 ---------------------------------------------------- 
2019/09/27 20:21:38 special profile: &{ID:23622678-065a-4da0-b943-10ddfae57859 Name:getprofiletest_name2 CreatedAt:2019-09-27 20:21:38.223209 +0200 CEST m=+0.000930003 UpdatedAt:2019-09-27 20:21:38.223209 +0200 CEST m=+0.000930003 Status:s Age:0 Sex:0} 
2019/09/27 20:21:38 ---------------------------------------------------- 
2019/09/27 20:21:38 normal profile's session: &{ID:390351ea-b05b-4c2b-9f3f-f8bedd425aa7 ProfileID:3c988595-b678-448c-b306-d280fbde69a2 CreatedAt:2019-09-27 20:21:38.223211 +0200 CEST m=+0.000931538 ExpiredAt:2019-09-27 20:31:38.223211 +0200 CEST m=+600.000931637} 
2019/09/27 20:21:38 ---------------------------------------------------- 
2019/09/27 20:21:38 special profile's session: &{ID:aa3eecfb-f2b7-4340-b89e-2e74d43e60f4 ProfileID:23622678-065a-4da0-b943-10ddfae57859 CreatedAt:2019-09-27 20:21:38.223215 +0200 CEST m=+0.000935920 ExpiredAt:2019-09-27 20:31:38.223215 +0200 CEST m=+600.000936004} 
2019/09/27 20:21:38 ---------------------------------------------------- 
2019/09/27 20:21:38 ---------------------------------------------------- 
2019/09/27 20:21:38 ---------------------------------------------------- 
2019/09/27 20:21:38 ---------------------------------------------------- 
```

You can then try to get normal player's profile using special player's session with `curl`
```
curl -X -v GET http://localhost:8080/profiles/get?session=aa3eecfb-f2b7-4340-b89e-2e74d43e60f4 -d '["3c988595-b678-448c-b306-d280fbde69a2"]'

{"profiles":[{"id":"4f3c186e-941a-4a1f-8823-39dd059218fe","name":"getprofiletest_name1","created_at":"2019-09-27T20:27:59.880304+02:00","updated_at":"2019-09-27T20:27:59.880304+02:00","status":"n","Age":0,"Sex":0}],"ok":true}
```

## API document 

`BASE_URL` is http:localhost if you run it in your local machine,
it could be an IP address if you run it via Docker, use `docker-machine ip default` to get the IP.

In case of failure, we get an error response like 

```json
{
    "ok": false,
    "error": "ERROR_MESSAGE"
}
```

### Get profiles in a batch

* **Description**
Get a set of profiles by the given profile IDs.

* **URL** GET `BASE_URL/profiles/get?session=SESSION_ID`

* **Payload** 
```json
["PROFILE_ID1","PROFILE_ID2","PROFILE_ID3" ...]
```

* **Response**
```json
{
    "ok": true,
    "profiles": [
        {
            "id": "PROFILE_ID1",
            "name": "PROFILE_NAME1",
            "created_at": "2019-09-22 10:22:11+00",
            "updated_at": "2019-09-23 11:33:11+00",
            "status": "n",
            ...
        },
        {
            "id": "PROFILE_ID2",
            "name": "PROFILE_NAME2",
            "created_at": "2019-09-01 09:22:11+00",
            "updated_at": "2019-09-02 12:33:11+00",
            "status": "s",
            ...
        }
    ]
}
```

### Ask friends

* **Description**
Ask be friend with another player.

* **URL** POST `BASE_URL/friends/submit?session=SESSION_ID`

* **Payload** 
```json
{
    "target_id": "PROFILE_ID",
    "message": "OPTIONAL"
}
```

* **Response**
```json
{
    "ok": true,
   "created": "2019-09-01 09:22:11+00",
   "message": "OPTIONAL"
}
```

### List my ask friends requests

* **Description**
List the ask friends requests that are made by other players

* **URL** GET `BASE_URL/friends/treate_submit?session=SESSION_ID`

* **Payload** NONE

* **Response**
```json
{
    "ok": true,
    "requests": [
        {
            "id": "INTERNAL_FRIEND_RELATION_ID1",
            "asker": "PROFILE_ID1",
            "receiver": "MY_PROFILE_ID",
            "created_at": "2019-09-01 09:22:11+00"
        },
        {
            "id": "INTERNAL_FRIEND_RELATION_ID2",
            "asker": "PROFILE_ID2",
            "receiver": "MY_PROFILE_ID",
            "created_at": "2019-09-03 10:22:11+00"
        },
    ]
}
```

### Treat a ask friends request

* **Description**
Accept or refuse an ask friends request.

* **URL** POST `BASE_URL/friends/treate_submit?session=SESSION_ID`

* **Payload** 
```json
{
    "action_id": "ASK_FRIENDS_REQUEST_ID",
    "action": "COULD BE i/r/a",
}
```

PS: the value of the field `action` is enum string: 
"i": ignore
"a": accept
"r": refuse

* **Response**

```json
{
    "ok": true
}
```

### List my friends

* **Description**
List the profiles of a player's friends. This endpoint support pagination.

* **URL** GET `BASE_URL/friends/list?session=SESSION_ID`

* **Payload** 
```json
{
    "target_id": "PROFILE_ID",
    "cursor": "TOKEN_OF_PAGE",
    "limit": "MAX_ENTITIES_PER_PAGE"
}
```

* **Response**

```json
{
    "ok": true,
    "friends": [
        {
            "id": "INTERNAL_FRIEND_RELATION_ID1",
            "my_id": "PROFILE_ID",
            "friend_id": "FRIEND_PROFILE_ID1",
            "created_at": "2019-09-01 09:22:11+00"
        },
        {
            "id": "INTERNAL_FRIEND_RELATION_ID2",
            "my_id": "PROFILE_ID",
            "friend_id": "FRIEND_PROFILE_ID2",
            "created_at": "2019-09-10 12:23:11+00"
        }
    ],
    "cursor": "TOKEN_OF_PAGE",
    "limit": "MAX_ENTITIES_PER_PAGE"
}
```

### Report player

* **Description**
Report a player with a motive.

* **URL** GET `BASE_URL/report/create?session=SESSION_ID`

* **Payload** 
```json
{
    "target_id": "PROFILE_ID",
    "motive": "MOTIVE",
}
```

* **Response**

```json
{
    "ok": true
}
```

## Tests

### Unit test

I only write the unit tests for the small and independant function, instead of add it for every function.
Because 
* there are service test and integration test to test the complex logic
* the current code pattern should be improved (see [To improve points.Code Pattern](#code-pattern) for the detail)

### Service test

This kind of test is to test a whole service business.

`profiles_get_test.go` is an example, and I should add the service tests for every endpoint.

### Integration test

This kind of test is to test some complex scenarios, that contain **more than one** service business. 

`friends_list_test.go` is an example, it tests the following scenarios:

* a normal player asks to be friends of a special player;
* the normal player re-asks the same request;
* the special player gets all his ask friends request then accepts them;
* the normal player lists all his friends;
* the special player lists all his friends;
* the normal player lists the friends of a special player;
* the special player lists the friends of a normal player.

## To improve points

### Code Pattern

I should split the files into different folders, for example:
```
| main.go
|
| handlers  ____ profiles_get.go
             |__ friends_list.go
             |__ ...
|
| types     ____ db_models.go
             |__ payloads.go
|
| ...
```

And the code in the handler files (ex:profiles_get.go) is contain the long function, they should be splited into small function.

### Other endpoints

Currently I seed some data(session, profile) into the fake db for the test, then there must be many other endpoints are necessary to process the social networks.

Fro example:
 * endpoints to manage session
 * endpoints to manege notification: for example when I ask be friends of another play, I must be notified when my request is accepted or refused.
 * CRUD profile

### Generic code

There are a lot of codes are very similar, specially for the part of handlers.  I should write some generic codes for them, to avoid of too many copy/paste codes.

### Test coverage

The current tests cover 67.3% code, it should be more (at least 80%).

### Docker image

We can then deploy the image to a server, for example to google cloud/AWS.

### Database

I implement only a fake db to simulate a nosql database, because this is a fast way to show my idea. But I should add a real database (ex: DynamoDB of AWS, Datastore of GCP, Open source) when I have time.

## I should do also 

Here are the importans points that are not required in the subject, but I think they are necessary for the production.

### Open source

I only a 2 external libraries in the current code, and they are used only for test.
It is possible to use more open source to make the code more simple.
For example, I can use a web performance (like [gin](https://github.com/gin-gonic/gin)) to manage the handler stuff.

### Metrics

It allows us to know better the server status.

* average execution time for each endpoint 
* the number of call for each endpoint 
* the number/scale of 4xx/5XX response 
* etc.

### Logs

* Alerts

For some expected cases, we should be alerted, specially in case of failure.

* Stacktraces

In case of failure, it allows me to find where is the problem.

### All TODO comments in the code