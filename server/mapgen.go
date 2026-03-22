package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/studio-imperium/atlas"
)

// Tiles
var WATER int8 = 0
var GRASS int8 = 1
var WOOD int8 = 2
var FLOOR int8 = 3
var DRYGRASS int8 = 4
var SAND int8 = 5
var SANDSTONE int8 = 6
var COLDGRASS int8 = 7
var SNOW int8 = 8
var ICE int8 = 9
var GRAVEL int8 = 10
var RUIN int8 = 11
var LAVA int8 = 12

var Snowy []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewFill(ICE),
	),
	atlas.NewBiome(
		atlas.NewCropCircle(3, ICE, SNOW),
		atlas.NewSelectiveBorder(SNOW, ICE),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(10, ICE, COLDGRASS),
		atlas.NewSelectiveBorder(SNOW, ICE),
		atlas.NewSelectiveBorder(SNOW, ICE),
	),
}
var Snowy2 []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewFill(COLDGRASS),
	),
	atlas.NewBiome(
		atlas.NewPattern(10, ICE, COLDGRASS),
		atlas.NewSelectiveBorder(SNOW, ICE),
	),
	atlas.NewBiome(
		atlas.NewPattern(4.1, ICE, COLDGRASS),
		atlas.NewSelectiveBorder(SNOW, ICE),
		atlas.NewSelectiveBorder(SNOW, COLDGRASS),
	),
	atlas.NewBiome(
		atlas.NewPattern(4.1, ICE, COLDGRASS),
		atlas.NewSelectiveBorder(SNOW, COLDGRASS),
		atlas.NewSelectiveBorder(SNOW, COLDGRASS),
	),
	atlas.NewBiome(
		atlas.NewFill(ICE),
	),
}
var Sandy []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewCropCircle(7, SANDSTONE, DRYGRASS),
		atlas.NewSelectiveBorder(SAND, DRYGRASS),
		atlas.NewSelectiveBorder(SAND, SANDSTONE),
		atlas.NewSelectiveBorder(SANDSTONE, DRYGRASS),
		atlas.NewSelectiveBorder(SANDSTONE, DRYGRASS),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(20, SANDSTONE, DRYGRASS),
		atlas.NewSelectiveBorder(SAND, SANDSTONE),
		atlas.NewSelectiveBorder(SAND, SANDSTONE),
	),
}
var Sandy2 []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewCropCircle(3.5, DRYGRASS, SAND),
		atlas.NewSelectiveBorder(SANDSTONE, SAND),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(20, SANDSTONE, DRYGRASS),
		atlas.NewSelectiveBorder(SAND, SANDSTONE),
		atlas.NewSelectiveBorder(SAND, SANDSTONE),
	),
}
var Sandy3 []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewCropCircle(3.5, DRYGRASS, SAND),
		atlas.NewSelectiveBorder(SANDSTONE, SAND),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(20, SANDSTONE, DRYGRASS),
		atlas.NewSelectiveBorder(SAND, SANDSTONE),
		atlas.NewSelectiveBorder(SAND, SANDSTONE),
	),
	atlas.NewBiome(
		atlas.NewCropCircle(7, SANDSTONE, DRYGRASS),
		atlas.NewSelectiveBorder(SAND, DRYGRASS),
		atlas.NewSelectiveBorder(SAND, SANDSTONE),
		atlas.NewSelectiveBorder(SANDSTONE, DRYGRASS),
		atlas.NewSelectiveBorder(SANDSTONE, DRYGRASS),
	),
}
var Hot []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewPattern(3.3, LAVA, GRAVEL),
	),
	atlas.NewBiome(
		atlas.NewPattern(15, LAVA, RUIN),
		atlas.NewSelectiveBorder(GRAVEL, LAVA),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(20, RUIN, LAVA),
		atlas.NewSelectiveBorder(GRAVEL, LAVA),
	),
}
var Beach []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewFill(GRASS),
	),
	atlas.NewBiome(
		atlas.NewFill(GRASS),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(6, GRASS, WATER, SAND),
		atlas.NewSelectiveBorder(SAND, GRASS),
		atlas.NewBorder(SAND),
		atlas.NewSelectiveExternalBorder(SAND, SAND),
	),
	atlas.NewBiome(
		atlas.NewPattern(6.1, GRASS, WATER),
		atlas.NewSelectiveBorder(SAND, GRASS),
		atlas.NewSelectiveBorder(SAND, WATER),
	),
	atlas.NewBiome(
		atlas.NewPattern(6.1, GRASS, WATER),
		atlas.NewSelectiveBorder(SAND, GRASS),
		atlas.NewSelectiveBorder(WATER, GRASS),
		atlas.NewSelectiveBorder(WATER, GRASS),
		atlas.NewSelectiveBorder(WATER, GRASS),
		atlas.NewSelectiveBorder(WATER, GRASS),
	),
	atlas.NewBiome(
		atlas.NewFill(WATER),
	),
}
var Glaciers []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewVoronoi(40, FLOOR, FLOOR, FLOOR, FLOOR, SNOW),
		atlas.NewSelectiveBorder(FLOOR, SNOW),
		atlas.NewBorder(FLOOR),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(40, FLOOR, FLOOR, FLOOR, FLOOR, FLOOR, FLOOR, FLOOR, SNOW),
		atlas.NewSelectiveBorder(FLOOR, SNOW),
		atlas.NewBorder(SNOW),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(40, FLOOR, FLOOR, FLOOR, SNOW, ICE),
		atlas.NewSelectiveBorder(FLOOR, SNOW),
		atlas.NewSelectiveBorder(SNOW, ICE),
		atlas.NewBorder(SNOW),
	),
	atlas.NewBiome(
		atlas.NewPattern(27, FLOOR, SNOW),
		atlas.NewSelectiveBorder(SNOW, ICE),
		atlas.NewSelectiveBorder(FLOOR, FLOOR),
		atlas.NewBorder(ICE),
		atlas.NewSelectiveExternalBorder(SNOW, ICE),
	),
	atlas.NewBiome(
		atlas.NewPattern(27, ICE, SNOW),
		atlas.NewSelectiveBorder(SNOW, ICE),
	),
	atlas.NewBiome(
		atlas.NewFill(SNOW),
	),
}

func newHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		biome := rand.Intn(4)
		seed := rand.Int63n(200)
		scale := 75 + rand.Intn(75)
		density := 50 + rand.Intn(75)
		world := atlas.NewWorld(scale, density, seed)

		fmt.Println(biome)

		if biome == 0 {
			if rand.Intn(2) == 0 {
				world.Infect(Snowy, 0.6)
			} else if rand.Intn(2) == 0 {
				world = atlas.NewWorld(scale, density, seed)
				world.Infect(Snowy2, 0.5)
			} else {
				world.Infect(Glaciers, 0.5)
			}
		}
		if biome == 1 {
			world.Infect(Hot, 0.4)
		}
		if biome == 2 {
			density += 100
			scale -= 20
			world = atlas.NewWorld(scale, density, seed)
			world.Infect(Beach, 0.6)
		}
		if biome == 3 {
			if rand.Intn(2) == 0 {
				world.Infect(Sandy, 0.2)
			} else if rand.Intn(2) == 0 {
				world.Infect(Sandy2, 0.2)
			} else {
				world.Infect(Sandy3, 0.4)
			}
		}

		if err := json.NewEncoder(w).Encode(world); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	fmt.Println("Listening on 7778")
	http.HandleFunc("/get_map", newHandler())
	log.Fatal(http.ListenAndServe(":7778", nil))
}
