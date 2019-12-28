# strings history app
ギター弦の管理用アプリのサーバサイド

## セットアップ
### mysql
```
docker run --name strings-mysql \
 -e MYSQL_ROOT_PASSWORD=pass \
 -e MYSQL_DATABASE=strings_history \
 -e MYSQL_ROOT_HOST='%' -e TZ='Asia/Tokyo' \
 -p 3306:3306 \
 mysql/mysql-server:5.7 \
 --character-set-server=utf8mb4 \ 
 --collation-server=utf8mb4_general_ci
```
