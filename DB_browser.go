// DO NOT EDIT THIS FILE. IT IS AUTO-GENERATED BY "gen/gen-db.go". //
/*
 * Copyright (c) 2018 wellwell.work, LLC by Zoe
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *     
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package surferua

func init () {
	browserDBSize = 3
	browserDB = []BrowserInfo{
		{
			Name: "Firefox",
			EngineInfo: EngineInfo{
				Name: "Gecko",
				VersionInfo: VersionInfo{
					Major: Endpoint{Start: 0, End: 0},
					Minor: Endpoint{Start: 0, End: 0},
					Patch: Endpoint{Start: 0, End: 0},
				},
			},
			VersionInfo: VersionInfo{
				Major: Endpoint{Start: 35, End: 56},
				Minor: Endpoint{Start: 35, End: 56},
				Patch: Endpoint{Start: 0, End: 3},
			},
		},
		{
			Name: "Chrome",
			EngineInfo: EngineInfo{
				Name: "AppleWebKit",
				VersionInfo: VersionInfo{
					Major: Endpoint{Start: 534, End: 603},
					Minor: Endpoint{Start: 35, End: 56},
					Patch: Endpoint{Start: 0, End: 0},
				},
			},
			VersionInfo: VersionInfo{
				Major: Endpoint{Start: 39, End: 64},
				Minor: Endpoint{Start: 0, End: 0},
				Patch: Endpoint{Start: 0, End: 3000},
			},
		},
		{
			Name: "Safari",
			EngineInfo: EngineInfo{
				Name: "WebKit",
				VersionInfo: VersionInfo{
					Major: Endpoint{Start: 534, End: 603},
					Minor: Endpoint{Start: 0, End: 21},
					Patch: Endpoint{Start: 0, End: 10},
				},
			},
			VersionInfo: VersionInfo{
				Major: Endpoint{Start: 5, End: 11},
				Minor: Endpoint{Start: 0, End: 2},
				Patch: Endpoint{Start: 0, End: 10},
			},
		},
	}
}


// Firefox returns Specified browser: Firefox
func (ua *UserAgent) Firefox () *UserAgent {
	ua.browser = browserDB[0].Random()
	return ua
}

// Chrome returns Specified browser: Chrome
func (ua *UserAgent) Chrome () *UserAgent {
	ua.browser = browserDB[1].Random()
	return ua
}

// Safari returns Specified browser: Safari
func (ua *UserAgent) Safari () *UserAgent {
	ua.browser = browserDB[2].Random()
	return ua
}

