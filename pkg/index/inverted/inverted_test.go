// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package inverted

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/apache/skywalking-banyandb/pkg/index/posting"
	"github.com/apache/skywalking-banyandb/pkg/index/posting/roaring"
	"github.com/apache/skywalking-banyandb/pkg/index/testcases"
	"github.com/apache/skywalking-banyandb/pkg/logger"
	"github.com/apache/skywalking-banyandb/pkg/test"
)

func TestStore_MatchTerm(t *testing.T) {
	tester := assert.New(t)
	path, fn := setUp(require.New(t))
	s, err := NewStore(StoreOpts{
		Path:   path,
		Logger: logger.GetLogger("test"),
	})
	defer func() {
		tester.NoError(s.Close())
		fn()
	}()
	tester.NoError(err)
	testcases.SetUp(tester, s)
	testcases.RunServiceName(t, s)
}

func TestStore_MatchTerm_AfterFlush(t *testing.T) {
	tester := assert.New(t)
	path, fn := setUp(require.New(t))
	s, err := NewStore(StoreOpts{
		Path:   path,
		Logger: logger.GetLogger("test"),
	})
	defer func() {
		tester.NoError(s.Close())
		fn()
	}()
	tester.NoError(err)
	testcases.SetUp(tester, s)
	tester.NoError(s.(*store).Flush())
	testcases.RunServiceName(t, s)
}

func TestStore_Iterator(t *testing.T) {
	tester := assert.New(t)
	path, fn := setUp(require.New(t))
	s, err := NewStore(StoreOpts{
		Path:   path,
		Logger: logger.GetLogger("test"),
	})
	defer func() {
		tester.NoError(s.Close())
		fn()
	}()
	tester.NoError(err)
	data := testcases.SetUpDuration(tester, s)
	testcases.RunDuration(t, data, s)
}

func TestStore_Iterator_AfterFlush(t *testing.T) {
	tester := assert.New(t)
	path, fn := setUp(require.New(t))
	s, err := NewStore(StoreOpts{
		Path:   path,
		Logger: logger.GetLogger("test"),
	})
	defer func() {
		tester.NoError(s.Close())
		fn()
	}()
	tester.NoError(err)
	data := testcases.SetUpDuration(tester, s)
	tester.NoError(s.(*store).Flush())
	testcases.RunDuration(t, data, s)
}

func TestStore_Iterator_Hybrid(t *testing.T) {
	tester := assert.New(t)
	path, fn := setUp(require.New(t))
	s, err := NewStore(StoreOpts{
		Path:   path,
		Logger: logger.GetLogger("test"),
	})
	defer func() {
		tester.NoError(s.Close())
		fn()
	}()
	tester.NoError(err)
	r := map[int]posting.List{
		50:   roaring.NewPostingList(),
		200:  nil,
		500:  roaring.NewPostingList(),
		1000: nil,
		2000: roaring.NewPostingList(),
	}
	data1 := testcases.SetUpPartialDuration(tester, s, r)
	tester.NoError(s.(*store).Flush())
	r = map[int]posting.List{
		50:   nil,
		200:  roaring.NewPostingList(),
		500:  nil,
		1000: roaring.NewPostingList(),
		2000: nil,
	}
	data := testcases.SetUpPartialDuration(tester, s, r)
	for i, list := range data {
		if list == nil {
			data[i] = data1[i]
		}
	}
	testcases.RunDuration(t, data, s)
}

func setUp(t *require.Assertions) (tempDir string, deferFunc func()) {
	t.NoError(logger.Init(logger.Logging{
		Env:   "dev",
		Level: "warn",
	}))
	tempDir, deferFunc = test.Space(t)
	return tempDir, deferFunc
}
