# Sidecar

## Install
`go get github.com/serenity-homelab/sidecar`

## Example

```go
sidecar.Configure("/vault/secrets") // not required

mapOfSecrets, err := sidecar.GetSecrets("secrets.json")

databaseSecrets, err := sidecar.GetDatabaseCreds("database.json")
fmt.Printf("username: %s    password: %s", databaseSecrets.Username, databaseSecrets.Password)
```

## Documentation


### GetSecrets

returns a map of all the secrets

`GetSecrets(path string) (map[string]string, error)`


### GetDatabaseCreds

returns a struct with the username and password

`GetDatabaseCreds(path string) (*DatabaseCreds, error)`