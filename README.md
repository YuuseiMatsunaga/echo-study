### 環境構築
1. envを作成してください
```
copy .env.example .env
```
2. コンテナを立ち上げてください
```
docker-compose up
```
3. migrationしてください
```
docker-compose exec api sql-migrate up
```

[ここにアクセスしてみてhello worldが出てたら完了です](http://localhost:8080)