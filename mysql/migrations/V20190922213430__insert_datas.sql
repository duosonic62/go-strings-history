-- +migrate Up
-- Insert Member Information
INSERT INTO member (id, uid, name, token, token_expired, is_deleted, version, created_at, updated_at) VALUES ('4E8F28EDA9C64673924A9C5966D082C8', 'C8E184B0DB5B45F5B26C337ADFE0BF9D', 'Lowell', 'F24BF136E0124199800CA25CB216891B', (now() + interval 1 day), false, 1, now(), now());
INSERT INTO member (id, uid, name, token, token_expired, is_deleted, version, created_at, updated_at)VALUES ('0DECD4D5EB29456EBA770DBA60A96A1E', 'D349036A46B84EE4943FEF16F85637FD', 'Bonny', '53B7AEEFBD2E41B68B95D54254D51150', (now() + interval 1 day), false, 1, now(), now());

-- Insert String Information
INSERT INTO guitar_string (id, name, description, maker, thin_gauge, thick_gauge, url, is_deleted, version, created_at, updated_at) VALUES ('4E8F28EDA9C64673924A9C5966D082C8', 'EHR370', 'SEMI-FLAT WOUND', 'D\'Addario', 11, 49,'https://www.soundhouse.co.jp/products/detail/item/26536/', false, 1, now(), now());
INSERT INTO guitar_string (id, name, description, maker, thin_gauge, thick_gauge, url, is_deleted, version, created_at, updated_at) VALUES ('4E8F28EDA9C64673924A9C5966D082C9', 'Eatrhwood', 'BRONZE WOUND', 'ERNIEBALL', 12, 54, 'https://www.soundhouse.co.jp/products/detail/item/32523/', false, 1, now(), now());

-- Insert Guitar Information
INSERT INTO guitar (id, member_id, name, description, body_type, maker, is_deleted, version, created_at, updated_at) VALUES ('A4A4EDF948D241C38088C2D4BD7A4AA1', '4E8F28EDA9C64673924A9C5966D082C8', 'Stratocaster 72', 'white body', 'STRATOCASTER', 'Fender', false, 1, now(), now());
