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

// BoolNative
func (w *WhereBuilder) BoolNative(condition bool, query string, args ...any) *WhereBuilder {
	if !condition {
		return w
	}
	return w.Native(query, args...)
}

// Native
// in 需要用 in (?)
// 其他 = >  直接用 ?
func (w *WhereBuilder) Native(query string, args ...any) *WhereBuilder {
	if w.err != nil {
		return w
	}
	query, args, err := processNativeExIn(query, args...)
	if err != nil {
		w.err = err
		return w
	}
	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  Native,
			query: query,
			args:  args,
		},
	})
	return w
}
