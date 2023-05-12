
# Code Competence Alterra

## API Reference

| URI | Method     | Description                |
| :-------- | :------- | :------------------------- |
| `/users/register` | `POST` | Mendaftarkan akun pengguna |
| `/users/login` | `POST` | Mendapatkan token saat pengguna masuk |
| `/items` | `GET` | Mendapatkan semua data barang |
| `/items/:id` | `GET` | Mendapatkan data barang berdasarkan ID |
| `/items` | `POST` | Menambahkan data barang |
| `/items/:id` | `PUT` | Mengubah data barang |
| `/items/:id` | `DELETE` | Menghapus data barang |
| `/items/category/:category_id` | `GET` | Mendapatkan semua data barang berdasarkan ID Kategori |
| `/items?keyword=item_name` | `GET` | Mendapatkan data barang berdasarkan nama barang |

---

#### Mendaftarkan akun pengguna

```http
  POST /users/register
```

#####

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/13.%20Register.png?raw=true)

#### Jika email telah digunakan

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/12.%20Register%20if%20email%20already%20exist.png?raw=true?raw=true)

---

#### Mendapatkan token saat pengguna masuk

```http
  POST /users/login
```

#####

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/14.%20Login.png?raw=true)

---

#### Mendapatkan semua data barang

```http
  GET /items
```

#####

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/08.%20Get%20All%20Items.png?raw=true)

#### Jika tidak memiliki akses token

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/08a.%20Get%20All%20Items%20if%20not%20authorized.png?raw=true)

---

#### Mendapatkan data barang berdasarkan ID

```http
  GET /items/:id
```

#####

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/07.%20Get%20Item%20By%20ID.png?raw=true)

#### Jika tidak memiliki akses token

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/07a.%20Get%20Item%20By%20ID%20if%20not%20authorized.png?raw=true)

---

#### Menambahkan data barang

```http
  POST /items
```

#####

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/05.%20Add%20Item.png?raw=true)

#### Jika tidak memiliki akses token

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/05a.%20Add%20Item%20if%20not%20authorized.png?raw=true)

---

#### Mengubah data barang

```http
  PUT /items/:id
```

#####

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/06.%20Update%20Item.png?raw=true)

#### Jika tidak memiliki akses token

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/06a.%20Update%20Item%20if%20not%20authorized.png?raw=true)

---

#### Menghapus data barang

```http
  DELETE /items/:id
```

#####

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/09.%20Delete%20Item.png?raw=true)

#### Jika tidak memiliki akses token

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/09a.%20Delete%20Item%20if%20not%20authorized.png?raw=true)

---

#### Mendapatkan semua data barang berdasarkan ID Kategori

```http
  GET /items/category/:category_id
```

#####

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/10.%20Get%20Item%20By%20Category%20ID.png?raw=true)

#### Jika tidak memiliki akses token

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/10a.%20Get%20Item%20By%20Category%20ID%20if%20not%20authoirzed.png?raw=true)

---

#### Mendapatkan data barang berdasarkan nama barang

```http
  GET /items?keyword=item_name
```

#####

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/11.%20Get%20Item%20by%20Name%20Query%20Param.png?raw=true)

#### Jika tidak memiliki akses token

![App Screenshot](https://github.com/nifz/go_mochammad-hanif/blob/main/code-competence/screenshot/11a.%20Get%20Item%20by%20Name%20Query%20Param%20if%20not%20authorized.png?raw=true)
