A. Web App 需求情境展示：

此章節描述你需要實現的 Web App 規格，由多項使用案例組成，而每一項使用案例中會描述其流程和對應的 HTTP API Document [附錄-4.2]。

A0. 情境摘要

你要開發一個簡易的會員系統，使用者透過在會員系統上輸入個人資料、郵件和密碼來註冊會員身份，註冊完會員身份之後，則可以透過輸入郵件帳號和密碼來登入，並進行額外會員專屬的操作。

A1. 會員註冊 (User registration)

流程：
使用者提供以下資料來註冊 (Register) 會員身份 [R1]：
郵件 (Email)：必須為郵件格式，字數限制 4~32。不能與已註冊的會員所擁有的郵件重複。
名稱 (Name)：字數限制 5~32。
密碼 (Password)：字數限制 5~32。
後端驗證註冊資料：
在以下兩種情況下，此請求會被視為不合法的請求。後端根據不同情況做出不同回應：
郵件已存在 (Email Duplication) [R3]：該筆註冊郵件與現存會員的郵件重複。
資料格式不正確（超出字數限制，或是註冊的郵件不遵守郵件格式）[R4]
如果請求合法則成功註冊。則首先先為這筆新的會員資料創建其會員編號，並回傳成功註冊的會員資料 [R2]：
會員編號 (Id)：該會員的流水號，從 1 開始算起，如果此系統中目前存在 3 位會員，則在此時註冊會員的話會獲得會員編號 4。
郵件和名稱則為使用者在註冊時提供的郵件和名稱。
API Document
使用者註冊 [R1] — HTTP Request
HTTP Headers:

content-type: application/json
HTTP Method: POST

HTTP Path: /api/users

Request Body:

{
    "email": "會員郵件",
    "name": "會員名稱",
    "password": "會員密碼",
}
成功註冊 [R2] — HTTP Response

HTTP Status Code: 200

HTTP Headers:

content-type: application/json
content-encoding: UTF-8
Response Body:

{
    "id": "會員編號",
    "email": "會員郵件",
    "name": "會員名稱",
}
註冊失敗案例 1（郵件已存在）[R3]：

HTTP Status Code: 400

HTTP Headers:

content-type: plain/text
content-encoding: UTF-8
Request Body:

Duplicate email
註冊失敗案例 2（資料格式不正確）[R4]：

HTTP Status Code: 400

HTTP Headers:

content-type: plain/text
content-encoding: UTF-8
Response Body:

Registration's format incorrect.
A2. 會員登入 (User Login)

流程：
使用者提供以下資料來登入 (Login) [R1]：
郵件 (Email)：必須為郵件格式，字數限制 4~32。
密碼 (Password)：字數限制 5~32。
後端驗證登入資料：
在以下兩種情況下，此請求會被視為不合法的請求。後端根據不同情況做出不同回應：
找不到相同郵件和密碼的會員 (Credentials Invalid) [R3]
資料格式不正確（超出字數限制，或是登入的郵件不遵守郵件格式）[R4]
若能找出相同郵件和密碼的會員，則判定為登入成功且回傳該會員資料 [R2]：
會員編號 (Id)：該會員在註冊時便決定好的編號。
郵件和名稱則為使用者在註冊時提供的郵件和名稱。
身份權杖 (Token)：後端隨機生出一組 36~60 字數的亂碼字串（必須保證在不同時間點會產生出絕對不同且不會重複的字串，可使用 UUID 相關套件），此份亂碼字串作為該會員的身份權杖，未來每當使用者在送出 HTTP 請求時，在標頭中帶上這個權杖，就能透過權杖向後端展現其會員的身份，並享有該會員的特殊權限。
API Document
使用者登入 [R1] — HTTP Request
HTTP Headers:

content-type: application/json
HTTP Method: POST

HTTP Path: /api/users/login

Request Body:

{
    "email": "會員郵件",
    "password": "會員密碼",
}
成功登入 [R2] — HTTP Response

HTTP Status Code: 200

HTTP Headers:

content-type: application/json
content-encoding: UTF-8
Response Body:

{
    "id": "會員編號",
    "email": "會員郵件",
    "name": "會員名稱",
    "token": "67cbbe1b-0c10-46e4-b4a8-9e8505c2453b" // 會員的權杖範例
}
登入失敗案例 1（找不到相同郵件和密碼的會員）[R3]：

HTTP Status Code: 400

HTTP Headers:

content-type: plain/text
content-encoding: UTF-8
Request Body:

Credentials Invalid
登入失敗案例 2（資料格式不正確）[R4]：

HTTP Status Code: 400

HTTP Headers:

content-type: plain/text
content-encoding: UTF-8
Response Body:

Login's format incorrect.
A3. 會員修改名稱 (User Rename)

流程：
使用者欲提供以下資料來修改會員名稱 [R1]：
名稱 (Name)：字數限制 5~32。
後端首先做身份驗證，解析請求標頭中的身份權杖：
如果此為權杖的格式不正確，導致後端無法正確解析，則視為身份驗證失敗，回傳 [R2]。
如果後端能成功解析此身份權杖所展現的會員身份，但此請求提出修改的會員對象為其他用戶，由於你無權修改其他會員的名稱，回傳 [R3]。
後端驗證修改名稱資料：
如果名稱的長度不正確，則回傳 [R5]
找到此會員，並成功修改其名稱，回傳 [R4]。
API Document
使用者欲修改名稱 [R1] — HTTP Request
HTTP Headers:

content-type: application/json
身份權杖 (token) 透過 header 攜帶：
Authorization: Bearer <token> （欄位名稱為 Authorization，值的部分首先先填入 Bearer，接著隔一個空白，再填入身份權杖的字串）
HTTP Method: PATCH

HTTP Path: /api/users/{userId}

將該會員的編號作為路徑變數 (Path Variable)，填入路徑中的 {userId} 段落。如果此使用者的會員編號為 3，則修改名稱請求的路徑為：/api/users/3。
Request Body:

{
    "newName": "欲修改的新名稱"
}
此為不合法的身份權杖（權杖格式不正確） [R2] — HTTP Response

HTTP Status Code: 401

HTTP Headers:

content-type: plain/text
content-encoding: UTF-8
Response Body:

Can't authenticate who you are.
無權修改其他人的名稱 [R3] — HTTP Response

HTTP Status Code: 403

HTTP Headers:

content-type: plain/text
content-encoding: UTF-8
Response Body:

Forbidden
成功修改名稱 [R4] — HTTP Response

HTTP Status Code: 204
HTTP Headers 中不存在 content-type 和 content-encoding ，並且也沒有 Response Body。
名稱的長度不正確 [R5] — HTTP Response

HTTP Status Code: 400

HTTP Headers:

content-type: plain/text
content-encoding: UTF-8
Response Body:

Name's format invalid.
A4. 用名稱關鍵字來查詢會員 (User Query)

流程
使用者欲查詢會員 [R1]，提供以下資料：
查詢關鍵字 (keyword)：為一字串，非必要參數
後端驗證此請求者是否來自於一「會員」，意即此請求須攜帶合法的會員權杖，否則無權查詢，中斷操作 [R2]。
後端回傳篩選後的會員列表作為查詢結果 [R3]
如果使用者沒有提供此查詢關鍵字，則回傳所有會員。
如果使用者有提供此查詢關鍵字，則回傳所有會員中，名稱包含此關鍵字的會員。
API Document
使用者欲查詢會員 [R1] — HTTP Request
HTTP Headers:

身份權杖 (token) 透過 header 攜帶：
Authorization: Bearer <token> （欄位名稱為 Authorization，值的部分首先先填入 Bearer，接著隔一個空白，再填入身份權杖的字串）
HTTP Method: GET

HTTP Path: /api/users

HTTP 查詢變數 (Query Parameter)：keyword 為一字串

舉例來說，如果此使用者欲查詢的關鍵字為 johnny ，則此請求的路徑為/api/users?keyword=johnny
Request Body: 無需提供

此為不合法的身份權杖（權杖格式不正確） [R2] — HTTP Response

HTTP Status Code: 401

HTTP Headers:

content-type: plain/text
content-encoding: UTF-8
Response Body:

Can't authenticate who you are.
回傳篩選後的會員列表作為查詢結果 [R3] — HTTP Response

HTTP Status Code: 200

HTTP Headers:

content-type: application/json
content-encoding: UTF-8
Response Body:

[
    {
        "id": "第 0 位會員的編號",
        "email": "第 0 位會員的郵件",
        "name": "第 0 位會員的姓名",
    },
    {
        "id": "第 1 位會員的編號",
        "email": "第 1 位會員的郵件",
        "name": "第 1 位會員的姓名",
    },
... // 以此類推，列入每一位會員在 JSON 陣列中
]
在初版需求中，僅需要以此四項使用案例作為雛形即可，你可以自由添加任何你認為有助於你展現雛形的使用案例。
