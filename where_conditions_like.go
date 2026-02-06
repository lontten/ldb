//  Copyright 2025 lontten lontten@163.com
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package ldb

type likeType int

const (
	LikeAnywhere   likeType = iota // 两边加通配符，匹配任意位置
	LikeEndsWith                   // 左边加通配符，匹配结尾
	LikeStartsWith                 // 右边加通配符，匹配开头
)

// columns 多个字段，任意一个字段满足 LIKE ?
func (w *WhereBuilder) _like(key *string, likeType likeType, noLike bool, columns ...string) *WhereBuilder {
	if w.err != nil {
		return w
	}
	if key == nil {
		return w
	}
	if *key == "" {
		return w
	}

	var k = ""
	switch likeType {
	case LikeAnywhere:
		k = "%" + *key + "%"
	case LikeEndsWith:
		k = "%" + *key
	case LikeStartsWith:
		k = *key + "%"
	}

	var likeTokenType = Like
	if noLike {
		likeTokenType = NotLike
	}

	likeW := W()
	for _, field := range columns {
		if noLike {
			likeW.andWheres = append(likeW.andWheres, WhereBuilder{
				clause: &Clause{
					Type:  likeTokenType,
					query: field,
					args:  []any{k},
				},
			})
		} else {
			likeW.wheres = append(likeW.wheres, WhereBuilder{
				clause: &Clause{
					Type:  likeTokenType,
					query: field,
					args:  []any{k},
				},
			})
		}
	}
	w.And(likeW)
	return w
}

// Like
// LIKE %?%
func (w *WhereBuilder) Like(key *string, field ...string) *WhereBuilder {
	w._like(key, LikeAnywhere, false, field...)
	return w
}

func (w *WhereBuilder) BoolLike(condition bool, key *string, field ...string) *WhereBuilder {
	if !condition {
		return w
	}
	return w.Like(key, field...)
}

// LikeLeft 左匹配
// LIKE %?
func (w *WhereBuilder) LikeLeft(key *string, field ...string) *WhereBuilder {
	w._like(key, LikeStartsWith, false, field...)
	return w
}
func (w *WhereBuilder) BoolLikeLeft(condition bool, key *string, field ...string) *WhereBuilder {
	if !condition {
		return w
	}
	return w.LikeLeft(key, field...)
}

// LikeRight 右匹配
// LIKE ?%
func (w *WhereBuilder) LikeRight(key *string, field ...string) *WhereBuilder {
	w._like(key, LikeEndsWith, false, field...)
	return w
}
func (w *WhereBuilder) BoolLikeRight(condition bool, key *string, field ...string) *WhereBuilder {
	if !condition {
		return w
	}
	return w.LikeLeft(key, field...)
}

// NoLike
// NOT LIKE %?%
func (w *WhereBuilder) NoLike(key *string, field ...string) *WhereBuilder {
	w._like(key, LikeAnywhere, true, field...)
	return w
}

func (w *WhereBuilder) BoolNoLike(condition bool, key *string, field ...string) *WhereBuilder {
	if !condition {
		return w
	}
	return w.NoLike(key, field...)
}

// NoLikeLeft 左匹配
// NOT LIKE %?
func (w *WhereBuilder) NoLikeLeft(key *string, field ...string) *WhereBuilder {
	w._like(key, LikeStartsWith, true, field...)
	return w
}
func (w *WhereBuilder) BoolNoLikeLeft(condition bool, key *string, field ...string) *WhereBuilder {
	if !condition {
		return w
	}
	return w.NoLikeLeft(key, field...)
}

// NoLikeRight 右匹配
// NOT LIKE ?%
func (w *WhereBuilder) NoLikeRight(key *string, field ...string) *WhereBuilder {
	w._like(key, LikeEndsWith, true, field...)
	return w
}
func (w *WhereBuilder) BoolNoLikeRight(condition bool, key *string, field ...string) *WhereBuilder {
	if !condition {
		return w
	}
	return w.NoLikeLeft(key, field...)
}
