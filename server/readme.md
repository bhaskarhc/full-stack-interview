1. New user login by mobile number
2. VAlidate user Mobile number by matching OTP.
3. Successful login/signup generate JWT token
4. Using JWT token in Header , client can access user related data .
5. user can create single/multiple stores
6. list all the stores of client-user .




Postgresql table schema :
UserTable: 
```
                                    Table "public.userdata"
  Column  |       Type        | Collation | Nullable |                 Default                 
----------+-------------------+-----------+----------+-----------------------------------------
 index    | integer           |           | not null | nextval('userdata_index_seq'::regclass)
 userid   | character varying |           |          | 
 mobile   | bigint            |           |          | 
 token    | character varying |           |          | 
 name     | character varying |           |          | 
 location | character varying |           |          | 
 shoptype | character varying |           |          | 
 otp      | integer           |           |          | 
Indexes:
    "userdata_pkey" PRIMARY KEY, btree (index)

```

Shops Table :

```
                                    Table "public.shopdata"
   Column    |       Type        | Collation | Nullable |               Default                
-------------+-------------------+-----------+----------+--------------------------------------
 id          | integer           |           | not null | nextval('shopdata_id_seq'::regclass)
 itemgroupid | character varying |           |          | 
 name        | character varying |           |          | 
 rating      | character varying |           |          | 
 status      | character varying |           |          | 
 accountid   | character varying |           |          | 
 type        | character varying |           |          | 
 description | character varying |           |          | 
Indexes:
    "shopdata_pkey" PRIMARY KEY, btree (id)
```


pre-requsite to run code: 
- install latest version of Go and create $GOPATH in local
- install postgresql in local <or> run postgresql as Docker container and get access credentials .
- Set psql database credentails as env's ,
```
{ export DBHOST="localhost"; export DBPORT="5432"; export DBUSER="postgres"; export DBPASS=""; export DBNAME="postgres"; }
```

- clone code from github 
```
git clone https://github.com/bhaskarhc/full-stack-interview.git
```
- start the  server .
```
go run /server/main.go
```
----------- Server started with localhost as host and port 3000  ---------

1. New user login by mobile number :
  ```
  api: "/user/new" 
  method : "POST"
  body: {
    "mobile": 9xxxxxxxxxxxx
  }
  ```
  response : 
  ```
  ⇒  curl -d '{"mobile":9988776655}' -H 'Content-Type: application/json' http://localhost:3000/user/new
Successfully created user 
 {0  9988776655     9879}%      
 ```

 - new user created with his mobile number and generated OTP

2. VAlidate user Mobile number by matching OTP. \n 3.Successful login/signup generate JWT token

 - User need to validate his profile with OTP, then user will get a JWT token .

```
api: "/token" 
  method : "POST"
  body: {
    "mobile": 9xxxxxxxxxxxx,
    "otp": 9879
  }
  ```

  Response :

  ```
   {token : {8989898001 7943}} {token : {eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJjbGllbnQiOjg5ODk4OTgwMDEsImV4cCI6MTYxMjM2NzIyM30.FqHECg-kkEs8GA5JMm_K_oeWFSU2QzG4-AvOw4n5L9w }}
   ```
3.Using JWT token in Header , client can access user related data .

  - listDown all shops data related to user
  ```
  api: "/shops" 
  method : "GET"
  Header: "Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJjbGllbnQiOjg5ODk4OTgwMDEsImV4cCI6MTYxMjM2NzIyM30.FqHECg-kkEs8GA5JMm_K_oeWFSU2QzG4-AvOw4n5L9w"
  ```

  Response :
  ```json
  {
    "shops": [
        {
            "id": 6,
            "accountID": 0,
            "itemGroupID": "",
            "location": "",
            "rating": 0,
            "status": "active",
            "name": "laddu-shop",
            "description": "samall scale shop",
            "type": "sweet"
        },
        {
            "id": 7,
            "accountID": 0,
            "itemGroupID": "",
            "location": "",
            "rating": 0,
            "status": "active",
            "name": "laddu-shop",
            "description": "samall scale shop",
            "type": "sweet"
        },
        {
            "id": 8,
            "accountID": 0,
            "itemGroupID": "",
            "location": "",
            "rating": 0,
            "status": "active",
            "name": "laddu-shop",
            "description": "samall scale shop",
            "type": "sweet"
        },
        {
            "id": 9,
            "accountID": 0,
            "itemGroupID": "",
            "location": "",
            "rating": 0,
            "status": "active",
            "name": "laddu-shop",
            "description": "samall scale shop",
            "type": "sweet"
        },
        {
            "id": 10,
            "accountID": 0,
            "itemGroupID": "",
            "location": "",
            "rating": 0,
            "status": "active",
            "name": "chemist-log",
            "description": "large scale shop",
            "type": "chemist"
        }
    ]
}
```
5. user can create single/multiple stores
```

```
api: "/shop/add" 
  method : "POST"
  body: {
    "name": "ram pan shop",
    "type": "pan-shop",
    "status": "active",
    "description": "",
    "accountid": 10
}
  ```

  Response :
  ```
  successfully shopdata saved .. ! 
 {0 1   0 active ram pan shop  pan-shop}
 ```




