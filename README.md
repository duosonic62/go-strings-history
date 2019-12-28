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

```sql 
drop table gorp_migrations;
drop table string_exchange_log;
drop table guitar_string;
drop table guitar;
drop table member;
```


## コード自動生成
### sqlboiler
モデルを自動生成
```
sqlboiler mysql
```

### mockgen
[gomock](https://github.com/golang/mock)を自動生成

```
mockgen \
    -source=[source file] \
    -destination=[destination file]
```

sample
```
$ mockgen \
    -source=internal/domain/factory/user.go \
    -destination=internal/domain/factory/mock_factory/mock_user.go
```