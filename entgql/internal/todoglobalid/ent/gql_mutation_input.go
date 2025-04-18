// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"entgo.io/contrib/entgql/internal/todoglobalid/ent/category"
	"entgo.io/contrib/entgql/internal/todoglobalid/ent/schema/schematype"
	"entgo.io/contrib/entgql/internal/todoglobalid/ent/todo"
	"github.com/google/uuid"
)

// CreateCategoryInput represents a mutation input for creating categories.
type CreateCategoryInput struct {
	Text           string
	Status         category.Status
	Config         *schematype.CategoryConfig
	Types          *schematype.CategoryTypes
	Duration       *time.Duration
	Count          *uint64
	Strings        []string
	TodoIDs        []int
	SubCategoryIDs []int
}

// Mutate applies the CreateCategoryInput on the CategoryMutation builder.
func (i *CreateCategoryInput) Mutate(m *CategoryMutation) {
	m.SetText(i.Text)
	m.SetStatus(i.Status)
	if v := i.Config; v != nil {
		m.SetConfig(v)
	}
	if v := i.Types; v != nil {
		m.SetTypes(v)
	}
	if v := i.Duration; v != nil {
		m.SetDuration(*v)
	}
	if v := i.Count; v != nil {
		m.SetCount(*v)
	}
	if v := i.Strings; v != nil {
		m.SetStrings(v)
	}
	if v := i.TodoIDs; len(v) > 0 {
		m.AddTodoIDs(v...)
	}
	if v := i.SubCategoryIDs; len(v) > 0 {
		m.AddSubCategoryIDs(v...)
	}
}

// SetInput applies the change-set in the CreateCategoryInput on the CategoryCreate builder.
func (c *CategoryCreate) SetInput(i CreateCategoryInput) *CategoryCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateCategoryInput represents a mutation input for updating categories.
type UpdateCategoryInput struct {
	Text                 *string
	Status               *category.Status
	ClearConfig          bool
	Config               *schematype.CategoryConfig
	ClearTypes           bool
	Types                *schematype.CategoryTypes
	ClearDuration        bool
	Duration             *time.Duration
	ClearCount           bool
	Count                *uint64
	ClearStrings         bool
	Strings              []string
	AppendStrings        []string
	ClearTodos           bool
	AddTodoIDs           []int
	RemoveTodoIDs        []int
	ClearSubCategories   bool
	AddSubCategoryIDs    []int
	RemoveSubCategoryIDs []int
}

// Mutate applies the UpdateCategoryInput on the CategoryMutation builder.
func (i *UpdateCategoryInput) Mutate(m *CategoryMutation) {
	if v := i.Text; v != nil {
		m.SetText(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if i.ClearConfig {
		m.ClearConfig()
	}
	if v := i.Config; v != nil {
		m.SetConfig(v)
	}
	if i.ClearTypes {
		m.ClearTypes()
	}
	if v := i.Types; v != nil {
		m.SetTypes(v)
	}
	if i.ClearDuration {
		m.ClearDuration()
	}
	if v := i.Duration; v != nil {
		m.SetDuration(*v)
	}
	if i.ClearCount {
		m.ClearCount()
	}
	if v := i.Count; v != nil {
		m.SetCount(*v)
	}
	if i.ClearStrings {
		m.ClearStrings()
	}
	if v := i.Strings; v != nil {
		m.SetStrings(v)
	}
	if i.AppendStrings != nil {
		m.AppendStrings(i.Strings)
	}
	if i.ClearTodos {
		m.ClearTodos()
	}
	if v := i.AddTodoIDs; len(v) > 0 {
		m.AddTodoIDs(v...)
	}
	if v := i.RemoveTodoIDs; len(v) > 0 {
		m.RemoveTodoIDs(v...)
	}
	if i.ClearSubCategories {
		m.ClearSubCategories()
	}
	if v := i.AddSubCategoryIDs; len(v) > 0 {
		m.AddSubCategoryIDs(v...)
	}
	if v := i.RemoveSubCategoryIDs; len(v) > 0 {
		m.RemoveSubCategoryIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateCategoryInput on the CategoryUpdate builder.
func (c *CategoryUpdate) SetInput(i UpdateCategoryInput) *CategoryUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateCategoryInput on the CategoryUpdateOne builder.
func (c *CategoryUpdateOne) SetInput(i UpdateCategoryInput) *CategoryUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// UpdateFriendshipInput represents a mutation input for updating friendships.
type UpdateFriendshipInput struct {
	CreatedAt *time.Time
	UserID    *int
	FriendID  *int
}

// Mutate applies the UpdateFriendshipInput on the FriendshipMutation builder.
func (i *UpdateFriendshipInput) Mutate(m *FriendshipMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
	if v := i.FriendID; v != nil {
		m.SetFriendID(*v)
	}
}

// SetInput applies the change-set in the UpdateFriendshipInput on the FriendshipUpdate builder.
func (c *FriendshipUpdate) SetInput(i UpdateFriendshipInput) *FriendshipUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateFriendshipInput on the FriendshipUpdateOne builder.
func (c *FriendshipUpdateOne) SetInput(i UpdateFriendshipInput) *FriendshipUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateTodoInput represents a mutation input for creating todos.
type CreateTodoInput struct {
	Status     todo.Status
	Priority   *int
	Text       string
	Init       map[string]interface{}
	Value      *int
	ParentID   *int
	ChildIDs   []int
	CategoryID *int
	SecretID   *int
}

// Mutate applies the CreateTodoInput on the TodoMutation builder.
func (i *CreateTodoInput) Mutate(m *TodoMutation) {
	m.SetStatus(i.Status)
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	m.SetText(i.Text)
	if v := i.Init; v != nil {
		m.SetInit(v)
	}
	if v := i.Value; v != nil {
		m.SetValue(*v)
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if v := i.ChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
	if v := i.CategoryID; v != nil {
		m.SetCategoryID(*v)
	}
	if v := i.SecretID; v != nil {
		m.SetSecretID(*v)
	}
}

// SetInput applies the change-set in the CreateTodoInput on the TodoCreate builder.
func (c *TodoCreate) SetInput(i CreateTodoInput) *TodoCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTodoInput represents a mutation input for updating todos.
type UpdateTodoInput struct {
	Status         *todo.Status
	Priority       *int
	Text           *string
	ClearInit      bool
	Init           map[string]interface{}
	Value          *int
	ClearParent    bool
	ParentID       *int
	ClearChildren  bool
	AddChildIDs    []int
	RemoveChildIDs []int
	ClearSecret    bool
	SecretID       *int
}

// Mutate applies the UpdateTodoInput on the TodoMutation builder.
func (i *UpdateTodoInput) Mutate(m *TodoMutation) {
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	if v := i.Text; v != nil {
		m.SetText(*v)
	}
	if i.ClearInit {
		m.ClearInit()
	}
	if v := i.Init; v != nil {
		m.SetInit(v)
	}
	if v := i.Value; v != nil {
		m.SetValue(*v)
	}
	if i.ClearParent {
		m.ClearParent()
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if i.ClearChildren {
		m.ClearChildren()
	}
	if v := i.AddChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
	if v := i.RemoveChildIDs; len(v) > 0 {
		m.RemoveChildIDs(v...)
	}
	if i.ClearSecret {
		m.ClearSecret()
	}
	if v := i.SecretID; v != nil {
		m.SetSecretID(*v)
	}
}

// SetInput applies the change-set in the UpdateTodoInput on the TodoUpdate builder.
func (c *TodoUpdate) SetInput(i UpdateTodoInput) *TodoUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTodoInput on the TodoUpdateOne builder.
func (c *TodoUpdateOne) SetInput(i UpdateTodoInput) *TodoUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Name             *string
	Username         *uuid.UUID
	Password         *string
	RequiredMetadata map[string]interface{}
	Metadata         map[string]interface{}
	GroupIDs         []int
	FriendIDs        []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Username; v != nil {
		m.SetUsername(*v)
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if v := i.RequiredMetadata; v != nil {
		m.SetRequiredMetadata(v)
	}
	if v := i.Metadata; v != nil {
		m.SetMetadata(v)
	}
	if v := i.GroupIDs; len(v) > 0 {
		m.AddGroupIDs(v...)
	}
	if v := i.FriendIDs; len(v) > 0 {
		m.AddFriendIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	Name             *string
	Username         *uuid.UUID
	ClearPassword    bool
	Password         *string
	RequiredMetadata map[string]interface{}
	ClearMetadata    bool
	Metadata         map[string]interface{}
	ClearGroups      bool
	AddGroupIDs      []int
	RemoveGroupIDs   []int
	ClearFriends     bool
	AddFriendIDs     []int
	RemoveFriendIDs  []int
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Username; v != nil {
		m.SetUsername(*v)
	}
	if i.ClearPassword {
		m.ClearPassword()
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if v := i.RequiredMetadata; v != nil {
		m.SetRequiredMetadata(v)
	}
	if i.ClearMetadata {
		m.ClearMetadata()
	}
	if v := i.Metadata; v != nil {
		m.SetMetadata(v)
	}
	if i.ClearGroups {
		m.ClearGroups()
	}
	if v := i.AddGroupIDs; len(v) > 0 {
		m.AddGroupIDs(v...)
	}
	if v := i.RemoveGroupIDs; len(v) > 0 {
		m.RemoveGroupIDs(v...)
	}
	if i.ClearFriends {
		m.ClearFriends()
	}
	if v := i.AddFriendIDs; len(v) > 0 {
		m.AddFriendIDs(v...)
	}
	if v := i.RemoveFriendIDs; len(v) > 0 {
		m.RemoveFriendIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
