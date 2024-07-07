module github.com/lcphutchinson/pokedexcli

go 1.22.4

replace (
	github.com/lcphutchinson/pokecache v0.0.0 => ./internal/pokecache
	github.com/lcphutchinson/pokecaller v0.0.0 => ./internal/pokecaller
	github.com/lcphutchinson/pokejson v0.0.0 => ./internal/pokejson
	github.com/lcphutchinson/pokedex v0.0.0 => ./internal/pokedex
)

require (
	github.com/lcphutchinson/pokecache v0.0.0
	github.com/lcphutchinson/pokecaller v0.0.0
	github.com/lcphutchinson/pokejson v0.0.0
	github.com/lcphutchinson/pokedex v0.0.0
)
