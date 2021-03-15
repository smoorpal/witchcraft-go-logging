// Copyright (c) 2021 Palantir Technologies. All rights reserved.
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

package wrapped1log

import (
	"io"

	"github.com/palantir/witchcraft-go-logging/wlog"
	"github.com/palantir/witchcraft-go-logging/wlog/svclog/svc1log"
)

type Logger interface {
	Service(params ...svc1log.Param) svc1log.Logger
}

func New(w io.Writer, level wlog.LogLevel, name, version string) Logger {
	return NewFromProvider(w, level, wlog.DefaultLoggerProvider(), name, version)
}

func NewFromProvider(w io.Writer, level wlog.LogLevel, creator wlog.LoggerProvider, name, version string) Logger {
	return &defaultLogger{
		name:        name,
		version:     version,
		logger:      creator.NewLogger(w),
		levellogger: creator.NewLeveledLogger(w, level),
	}
}