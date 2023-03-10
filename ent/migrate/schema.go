// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ProfilesColumns holds the columns for the "profiles" table.
	ProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gender", Type: field.TypeString, Nullable: true, Default: "unknown"},
		{Name: "age", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt, Unique: true},
	}
	// ProfilesTable holds the schema information for the "profiles" table.
	ProfilesTable = &schema.Table{
		Name:       "profiles",
		Columns:    ProfilesColumns,
		PrimaryKey: []*schema.Column{ProfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "profiles_users_profile",
				Columns:    []*schema.Column{ProfilesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ProfilesTable,
		UsersTable,
	}
)

func init() {
	ProfilesTable.ForeignKeys[0].RefTable = UsersTable
}
