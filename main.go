package main

import (
	api "fetch-data/apis"
	myError "fetch-data/errors"
	fetch "fetch-data/fetch"

	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	defer myError.CapturePanic(true, nil)

	wg.Add(1)
	go jsonfunc()

	wg.Add(1)
	go poke()

	wg.Wait()
}

func jsonfunc() {
	defer myError.CapturePanic(true, func(_ *any) {
		wg.Done()
	})

	/* api.JSON_PLACEHOLDER */
	jsonRes := fetch.Get(api.JSON_PLACEHOLDER)
	jd := new(api.PostsData)

	jsonRes.Body.Json(&jd)
	
	fmt.Printf("%+v\n\n", jd)
	wg.Done()
}

func poke() {
	defer myError.CapturePanic(true, func(_ *any) {
		wg.Done()
	})

	/* api.POKEDEX */
	pokeRes := fetch.Get(api.POKEDEX)
	pd := new(api.PokedexData)

	pokeRes.Body.Json(&pd)

	fmt.Printf("%+v\n\n", pd)
	wg.Done()
}
