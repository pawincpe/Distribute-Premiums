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
		allocated := make(map[int]int)

		// ขั้นตอนที่ 4: การจัดสรรตามสัดส่วนและเติมเศษให้ครบ
		totalProductQty := 0.0
		for _, premium := range product.Premiums {
			totalProductQty += premium.QtyUse
		}

		// 1. คำนวณสัดส่วน, ปัดเศษลง, เก็บเศษ
		type allocInfo struct {
			code      int
			floorVal  int
			frac      float64
			available int
		}
		allocs := make([]allocInfo, 0, len(product.Premiums))
		sumFloor := 0
		for _, premium := range product.Premiums {
			proportion := premium.QtyUse / totalProductQty
			allocF := proportion * float64(product.Qty)
			floorVal := int(math.Floor(allocF))
			frac := allocF - float64(floorVal)
			available := promotionRounded[premium.PromotionCode]
			if floorVal > available {
				floorVal = available
				frac = 0
			}
			allocs = append(allocs, allocInfo{premium.PromotionCode, floorVal, frac, available})
			sumFloor += floorVal
		}

		// 2. แจกจ่ายเศษที่เหลือให้ครบ Qty
		remaining := product.Qty - sumFloor
		for remaining > 0 && len(allocs) > 0 {
			found := false
			for i := 0; i < len(allocs) && remaining > 0; i++ {
				if allocs[i].floorVal < allocs[i].available && promotionRounded[allocs[i].code] > 0 {
					allocs[i].floorVal++
					promotionRounded[allocs[i].code]--
					remaining--
					found = true
				}
			}
			if !found {
				break // ไม่มีใครรับเศษได้แล้ว
			}
		}

		// 3. บันทึกผลลัพธ์
		for _, a := range allocs {
			if a.floorVal > 0 {
				allocated[a.code] = a.floorVal
			}
		}

		// 4. ตรวจสอบและปรับแก้ในบรรทัดสุดท้าย
		totalAllocated := 0
		for _, qty := range allocated {
			totalAllocated += qty
		}

		// ถ้าผลรวมไม่เท่ากับ Qty ให้ปรับแก้
		if totalAllocated != product.Qty {
			diff := product.Qty - totalAllocated

			if diff > 0 {
				// ต้องเพิ่ม diff หน่วย
				// หา promotion ที่มี available มากที่สุด
				bestCode := -1
				bestAvailable := 0
				for code, qty := range allocated {
					available := promotionRounded[code] + qty // จำนวนที่ยังไม่ถูกใช้
					if available > bestAvailable {
						bestAvailable = available
						bestCode = code
					}
				}
				if bestCode != -1 {
					allocated[bestCode] += diff
				}
			} else if diff < 0 {
				// ต้องลด diff หน่วย
				// หา promotion ที่มี qty มากที่สุด
				bestCode := -1
				bestQty := 0
				for code, qty := range allocated {
					if qty > bestQty {
						bestQty = qty
						bestCode = code
					}
				}
				if bestCode != -1 {
					allocated[bestCode] += diff // diff เป็นลบอยู่แล้ว
				}
			}
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
