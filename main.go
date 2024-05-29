package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

type ValorantPlayerSettings struct {
	Team       string
	Player     string
	Mouse      string
	Hz         string
	DPI        string
	Sens       string
	eDPI       string
	ScopedSens string
	Monitor    string
	Resolution string
	Mousepad   string
	Keyboard   string
	Headset    string
}
type RS6PlayerSettings struct {
	Team        string
	Player      string
	Mouse       string
	Hz          string
	DPI         string
	Multiplier  string
	HorzSens    string
	VertSens    string
	WinSens     string
	ADS         string
	X1x         string
	X25x        string
	X35x        string
	X5x         string
	X12x        string
	FOV         string
	Monitor     string
	Resolution  string
	AspectRatio string
	Mousepad    string
	Keyboard    string
}
type FortnitePlayerSettings struct {
	Team          string
	Player        string
	Mouse         string
	DPI           string
	XAxisSens     string
	YAxisSens     string
	TargetingSens string
	ScopeSens     string
	Monitor       string
	Resolution    string
	Mousepad      string
	Keyboard      string
	Headset       string
}
type CSGOPlayerSettings struct {
	Team        string
	Player      string
	Role        string
	Mouse       string
	Hz          string
	DPI         string
	Sens        string
	eDPI        string
	ZoomSens    string
	Monitor     string
	Resolution  string
	AspectRatio string
	ScalingMode string
	Mousepad    string
	Keyboard    string
	Headset     string
}

func ParseValorantCSV(filename string) ([]ValorantPlayerSettings, error) {
	file, err := os.Open("data/" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var players []ValorantPlayerSettings

	for _, record := range records {
		player := ValorantPlayerSettings{
			Team:       record[0],
			Player:     record[1],
			Mouse:      record[2],
			Hz:         record[3],
			DPI:        record[4],
			Sens:       record[5],
			eDPI:       record[6],
			ScopedSens: record[7],
			Monitor:    record[8],
			Resolution: record[9],
			Mousepad:   record[10],
			Keyboard:   record[11],
			Headset:    record[12],
		}
		players = append(players, player)
	}

	return players, nil
}
func ParseRS6CSV(filename string) ([]RS6PlayerSettings, error) {
	file, err := os.Open("data/" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var players []RS6PlayerSettings

	for _, record := range records {
		player := RS6PlayerSettings{
			Team:        record[0],
			Player:      record[1],
			Mouse:       record[2],
			Hz:          record[3],
			DPI:         record[4],
			Multiplier:  record[5],
			HorzSens:    record[6],
			VertSens:    record[7],
			WinSens:     record[8],
			ADS:         record[9],
			X1x:         record[10],
			X25x:        record[11],
			X35x:        record[12],
			X5x:         record[13],
			X12x:        record[14],
			FOV:         record[15],
			Monitor:     record[16],
			Resolution:  record[17],
			AspectRatio: record[18],
			Mousepad:    record[19],
			Keyboard:    record[20],
		}
		players = append(players, player)
	}

	return players, nil
}
func ParseFortniteCSV(filename string) ([]FortnitePlayerSettings, error) {
	file, err := os.Open("data/" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var players []FortnitePlayerSettings

	for _, record := range records {
		player := FortnitePlayerSettings{
			Team:          record[0],
			Player:        record[1],
			Mouse:         record[2],
			DPI:           record[3],
			XAxisSens:     record[4],
			YAxisSens:     record[5],
			TargetingSens: record[6],
			ScopeSens:     record[7],
			Monitor:       record[8],
			Resolution:    record[9],
			Mousepad:      record[10],
			Keyboard:      record[11],
			Headset:       record[12],
		}
		players = append(players, player)
	}

	return players, nil
}
func ParseCSGOCsv(filename string) ([]CSGOPlayerSettings, error) {
	file, err := os.Open("data/" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var players []CSGOPlayerSettings

	for _, record := range records {
		player := CSGOPlayerSettings{
			Team:        record[0],
			Player:      record[1],
			Role:        record[2],
			Mouse:       record[3],
			Hz:          record[4],
			DPI:         record[5],
			Sens:        record[6],
			eDPI:        record[7],
			ZoomSens:    record[8],
			Monitor:     record[9],
			Resolution:  record[10],
			AspectRatio: record[11],
			ScalingMode: record[12],
			Mousepad:    record[13],
			Keyboard:    record[14],
			Headset:     record[15],
		}
		players = append(players, player)
	}

	return players, nil
}

var (
	ValorantPlayers, _ = ParseValorantCSV("valorant.csv")
	CSGOPlayers, _     = ParseCSGOCsv("csgo.csv")
	RS6Players, _      = ParseRS6CSV("rs6.csv")
	FortnitePlayers, _ = ParseFortniteCSV("fortnite.csv")
)

func SearchCSGOPlayer(playerName string) ([]CSGOPlayerSettings, string) {
	players := CSGOPlayers
	var result []CSGOPlayerSettings
	for _, player := range players {
		if strings.EqualFold(player.Player, playerName) {
			return append(result, player), ""
		}
	}

	var similarPlayers []string
	for _, player := range players {
		if strings.Contains(strings.ToLower(player.Player), strings.ToLower(playerName)) {
			similarPlayers = append(similarPlayers, player.Player)
		}
	}

	if len(similarPlayers) > 0 {
		suggestion := fmt.Sprintf("Did you mean \"%s\"?", similarPlayers[0])
		return nil, suggestion
	}

	return nil, "Player not found"
}
func SearchValorantPlayer(playerName string) ([]ValorantPlayerSettings, string) {
	players := ValorantPlayers
	var result []ValorantPlayerSettings
	for _, player := range players {
		if strings.EqualFold(player.Player, playerName) {
			return append(result, player), ""
		}
	}

	var similarPlayers []string
	for _, player := range players {
		if strings.Contains(strings.ToLower(player.Player), strings.ToLower(playerName)) {
			similarPlayers = append(similarPlayers, player.Player)
		}
	}

	if len(similarPlayers) > 0 {
		suggestion := fmt.Sprintf("Did you mean \"%s\"?", similarPlayers[0])
		return nil, suggestion
	}

	return nil, "Player not found"
}
func SearchRS6Player(playerName string) ([]RS6PlayerSettings, string) {
	players := RS6Players
	var result []RS6PlayerSettings
	for _, player := range players {
		if strings.EqualFold(player.Player, playerName) {
			return append(result, player), ""
		}
	}

	var similarPlayers []string
	for _, player := range players {
		if strings.Contains(strings.ToLower(player.Player), strings.ToLower(playerName)) {
			similarPlayers = append(similarPlayers, player.Player)
		}
	}

	if len(similarPlayers) > 0 {
		suggestion := fmt.Sprintf("Did you mean \"%s\"?", similarPlayers[0])
		return nil, suggestion
	}

	return nil, "Player not found"
}
func SearchFortnitePlayer(playerName string) ([]FortnitePlayerSettings, string) {
	players := FortnitePlayers
	var result []FortnitePlayerSettings
	for _, player := range players {
		if strings.EqualFold(player.Player, playerName) {
			return append(result, player), ""
		}
	}

	var similarPlayers []string
	for _, player := range players {
		if strings.Contains(strings.ToLower(player.Player), strings.ToLower(playerName)) {
			similarPlayers = append(similarPlayers, player.Player)
		}
	}

	if len(similarPlayers) > 0 {
		suggestion := fmt.Sprintf("Did you mean \"%s\"?", similarPlayers[0])
		return nil, suggestion
	}

	return nil, "Player not found"
}

func getTopItems(field string, playersLists ...interface{}) map[string]float64 {
	itemCount := make(map[string]int)
	totalPlayers := 0

	for _, players := range playersLists {
		v := reflect.ValueOf(players)
		if v.Kind() != reflect.Slice {
			continue // Skip if not a slice
		}
		totalPlayers += v.Len()
		for i := 0; i < v.Len(); i++ {
			player := v.Index(i)
			fieldValue := player.FieldByName(field).String()
			if fieldValue != "" && fieldValue != "<invalid Value>" {
				itemCount[fieldValue]++
			}
		}
	}

	topItems := make(map[string]float64)
	for item, count := range itemCount {
		if item != "" {
			percentage := float64(count) / float64(totalPlayers) * 100
			topItems[item] = percentage
		}
	}

	// Sort the top items by usage percentage
	var sortedItems []struct {
		Item       string
		Percentage float64
	}
	for item, percentage := range topItems {
		sortedItems = append(sortedItems, struct {
			Item       string
			Percentage float64
		}{Item: item, Percentage: percentage})
	}
	sort.Slice(sortedItems, func(i, j int) bool {
		return sortedItems[i].Percentage > sortedItems[j].Percentage
	})

	// Take only the top three items
	topThree := make(map[string]float64)
	for _, item := range sortedItems[:3] {
		topThree[item.Item] = item.Percentage
	}

	return topThree
}
func topItems() {
	topMousepads := getTopItems("Mousepad", ValorantPlayers, RS6Players, FortnitePlayers, CSGOPlayers)
	fmt.Println("Top mousepads:")
	for item, percentage := range topMousepads {
		fmt.Printf("%s: %.2f%%\n", item, percentage)
	}
	topKeyboards := getTopItems("Keyboard", ValorantPlayers, RS6Players, FortnitePlayers, CSGOPlayers)
	fmt.Println("\nTop keyboards:")
	for item, percentage := range topKeyboards {
		fmt.Printf("%s: %.2f%%\n", item, percentage)
	}

	topHeadsets := getTopItems("Headset", ValorantPlayers, RS6Players, FortnitePlayers, CSGOPlayers)
	fmt.Println("\nTop headsets:")
	for item, percentage := range topHeadsets {
		fmt.Printf("%s: %.2f%%\n", item, percentage)
	}
}

func main() {
	topItems()
	csgoPlayer, suggestion := SearchCSGOPlayer("Frozen")
	fmt.Println("CS:GO Player:", csgoPlayer, suggestion)

	valorantPlayer, suggestion := SearchValorantPlayer("Tenz")
	fmt.Println("Valorant Player:", valorantPlayer, suggestion)

	rs6Player, suggestion := SearchRS6Player("Jynxzi")
	fmt.Println("RS6 Player:", rs6Player, suggestion)

	fortnitePlayer, suggestion := SearchFortnitePlayer("Cizzorz")
	fmt.Println("Fortnite Player:", fortnitePlayer, suggestion)
}
