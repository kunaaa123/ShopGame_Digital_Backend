# 🎮 ShopGame API

API สำหรับจัดการข้อมูลเกมดิจิทัล สร้างด้วย Go พร้อมระบบจัดการข้อมูลพื้นฐาน (CRUD)

## 🌟 คุณสมบัติหลัก

- API endpoints สำหรับจัดการข้อมูลเกม
- ใช้สถาปัตยกรรมแบบ Clean Architecture
- ใช้ MySQL เป็นฐานข้อมูลร่วมกับ GORM
- ใช้ Singleton pattern สำหรับการเชื่อมต่อฐานข้อมูล
- รองรับการทำงานแบบ CRUD
- มีการทดสอบระบบครบถ้วน

## 📁 โครงสร้างโปรเจค

```
├── cmd/
│   └── main.go                 # จุดเริ่มต้นโปรแกรม
├── internal/
│   ├── adapter/
│   │   ├── handler/            # ส่วนจัดการ HTTP
│   │   └── repository/         # ส่วนจัดการฐานข้อมูล
│   ├── domain/                 # โมเดลและอินเตอร์เฟส
│   ├── service/                # ส่วนจัดการธุรกิจ
│   └── db/                     # การตั้งค่าฐานข้อมูล
└── test/                       # การทดสอบ
```

## 🛠 เทคโนโลยีที่ใช้

- **Go** - ภาษาที่ใช้พัฒนา
- **Gin** - เฟรมเวิร์ค HTTP
- **GORM** - ORM สำหรับจัดการฐานข้อมูล
- **MySQL** - ระบบฐานข้อมูล
- **Testify** - เฟรมเวิร์คสำหรับการทดสอบ

## 📋 ความต้องการของระบบ

- Go 1.16 ขึ้นไป
- MySQL 5.7 ขึ้นไป
- Git

## ⚙️ วิธีการติดตั้ง

1. โคลนโปรเจค:
```bash
git clone https://github.com/yourusername/ShopGame.git
```

2. เข้าไปที่โฟลเดอร์โปรเจค:
```bash
cd ShopGame
```

3. ติดตั้ง dependencies:
```bash
go mod download
```

4. ตั้งค่าฐานข้อมูลใน `internal/db/mysql.go`:
```go
func NewDBConfig() *DBConfig {
    return &DBConfig{
        Host:     "localhost",
        Port:     "3306",
        User:     "root",
        Password: "",
        DBName:   "gameshop",
    }
}
```

5. รันโปรแกรม:
```bash
go run cmd/main.go
```

## 🚀 API Endpoints

| Method | Endpoint | คำอธิบาย |
|--------|----------|----------|
| GET    | /games   | ดึงข้อมูลเกมทั้งหมด |
| GET    | /games/:id | ดึงข้อมูลเกมตาม ID |
| POST   | /games   | เพิ่มเกมใหม่ |
| PUT    | /games/:id | แก้ไขข้อมูลเกม |
| DELETE | /games/:id | ลบข้อมูลเกม |

### ตัวอย่าง Request Body (POST/PUT)
```json
{
    "name": "Elden Ring",
    "description": "เกม Action RPG",
    "price": 59.99,
    "image": "elden-ring.jpg"
}
```

## ✅ การทดสอบ

รันการทดสอบทั้งหมด:
```bash
go test ./test/... -v
```

## 📝 กฎการตรวจสอบข้อมูล

- ชื่อเกมห้ามว่าง
- ราคาต้องมากกว่า 0
- รูปภาพต้องมีนามสกุลถูกต้อง (.jpg หรือ .png)

## 🤝 การมีส่วนร่วม

ยินดีรับข้อเสนอแนะและการแก้ไขเพิ่มเติม

## 📜 สัญญาอนุญาต

โปรเจคนี้ใช้สัญญาอนุญาตแบบ MIT License