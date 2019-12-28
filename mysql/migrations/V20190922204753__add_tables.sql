-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE member
(
    id            VARCHAR(36) NOT NULL comment 'メンバーid',
    uid           VARCHAR(36) NOT NULL comment '外部認証紐付け用uid',
    name          VARCHAR(32) NOT NULL comment 'メンバー名',
    token         VARCHAR(36) NOT NULL comment 'アクセストークン',
    token_expired DATETIME    NOT NULL comment 'トークン有効期限',
    is_deleted    BOOLEAN     NOT NULL comment '削除フラグ',
    version       INT         NOT NULL comment 'バージョン',
    created_at    DATETIME    NOT NULL comment '登録日時',
    updated_at    DATETIME comment '更新日時',
    PRIMARY KEY (id)
) comment ='メンバー情報'
    /*! engine = InnoDb */;
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE TABLE guitar
(
    id          VARCHAR(36)  NOT NULL comment 'ギターid',
    member_id   VARCHAR(36)  NOT NULL comment 'メンバーid(外部キー)',
    name        VARCHAR(256) NOT NULL comment 'ギター名',
    description VARCHAR(2048) comment '説明',
    body_type   VARCHAR(256) comment 'ギター機種',
    maker       VARCHAR(256) comment 'メーカー',
    image_url   VARCHAR(1024) comment 'ギター画像',
    is_deleted  BOOLEAN      NOT NULL comment '削除フラグ',
    version     INT          NOT NULL comment 'バージョン',
    created_at  DATETIME     NOT NULL comment '登録日時',
    updated_at  DATETIME comment '更新日時',
    PRIMARY KEY (id),
    CONSTRAINT fk_guitar_member_id
        FOREIGN KEY (member_id)
            REFERENCES member (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
) comment ='ギター情報'
    /*! engine = InnoDb */;
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE TABLE guitar_string
(
    id          VARCHAR(36)  NOT NULL comment '弦id',
    name        VARCHAR(256) NOT NULL comment '弦名',
    description VARCHAR(2048) comment '説明',
    maker       VARCHAR(256) comment 'メーカー',
    thin_gauge  TINYINT UNSIGNED comment '細い弦のゲージ',
    thick_gauge TINYINT UNSIGNED comment '太い弦のゲージ',
    url         VARCHAR(1024) comment '商品URL',
    is_deleted  BOOLEAN      NOT NULL comment '削除フラグ',
    version     INT          NOT NULL comment 'バージョン',
    created_at  DATETIME     NOT NULL comment '登録日時',
    updated_at  DATETIME comment '更新日時',
    PRIMARY KEY (id)
) comment ='弦情報'
    /*! engine = InnoDb */;
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE TABLE string_exchange_log
(
    id           VARCHAR(36) NOT NULL comment '弦交換id',
    string_id    VARCHAR(36) NOT NULL comment '弦id',
    guitar_id    VARCHAR(36) NOT NULL comment 'ギターid',
    is_exchanged BOOLEAN     NOT NULL comment '交換フラグ',
    is_deleted   BOOLEAN     NOT NULL comment '削除フラグ',
    version      INT         NOT NULL comment 'バージョン',
    created_at   DATETIME    NOT NULL comment '登録日時',
    updated_at   DATETIME comment '更新日時',
    PRIMARY KEY (id),
    CONSTRAINT fk_exchange_string_id
        FOREIGN KEY (string_id)
            REFERENCES guitar_string (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE,
    CONSTRAINT fk_exchange_guitar_id
        FOREIGN KEY (guitar_id)
            REFERENCES guitar (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
) comment ='弦交換情報'
    /*! engine = InnoDb */;
-- +migrate StatementEnd
