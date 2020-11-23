module main

go 1.15

replace example/go => ./go

require (
	example/go v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/lib/pq v1.8.0
)
