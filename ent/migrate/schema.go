// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ConversationsColumns holds the columns for the "conversations" table.
	ConversationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Size: 2147483647},
	}
	// ConversationsTable holds the schema information for the "conversations" table.
	ConversationsTable = &schema.Table{
		Name:       "conversations",
		Columns:    ConversationsColumns,
		PrimaryKey: []*schema.Column{ConversationsColumns[0]},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "message", Type: field.TypeString},
		{Name: "read", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "message_conversation", Type: field.TypeUUID, Nullable: true},
		{Name: "message_user", Type: field.TypeUUID, Nullable: true},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_conversations_conversation",
				Columns:    []*schema.Column{MessagesColumns[4]},
				RefColumns: []*schema.Column{ConversationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "messages_users_user",
				Columns:    []*schema.Column{MessagesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "login", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// ConversationUserColumns holds the columns for the "conversation_user" table.
	ConversationUserColumns = []*schema.Column{
		{Name: "conversation_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// ConversationUserTable holds the schema information for the "conversation_user" table.
	ConversationUserTable = &schema.Table{
		Name:       "conversation_user",
		Columns:    ConversationUserColumns,
		PrimaryKey: []*schema.Column{ConversationUserColumns[0], ConversationUserColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "conversation_user_conversation_id",
				Columns:    []*schema.Column{ConversationUserColumns[0]},
				RefColumns: []*schema.Column{ConversationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "conversation_user_user_id",
				Columns:    []*schema.Column{ConversationUserColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ConversationsTable,
		MessagesTable,
		UsersTable,
		ConversationUserTable,
	}
)

func init() {
	MessagesTable.ForeignKeys[0].RefTable = ConversationsTable
	MessagesTable.ForeignKeys[1].RefTable = UsersTable
	ConversationUserTable.ForeignKeys[0].RefTable = ConversationsTable
	ConversationUserTable.ForeignKeys[1].RefTable = UsersTable
}