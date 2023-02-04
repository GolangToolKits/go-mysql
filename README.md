# go_mysql
An easy to use Mockable db interface for mysql

```go

    mm := &MyDB{
	    Host:     "localhost:3306",
	    User:     "test",
	    Password: "test",
	    Database: "test",				
    }
    m := mm.New()
    m.Connect()
    m.Insert(query, args...)
    m.Update(query, args...)
    m.Get(query, args...)
    m.GetList(query, args...)
    m.Delete(query, args...)

```

## Also Supports Transactions

```go
 mm := &MyDB{
	    Host:     "localhost:3306",
	    User:     "test",
	    Password: "test",
	    Database: "test",				
    }
    m := mm.New()
    m.Connect()
    tr := m.BeginTransaction()
    tr.Insert(query, args...)
    tr.Update(query, args...)
    tr.Get(query, args...)
    tr.GetList(query, args...)
    tr.Delete(query, args...)
    tr.Commit()
    //Or
    tr.Rollback()


```
