/*******************************************************************************
 * Copyright 2019 Samsung Electronics All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *******************************************************************************/

package requestervalidator

import (
	"testing"
)

func TestCheckRequester(t *testing.T) {
	serviceName := "test"
	requesters := []string{"test1", "test2"}

	RequesterValidator{}.StoreRequesterInfo(serviceName, requesters)

	t.Run("Error", func(t *testing.T) {
		t.Run("NotStoredService", func(t *testing.T) {
			err := RequesterValidator{}.CheckRequester("notStored", "")
			if err == nil {
				t.Error("unexpected succeed")
			}
		})
		t.Run("NotAllowedRequester", func(t *testing.T) {
			err := RequesterValidator{}.CheckRequester("test", "notAllowed")
			if err == nil {
				t.Error("unexpected succeed")
			}
		})
	})
}

func TestStoreRequesterInfo(t *testing.T) {
	serviceName := "test"
	requesters := []string{"test1", "test2"}

	RequesterValidator{}.StoreRequesterInfo(serviceName, requesters)

	t.Run("Success", func(t *testing.T) {
		t.Run("StoreOnce", func(t *testing.T) {
			for _, req := range requesters {
				err := RequesterValidator{}.CheckRequester(serviceName, req)
				if err != nil {
					t.Error("unexpected error")
				}
			}
		})
	})
}
