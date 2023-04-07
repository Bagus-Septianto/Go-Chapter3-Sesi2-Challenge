# Registrasi
endpoint : `localhost:8080/users/register`, method : `POST`
1.  - input : `{"username": "u1", "password": "12345678"}`
    - response : `{
  "id": 5,
  "role": "",
  "username": "u1"
}`
---
2.  - input : `{"username": "u1", "password": "12345678"}`
    - response : `{
  "error": "Bad Request",
  "message": "duplicated key not allowed"
}`
---
3.  - input : `{"username": "u1", "password": "12345678", "role": "admin"}`
    - response : `{
  "error": "Bad Request",
  "message": "duplicated key not allowed"
}`
---
4.  - input : `{"username": "admin", "password": "12345678", "role": "admin"}`
    - response : `{
  "id": 8,
  "role": "admin",
  "username": "admin"
}`
---
##### Note (Registrasi) : tiap request baik gagal (Bad Request) maupun berhasil (Created) akan meng*increment* id user.
<br>

# Login
endpoint : `localhost:8080/users/login`, method : `POST`
1.  - input : `{"username": "u1", "password": "12345678"}`
    - response : `{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwicm9sZSI6IiIsInVzZXJuYW1lIjoidTEifQ.Hn17TEbivlECPNkC-GW-Bi8Q3IFt6LApWmAG7Ed07eo"
}`
---
2.  - input : `{"username": "u1u", "password": "12345678"}`
    - response : `{
  "error": "Unauthorized",
  "message": "invaild username/password"
}`
---
3.  - input : `{"username": "u1", "password": "12345678x"}`
    - response : `{
  "error": "Unauthorized",
  "message": "invaild username/password"
}`
---
4.  - input : `{"username": "u1u", "password": "12345678x"}`
    - response : `{
  "error": "Unauthorized",
  "message": "invaild username/password"
}`
---
5.  - input : `{"username": "admin", "password": "12345678"}`
    - response : `{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OCwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJhZG1pbiJ9.kd4_-IssetDYXZXaOfkAYmr2OlE43VcP5B8Xeou10hI"
}`
---
##### Note : kelupaan status code, males test ulang karena manual. jadi ya bagian register dan login tidak ada status code.
<br>

# Create Product
endpoint : `localhost:8080/products`, method : `POST`
1.  - input : `{"title": "title", "description": "desc"}`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwicm9sZSI6IiIsInVzZXJuYW1lIjoidTEifQ.Hn17TEbivlECPNkC-GW-Bi8Q3IFt6LApWmAG7Ed07eo` (non-admin)
    - status : 201 (Created)
    - response : `{
  "id": 4,
  "created_at": "2023-04-07T18:00:13.6213561+07:00",
  "updated_at": "2023-04-07T18:00:13.6213561+07:00",
  "title": "title",
  "description": "desc",
  "UserID": 5,
  "User": null
}`
---
2.  - input : `{"title": "title", "description": "desc"}`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OCwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJhZG1pbiJ9.kd4_-IssetDYXZXaOfkAYmr2OlE43VcP5B8Xeou10hI` (admin)
    - status : 201 (Created)
    - response : `{
  "id": 5,
  "created_at": "2023-04-07T18:01:45.3864517+07:00",
  "updated_at": "2023-04-07T18:01:45.3864517+07:00",
  "title": "title",
  "description": "desc",
  "UserID": 8,
  "User": null
}`
---
3.  - input : `{"title": "titlex"}`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwicm9sZSI6IiIsInVzZXJuYW1lIjoidTEifQ.Hn17TEbivlECPNkC-GW-Bi8Q3IFt6LApWmAG7Ed07eo` (non-admin)
    - status : 400 (Bad Request)
    - response : `{
  "error": "Bad Request",
  "message": "Description of your product is required"
}`
---
4.  - input : `{"description": "descx"}`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwicm9sZSI6IiIsInVzZXJuYW1lIjoidTEifQ.Hn17TEbivlECPNkC-GW-Bi8Q3IFt6LApWmAG7Ed07eo` (non-admin)
    - status : 400 (Bad Request)
    - response : `{
  "error": "Bad Request",
  "message": "Title of your product is required"
}`
---
##### Note : at this point i'm lazy. realizing the possibility are endless
<br>

# Read Product
method : `GET`
1.  - endpoint : `localhost:8080/products/4`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwicm9sZSI6IiIsInVzZXJuYW1lIjoidTEifQ.Hn17TEbivlECPNkC-GW-Bi8Q3IFt6LApWmAG7Ed07eo` (non-admin)
    - status : 200 (OK)
    - response : `{
  "id": 4,
  "created_at": "2023-04-07T18:00:13.621356+07:00",
  "updated_at": "2023-04-07T18:00:13.621356+07:00",
  "title": "title",
  "description": "desc",
  "UserID": 5,
  "User": null
}`
---
2.  - endpoint : `localhost:8080/products/5`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwicm9sZSI6IiIsInVzZXJuYW1lIjoidTEifQ.Hn17TEbivlECPNkC-GW-Bi8Q3IFt6LApWmAG7Ed07eo` (non-admin)
    - status : 401 (Unauthorized)
    - response : `{
  "error": "Unauthorized",
  "message": "you are not allowed to access this data"
}`
---
3.  - endpoint : `localhost:8080/products/4`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OCwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJhZG1pbiJ9.kd4_-IssetDYXZXaOfkAYmr2OlE43VcP5B8Xeou10hI` (admin)
    - status : 200 (OK)
    - response : `{
  "id": 4,
  "created_at": "2023-04-07T18:00:13.621356+07:00",
  "updated_at": "2023-04-07T18:00:13.621356+07:00",
  "title": "title",
  "description": "desc",
  "UserID": 5,
  "User": null
}`
---
4.  - endpoint : `localhost:8080/products/5`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OCwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJhZG1pbiJ9.kd4_-IssetDYXZXaOfkAYmr2OlE43VcP5B8Xeou10hI` (admin)
    - status : 200 (OK)
    - response : `{
  "id": 5,
  "created_at": "2023-04-07T18:01:45.386451+07:00",
  "updated_at": "2023-04-07T18:01:45.386451+07:00",
  "title": "title",
  "description": "desc",
  "UserID": 8,
  "User": null
}`
---
<br>

# Update Product
method : `PUT`
1.  - input : `{"title": "title Update", "description": "desc Update"}`
    - endpoint : `localhost:8080/products/4`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwicm9sZSI6IiIsInVzZXJuYW1lIjoidTEifQ.Hn17TEbivlECPNkC-GW-Bi8Q3IFt6LApWmAG7Ed07eo` (non-admin)
    - status : 401 (Unauthorized)
    - response : `{
  "error": "Unauthorized",
  "message": "you are not allowed to access this data"
}`
---
2.  - input : `{"title": "title Update", "description": "desc Update"}`
    - endpoint : `localhost:8080/products/5`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwicm9sZSI6IiIsInVzZXJuYW1lIjoidTEifQ.Hn17TEbivlECPNkC-GW-Bi8Q3IFt6LApWmAG7Ed07eo` (non-admin)
    - status : 401 (Unauthorized)
    - response : `{
  "error": "Unauthorized",
  "message": "you are not allowed to access this data"
}`
---
3.  - input : `{"title": "title Update", "description": "desc Update"}`
    - endpoint : `localhost:8080/products/4`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OCwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJhZG1pbiJ9.kd4_-IssetDYXZXaOfkAYmr2OlE43VcP5B8Xeou10hI` (admin)
    - status : 200 (OK)
    - response : `{
  "id": 4,
  "created_at": null,
  "updated_at": "2023-04-07T18:20:08.4577956+07:00",
  "title": "title Update",
  "description": "desc Update",
  "UserID": 8,
  "User": null
}`
---
4.  - input : `{"title": "title Updatex", "description": "desc Updatex"}`
    - endpoint : `localhost:8080/products/5`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OCwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJhZG1pbiJ9.kd4_-IssetDYXZXaOfkAYmr2OlE43VcP5B8Xeou10hI` (admin)
    - status : 200 (OK)
    - response : `{
  "id": 5,
  "created_at": null,
  "updated_at": "2023-04-07T18:20:52.8215401+07:00",
  "title": "title Updatex",
  "description": "desc Updatex",
  "UserID": 8,
  "User": null
}`
---
<br>

# Delete Product
method : `DELETE`
1.  - endpoint : `localhost:8080/products/4`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwicm9sZSI6IiIsInVzZXJuYW1lIjoidTEifQ.Hn17TEbivlECPNkC-GW-Bi8Q3IFt6LApWmAG7Ed07eo` (non-admin)
    - status : 401 (Unauthorized)
    - response : `{
  "error": "Unauthorized",
  "message": "you are not allowed to access this data"
}`
---
2.  - endpoint : `localhost:8080/products/5`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwicm9sZSI6IiIsInVzZXJuYW1lIjoidTEifQ.Hn17TEbivlECPNkC-GW-Bi8Q3IFt6LApWmAG7Ed07eo` (non-admin)
    - status : 401 (Unauthorized)
    - response : `{
  "error": "Unauthorized",
  "message": "you are not allowed to access this data"
}`
---
3.  - endpoint : `localhost:8080/products/4`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OCwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJhZG1pbiJ9.kd4_-IssetDYXZXaOfkAYmr2OlE43VcP5B8Xeou10hI` (admin)
    - status : 204 (No Content)
---
4.  - endpoint : `localhost:8080/products/5`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OCwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJhZG1pbiJ9.kd4_-IssetDYXZXaOfkAYmr2OlE43VcP5B8Xeou10hI` (admin)
    - status : 204 (No Content)
---
<br>

# Read Product setelah Delete
method : `GET`
1.  - endpoint : `localhost:8080/products/4`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OCwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJhZG1pbiJ9.kd4_-IssetDYXZXaOfkAYmr2OlE43VcP5B8Xeou10hI` (admin)
    - status : 404 (Not Found)
    - response : `{
  "error": "Data Not Found",
  "message": "data doesn't exist"
}`
---
2.  - endpoint : `localhost:8080/products/5`
    - auth : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OCwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJhZG1pbiJ9.kd4_-IssetDYXZXaOfkAYmr2OlE43VcP5B8Xeou10hI` (admin)
    - status : 404 (Not Found)
    - response : `{
  "error": "Data Not Found",
  "message": "data doesn't exist"
}`
---
## Checklist
1. fungsi registrasi : sesuai ekspektasi
2. fungsi login : sesuai ekspektasi
3. fungsi Create Product non-admin : sesuai ekspektasi
4. fungsi Create Product admin : sesuai ekspektasi
5. fungsi Read Product non-admin : sesuai ekspektasi (hanya bisa read product dengan userID yang sama dengan user yang terautentikasi)
6. fungsi Read Product admin : sesuai ekspektasi (bisa read product apa saja)
7. fungsi Update Product non-admin : sesuai ekspektasi (Unauthorized)
8. fungsi Update Product admin : sesuai ekspektasi (bisa update product)
9. fungsi Delete Product non-admin : sesuai ekspektasi (Unauthorized)
10. fungsi Delete Product admin : sesuai ekspektasi (bisa delete product) 