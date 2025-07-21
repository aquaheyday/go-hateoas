# go-hateoas

A lightweight Go library for building HATEOAS-compliant JSON APIs.

## ✨ Features

- Wraps your resource with `_links` in HAL-style  
- Simple builder-style chaining (`Wrap().Link().Build()`)  
- JSON encoding helper  
- Lightweight, no external dependencies  

## 🚀 Install

```bash
go get github.com/aquaheyday/go-hateoas
```

## 📦 Basic Usage

```go
import "github.com/aquaheyday/go-hateoas"

type Book struct {
	ID    int
	Title string
}

book := Book{ID: 1, Title: "Clean Architecture"}

res := hateoas.Wrap(book).
	Link("self", "/books/1").
	Link("author", "/authors/42").
	Build()

jsonBytes, _ := json.Marshal(res)
```

**Output**:

```json
{
  "data": {
    "id": 1,
    "title": "Clean Architecture"
  },
  "_links": {
    "self": {
      "href": "/books/1"
    },
    "author": {
      "href": "/authors/42"
    }
  }
}
```

## 🧪 With Gin (Optional)

```go
hateoas.RespondWithHATEOAS(c, http.StatusOK, user, links)
```

## 📜 License

MIT
