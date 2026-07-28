README.md file
insteall postgress & write the ~/.gatorconfig file & insure postgress db entry is clean;
use goose to upmigrate to latest db scema
postgres://postgres:postgres@localhost:5432/gator?sslmode=disable
goose postgres postgres://postgres:postgres@localhost:5432/gator?sslmode=disable up
sqlc generate # generate internal go sql code 

