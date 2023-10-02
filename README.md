# tempWriteFile

## features
1. local: 把 mock.jpg 讀入, 然後內容原封不動寫出 result.jpg
2. tcp connection: 從 client 把 mock.jpg , 傳送到 server, 存成 result.jpg

## 1. Run Local
```bash
go run .
```
預期結果: 多一張 result.jpg

## 2. Run with TCP Connection
server 端
```bash
cd server && go run .
```

client 端
```bash
cd client && go run .
```
預期結果：server 資料夾下多一個 result.jpg 檔案且可正常打開