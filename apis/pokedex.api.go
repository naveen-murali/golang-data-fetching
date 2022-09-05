package apis

const POKEDEX = "http://pokeapi.co/api/v2/pokedex/kanto/"

/* Subs Types--------------------------------------------- */

type shared struct {
	Name string
	Url  string
}
type Language shared
type Pokemon_species Language

/* Main Types--------------------------------------------- */

type Description struct {
	Description string
	Language    Language
}

type Name struct {
	Name     string
	Language Language
}

type Pokemon_entries struct {
	Entry_number    uint8
	Pokemon_species Pokemon_species
}

type Region Language
type Version_groups Language

type PokedexData struct {
	Id              uint8
	Descriptions    []Description
	Is_main_series  bool
	Name            string
	Names           []Name
	Pokemon_entries []Pokemon_entries
	Region          Region
	Version_groups  []Version_groups
}
