package constants

import (
  "time"
)
const(
  DBUSER      = "root"
  DBPASSWORD  = "pssnoida123"
  DBNAME      = "MyBlog"
)

const (
    JWTTokensExpiry = 100 * time.Hour // Define the token expiration time
)
var JWTSecret       =  []byte("WDcNvF0NkrBlxIoqCwTlxFYNY5XulNCd") //secret key
