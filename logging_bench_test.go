// Copyright 2016 Jim Zhang (jim.zoumo@gmail.com)
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
package logdog

import "testing"

func createLogger() *Logger {
	return NewLogger().AddHandler(NewStreamHandler().discardOutput())
}

func BenchmarkLogWithoutFields(b *testing.B) {
	b.ResetTimer()
	logger := createLogger()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("test")
		}
	})
}

func BenchmarkLogWithFields(b *testing.B) {
	b.ResetTimer()
	logger := createLogger()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("test", smallFields)
		}
	})
}
