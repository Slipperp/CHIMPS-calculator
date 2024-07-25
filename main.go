package main

import (
	"fmt"
	"math"
	"strconv"

	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func calculateMoney(startingRound int, endRound int) int {
	roundCash := [...]int{20, 35, 35, 71, 59, 57, 75, 92, 90, 204, 78, 80, 169, 145, 151, 152, 48, 240, 141, 66, 230, 176, 154, 43, 210, 207, 535, 138, 260, 207, 406, 495, 72, 778, 1015, 760, 1202, 1139, 1620, 381, 2040, 517, 1135, 1150, 2277, 570, 1490, 2695, 4609, 2866, 947, 1443, 771, 2043, 2328, 1130, 1702, 2140, 2000, 762, 1071, 1224, 2663, 685, 2906, 838, 856, 609, 1222, 2448, 1332, 1332, 1219, 2870, 2492, 1140, 2363, 4684, 6530, 1220, 5185, 4575, 4566, 6860, 2440, 762, 2440, 3126, 1982, 149, 4000, 4345, 1753, 7473, 3523, 9759, 1220, 9455, 2628, 1334}
	money := 0
	startingRound--
	endRound--

	for i := startingRound; i <= endRound; i++ {
		money = money + roundCash[i] + (101 + i)
	}

	return money
}

func calculatePath(top int, mid int, bot int, monkey [3][6]int, topHave int, midHave int, botHave int) int {
	costTop := 0
	costMid := 0
	costBot := 0

	for i := topHave + 1; i <= top; i++ {
		costTop += monkey[0][i]
	}

	for i := midHave + 1; i <= mid; i++ {
		costMid += monkey[1][i]
	}

	for i := botHave + 1; i <= bot; i++ {
		costBot += monkey[2][i]
	}

	cost := costTop + costMid + costBot
	return cost
}

func calculateTower(cost int, startingR int, money int) int {
	roundCash := [...]int{20, 35, 35, 71, 59, 57, 75, 92, 90, 204, 78, 80, 169, 145, 151, 152, 48, 240, 141, 66, 230, 176, 154, 43, 210, 207, 535, 138, 260, 207, 406, 495, 72, 778, 1015, 760, 1202, 1139, 1620, 381, 2040, 517, 1135, 1150, 2277, 570, 1490, 2695, 4609, 2866, 947, 1443, 771, 2043, 2328, 1130, 1702, 2140, 2000, 762, 1071, 1224, 2663, 685, 2906, 838, 856, 609, 1222, 2448, 1332, 1332, 1219, 2870, 2492, 1140, 2363, 4684, 6530, 1220, 5185, 4575, 4566, 6860, 2440, 762, 2440, 3126, 1982, 149, 4000, 4345, 1753, 7473, 3523, 9759, 1220, 9455, 2628, 1334}

	endRound := startingR
	startingR--

	for i := startingR; money < cost; i++ {
		money += roundCash[i] + (101 + i)
		endRound++
	}

	endRound--
	return endRound
}

func calculateXp(difficulty string, round int, heroType float64) [20]int {
	xpNeeded := []float64{180, 460, 1000, 1860, 3280, 5180, 8320, 9380, 13620, 16380, 14400, 16650, 14940, 16380, 17820, 19260, 20700, 16470, 17280}
	xpGained := []float64{40, 60, 80, 100, 120, 140, 160, 180, 200, 220, 240, 260, 280, 300, 320, 340, 360, 380, 400, 420, 460, 500, 540, 580, 620, 660, 700, 740, 780, 820, 860, 900, 940, 980, 1020, 1060, 1100, 1140, 1180, 1220, 1260, 1300, 1340, 1380, 1420, 1460, 1500, 1540, 1580, 1620, 1710, 1800, 1890, 1980, 2070, 2160, 2250, 2340, 2430, 2520, 2610, 2700, 2790, 2880, 2970, 3060, 3150, 3240, 3330, 3420, 3510, 3600, 3690, 3780, 3870, 3960, 4050, 4140, 4230, 4320, 4410, 4500, 4590, 4680, 4770, 4860, 4950, 5040, 5130, 5220, 5310, 5400, 5490, 5580, 5670, 5760, 5850, 5940, 6030, 6120}
	imGoingInsane := [20]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if difficulty == "Intermediate" {
		for i := 0; i < len(xpGained); i++ {
			xpGained[i] = xpGained[i] * 1.1
		}
	} else if difficulty == "Advanced" {
		for i := 0; i < len(xpGained); i++ {
			xpGained[i] = xpGained[i] * 1.2
		}
	} else if difficulty == "Expert" {
		for i := 0; i < len(xpGained); i++ {
			xpGained[i] = xpGained[i] * 1.3
		}
	}

	if heroType == 1.425 {
		for i := 0; i < len(xpNeeded); i++ {
			xpNeeded[i] = xpNeeded[i] * heroType
		}
	} else if heroType == 1.5 {
		for i := 0; i < len(xpNeeded); i++ {
			xpNeeded[i] = xpNeeded[i] * heroType
		}
	} else if heroType == 1.71 {
		for i := 0; i < len(xpNeeded); i++ {
			xpNeeded[i] = xpNeeded[i] * heroType
		}
	}

	imGoingInsane[0] = round
	round--
	level := 0

	for i := round; i < len(xpGained); {
		xpForLevel := xpNeeded[level]
		for xpForLevel >= 0 {
			xpForLevel -= xpGained[i]

			i++
			if i == 100 {
				return imGoingInsane
			}

			if xpForLevel < 0 {
				leftOverXp := math.Abs(xpForLevel)
				xpGained[i] = xpGained[i] + leftOverXp
			}

		}
		level++
		imGoingInsane[level] = i + 1

		if level == 19 {
			break
		}
	}
	return imGoingInsane
}

func main() {
	dartMonkey := [3][6]int{
		{0, 150, 235, 325, 1945, 16200},
		{0, 110, 205, 430, 8100, 48600},
		{0, 95, 215, 675, 2160, 23220},
	}

	boomer := [3][6]int{
		{0, 215, 300, 650, 3240, 31750},
		{0, 190, 270, 1565, 4535, 37800},
		{0, 110, 325, 1405, 2590, 54000},
	}

	bombShooter := [3][6]int{
		{0, 380, 700, 1190, 3455, 59400},
		{0, 270, 430, 1190, 3350, 27540},
		{0, 215, 325, 865, 3025, 30240},
	}

	tackShooter := [3][6]int{
		{0, 160, 325, 650, 3780, 49140},
		{0, 110, 245, 595, 2915, 16200},
		{0, 120, 120, 485, 3455, 21600},
	}

	iceMonkey := [3][6]int{
		{0, 160, 380, 1620, 2375, 30240},
		{0, 245, 485, 3025, 4105, 20735},
		{0, 190, 245, 2430, 2970, 32400},
	}

	glueGunner := [3][6]int{
		{0, 215, 325, 2700, 5400, 23760},
		{0, 110, 1045, 2270, 4160, 17280},
		{0, 300, 430, 3890, 3670, 25920},
	}

	sniperMonkey := [3][6]int{
		{0, 380, 1405, 3240, 6100, 34560},
		{0, 270, 485, 2590, 8210, 15660},
		{0, 485, 485, 3130, 4430, 15875},
	}

	monkeySub := [3][6]int{
		{0, 140, 540, 755, 2485, 33480},
		{0, 485, 325, 1460, 14040, 31320},
		{0, 485, 1080, 1190, 3240, 27000},
	}

	monkeyBuccaneer := [3][6]int{
		{0, 295, 460, 3295, 7450, 26460},
		{0, 595, 540, 970, 5290, 28080},
		{0, 215, 380, 2590, 5940, 24840},
	}

	monkeyAce := [3][6]int{
		{0, 700, 700, 1080, 3240, 45900},
		{0, 215, 380, 970, 19440, 32400},
		{0, 540, 595, 2755, 25270, 91800},
	}

	heliPilot := [3][6]int{
		{0, 865, 540, 1890, 21170, 48600},
		{0, 325, 650, 3780, 10260, 32400},
		{0, 270, 380, 3240, 9180, 37800},
	}

	mortarMonkey := [3][6]int{
		{0, 540, 540, 970, 7020, 38880},
		{0, 325, 540, 970, 6370, 34560},
		{0, 215, 540, 970, 11340, 43200},
	}

	dartlingGunner := [3][6]int{
		{0, 325, 970, 3940, 11990, 86400},
		{0, 270, 1025, 5185, 5995, 64800},
		{0, 160, 1295, 3670, 12960, 62640},
	}

	wizardMonkey := [3][6]int{
		{0, 160, 485, 1510, 10800, 34560},
		{0, 325, 865, 3565, 7560, 54000},
		{0, 325, 325, 1620, 3025, 28620},
	}

	superMonkey := [3][6]int{
		{0, 2160, 2700, 21600, 108000, 540000},
		{0, 1620, 2050, 8100, 27000, 75600},
		{0, 3240, 1295, 6050, 60000, 216000},
	}

	ninjaMonkey := [3][6]int{
		{0, 380, 380, 970, 2970, 37800},
		{0, 270, 430, 1295, 5615, 23760},
		{0, 325, 485, 2430, 5400, 43200},
	}

	alchemist := [3][6]int{
		{0, 270, 380, 1405, 3185, 64800},
		{0, 270, 515, 3240, 4860, 48600},
		{0, 700, 485, 1080, 2970, 43200},
	}

	druid := [3][6]int{
		{0, 380, 920, 1835, 4860, 64800},
		{0, 270, 380, 1135, 5290, 37800},
		{0, 110, 325, 650, 2700, 48600},
	}

	spikeFactory := [3][6]int{
		{0, 865, 650, 2485, 10260, 135000},
		{0, 650, 865, 2700, 5400, 45360},
		{0, 160, 430, 1405, 3890, 32400},
	}

	monkeyVillage := [3][6]int{
		{0, 430, 1620, 865, 2700, 27000},
		{0, 270, 2160, 8100, 21600, 43200},
		{0, 540, 540, 10800, 3240, 5400},
	}

	engineer := [3][6]int{
		{0, 540, 430, 620, 2700, 34560},
		{0, 270, 380, 970, 14580, 77760},
		{0, 485, 235, 485, 3890, 51840},
	}

	startingRound := 1
	endRound := 1
	playerMoney := 0

	a := app.New()
	w := a.NewWindow("BTD6 Calculator")

	label1 := widget.NewLabel("")
	label2 := widget.NewLabel("")
	label3 := widget.NewLabel("")
	label4 := widget.NewLabel("")
	label5 := widget.NewLabel("Upgrades you own")
	label5.Hide()

	quitButton := widget.NewButton("Quit", func() {
		a.Quit()
	})

	originalOptions := []string{"Dart Monkey", "Boomerang Monkey", "Bomb Shooter", "Tack Shooter", "Ice Monkey", "Glue Gunner", "Sniper Monkey", "Monkey Sub", "Monkey Buccaneer", "Monkey Ace", "Heli Pilot", "Mortar Monkey", "Dartling Gunner", "Wizard Monkey", "Super Monkey", "Ninja Monkey", "Alchemist", "Druid", "Spike Factory", "Monkey Village", "Engineer"}
	options2 := originalOptions

	entry := widget.NewEntry()
	selectWidget := widget.NewSelect(options2, func(selected string) {
		entry.SetText(selected)
	})

	selectWidget.Hide()

	entry.OnChanged = func(s string) {
		if s == "" {
			selectWidget.Hide()
			return
		}

		filtered := []string{}
		for _, option2 := range originalOptions {
			if strings.HasPrefix(strings.ToLower(option2), strings.ToLower(s)) {
				filtered = append(filtered, option2)
			}
		}

		if len(filtered) == 0 {
			selectWidget.Hide()
			return
		}

		options2 = filtered
		selectWidget.Options = options2
		selectWidget.Refresh()
		selectWidget.Show()
	}

	sRound := widget.NewEntry()
	sRound.SetPlaceHolder("Enter starting round")

	eRound := widget.NewEntry()
	eRound.SetPlaceHolder("Enter end round")

	pMoney := widget.NewEntry()
	pMoney.SetPlaceHolder("Enter money on hand")

	start := widget.NewEntry()
	start.SetPlaceHolder("Enter your round")

	cash := widget.NewEntry()
	cash.SetPlaceHolder("Enter your cash")

	options := []string{"0", "1", "2", "3", "4", "5"}

	topWanted := 0
	midWanted := 0
	botWanted := 0

	top := 0
	mid := 0
	bot := 0

	selectWidget1 := widget.NewSelect(options, func(selected string) {
		topPath := selected

		if num, err := strconv.Atoi(topPath); err == nil {
			topWanted = num
		}

	})
	selectWidget2 := widget.NewSelect(options, func(selected string) {
		midPath := selected

		if num, err := strconv.Atoi(midPath); err == nil {
			midWanted = num
		}
	})
	selectWidget3 := widget.NewSelect(options, func(selected string) {
		botPath := selected

		if num, err := strconv.Atoi(botPath); err == nil {
			botWanted = num
		}
	})
	selectWidget1.PlaceHolder = "Top path upgrades"
	selectWidget2.PlaceHolder = "Mid path upgrades"
	selectWidget3.PlaceHolder = "Bot path upgrades"

	selectWidget1Owned := widget.NewSelect(options, func(selected string) {
		topPathOwned := selected

		if num, err := strconv.Atoi(topPathOwned); err == nil {
			top = num
		}

	})
	selectWidget1Owned.Hide()

	selectWidget2Owned := widget.NewSelect(options, func(selected string) {
		midPathOwned := selected

		if num, err := strconv.Atoi(midPathOwned); err == nil {
			mid = num
		}
	})
	selectWidget2Owned.Hide()

	selectWidget3Owned := widget.NewSelect(options, func(selected string) {
		botPathOwned := selected

		if num, err := strconv.Atoi(botPathOwned); err == nil {
			bot = num
		}
	})
	selectWidget3Owned.Hide()

	selectWidget1Owned.PlaceHolder = "Top path upgrades"
	selectWidget2Owned.PlaceHolder = "Mid path upgrades"
	selectWidget3Owned.PlaceHolder = "Bot path upgrades"

	check := widget.NewCheck("I have the tower placed", func(checked bool) {
		if checked {
			selectWidget1Owned.Show()
			selectWidget2Owned.Show()
			selectWidget3Owned.Show()
			label5.Show()
		} else {
			selectWidget1Owned.Hide()
			selectWidget2Owned.Hide()
			selectWidget3Owned.Hide()
			label5.Hide()
		}

		w.Content().Refresh()
	})

	calculatePathsButton := widget.NewButton("Calculate Path Upgrades", func() {
		currRoundStr := start.Text
		currRoundInt := 1

		if num, err := strconv.Atoi(currRoundStr); err == nil {
			currRoundInt = num
		}

		cashStr := cash.Text
		cashInt := 0

		if num, err := strconv.Atoi(cashStr); err == nil {
			cashInt = num
		}

		tower := entry.Text
		cost := 0

		if currRoundInt < 1 {
			message3 := "Invalid round"
			label3.SetText(message3)
			return
		} else if cashInt < 0 {
			message3 := "Invalid cash amount"
			label3.SetText(message3)
			return
		}

		if currRoundInt > 100 {
			message3 := "Invalid round"
			label3.SetText(message3)
			return
		}

		if topWanted+midWanted+botWanted > 7 {
			message3 := "Invalid tower upgrade path"
			label3.SetText(message3)
			return
		}

		if top+mid+bot > 7 {
			message3 := "Invalid tower upgrade path"
			label3.SetText(message3)
			return
		}

		if topWanted > 0 && midWanted > 0 && botWanted > 0 {
			message3 := "Invalid tower upgrade path"
			label3.SetText(message3)
			return
		}

		if top > 0 && mid > 0 && bot > 0 {
			message3 := "Invalid tower upgrade path"
			label3.SetText(message3)
			return
		}

		switch tower {
		case "Dart Monkey":
			cost = calculatePath(topWanted, midWanted, botWanted, dartMonkey, top, mid, bot)
			if !check.Checked {
				cost += 215
			}
		case "Boomerang Monkey":
			cost = calculatePath(topWanted, midWanted, botWanted, boomer, top, mid, bot)
			if !check.Checked {
				cost += 340
			}
		case "Bomb Shooter":
			cost = calculatePath(topWanted, midWanted, botWanted, bombShooter, top, mid, bot)
			if !check.Checked {
				cost += 565
			}
		case "Tack Shooter":
			cost = calculatePath(topWanted, midWanted, botWanted, tackShooter, top, mid, bot)
			if !check.Checked {
				cost += 280
			}
		case "Ice Monkey":
			cost = calculatePath(topWanted, midWanted, botWanted, iceMonkey, top, mid, bot)
			if !check.Checked {
				cost += 540
			}
		case "Glue Gunner":
			cost = calculatePath(topWanted, midWanted, botWanted, glueGunner, top, mid, bot)
			if !check.Checked {
				cost += 245
			}
		case "Sniper Monkey":
			cost = calculatePath(topWanted, midWanted, botWanted, sniperMonkey, top, mid, bot)
			if !check.Checked {
				cost += 380
			}
		case "Monkey Sub":
			cost = calculatePath(topWanted, midWanted, botWanted, monkeySub, top, mid, bot)
			if !check.Checked {
				cost += 350
			}
		case "Monkey Buccaneer":
			cost = calculatePath(topWanted, midWanted, botWanted, monkeyBuccaneer, top, mid, bot)
			if !check.Checked {
				cost += 430
			}
		case "Monkey Ace":
			cost = calculatePath(topWanted, midWanted, botWanted, monkeyAce, top, mid, bot)
			if !check.Checked {
				cost += 865
			}
		case "Heli Pilot":
			cost = calculatePath(topWanted, midWanted, botWanted, heliPilot, top, mid, bot)
			if !check.Checked {
				cost += 1730
			}
		case "Mortar Monkey":
			cost = calculatePath(topWanted, midWanted, botWanted, mortarMonkey, top, mid, bot)
			if !check.Checked {
				cost += 810
			}
		case "Dartling Gunner":
			cost = calculatePath(topWanted, midWanted, botWanted, dartlingGunner, top, mid, bot)
			if !check.Checked {
				cost += 920
			}
		case "Wizard Monkey":
			cost = calculatePath(topWanted, midWanted, botWanted, wizardMonkey, top, mid, bot)
			if !check.Checked {
				cost += 350
			}
		case "Super Monkey":
			cost = calculatePath(topWanted, midWanted, botWanted, superMonkey, top, mid, bot)
			if !check.Checked {
				cost += 2700
			}
		case "Ninja Monkey":
			cost = calculatePath(topWanted, midWanted, botWanted, ninjaMonkey, top, mid, bot)
			if !check.Checked {
				cost += 430
			}
		case "Alchemist":
			cost = calculatePath(topWanted, midWanted, botWanted, alchemist, top, mid, bot)
			if !check.Checked {
				cost += 595
			}
		case "Druid":
			cost = calculatePath(topWanted, midWanted, botWanted, druid, top, mid, bot)
			if !check.Checked {
				cost += 430
			}
		case "Spike Factory":
			cost = calculatePath(topWanted, midWanted, botWanted, spikeFactory, top, mid, bot)
			if !check.Checked {
				cost += 1080
			}
		case "Monkey Village":
			cost = calculatePath(topWanted, midWanted, botWanted, monkeyVillage, top, mid, bot)
			if !check.Checked {
				cost += 1295
			}
		case "Engineer":
			cost = calculatePath(topWanted, midWanted, botWanted, engineer, top, mid, bot)
			if !check.Checked {
				cost += 380
			}
		}

		fuckThisCalc := calculateTower(cost, currRoundInt, cashInt)

		message3 := fmt.Sprintf("Total cost: %d. You need %d more money", cost, cost-cashInt)
		label3.SetText(message3)

		message4 := fmt.Sprintf("You can buy this upgrade after round: %d", fuckThisCalc)
		label4.SetText(message4)
	})

	calculateButton := widget.NewButton("Calculate Earnings", func() {
		input1 := sRound.Text

		if num, err := strconv.Atoi(input1); err == nil {
			startingRound = num
		}

		input2 := eRound.Text

		if num, err := strconv.Atoi(input2); err == nil {
			endRound = num
		}

		input3 := pMoney.Text

		if num, err := strconv.Atoi(input3); err == nil {
			playerMoney = num
		}

		if startingRound > 100 || startingRound < 0 {
			message1 := "Invalid starting round"
			label1.SetText(message1)

		} else if endRound > 100 || endRound < 0 {
			message1 := "Invalid end round"
			label1.SetText(message1)

		} else if endRound < startingRound {
			message1 := "End round must be greater than starting round"
			label1.SetText(message1)

		} else {
			x := calculateMoney(startingRound, endRound)
			message1 := fmt.Sprintf("Money earned from round %d to %d: %d", startingRound, endRound, x)

			message2 := fmt.Sprintf("Total money at the end: %d", playerMoney+x)

			label1.SetText(message1)

			if playerMoney > 0 {
				label2.SetText(message2)
			}
		}

	})

	text := widget.NewLabel("Money from rounds calculator")
	someText := widget.NewLabel("Tower cost calculator")
	entry.SetPlaceHolder("Enter tower name")

	originalOptions3 := []string{"Quincy", "Gwendolin", "Striker Jones", "Obyn", "Etienne", "Geraldo", "Ezili", "Pat Fusty", "Admiral Brickell", "Sauda", "Benjamin", "Psi", "Captain Churchill", "Adora", "Corvus"}
	options3 := originalOptions3

	entry2 := widget.NewEntry()
	selectWidget15 := widget.NewSelect(options3, func(selected string) {
		entry2.SetText(selected)
	})

	selectWidget15.Hide()

	entry2.OnChanged = func(s string) {
		if s == "" {
			selectWidget15.Hide()
			return
		}

		filtered := []string{}
		for _, option := range originalOptions3 {
			if strings.HasPrefix(strings.ToLower(option), strings.ToLower(s)) {
				filtered = append(filtered, option)
			}
		}

		if len(filtered) == 0 {
			selectWidget15.Hide()
			return
		}

		options3 = filtered
		selectWidget15.Options = options3
		selectWidget15.Refresh()
		selectWidget15.Show()
	}

	entry2.SetPlaceHolder("Enter hero name")

	heroRound := widget.NewEntry()
	heroRound.SetPlaceHolder("Enter round")

	difficulty := "Begginer"
	Difficulties := []string{"Begginer", "Intermediate", "Advanced", "Expert"}
	mapDifficulty := widget.NewSelect(Difficulties, func(s string) {
		difficulty = s
	})

	level1 := widget.NewLabel("")
	level2 := widget.NewLabel("")
	level3 := widget.NewLabel("")
	level4 := widget.NewLabel("")
	level5 := widget.NewLabel("")
	level6 := widget.NewLabel("")
	level7 := widget.NewLabel("")
	level8 := widget.NewLabel("")
	level9 := widget.NewLabel("")
	level10 := widget.NewLabel("")
	level11 := widget.NewLabel("")
	level12 := widget.NewLabel("")
	level13 := widget.NewLabel("")
	level14 := widget.NewLabel("")
	level15 := widget.NewLabel("")
	level16 := widget.NewLabel("")
	level17 := widget.NewLabel("")
	level18 := widget.NewLabel("")
	level19 := widget.NewLabel("")
	level20 := widget.NewLabel("")

	xpButton := widget.NewButton("Calculate hero leveling", func() {
		heroPlacement := 1

		if num, err := strconv.Atoi(heroRound.Text); err == nil {
			heroPlacement = num
		}

		if heroPlacement > 100 {
			level1.SetText("Invalid round")
			return
		}

		heroMultiplayer := 1.0

		if entry2.Text == "Quincy" || entry2.Text == "Gwendolin" || entry2.Text == "Striker Jones" || entry2.Text == "Obyn" || entry2.Text == "Etienne" || entry2.Text == "Geraldo" {
			heroMultiplayer = 1
		} else if entry2.Text == "Ezili" || entry2.Text == "Pat Fusty" || entry2.Text == "Admiral Brickell" || entry2.Text == "Sauda" || entry2.Text == "Rosalia" {
			heroMultiplayer = 1.425
		} else if entry2.Text == "Benjamin" || entry2.Text == "Psi" {
			heroMultiplayer = 1.5
		} else if entry2.Text == "Corvus" || entry2.Text == "Adora" || entry2.Text == "Captain Churchill" {
			heroMultiplayer = 1.71
		} else {
			level1.SetText("Invalid hero name")
			return
		}

		imGoingInsane := calculateXp(difficulty, heroPlacement, heroMultiplayer)

		messages := make([]string, 0)

		for i := 0; i < len(imGoingInsane); i++ {
			if imGoingInsane[i] > 0 {
				message := fmt.Sprintf("Level %d: round %d", i+1, imGoingInsane[i])
				messages = append(messages, message)
			}
		}

		level1.SetText(messages[0])
		if len(messages) >= 2 {
			level2.SetText(messages[1])
		}
		if len(messages) >= 3 {
			level3.SetText(messages[2])
		}
		if len(messages) >= 4 {
			level4.SetText(messages[3])
		}
		if len(messages) >= 5 {
			level5.SetText(messages[4])
		}
		if len(messages) >= 6 {
			level6.SetText(messages[5])
		}
		if len(messages) >= 7 {
			level7.SetText(messages[6])
		}
		if len(messages) >= 8 {
			level8.SetText(messages[7])
		}
		if len(messages) >= 9 {
			level9.SetText(messages[8])
		}
		if len(messages) >= 10 {
			level10.SetText(messages[9])
		}
		if len(messages) >= 11 {
			level11.SetText(messages[10])
		}
		if len(messages) >= 12 {
			level12.SetText(messages[11])
		}
		if len(messages) >= 13 {
			level13.SetText(messages[12])
		}
		if len(messages) >= 14 {
			level14.SetText(messages[13])
		}
		if len(messages) >= 15 {
			level15.SetText(messages[14])
		}
		if len(messages) >= 16 {
			level16.SetText(messages[15])
		}
		if len(messages) >= 17 {
			level17.SetText(messages[16])
		}
		if len(messages) >= 18 {
			level18.SetText(messages[17])
		}
		if len(messages) >= 19 {
			level19.SetText(messages[18])
		}
		if len(messages) >= 20 {
			level20.SetText(messages[19])
		}
	})

	content := container.NewVBox(
		container.NewCenter(text),
		container.NewGridWithColumns(2,
			widget.NewLabel("Starting Round:"), sRound,
			widget.NewLabel("End Round:"), eRound,
			widget.NewLabel("Money on Hand:"), pMoney,
		),
		container.NewCenter(calculateButton),
		label1,
		label2,
	)

	setContent := func(elements ...fyne.CanvasObject) {
		content.Objects = elements
		content.Refresh()
	}

	btn1 := widget.NewButton("Round cash", func() {
		setContent(
			container.NewCenter(text),
			container.NewGridWithColumns(2,
				widget.NewLabel("Starting Round:"), sRound,
				widget.NewLabel("End Round:"), eRound,
				widget.NewLabel("Money on Hand:"), pMoney,
			),
			container.NewCenter(calculateButton),
			label1,
			label2,
		)
	})

	uselessButton := widget.NewButton("Calculate Paths Button", func() {

	})
	uselessButton.Hide()

	btn2 := widget.NewButton("Tower cost", func() {
		setContent(
			container.NewVBox(
				container.NewCenter(someText),
				container.NewVBox(entry, selectWidget),
				check,
				container.NewGridWithColumns(3,
					container.NewVBox(
						widget.NewLabel("Tower to buy:"),
						selectWidget1,
						selectWidget2,
						selectWidget3,
					),
					container.NewCenter(uselessButton),
					container.NewVBox(
						label5,
						selectWidget1Owned,
						selectWidget2Owned,
						selectWidget3Owned,
					),
				),
				start,
				cash,
				container.NewCenter(calculatePathsButton),
			),
			label3,
			label4,
		)
	})

	heroLevelText := widget.NewLabel("Hero leveling calculator")

	btn3 := widget.NewButton("Hero level", func() {
		setContent(
			container.NewCenter(heroLevelText),
			container.NewVBox(entry2, selectWidget15),
			container.NewGridWithColumns(2,
				widget.NewLabel("When hero was placed: "), heroRound,
				widget.NewLabel("Map difficulty: "), mapDifficulty,
			),
			container.NewCenter(xpButton),

			container.NewCenter(
				container.NewHBox(
					container.NewVBox(
						level1,
						level2,
						level3,
						level4,
						level5,
					),
					container.NewVBox(
						level6,
						level7,
						level8,
						level9,
						level10,
					),
					container.NewVBox(
						level11,
						level12,
						level13,
						level14,
						level15,
					),
					container.NewVBox(
						level16,
						level17,
						level18,
						level19,
						level20,
					),
				),
			),
		)
	})

	buttonContainer := container.New(layout.NewGridLayoutWithColumns(3), btn1, btn2, btn3)

	bottomContainer := container.NewVBox(
		widget.NewSeparator(),
		container.NewCenter(quitButton),
	)

	mainContainer := container.NewBorder(buttonContainer, bottomContainer, nil, nil, content)

	w.Resize(fyne.NewSize(590, 560))
	w.SetContent(mainContainer)
	w.ShowAndRun()
}
