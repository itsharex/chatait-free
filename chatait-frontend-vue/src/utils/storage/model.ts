/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface Token {
  accessToken: string
  accessTokenExpire: number
  accessTokenExpireIn: number
  refreshToken: string
  refreshTokenExpire: number
  refreshTokenExpireIn: number
}

export interface UserInfo {
  id: string
  username: string
  nickname: string
  avatar: string
}
