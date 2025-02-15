// Copyright 2020 PingCAP, Inc.
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

package util

import (
	"fmt"

	"github.com/pingcap/tidb/pkg/expression"
	"github.com/pingcap/tidb/pkg/sessionctx"
	"github.com/pingcap/tidb/pkg/util/size"
)

// ByItems wraps a "by" item.
type ByItems struct {
	Expr expression.Expression
	Desc bool
}

// String implements fmt.Stringer interface.
func (by *ByItems) String() string {
	if by.Desc {
		return fmt.Sprintf("%s true", by.Expr)
	}
	return by.Expr.String()
}

// Clone makes a copy of ByItems.
func (by *ByItems) Clone() *ByItems {
	return &ByItems{Expr: by.Expr.Clone(), Desc: by.Desc}
}

// Equal checks whether two ByItems are equal.
func (by *ByItems) Equal(ctx sessionctx.Context, other *ByItems) bool {
	return by.Expr.Equal(ctx, other.Expr) && by.Desc == other.Desc
}

// MemoryUsage return the memory usage of ByItems.
func (by *ByItems) MemoryUsage() (sum int64) {
	if by == nil {
		return
	}

	sum = size.SizeOfBool
	if by.Expr != nil {
		sum += by.Expr.MemoryUsage()
	}
	return sum
}
