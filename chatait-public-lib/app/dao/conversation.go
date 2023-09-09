// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao/internal"
)

// conversationDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type conversationDao struct {
	*internal.ConversationDao
}

var (
	// Conversation is globally public accessible object for table c_conversation operations.
	Conversation conversationDao
)

func init() {
	Conversation = conversationDao{
		internal.NewConversationDao(),
	}
}

// Fill with you ideas below.
