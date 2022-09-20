package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
)

var (
	titleMatcher = regexp.MustCompile(`--- scanner \d+ ---`)
	emptyMatcher = regexp.MustCompile(`\s*`)
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var (
		id     int
		origBS []beaconScanner
		input  beaconScanner
	)
	for scanner.Scan() {
		if titleMatcher.MatchString(scanner.Text()) {
			input = beaconScanner{id: id}
			continue
		}
		if len(scanner.Text()) == 0 {
			id++
			origBS = append(origBS, input)
			continue
		}
		var x, y, z int
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d,%d", &x, &y, &z)
		if err != nil {
			log.Fatalf("Failed to scan coordinate %s: %v", scanner.Text(), err)
		}
		input.beacons = append(input.beacons, beacon{point{x, y, z}, map[int]bool{}})
	}
	origBS = append(origBS, input)

	for _, s := range origBS {
		for i, b := range s.beacons {
			for j, other := range s.beacons {
				if i == j {
					continue
				}
				b.distMap[b.dist(other.point)] = true
			}
		}
	}

	var normBS []beaconScanner
	normBS = append(normBS, origBS[0])
	origBS[0] = origBS[len(origBS)-1]
	origBS = origBS[:len(origBS)-1]
	for len(origBS) > 0 {
		fmt.Printf("Norm Size: %d Orig Size: %d\n", len(normBS), len(origBS))
		found, index := findMatchAndNormalize(normBS, origBS)
		if !found {
			log.Fatal("Failed to find a match.")
		}
		normBS = append(normBS, origBS[index])
		origBS[index] = origBS[len(origBS)-1]
		origBS = origBS[:len(origBS)-1]
	}

	var maxDist int
	for i := 0; i < len(normBS); i++ {
		for j := i + 1; j < len(normBS); j++ {
			dist := math.Abs(float64(normBS[i].x-normBS[j].x)) + math.Abs(float64(normBS[i].y-normBS[j].y)) + math.Abs(float64(normBS[i].z-normBS[j].z))
			if int(dist) > maxDist {
				maxDist = int(dist)
			}
		}
	}

	fmt.Printf("Maximum scanner distance is %d\n", maxDist)
}

func findMatchAndNormalize(normBS, origBS []beaconScanner) (bool, int) {
	for _, norm := range normBS {
		for j, orig := range origBS {
			found, matches := isScannerMatch(norm, orig)
			if found {
				fmt.Printf("Found scanner match between %d and %d.\n", norm.id, orig.id)
				bn, err := newBeaconNormalizer(matches)
				fmt.Println("Normalizer:", bn.shiftX, bn.shiftY, bn.shiftZ)
				if err != nil {
					return false, 0
				}
				orig.point = point{bn.shiftX, bn.shiftY, bn.shiftZ}
				orig.normalize(bn)
				origBS[j] = orig
				return true, j
			}
		}
	}
	return false, 0
}

func isScannerMatch(norm, orig beaconScanner) (bool, []pointPair) {
	var pointMatches []pointPair
	for _, nb := range norm.beacons {
		for _, ob := range orig.beacons {
			count := 0
			for normDist := range nb.distMap {
				if ob.distMap[normDist] {
					count++
				}
			}
			if count == 11 {
				pointMatches = append(pointMatches, pointPair{nb.point, ob.point})
				// fmt.Printf("Found a match between Scanner %d (%d,%d,%d) and Scanner %d (%d,%d,%d)\n", norm.id, nb.x, nb.y, nb.z, orig.id, ob.x, ob.y, ob.z)
			}
		}
	}
	if len(pointMatches) < 12 {
		return false, nil
	}
	return true, pointMatches
}

func newBeaconNormalizer(pointMatches []pointPair) (beaconNormalizer, error) {
	var pn beaconNormalizer
	for i := 0; i < len(pointMatches)-1; i++ {
		firstMatch, secondMatch := pointMatches[i], pointMatches[i+1]
		firstDiff := firstMatch.first.diff(secondMatch.first)
		absFirstDiff := firstDiff.abs()
		// skip these points if not all shifts are unique
		if absFirstDiff.x == absFirstDiff.y || absFirstDiff.x == absFirstDiff.z || absFirstDiff.y == absFirstDiff.z {
			fmt.Println("Skipping points.")
			continue
		}
		secondDiff := firstMatch.second.diff(secondMatch.second)
		absSecondDiff := secondDiff.abs()
		if absFirstDiff.x == absSecondDiff.x {
			pn.mappers = append(pn.mappers, func(base point, norm *point) {
				norm.x = base.x
			})
			pn.invX = firstDiff.x / secondDiff.x
			pn.shiftX = firstMatch.first.x - firstMatch.second.x*pn.invX
		}
		if absFirstDiff.x == absSecondDiff.y {
			pn.mappers = append(pn.mappers, func(base point, norm *point) {
				norm.x = base.y
			})
			pn.invX = firstDiff.x / secondDiff.y
			pn.shiftX = firstMatch.first.x - firstMatch.second.y*pn.invX
		}
		if absFirstDiff.x == absSecondDiff.z {
			pn.mappers = append(pn.mappers, func(base point, norm *point) {
				norm.x = base.z
			})
			pn.invX = firstDiff.x / secondDiff.z
			pn.shiftX = firstMatch.first.x - firstMatch.second.z*pn.invX
		}
		if absFirstDiff.y == absSecondDiff.y {
			pn.mappers = append(pn.mappers, func(base point, norm *point) {
				norm.y = base.y
			})
			pn.invY = firstDiff.y / secondDiff.y
			pn.shiftY = firstMatch.first.y - firstMatch.second.y*pn.invY
		}
		if absFirstDiff.y == absSecondDiff.x {
			pn.mappers = append(pn.mappers, func(base point, norm *point) {
				norm.y = base.x
			})
			pn.invY = firstDiff.y / secondDiff.x
			pn.shiftY = firstMatch.first.y - firstMatch.second.x*pn.invY
		}
		if absFirstDiff.y == absSecondDiff.z {
			pn.mappers = append(pn.mappers, func(base point, norm *point) {
				norm.y = base.z
			})
			pn.invY = firstDiff.y / secondDiff.z
			pn.shiftY = firstMatch.first.y - firstMatch.second.z*pn.invY
		}
		if absFirstDiff.z == absSecondDiff.z {
			pn.mappers = append(pn.mappers, func(base point, norm *point) {
				norm.z = base.z
			})
			pn.invZ = firstDiff.z / secondDiff.z
			pn.shiftZ = firstMatch.first.z - firstMatch.second.z*pn.invZ
		}
		if absFirstDiff.z == absSecondDiff.x {
			pn.mappers = append(pn.mappers, func(base point, norm *point) {
				norm.z = base.x
			})
			pn.invZ = firstDiff.z / secondDiff.x
			pn.shiftZ = firstMatch.first.z - firstMatch.second.x*pn.invZ
		}
		if absFirstDiff.z == absSecondDiff.y {
			pn.mappers = append(pn.mappers, func(base point, norm *point) {
				norm.z = base.y
			})
			pn.invZ = firstDiff.z / secondDiff.y
			pn.shiftZ = firstMatch.first.z - firstMatch.second.y*pn.invZ
		}
		return pn, nil
	}
	fmt.Println("Failed to create normailizer.")
	return pn, errors.New("Unable to create beacon normailizer from point matches.")
}

type pointMapper func(point, *point)

type beaconNormalizer struct {
	invX, invY, invZ       int
	shiftX, shiftY, shiftZ int
	mappers                []pointMapper
}

func (n beaconNormalizer) normalize(base point) point {
	var norm point
	for _, m := range n.mappers {
		m(base, &norm)
	}
	norm.x = norm.x*n.invX + n.shiftX
	norm.y = norm.y*n.invY + n.shiftY
	norm.z = norm.z*n.invZ + n.shiftZ
	return norm
}

type beaconScanner struct {
	id int
	point
	beacons []beacon
}

func (s beaconScanner) normalize(bn beaconNormalizer) {
	for i, b := range s.beacons {
		norm := bn.normalize(b.point)
		s.beacons[i] = beacon{norm, b.distMap}
	}
}

type beacon struct {
	point
	distMap map[int]bool
}

type point struct {
	x, y, z int
}

func (p point) diff(other point) point {
	return point{p.x - other.x, p.y - other.y, p.z - other.z}
}

func (p point) abs() point {
	x := int(math.Abs(float64(p.x)))
	y := int(math.Abs(float64(p.y)))
	z := int(math.Abs(float64(p.z)))
	return point{x, y, z}
}

func (p point) dist(other point) int {
	x := p.x - other.x
	y := p.y - other.y
	z := p.z - other.z
	return x*x + y*y + z*z
}

type pointPair struct {
	first, second point
}
