# 📦 Inventory Management API
Welcome to the Inventory Management API! This is a simple and efficient RESTful API built using Golang, designed to manage users, items (Barang), and inventory mutations (Mutasi). This API uses PostgreSQL as its database, making it robust and reliable.

## 🎯 Features
- User Management: Create, read, update, and delete users.
- Item Management: Handle inventory items with CRUD operations.
- Mutation Tracking: Track mutations (additions, removals, transfers) of inventory items.
- Authentication: Secure login with JWT-based token authentication.
- Relation Mapping: Proper mapping of relationships between Users, Items, and Mutations.
- Error Handling: Friendly and clear error messages for a smooth developer experience.

## 📚 Daftar Library
Proyek ini menggunakan berbagai library Python untuk mengotomatiskan proses login:

- requests: Mendownload gambar CAPTCHA.
- Pillow (PIL): Manipulasi gambar.
- pytesseract: Optical Character Recognition (OCR) untuk menyelesaikan CAPTCHA.
- selenium: Automatisasi interaksi dengan browser.
- opencv-python: Pemrosesan dan preprocessing gambar untuk meningkatkan akurasi OCR.
- numpy: Manipulasi array gambar.

## 🛠️ Getting Started
# Prerequisites
- Golang: Make sure you have Golang installed. Download Golang
1. Clone the Repository:
```bash
git clone https://github.com/yourusername/inventory-management-api.git
cd inventory-management-api
```

2. Install Dependencies:
```bash
go mod tidy
```

3. Environment Setup:
Create a .env file in the root directory and add your configuration:


```bash
DSN=root:@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local

```
4. Run the Server:

```bash
go run main.go
```
## 🎯 Postman Documentation
https://documenter.getpostman.com/view/31882906/2sAXjNXAtY