package main

import (
	"errors"
)

var ErrNoAvatarURL = errors.New("chat: アバターのURLが取得できません。")

// Avatarはユーザのプロフィール画像を表す型です。
type Avatar interface {
	// GetAvatarURLは指定されたクライアントのアバターのURLを返します。
	// 問題が発生した場合にはエラーを返します。特に、URLが取得できなかった
	// 場合にはErrNoAvatarURLを返します。
	GetAvatarURL(c *client) (string, error)
}
