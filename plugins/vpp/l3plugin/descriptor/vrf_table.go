// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package descriptor

import (
	"fmt"
	"math"
	"strings"

	"github.com/ligato/cn-infra/idxmap"

	"github.com/pkg/errors"

	"github.com/ligato/cn-infra/logging"
	l3 "github.com/ligato/vpp-agent/api/models/vpp/l3"
	kvs "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	"github.com/ligato/vpp-agent/plugins/vpp/l3plugin/descriptor/adapter"
	"github.com/ligato/vpp-agent/plugins/vpp/l3plugin/vppcalls"
	"github.com/ligato/vpp-agent/plugins/vpp/l3plugin/vrfidx"
)

const (
	// VrfTableDescriptorName is the name of the descriptor for VRF tables.
	VrfTableDescriptorName = "vpp-vrf-table"

	// how many characters a VRF table label is allowed to have
	//  - determined by much fits into the binary API (64 null-terminated character string)
	labelLengthLimit = 63
)

// A list of non-retriable errors:
var (
	// ErrVrfTableLabelTooLong is returned when VRF table label exceeds the length limit.
	ErrVrfTableLabelTooLong = errors.New("VPP VRF table label exceeds the length limit (63 characters)")
)

// VrfTableDescriptor teaches KVScheduler how to configure VPP VRF tables.
type VrfTableDescriptor struct {
	log       logging.Logger
	vtHandler vppcalls.VrfTableVppAPI
}

// NewVrfTableDescriptor creates a new instance of the VrfTable descriptor.
func NewVrfTableDescriptor(
	vtHandler vppcalls.VrfTableVppAPI, log logging.PluginLogger) *kvs.KVDescriptor {

	ctx := &VrfTableDescriptor{
		vtHandler: vtHandler,
		log:       log.NewLogger("vrf-table-descriptor"),
	}
	typedDescr := &adapter.VrfTableDescriptor{
		Name:               VrfTableDescriptorName,
		NBKeyPrefix:        l3.ModelVrfTable.KeyPrefix(),
		ValueTypeName:      l3.ModelVrfTable.ProtoName(),
		KeySelector:        l3.ModelVrfTable.IsKeyValid,
		KeyLabel:           l3.ModelVrfTable.StripKeyPrefix,
		WithMetadata:       true,
		MetadataMapFactory: ctx.MetadataFactory,
		ValueComparator:    ctx.EquivalentVrfTables,
		Validate:           ctx.Validate,
		Create:             ctx.Create,
		Delete:             ctx.Delete,
		Retrieve:           ctx.Retrieve,
	}
	return adapter.NewVrfTableDescriptor(typedDescr)
}

// EquivalentVrfTables is a comparison function for l3.VrfTable.
func (d *VrfTableDescriptor) EquivalentVrfTables(key string, oldVrfTable, newVrfTable *l3.VrfTable) bool {
	if getVrfTableLabel(oldVrfTable) != getVrfTableLabel(newVrfTable) {
		return false
	}
	return true
}

// MetadataFactory is a factory for index-map customized for VRFs.
func (d *VrfTableDescriptor) MetadataFactory() idxmap.NamedMappingRW {
	return vrfidx.NewVRFIndex(d.log, "vpp-vrf-index")
}

// Validate validates configuration of VPP VRF table.
func (d *VrfTableDescriptor) Validate(key string, vrfTable *l3.VrfTable) (err error) {
	if len(vrfTable.Label) > labelLengthLimit {
		return kvs.NewInvalidValueError(ErrVrfTableLabelTooLong, "label")
	}
	return nil
}

// Create adds VPP VRF table.
func (d *VrfTableDescriptor) Create(key string, vrfTable *l3.VrfTable) (metadata *vrfidx.VRFMetadata, err error) {
	if vrfTable.Id == 0 {
		// nothing to do, automatically created by VPP
	}
	err = d.vtHandler.AddVrfTable(vrfTable)
	if err != nil {
		return nil, err
	}

	// fill the metadata
	metadata = &vrfidx.VRFMetadata{
		Index:    vrfTable.Id,
		Protocol: vrfTable.Protocol,
	}

	return metadata, nil
}

// Delete removes VPP VRF table.
func (d *VrfTableDescriptor) Delete(key string, vrfTable *l3.VrfTable, metadata *vrfidx.VRFMetadata) error {
	if vrfTable.Id == 0 {
		// nothing to do, VRF ID=0 always exists
	}
	err := d.vtHandler.DelVrfTable(vrfTable)
	if err != nil {
		return err
	}

	return nil
}

// Retrieve returns all configured VPP VRF tables.
func (d *VrfTableDescriptor) Retrieve(correlate []adapter.VrfTableKVWithMetadata) (
	retrieved []adapter.VrfTableKVWithMetadata, err error,
) {
	tables, err := d.vtHandler.DumpVrfTables()
	if err != nil {
		return nil, errors.Errorf("failed to dump VPP VRF tables: %v", err)
	}

	for _, table := range tables {
		origin := kvs.UnknownOrigin
		// VRF with ID=0 and ID=MaxUint32 are special
		// and should not be removed automatically
		if table.Id > 0 && table.Id < math.MaxUint32 {
			origin = kvs.FromNB
		}
		retrieved = append(retrieved, adapter.VrfTableKVWithMetadata{
			Key:   l3.VrfTableKey(table.Id, table.Protocol),
			Value: table,
			Metadata: &vrfidx.VRFMetadata{
				Index:    table.Id,
				Protocol: table.Protocol,
			},
			Origin: origin,
		})
	}

	return retrieved, nil
}

func getVrfTableLabel(vrfTable *l3.VrfTable) string {
	if vrfTable.Label == "" {
		// label generated by VPP (e.g. "ipv4-VRF:0")
		return fmt.Sprintf("%s-VRF:%d",
			strings.ToLower(vrfTable.Protocol.String()), vrfTable.Id)
	}
	return vrfTable.Label
}
