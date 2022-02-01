/*
 Navicat MySQL Data Transfer

 Source Server         : 我得
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : hemacms

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 01/02/2022 03:47:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for hm_auth_role
-- ----------------------------
DROP TABLE IF EXISTS `hm_auth_role`;
CREATE TABLE `hm_auth_role`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `rules` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` int NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `en` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_auth_role
-- ----------------------------
INSERT INTO `hm_auth_role` VALUES (1, '超级管理员', '1,3,5,18,35,36,37,42,46,6,10,11,12,13,14,15,16,4,7,19,20,21,22,23,24,25,8,26,27,28,29,30,31,32,33,34,43,47,48,44,45,2,68,69,78,70,71,74,72,73', 1, 'sdfasdf', 'Super admin');
INSERT INTO `hm_auth_role` VALUES (2, '测试', '1,3,5,18,35,36,37,42,67,46,6,10,11,12,13,14,15,16,4,7,19,20,21,22,23,24,25,8,26,27,28,29,30,31,32,33,34,43,47,49,50,48,52,53,54,66,44,45,2,40,41,17,55,60,56,61,65,62,57,58,59,63,64,68,69,78,70,71,74,72,73', 1, 'asdfasdf', 'Test');

-- ----------------------------
-- Table structure for hm_auth_role_access
-- ----------------------------
DROP TABLE IF EXISTS `hm_auth_role_access`;
CREATE TABLE `hm_auth_role_access`  (
  `uid` int UNSIGNED NOT NULL,
  `role_id` int UNSIGNED NOT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_auth_role_access
-- ----------------------------
INSERT INTO `hm_auth_role_access` VALUES (1, 1);
INSERT INTO `hm_auth_role_access` VALUES (2, 2);
INSERT INTO `hm_auth_role_access` VALUES (3, 2);

-- ----------------------------
-- Table structure for hm_auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `hm_auth_rule`;
CREATE TABLE `hm_auth_rule`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `pid` int NULL DEFAULT NULL,
  `isshow` int NULL DEFAULT NULL,
  `sort` int NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `level` int NULL DEFAULT NULL,
  `en` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 78 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_auth_rule
-- ----------------------------
INSERT INTO `hm_auth_rule` VALUES (1, 'site', '设置', 0, 1, 1, 'icon-1012333', 1, 'Site');
INSERT INTO `hm_auth_rule` VALUES (2, 'content', '内容', 0, 1, 1, 'icon-book_addresses_key', 1, 'Content');
INSERT INTO `hm_auth_rule` VALUES (3, 'adsense/site', '站长设置', 1, 1, 1, 'icon-ipod_connect', 1, 'Adsense site');
INSERT INTO `hm_auth_rule` VALUES (4, 'account/site', '用户设置', 1, 1, 2, 'icon-user_key', 1, 'Account site');
INSERT INTO `hm_auth_rule` VALUES (5, 'site/system', '系统设置', 3, 1, 1, 'icon-shape_square_key', 1, 'System site');
INSERT INTO `hm_auth_rule` VALUES (6, 'site/menu', '菜单设置', 3, 1, 2, 'icon-calendar_star', 1, 'Menu site');
INSERT INTO `hm_auth_rule` VALUES (7, 'site/user', '用户管理', 4, 1, 1, 'icon-user_orange', 1, 'Account manager');
INSERT INTO `hm_auth_rule` VALUES (8, 'site/role', '角色管理', 4, 1, 2, 'icon-group_key', 1, 'Role manager');
INSERT INTO `hm_auth_rule` VALUES (10, 'site/getMenu', '获取所有菜单', 6, 0, 1, 'icon-cabinet', 1, 'get all menus');
INSERT INTO `hm_auth_rule` VALUES (11, 'site/sortMenu', '菜单排序', 6, 0, 2, 'icon-cabinet', 1, 'menu sort');
INSERT INTO `hm_auth_rule` VALUES (12, 'site/addMenu', '添加菜单页面', 6, 0, 3, 'icon-cabinet', 1, 'Add menu page');
INSERT INTO `hm_auth_rule` VALUES (13, 'site/addMenuSubmit', '添加菜单提交', 6, 0, 4, 'icon-cabinet', 1, 'Add menu submit');
INSERT INTO `hm_auth_rule` VALUES (14, 'site/editMenu', '修改菜单页面', 6, 0, 5, 'icon-cabinet', 1, 'Edit menu page');
INSERT INTO `hm_auth_rule` VALUES (15, 'site/editMenuSubmit', '修改菜单提交', 6, 0, 6, 'icon-cabinet', 1, 'Edit menu submit');
INSERT INTO `hm_auth_rule` VALUES (16, 'site/delMenuSubmit', '删除菜单', 6, 0, 7, 'icon-cabinet', 1, 'Delete menu');
INSERT INTO `hm_auth_rule` VALUES (17, 'content/articleManage', '文章管理', 2, 1, 2, 'icon-table_key', 1, 'Article manage');
INSERT INTO `hm_auth_rule` VALUES (18, 'site/editSystem', '修改系统设置', 5, 0, 1, 'icon-cabinet', 1, 'Edit system site');
INSERT INTO `hm_auth_rule` VALUES (19, 'site/getUser', '获取所有用户', 7, 0, 1, 'icon-cabinet', 1, 'Get all accounts');
INSERT INTO `hm_auth_rule` VALUES (20, 'site/addUser', '添加用户页面', 7, 0, 2, 'icon-cabinet', 1, 'Add account  page');
INSERT INTO `hm_auth_rule` VALUES (21, 'site/addUserSubmit', '添加用户提交', 7, 0, 3, 'icon-cabinet', 1, 'Add account submit');
INSERT INTO `hm_auth_rule` VALUES (22, 'site/editUser', '修改用户页面', 7, 0, 4, 'icon-cabinet', 1, 'Edit account page');
INSERT INTO `hm_auth_rule` VALUES (23, 'site/editUserSubmit', '修改用户提交', 7, 0, 5, 'icon-cabinet', 1, 'Edit account submit');
INSERT INTO `hm_auth_rule` VALUES (24, 'site/delUser', '删除单个用户', 7, 0, 6, 'icon-cabinet', 1, 'Delete single account');
INSERT INTO `hm_auth_rule` VALUES (25, 'site/batchDelUser', '批量删除用户', 7, 0, 7, 'icon-cabinet', 1, 'Batch delete accounts');
INSERT INTO `hm_auth_rule` VALUES (26, 'site/getRole', '获取所有角色', 8, 0, 1, 'icon-cabinet', 1, 'Get all roles');
INSERT INTO `hm_auth_rule` VALUES (27, 'site/setRole', '权限设置页面', 8, 0, 2, 'icon-cabinet', 1, 'Permission settings page');
INSERT INTO `hm_auth_rule` VALUES (28, 'site/setRoleSubmit', '权限授权', 8, 0, 3, 'icon-cabinet', 1, 'Authorization of authority');
INSERT INTO `hm_auth_rule` VALUES (29, 'site/addRole', '添加角色页面', 8, 0, 4, 'icon-cabinet', 1, 'Add role page');
INSERT INTO `hm_auth_rule` VALUES (30, 'site/addRoleSubmit', '添加角色提交', 8, 0, 5, 'icon-cabinet', 1, 'Add role submit');
INSERT INTO `hm_auth_rule` VALUES (31, 'site/editRole', '修改角色页面', 8, 0, 6, 'icon-cabinet', 1, 'Edit role page');
INSERT INTO `hm_auth_rule` VALUES (32, 'site/editRoleSubmit', '修改角色提交', 8, 0, 7, 'icon-cabinet', 1, 'Edit role submit');
INSERT INTO `hm_auth_rule` VALUES (33, 'site/delRole', '删除单个角色', 8, 0, 8, 'icon-cabinet', 1, 'Delete single role');
INSERT INTO `hm_auth_rule` VALUES (34, 'site/batchDelRole', '批量删除角色', 8, 0, 9, 'icon-cabinet', 1, 'Batch delete roles');
INSERT INTO `hm_auth_rule` VALUES (35, 'index', '系统主页', 5, 0, 2, 'icon-cabinet', 1, 'System home');
INSERT INTO `hm_auth_rule` VALUES (36, 'home', '首页', 5, 0, 3, 'icon-cabinet', 1, 'Home');
INSERT INTO `hm_auth_rule` VALUES (37, 'loginOut', '登出', 5, 0, 5, 'icon-cabinet', 1, 'Sign up');
INSERT INTO `hm_auth_rule` VALUES (40, 'content/columnManage', '栏目管理', 2, 1, 1, 'icon-book', 1, 'Column manage');
INSERT INTO `hm_auth_rule` VALUES (41, 'content/columnList', '栏目列表', 40, 1, 1, 'icon-bookmark_go', 1, 'Column list');
INSERT INTO `hm_auth_rule` VALUES (42, 'tabNoAuth', 'tab权限验证', 5, 0, 6, 'icon-cabinet', 1, 'tab auth verify');
INSERT INTO `hm_auth_rule` VALUES (43, 'log', '日志管理', 1, 1, 4, 'icon-folder_table', 1, 'Log manager');
INSERT INTO `hm_auth_rule` VALUES (44, 'site/secondDevelop', '二次开发', 1, 1, 5, 'icon-telephone_edit', 1, 'Second develop');
INSERT INTO `hm_auth_rule` VALUES (45, 'site/colorSchemes', '配色方案', 44, 1, 1, 'icon-stop', 1, 'Color schemes');
INSERT INTO `hm_auth_rule` VALUES (46, 'ajaxPolling', '异步轮询', 5, 0, 8, 'icon-cabinet', 1, 'ajax polling');
INSERT INTO `hm_auth_rule` VALUES (47, 'site/loginLog', '登录日志', 43, 1, 1, 'icon-user_earth', 1, 'Login log');
INSERT INTO `hm_auth_rule` VALUES (48, 'site/oprateLog', '操作日志', 43, 1, 2, 'icon-date_edit', 1, 'Oprate log');
INSERT INTO `hm_auth_rule` VALUES (49, 'site/getLoginLog', '获取登录日志', 47, 0, 1, 'icon-cabinet', 1, 'Get login log');
INSERT INTO `hm_auth_rule` VALUES (50, 'site/delLoginLog', '删除上月登录日志', 47, 0, 2, 'icon-cabinet', 1, 'Delete last month login log');
INSERT INTO `hm_auth_rule` VALUES (52, 'site/getOprateLog', '获取操作日志', 48, 0, 1, 'icon-cabinet', 1, 'Get oprate log');
INSERT INTO `hm_auth_rule` VALUES (53, 'site/delOprateLog', '删除上月操作日志', 48, 0, 2, 'icon-cabinet', 1, 'Delete last month oprate log');
INSERT INTO `hm_auth_rule` VALUES (54, 'site/delAllOprateLog', '删除全部操作日志', 48, 0, 3, 'icon-cabinet', 1, 'Delete all oprate log');
INSERT INTO `hm_auth_rule` VALUES (55, 'content/articleList', '文章列表', 17, 1, 1, 'icon-table_column', 1, 'Article list');
INSERT INTO `hm_auth_rule` VALUES (56, 'content/tagClass', '标签分类', 17, 1, 3, 'icon-table_column_add', 1, 'Tag class');
INSERT INTO `hm_auth_rule` VALUES (57, 'content/collectionManage', '采集管理', 2, 1, 3, 'icon-transmit', 1, 'Collection manage');
INSERT INTO `hm_auth_rule` VALUES (58, 'content/collectionRules', '采集规则', 57, 1, 1, 'icon-transmit_go', 1, 'Collection rules');
INSERT INTO `hm_auth_rule` VALUES (59, 'content/collector', '采集器', 57, 1, 2, 'icon-transmit_edit', 1, 'Collector');
INSERT INTO `hm_auth_rule` VALUES (60, 'content/pendArticles', '待审核文章', 17, 1, 2, 'icon-table_delete', 1, 'Pend review articles');
INSERT INTO `hm_auth_rule` VALUES (61, 'content/commentManage', '评论管理', 17, 1, 4, 'icon-table_refresh', 1, 'Comment manage');
INSERT INTO `hm_auth_rule` VALUES (62, 'content/articleRecycled', '文章回收站', 17, 1, 6, 'icon-table_save', 1, 'Articles Recycled');
INSERT INTO `hm_auth_rule` VALUES (63, 'content/cacheManage', '缓存管理', 2, 1, 4, 'icon-chart_bar', 1, 'Cache manage');
INSERT INTO `hm_auth_rule` VALUES (64, 'content/cacheList', '缓存列表', 63, 1, 1, 'icon-chart_bar_add', 1, 'Cache list');
INSERT INTO `hm_auth_rule` VALUES (65, 'content/attachManage', '附件管理', 17, 1, 5, 'icon-page_landscape_shot', 1, 'Attach manage');
INSERT INTO `hm_auth_rule` VALUES (66, 'site/updateLog', '更新日志', 43, 1, 3, 'icon-box_world', 1, 'Update log');
INSERT INTO `hm_auth_rule` VALUES (67, 'ajaxUpdateLog', '异步获取更新日志', 5, 0, 8, 'icon-cabinet', 1, 'Ajax get update log');
INSERT INTO `hm_auth_rule` VALUES (68, 'enterprise', '企业', 0, 1, 1, 'icon-2012080404391', 1, 'Enterprise');
INSERT INTO `hm_auth_rule` VALUES (69, 'enterprise/navManage', '导航管理', 68, 1, 1, 'icon-20130406125647919_easyicon_net_16', 1, 'Nav Manage');
INSERT INTO `hm_auth_rule` VALUES (70, 'enterprise/navList', '导航列表', 69, 1, 2, 'icon-2012080404218', 1, 'Nav List');
INSERT INTO `hm_auth_rule` VALUES (71, 'enterprise/contentManage', '内容管理', 69, 1, 3, 'icon-book_addresses_edit', 1, 'content manage');
INSERT INTO `hm_auth_rule` VALUES (72, 'enterprise/site', '企业设置', 68, 1, 2, 'icon-application_tile_horizontal', 1, 'Enterprise Site');
INSERT INTO `hm_auth_rule` VALUES (73, 'enterprise/baseInfo', '基础信息', 72, 1, 1, 'icon-bell_start', 1, 'Base Info');
INSERT INTO `hm_auth_rule` VALUES (74, 'enterprise/sliderManage', '幻灯管理', 69, 1, 4, 'icon-application_view_list', 1, 'Slider Manage');
INSERT INTO `hm_auth_rule` VALUES (75, 'site/imageManage', '图片管理', 77, 1, 1, 'icon-flag_ht', 1, 'Image Manage');
INSERT INTO `hm_auth_rule` VALUES (76, 'site/fileManage', '文件管理', 77, 1, 2, 'icon-email_edit', 1, 'File Manage');
INSERT INTO `hm_auth_rule` VALUES (77, 'upload', '上传管理', 1, 1, 3, 'icon-20130406125647919_easyicon_net_16', 1, 'Upload Manage');
INSERT INTO `hm_auth_rule` VALUES (78, 'enterprise/templateType', '模板类型', 69, 1, 1, 'icon-application_view_tile', 1, 'Template Type');

-- ----------------------------
-- Table structure for hm_enterprise_nav
-- ----------------------------
DROP TABLE IF EXISTS `hm_enterprise_nav`;
CREATE TABLE `hm_enterprise_nav`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `image` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `external_link` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `dir` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `type` int NULL DEFAULT 1,
  `template` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `seo_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `seo_keyword` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `seo_describe` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `pid` int NULL DEFAULT 0,
  `isshow` int NULL DEFAULT 1,
  `sort` int NULL DEFAULT 1,
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `level` int NULL DEFAULT 1,
  `en` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_enterprise_nav
-- ----------------------------
INSERT INTO `hm_enterprise_nav` VALUES (1, 'index', '首页', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, 1, 'icon-2012080404391', 1, 'Home');
INSERT INTO `hm_enterprise_nav` VALUES (2, 'works', '工作', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, 2, 'icon-2012080404391', 1, 'Works');
INSERT INTO `hm_enterprise_nav` VALUES (3, 'worksGridText', '工作表格与文本', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 2, 1, 1, 'icon-2012080404391', 1, 'Works grid with text');
INSERT INTO `hm_enterprise_nav` VALUES (4, 'worksGridW', '工作百奥和文本', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 2, 1, 2, 'icon-2012080404391', 1, 'Works grid wo text');
INSERT INTO `hm_enterprise_nav` VALUES (5, 'services', '服务', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, 3, 'icon-2012080404391', 1, 'Services');
INSERT INTO `hm_enterprise_nav` VALUES (6, 'blog', '博客', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, 4, 'icon-2012080404391', 1, 'Blog');
INSERT INTO `hm_enterprise_nav` VALUES (7, 'about', '关于', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, 5, 'icon-2012080404391', 1, 'About');
INSERT INTO `hm_enterprise_nav` VALUES (8, 'shop', '商店', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, 8, 'icon-2012080404391', 1, 'Shop');
INSERT INTO `hm_enterprise_nav` VALUES (9, 'contact', '联系方式', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, 9, 'icon-2012080404391', 1, 'Contact');
INSERT INTO `hm_enterprise_nav` VALUES (10, 'junlunjue', '功昆仑决', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 2, 1, 3, 'icon-2012080404391', 1, 'junlunjue');
INSERT INTO `hm_enterprise_nav` VALUES (11, '/function/ntroduction', '功能简介', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, 13, 'icon-2012080404391', 1, 'Function Introduction');

-- ----------------------------
-- Table structure for hm_enterprise_slider
-- ----------------------------
DROP TABLE IF EXISTS `hm_enterprise_slider`;
CREATE TABLE `hm_enterprise_slider`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `url` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `sort` int NULL DEFAULT 1,
  `image` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_enterprise_slider
-- ----------------------------
INSERT INTO `hm_enterprise_slider` VALUES (1, '菜单设置', 'http://www.baidu.com', 2, '20220121014122495426500.jpg');
INSERT INTO `hm_enterprise_slider` VALUES (2, '后台首页', 'http://www.sina.com.cn', 1, '20220121014102480426500.jpg');
INSERT INTO `hm_enterprise_slider` VALUES (3, '权限管理', 'http://www.sina.com', 3, '20220121014135211426500.jpg');
INSERT INTO `hm_enterprise_slider` VALUES (4, '牛逼的', 'http://baidu.com', 4, '20220121014146225426500.jpg');
INSERT INTO `hm_enterprise_slider` VALUES (5, 'sfafasfdsa', 'sadfasdf', 1, '20220131020926907000000.jpg');
INSERT INTO `hm_enterprise_slider` VALUES (6, 'sfafasfdsa', 'http://www.baidu.com', 1, '20220131020926907000000.jpg');
INSERT INTO `hm_enterprise_slider` VALUES (7, 'sfafasfdsa', 'http://www.baidu.com', 1, '20220131020926907000000.jpg');
INSERT INTO `hm_enterprise_slider` VALUES (8, 'sadfsadf', 'http://www.baiud.com', 2, '20220131021325901000000.jpg');
INSERT INTO `hm_enterprise_slider` VALUES (9, 'sadfsadf', 'http://www.baiud.com', 2, '20220131021325901000000.jpg');
INSERT INTO `hm_enterprise_slider` VALUES (10, 'sadfsadf', 'http://www.baiud.com', 2, '20220131021325901000000.jpg');
INSERT INTO `hm_enterprise_slider` VALUES (11, 'sadfsadf', 'http://www.baiud.com', 2, '20220131021325901000000.jpg');
INSERT INTO `hm_enterprise_slider` VALUES (12, 'sadfsad', 'http://www.baidu.com', 1, '20220131022409285000000.jpg');
INSERT INTO `hm_enterprise_slider` VALUES (13, 'sadfsad', 'http://www.baidu.com', 1, '20220131022409285000000.jpg');
INSERT INTO `hm_enterprise_slider` VALUES (14, 'sadfsad', 'http://www.baidu.com', 1, '20220131022409285000000.jpg');

-- ----------------------------
-- Table structure for hm_enterprise_textmodel
-- ----------------------------
DROP TABLE IF EXISTS `hm_enterprise_textmodel`;
CREATE TABLE `hm_enterprise_textmodel`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `cn_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `cn_text` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `en_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `en_text` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `nid` int NULL DEFAULT NULL,
  `sort` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_enterprise_textmodel
-- ----------------------------
INSERT INTO `hm_enterprise_textmodel` VALUES (1, '创建自己的模板', '在很远的地方，远离Vokalia和Consonantia等国家的山脉背后，有盲文。', 'Create your own template', 'Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the blind texts.', '#iconUIshejiwang', 11, 1);
INSERT INTO `hm_enterprise_textmodel` VALUES (2, '自动备份数据', '遥远的地方，远离Vokalia和Consonantia等国家的山脉背后，有盲文。', 'Automatic Backup Data', 'Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the blind texts.', '#iconzhanku1231', 11, 2);
INSERT INTO `hm_enterprise_textmodel` VALUES (3, '页面构建器', '遥远的地方，远离Vokalia和Consonantia等国家的山脉背后，有盲文。', 'Page Builder', 'Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the blind texts.', '#iconnituwang', 11, 3);
INSERT INTO `hm_enterprise_textmodel` VALUES (4, '萨达刚发生', '阿士大夫撒打算', 'asdfsadfsadf', 'asdfsadfsadf', '#iconword', 11, 4);
INSERT INTO `hm_enterprise_textmodel` VALUES (5, '进口量的台风过后就', '进口量的台风过后就', 'sadfasdfojhdfk', 'sadfasdfojhdfk', '#iconword', 11, 5);
INSERT INTO `hm_enterprise_textmodel` VALUES (6, '老地方客户就是不过来', '老地方客户就是不过来', 'sadfasdfojhdfk', 'sadfasdfojhdfk', '#iconword', 11, 6);
INSERT INTO `hm_enterprise_textmodel` VALUES (7, '了人工', '了人工', 'sadfsadfsa', 'asfdsadfsadf', '#iconword', 11, 7);
INSERT INTO `hm_enterprise_textmodel` VALUES (8, '了人工电风扇', '了人工电风扇', 'sdfsadfsa', 'asdfsdfds', '#iconword', 11, 8);
INSERT INTO `hm_enterprise_textmodel` VALUES (9, '领导风格看', '领导风格看', 'sadfsadfas', 'asfdasdfasdf', '#iconword', 11, 9);
INSERT INTO `hm_enterprise_textmodel` VALUES (10, '领导是否感觉到是', '领导是否感觉到是', 'asdfsadfdsa', 'asdfsadf', '#iconword', 11, 10);
INSERT INTO `hm_enterprise_textmodel` VALUES (11, '六点十分攻击力', '六点十分攻击力', 'sdfsadf', 'asdfsadf', '#iconword', 11, 11);

-- ----------------------------
-- Table structure for hm_enterprise_tptype
-- ----------------------------
DROP TABLE IF EXISTS `hm_enterprise_tptype`;
CREATE TABLE `hm_enterprise_tptype`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `cn` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `en` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_enterprise_tptype
-- ----------------------------
INSERT INTO `hm_enterprise_tptype` VALUES (1, '单页模型', 'Single Page Model', 1);
INSERT INTO `hm_enterprise_tptype` VALUES (2, '文章模型', 'Article Model', 2);
INSERT INTO `hm_enterprise_tptype` VALUES (3, '下载模型', 'Download Model', 3);
INSERT INTO `hm_enterprise_tptype` VALUES (4, '本文模型', 'Text Model', 4);

-- ----------------------------
-- Table structure for hm_image
-- ----------------------------
DROP TABLE IF EXISTS `hm_image`;
CREATE TABLE `hm_image`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `time` int NOT NULL,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `width` int NULL DEFAULT NULL,
  `height` int NULL DEFAULT NULL,
  `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `size` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_image
-- ----------------------------
INSERT INTO `hm_image` VALUES (1, '20220131014212774000000.jpg', 1643564532, 'admin', 120, 120, 'jpg', 141237);
INSERT INTO `hm_image` VALUES (2, '20220131014854921000000.jpg', 1643564934, 'admin', 120, 120, 'jpg', 186759);
INSERT INTO `hm_image` VALUES (3, '20220131014922230000000.jpg', 1643564962, 'admin', 120, 120, 'jpg', 186759);
INSERT INTO `hm_image` VALUES (4, '20220131014929701000000.jpg', 1643564969, 'admin', 120, 120, 'jpg', 186759);
INSERT INTO `hm_image` VALUES (5, '20220131014956443000000.jpg', 1643564996, 'admin', 120, 120, 'jpg', 141237);
INSERT INTO `hm_image` VALUES (6, '20220131015017442000000.jpg', 1643565017, 'admin', 120, 120, 'jpg', 186759);
INSERT INTO `hm_image` VALUES (7, '20220131015052392000000.jpg', 1643565052, 'admin', 120, 120, 'jpg', 186759);
INSERT INTO `hm_image` VALUES (8, '20220131015303962000000.jpg', 1643565183, 'admin', 120, 120, 'jpg', 186759);
INSERT INTO `hm_image` VALUES (9, '20220131020926907000000.jpg', 1643566166, 'admin', 120, 120, 'jpg', 186759);
INSERT INTO `hm_image` VALUES (10, '20220131021325901000000.jpg', 1643566405, 'admin', 120, 120, 'jpg', 186759);
INSERT INTO `hm_image` VALUES (11, '20220131022409285000000.jpg', 1643567049, 'admin', 120, 120, 'jpg', 186759);

-- ----------------------------
-- Table structure for hm_login_log
-- ----------------------------
DROP TABLE IF EXISTS `hm_login_log`;
CREATE TABLE `hm_login_log`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `login_time` int NOT NULL,
  `login_ip` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `status` tinyint(1) NOT NULL,
  `area` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `country` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `useragent` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `uid` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 68 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_login_log
-- ----------------------------
INSERT INTO `hm_login_log` VALUES (62, 'admin', 1643563896, '127.0.0.1', 1, '', '', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', 1);
INSERT INTO `hm_login_log` VALUES (63, 'admin2', 1643564154, '127.0.0.1', 1, '', '', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', 3);
INSERT INTO `hm_login_log` VALUES (64, 'admin', 1643564312, '127.0.0.1', 1, '', '', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', 1);
INSERT INTO `hm_login_log` VALUES (65, 'admin', 1643564429, '127.0.0.1', 1, '', '', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', 1);
INSERT INTO `hm_login_log` VALUES (66, 'admin', 1643565693, '127.0.0.1', 1, '', '', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', 1);
INSERT INTO `hm_login_log` VALUES (67, 'admin', 1643567323, '127.0.0.1', 1, '', '', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', 1);
INSERT INTO `hm_login_log` VALUES (68, 'admin', 1643568331, '127.0.0.1', 1, '', '', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', 1);
INSERT INTO `hm_login_log` VALUES (69, 'admin2', 1643657679, '127.0.0.1', 1, '', '', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', 3);
INSERT INTO `hm_login_log` VALUES (70, 'admin', 1643657704, '127.0.0.1', 1, '', '', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', 1);

-- ----------------------------
-- Table structure for hm_oprate_log
-- ----------------------------
DROP TABLE IF EXISTS `hm_oprate_log`;
CREATE TABLE `hm_oprate_log`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `oprate_time` int NOT NULL,
  `oprate_ip` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `useragent` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `detail` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `info` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `method` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `excute_time` int NOT NULL,
  `status` tinyint(1) NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 82 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_oprate_log
-- ----------------------------
INSERT INTO `hm_oprate_log` VALUES (1, 'admin2', 1643484484, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*LoginCtrl).LoginOut<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/login.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>133', 'no problem', '/intendant/loginOut', 'POST', 213, 1);
INSERT INTO `hm_oprate_log` VALUES (2, 'admin2', 1643484621, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*LoginCtrl).LoginOut<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/login.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>133', 'no problem', '/intendant/loginOut', 'POST', 683, 1);
INSERT INTO `hm_oprate_log` VALUES (3, 'admin2', 1643484784, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*LoginCtrl).LoginOut<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/login.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>133', 'no problem', '/intendant/loginOut', 'POST', 784, 1);
INSERT INTO `hm_oprate_log` VALUES (4, 'admin2', 1643484819, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*LoginCtrl).LoginOut<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/login.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>133', 'no problem', '/intendant/loginOut', 'POST', 671, 1);
INSERT INTO `hm_oprate_log` VALUES (5, 'admin2', 1643561969, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*LoginCtrl).LoginOut<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/login.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>136', 'no problem', '/intendant/loginOut', 'POST', 539, 1);
INSERT INTO `hm_oprate_log` VALUES (6, 'admin', 1643562701, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*LoginCtrl).LoginOut<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/login.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>137', 'no problem', '/intendant/loginOut', 'POST', 595, 1);
INSERT INTO `hm_oprate_log` VALUES (7, 'admin', 1643562801, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1627', 'sql: Scan error on column index 5, name \"area\": converting NULL to string is unsupported', '/intendant/site/getLoginLog', 'POST', 90, 4);
INSERT INTO `hm_oprate_log` VALUES (8, 'admin', 1643562801, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1637', 'no problem', '/intendant/site/getLoginLog', 'POST', 97, 1);
INSERT INTO `hm_oprate_log` VALUES (9, 'admin', 1643562805, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1637', 'no problem', '/intendant/site/getLoginLog', 'POST', 728, 1);
INSERT INTO `hm_oprate_log` VALUES (10, 'admin', 1643562808, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1637', 'no problem', '/intendant/site/getLoginLog', 'POST', 961, 1);
INSERT INTO `hm_oprate_log` VALUES (11, 'admin', 1643562812, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1637', 'no problem', '/intendant/site/getLoginLog', 'POST', 140, 1);
INSERT INTO `hm_oprate_log` VALUES (12, 'admin', 1643562909, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1637', 'no problem', '/intendant/site/getLoginLog', 'POST', 854, 1);
INSERT INTO `hm_oprate_log` VALUES (13, 'admin', 1643562910, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1637', 'no problem', '/intendant/site/getLoginLog', 'POST', 699, 1);
INSERT INTO `hm_oprate_log` VALUES (14, 'admin', 1643563024, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1637', 'no problem', '/intendant/site/getLoginLog', 'POST', 417, 1);
INSERT INTO `hm_oprate_log` VALUES (15, 'admin', 1643563050, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1637', 'no problem', '/intendant/site/getLoginLog', 'POST', 860, 1);
INSERT INTO `hm_oprate_log` VALUES (16, 'admin', 1643563056, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*LoginCtrl).LoginOut<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/login.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>137', 'no problem', '/intendant/loginOut', 'POST', 19, 1);
INSERT INTO `hm_oprate_log` VALUES (17, 'admin2', 1643563067, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1637', 'no problem', '/intendant/site/getLoginLog', 'POST', 581, 1);
INSERT INTO `hm_oprate_log` VALUES (18, 'admin2', 1643563070, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*LoginCtrl).LoginOut<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/login.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>137', 'no problem', '/intendant/loginOut', 'POST', 519, 1);
INSERT INTO `hm_oprate_log` VALUES (19, 'admin', 1643563088, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1637', 'no problem', '/intendant/site/getLoginLog', 'POST', 163, 1);
INSERT INTO `hm_oprate_log` VALUES (20, 'admin', 1643563388, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1627', 'sql: Scan error on column index 5, name \"area\": converting NULL to string is unsupported', '/intendant/site/getLoginLog', 'POST', 361, 4);
INSERT INTO `hm_oprate_log` VALUES (21, 'admin', 1643563388, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1638', 'no problem', '/intendant/site/getLoginLog', 'POST', 887, 1);
INSERT INTO `hm_oprate_log` VALUES (22, 'admin', 1643563422, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1627', 'sql: Scan error on column index 5, name \"area\": converting NULL to string is unsupported', '/intendant/site/getLoginLog', 'POST', 76, 4);
INSERT INTO `hm_oprate_log` VALUES (23, 'admin', 1643563422, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1638', 'no problem', '/intendant/site/getLoginLog', 'POST', 126, 1);
INSERT INTO `hm_oprate_log` VALUES (24, 'admin', 1643563428, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1638', 'no problem', '/intendant/site/getLoginLog', 'POST', 342, 1);
INSERT INTO `hm_oprate_log` VALUES (25, 'admin', 1643563541, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1638', 'no problem', '/intendant/site/getLoginLog', 'POST', 753, 1);
INSERT INTO `hm_oprate_log` VALUES (26, 'admin', 1643563544, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1638', 'no problem', '/intendant/site/getLoginLog', 'POST', 804, 1);
INSERT INTO `hm_oprate_log` VALUES (27, 'admin', 1643563592, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1638', 'no problem', '/intendant/site/getLoginLog', 'POST', 20, 1);
INSERT INTO `hm_oprate_log` VALUES (28, 'admin', 1643563600, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1638', 'no problem', '/intendant/site/getLoginLog', 'POST', 705, 1);
INSERT INTO `hm_oprate_log` VALUES (29, 'admin', 1643563602, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1638', 'no problem', '/intendant/site/getLoginLog', 'POST', 727, 1);
INSERT INTO `hm_oprate_log` VALUES (30, 'admin', 1643563604, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1638', 'no problem', '/intendant/site/getLoginLog', 'POST', 543, 1);
INSERT INTO `hm_oprate_log` VALUES (31, 'admin', 1643563716, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1638', 'no problem', '/intendant/site/getLoginLog', 'POST', 855, 1);
INSERT INTO `hm_oprate_log` VALUES (32, 'admin', 1643563718, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1638', 'no problem', '/intendant/site/getLoginLog', 'POST', 281, 1);
INSERT INTO `hm_oprate_log` VALUES (33, 'admin', 1643563817, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1638', 'no problem', '/intendant/site/getLoginLog', 'POST', 462, 1);
INSERT INTO `hm_oprate_log` VALUES (34, 'admin', 1643563878, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1638', 'no problem', '/intendant/site/getLoginLog', 'POST', 202, 1);
INSERT INTO `hm_oprate_log` VALUES (35, 'admin', 1643563892, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*LoginCtrl).LoginOut<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/login.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>133', 'no problem', '/intendant/loginOut', 'POST', 107, 1);
INSERT INTO `hm_oprate_log` VALUES (36, 'admin', 1643563902, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1638', 'no problem', '/intendant/site/getLoginLog', 'POST', 404, 1);
INSERT INTO `hm_oprate_log` VALUES (37, 'admin', 1643563904, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1638', 'no problem', '/intendant/site/getLoginLog', 'POST', 528, 1);
INSERT INTO `hm_oprate_log` VALUES (38, 'admin', 1643564075, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1639', 'no problem', '/intendant/site/getLoginLog', 'POST', 509, 1);
INSERT INTO `hm_oprate_log` VALUES (39, 'admin', 1643564100, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1639', 'no problem', '/intendant/site/getLoginLog', 'POST', 115, 1);
INSERT INTO `hm_oprate_log` VALUES (40, 'admin', 1643564147, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*LoginCtrl).LoginOut<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/login.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>133', 'no problem', '/intendant/loginOut', 'POST', 862, 1);
INSERT INTO `hm_oprate_log` VALUES (41, 'admin2', 1643564160, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1639', 'no problem', '/intendant/site/getLoginLog', 'POST', 371, 1);
INSERT INTO `hm_oprate_log` VALUES (42, 'admin2', 1643564169, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1639', 'no problem', '/intendant/site/getLoginLog', 'POST', 619, 1);
INSERT INTO `hm_oprate_log` VALUES (43, 'admin2', 1643564291, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1639', 'no problem', '/intendant/site/getLoginLog', 'POST', 781, 1);
INSERT INTO `hm_oprate_log` VALUES (44, 'admin2', 1643564308, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*LoginCtrl).LoginOut<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/login.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>133', 'no problem', '/intendant/loginOut', 'POST', 103, 1);
INSERT INTO `hm_oprate_log` VALUES (45, 'admin', 1643564318, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1639', 'no problem', '/intendant/site/getLoginLog', 'POST', 197, 1);
INSERT INTO `hm_oprate_log` VALUES (46, 'admin', 1643564420, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1639', 'no problem', '/intendant/site/getLoginLog', 'POST', 123, 1);
INSERT INTO `hm_oprate_log` VALUES (47, 'admin', 1643564424, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*LoginCtrl).LoginOut<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/login.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>134', 'no problem', '/intendant/loginOut', 'POST', 892, 1);
INSERT INTO `hm_oprate_log` VALUES (48, 'admin', 1643564434, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1639', 'no problem', '/intendant/site/getLoginLog', 'POST', 522, 1);
INSERT INTO `hm_oprate_log` VALUES (49, 'admin', 1643564457, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetImage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>325', 'no found data', '/intendant/site/getImage', 'POST', 670, 54);
INSERT INTO `hm_oprate_log` VALUES (50, 'admin', 1643564532, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).UploadImages<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>145', '20220131014212774000000.jpg', '/intendant/site/uploadImages', 'POST', 788, 1);
INSERT INTO `hm_oprate_log` VALUES (51, 'admin', 1643564538, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetImage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>325', 'no problem', '/intendant/site/getImage', 'POST', 857, 1);
INSERT INTO `hm_oprate_log` VALUES (52, 'admin', 1643564542, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetUser<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>913', 'no problem', '/intendant/site/getUser', 'POST', 882, 1);
INSERT INTO `hm_oprate_log` VALUES (53, 'admin', 1643564934, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).UploadImage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>91', '20220131014854921000000.jpg', '/intendant/site/uploadImage', 'POST', 934, 1);
INSERT INTO `hm_oprate_log` VALUES (54, 'admin', 1643564962, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).UploadImage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>91', '20220131014922230000000.jpg', '/intendant/site/uploadImage', 'POST', 243, 1);
INSERT INTO `hm_oprate_log` VALUES (55, 'admin', 1643564969, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).UploadImage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>91', '20220131014929701000000.jpg', '/intendant/site/uploadImage', 'POST', 709, 1);
INSERT INTO `hm_oprate_log` VALUES (56, 'admin', 1643564985, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetImage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>325', 'no problem', '/intendant/site/getImage', 'POST', 146, 1);
INSERT INTO `hm_oprate_log` VALUES (57, 'admin', 1643564996, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).UploadImages<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>145', '20220131014956443000000.jpg', '/intendant/site/uploadImages', 'POST', 453, 1);
INSERT INTO `hm_oprate_log` VALUES (58, 'admin', 1643565000, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetImage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>325', 'no problem', '/intendant/site/getImage', 'POST', 838, 1);
INSERT INTO `hm_oprate_log` VALUES (59, 'admin', 1643565017, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).UploadImage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>91', '20220131015017442000000.jpg', '/intendant/site/uploadImage', 'POST', 456, 1);
INSERT INTO `hm_oprate_log` VALUES (60, 'admin', 1643565052, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).UploadImage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>91', '20220131015052392000000.jpg', '/intendant/site/uploadImage', 'POST', 402, 1);
INSERT INTO `hm_oprate_log` VALUES (61, 'admin', 1643565183, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).UploadImage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>91', '20220131015303962000000.jpg', '/intendant/site/uploadImage', 'POST', 972, 1);
INSERT INTO `hm_oprate_log` VALUES (62, 'admin', 1643565251, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1639', 'no problem', '/intendant/site/getLoginLog', 'POST', 398, 1);
INSERT INTO `hm_oprate_log` VALUES (63, 'admin', 1643566166, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).UploadImage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>91', '20220131020926907000000.jpg', '/intendant/site/uploadImage', 'POST', 925, 1);
INSERT INTO `hm_oprate_log` VALUES (64, 'admin', 1643566194, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetImage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>325', 'no problem', '/intendant/site/getImage', 'POST', 209, 1);
INSERT INTO `hm_oprate_log` VALUES (65, 'admin', 1643566405, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).UploadImage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>91', '20220131021325901000000.jpg', '/intendant/site/uploadImage', 'POST', 916, 1);
INSERT INTO `hm_oprate_log` VALUES (66, 'admin', 1643567049, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).UploadImage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>91', '20220131022409285000000.jpg', '/intendant/site/uploadImage', 'POST', 300, 1);
INSERT INTO `hm_oprate_log` VALUES (67, 'admin', 1643567329, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1639', 'no problem', '/intendant/site/getLoginLog', 'POST', 489, 1);
INSERT INTO `hm_oprate_log` VALUES (68, 'admin', 1643568094, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1639', 'no problem', '/intendant/site/getLoginLog', 'POST', 414, 1);
INSERT INTO `hm_oprate_log` VALUES (73, 'admin', 1643568520, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1639', 'no problem', '/intendant/site/getLoginLog', 'POST', 612, 1);
INSERT INTO `hm_oprate_log` VALUES (74, 'admin', 1643568581, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1639', 'no problem', '/intendant/site/getLoginLog', 'POST', 948, 1);
INSERT INTO `hm_oprate_log` VALUES (75, 'admin', 1643568741, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).GetSlider<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>892', 'no problem', '/intendant/enterprise/getSlider', 'POST', 1002, 1);
INSERT INTO `hm_oprate_log` VALUES (76, 'admin', 1643568780, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).GetTemplateType<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>83', 'no problem', '/intendant/enterprise/getTemplateType', 'POST', 865, 1);
INSERT INTO `hm_oprate_log` VALUES (77, 'admin', 1643568784, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).GetNav<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>316', 'sql: Scan error on column index 3, name \"image\": converting NULL to string is unsupported', '/intendant/enterprise/getNav', 'POST', 118, 4);
INSERT INTO `hm_oprate_log` VALUES (78, 'admin', 1643568806, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).GetSlider<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>892', 'no problem', '/intendant/enterprise/getSlider', 'POST', 437, 1);
INSERT INTO `hm_oprate_log` VALUES (79, 'admin', 1643568906, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/enterprise.(*IndexCtrl).Index<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/enterprise/index.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>27', 'sql: Scan error on column index 3, name \"image\": converting NULL to string is unsupported', '/enterprise/index', 'GET', 0, 4);
INSERT INTO `hm_oprate_log` VALUES (80, 'admin', 1643568915, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*IndexCtrl).AjaxUpdateLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/index.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>99', 'no problem', '/intendant/ajaxUpdateLog', 'POST', 239, 1);
INSERT INTO `hm_oprate_log` VALUES (81, 'admin', 1643568918, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetImage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>325', 'no problem', '/intendant/site/getImage', 'POST', 893, 1);
INSERT INTO `hm_oprate_log` VALUES (82, 'admin', 1643568921, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetLoginLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1639', 'no problem', '/intendant/site/getLoginLog', 'POST', 821, 1);
INSERT INTO `hm_oprate_log` VALUES (83, 'admin', 1643656505, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*IndexCtrl).AjaxUpdateLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/index.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>99', 'no problem', '/intendant/ajaxUpdateLog', 'POST', 29, 1);
INSERT INTO `hm_oprate_log` VALUES (84, 'admin', 1643656507, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).GetSlider<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>892', 'no problem', '/intendant/enterprise/getSlider', 'POST', 984, 1);
INSERT INTO `hm_oprate_log` VALUES (85, 'admin', 1643656522, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).GetTemplateType<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>83', 'no problem', '/intendant/enterprise/getTemplateType', 'POST', 460, 1);
INSERT INTO `hm_oprate_log` VALUES (86, 'admin', 1643656523, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).GetNav<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>316', 'sql: Scan error on column index 3, name \"image\": converting NULL to string is unsupported', '/intendant/enterprise/getNav', 'POST', 472, 4);
INSERT INTO `hm_oprate_log` VALUES (87, 'admin', 1643657061, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).ContentManage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>654', 'sql: Scan error on column index 3, name \"image\": converting NULL to string is unsupported', '/intendant/enterprise/contentManage', 'GET', 296, 4);
INSERT INTO `hm_oprate_log` VALUES (88, 'admin', 1643657284, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*IndexCtrl).AjaxUpdateLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/index.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>99', 'no problem', '/intendant/ajaxUpdateLog', 'POST', 783, 1);
INSERT INTO `hm_oprate_log` VALUES (89, 'admin', 1643657344, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).ContentManage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>654', 'sql: Scan error on column index 3, name \"image\": converting NULL to string is unsupported', '/intendant/enterprise/contentManage', 'GET', 892, 4);
INSERT INTO `hm_oprate_log` VALUES (90, 'admin', 1643657465, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).ContentManage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>654', 'sql: Scan error on column index 3, name \"image\": converting NULL to string is unsupported', '/intendant/enterprise/contentManage', 'GET', 164, 4);
INSERT INTO `hm_oprate_log` VALUES (91, 'admin', 1643657530, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetUser<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>913', 'no problem', '/intendant/site/getUser', 'POST', 587, 1);
INSERT INTO `hm_oprate_log` VALUES (92, 'admin', 1643657546, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetUser<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>913', 'no problem', '/intendant/site/getUser', 'POST', 47, 1);
INSERT INTO `hm_oprate_log` VALUES (93, 'admin', 1643657548, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetRole<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1263', 'no problem', '/intendant/site/getRole', 'POST', 619, 1);
INSERT INTO `hm_oprate_log` VALUES (94, 'admin', 1643657571, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).SetRoleSubmit<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1346', 'no problem', '/intendant/site/setRoleSubmit', 'POST', 959, 1);
INSERT INTO `hm_oprate_log` VALUES (95, 'admin', 1643657593, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).ContentManage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>654', 'sql: Scan error on column index 3, name \"image\": converting NULL to string is unsupported', '/intendant/enterprise/contentManage', 'GET', 358, 4);
INSERT INTO `hm_oprate_log` VALUES (96, 'admin', 1643657615, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).GetTemplateType<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>83', 'no problem', '/intendant/enterprise/getTemplateType', 'POST', 565, 1);
INSERT INTO `hm_oprate_log` VALUES (97, 'admin', 1643657655, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).GetRole<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1263', 'no problem', '/intendant/site/getRole', 'POST', 273, 1);
INSERT INTO `hm_oprate_log` VALUES (98, 'admin', 1643657667, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).SetRoleSubmit<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>1346', 'no problem', '/intendant/site/setRoleSubmit', 'POST', 497, 1);
INSERT INTO `hm_oprate_log` VALUES (99, 'admin', 1643657674, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*LoginCtrl).LoginOut<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/login.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>134', 'no problem', '/intendant/loginOut', 'POST', 958, 1);
INSERT INTO `hm_oprate_log` VALUES (100, 'admin2', 1643657683, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*IndexCtrl).AjaxUpdateLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/index.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>99', 'no problem', '/intendant/ajaxUpdateLog', 'POST', 297, 1);
INSERT INTO `hm_oprate_log` VALUES (101, 'admin2', 1643657699, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*LoginCtrl).LoginOut<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/login.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>134', 'no problem', '/intendant/loginOut', 'POST', 943, 1);
INSERT INTO `hm_oprate_log` VALUES (102, 'admin', 1643657708, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*IndexCtrl).AjaxUpdateLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/index.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>99', 'no problem', '/intendant/ajaxUpdateLog', 'POST', 8, 1);
INSERT INTO `hm_oprate_log` VALUES (103, 'admin', 1643657726, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).ContentManage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>654', 'sql: Scan error on column index 3, name \"image\": converting NULL to string is unsupported', '/intendant/enterprise/contentManage', 'GET', 420, 4);
INSERT INTO `hm_oprate_log` VALUES (104, 'admin', 1643657888, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).ContentManage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>654', 'sql: Scan error on column index 4, name \"external_link\": converting NULL to string is unsupported', '/intendant/enterprise/contentManage', 'GET', 90, 4);
INSERT INTO `hm_oprate_log` VALUES (105, 'admin', 1643657931, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).ContentManage<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>654', 'sql: Scan error on column index 6, name \"type\": converting NULL to int is unsupported', '/intendant/enterprise/contentManage', 'GET', 304, 4);
INSERT INTO `hm_oprate_log` VALUES (106, 'admin', 1643658023, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*IndexCtrl).AjaxUpdateLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/index.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>99', 'no problem', '/intendant/ajaxUpdateLog', 'POST', 431, 1);
INSERT INTO `hm_oprate_log` VALUES (107, 'admin', 1643658036, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).GetSlider<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>892', 'no problem', '/intendant/enterprise/getSlider', 'POST', 425, 1);
INSERT INTO `hm_oprate_log` VALUES (108, 'admin', 1643658050, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).GetTemplateType<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>83', 'no problem', '/intendant/enterprise/getTemplateType', 'POST', 495, 1);
INSERT INTO `hm_oprate_log` VALUES (109, 'admin', 1643658352, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*IndexCtrl).AjaxUpdateLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/index.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>99', 'no problem', '/intendant/ajaxUpdateLog', 'POST', 456, 1);
INSERT INTO `hm_oprate_log` VALUES (110, 'admin', 1643658354, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*EnterpriseCtrl).GetSlider<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/enterprise.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>892', 'no problem', '/intendant/enterprise/getSlider', 'POST', 444, 1);
INSERT INTO `hm_oprate_log` VALUES (111, 'admin', 1643658403, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*SiteCtrl).EditMenuSubmit<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/site.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>672', 'no problem', '/intendant/site/editMenuSubmit', 'POST', 588, 1);
INSERT INTO `hm_oprate_log` VALUES (112, 'admin', 1643658407, '127.0.0.1', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36', '<span style=\'color:#3498db\'>Package&nbsp;:&nbsp;</span>hemacms/ctrl/admin.(*IndexCtrl).AjaxUpdateLog<br/><span style=\'color:#e67e22\'>File&nbsp;:&nbsp;</span>F:/gopath/src/hemacms/ctrl/admin/index.go<br/><span style=\'color:#9b59b6\'>Line&nbsp;:&nbsp;</span>99', 'no problem', '/intendant/ajaxUpdateLog', 'POST', 546, 1);

-- ----------------------------
-- Table structure for hm_update_log
-- ----------------------------
DROP TABLE IF EXISTS `hm_update_log`;
CREATE TABLE `hm_update_log`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `time` int NOT NULL,
  `cn` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `en` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_update_log
-- ----------------------------

-- ----------------------------
-- Table structure for hm_user
-- ----------------------------
DROP TABLE IF EXISTS `hm_user`;
CREATE TABLE `hm_user`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `nickname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `create_time` int NULL DEFAULT 0,
  `create_ip` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '127.0.0.1',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `status` tinyint(1) NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_user
-- ----------------------------
INSERT INTO `hm_user` VALUES (1, 'admin', '河马', '90b9aa7e25f80cf4f64e990b78a9fc5ebd6cecad', 'iteny@msn.com', 0, '127.0.0.1', '', 1);
INSERT INTO `hm_user` VALUES (2, 'testme', '测试', '90b9aa7e25f80cf4f64e990b78a9fc5ebd6cecad', 'test@msn.com', 0, '127.0.0.1', '', 1);
INSERT INTO `hm_user` VALUES (3, 'admin2', 'lasdgjksaldjg', '90b9aa7e25f80cf4f64e990b78a9fc5ebd6cecad', 'lasfgjlasd@qq.com', 1524583419, '127.0.0.1', '', 1);

SET FOREIGN_KEY_CHECKS = 1;
