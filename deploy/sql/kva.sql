/*
 Navicat Premium Data Transfer

 Source Server         : im测试服
 Source Server Type    : MySQL
 Source Server Version : 50742
 Source Host           : localhost:13306
 Source Schema         : kva

 Target Server Type    : MySQL
 Target Server Version : 50742
 File Encoding         : 65001

 Date: 07/09/2023 16:33:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- Solve problem insert default datetime error incorrect datetime value: '0000-00-00 00:00:00'
-- 解决orm插入datetime类型默认值'0000-00-00 00:00:00'会报错问题
SET @@global.sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 167 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (58, 'p', 'admin', '/api.admin.v1.Api/AllApi', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (61, 'p', 'admin', '/api.admin.v1.Api/CreateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (63, 'p', 'admin', '/api.admin.v1.Api/DeleteApi', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (60, 'p', 'admin', '/api.admin.v1.Api/GetApi', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (59, 'p', 'admin', '/api.admin.v1.Api/GetPolicyPathByRoleKey', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (57, 'p', 'admin', '/api.admin.v1.Api/ListApi', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (62, 'p', 'admin', '/api.admin.v1.Api/UpdateApi', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (163, 'p', 'admin', '/api.admin.v1.black/blackList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (164, 'p', 'admin', '/api.admin.v1.black/DeleteBlack', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (21, 'p', 'admin', '/api.admin.v1.Dept/CreateDept', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (23, 'p', 'admin', '/api.admin.v1.Dept/DeleteDept', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (20, 'p', 'admin', '/api.admin.v1.Dept/GetDeptTree', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (17, 'p', 'admin', '/api.admin.v1.Dept/ListDept', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (18, 'p', 'admin', '/api.admin.v1.Dept/ListDept:deptId', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (19, 'p', 'admin', '/api.admin.v1.Dept/RoleDeptTreeSelect', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (22, 'p', 'admin', '/api.admin.v1.Dept/UpdateDept', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (33, 'p', 'admin', '/api.admin.v1.DictData/CreateDictData', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (35, 'p', 'admin', '/api.admin.v1.DictData/DeleteDictData', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (32, 'p', 'admin', '/api.admin.v1.DictData/GetDictData', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (31, 'p', 'admin', '/api.admin.v1.DictData/GetDictData1', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (30, 'p', 'admin', '/api.admin.v1.DictData/ListDictData', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (34, 'p', 'admin', '/api.admin.v1.DictData/UpdateDictData', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (26, 'p', 'admin', '/api.admin.v1.DictType/CreateDictType', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (28, 'p', 'admin', '/api.admin.v1.DictType/DeleteDictType', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (25, 'p', 'admin', '/api.admin.v1.DictType/GetDictType', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (24, 'p', 'admin', '/api.admin.v1.DictType/ListDictType', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (27, 'p', 'admin', '/api.admin.v1.DictType/UpdateDictType', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (78, 'p', 'admin', '/api.admin.v1.DiscoveryPage/DiscoveryPageCreate', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (79, 'p', 'admin', '/api.admin.v1.DiscoveryPage/DiscoveryPageDelete', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (80, 'p', 'admin', '/api.admin.v1.DiscoveryPage/DiscoveryPageList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (81, 'p', 'admin', '/api.admin.v1.DiscoveryPage/DiscoveryPageUpdate', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (166, 'p', 'admin', '/api.admin.v1.FeedBack/FeedBackList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (165, 'p', 'admin', '/api.admin.v1.FeedBack/HandleFeedBack', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (159, 'p', 'admin', '/api.admin.v1.Friend/AddBlack', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (160, 'p', 'admin', '/api.admin.v1.Friend/AddFriend', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (161, 'p', 'admin', '/api.admin.v1.Friend/DeleteFriend', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (162, 'p', 'admin', '/api.admin.v1.Friend/FriendList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (82, 'p', 'admin', '/api.admin.v1.FunctionalConfig/FunctionalConfigGet', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (83, 'p', 'admin', '/api.admin.v1.FunctionalConfig/FunctionalConfigUpdate', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (86, 'p', 'admin', '/api.admin.v1.GroupManager/GroupBan', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (87, 'p', 'admin', '/api.admin.v1.GroupManager/GroupCreate', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (88, 'p', 'admin', '/api.admin.v1.GroupManager/GroupDismiss', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (89, 'p', 'admin', '/api.admin.v1.GroupManager/GroupInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (101, 'p', 'admin', '/api.admin.v1.GroupManager/GroupJoined', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (90, 'p', 'admin', '/api.admin.v1.GroupManager/GroupList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (91, 'p', 'admin', '/api.admin.v1.GroupManager/GroupMember', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (92, 'p', 'admin', '/api.admin.v1.GroupManager/GroupMemberAdd', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (93, 'p', 'admin', '/api.admin.v1.GroupManager/GroupMemberAdmin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (100, 'p', 'admin', '/api.admin.v1.GroupManager/GroupMemberInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (94, 'p', 'admin', '/api.admin.v1.GroupManager/GroupMemberKickOut', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (95, 'p', 'admin', '/api.admin.v1.GroupManager/GroupMemberLeader', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (96, 'p', 'admin', '/api.admin.v1.GroupManager/GroupMemberMute', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (97, 'p', 'admin', '/api.admin.v1.GroupManager/GroupMute', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (98, 'p', 'admin', '/api.admin.v1.GroupManager/GroupStatusNormal', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (99, 'p', 'admin', '/api.admin.v1.GroupManager/GroupUpdate', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (123, 'p', 'admin', '/api.admin.v1.InteratedTask/DeleteHandleLog', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (124, 'p', 'admin', '/api.admin.v1.InteratedTask/GetHandleLogList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (129, 'p', 'admin', '/api.admin.v1.InteratedTask/PrivilegeUserAddUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (130, 'p', 'admin', '/api.admin.v1.InteratedTask/PrivilegeUserBatchAddFriend', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (131, 'p', 'admin', '/api.admin.v1.InteratedTask/PrivilegeUserBatchAddGroup', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (132, 'p', 'admin', '/api.admin.v1.InteratedTask/PrivilegeUserDeleteUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (137, 'p', 'admin', '/api.admin.v1.InteratedTask/PrivilegeUserDisableUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (133, 'p', 'admin', '/api.admin.v1.InteratedTask/PrivilegeUserEditUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (134, 'p', 'admin', '/api.admin.v1.InteratedTask/PrivilegeUserEnableUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (135, 'p', 'admin', '/api.admin.v1.InteratedTask/PrivilegeUserIpWhitelist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (136, 'p', 'admin', '/api.admin.v1.InteratedTask/PrivilegeUserList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (125, 'p', 'admin', '/api.admin.v1.InteratedTask/SystemMsgAdd', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (126, 'p', 'admin', '/api.admin.v1.InteratedTask/SystemMsgDelete', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (127, 'p', 'admin', '/api.admin.v1.InteratedTask/SystemMsgList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (128, 'p', 'admin', '/api.admin.v1.InteratedTask/SystemMsgRetry', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (153, 'p', 'admin', '/api.admin.v1.IpBlackList/BatchUnbanIpBlackList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (155, 'p', 'admin', '/api.admin.v1.IpBlackList/CreateIpBlackList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (156, 'p', 'admin', '/api.admin.v1.IpBlackList/DeleteIpBlackList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (157, 'p', 'admin', '/api.admin.v1.IpBlackList/GetIpBlackListPaginationData', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (158, 'p', 'admin', '/api.admin.v1.IpBlackList/IpBlackListRelationUsers', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (154, 'p', 'admin', '/api.admin.v1.IpBlackList/UnbanIpBlackList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (147, 'p', 'admin', '/api.admin.v1.IpWhiteList/BatchDeleteIpWhiteList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (148, 'p', 'admin', '/api.admin.v1.IpWhiteList/CreateIpWhiteList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (149, 'p', 'admin', '/api.admin.v1.IpWhiteList/DeleteIpWhiteList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (150, 'p', 'admin', '/api.admin.v1.IpWhiteList/GetPaginationData', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (151, 'p', 'admin', '/api.admin.v1.IpWhiteList/RemoveIpWhiteList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (152, 'p', 'admin', '/api.admin.v1.IpWhiteList/UpdateIpWhiteList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (84, 'p', 'admin', '/api.admin.v1.LoginDevice/DeviceList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (85, 'p', 'admin', '/api.admin.v1.LoginDevice/DeviceUserLogout', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (41, 'p', 'admin', '/api.admin.v1.Menus/CreateMenus', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (43, 'p', 'admin', '/api.admin.v1.Menus/DeleteMenus', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (40, 'p', 'admin', '/api.admin.v1.Menus/GetMenus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (36, 'p', 'admin', '/api.admin.v1.Menus/GetMenusTree', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (39, 'p', 'admin', '/api.admin.v1.Menus/ListMenus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (38, 'p', 'admin', '/api.admin.v1.Menus/RoleMenuTreeSelect', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (42, 'p', 'admin', '/api.admin.v1.Menus/UpdateMenus', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (139, 'p', 'admin', '/api.admin.v1.Msg/GroupList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (138, 'p', 'admin', '/api.admin.v1.Msg/SingleList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (54, 'p', 'admin', '/api.admin.v1.Roles/ChangeRoleStatus', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (51, 'p', 'admin', '/api.admin.v1.Roles/CreateRoles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (55, 'p', 'admin', '/api.admin.v1.Roles/DataScope', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (53, 'p', 'admin', '/api.admin.v1.Roles/DeleteRoles', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (50, 'p', 'admin', '/api.admin.v1.Roles/GetRoles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (49, 'p', 'admin', '/api.admin.v1.Roles/ListRoles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (52, 'p', 'admin', '/api.admin.v1.Roles/UpdateRoles', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (140, 'p', 'admin', '/api.admin.v1.Sensitive/BatchDeleteSensitive', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (141, 'p', 'admin', '/api.admin.v1.Sensitive/CreateSensitive', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (142, 'p', 'admin', '/api.admin.v1.Sensitive/DeleteSensitive', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (143, 'p', 'admin', '/api.admin.v1.Sensitive/DisableSensitive', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (144, 'p', 'admin', '/api.admin.v1.Sensitive/EnableSensitive', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (145, 'p', 'admin', '/api.admin.v1.Sensitive/ListSensitive', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (146, 'p', 'admin', '/api.admin.v1.Sensitive/UpdateSensitive', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (46, 'p', 'admin', '/api.admin.v1.SysPost/CreatePost', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (48, 'p', 'admin', '/api.admin.v1.SysPost/DeletePost', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (44, 'p', 'admin', '/api.admin.v1.SysPost/ListPost', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (47, 'p', 'admin', '/api.admin.v1.SysPost/UpdatePost', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (12, 'p', 'admin', '/api.admin.v1.Sysuser/Auth', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (13, 'p', 'admin', '/api.admin.v1.Sysuser/ChangeStatus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1, 'p', 'admin', '/api.admin.v1.Sysuser/ChangeStatus', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (8, 'p', 'admin', '/api.admin.v1.Sysuser/CreateSysuser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (16, 'p', 'admin', '/api.admin.v1.Sysuser/DeleteSysuser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (2, 'p', 'admin', '/api.admin.v1.Sysuser/DeleteSysuser:userId', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (6, 'p', 'admin', '/api.admin.v1.Sysuser/GetPostInit', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (5, 'p', 'admin', '/api.admin.v1.Sysuser/GetSysuser', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (11, 'p', 'admin', '/api.admin.v1.Sysuser/GetUserGoogleSecret', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (7, 'p', 'admin', '/api.admin.v1.Sysuser/GetUserRolePost', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (15, 'p', 'admin', '/api.admin.v1.Sysuser/ListSysuser', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (14, 'p', 'admin', '/api.admin.v1.Sysuser/Logout', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (3, 'p', 'admin', '/api.admin.v1.Sysuser/UpdateAvatar', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (4, 'p', 'admin', '/api.admin.v1.Sysuser/UpdatePassword', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (9, 'p', 'admin', '/api.admin.v1.Sysuser/UpdateSysuser', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (102, 'p', 'admin', '/api.admin.v1.User/AllowAddFriend', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (103, 'p', 'admin', '/api.admin.v1.User/AllowCreateGroup', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (104, 'p', 'admin', '/api.admin.v1.User/CreateUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (105, 'p', 'admin', '/api.admin.v1.User/DeleteUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (106, 'p', 'admin', '/api.admin.v1.User/FreezeAccount', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (107, 'p', 'admin', '/api.admin.v1.User/FriendMute', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (108, 'p', 'admin', '/api.admin.v1.User/GetMuteInfo', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (109, 'p', 'admin', '/api.admin.v1.User/GetPaginationUsers', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (111, 'p', 'admin', '/api.admin.v1.User/GetUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (110, 'p', 'admin', '/api.admin.v1.User/GetUserAccountInfo', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (112, 'p', 'admin', '/api.admin.v1.User/GroupMute', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (113, 'p', 'admin', '/api.admin.v1.User/ProhibitAddFriend', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (114, 'p', 'admin', '/api.admin.v1.User/ProhibitCreateGroup', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (115, 'p', 'admin', '/api.admin.v1.User/RemoveFriendMute', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (116, 'p', 'admin', '/api.admin.v1.User/RemoveGroupMute', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (117, 'p', 'admin', '/api.admin.v1.User/UnFreezeAccount', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (118, 'p', 'admin', '/api.admin.v1.User/UpdateSelfSignature', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (119, 'p', 'admin', '/api.admin.v1.User/UpdateUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (120, 'p', 'admin', '/api.admin.v1.User/UpdateUserAvatar', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (121, 'p', 'admin', '/api.admin.v1.User/UpdateUserGender', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (122, 'p', 'admin', '/api.admin.v1.User/UpdateUserNickname', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (77, 'p', 'admin', '/job/changeStatus', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (72, 'p', 'admin', '/log/logJob/:logId', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (71, 'p', 'admin', '/log/logJob/all', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (70, 'p', 'admin', '/log/logJob/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (65, 'p', 'admin', '/log/logLogin/:infoId', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (66, 'p', 'admin', '/log/logLogin/all', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (64, 'p', 'admin', '/log/logLogin/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (68, 'p', 'admin', '/log/logOper/:operId', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (69, 'p', 'admin', '/log/logOper/all', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (67, 'p', 'admin', '/log/logOper/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (29, 'p', 'admin', '/system/dict/type/export', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (37, 'p', 'admin', '/system/menu/menuRole', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (74, 'p', 'admin', '/system/notice', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (75, 'p', 'admin', '/system/notice', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (76, 'p', 'admin', '/system/notice/:noticeId', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (73, 'p', 'admin', '/system/notice/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (45, 'p', 'admin', '/system/post/:postId', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (56, 'p', 'admin', '/system/role/export', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (10, 'p', 'admin', '/system/user/export', 'GET', '', '', '');

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'jwt',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------

-- ----------------------------
-- Table structure for log_jobs
-- ----------------------------
DROP TABLE IF EXISTS `log_jobs`;
CREATE TABLE `log_jobs`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '任务名称',
  `job_group` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分组',
  `entry_id` int(11) NOT NULL COMMENT '任务id',
  `invoke_target` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '调用方法',
  `log_info` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '日志信息',
  `status` tinyint(2) NOT NULL DEFAULT 1 COMMENT '1=正常 2=异常',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of log_jobs
-- ----------------------------

-- ----------------------------
-- Table structure for log_logins
-- ----------------------------
DROP TABLE IF EXISTS `log_logins`;
CREATE TABLE `log_logins`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `status` tinyint(2) NOT NULL DEFAULT 1 COMMENT '1=正常 2=异常',
  `ipaddr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'ip地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '归属地',
  `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '浏览器',
  `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '系统',
  `platform` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '固件',
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新者',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息',
  `login_time` timestamp NULL DEFAULT NULL COMMENT '登录时间',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of log_logins
-- ----------------------------

-- ----------------------------
-- Table structure for log_opers
-- ----------------------------
DROP TABLE IF EXISTS `log_opers`;
CREATE TABLE `log_opers`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '操作的模块',
  `business_type` tinyint(2) NOT NULL DEFAULT 0 COMMENT '0其它 1新增 2修改 3删除',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作url',
  `method` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求方法',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作人员',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户id',
  `ip` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作IP',
  `agent` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '代理',
  `latency` int(11) NOT NULL DEFAULT 0 COMMENT '延迟',
  `resp` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求参数',
  `status` tinyint(2) NOT NULL DEFAULT 1 COMMENT '1=正常 2=异常',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '错误信息',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of log_opers
-- ----------------------------

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'api组',
  `method` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '方法',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 215 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES (2, '/api.admin.v1.Sysuser/ChangeStatus', '修改用户状态', 'user', 'PUT', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (3, '/api.admin.v1.Sysuser/DeleteSysuser:userId', '删除用户信息', 'user', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (4, '/api.admin.v1.Dept/ListDept', '获取部门列表', 'dept', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (5, '/api.admin.v1.Dept/ListDept:deptId', '获取部门信息', 'dept', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (6, '/api.admin.v1.Sysuser/UpdateAvatar', '上传头像', 'user', 'POST', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (7, '/api.admin.v1.Sysuser/UpdatePassword', '修改密码', 'user', 'PUT', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (8, '/api.admin.v1.Sysuser/GetSysuser', '通过id获取用户信息', 'user', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (9, '/api.admin.v1.Sysuser/GetPostInit', '获取初始化角色岗位信息(添加用户初始化)', 'user', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (10, '/api.admin.v1.Sysuser/GetUserRolePost', '获取用户角色岗位信息', 'user', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (11, '/api.admin.v1.Sysuser/CreateSysuser', '添加用户信息', 'user', 'POST', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (12, '/api.admin.v1.Sysuser/UpdateSysuser', '修改用户信息', 'user', 'PUT', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (13, '/system/user/export', '导出用户信息', 'user', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (14, '/api.admin.v1.Dept/RoleDeptTreeSelect', '获取角色部门树', 'dept', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (15, '/api.admin.v1.Dept/GetDeptTree', '获取所有部门树', 'dept', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (16, '/api.admin.v1.Dept/CreateDept', '添加部门信息', 'dept', 'POST', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (17, '/api.admin.v1.Dept/UpdateDept', '修改部门信息', 'dept', 'PUT', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (18, '/api.admin.v1.Dept/DeleteDept', '删除部门信息', 'dept', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (25, '/api.admin.v1.DictType/ListDictType', '获取字典类型分页列表', 'dict', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (26, '/api.admin.v1.DictType/GetDictType', '获取字典类型信息', 'dict', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (27, '/api.admin.v1.DictType/CreateDictType', '添加字典类型信息', 'dict', 'POST', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (28, '/api.admin.v1.DictType/UpdateDictType', '修改字典类型信息', 'dict', 'PUT', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (29, '/api.admin.v1.DictType/DeleteDictType', '删除字典类型信息', 'dict', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (30, '/system/dict/type/export', '导出字典类型信息', 'dict', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (31, '/api.admin.v1.DictData/ListDictData', '获取字典数据分页列表', 'dict', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (32, '/api.admin.v1.DictData/GetDictData1', '获取字典数据列表通过字典类型', 'dict', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (33, '/api.admin.v1.DictData/GetDictData', '获取字典数据信息', 'dict', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (34, '/api.admin.v1.DictData/CreateDictData', '添加字典数据信息', 'dict', 'POST', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (35, '/api.admin.v1.DictData/UpdateDictData', '修改字典数据信息', 'dict', 'PUT', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (36, '/api.admin.v1.DictData/DeleteDictData', '删除字典数据信息', 'dict', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (37, '/api.admin.v1.Menus/GetMenusTree', '获取菜单树', 'menu', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (38, '/system/menu/menuRole', '获取角色菜单', 'menu', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (39, '/api.admin.v1.Menus/RoleMenuTreeSelect', '获取角色菜单树', 'menu', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (41, '/api.admin.v1.Menus/ListMenus', '获取菜单列表', 'menu', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (42, '/api.admin.v1.Menus/GetMenus', '获取菜单信息', 'menu', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (43, '/api.admin.v1.Menus/CreateMenus', '添加菜单信息', 'menu', 'POST', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (44, '/api.admin.v1.Menus/UpdateMenus', '修改菜单信息', 'menu', 'PUT', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (45, '/api.admin.v1.Menus/DeleteMenus', '删除菜单信息', 'menu', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (46, '/api.admin.v1.SysPost/ListPost', '获取岗位分页列表', 'post', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (47, '/system/post/:postId', '获取岗位信息', 'post', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (48, '/api.admin.v1.SysPost/CreatePost', '添加岗位信息', 'post', 'POST', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (49, '/api.admin.v1.SysPost/UpdatePost', '修改岗位信息', 'post', 'PUT', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (50, '/api.admin.v1.SysPost/DeletePost', '删除岗位信息', 'post', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (51, '/api.admin.v1.Roles/ListRoles', '获取角色分页列表', 'role', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (52, '/api.admin.v1.Roles/GetRoles', '获取角色信息', 'role', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (53, '/api.admin.v1.Roles/CreateRoles', '添加角色信息', 'role', 'POST', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (54, '/api.admin.v1.Roles/UpdateRoles', '修改角色信息', 'role', 'PUT', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (55, '/api.admin.v1.Roles/DeleteRoles', '删除角色信息', 'role', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (56, '/api.admin.v1.Roles/ChangeRoleStatus', '修改角色状态', 'role', 'PUT', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (57, '/api.admin.v1.Roles/DataScope', '修改角色部门权限', 'role', 'PUT', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (58, '/system/role/export', '导出角色信息', 'role', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (59, '/api.admin.v1.Api/ListApi', '获取api分页列表1', 'api', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (60, '/api.admin.v1.Api/AllApi', '获取所有api', 'api', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (61, '/api.admin.v1.Api/GetPolicyPathByRoleKey', '获取角色拥有的api权限', 'api', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (62, '/api.admin.v1.Api/GetApi', '获取api信息', 'api', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (63, '/api.admin.v1.Api/CreateApi', '添加api信息', 'api', 'POST', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (64, '/api.admin.v1.Api/UpdateApi', '修改api信息', 'api', 'PUT', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (65, '/api.admin.v1.Api/DeleteApi', '删除api信息', 'api', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (66, '/log/logLogin/list', '获取登录日志', 'log', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (67, '/log/logLogin/:infoId', '删除日志', 'log', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (68, '/log/logLogin/all', '清空所有', 'log', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (69, '/log/logOper/list', '操作日志列表', 'log', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (70, '/log/logOper/:operId', '删除', 'log', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (71, '/log/logOper/all', '清空', 'log', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (79, '/log/logJob/list', '任务日志列表', 'log', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (80, '/log/logJob/all', '清空任务日志', 'log', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (81, '/log/logJob/:logId', '删除任务日志', 'log', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (82, '/system/notice/list', '获取通知分页列表', 'notice', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (83, '/system/notice', '添加通知信息', 'notice', 'POST', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (84, '/system/notice', '修改通知信息', 'notice', 'PUT', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (85, '/system/notice/:noticeId', '删除通知信息', 'notice', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (86, '/job/changeStatus', '修改状态', 'job', 'PUT', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (119, '/api.admin.v1.Sysuser/GetUserGoogleSecret', '获取用户的 Google 验证码密钥', 'user', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (120, '/api.admin.v1.Sysuser/Auth', '获取用户授权信息', 'user', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (121, '/api.admin.v1.Sysuser/ChangeStatus', '用户更换状态', 'user', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (122, '/api.admin.v1.Sysuser/Logout', '用户退出', 'user', 'POST', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (123, '/api.admin.v1.Sysuser/ListSysuser', '获取用户列表', 'user', 'GET', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);
INSERT INTO `sys_apis` VALUES (124, '/api.admin.v1.Sysuser/DeleteSysuser', '删除用户', 'user', 'DELETE', '2023-09-07 16:33:04', '2023-09-07 16:33:20', NULL);

-- ----------------------------
-- Table structure for sys_depts
-- ----------------------------
DROP TABLE IF EXISTS `sys_depts`;
CREATE TABLE `sys_depts`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `parent_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '上级部门',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '部门路径',
  `dept_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '部门名称',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `leader` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '状态 1=正常 2-冻结',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '修改人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_depts
-- ----------------------------
INSERT INTO `sys_depts` VALUES (1, 0, '/0/1', '熊猫科技', 0, 'xm', '18353366836', '342@qq.com', 1, '', 'admin', '2021-12-01 17:31:53', '2023-09-04 11:21:50', NULL);
INSERT INTO `sys_depts` VALUES (2, 1, '/0/1/2', '研发部', 1, 'panda', '18353366543', 'ewr@qq.com', 0, '', 'admin', '2021-12-01 17:37:43', '2023-09-04 11:51:00', NULL);
INSERT INTO `sys_depts` VALUES (3, 2, '/0/1/2/3', '营销部', 2, 'panda', '18353333333', '342@qq.com', 1, '', 'admin', '2021-12-24 10:46:24', '2023-09-04 11:45:40', NULL);
INSERT INTO `sys_depts` VALUES (4, 0, '/0/4', 'GO', 0, 'leaderGO', '19312334122', 'email@email.com   ', 1, '', 'admin', '2023-08-23 12:14:41', '2023-08-23 12:15:41', '2023-08-23 12:16:48');
INSERT INTO `sys_depts` VALUES (5, 0, '/0/4', '', 0, '', '', '', 1, '', '', '2023-09-04 11:36:04', '2023-09-04 11:36:04', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT,
  `dict_sort` int(11) NULL DEFAULT NULL COMMENT '排序',
  `dict_label` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `dict_value` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '值',
  `dict_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典类型',
  `status` tinyint(2) NULL DEFAULT NULL COMMENT '状态（0正常 1停用）',
  `css_class` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'CssClass',
  `list_class` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ListClass',
  `is_default` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'IsDefault',
  `create_by` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `remark` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 0, '男', '0', 'sys_user_sex', 0, '', '', '', 'admin', '', '男', '2021-11-30 14:58:18', '2021-11-30 14:58:18', NULL);
INSERT INTO `sys_dict_data` VALUES (2, 1, '女', '1', 'sys_user_sex', 0, '', '', '', 'admin', '', '女生', '2021-11-30 15:09:11', '2021-11-30 15:10:28', NULL);
INSERT INTO `sys_dict_data` VALUES (3, 2, '未知', '2', 'sys_user_sex', 0, '', '', '', 'admin', '', '未知', '2021-11-30 15:09:11', '2021-11-30 15:10:28', NULL);
INSERT INTO `sys_dict_data` VALUES (4, 0, '正常', '1', 'sys_normal_disable', 0, '', '', '', 'admin', '', '', '2021-12-01 15:58:50', '2021-12-01 15:58:50', NULL);
INSERT INTO `sys_dict_data` VALUES (5, 1, '停用', '2', 'sys_normal_disable', 0, '', '', '', 'admin', '', '', '2021-12-01 15:59:08', '2021-12-01 15:59:08', NULL);
INSERT INTO `sys_dict_data` VALUES (6, 0, '目录', 'M', 'sys_menu_type', 0, '', '', '', 'admin', '', '', '2021-12-02 09:49:12', '2021-12-02 09:49:12', NULL);
INSERT INTO `sys_dict_data` VALUES (7, 1, '菜单', 'C', 'sys_menu_type', 0, '', '', '', 'admin', '', '', '2021-12-02 09:49:35', '2021-12-02 09:49:52', NULL);
INSERT INTO `sys_dict_data` VALUES (8, 2, '按钮', 'F', 'sys_menu_type', 0, '', '', '', 'admin', '', '', '2021-12-02 09:49:35', '2021-12-02 09:49:35', NULL);
INSERT INTO `sys_dict_data` VALUES (9, 0, '显示', '0', 'sys_show_hide', 0, '', '', '', 'admin', '', '', '2021-12-02 09:56:40', '2021-12-02 09:56:40', NULL);
INSERT INTO `sys_dict_data` VALUES (10, 0, '隐藏', '1', 'sys_show_hide', 0, '', '', '', 'admin', '', '', '2021-12-02 09:56:52', '2021-12-02 09:56:52', NULL);
INSERT INTO `sys_dict_data` VALUES (11, 0, '是', '0', 'sys_num_yes_no', 0, '', '', '', 'admin', '', '', '2021-12-02 10:16:16', '2021-12-02 10:16:16', NULL);
INSERT INTO `sys_dict_data` VALUES (12, 1, '否', '1', 'sys_num_yes_no', 0, '', '', '', 'admin', '', '', '2021-12-02 10:16:26', '2021-12-02 10:16:26', NULL);
INSERT INTO `sys_dict_data` VALUES (13, 0, '是', '0', 'sys_yes_no', 0, '', '', '', 'admin', '', '', '2021-12-04 13:48:15', '2021-12-04 13:48:15', NULL);
INSERT INTO `sys_dict_data` VALUES (14, 0, '否', '1', 'sys_yes_no', 0, '', '', '', 'admin', '', '', '2021-12-04 13:48:21', '2021-12-04 13:48:21', NULL);
INSERT INTO `sys_dict_data` VALUES (15, 0, '创建(POST)', 'POST', 'sys_method_api', 0, '', '', '', 'admin', '', '', '2021-12-08 17:22:05', '2021-12-09 09:29:52', NULL);
INSERT INTO `sys_dict_data` VALUES (16, 1, '查询(GET)', 'GET', 'sys_method_api', 0, '', '', '', 'admin', '', '', '2021-12-08 17:22:24', '2021-12-09 09:29:59', NULL);
INSERT INTO `sys_dict_data` VALUES (17, 2, '修改(PUT)', 'PUT', 'sys_method_api', 0, '', '', '', 'admin', '', '', '2021-12-08 17:22:40', '2021-12-09 09:30:06', NULL);
INSERT INTO `sys_dict_data` VALUES (18, 3, '删除(DELETE)', 'DELETE', 'sys_method_api', 0, '', '', '', 'admin', '', '', '2021-12-08 17:22:54', '2021-12-09 09:30:13', NULL);
INSERT INTO `sys_dict_data` VALUES (19, 0, '成功', '0', 'sys_common_status', 0, '', '', '', 'admin', '', '', '2021-12-17 11:01:52', '2021-12-17 11:01:52', NULL);
INSERT INTO `sys_dict_data` VALUES (20, 0, '失败', '1', 'sys_common_status', 0, '', '', '', 'admin', '', '', '2021-12-17 11:02:08', '2021-12-17 11:02:08', NULL);
INSERT INTO `sys_dict_data` VALUES (21, 0, '其他', '0', 'sys_oper_type', 0, '', '', '', 'admin', '', '', '2021-12-17 11:30:07', '2021-12-17 11:30:07', NULL);
INSERT INTO `sys_dict_data` VALUES (22, 0, '新增', '1', 'sys_oper_type', 0, '', '', '', 'admin', '', '', '2021-12-17 11:30:21', '2021-12-17 11:30:21', NULL);
INSERT INTO `sys_dict_data` VALUES (23, 0, '修改', '2', 'sys_oper_type', 0, '', '', '', 'admin', '', '', '2021-12-17 11:30:32', '2021-12-17 11:30:32', NULL);
INSERT INTO `sys_dict_data` VALUES (24, 0, '删除', '3', 'sys_oper_type', 0, '', '', '', 'admin', '', '', '2021-12-17 11:30:40', '2021-12-17 11:30:40', NULL);
INSERT INTO `sys_dict_data` VALUES (25, 0, '默认', 'DEFAULT', 'sys_job_group', 0, '', '', '', 'panda', '', '', '2021-12-24 15:15:31', '2021-12-24 15:15:31', NULL);
INSERT INTO `sys_dict_data` VALUES (26, 1, '系统', 'SYSTEM', 'sys_job_group', 0, '', '', '', 'panda', '', '', '2021-12-24 15:15:50', '2021-12-24 15:15:50', NULL);
INSERT INTO `sys_dict_data` VALUES (27, 0, '发布通知', '1', 'sys_notice_type', 0, '', '', '', 'panda', '', '', '2021-12-26 15:24:07', '2021-12-26 15:24:07', NULL);
INSERT INTO `sys_dict_data` VALUES (28, 0, '任免通知', '2', 'sys_notice_type', 0, '', '', '', 'panda', '', '', '2021-12-26 15:24:18', '2021-12-26 15:24:18', NULL);
INSERT INTO `sys_dict_data` VALUES (29, 0, '事物通知', '3', 'sys_notice_type', 0, '', '', '', 'panda', '', '', '2021-12-26 15:24:46', '2021-12-26 15:24:46', NULL);
INSERT INTO `sys_dict_data` VALUES (30, 0, '审批通知', '4', 'sys_notice_type', 0, '', '', '', 'panda', '', '', '2021-12-26 15:25:08', '2021-12-26 15:25:08', NULL);
INSERT INTO `sys_dict_data` VALUES (31, 0, '阿里云', '0', 'res_oss_category', 0, '', '', '', 'panda', '', '', '2022-01-13 15:44:01', '2022-01-13 15:44:01', NULL);
INSERT INTO `sys_dict_data` VALUES (32, 1, '七牛云', '1', 'res_oss_category', 0, '', '', '', 'panda', '', '', '2022-01-13 15:44:18', '2022-01-13 15:44:18', NULL);
INSERT INTO `sys_dict_data` VALUES (33, 2, '腾讯云', '2', 'res_oss_category', 0, '', '', '', 'panda', '', '', '2022-01-13 15:44:39', '2022-01-13 15:44:39', NULL);
INSERT INTO `sys_dict_data` VALUES (34, 0, '阿里云', '0', 'res_sms_category', 0, '', '', '', 'panda', '', '', '2022-01-13 15:47:30', '2022-01-13 15:47:30', NULL);
INSERT INTO `sys_dict_data` VALUES (35, 1, '腾讯云', '1', 'res_sms_category', 0, '', '', '', 'panda', '', '', '2022-01-13 15:47:39', '2022-01-13 15:47:39', NULL);
INSERT INTO `sys_dict_data` VALUES (36, 0, '163邮箱', '0', 'res_mail_category', 0, '', '', '', 'panda', '', '', '2022-01-14 15:43:42', '2022-01-14 15:43:42', NULL);
INSERT INTO `sys_dict_data` VALUES (37, 0, 'qq邮箱', '1', 'res_mail_category', 0, '', '', '', 'panda', '', '', '2022-01-14 15:44:08', '2022-01-14 15:44:08', NULL);
INSERT INTO `sys_dict_data` VALUES (38, 0, '企业邮箱', '2', 'res_mail_category', 0, '', '', '', 'panda', '', '', '2022-01-14 15:44:20', '2022-01-14 15:44:20', NULL);

-- ----------------------------
-- Table structure for sys_dict_types
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_types`;
CREATE TABLE `sys_dict_types`  (
  `dict_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `dict_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `dict_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型',
  `status` tinyint(2) NULL DEFAULT NULL COMMENT '状态',
  `create_by` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `remark` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`dict_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 33 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dict_types
-- ----------------------------
INSERT INTO `sys_dict_types` VALUES (1, '用户性别', 'sys_user_sex', 0, 'admin', '', '性别列表', '2021-11-30 14:02:52', '2021-11-30 14:07:55', '2021-11-30 14:11:54');
INSERT INTO `sys_dict_types` VALUES (2, '用户性别', 'sys_user_sex', 0, 'admin', '', '用户性别列表', '2021-11-30 14:12:33', '2021-11-30 14:12:33', '2021-11-30 14:14:19');
INSERT INTO `sys_dict_types` VALUES (3, '的心', 'sfd', 0, 'admin', '', 'fs', '2021-11-30 14:13:22', '2021-11-30 14:13:22', '2021-11-30 14:14:19');
INSERT INTO `sys_dict_types` VALUES (4, '用户性别', 'sys_user_sex', 0, 'admin', '', '性别字典', '2021-11-30 14:15:25', '2021-11-30 14:15:25', NULL);
INSERT INTO `sys_dict_types` VALUES (5, 'df', 'da', 0, 'admin', '', 'sd', '2021-11-30 15:54:33', '2021-11-30 15:54:33', '2021-11-30 15:54:40');
INSERT INTO `sys_dict_types` VALUES (6, '系统开关', 'sys_normal_disable', 0, 'admin', '', '开关列表', '2021-12-01 15:57:58', '2021-12-01 15:57:58', NULL);
INSERT INTO `sys_dict_types` VALUES (7, '菜单类型', 'sys_menu_type', 0, 'admin', '', '菜单类型列表', '2021-12-02 09:48:48', '2021-12-02 09:56:12', NULL);
INSERT INTO `sys_dict_types` VALUES (8, '菜单状态', 'sys_show_hide', 0, 'admin', '', '菜单状态列表', '2021-12-02 09:55:59', '2021-12-02 09:55:59', NULL);
INSERT INTO `sys_dict_types` VALUES (9, '数字是否', 'sys_num_yes_no', 0, 'admin', '', '数字是否列表', '2021-12-02 10:13:29', '2021-12-02 10:13:40', '2021-12-02 10:15:07');
INSERT INTO `sys_dict_types` VALUES (10, '数字是否', 'sys_num_yes_no', 0, 'admin', '', '数字是否', '2021-12-02 10:13:29', '2021-12-02 10:13:29', NULL);
INSERT INTO `sys_dict_types` VALUES (11, '状态是否', 'sys_yes_no', 0, 'admin', '', '状态是否', '2021-12-04 13:47:57', '2021-12-04 13:47:57', NULL);
INSERT INTO `sys_dict_types` VALUES (12, '网络请求方法', 'sys_method_api', 0, 'admin', '', '网络请求方法列表', '2021-12-08 17:21:27', '2021-12-08 17:21:27', NULL);
INSERT INTO `sys_dict_types` VALUES (13, '成功失败', 'sys_common_status', 0, 'admin', '', '是否成功失败', '2021-12-17 10:10:03', '2021-12-17 10:10:03', NULL);
INSERT INTO `sys_dict_types` VALUES (27, '操作分类', 'sys_oper_type', 0, 'admin', '', '操作分类列表', '2021-12-17 11:29:50', '2021-12-17 11:29:50', NULL);
INSERT INTO `sys_dict_types` VALUES (28, '任务组', 'sys_job_group', 0, 'panda', '', '系统任务，开机自启', '2021-12-24 15:14:56', '2021-12-24 15:14:56', NULL);
INSERT INTO `sys_dict_types` VALUES (29, '通知类型', 'sys_notice_type', 0, 'panda', '', '通知类型列表', '2021-12-26 15:23:26', '2021-12-26 15:23:26', NULL);
INSERT INTO `sys_dict_types` VALUES (30, 'oss分类', 'res_oss_category', 0, 'panda', '', 'oss分类列表', '2022-01-13 15:43:29', '2022-01-13 15:43:29', NULL);
INSERT INTO `sys_dict_types` VALUES (31, 'sms分类', 'res_sms_category', 0, 'panda', '', 'sms分类列表', '2021-12-26 15:23:26', '2022-01-13 15:47:13', NULL);
INSERT INTO `sys_dict_types` VALUES (32, 'mail分类', 'res_mail_category', 0, 'panda', '', 'mail分类列表', '2022-01-14 15:43:17', '2022-01-14 15:43:17', NULL);

-- ----------------------------
-- Table structure for sys_discovery
-- ----------------------------
DROP TABLE IF EXISTS `sys_discovery`;
CREATE TABLE `sys_discovery`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '名称',
  `picture` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '发现页图片',
  `rank` int(11) NULL DEFAULT 1 COMMENT '排序',
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '发现页链接',
  `status` int(11) NULL DEFAULT 1 COMMENT '发现页状态, 1-正常，2-异常',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_discovery
-- ----------------------------
INSERT INTO `sys_discovery` VALUES (8, '啊实', 'http://oss.nfdx.xyz/files/82cf620a-b726-3405-bbe0-dad3e0d3755b.jpg', 1, '发', 1);
INSERT INTO `sys_discovery` VALUES (9, '发广告', 'http://oss.nfdx.xyz/files/74028856-4d7a-3922-9ea5-6c4430e94f72.jpg', 1, '1231321', 1);

-- ----------------------------
-- Table structure for sys_jobs
-- ----------------------------
DROP TABLE IF EXISTS `sys_jobs`;
CREATE TABLE `sys_jobs`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `job_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '任务名称',
  `job_group` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '任务组',
  `job_type` tinyint(2) NULL DEFAULT 0 COMMENT '任务类型',
  `cron_expression` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'cron表达式',
  `invoke_target` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '调用目标',
  `args` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '目标参数',
  `misfire_policy` tinyint(2) NULL DEFAULT 0 COMMENT '执行策略',
  `concurrent` tinyint(2) NULL DEFAULT 2 COMMENT '是否并发 1=是 2=否',
  `status` tinyint(2) NULL DEFAULT 1 COMMENT '1=正常 2=异常',
  `entry_id` int(11) NULL DEFAULT 0 COMMENT 'job启动时返回的id',
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_jobs
-- ----------------------------
INSERT INTO `sys_jobs` VALUES (1, 'testcron', 'SYSTEM', 2, ' 0/10 * * * * ?', 'cronHandle', 'aaa', 0, 2, 1, 1, 'panda', '', '2021-12-24 16:06:09', '2022-01-22 08:11:24', NULL);

-- ----------------------------
-- Table structure for sys_menu_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu_btns`;
CREATE TABLE `sys_menu_btns`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单ID',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '按钮关键key',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '按钮描述',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menu_btns
-- ----------------------------

-- ----------------------------
-- Table structure for sys_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_menus`;
CREATE TABLE `sys_menus`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '附加属性',
  `parent_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '父级菜单id',
  `sort` int(4) NOT NULL DEFAULT 0 COMMENT '排序',
  `icon` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `is_iframe` tinyint(2) NOT NULL DEFAULT 2 COMMENT '是否为内嵌 1=是 2=否',
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '超链接',
  `menu_type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '菜单类型（1目录 2菜单 3按钮）',
  `hidden` tinyint(2) NOT NULL DEFAULT 1 COMMENT '显示状态（0显示 1隐藏）',
  `keep_alive` tinyint(2) NOT NULL DEFAULT 1 COMMENT '是否缓存组件状态（1是 2否）',
  `is_affix` tinyint(2) NOT NULL DEFAULT 1 COMMENT '是否固定在 tagsView 栏上（1是 2否）',
  `permission` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限标识',
  `status` tinyint(2) NOT NULL DEFAULT 1 COMMENT '菜单状态 1=正常 2=停用',
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新人',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 52 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
INSERT INTO `sys_menus` VALUES (1, '系统设置', '', 0, 1, 'elementSetting', '/system', 'Layout', 1, '', 'M', 1, 1, 1, '', 1, '', 'admin', '', '2021-12-02 11:04:08', '2023-09-07 10:29:24', NULL);
INSERT INTO `sys_menus` VALUES (3, '用户管理', '', 1, 1, 'elementUser', '/system/user', '/system/user/index', 1, '', 'C', 0, 1, 1, 'system:user:list', 0, 'admin', 'panda', '', '2021-12-02 14:07:56', '2021-12-28 13:32:44', NULL);
INSERT INTO `sys_menus` VALUES (4, '添加用户', '', 3, 1, '', '', '', 2, '', 'F', 1, 1, 1, 'system:user:add', 1, '', 'admin', '', '2021-12-03 13:36:33', '2023-08-31 15:41:25', NULL);
INSERT INTO `sys_menus` VALUES (5, '编辑用户', '', 3, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:user:edit', 0, 'admin', '', '', '2021-12-03 13:48:13', '2021-12-03 13:48:13', NULL);
INSERT INTO `sys_menus` VALUES (6, '角色管理', '', 1, 2, 'elementUserFilled', '/system/role', '/system/role/index', 1, '', 'C', 0, 1, 1, 'system:role:list', 0, '', 'panda', '', '2021-12-03 13:51:55', '2022-07-16 10:23:21', NULL);
INSERT INTO `sys_menus` VALUES (7, '菜单管理', '', 1, 2, 'iconfont icon-juxingkaobei', '/system/menu', '/system/menu/index', 1, '', 'C', 0, 1, 1, 'system:menu:list', 0, 'admin', 'panda', '', '2021-12-03 13:54:44', '2021-12-28 13:33:19', NULL);
INSERT INTO `sys_menus` VALUES (8, '组织管理', '', 1, 3, 'iconfont icon-jiliandongxuanzeqi', '/system/dept', '/system/dept/index', 1, '', 'C', 1, 1, 1, 'system:dept:list', 1, '', 'admin', '', '2021-12-03 13:58:36', '2023-09-04 14:39:44', NULL);
INSERT INTO `sys_menus` VALUES (9, '岗位管理', '', 1, 4, 'iconfont icon-neiqianshujuchucun', '/system/post', '/system/post/index', 1, '', 'C', 1, 1, 1, 'system:post:list', 1, '', 'admin', '', '2021-12-03 13:54:44', '2023-09-07 10:16:11', NULL);
INSERT INTO `sys_menus` VALUES (10, '字典管理', '', 1, 5, 'elementCellphone', '/system/dict', '/system/dict/index', 1, '', 'C', 0, 1, 1, 'system:dict:list', 0, 'admin', 'panda', '', '2021-12-03 13:54:44', '2021-12-28 13:40:50', NULL);
INSERT INTO `sys_menus` VALUES (11, '参数管理', '', 1, 6, 'elementDocumentCopy', '/system/config', '/system/config/index', 1, '', 'C', 0, 1, 1, 'system:config:list', 0, 'admin', 'panda', '', '2021-12-03 13:54:44', '2021-12-28 13:41:05', NULL);
INSERT INTO `sys_menus` VALUES (12, '个人中心', '', 5, 10, 'elementAvatar', '/personal', '/personal/index', 1, '', 'M', 1, 1, 1, '', 1, '', 'admin', '', '2021-12-03 14:12:43', '2023-09-04 10:52:30', NULL);
INSERT INTO `sys_menus` VALUES (13, '添加配置', '', 11, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:config:add', 0, 'admin', '', '', '2021-12-06 17:19:19', '2021-12-06 17:19:19', NULL);
INSERT INTO `sys_menus` VALUES (14, '修改配置', '', 11, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:config:edit', 0, 'admin', '', '', '2021-12-06 17:20:30', '2021-12-06 17:20:30', NULL);
INSERT INTO `sys_menus` VALUES (15, '删除配置', '', 11, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:config:delete', 0, 'admin', '', '', '2021-12-06 17:23:52', '2021-12-06 17:23:52', NULL);
INSERT INTO `sys_menus` VALUES (16, '导出配置', '', 11, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:config:export', 0, 'admin', '', '', '2021-12-06 17:24:41', '2021-12-06 17:24:41', NULL);
INSERT INTO `sys_menus` VALUES (17, '新增角色', '', 6, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:role:add', 0, 'admin', '', '', '2021-12-06 17:43:35', '2021-12-06 17:43:35', NULL);
INSERT INTO `sys_menus` VALUES (18, '删除角色', '', 6, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:role:delete', 0, 'admin', '', '', '2021-12-06 17:44:10', '2021-12-06 17:44:10', NULL);
INSERT INTO `sys_menus` VALUES (19, '修改角色', '', 6, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:role:edit', 0, 'admin', '', '', '2021-12-06 17:44:48', '2021-12-06 17:44:48', NULL);
INSERT INTO `sys_menus` VALUES (20, '导出角色', '', 6, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:role:export', 0, 'admin', '', '', '2021-12-06 17:45:25', '2021-12-06 17:45:25', NULL);
INSERT INTO `sys_menus` VALUES (21, '添加菜单', '', 7, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:menu:add', 0, 'admin', '', '', '2021-12-06 17:46:01', '2021-12-06 17:46:01', NULL);
INSERT INTO `sys_menus` VALUES (22, '修改菜单', '', 7, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:menu:edit', 0, 'admin', '', '', '2021-12-06 17:46:24', '2021-12-06 17:46:24', NULL);
INSERT INTO `sys_menus` VALUES (23, '删除菜单', '', 7, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:menu:delete', 0, 'admin', '', '', '2021-12-06 17:46:47', '2021-12-06 17:46:47', NULL);
INSERT INTO `sys_menus` VALUES (24, '添加部门', '', 8, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:dept:add', 0, 'admin', '', '', '2021-12-07 09:33:58', '2021-12-07 09:33:58', NULL);
INSERT INTO `sys_menus` VALUES (25, '编辑部门', '', 8, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:dept:edit', 0, 'admin', '', '', '2021-12-07 09:34:39', '2021-12-07 09:34:39', NULL);
INSERT INTO `sys_menus` VALUES (26, '删除部门', '', 8, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:dept:delete', 0, 'admin', 'admin', '', '2021-12-07 09:35:09', '2021-12-07 09:36:26', NULL);
INSERT INTO `sys_menus` VALUES (28, '添加岗位', '', 9, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:post:add', 0, 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (29, '编辑岗位', '', 9, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:post:edit', 0, 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (30, '删除岗位', '', 9, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:post:delete', 0, 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (31, '导出岗位', '', 9, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:post:export', 0, 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (32, '添加字典类型', '', 10, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:dictT:add', 0, 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (33, '编辑字典类型', '', 10, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:dictT:edit', 0, 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (34, '删除字典类型', '', 10, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:dictT:delete', 0, 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (35, '导出字典类型', '', 10, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:dictT:export', 0, 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (36, '新增字典数据', '', 10, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:dictD:add', 0, 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (37, '修改字典数据', '', 10, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:dictD:edit', 0, 'admin', '', '', '2021-12-07 09:48:04', '2021-12-07 09:48:04', NULL);
INSERT INTO `sys_menus` VALUES (38, '删除字典数据', '', 10, 1, '', '', '', 0, '', 'F', 0, 0, 0, 'system:dictD:delete', 0, 'admin', '', '', '2021-12-07 09:48:42', '2021-12-07 09:48:42', NULL);
INSERT INTO `sys_menus` VALUES (39, 'API管理', '', 1, 2, 'iconfont icon-siweidaotu', '/system/api', '/system/api/index', 1, '', 'C', 0, 1, 1, 'system:api:list', 0, '', 'panda', '', '2021-12-09 09:09:13', '2022-07-16 10:23:42', NULL);
INSERT INTO `sys_menus` VALUES (40, '添加api', '', 39, 1, '', '/system/api', '', 0, '', 'F', 0, 0, 0, 'system:api:add', 0, 'admin', '', '', '2021-12-09 09:09:54', '2021-12-09 09:09:54', NULL);
INSERT INTO `sys_menus` VALUES (41, '编辑api', '', 39, 1, '', '/system/api', '', 0, '', 'F', 0, 0, 0, 'system:api:edit', 0, 'admin', '', '', '2021-12-09 09:10:38', '2021-12-09 09:10:38', NULL);
INSERT INTO `sys_menus` VALUES (42, '删除api', '', 39, 1, '', '/system/api', '', 0, '', 'F', 0, 0, 0, 'system:api:delete', 0, 'admin', '', '', '2021-12-09 09:11:11', '2021-12-09 09:11:11', NULL);
INSERT INTO `sys_menus` VALUES (43, 'im用户管理', '', 0, 2, 'elementUser', '', 'Layout', 2, '', 'M', 1, 1, 1, '', 1, '', 'admin', '', '2023-09-06 15:22:10', '2023-09-07 10:33:27', '2023-09-07 16:29:34');
INSERT INTO `sys_menus` VALUES (44, '产品配置', '', 0, 3, 'elementSetUp', '', 'Layout', 2, '', 'M', 1, 1, 1, '', 1, '', 'admin', '', '2023-09-06 16:49:18', '2023-09-06 17:05:34', NULL);
INSERT INTO `sys_menus` VALUES (45, '群组管理', '', 0, 4, 'elementAvatar', '', 'Layout', 2, '', 'M', 1, 1, 1, '', 1, '', 'admin', '', '2023-09-06 16:53:42', '2023-09-06 17:05:56', NULL);
INSERT INTO `sys_menus` VALUES (46, '消息管理', '', 0, 5, 'elementMessage', '/msg/index', '/msg/index', 1, '', 'M', 1, 1, 1, 'system:msg', 1, '', 'admin', '', '2023-09-06 16:54:06', '2023-09-07 11:00:07', NULL);
INSERT INTO `sys_menus` VALUES (47, '系统消息', '', 0, 6, 'elementMessageBox', '', 'Layout', 2, '', 'M', 1, 1, 1, '', 1, '', 'admin', '', '2023-09-06 16:54:37', '2023-09-06 17:03:51', NULL);
INSERT INTO `sys_menus` VALUES (48, '敏感词管理', '', 0, 7, 'elementInfoFilled', '', 'Layout', 2, '', 'M', 1, 1, 1, '', 1, '', 'admin', '', '2023-09-06 16:58:56', '2023-09-06 17:04:18', '2023-09-07 16:29:41');
INSERT INTO `sys_menus` VALUES (49, '操作日志', '', 0, 8, 'elementComment', '', 'Layout', 2, '', 'M', 1, 1, 1, '', 1, '', 'admin', '', '2023-09-06 16:59:13', '2023-09-06 17:04:42', NULL);
INSERT INTO `sys_menus` VALUES (50, '反馈管理', '', 0, 9, 'elementEdit', '', 'Layout', 2, '', 'M', 1, 1, 1, '', 1, '', 'admin', '', '2023-09-06 16:59:32', '2023-09-06 17:07:32', NULL);
INSERT INTO `sys_menus` VALUES (51, '私聊消息', '', 46, 1, 'iconfont icon-tongzhi1', '/msg/single_list', '/msg/single_list', 1, '', 'C', 0, 1, 1, 'system:msg:index', 0, '', 'admin', '', '2023-09-07 10:52:10', '2023-09-07 10:59:30', NULL);

-- ----------------------------
-- Table structure for sys_posts
-- ----------------------------
DROP TABLE IF EXISTS `sys_posts`;
CREATE TABLE `sys_posts`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `post_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '岗位名称',
  `post_code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '岗位代码',
  `sort` int(4) NOT NULL DEFAULT 0 COMMENT '岗位排序',
  `status` tinyint(2) NOT NULL DEFAULT 1 COMMENT '状态 1=正常 2=冻结',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '修改人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_posts
-- ----------------------------
INSERT INTO `sys_posts` VALUES (1, '首席执行官', 'CEO', 1, 1, '首席执行官', '', 'admin', '2021-12-02 09:21:44', '2023-09-04 09:49:50', NULL);
INSERT INTO `sys_posts` VALUES (2, '首席技术执行官', 'CTO', 2, 0, '', 'panda', '', '2021-12-02 09:21:44', '2022-07-16 17:37:42', NULL);
INSERT INTO `sys_posts` VALUES (3, 'testPost2', '111', 0, 1, 'this is a remark', '', 'admin', '2023-08-23 12:18:20', '2023-08-23 12:18:56', '2023-08-23 12:19:26');
INSERT INTO `sys_posts` VALUES (7, 'testPost2', '111', 3, 1, 'this is a remark', '', 'admin', '2023-08-31 20:54:03', '2023-09-01 16:14:55', NULL);

-- ----------------------------
-- Table structure for sys_role_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_btns`;
CREATE TABLE `sys_role_btns`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单ID',
  `btn_id` bigint(20) NOT NULL COMMENT '菜单按钮ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_btns
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role_depts
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_depts`;
CREATE TABLE `sys_role_depts`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `role_id` bigint(20) NOT NULL COMMENT '角色id',
  `dept_id` bigint(20) NOT NULL COMMENT '部门id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_depts
-- ----------------------------
INSERT INTO `sys_role_depts` VALUES (1, 1, 1);

-- ----------------------------
-- Table structure for sys_role_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menus`;
CREATE TABLE `sys_role_menus`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `role_id` bigint(20) NOT NULL COMMENT '角色id',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单id',
  `role_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5920 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_menus
-- ----------------------------
INSERT INTO `sys_role_menus` VALUES (4931, 2, 1, '管理员');
INSERT INTO `sys_role_menus` VALUES (4932, 2, 3, '管理员');
INSERT INTO `sys_role_menus` VALUES (4933, 2, 4, '管理员');
INSERT INTO `sys_role_menus` VALUES (4934, 2, 5, '管理员');
INSERT INTO `sys_role_menus` VALUES (4935, 2, 6, '管理员');
INSERT INTO `sys_role_menus` VALUES (4936, 2, 17, '管理员');
INSERT INTO `sys_role_menus` VALUES (4937, 2, 18, '管理员');
INSERT INTO `sys_role_menus` VALUES (4938, 2, 19, '管理员');
INSERT INTO `sys_role_menus` VALUES (4939, 2, 20, '管理员');
INSERT INTO `sys_role_menus` VALUES (4940, 2, 7, '管理员');
INSERT INTO `sys_role_menus` VALUES (4941, 2, 21, '管理员');
INSERT INTO `sys_role_menus` VALUES (4942, 2, 22, '管理员');
INSERT INTO `sys_role_menus` VALUES (4943, 2, 23, '管理员');
INSERT INTO `sys_role_menus` VALUES (4944, 2, 39, '管理员');
INSERT INTO `sys_role_menus` VALUES (4945, 2, 40, '管理员');
INSERT INTO `sys_role_menus` VALUES (4946, 2, 41, '管理员');
INSERT INTO `sys_role_menus` VALUES (4947, 2, 42, '管理员');
INSERT INTO `sys_role_menus` VALUES (4948, 2, 8, '管理员');
INSERT INTO `sys_role_menus` VALUES (4949, 2, 24, '管理员');
INSERT INTO `sys_role_menus` VALUES (4950, 2, 25, '管理员');
INSERT INTO `sys_role_menus` VALUES (4951, 2, 26, '管理员');
INSERT INTO `sys_role_menus` VALUES (4952, 2, 9, '管理员');
INSERT INTO `sys_role_menus` VALUES (4953, 2, 28, '管理员');
INSERT INTO `sys_role_menus` VALUES (4954, 2, 29, '管理员');
INSERT INTO `sys_role_menus` VALUES (4955, 2, 30, '管理员');
INSERT INTO `sys_role_menus` VALUES (4956, 2, 31, '管理员');
INSERT INTO `sys_role_menus` VALUES (4957, 2, 10, '管理员');
INSERT INTO `sys_role_menus` VALUES (4958, 2, 32, '管理员');
INSERT INTO `sys_role_menus` VALUES (4959, 2, 33, '管理员');
INSERT INTO `sys_role_menus` VALUES (4960, 2, 34, '管理员');
INSERT INTO `sys_role_menus` VALUES (4961, 2, 35, '管理员');
INSERT INTO `sys_role_menus` VALUES (4962, 2, 36, '管理员');
INSERT INTO `sys_role_menus` VALUES (4963, 2, 37, '管理员');
INSERT INTO `sys_role_menus` VALUES (4964, 2, 38, '管理员');
INSERT INTO `sys_role_menus` VALUES (4965, 2, 11, '管理员');
INSERT INTO `sys_role_menus` VALUES (4966, 2, 13, '管理员');
INSERT INTO `sys_role_menus` VALUES (4967, 2, 14, '管理员');
INSERT INTO `sys_role_menus` VALUES (4968, 2, 15, '管理员');
INSERT INTO `sys_role_menus` VALUES (4969, 2, 16, '管理员');
INSERT INTO `sys_role_menus` VALUES (5872, 1, 1, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5873, 1, 3, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5874, 1, 4, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5875, 1, 5, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5876, 1, 6, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5877, 1, 17, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5878, 1, 18, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5879, 1, 19, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5880, 1, 20, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5881, 1, 7, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5882, 1, 21, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5883, 1, 22, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5884, 1, 23, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5885, 1, 39, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5886, 1, 40, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5887, 1, 41, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5888, 1, 42, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5889, 1, 8, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5890, 1, 24, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5891, 1, 25, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5892, 1, 26, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5893, 1, 9, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5894, 1, 28, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5895, 1, 29, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5896, 1, 30, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5897, 1, 31, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5898, 1, 10, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5899, 1, 32, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5900, 1, 33, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5901, 1, 34, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5902, 1, 35, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5903, 1, 36, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5904, 1, 37, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5905, 1, 38, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5906, 1, 11, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5907, 1, 13, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5908, 1, 14, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5909, 1, 15, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5910, 1, 16, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5911, 1, 43, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5912, 1, 44, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5913, 1, 45, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5914, 1, 46, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5915, 1, 51, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5916, 1, 47, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5917, 1, 48, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5918, 1, 49, '超管理员');
INSERT INTO `sys_role_menus` VALUES (5919, 1, 50, '超管理员');

-- ----------------------------
-- Table structure for sys_roles
-- ----------------------------
DROP TABLE IF EXISTS `sys_roles`;
CREATE TABLE `sys_roles`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `parent_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '父角色ID',
  `role_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `status` tinyint(2) NOT NULL DEFAULT 1 COMMENT '1=正常 2=异常',
  `role_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色代码',
  `data_scope` tinyint(2) NOT NULL DEFAULT 1 COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `role_sort` int(4) NOT NULL DEFAULT 0 COMMENT '角色排序',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '默认菜单',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_roles
-- ----------------------------
INSERT INTO `sys_roles` VALUES (1, 0, '超管理员', 1, 'admin', 1, 0, '', '', '', 'admin', '2021-12-02 16:03:26', '2023-09-07 11:01:23', NULL);
INSERT INTO `sys_roles` VALUES (2, 0, '管理员', 0, 'manage', 1, 0, '', '', '', 'admin', '2021-12-19 16:06:20', '2023-09-05 11:10:50', NULL);

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `uuid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户UUID',
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名(登入)',
  `nick_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机',
  `role_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '角色id',
  `salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '盐',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `sex` tinyint(2) NOT NULL DEFAULT 0 COMMENT '性别 0-未知 1-男 2-女',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `dept_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '部门id',
  `post_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '岗位id',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `status` tinyint(2) NOT NULL DEFAULT 1 COMMENT '1=正常 2=异常',
  `role_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '多角色',
  `post_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '多岗位',
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `secret` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'google密钥',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, '1-1-1-1', 'admin', 'admin', '$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2', '18888888888', 1, '', '', 0, 'example@email.com', 3, 1, 'remark', 1, '1', '1', 'admin', 'admin', '2021-12-03 09:46:55', '2023-09-04 10:40:54', NULL, '45RNXQTW2EMJ2EOAQ26UYML2D2K2IPYT');
INSERT INTO `sys_users` VALUES (2, 'd26733e4-ee09-4d98-b462-93bae039209c', 'test2', 'test2', '$2a$10$yiQ9u0lh7wsGjchqMGVOE.lp3KO99R5nw0Kc1DWQC6THI6d.JzNP.', '13312312311', 1, '', '', 1, 'email@email.com', 2, 1, 'this is a remark2', 1, '1', '1', 'admin', 'admin', '2023-08-23 11:38:47', '2023-09-07 10:02:50', NULL, 'K6SSMXVIX6WRDBEIPX2ZDHLR6XSCEAKN');
INSERT INTO `sys_users` VALUES (3, 'b3614db9-80a8-4892-9f65-0a6e70a00a2d', 'dahe', 'dahe', '$2a$10$iCr0rC6esWA91xCiImLZ5uMxjnW45VVhFzR2e9IPVg4QKY/XhvqEu', '13777788880', 1, '', '', 0, 'dahe@gmail.com', 3, 1, 'ewtwet', 1, '1', '1', 'admin', 'admin', '2023-08-24 08:54:34', '2023-09-04 11:16:19', NULL, '5KPR5XMSMTZTE6WQBHLFXYNKA64EOBUH');

SET FOREIGN_KEY_CHECKS = 1;
