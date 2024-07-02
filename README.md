# NFT-Bidding-Platform

<h2>Import Packages</h2>

```bash
# For interacting with Google Cloud Storage services
go get cloud.google.com/go/storage@v1.40.0

# A robust library for sending emails
go get github.com/go-mail/mail@v2.3.1

# For validating structs and fields against certain criteria
go get github.com/go-playground/validator/v10@v10.19.0

# Provides methods for creating and validating JSON Web Tokens (JWT)
go get github.com/golang-jwt/jwt/v5@v5.2.1

# Used for generating, parsing, and inspecting UUIDs
go get github.com/google/uuid@v1.6.0

# Loads environment variables from a `.env` file
go get github.com/joho/godotenv@v1.5.1

# A PDF document generator
go get github.com/jung-kurt/gofpdf@v1.16.2

# A high performance, extensible, minimalist Go web framework
go get github.com/labstack/echo/v4@v4.11.4

# The official Go SDK for interacting with Razorpay payment gateway
go get github.com/razorpay/razorpay-go@v1.3.2

# A cron library for scheduling jobs to run at fixed times or intervals
go get github.com/robfig/cron/v3@v3.0.1

# The official MongoDB driver for the Go language
go get go.mongodb.org/mongo-driver@v1.14.0

# Provides additional cryptography packages
go get golang.org/x/crypto@v0.21.0

# The Go implementation of gRPC, a high performance, open-source universal RPC framework
go get google.golang.org/grpc@v1.62.1
```

<h2>Start App in Terminal</h2>

```bash
go run main.go ./env/dev/.env.auth
```
```bash
go run main.go ./env/dev/.env.nft
```
```bash
go run main.go ./env/dev/.env.user
```


<h2>Generate a Proto File Command</h2>

<p>Auth</p>

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    modules/auth/authPb/authPb.proto
```

<p>NFT</p>

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    modules/nft/nftPb/nftPb.proto
```

<p>User</p>

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    modules/user/userPb/userPb.proto
```
