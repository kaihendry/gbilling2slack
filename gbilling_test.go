/**
 * Copyright (c) 2019-present Future Corporation
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
package gbilling2slack

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
)

func TestNotifyBilling(t *testing.T) {
	type args struct {
		ctx context.Context
		msg *pubsub.Message
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// source .env
		{"local test", args{context.Background(), &pubsub.Message{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NotifyBilling(tt.args.ctx, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("NotifyBilling() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}