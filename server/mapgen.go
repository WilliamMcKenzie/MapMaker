package main

import (
	"encoding/json"
	"fmt"
	"log"
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
		atlas.NewCropCircle(7, ICE, COLDGRASS),
		atlas.NewSelectiveBorder(SNOW, COLDGRASS),
		atlas.NewSelectiveBorder(SNOW, ICE),
		atlas.NewSelectiveBorder(ICE, COLDGRASS),
		atlas.NewSelectiveBorder(ICE, COLDGRASS),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(20, ICE, COLDGRASS),
		atlas.NewSelectiveBorder(SNOW, ICE),
		atlas.NewSelectiveBorder(SNOW, ICE),
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
var Hot []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewPattern(6.5, LAVA, RUIN),
		atlas.NewSelectiveBorder(GRAVEL, LAVA),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(20, RUIN, LAVA),
		atlas.NewSelectiveBorder(GRAVEL, LAVA),
	),
}
var Beach []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewFill(SAND),
	),
	atlas.NewBiome(
		atlas.NewFill(SAND),
		atlas.NewBorder(SANDSTONE),
	),
	atlas.NewBiome(
		atlas.NewFill(SAND),
		atlas.NewBorder(GRASS),
		atlas.NewSelectiveExternalBorder(GRASS, GRASS),
		atlas.NewSelectiveBorder(SANDSTONE, SAND),
		atlas.NewSelectiveBorder(SANDSTONE, SAND),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(40, GRASS, SANDSTONE, SAND),
		atlas.NewSelectiveExternalBorder(GRASS, GRASS),
		atlas.NewSelectiveExternalBorder(SANDSTONE, SAND),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(4, GRASS, WATER),
		atlas.NewBorder(GRASS),
	),
	atlas.NewBiome(
		atlas.NewFill(WATER),
	),
}

// var Sandy []atlas.Biome = []atlas.Biome{
// 	atlas.NewBiome(
// 		atlas.NewCropCircle(3.2, DRYGRASS, SAND),
// 		atlas.NewSelectiveBorder(SANDSTONE, SAND),
// 	),
// 	atlas.NewBiome(
// 		atlas.NewVoronoi(40, SAND, SANDSTONE, DRYGRASS),
// 		atlas.NewSelectiveBorder(SANDSTONE, DRYGRASS),
// 	),
// 	atlas.NewBiome(
// 		atlas.NewCropCircle(40, DRYGRASS, SAND),
// 		atlas.NewSelectiveBorder(SANDSTONE, SAND),
// 		atlas.NewSelectiveBorder(SAND, DRYGRASS),
// 		atlas.NewSelectiveBorder(SAND, DRYGRASS),
// 		atlas.NewSelectiveBorder(SAND, DRYGRASS),
// 		atlas.NewSelectiveBorder(SANDSTONE, SAND),
// 	),
// }

func newHandler(world *atlas.World) func(http.ResponseWriter, *http.Request) {
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

		if err := json.NewEncoder(w).Encode(world); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	world1 := atlas.NewWorld(100, 100, 8)
	world1.Infect(Snowy, 0.5)
	world2 := atlas.NewWorld(100, 100, 3)
	world2.Infect(Hot, 0.2)
	world3 := atlas.NewWorld(100, 100, 8)
	world3.Infect(Beach, 0.5)
	world4 := atlas.NewWorld(100, 100, 21)
	world4.Infect(Sandy, 0.2)

	fmt.Println("Listening on 7778")
	http.HandleFunc("/get_map_1", newHandler(world1))
	http.HandleFunc("/get_map_2", newHandler(world2))
	http.HandleFunc("/get_map_3", newHandler(world3))
	http.HandleFunc("/get_map_4", newHandler(world4))
	log.Fatal(http.ListenAndServe(":7778", nil))
}
