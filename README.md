# hospital-information-service-backend

## step

```bash
    source ~/.zshrc
    which go
    go env GOROOT
    go env GOPATH

    go mod init github.com/naveeharn/hospital-information-service-backend

    go get -u github.com/gin-gonic/gin
    go get -u github.com/gin-contrib/cors
    go get -u github.com/go-playground/assert/v2
    go get -u github.com/joho/godotenv
    go get -u github.com/mashingan/smapping
    go get -u golang.org/x/crypto
    go get -u gorm.io/driver/postgres
    go get -u gorm.io/gorm

    go get -u github.com/jackc/pgx/v5/stdlib
```

```bash
    go install github.com/air-verse/air@latest
    go install github.com/gin-contrib/cors@latest
    go install golang.org/x/crypto@latest
    air init
    air

    go run cmd/server/main.go
    


```