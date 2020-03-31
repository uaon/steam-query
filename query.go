package main

import (
	"flag"
	"fmt"
	"sort"

	steam "github.com/kidoman/go-steam"
)

// address list should change it to a var and GET for REST
var addresses = []string{
	"142.44.143.136:2323",
	"144.217.11.41:2313",
	"144.217.11.41:2323",
	"144.217.11.41:2333",
	"144.217.11.41:2343",
	"144.217.11.41:2303",
}

type SteamPlayers steam.PlayersInfoResponse

func (d SteamPlayers) Len() int {
	return len(d.Players)
}
func (d SteamPlayers) Swap(i, j int) {
	d.Players[i], d.Players[j] = d.Players[j], d.Players[i]
}
func (d SteamPlayers) Less(i, j int) bool {
	return d.Players[i].Score < d.Players[j].Score
}

func main() {
	flag.Parse()
	for _, addr := range addresses {
		server, err := steam.Connect(addr)
		if err != nil {
			panic(err)
		}
		defer server.Close()
		ping, err := server.Ping()
		if err != nil {
			fmt.Printf("warcrimes were committed %v: %v", addr, err)
			continue
		}
		info, err := server.Info()
		if err != nil {
			fmt.Printf("info err %v: %v", addr, err)
			continue
		}
		playersInfo, err := server.PlayersInfo()
		if err != nil {
			fmt.Printf("player info err %v: %v", addr, err)
			continue
		}
		fmt.Printf(" Server | %v\n", info.Name)
		fmt.Printf("   Name | %s\n", info.Game)
		fmt.Printf("   Ping | %s\n", ping)
		fmt.Printf("   Type | %s\n", info.ServerType)
		fmt.Printf("Version | %s\n", info.Version)
		fmt.Printf("Players | %d/%d\n", info.Players, info.MaxPlayers)
		fmt.Printf("    Map | %s\n", info.Map)
		fmt.Printf("     IP | %s\n", addr)

		if len(playersInfo.Players) > 0 {
			fmt.Printf("--------|-----------------------------------------\n")
			fmt.Printf("  Score | Time On | Player Name\n")
			fmt.Printf("--------|-----------------------------------------\n")
			var temp SteamPlayers
			temp.Players = playersInfo.Players
			sort.Sort(sort.Reverse(temp))
			for _, player := range temp.Players {
				fmt.Printf("%7d | %7d | %s\n", player.Score, int(player.Duration/60), player.Name)
			}
			fmt.Printf("\n")
		}
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
