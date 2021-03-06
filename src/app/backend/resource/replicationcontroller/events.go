// Copyright 2017 The Kubernetes Dashboard Authors.
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

package replicationcontroller

import (
	"github.com/kubernetes/dashboard/src/app/backend/resource/common"
	"github.com/kubernetes/dashboard/src/app/backend/resource/dataselect"
	"github.com/kubernetes/dashboard/src/app/backend/resource/event"
	client "k8s.io/client-go/kubernetes"
)

// GetReplicationControllerEvents returns events for particular namespace and replication
// controller or error if occurred.
func GetReplicationControllerEvents(client client.Interface, dsQuery *dataselect.DataSelectQuery,
	namespace, replicationControllerName string) (*common.EventList, error) {

	// Get events for replication controller.
	rsEvents, err := event.GetEvents(client, namespace, replicationControllerName)
	if err != nil {
		return event.EmptyEventList, err
	}

	events := event.CreateEventList(rsEvents, dsQuery)
	return &events, nil
}
