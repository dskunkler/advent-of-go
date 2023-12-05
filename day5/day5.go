package day5

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1() {
	content,err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("woooops")
		os.Exit(1)
	}
	text:=string(content)
	splitRe:=regexp.MustCompile(`\n\s*\n`)
	sections:= splitRe.Split(text,-1)

	// Make maps
	seeds:= make(map[int]int)
	stsSection:= strings.Split(sections[1],":")[1]
	lineRe :=regexp.MustCompile(`(\d+) (\d+) (\d+)`)
	stsLines := lineRe.FindAllString(stsSection,-1)
	sToS :=make(map[int]int)

	sToF :=make(map[int]int)
	stfSection := strings.Split(sections[2],":")[1]
	stfLines:= lineRe.FindAllString(stfSection,-1)
	fToW :=make(map[int]int)
	ftwSection :=strings.Split(sections[3],":")[1]
	ftwLines:=lineRe.FindAllString(ftwSection,-1)
	
	wToL :=make(map[int]int)
	wtlSection :=strings.Split(sections[4],":")[1]
	wtlLines:=lineRe.FindAllString(wtlSection,-1)

	LToT :=make(map[int]int)
	lttSection :=strings.Split(sections[5],":")[1]
	lttLines:=lineRe.FindAllString(lttSection,-1)

	tToH :=make(map[int]int)
	tthSection :=strings.Split(sections[6],":")[1]
	tthLines:=lineRe.FindAllString(tthSection,-1)

	hToL :=make(map[int]int)
	htlSection :=strings.Split(sections[7],":")[1]
	htlLines:=lineRe.FindAllString(htlSection,-1)

	// Add seeds
	digRe := regexp.MustCompile(`(\d+)`)
	// fmt.Println(sections[1])
	seedy:= digRe.FindAllString(sections[0],-1)
	for _, s:= range seedy {
		tempNum, _ :=strconv.Atoi(s)
		seeds[tempNum] = 1
	}

	// For each seed, find its appropriate soil
// Seed to Soil
	for seed := range seeds {
		fmt.Println(seed)
		noMatch:=true
		for _, line := range stsLines {
			parts := strings.Split(line, " ")
			targ,source,rng := parts[0], parts[1], parts[2]
			targNum, _ := strconv.Atoi(targ)
			sourceNum, _ := strconv.Atoi(source)
			rngNum, _ := strconv.Atoi(rng)
			if seed >= sourceNum && seed <= rngNum + sourceNum {
				noMatch = false
				dif :=  targNum - sourceNum
				sToS[seed+dif] = 1
				break
			}
		}
		if noMatch {
			sToS[seed] = 1
		}
	}

	// for each soil, find its appropriate fertilizer

	for currVal,_ := range sToS {
		noMatch:=true
		for _, line := range stfLines {
			parts := strings.Split(line, " ")
			targ,source,rng := parts[0], parts[1], parts[2]
			targNum, _ := strconv.Atoi(targ)
			sourceNum, _ := strconv.Atoi(source)
			rngNum, _ := strconv.Atoi(rng)
			if currVal >= sourceNum && currVal <= rngNum + sourceNum {
				noMatch = false
				dif :=  targNum - sourceNum
				sToF[currVal+dif] = 1
				break
			}
		}
		if noMatch {
			sToF[currVal] = 1
		}
	}
	fmt.Println(sToF)

	// for each fert, find right water
		for currVal,_ := range sToF {
		noMatch:=true
		for _, line := range ftwLines {
			parts := strings.Split(line, " ")
			targ,source,rng := parts[0], parts[1], parts[2]
			targNum, _ := strconv.Atoi(targ)
			sourceNum, _ := strconv.Atoi(source)
			rngNum, _ := strconv.Atoi(rng)
			if currVal >= sourceNum && currVal <= rngNum + sourceNum {
				noMatch = false
				dif :=  targNum - sourceNum
				fToW[currVal+dif] = 1
				break
			}
		}
		if noMatch {
			fToW[currVal] = 1
		}
	}
	fmt.Println(fToW)
	
	// for each water find right light
		for currVal,_ := range fToW {
		noMatch:=true
		for _, line := range wtlLines {
			parts := strings.Split(line, " ")
			targ,source,rng := parts[0], parts[1], parts[2]
			targNum, _ := strconv.Atoi(targ)
			sourceNum, _ := strconv.Atoi(source)
			rngNum, _ := strconv.Atoi(rng)
			if currVal >= sourceNum && currVal <= rngNum + sourceNum {
				noMatch = false
				dif :=  targNum - sourceNum
				wToL[currVal+dif] = 1
				break
			}
		}
		if noMatch {
			wToL[currVal] = 1
		}
	}
fmt.Println(wToL)
	// for each light find right temp
	for currVal,_ := range wToL{
		noMatch:=true
		for _, line := range lttLines {
			parts := strings.Split(line, " ")
			targ,source,rng := parts[0], parts[1], parts[2]
			targNum, _ := strconv.Atoi(targ)
			sourceNum, _ := strconv.Atoi(source)
			rngNum, _ := strconv.Atoi(rng)
			if currVal >= sourceNum && currVal <= rngNum + sourceNum {
				noMatch = false
				dif :=  targNum - sourceNum
				LToT[currVal+dif] = 1
				break
			}
		}
		if noMatch {
			LToT[currVal] = 1
		}
	}
fmt.Println(LToT)

	
	// for each temp find right humidity
	for currVal,_ := range LToT{
		noMatch:=true
		for _, line := range tthLines {
			parts := strings.Split(line, " ")
			targ,source,rng := parts[0], parts[1], parts[2]
			targNum, _ := strconv.Atoi(targ)
			sourceNum, _ := strconv.Atoi(source)
			rngNum, _ := strconv.Atoi(rng)
			if currVal >= sourceNum && currVal <= rngNum + sourceNum {
				noMatch = false
				dif :=  targNum - sourceNum
				tToH[currVal+dif] = 1
				break
			}
		}
		if noMatch {
			tToH[currVal] = 1
		}
	}
fmt.Println(tToH)

	// for each humidity find right loc
	for currVal,_ := range tToH{
		noMatch:=true
		for _, line := range htlLines {
			parts := strings.Split(line, " ")
			targ,source,rng := parts[0], parts[1], parts[2]
			targNum, _ := strconv.Atoi(targ)
			sourceNum, _ := strconv.Atoi(source)
			rngNum, _ := strconv.Atoi(rng)
			if currVal >= sourceNum && currVal <= rngNum + sourceNum {
				noMatch = false
				dif :=  targNum - sourceNum
				hToL[currVal+dif] = 1
				break
			}
		}
		if noMatch {
			hToL[currVal] = 1
		}
	}
fmt.Println(hToL)
minVal := math.MaxInt64
for key,_ :=range hToL {
	if key < minVal {
		minVal = key
	}
}
fmt.Println(minVal)
}

func Part2() {
	content,err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("woooops")
		os.Exit(1)
	}
	text:=string(content)
	splitRe:=regexp.MustCompile(`\n\s*\n`)
	sections:= splitRe.Split(text,-1)

	// Make maps
	seeds:= make(map[int]int)
	stsSection:= strings.Split(sections[1],":")[1]
	lineRe :=regexp.MustCompile(`(\d+) (\d+) (\d+)`)
	stsLines := lineRe.FindAllString(stsSection,-1)
	sToS :=make(map[int]int)

	sToF :=make(map[int]int)
	stfSection := strings.Split(sections[2],":")[1]
	stfLines:= lineRe.FindAllString(stfSection,-1)
	fToW :=make(map[int]int)
	ftwSection :=strings.Split(sections[3],":")[1]
	ftwLines:=lineRe.FindAllString(ftwSection,-1)
	
	wToL :=make(map[int]int)
	wtlSection :=strings.Split(sections[4],":")[1]
	wtlLines:=lineRe.FindAllString(wtlSection,-1)

	LToT :=make(map[int]int)
	lttSection :=strings.Split(sections[5],":")[1]
	lttLines:=lineRe.FindAllString(lttSection,-1)

	tToH :=make(map[int]int)
	tthSection :=strings.Split(sections[6],":")[1]
	tthLines:=lineRe.FindAllString(tthSection,-1)

	hToL :=make(map[int]int)
	htlSection :=strings.Split(sections[7],":")[1]
	htlLines:=lineRe.FindAllString(htlSection,-1)

	// Add seeds
	// fmt.Println(sections[1])
	digRe := regexp.MustCompile(`(\d+)`)
	ddRe := regexp.MustCompile(`(\d+) (\d+)`)
	seedy:= ddRe.FindAllString(sections[0],-1)
	// fmt.Println("seedy",seedy)
	for _, s:= range seedy {
		// fmt.Println("s",s)
		rng := digRe.FindAllString(s,-1)
		start,_:= strconv.Atoi(rng[0])
		spread,_:=strconv.Atoi(rng[1])
		for i:=start; i < start+spread; i++ {
			seeds[i] = 1
		}
	}
	// fmt.Println(seeds)

	// For each seed, find its appropriate soil
// Seed to Soil
	for seed := range seeds {
		// fmt.Println(seed)
		noMatch:=true
		for _, line := range stsLines {
			parts := strings.Split(line, " ")
			targ,source,rng := parts[0], parts[1], parts[2]
			targNum, _ := strconv.Atoi(targ)
			sourceNum, _ := strconv.Atoi(source)
			rngNum, _ := strconv.Atoi(rng)
			if seed >= sourceNum && seed <= rngNum + sourceNum {
				noMatch = false
				dif :=  targNum - sourceNum
				sToS[seed+dif] = 1
				break
			}
		}
		if noMatch {
			sToS[seed] = 1
		}
	}

	// for each soil, find its appropriate fertilizer

	for currVal,_ := range sToS {
		noMatch:=true
		for _, line := range stfLines {
			parts := strings.Split(line, " ")
			targ,source,rng := parts[0], parts[1], parts[2]
			targNum, _ := strconv.Atoi(targ)
			sourceNum, _ := strconv.Atoi(source)
			rngNum, _ := strconv.Atoi(rng)
			if currVal >= sourceNum && currVal <= rngNum + sourceNum {
				noMatch = false
				dif :=  targNum - sourceNum
				sToF[currVal+dif] = 1
				break
			}
		}
		if noMatch {
			sToF[currVal] = 1
		}
	}
	// fmt.Println(sToF)

	// for each fert, find right water
		for currVal,_ := range sToF {
		noMatch:=true
		for _, line := range ftwLines {
			parts := strings.Split(line, " ")
			targ,source,rng := parts[0], parts[1], parts[2]
			targNum, _ := strconv.Atoi(targ)
			sourceNum, _ := strconv.Atoi(source)
			rngNum, _ := strconv.Atoi(rng)
			if currVal >= sourceNum && currVal <= rngNum + sourceNum {
				noMatch = false
				dif :=  targNum - sourceNum
				fToW[currVal+dif] = 1
				break
			}
		}
		if noMatch {
			fToW[currVal] = 1
		}
	}
	// fmt.Println(fToW)
	
	// for each water find right light
		for currVal,_ := range fToW {
		noMatch:=true
		for _, line := range wtlLines {
			parts := strings.Split(line, " ")
			targ,source,rng := parts[0], parts[1], parts[2]
			targNum, _ := strconv.Atoi(targ)
			sourceNum, _ := strconv.Atoi(source)
			rngNum, _ := strconv.Atoi(rng)
			if currVal >= sourceNum && currVal <= rngNum + sourceNum {
				noMatch = false
				dif :=  targNum - sourceNum
				wToL[currVal+dif] = 1
				break
			}
		}
		if noMatch {
			wToL[currVal] = 1
		}
	}
// fmt.Println(wToL)
	// for each light find right temp
	for currVal,_ := range wToL{
		noMatch:=true
		for _, line := range lttLines {
			parts := strings.Split(line, " ")
			targ,source,rng := parts[0], parts[1], parts[2]
			targNum, _ := strconv.Atoi(targ)
			sourceNum, _ := strconv.Atoi(source)
			rngNum, _ := strconv.Atoi(rng)
			if currVal >= sourceNum && currVal <= rngNum + sourceNum {
				noMatch = false
				dif :=  targNum - sourceNum
				LToT[currVal+dif] = 1
				break
			}
		}
		if noMatch {
			LToT[currVal] = 1
		}
	}
// fmt.Println(LToT)

	
	// for each temp find right humidity
	for currVal,_ := range LToT{
		noMatch:=true
		for _, line := range tthLines {
			parts := strings.Split(line, " ")
			targ,source,rng := parts[0], parts[1], parts[2]
			targNum, _ := strconv.Atoi(targ)
			sourceNum, _ := strconv.Atoi(source)
			rngNum, _ := strconv.Atoi(rng)
			if currVal >= sourceNum && currVal <= rngNum + sourceNum {
				noMatch = false
				dif :=  targNum - sourceNum
				tToH[currVal+dif] = 1
				break
			}
		}
		if noMatch {
			tToH[currVal] = 1
		}
	}
// fmt.Println(tToH)

	// for each humidity find right loc
	for currVal,_ := range tToH{
		noMatch:=true
		for _, line := range htlLines {
			parts := strings.Split(line, " ")
			targ,source,rng := parts[0], parts[1], parts[2]
			targNum, _ := strconv.Atoi(targ)
			sourceNum, _ := strconv.Atoi(source)
			rngNum, _ := strconv.Atoi(rng)
			if currVal >= sourceNum && currVal <= rngNum + sourceNum {
				noMatch = false
				dif :=  targNum - sourceNum
				hToL[currVal+dif] = 1
				break
			}
		}
		if noMatch {
			hToL[currVal] = 1
		}
	}
// fmt.Println(hToL)
minVal := math.MaxInt64
for key,_ :=range hToL {
	if key < minVal {
		minVal = key
	}
}
fmt.Println(minVal)
}