package distribute

import (
	"math"
	"sort"
)

// Premium เก็บข้อมูลโปรโมชั่นสำหรับ output (จำนวนเต็ม)
type Premium struct {
	PromotionCode int // รหัสโปรโมชั่น
	QtyUse        int // จำนวนสินค้าที่ใช้กับโปรโมชั่นนี้ (จำนวนเต็ม)
}

// Product เก็บข้อมูลสินค้าสำหรับ output
type Product struct {
	Product  string    // ชื่อสินค้า
	Qty      int       // จำนวนสินค้าที่ซื้อ
	Premiums []Premium // รายการโปรโมชั่นที่จัดสรรแล้ว
}

// InputPremium เก็บข้อมูลโปรโมชั่นสำหรับ input (ทศนิยม)
type InputPremium struct {
	PromotionCode int     // รหัสโปรโมชั่น
	QtyUse        float64 // จำนวนสินค้าที่ใช้กับโปรโมชั่นนี้ (ทศนิยม)
}

// InputProduct เก็บข้อมูลสินค้าสำหรับ input
type InputProduct struct {
	Product  string         // ชื่อสินค้า
	Qty      int            // จำนวนสินค้าที่ซื้อ
	Premiums []InputPremium // รายการโปรโมชั่นที่เกี่ยวข้อง
}

// DistributePremiums จัดสรรสินค้ากับโปรโมชั่นตามสัดส่วนที่เหมาะสม
// รับ input เป็นรายการสินค้าที่มีโปรโมชั่นและจำนวนที่ต้องการ
// ส่งคืนรายการสินค้าที่มีการจัดสรรโปรโมชั่นเป็นจำนวนเต็มแล้ว
func DistributePremiums(products []InputProduct) []Product {
	// ขั้นตอนที่ 1: คำนวณยอดรวม qtyUse ของแต่ละโปรโมชั่นจากทุกสินค้า
	promotionTotals := make(map[int]float64)
	for _, product := range products {
		for _, premium := range product.Premiums {
			promotionTotals[premium.PromotionCode] += premium.QtyUse
		}
	}

	// ขั้นตอนที่ 2: ปัดเศษยอดรวมโปรโมชั่นให้เป็นจำนวนเต็ม
	promotionRounded := make(map[int]int)
	for code, total := range promotionTotals {
		promotionRounded[code] = int(math.Round(total))
	}

	// ขั้นตอนที่ 3: สร้างผลลัพธ์สินค้า
	output := make([]Product, len(products))

	for i, product := range products {
		remainingQty := product.Qty
		allocated := make(map[int]int)

		// ขั้นตอนที่ 4: การจัดสรรครั้งแรก - จัดสรรตามสัดส่วนเดิม
		// เรียงลำดับโปรโมชั่นตาม qtyUse (จากมากไปน้อย) เพื่อให้ความสำคัญกับโปรโมชั่นที่มีการจัดสรรมากกว่า
		sortedPremiums := make([]InputPremium, len(product.Premiums))
		copy(sortedPremiums, product.Premiums)
		sort.SliceStable(sortedPremiums, func(i, j int) bool {
			return sortedPremiums[i].QtyUse > sortedPremiums[j].QtyUse
		})

		// จัดสรรตามความจุโปรโมชั่นที่เหลือ
		for _, premium := range sortedPremiums {
			code := premium.PromotionCode
			available := promotionRounded[code]

			if available > 0 && remainingQty > 0 {
				// จัดสรรอย่างน้อย 1 ถ้าโปรโมชั่นนี้มีการจัดสรรใดๆ
				alloc := 1
				allocated[code] = alloc
				promotionRounded[code] -= alloc
				remainingQty -= alloc
			}
		}

		// ขั้นตอนที่ 5: การจัดสรรครั้งที่สอง - เติมส่วนที่เหลือด้วยโปรโมชั่นที่มีอยู่
		for remainingQty > 0 {
			// หาโปรโมชั่นที่ยังมีความจุเหลือ
			availablePromotions := make([]int, 0)
			for code, available := range promotionRounded {
				if available > 0 {
					availablePromotions = append(availablePromotions, code)
				}
			}

			if len(availablePromotions) == 0 {
				break // ไม่มีโปรโมชั่นเหลือ
			}

			// เรียงลำดับตามความจุที่เหลือ (จากมากไปน้อย)
			sort.SliceStable(availablePromotions, func(i, j int) bool {
				return promotionRounded[availablePromotions[i]] > promotionRounded[availablePromotions[j]]
			})

			// จัดสรรให้กับโปรโมชั่นที่มีความจุเหลือมากที่สุด
			selectedCode := availablePromotions[0]
			allocated[selectedCode]++
			promotionRounded[selectedCode]--
			remainingQty--
		}

		// ขั้นตอนที่ 6: แปลง map ที่จัดสรรแล้วเป็น slice ของ Premium
		premiums := make([]Premium, 0, len(allocated))
		for code, qty := range allocated {
			premiums = append(premiums, Premium{
				PromotionCode: code,
				QtyUse:        qty,
			})
		}

		// เรียงลำดับโปรโมชั่นตามรหัสโปรโมชั่นเพื่อผลลัพธ์ที่สม่ำเสมอ
		sort.SliceStable(premiums, func(i, j int) bool {
			return premiums[i].PromotionCode < premiums[j].PromotionCode
		})

		output[i] = Product{
			Product:  product.Product,
			Qty:      product.Qty,
			Premiums: premiums,
		}
	}

	return output
}
