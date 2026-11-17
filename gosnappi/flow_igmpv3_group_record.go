package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIgmpv3GroupRecord *****
type flowIgmpv3GroupRecord struct {
	validation
	obj                    *otg.FlowIgmpv3GroupRecord
	marshaller             marshalFlowIgmpv3GroupRecord
	unMarshaller           unMarshalFlowIgmpv3GroupRecord
	recordTypeHolder       PatternFlowIgmpv3GroupRecordRecordType
	auxDataLengthHolder    PatternFlowIgmpv3GroupRecordAuxDataLength
	numberOfSourcesHolder  PatternFlowIgmpv3GroupRecordNumberOfSources
	multicastAddressHolder PatternFlowIgmpv3GroupRecordMulticastAddress
	sourcesHolder          FlowIgmpv3GroupRecordFlowIgmpv3SourceIter
}

func NewFlowIgmpv3GroupRecord() FlowIgmpv3GroupRecord {
	obj := flowIgmpv3GroupRecord{obj: &otg.FlowIgmpv3GroupRecord{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIgmpv3GroupRecord) msg() *otg.FlowIgmpv3GroupRecord {
	return obj.obj
}

func (obj *flowIgmpv3GroupRecord) setMsg(msg *otg.FlowIgmpv3GroupRecord) FlowIgmpv3GroupRecord {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIgmpv3GroupRecord struct {
	obj *flowIgmpv3GroupRecord
}

type marshalFlowIgmpv3GroupRecord interface {
	// ToProto marshals FlowIgmpv3GroupRecord to protobuf object *otg.FlowIgmpv3GroupRecord
	ToProto() (*otg.FlowIgmpv3GroupRecord, error)
	// ToPbText marshals FlowIgmpv3GroupRecord to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIgmpv3GroupRecord to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIgmpv3GroupRecord to JSON text
	ToJson() (string, error)
}

type unMarshalflowIgmpv3GroupRecord struct {
	obj *flowIgmpv3GroupRecord
}

type unMarshalFlowIgmpv3GroupRecord interface {
	// FromProto unmarshals FlowIgmpv3GroupRecord from protobuf object *otg.FlowIgmpv3GroupRecord
	FromProto(msg *otg.FlowIgmpv3GroupRecord) (FlowIgmpv3GroupRecord, error)
	// FromPbText unmarshals FlowIgmpv3GroupRecord from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIgmpv3GroupRecord from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIgmpv3GroupRecord from JSON text
	FromJson(value string) error
}

func (obj *flowIgmpv3GroupRecord) Marshal() marshalFlowIgmpv3GroupRecord {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIgmpv3GroupRecord{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIgmpv3GroupRecord) Unmarshal() unMarshalFlowIgmpv3GroupRecord {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIgmpv3GroupRecord{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIgmpv3GroupRecord) ToProto() (*otg.FlowIgmpv3GroupRecord, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIgmpv3GroupRecord) FromProto(msg *otg.FlowIgmpv3GroupRecord) (FlowIgmpv3GroupRecord, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIgmpv3GroupRecord) ToPbText() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (m *unMarshalflowIgmpv3GroupRecord) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalflowIgmpv3GroupRecord) ToYaml() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshalflowIgmpv3GroupRecord) FromYaml(value string) error {
	if value == "" {
		value = "{}"
	}
	data, err := yaml.YAMLToJSON([]byte(value))
	if err != nil {
		return err
	}
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	uError := opts.Unmarshal([]byte(data), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalflowIgmpv3GroupRecord) ToJson() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshalflowIgmpv3GroupRecord) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *flowIgmpv3GroupRecord) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIgmpv3GroupRecord) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIgmpv3GroupRecord) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIgmpv3GroupRecord) Clone() (FlowIgmpv3GroupRecord, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIgmpv3GroupRecord()
	data, err := proto.Marshal(obj.msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *flowIgmpv3GroupRecord) setNil() {
	obj.recordTypeHolder = nil
	obj.auxDataLengthHolder = nil
	obj.numberOfSourcesHolder = nil
	obj.multicastAddressHolder = nil
	obj.sourcesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIgmpv3GroupRecord is a group record of the Version 3 Membership Report (RFC 9776 Section 4.2.5).
type FlowIgmpv3GroupRecord interface {
	Validation
	// msg marshals FlowIgmpv3GroupRecord to protobuf object *otg.FlowIgmpv3GroupRecord
	// and doesn't set defaults
	msg() *otg.FlowIgmpv3GroupRecord
	// setMsg unmarshals FlowIgmpv3GroupRecord from protobuf object *otg.FlowIgmpv3GroupRecord
	// and doesn't set defaults
	setMsg(*otg.FlowIgmpv3GroupRecord) FlowIgmpv3GroupRecord
	// provides marshal interface
	Marshal() marshalFlowIgmpv3GroupRecord
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIgmpv3GroupRecord
	// validate validates FlowIgmpv3GroupRecord
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIgmpv3GroupRecord, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RecordType returns PatternFlowIgmpv3GroupRecordRecordType, set in FlowIgmpv3GroupRecord.
	// PatternFlowIgmpv3GroupRecordRecordType is record type, which tells whether the record reports the current state, a filter mode change or a source list change of the group.
	RecordType() PatternFlowIgmpv3GroupRecordRecordType
	// SetRecordType assigns PatternFlowIgmpv3GroupRecordRecordType provided by user to FlowIgmpv3GroupRecord.
	// PatternFlowIgmpv3GroupRecordRecordType is record type, which tells whether the record reports the current state, a filter mode change or a source list change of the group.
	SetRecordType(value PatternFlowIgmpv3GroupRecordRecordType) FlowIgmpv3GroupRecord
	// HasRecordType checks if RecordType has been set in FlowIgmpv3GroupRecord
	HasRecordType() bool
	// AuxDataLength returns PatternFlowIgmpv3GroupRecordAuxDataLength, set in FlowIgmpv3GroupRecord.
	// PatternFlowIgmpv3GroupRecordAuxDataLength is length of aux_data in 32-bit words (RFC 9776 Section 4.2.7). When auto, it is set to the length of aux_data in 32-bit words, so zero when aux_data is omitted. A conforming sender uses zero.
	AuxDataLength() PatternFlowIgmpv3GroupRecordAuxDataLength
	// SetAuxDataLength assigns PatternFlowIgmpv3GroupRecordAuxDataLength provided by user to FlowIgmpv3GroupRecord.
	// PatternFlowIgmpv3GroupRecordAuxDataLength is length of aux_data in 32-bit words (RFC 9776 Section 4.2.7). When auto, it is set to the length of aux_data in 32-bit words, so zero when aux_data is omitted. A conforming sender uses zero.
	SetAuxDataLength(value PatternFlowIgmpv3GroupRecordAuxDataLength) FlowIgmpv3GroupRecord
	// HasAuxDataLength checks if AuxDataLength has been set in FlowIgmpv3GroupRecord
	HasAuxDataLength() bool
	// NumberOfSources returns PatternFlowIgmpv3GroupRecordNumberOfSources, set in FlowIgmpv3GroupRecord.
	// PatternFlowIgmpv3GroupRecordNumberOfSources is number of source addresses in this record. When auto, it is set to the number of entries in sources.
	NumberOfSources() PatternFlowIgmpv3GroupRecordNumberOfSources
	// SetNumberOfSources assigns PatternFlowIgmpv3GroupRecordNumberOfSources provided by user to FlowIgmpv3GroupRecord.
	// PatternFlowIgmpv3GroupRecordNumberOfSources is number of source addresses in this record. When auto, it is set to the number of entries in sources.
	SetNumberOfSources(value PatternFlowIgmpv3GroupRecordNumberOfSources) FlowIgmpv3GroupRecord
	// HasNumberOfSources checks if NumberOfSources has been set in FlowIgmpv3GroupRecord
	HasNumberOfSources() bool
	// MulticastAddress returns PatternFlowIgmpv3GroupRecordMulticastAddress, set in FlowIgmpv3GroupRecord.
	// PatternFlowIgmpv3GroupRecordMulticastAddress is multicast group address that this record reports on.
	MulticastAddress() PatternFlowIgmpv3GroupRecordMulticastAddress
	// SetMulticastAddress assigns PatternFlowIgmpv3GroupRecordMulticastAddress provided by user to FlowIgmpv3GroupRecord.
	// PatternFlowIgmpv3GroupRecordMulticastAddress is multicast group address that this record reports on.
	SetMulticastAddress(value PatternFlowIgmpv3GroupRecordMulticastAddress) FlowIgmpv3GroupRecord
	// HasMulticastAddress checks if MulticastAddress has been set in FlowIgmpv3GroupRecord
	HasMulticastAddress() bool
	// Sources returns FlowIgmpv3GroupRecordFlowIgmpv3SourceIterIter, set in FlowIgmpv3GroupRecord
	Sources() FlowIgmpv3GroupRecordFlowIgmpv3SourceIter
	// AuxData returns string, set in FlowIgmpv3GroupRecord.
	AuxData() string
	// SetAuxData assigns string provided by user to FlowIgmpv3GroupRecord
	SetAuxData(value string) FlowIgmpv3GroupRecord
	// HasAuxData checks if AuxData has been set in FlowIgmpv3GroupRecord
	HasAuxData() bool
	setNil()
}

// description is TBD
// RecordType returns a PatternFlowIgmpv3GroupRecordRecordType
func (obj *flowIgmpv3GroupRecord) RecordType() PatternFlowIgmpv3GroupRecordRecordType {
	if obj.obj.RecordType == nil {
		obj.obj.RecordType = NewPatternFlowIgmpv3GroupRecordRecordType().msg()
	}
	if obj.recordTypeHolder == nil {
		obj.recordTypeHolder = &patternFlowIgmpv3GroupRecordRecordType{obj: obj.obj.RecordType}
	}
	return obj.recordTypeHolder
}

// description is TBD
// RecordType returns a PatternFlowIgmpv3GroupRecordRecordType
func (obj *flowIgmpv3GroupRecord) HasRecordType() bool {
	return obj.obj.RecordType != nil
}

// description is TBD
// SetRecordType sets the PatternFlowIgmpv3GroupRecordRecordType value in the FlowIgmpv3GroupRecord object
func (obj *flowIgmpv3GroupRecord) SetRecordType(value PatternFlowIgmpv3GroupRecordRecordType) FlowIgmpv3GroupRecord {

	obj.recordTypeHolder = nil
	obj.obj.RecordType = value.msg()

	return obj
}

// description is TBD
// AuxDataLength returns a PatternFlowIgmpv3GroupRecordAuxDataLength
func (obj *flowIgmpv3GroupRecord) AuxDataLength() PatternFlowIgmpv3GroupRecordAuxDataLength {
	if obj.obj.AuxDataLength == nil {
		obj.obj.AuxDataLength = NewPatternFlowIgmpv3GroupRecordAuxDataLength().msg()
	}
	if obj.auxDataLengthHolder == nil {
		obj.auxDataLengthHolder = &patternFlowIgmpv3GroupRecordAuxDataLength{obj: obj.obj.AuxDataLength}
	}
	return obj.auxDataLengthHolder
}

// description is TBD
// AuxDataLength returns a PatternFlowIgmpv3GroupRecordAuxDataLength
func (obj *flowIgmpv3GroupRecord) HasAuxDataLength() bool {
	return obj.obj.AuxDataLength != nil
}

// description is TBD
// SetAuxDataLength sets the PatternFlowIgmpv3GroupRecordAuxDataLength value in the FlowIgmpv3GroupRecord object
func (obj *flowIgmpv3GroupRecord) SetAuxDataLength(value PatternFlowIgmpv3GroupRecordAuxDataLength) FlowIgmpv3GroupRecord {

	obj.auxDataLengthHolder = nil
	obj.obj.AuxDataLength = value.msg()

	return obj
}

// description is TBD
// NumberOfSources returns a PatternFlowIgmpv3GroupRecordNumberOfSources
func (obj *flowIgmpv3GroupRecord) NumberOfSources() PatternFlowIgmpv3GroupRecordNumberOfSources {
	if obj.obj.NumberOfSources == nil {
		obj.obj.NumberOfSources = NewPatternFlowIgmpv3GroupRecordNumberOfSources().msg()
	}
	if obj.numberOfSourcesHolder == nil {
		obj.numberOfSourcesHolder = &patternFlowIgmpv3GroupRecordNumberOfSources{obj: obj.obj.NumberOfSources}
	}
	return obj.numberOfSourcesHolder
}

// description is TBD
// NumberOfSources returns a PatternFlowIgmpv3GroupRecordNumberOfSources
func (obj *flowIgmpv3GroupRecord) HasNumberOfSources() bool {
	return obj.obj.NumberOfSources != nil
}

// description is TBD
// SetNumberOfSources sets the PatternFlowIgmpv3GroupRecordNumberOfSources value in the FlowIgmpv3GroupRecord object
func (obj *flowIgmpv3GroupRecord) SetNumberOfSources(value PatternFlowIgmpv3GroupRecordNumberOfSources) FlowIgmpv3GroupRecord {

	obj.numberOfSourcesHolder = nil
	obj.obj.NumberOfSources = value.msg()

	return obj
}

// description is TBD
// MulticastAddress returns a PatternFlowIgmpv3GroupRecordMulticastAddress
func (obj *flowIgmpv3GroupRecord) MulticastAddress() PatternFlowIgmpv3GroupRecordMulticastAddress {
	if obj.obj.MulticastAddress == nil {
		obj.obj.MulticastAddress = NewPatternFlowIgmpv3GroupRecordMulticastAddress().msg()
	}
	if obj.multicastAddressHolder == nil {
		obj.multicastAddressHolder = &patternFlowIgmpv3GroupRecordMulticastAddress{obj: obj.obj.MulticastAddress}
	}
	return obj.multicastAddressHolder
}

// description is TBD
// MulticastAddress returns a PatternFlowIgmpv3GroupRecordMulticastAddress
func (obj *flowIgmpv3GroupRecord) HasMulticastAddress() bool {
	return obj.obj.MulticastAddress != nil
}

// description is TBD
// SetMulticastAddress sets the PatternFlowIgmpv3GroupRecordMulticastAddress value in the FlowIgmpv3GroupRecord object
func (obj *flowIgmpv3GroupRecord) SetMulticastAddress(value PatternFlowIgmpv3GroupRecordMulticastAddress) FlowIgmpv3GroupRecord {

	obj.multicastAddressHolder = nil
	obj.obj.MulticastAddress = value.msg()

	return obj
}

// List of unicast source addresses of this record.
// Sources returns a []FlowIgmpv3Source
func (obj *flowIgmpv3GroupRecord) Sources() FlowIgmpv3GroupRecordFlowIgmpv3SourceIter {
	if len(obj.obj.Sources) == 0 {
		obj.obj.Sources = []*otg.FlowIgmpv3Source{}
	}
	if obj.sourcesHolder == nil {
		obj.sourcesHolder = newFlowIgmpv3GroupRecordFlowIgmpv3SourceIter(&obj.obj.Sources).setMsg(obj)
	}
	return obj.sourcesHolder
}

type flowIgmpv3GroupRecordFlowIgmpv3SourceIter struct {
	obj                   *flowIgmpv3GroupRecord
	flowIgmpv3SourceSlice []FlowIgmpv3Source
	fieldPtr              *[]*otg.FlowIgmpv3Source
}

func newFlowIgmpv3GroupRecordFlowIgmpv3SourceIter(ptr *[]*otg.FlowIgmpv3Source) FlowIgmpv3GroupRecordFlowIgmpv3SourceIter {
	return &flowIgmpv3GroupRecordFlowIgmpv3SourceIter{fieldPtr: ptr}
}

type FlowIgmpv3GroupRecordFlowIgmpv3SourceIter interface {
	setMsg(*flowIgmpv3GroupRecord) FlowIgmpv3GroupRecordFlowIgmpv3SourceIter
	Items() []FlowIgmpv3Source
	Add() FlowIgmpv3Source
	Append(items ...FlowIgmpv3Source) FlowIgmpv3GroupRecordFlowIgmpv3SourceIter
	Set(index int, newObj FlowIgmpv3Source) FlowIgmpv3GroupRecordFlowIgmpv3SourceIter
	Clear() FlowIgmpv3GroupRecordFlowIgmpv3SourceIter
	clearHolderSlice() FlowIgmpv3GroupRecordFlowIgmpv3SourceIter
	appendHolderSlice(item FlowIgmpv3Source) FlowIgmpv3GroupRecordFlowIgmpv3SourceIter
}

func (obj *flowIgmpv3GroupRecordFlowIgmpv3SourceIter) setMsg(msg *flowIgmpv3GroupRecord) FlowIgmpv3GroupRecordFlowIgmpv3SourceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowIgmpv3Source{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowIgmpv3GroupRecordFlowIgmpv3SourceIter) Items() []FlowIgmpv3Source {
	return obj.flowIgmpv3SourceSlice
}

func (obj *flowIgmpv3GroupRecordFlowIgmpv3SourceIter) Add() FlowIgmpv3Source {
	newObj := &otg.FlowIgmpv3Source{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowIgmpv3Source{obj: newObj}
	newLibObj.setDefault()
	obj.flowIgmpv3SourceSlice = append(obj.flowIgmpv3SourceSlice, newLibObj)
	return newLibObj
}

func (obj *flowIgmpv3GroupRecordFlowIgmpv3SourceIter) Append(items ...FlowIgmpv3Source) FlowIgmpv3GroupRecordFlowIgmpv3SourceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowIgmpv3SourceSlice = append(obj.flowIgmpv3SourceSlice, item)
	}
	return obj
}

func (obj *flowIgmpv3GroupRecordFlowIgmpv3SourceIter) Set(index int, newObj FlowIgmpv3Source) FlowIgmpv3GroupRecordFlowIgmpv3SourceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowIgmpv3SourceSlice[index] = newObj
	return obj
}
func (obj *flowIgmpv3GroupRecordFlowIgmpv3SourceIter) Clear() FlowIgmpv3GroupRecordFlowIgmpv3SourceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowIgmpv3Source{}
		obj.flowIgmpv3SourceSlice = []FlowIgmpv3Source{}
	}
	return obj
}
func (obj *flowIgmpv3GroupRecordFlowIgmpv3SourceIter) clearHolderSlice() FlowIgmpv3GroupRecordFlowIgmpv3SourceIter {
	if len(obj.flowIgmpv3SourceSlice) > 0 {
		obj.flowIgmpv3SourceSlice = []FlowIgmpv3Source{}
	}
	return obj
}
func (obj *flowIgmpv3GroupRecordFlowIgmpv3SourceIter) appendHolderSlice(item FlowIgmpv3Source) FlowIgmpv3GroupRecordFlowIgmpv3SourceIter {
	obj.flowIgmpv3SourceSlice = append(obj.flowIgmpv3SourceSlice, item)
	return obj
}

// Auxiliary data in hex, a multiple of 4 octets (maximum 255 32-bit words). IGMPv3 does not define any auxiliary data, so a conforming sender omits it (RFC 9776 Section 4.2.11).
// AuxData returns a string
func (obj *flowIgmpv3GroupRecord) AuxData() string {

	return *obj.obj.AuxData

}

// Auxiliary data in hex, a multiple of 4 octets (maximum 255 32-bit words). IGMPv3 does not define any auxiliary data, so a conforming sender omits it (RFC 9776 Section 4.2.11).
// AuxData returns a string
func (obj *flowIgmpv3GroupRecord) HasAuxData() bool {
	return obj.obj.AuxData != nil
}

// Auxiliary data in hex, a multiple of 4 octets (maximum 255 32-bit words). IGMPv3 does not define any auxiliary data, so a conforming sender omits it (RFC 9776 Section 4.2.11).
// SetAuxData sets the string value in the FlowIgmpv3GroupRecord object
func (obj *flowIgmpv3GroupRecord) SetAuxData(value string) FlowIgmpv3GroupRecord {

	obj.obj.AuxData = &value
	return obj
}

func (obj *flowIgmpv3GroupRecord) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RecordType != nil {

		obj.RecordType().validateObj(vObj, set_default)
	}

	if obj.obj.AuxDataLength != nil {

		obj.AuxDataLength().validateObj(vObj, set_default)
	}

	if obj.obj.NumberOfSources != nil {

		obj.NumberOfSources().validateObj(vObj, set_default)
	}

	if obj.obj.MulticastAddress != nil {

		obj.MulticastAddress().validateObj(vObj, set_default)
	}

	if len(obj.obj.Sources) != 0 {

		if set_default {
			obj.Sources().clearHolderSlice()
			for _, item := range obj.obj.Sources {
				obj.Sources().appendHolderSlice(&flowIgmpv3Source{obj: item})
			}
		}
		for _, item := range obj.Sources().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.AuxData != nil {

		if len(*obj.obj.AuxData) > 2040 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of FlowIgmpv3GroupRecord.AuxData <= 2040 but Got %d",
					len(*obj.obj.AuxData)))
		}

	}

}

func (obj *flowIgmpv3GroupRecord) setDefault() {

}
