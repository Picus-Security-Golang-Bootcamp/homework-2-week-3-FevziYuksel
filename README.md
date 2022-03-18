## Homework | Week 3

Simple library application 
We have a book list
Book filed are like these;
```
- Book ID
- Book Name
- Page Number
- Stock Number
- Price
- Stock Code
- ISBN
- Author Info (ID and Name)
```

1. List all books (list)
2. List books which contain entered word (search)
3. Print books by ID
4. Delete books by entered ID. 
5. Buy a book by ID and print its last fields.

If the command is invalid Usage would be printed. 


### list command
```
go run main.go -command list
```

### search command 
```
go run main.go -command search -name <bookName>
```

### get command
```
go run main.go -command get -ID <bookID>
```

### delete command
```
go run main.go -command delete -ID <bookID>
```

### buy command
```
go run main.go -command buy -ID <bookID> -quantity <quantity>
```

