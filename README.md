# Myosin

## Overview

This project is a comprehensive tool designed to parse and analyze the equipment settings of professional players across multiple popular eSports games including Valorant, Rainbow Six Siege (RS6), Fortnite, and Counter-Strike: Global Offensive (CS:GO). It provides functionalities to retrieve specific player settings from CSV data files, search for players by name, and display commonly used equipment based on the parsed data.

## Features

### Player Settings Parsing

- **Game-specific Parsing**: Separate functions to parse player settings for each game from respective CSV files, accommodating the unique structure and attributes specific to each game.
- **Supported Games**: Valorant, RS6, Fortnite, and CS:GO.

### Search Functionality

- **Case-Insensitive Search**: Users can search for players by name across any of the supported games. The search is case-insensitive, enhancing user experience by providing flexibility in query input.
- **Similar Name Suggestions**: If the exact player name isn't found, the program suggests similar names based on partial matches, guiding users towards potential correct entries.

### Equipment Usage Analysis

- **Top Equipment Items**: Analyzes and identifies the most commonly used mousepads, keyboards, and headsets among the players, providing insights into popular gear choices in the professional scene.
- **Percentage Calculations**: Calculates what percentage of players use a particular item, giving a clear picture of equipment trends and preferences.

## Understanding the Project Name: "Myosin"

The term "myosin" is integral to the core functionality of this project, symbolizing both strength and precision in performance. Derived from the Greek word "μυς" (mus), meaning muscle, myosin is a type of motor protein that plays a crucial role in the contraction of muscle fibers and a variety of cellular movements. In the context of our project, "myosin" represents the robust and efficient processing of data — akin to the way myosin proteins facilitate muscle contractions by converting chemical energy into mechanical work. This analogy highlights the project's capability to handle and analyze complex eSports player data effectively, delivering insights with precision and speed.

## Usage

### Prerequisites

- Go programming environment
- CSV files with player settings data placed under a `data` directory

### Running the Program

Execute the program by running the `main` function. This will:
1. Parse player settings from CSV files.
2. Print out the top used equipment items.
3. Search for specific players in each game and provide output or suggestions.

### Example

```bash
go run main.go
```

Output includes:
- Top equipment items by usage percentage.
- Search results for specific players or suggestions for similar player names.

### Example Response
```golang
Top mousepads:
SteelSeries QcK Heavy: 10.78%
Logitech G640 Original: 6.29%
Razer Gigantus V2: 5.72%

Top keyboards:
Razer Huntsman V3 Pro TKL: 6.94%
Logitech G Pro X Keyboard: 6.45%
Wooting 60 HE: 12.66%

Top headsets:
HyperX Cloud II: 21.24%
Logitech G Pro X Headset: 12.34%
Razer BlackShark V2 Pro Black: 5.47%

CS:GO Player: [{FaZe Clan Frozen Rifler Razer Deathadder V3 Pro Black 4000 400 2 800 0.8 ZOWIE XL2566K 1920x1080 16:9 Native Razer Gigantus V2 Razer Huntsman V3 Pro TKL Razer BlackShark V2 Pro Black}]
Valorant Player: [{Sentinels TenZ Ninjutso Sora V2 1000 1600 0.24 384 1 ZOWIE XL2586X 1920x1080 Custom Mousepad Wooting 60 HE Xtrfy H1}]
RS6 Player: [{Content Creator Jynxzi     100 35  23 23 38 41 49 78 90 Acer XF240Q 1920x1080 16:9  }]
Fortnite Player: [{FaZe Clan Cizzorz FinalMouse Air58 Ninja CBR 400 8.0% 8.0% 100.0% 100.0% Alienware AW2518H 1920x1080 SteelSeries QcK+ Razer Huntsman Elite HyperX Cloud Alpha}]
```
