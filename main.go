package main

import (
	"fmt"

	"app/pkg/distribute"
)

func main() {
	fmt.Println("=== Premium Distribution Test Cases ===")

	// Test Case 1: Basic distribution
	fmt.Println("Test Case 1: Basic Distribution")
	testCase1()

	// Test Case 2: Multiple products with different quantities
	fmt.Println("\nTest Case 2: Multiple Products with Different Quantities")
	testCase2()

	// Test Case 3: Complex distribution with many promotions
	fmt.Println("\nTest Case 3: Complex Distribution")
	testCase3()

	// Test Case 4: Edge case with small quantities
	fmt.Println("\nTest Case 4: Edge Case with Small Quantities")
	testCase4()

	// Test Case 5: User requested case
	fmt.Println("\nTest Case 5: User Requested Case")
	testCase5()

	// Test Case 6: Large quantity distribution
	fmt.Println("\nTest Case 6: Large Quantity Distribution")
	testCase6()

	// Test Case 7: User provided case with 100 qty and 3 promotions
	fmt.Println("\nTest Case 7: User Provided Case with 100 Qty and 3 Promotions")
	testCase7()
}

func testCase1() {
	input := []distribute.InputProduct{
		{
			Product: "A",
			Qty:     3,
			Premiums: []distribute.InputPremium{
				{PromotionCode: 110, QtyUse: 1.5},
				{PromotionCode: 111, QtyUse: 1.5},
			},
		},
		{
			Product: "B",
			Qty:     2,
			Premiums: []distribute.InputPremium{
				{PromotionCode: 110, QtyUse: 1.0},
				{PromotionCode: 112, QtyUse: 1.0},
			},
		},
	}

	output := distribute.DistributePremiums(input)
	printResults(input, output)
}

func testCase2() {
	input := []distribute.InputProduct{
		{
			Product: "Product1",
			Qty:     5,
			Premiums: []distribute.InputPremium{
				{PromotionCode: 201, QtyUse: 2.3},
				{PromotionCode: 202, QtyUse: 1.7},
				{PromotionCode: 203, QtyUse: 1.0},
			},
		},
		{
			Product: "Product2",
			Qty:     3,
			Premiums: []distribute.InputPremium{
				{PromotionCode: 201, QtyUse: 1.5},
				{PromotionCode: 204, QtyUse: 1.5},
			},
		},
		{
			Product: "Product3",
			Qty:     2,
			Premiums: []distribute.InputPremium{
				{PromotionCode: 202, QtyUse: 1.2},
				{PromotionCode: 203, QtyUse: 0.8},
			},
		},
	}

	output := distribute.DistributePremiums(input)
	printResults(input, output)
}

func testCase3() {
	input := []distribute.InputProduct{
		{
			Product: "Laptop",
			Qty:     10,
			Premiums: []distribute.InputPremium{
				{PromotionCode: 301, QtyUse: 4.2},
				{PromotionCode: 302, QtyUse: 3.1},
				{PromotionCode: 303, QtyUse: 2.7},
			},
		},
		{
			Product: "Mouse",
			Qty:     8,
			Premiums: []distribute.InputPremium{
				{PromotionCode: 301, QtyUse: 2.8},
				{PromotionCode: 304, QtyUse: 3.2},
				{PromotionCode: 305, QtyUse: 2.0},
			},
		},
		{
			Product: "Keyboard",
			Qty:     6,
			Premiums: []distribute.InputPremium{
				{PromotionCode: 302, QtyUse: 2.5},
				{PromotionCode: 303, QtyUse: 1.8},
				{PromotionCode: 305, QtyUse: 1.7},
			},
		},
	}

	output := distribute.DistributePremiums(input)
	printResults(input, output)
}

func testCase4() {
	input := []distribute.InputProduct{
		{
			Product: "Small",
			Qty:     1,
			Premiums: []distribute.InputPremium{
				{PromotionCode: 401, QtyUse: 0.3},
				{PromotionCode: 402, QtyUse: 0.7},
			},
		},
		{
			Product: "Tiny",
			Qty:     1,
			Premiums: []distribute.InputPremium{
				{PromotionCode: 401, QtyUse: 0.4},
				{PromotionCode: 403, QtyUse: 0.6},
			},
		},
	}

	output := distribute.DistributePremiums(input)
	printResults(input, output)
}

func testCase5() {
	input := []distribute.InputProduct{
		{
			Product: "A",
			Qty:     1,
			Premiums: []distribute.InputPremium{
				{PromotionCode: 110, QtyUse: 0.5},
				{PromotionCode: 111, QtyUse: 0.5},
			},
		},
		{
			Product: "B",
			Qty:     1,
			Premiums: []distribute.InputPremium{
				{PromotionCode: 110, QtyUse: 0.5},
				{PromotionCode: 111, QtyUse: 0.5},
			},
		},
	}

	output := distribute.DistributePremiums(input)
	printResults(input, output)
}

func testCase6() {
	input := []distribute.InputProduct{
		{
			Product: "A",
			Qty:     29,
			Premiums: []distribute.InputPremium{
				{PromotionCode: 110, QtyUse: 17.5},
				{PromotionCode: 111, QtyUse: 11.4},
			},
		},
		{
			Product: "B",
			Qty:     1,
			Premiums: []distribute.InputPremium{
				{PromotionCode: 110, QtyUse: 0.5},
				{PromotionCode: 111, QtyUse: 0.5},
			},
		},
	}

	output := distribute.DistributePremiums(input)
	printResults(input, output)
}

func testCase7() {
	input := []distribute.InputProduct{
		{
			Product: "A",
			Qty:     100,
			Premiums: []distribute.InputPremium{
				{PromotionCode: 110, QtyUse: 33.33},
				{PromotionCode: 111, QtyUse: 33.33},
				{PromotionCode: 112, QtyUse: 33.33},
			},
		},
	}

	output := distribute.DistributePremiums(input)
	printResults(input, output)
}

func printResults(input []distribute.InputProduct, output []distribute.Product) {
	fmt.Println("Input:")
	for _, p := range input {
		fmt.Printf("  Product: %s, Qty: %d\n", p.Product, p.Qty)
		for _, prem := range p.Premiums {
			fmt.Printf("    PromotionCode: %d, QtyUse: %.1f\n", prem.PromotionCode, prem.QtyUse)
		}
	}

	fmt.Println("\nOutput:")
	for _, p := range output {
		fmt.Printf("  Product: %s, Qty: %d\n", p.Product, p.Qty)
		totalQty := 0
		for _, prem := range p.Premiums {
			fmt.Printf("    PromotionCode: %d, QtyUse: %d\n", prem.PromotionCode, prem.QtyUse)
			totalQty += prem.QtyUse
		}
		fmt.Printf("    Total QtyUse: %d (should equal Qty: %d)\n", totalQty, p.Qty)
	}

	// Verify promotion totals
	fmt.Println("\nPromotion Totals:")
	promotionTotals := make(map[int]int)
	for _, p := range output {
		for _, prem := range p.Premiums {
			promotionTotals[prem.PromotionCode] += prem.QtyUse
		}
	}
	for code, total := range promotionTotals {
		fmt.Printf("  PromotionCode %d: Total QtyUse = %d\n", code, total)
	}
	fmt.Println("---")
}
