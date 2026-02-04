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

import (
	"fmt"

	"github.com/lontten/ldb/v2/utils"
)

// In
// IN (?)
// arg 为数组；当arg不是数组时，自动封装成只有一个元素的数组。
func (w *WhereBuilder) In(field string, arg any, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}

	isNil := utils.IsNil(arg)
	if isNil {
		w.err = fmt.Errorf("invalid use of In: argument for field '%s' is nil", field)
		return w
	}

	_, args := processArrArg(arg)

	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  In,
			query: field,
			args:  args,
		},
	})
	return w
}

// NotIn
// NOT IN (?)
// arg 为数组；当arg不是数组时，自动封装成只有一个元素的数组。
func (w *WhereBuilder) NotIn(field string, arg any, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}

	isNil := utils.IsNil(arg)
	if isNil {
		w.err = fmt.Errorf("invalid use of NotIn: argument for field '%s' is nil", field)
		return w
	}

	_, args := processArrArg(arg)

	argsLen := len(args)
	if argsLen == 0 {
		return w
	}

	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  NotIn,
			query: field,
			args:  args,
		},
	})
	return w
}
