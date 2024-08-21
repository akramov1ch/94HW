# Uy vazifasi: RESTful API bilan `expvar` va ma'lumotlar bazasi integratsiyasi yordamida metrikalarni boshqarish

## Maqsad
Kutubxonadagi kitoblarni boshqarish uchun `RESTful API` yarating, `SQL` ma'lumotlar bazasi (masalan `PostgreSQL`) bilan integratsiyalang va `expvar` paketidan foydalanib, ilova metrikalarini yig'ing va ularga kirishni ta'minlang.

## Talablar
1. **RESTful API endpoints:**
    - `GET /books`: Barcha kitoblarni qaytarish.
    - `POST /books`: Yangi kitob qo'shish.
    - `GET /books/{id}`: Ma'lum bir ID bo'yicha kitobni qaytarish.
    - `PUT /books/{id}`: Ma'lum bir ID bo'yicha kitobni yangilash.
    - `DELETE /books/{id}`: Ma'lum bir ID bo'yicha kitobni o'chirish.

2. **Ma'lumotlar bazasi operatsiyalari (PostgreSQL):**
    - Kitoblar uchun `CRUD` operatsiyalarni amalga oshirish (yaratish, o'qish, yangilash, o'chirish).
    - Jadval sxemasi:
        - `id`: Asosiy kalit.
        - `title`: Kitobning nomi.
        - `author`: Muallifning ismi.
        - `published_date`: Nashr sanasi.
        - `isbn`: ISBN raqami.
    - Ma'lumotlar bazasi migratsiyalarini qo'shing.
    
3. **Metrikalar bilan ishlash (expvar):**
    - Umumiy so'rovlar sonini kuzatib boring (`RequestCount`).
    - Muvaffaqiyatli javoblar sonini sanang (`SuccessCount`).
    - Xatolarni kuzating (`ErrorCount`). 
    - Ma'lumotlar bazasiga oid qo'shimcha metrikalar qo'shing.


 















