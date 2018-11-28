// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/gogo/protobuf/proto"
	. "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	"github.com/ligato/vpp-agent/plugins/vppv2/model/ipsec"
)

////////// type-safe key-value pair with metadata //////////

type SAKVWithMetadata struct {
	Key      string
	Value    *ipsec.SecurityAssociation
	Metadata interface{}
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type SADescriptor struct {
	Name               string
	KeySelector        KeySelector
	ValueTypeName      string
	KeyLabel           func(key string) string
	ValueComparator    func(key string, oldValue, newValue *ipsec.SecurityAssociation) bool
	NBKeyPrefix        string
	WithMetadata       bool
	MetadataMapFactory MetadataMapFactory
	Add                func(key string, value *ipsec.SecurityAssociation) (metadata interface{}, err error)
	Delete             func(key string, value *ipsec.SecurityAssociation, metadata interface{}) error
	Modify             func(key string, oldValue, newValue *ipsec.SecurityAssociation, oldMetadata interface{}) (newMetadata interface{}, err error)
	ModifyWithRecreate func(key string, oldValue, newValue *ipsec.SecurityAssociation, metadata interface{}) bool
	Update             func(key string, value *ipsec.SecurityAssociation, metadata interface{}) error
	IsRetriableFailure func(err error) bool
	Dependencies       func(key string, value *ipsec.SecurityAssociation) []Dependency
	DerivedValues      func(key string, value *ipsec.SecurityAssociation) []KeyValuePair
	Dump               func(correlate []SAKVWithMetadata) ([]SAKVWithMetadata, error)
	DumpDependencies   []string /* descriptor name */
}

////////// Descriptor adapter //////////

type SADescriptorAdapter struct {
	descriptor *SADescriptor
}

func NewSADescriptor(typedDescriptor *SADescriptor) *KVDescriptor {
	adapter := &SADescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:               typedDescriptor.Name,
		KeySelector:        typedDescriptor.KeySelector,
		ValueTypeName:      typedDescriptor.ValueTypeName,
		KeyLabel:           typedDescriptor.KeyLabel,
		NBKeyPrefix:        typedDescriptor.NBKeyPrefix,
		WithMetadata:       typedDescriptor.WithMetadata,
		MetadataMapFactory: typedDescriptor.MetadataMapFactory,
		IsRetriableFailure: typedDescriptor.IsRetriableFailure,
		DumpDependencies:   typedDescriptor.DumpDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Add != nil {
		descriptor.Add = adapter.Add
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Modify != nil {
		descriptor.Modify = adapter.Modify
	}
	if typedDescriptor.ModifyWithRecreate != nil {
		descriptor.ModifyWithRecreate = adapter.ModifyWithRecreate
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	if typedDescriptor.Dump != nil {
		descriptor.Dump = adapter.Dump
	}
	return descriptor
}

func (da *SADescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castSAValue(key, oldValue)
	typedNewValue, err2 := castSAValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *SADescriptorAdapter) Add(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castSAValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Add(key, typedValue)
}

func (da *SADescriptorAdapter) Modify(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castSAValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castSAValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castSAMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Modify(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *SADescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castSAValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castSAMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *SADescriptorAdapter) ModifyWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castSAValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castSAValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castSAMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.ModifyWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *SADescriptorAdapter) Update(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castSAValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castSAMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Update(key, typedValue, typedMetadata)
}

func (da *SADescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castSAValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

func (da *SADescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castSAValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *SADescriptorAdapter) Dump(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []SAKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castSAValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castSAMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			SAKVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedDump, err := da.descriptor.Dump(correlateWithType)
	if err != nil {
		return nil, err
	}
	var dump []KVWithMetadata
	for _, typedKVWithMetadata := range typedDump {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		dump = append(dump, kvWithMetadata)
	}
	return dump, err
}

////////// Helper methods //////////

func castSAValue(key string, value proto.Message) (*ipsec.SecurityAssociation, error) {
	typedValue, ok := value.(*ipsec.SecurityAssociation)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castSAMetadata(key string, metadata Metadata) (interface{}, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(interface{})
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}