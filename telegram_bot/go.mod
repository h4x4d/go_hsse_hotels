module telegram_bot

go 1.23.1

replace github.com/h4x4d/go_hsse_hotels/pkg => ../pkg

require (
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/h4x4d/go_hsse_hotels/pkg v0.0.0-00010101000000-000000000000
	github.com/jackc/pgx/v5 v5.7.1
	github.com/joho/godotenv v1.5.1
)

require (
	github.com/Nerzal/gocloak/v13 v13.9.0 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/text v0.18.0 // indirect
)
