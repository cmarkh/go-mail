# Mail

Provides basic email sending functionality 

`go get` it as `github.com/cmarkh/go-mail`, import it as
`"github.com/cmarkh/go-mail"`, use it as `mail`.

## Usage

Send from gmail (I provide host and port):
```go
account := mail.NewGmail("First LastName", "from@example.com", "password")
```

Specifiy host: 
```go
account := mail.Account{"First LastName", "from@example.com", "password", "smtp.example.com", 587}
```

One recipient: 
```go
err := account.Send("Subject", "Body", "to@example.com")
```

Multiple recipients: 
```go
err := account.Send("Subject", "Body", "to@example.com", "another@example.com")
```