# Golang REST API
> Simple RESTful API to CRUD post. using mysql database and gorm as the ORM

## Quick Starts
```bash
# install mux router
$ go get -u github.com/gorilla/mux

# install dotenv
$ go get -u github.com/alexsasharegan/dotenv

# install mysql driver
$ go get -u github.com/go-sql-driver/mysql

# install gorm
$ go get -u github.com/jinzhu/gorm
```

```bash
$ go build

$ ./go-gorm-restapi
```
And open localhost:8081 in your browser

## Endpoints
### Get All Posts
``` bash
GET api/posts
```
### Get Single Post
``` bash
GET api/posts/{id}
```

### Delete Post
``` bash
DELETE api/posts/{id}
```

### Create Post
``` bash
POST api/posts

# Request sample
# {
#   "Title":"Post Three",
#   "Body":"Body of Post Three"
# }
```

### Update Post
``` bash
PUT api/posts/{id}

# Request sample
# {
#   "Title":"Updated Post",
#   "Body":"With Updated Body"
# }

```