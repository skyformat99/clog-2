// Copyright 2017 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package clog

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_slack_Init(t *testing.T) {
	Convey("Init Slack logger", t, func() {
		Convey("Mismatched config object", func() {
			err := New(SLACK, struct{}{})
			So(err, ShouldNotBeNil)
			_, ok := err.(ErrConfigObject)
			So(ok, ShouldBeTrue)
		})

		Convey("Valid config object", func() {
			So(New(SLACK, SlackConfig{
				URL: "https://slack.com",
			}), ShouldBeNil)

			Convey("Incorrect level", func() {
				err := New(SLACK, SlackConfig{
					Level: LEVEL(-1),
				})
				So(err, ShouldNotBeNil)
				_, ok := err.(ErrInvalidLevel)
				So(ok, ShouldBeTrue)
			})
		})
	})
}

func Test_buildSlackAttchment(t *testing.T) {
	Convey("Build Slack attachment", t, func() {
		payload, err := buildSlackPayload(&Message{
			Level: INFO,
			Body:  "test message",
		})
		So(err, ShouldBeNil)
		So(payload, ShouldEqual, `{"attachments":[{"text":"test message","color":"#3aa3e3"}]}`)
	})
}
