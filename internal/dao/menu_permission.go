// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"coderblog-interface/internal/dao/internal"
)

// internalMenuPermissionDao is internal type for wrapping internal DAO implements.
type internalMenuPermissionDao = *internal.MenuPermissionDao

// menuPermissionDao is the data access object for table menu_permission.
// You can define custom methods on it to extend its functionality as you wish.
type menuPermissionDao struct {
	internalMenuPermissionDao
}

var (
	// MenuPermission is globally public accessible object for table menu_permission operations.
	MenuPermission = menuPermissionDao{
		internal.NewMenuPermissionDao(),
	}
)

// Fill with you ideas below.
