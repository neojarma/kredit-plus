/*
 Navicat Premium Data Transfer

 Source Server         : mysql-docker
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : localhost:3306
 Source Schema         : XYZ_multi_finance

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 29/08/2024 15:31:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for assets
-- ----------------------------
DROP TABLE IF EXISTS `assets`;
CREATE TABLE `assets`  (
  `id_assets` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `nama_asset` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `type_asset` enum('WHITE_GOODS','MOBILE','MOTOR') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `otr_amount` double NULL DEFAULT NULL,
  `bunga_persen` float NULL DEFAULT NULL,
  `admin_fee` double NULL DEFAULT NULL,
  PRIMARY KEY (`id_assets`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of assets
-- ----------------------------
INSERT INTO `assets` VALUES ('HHRySZn1Md6UOz9', 'Mesin Cuci', 'WHITE_GOODS', 1500000, 2, 80000);
INSERT INTO `assets` VALUES ('WX0yrXGiPQODshF', 'Kulkas', 'WHITE_GOODS', 1300000, 2, 70000);

-- ----------------------------
-- Table structure for auths
-- ----------------------------
DROP TABLE IF EXISTS `auths`;
CREATE TABLE `auths`  (
  `id_auth` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `id_user` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of auths
-- ----------------------------
INSERT INTO `auths` VALUES ('8isdMazahMgt6od', 'WFIdLL6qyL0BiB8', 'neojarma@mail.com', '$2a$14$xPdnbbjHyR4c8KddVlw5SudSYZdf6xpvdYNYXSdbMXfnUaYOkpYGq');

-- ----------------------------
-- Table structure for limit_tenors
-- ----------------------------
DROP TABLE IF EXISTS `limit_tenors`;
CREATE TABLE `limit_tenors`  (
  `id_limit_tenor` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `1_bulan` double NULL DEFAULT NULL,
  `2_bulan` double NULL DEFAULT NULL,
  `3_bulan` double NULL DEFAULT NULL,
  `6_bulan` double NULL DEFAULT NULL,
  `salary_min` double NULL DEFAULT NULL,
  `salary_max` double NULL DEFAULT NULL,
  PRIMARY KEY (`id_limit_tenor`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of limit_tenors
-- ----------------------------
INSERT INTO `limit_tenors` VALUES ('LCjpY7zQUgMwqC3', 1000000, 1200000, 1500000, 2000000, 2000000, 5000000);
INSERT INTO `limit_tenors` VALUES ('wvlnlBbS7lUuyp1', 100000, 200000, 500000, 700000, 0, 2000000);

-- ----------------------------
-- Table structure for peminjamans
-- ----------------------------
DROP TABLE IF EXISTS `peminjamans`;
CREATE TABLE `peminjamans`  (
  `id_peminjaman` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `id_user` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `id_assets` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `id_limit_tenor` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `total_pembayaran` double NULL DEFAULT NULL,
  `total_bunga` double NULL DEFAULT NULL,
  `status` enum('PAID','ACTIVE') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `tanggal_peminjaman` datetime NULL DEFAULT NULL,
  `tanggal_akhir_pembayaran` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `sisa_pembayaran` double NULL DEFAULT NULL,
  `lama_tenor` int NULL DEFAULT NULL,
  `otr` double NULL DEFAULT NULL,
  PRIMARY KEY (`id_peminjaman`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of peminjamans
-- ----------------------------
INSERT INTO `peminjamans` VALUES ('0rKShVISDF1lDZl', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 'dn2yMlaDhFsHnjD', 1689600, 189600, 'PAID', '2024-08-29 11:31:25', '2025-03-01 00:00:00', '2024-08-29 11:36:20', 0, 6, 1500000);
INSERT INTO `peminjamans` VALUES ('CZeLzCjxeYFs9B8', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 'dn2yMlaDhFsHnjD', 1689600, 189600, 'ACTIVE', '2024-08-29 15:23:24', '2025-03-01 00:00:00', '2024-08-29 15:24:35', -281600, 6, 1500000);
INSERT INTO `peminjamans` VALUES ('MVI2uJmRYtJd2oh', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 'dn2yMlaDhFsHnjD', 1689600, 189600, 'PAID', '2024-08-29 11:00:17', '2025-03-01 00:00:00', '2024-08-29 11:29:04', 0, 6, 1500000);

-- ----------------------------
-- Table structure for penagihans
-- ----------------------------
DROP TABLE IF EXISTS `penagihans`;
CREATE TABLE `penagihans`  (
  `id_penagihan` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `id_peminjaman` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `id_user` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `id_asset` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `nominal_pembayaran` double NULL DEFAULT NULL,
  `status` enum('UNPAID','PAID') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `tanggal_pembayaran` datetime NULL DEFAULT NULL,
  `tanggal_jatuh_tempo` datetime NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id_penagihan`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of penagihans
-- ----------------------------
INSERT INTO `penagihans` VALUES ('1HusC7zY8CSFXFi', '0rKShVISDF1lDZl', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'PAID', '2024-08-29 11:35:31', '2024-12-29 00:00:00', '2024-08-29 11:31:26', '2024-08-29 11:35:31');
INSERT INTO `penagihans` VALUES ('3xCmIPryUc6UxYu', 'MVI2uJmRYtJd2oh', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'PAID', '2024-08-29 11:01:00', '2024-12-29 00:00:00', '2024-08-29 11:00:17', '2024-08-29 11:01:00');
INSERT INTO `penagihans` VALUES ('59W2Qjse9jzpwG5', 'MVI2uJmRYtJd2oh', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'PAID', '2024-08-29 11:04:46', '2025-01-29 00:00:00', '2024-08-29 11:00:17', '2024-08-29 11:04:46');
INSERT INTO `penagihans` VALUES ('5Q79Mo11tvtikM0', '0rKShVISDF1lDZl', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'PAID', '2024-08-29 11:35:52', '2025-01-29 00:00:00', '2024-08-29 11:31:26', '2024-08-29 11:35:52');
INSERT INTO `penagihans` VALUES ('9xmUl13vgjjwJMO', 'CZeLzCjxeYFs9B8', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'PAID', '2024-08-29 15:24:35', '2025-03-01 00:00:00', '2024-08-29 15:23:25', '2024-08-29 15:24:35');
INSERT INTO `penagihans` VALUES ('ArafkFvCjLEzWjy', 'CZeLzCjxeYFs9B8', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'UNPAID', NULL, '2024-09-29 00:00:00', '2024-08-29 15:23:25', NULL);
INSERT INTO `penagihans` VALUES ('eIiOrUSX1KbSqrU', '0rKShVISDF1lDZl', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'PAID', '2024-08-29 11:36:00', '2025-03-01 00:00:00', '2024-08-29 11:31:26', '2024-08-29 11:36:00');
INSERT INTO `penagihans` VALUES ('gmXY7sQzggLBNZz', '0rKShVISDF1lDZl', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'PAID', '2024-08-29 11:36:07', '2024-10-29 00:00:00', '2024-08-29 11:31:26', '2024-08-29 11:36:07');
INSERT INTO `penagihans` VALUES ('li7yc3fzyXgJeB4', 'MVI2uJmRYtJd2oh', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'PAID', '2024-08-29 11:04:57', '2024-09-29 00:00:00', '2024-08-29 11:00:17', '2024-08-29 11:04:57');
INSERT INTO `penagihans` VALUES ('m8VqkwF04joMm6J', '0rKShVISDF1lDZl', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'PAID', '2024-08-29 11:36:15', '2024-11-29 00:00:00', '2024-08-29 11:31:26', '2024-08-29 11:36:15');
INSERT INTO `penagihans` VALUES ('OOMHGiueR90sr2g', 'CZeLzCjxeYFs9B8', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'UNPAID', NULL, '2024-10-29 00:00:00', '2024-08-29 15:23:25', NULL);
INSERT INTO `penagihans` VALUES ('PbzQpWXVgbvX8Mr', 'CZeLzCjxeYFs9B8', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'UNPAID', NULL, '2025-01-29 00:00:00', '2024-08-29 15:23:25', NULL);
INSERT INTO `penagihans` VALUES ('PY0PgXgFA2EZtnM', 'MVI2uJmRYtJd2oh', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'PAID', '2024-08-29 11:05:09', '2025-03-01 00:00:00', '2024-08-29 11:00:17', '2024-08-29 11:05:09');
INSERT INTO `penagihans` VALUES ('Q62em9yYqYbJWjo', 'CZeLzCjxeYFs9B8', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'UNPAID', NULL, '2024-12-29 00:00:00', '2024-08-29 15:23:25', NULL);
INSERT INTO `penagihans` VALUES ('qvYQiCZZ8ZQvigt', 'MVI2uJmRYtJd2oh', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'PAID', '2024-08-29 11:05:15', '2024-10-29 00:00:00', '2024-08-29 11:00:17', '2024-08-29 11:05:15');
INSERT INTO `penagihans` VALUES ('rblRGbDIQSDdw2W', '0rKShVISDF1lDZl', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'PAID', '2024-08-29 11:36:20', '2024-09-29 00:00:00', '2024-08-29 11:31:26', '2024-08-29 11:36:20');
INSERT INTO `penagihans` VALUES ('SpDlXuVMCIYXQeN', 'MVI2uJmRYtJd2oh', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'PAID', '2024-08-29 11:29:04', '2024-11-29 00:00:00', '2024-08-29 11:00:17', '2024-08-29 11:29:04');
INSERT INTO `penagihans` VALUES ('vIM7IWcQSTVzZM8', 'CZeLzCjxeYFs9B8', 'WFIdLL6qyL0BiB8', 'HHRySZn1Md6UOz9', 281600, 'UNPAID', NULL, '2024-11-29 00:00:00', '2024-08-29 15:23:25', NULL);

-- ----------------------------
-- Table structure for user_tenors
-- ----------------------------
DROP TABLE IF EXISTS `user_tenors`;
CREATE TABLE `user_tenors`  (
  `id_user_tenor` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `id_user` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `id_limit_tenor` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `1_bulan` double NULL DEFAULT NULL,
  `2_bulan` double NULL DEFAULT NULL,
  `3_bulan` double NULL DEFAULT NULL,
  `6_bulan` double NULL DEFAULT NULL,
  PRIMARY KEY (`id_user_tenor`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_tenors
-- ----------------------------
INSERT INTO `user_tenors` VALUES ('dn2yMlaDhFsHnjD', 'WFIdLL6qyL0BiB8', 'LCjpY7zQUgMwqC3', 1000000, 1200000, 1500000, 100000);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id_user` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `nik` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `full_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `legal_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `tempat_lahir` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `tanggal_lahir` date NULL DEFAULT NULL,
  `gaji` double NULL DEFAULT NULL,
  `foto_ktp_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `foto_selfie_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id_user`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('WFIdLL6qyL0BiB8', '3123123', 'neo jarma', 'neo jarma', 'bandung', '2024-08-28', 3000000, 'assets\\WFIdLL6qyL0BiB8\\KTP\\1.jpg', 'assets\\WFIdLL6qyL0BiB8\\Selfie\\1.jpg');

SET FOREIGN_KEY_CHECKS = 1;
