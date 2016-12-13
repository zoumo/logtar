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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadJSONConfig(t *testing.T) {

	// var log_config interface{}
	config := []byte(`{
        "disableExistingLoggers": true,
        "handlers": {
            "null": {
                "class": "NullHandler",
                "level": "DEBUG"
            },
            "console": {
                "class": "StreamHandler",
                "formatter": "default",
                "level": "INFO"
            }
        },
        "loggers": {
            "app": {
                "level": "DEBUG",
                "handlers": ["null", "console"]
            }
        }

    }`)

	err := LoadJSONConfig(config)
	assert.Nil(t, err)
	var h Handler
	h = GetHandler("null")
	assert.NotNil(t, h)
	h = GetHandler("console")
	assert.NotNil(t, h)
	var l *Logger
	l = GetLogger("app")
	assert.NotNil(t, l)

}
