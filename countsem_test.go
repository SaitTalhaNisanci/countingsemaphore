// Copyright (c) 2008-2018, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package countingsemaphore

import (
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	size := 10
	sem := New(size)
	counter := 0
	go func() {
		for i := 0 ; i < size*2; i++ {
			sem.Lock()
			counter ++
		}
	}()
	time.Sleep(time.Second)
	if counter != size {
		t.Errorf("expected %d got %d", size, counter)
	}

}

func TestUnlock(t *testing.T) {
	size := 10
	sem := New(size)
	counter := 0
	go func() {
		for i := 0 ; i < size*2; i++ {
			sem.Lock()
			counter++
		}
	}()
	time.Sleep(time.Second)
	if counter != size {
		t.Errorf("expected %d got %d", size, counter)
	}

	go func() {
		for i := 0 ; i < size; i++ {
			sem.Unlock()
		}
	}()

	time.Sleep(time.Second)
	if counter != size*2 {
		t.Errorf("expected %d got %d", size*2, counter)
	}

}
