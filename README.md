# Simple GO Lang REST API

> Simple RESTful API to create, read, update and delete books. No database implementation yet

## Quick Start


``` bash
# Install mux router
go get -u github.com/gorilla/mux
```

``` bash
go build
./goRestApi
```

## Endpoints
    I think after this experiment, the Golan is just Bullshit HHH.
    friends, please keep using PHP or Django or nodeTs

### Get All Books
``` bash
GET api/books
```
### Get Single Book
``` bash
GET api/book/{id}
```

### Delete Book
``` bash
DELETE api/book/{id}
```

### Create Book
``` bash
POST api/book

# Request sample
# {
#   "isbn":"89",
#   "title":"Book Three",
#   "author":{"firstname":"Elteyab",  "lastname":"Ali"}
# }
```

### Update Book
``` bash
PUT api/book/{id}

# Request sample
# {
#   "isbn":"89",
#   "title":"Updated Title",
#   "author":{"firstname":"Elteyab",  "lastname":"Hassan"}
# }

```


```

## App Info

### Author

Elteyab Hassan
[DetaElectronic](http://www.DetaElectroinc.sd)

### Version

1.0.0

### License

This project is licensed under the BSD License
