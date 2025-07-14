# Premium Distributor

ระบบจัดสรรสินค้ากับโปรโมชั่นของแถมที่ช่วยกระจายจำนวนสินค้าตามสัดส่วนที่เหมาะสม

## 🎯 วัตถุประสงค์

ระบบนี้ถูกออกแบบมาเพื่อแก้ปัญหาการจัดสรรสินค้าหลายรายการให้กับโปรโมชั่นหลายรายการ โดยมีเงื่อนไขสำคัญดังนี้:

- ✅ เกลี่ย qtyUse ของแต่ละ promotionCode ให้เป็นจำนวนเต็ม
- ✅ กระจายของซื้อให้ครบทุกโปรโมชั่นตามสัดส่วนที่เหมาะสม
- ✅ แต่ละสินค้าต้องมี qtyUse รวมเท่ากับ qty
- ✅ แต่ละโปรโมชั่นต้องได้รับจำนวน qtyUse รวมตามที่คำนวณจาก input
- ✅ ไม่มีโปรโมชั่นซ้ำในรายการเดียวกัน

## 📁 โครงสร้างโปรเจค

```
distribute-premiums/
├── main.go                 # ไฟล์หลักสำหรับทดสอบ
├── go.mod                  # Go module configuration
├── README.md              # เอกสารนี้
└── pkg/
    └── distribute/
        └── distribute.go   # Logic หลักสำหรับการจัดสรร
```

## 🔧 การติดตั้ง

```bash
# Clone หรือ download โปรเจค
cd distribute-premiums

# รันโปรแกรมทดสอบ
go run main.go

# หรือ build และรัน
go build
./distribute-premiums
```

## 📊 โครงสร้างข้อมูล

### Input Types

```go
// ข้อมูลโปรโมชั่นสำหรับ input
type InputPremium struct {
    PromotionCode int     // รหัสโปรโมชั่น
    QtyUse        float64 // จำนวนสินค้าที่ใช้กับโปรโมชั่นนี้ (ทศนิยม)
}

// ข้อมูลสินค้าสำหรับ input
type InputProduct struct {
    Product  string           // ชื่อสินค้า
    Qty      int             // จำนวนสินค้าที่ซื้อ
    Premiums []InputPremium  // รายการโปรโมชั่นที่เกี่ยวข้อง
}
```

### Output Types

```go
// ข้อมูลโปรโมชั่นสำหรับ output
type Premium struct {
    PromotionCode int // รหัสโปรโมชั่น
    QtyUse        int // จำนวนสินค้าที่ใช้กับโปรโมชั่นนี้ (จำนวนเต็ม)
}

// ข้อมูลสินค้าสำหรับ output
type Product struct {
    Product  string    // ชื่อสินค้า
    Qty      int      // จำนวนสินค้าที่ซื้อ
    Premiums []Premium // รายการโปรโมชั่นที่จัดสรรแล้ว
}
```

## 💻 การใช้งาน

### ตัวอย่างพื้นฐาน

```go
package main

import (
    "fmt"
    "app/pkg/distribute"
)

func main() {
    // สร้าง input data
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
    }

    // เรียกใช้ function จัดสรร
    output := distribute.DistributePremiums(input)
    
    // แสดงผลลัพธ์
    for _, product := range output {
        fmt.Printf("Product: %s, Qty: %d\n", product.Product, product.Qty)
        for _, premium := range product.Premiums {
            fmt.Printf("  PromotionCode: %d, QtyUse: %d\n", 
                premium.PromotionCode, premium.QtyUse)
        }
    }
}
```

## 🔍 ตัวอย่างผลลัพธ์

### ตัวอย่างที่ 1: การจัดสรรพื้นฐาน

**Input:**
```
Product: A, Qty: 3
  PromotionCode: 110, QtyUse: 1.5
  PromotionCode: 111, QtyUse: 1.5
Product: B, Qty: 2
  PromotionCode: 110, QtyUse: 1.0
  PromotionCode: 112, QtyUse: 1.0
```

**Output:**
```
Product: A, Qty: 3
  PromotionCode: 110, QtyUse: 2
  PromotionCode: 111, QtyUse: 1
  Total QtyUse: 3 (should equal Qty: 3)
Product: B, Qty: 2
  PromotionCode: 110, QtyUse: 1
  PromotionCode: 112, QtyUse: 1
  Total QtyUse: 2 (should equal Qty: 2)

Promotion Totals:
  PromotionCode 110: Total QtyUse = 3
  PromotionCode 111: Total QtyUse = 1
  PromotionCode 112: Total QtyUse = 1
```

### ตัวอย่างที่ 2: กรณีพิเศษ

**Input:**
```
Product: A, Qty: 1
  PromotionCode: 110, QtyUse: 0.5
  PromotionCode: 111, QtyUse: 0.5
Product: B, Qty: 1
  PromotionCode: 110, QtyUse: 0.5
  PromotionCode: 111, QtyUse: 0.5
```

**Output:**
```
Product: A, Qty: 1
  PromotionCode: 110, QtyUse: 1
  Total QtyUse: 1 (should equal Qty: 1)
Product: B, Qty: 1
  PromotionCode: 111, QtyUse: 1
  Total QtyUse: 1 (should equal Qty: 1)

Promotion Totals:
  PromotionCode 110: Total QtyUse = 1
  PromotionCode 111: Total QtyUse = 1
```

## 🧮 วิธีการทำงาน

### ขั้นตอนการประมวลผล

1. **คำนวณยอดรวม (Calculate Totals)**
   - รวม qtyUse ของแต่ละ promotionCode จากทุกสินค้า
   - ตัวอย่าง: PromotionCode 110 = 1.5 + 1.0 = 2.5

2. **ปัดเศษ (Round Totals)**
   - ปัดเศษยอดรวมให้เป็นจำนวนเต็ม
   - ตัวอย่าง: 2.5 → 3

3. **จัดสรรครั้งแรก (First Allocation)**
   - จัดสรรตามสัดส่วนเดิม โดยเรียงลำดับจากมากไปน้อย
   - ให้ความสำคัญกับโปรโมชั่นที่มี qtyUse สูงกว่า

4. **จัดสรรครั้งที่สอง (Second Allocation)**
   - กระจายส่วนที่เหลือให้กับโปรโมชั่นที่มีความจุเหลือ
   - ใช้ algorithm แบบ greedy เพื่อให้ได้ผลลัพธ์ที่ดีที่สุด

5. **ตรวจสอบ (Validation)**
   - รับประกันว่าแต่ละสินค้ามี qtyUse รวมเท่ากับ qty
   - ตรวจสอบว่าไม่มีโปรโมชั่นซ้ำในรายการเดียวกัน

### Algorithm Pseudo Code

```
1. Calculate total qtyUse for each promotion across all products
2. Round promotion totals to integers
3. For each product:
   a. Sort premiums by qtyUse (descending)
   b. Allocate based on available promotion capacity
   c. Fill remaining qty with available promotions
4. Convert allocated map to Premium slice
5. Sort premiums by promotion code for consistent output
```

## 🧪 การทดสอบ

โปรแกรมมี test cases 5 กรณีที่ครอบคลุมสถานการณ์ต่างๆ:

1. **Basic Distribution** - การจัดสรรพื้นฐาน
2. **Multiple Products** - สินค้าหลายรายการที่มีจำนวนต่างกัน
3. **Complex Distribution** - การจัดสรรที่ซับซ้อน
4. **Edge Cases** - กรณีพิเศษที่มีจำนวนน้อย
5. **User Requested Case** - กรณีที่ผู้ใช้ต้องการทดสอบ

### รันการทดสอบ

```bash
go run main.go
```

แต่ละ test case จะแสดง:
- ✅ Input ที่ใช้
- ✅ Output ที่ได้
- ✅ ยอดรวมของแต่ละโปรโมชั่น
- ✅ การตรวจสอบว่า qtyUse รวมเท่ากับ qty

## 🔧 การปรับแต่ง

### การเพิ่ม Test Case ใหม่

```go
func testCase6() {
    input := []distribute.InputProduct{
        {
            Product: "YourProduct",
            Qty:     5,
            Premiums: []distribute.InputPremium{
                {PromotionCode: 999, QtyUse: 2.5},
                {PromotionCode: 888, QtyUse: 2.5},
            },
        },
    }

    output := distribute.DistributePremiums(input)
    printResults(input, output)
}
```

### การปรับแต่ง Algorithm

คุณสามารถปรับแต่ง logic ใน `pkg/distribute/distribute.go` ได้ตามความต้องการ:

- เปลี่ยนวิธีการเรียงลำดับโปรโมชั่น
- ปรับวิธีการจัดสรรส่วนที่เหลือ
- เพิ่มเงื่อนไขพิเศษสำหรับการจัดสรร

## 📝 หมายเหตุสำคัญ

1. **การปัดเศษ**: ใช้ `math.Round()` เพื่อปัดเศษให้เป็นจำนวนเต็ม
2. **การจัดสรร**: ใช้ algorithm แบบ greedy เพื่อให้ได้ผลลัพธ์ที่ดีที่สุด
3. **การตรวจสอบ**: ระบบจะตรวจสอบความถูกต้องของผลลัพธ์อัตโนมัติ
4. **ประสิทธิภาพ**: Algorithm มีความซับซ้อน O(n log n) สำหรับการเรียงลำดับ

---

**Version**: 1.0.0  
**Last Updated**: 2024  
**Language**: Go 1.23.5 