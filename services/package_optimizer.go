package services

import (
	"math"
)

type packOptimizer struct {
	packSizes []int
	order     int
}

func newPackOptimizer(packSizes []int, order int) *packOptimizer {
	return &packOptimizer{
		packSizes: packSizes,
		order:     order,
	}
}

func (p *packOptimizer) findOptimalPacks(order int, packSizes []int, currentSolution []int) []int {
	if order == 0 {
		return currentSolution
	}
	if order < 0 || len(packSizes) == 0 {
		return nil
	}

	withCurrentPack := p.findOptimalPacks(order-packSizes[0], packSizes, append(currentSolution, packSizes[0]))

	withoutCurrentPack := p.findOptimalPacks(order, packSizes[1:], currentSolution)

	if withCurrentPack == nil {
		return withoutCurrentPack
	} else if withoutCurrentPack == nil {
		return withCurrentPack
	} else {
		if len(withCurrentPack) <= len(withoutCurrentPack) {
			return withCurrentPack
		} else {
			return withoutCurrentPack
		}
	}
}

// PackSplitter represents an instance of the package splitting algorithm
type PackSplitter struct {
	packSizes []int
	order     int
}

// NewPackSplitter creates a new PackSplitter instance
func NewPackSplitter(packSizes []int, order int) *PackSplitter {
	return &PackSplitter{
		packSizes: packSizes,
		order:     order,
	}
}

func (p *PackSplitter) minimumRemaining() int {
	result := int(math.Ceil(float64(p.order) / float64(p.packSizes[0])))
	minRemain := result*p.packSizes[0] - p.order
	for i := 1; i < len(p.packSizes); i++ {
		result = int(math.Ceil(float64(p.order) / float64(p.packSizes[i])))
		remain := result*p.packSizes[i] - p.order
		if minRemain > remain {
			minRemain = remain
		}
	}
	return minRemain
}

func (p *PackSplitter) splitOrderIntoPacksArray() []int {
	minRemainining := p.minimumRemaining()
	var bestOptimalSolution []int
	for i := 0; i < minRemainining+1; i++ {
		exceededOrder := p.order + i
		packOptimizer := newPackOptimizer(p.packSizes, exceededOrder)
		solution := packOptimizer.findOptimalPacks(packOptimizer.order, packOptimizer.packSizes, []int{})
		if solution != nil {
			bestOptimalSolution = solution
			break
		}
	}
	return bestOptimalSolution
}

func countPackAmountsFromResultArray(result []int) map[int]int {
	//Create a   dictionary of values for each element
	packSizesWithAmounts := make(map[int]int)
	for _, packSize := range result {
		packSizesWithAmounts[packSize] = packSizesWithAmounts[packSize] + 1
	}
	return packSizesWithAmounts
}

// SplitOrderIntoPacks splits the order amount into packs using PackSplitter.packSizes
func (p *PackSplitter) SplitOrderIntoPacks() map[int]int {
	packSizesArray := p.splitOrderIntoPacksArray()
	packSizesWithAmount := countPackAmountsFromResultArray(packSizesArray)
	return packSizesWithAmount
}
