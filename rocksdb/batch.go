//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package rocksdb

import "github.com/nick13jaremek/grocksdb"

type Batch struct {
	batch *grocksdb.WriteBatch
}

func (b *Batch) Set(key, val []byte) {
	b.batch.Put(key, val)
}

func (b *Batch) Delete(key []byte) {
	b.batch.Delete(key)
}

func (b *Batch) Merge(key, val []byte) {
	b.batch.Merge(key, val)
}

func (b *Batch) Reset() {
	b.batch.Clear()
}

func (b *Batch) Close() error {
	b.batch.Destroy()
	b.batch = nil
	return nil
}
