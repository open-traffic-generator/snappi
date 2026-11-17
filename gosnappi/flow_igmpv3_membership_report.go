package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIgmpv3MembershipReport *****
type flowIgmpv3MembershipReport struct {
	validation
	obj                        *otg.FlowIgmpv3MembershipReport
	marshaller                 marshalFlowIgmpv3MembershipReport
	unMarshaller               unMarshalFlowIgmpv3MembershipReport
	typeHolder                 PatternFlowIgmpv3MembershipReportType
	reservedHolder             PatternFlowIgmpv3MembershipReportReserved
	checksumHolder             PatternFlowIgmpv3MembershipReportChecksum
	flagsHolder                PatternFlowIgmpv3MembershipReportFlags
	numberOfGroupRecordsHolder PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	groupRecordsHolder         FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter
}

func NewFlowIgmpv3MembershipReport() FlowIgmpv3MembershipReport {
	obj := flowIgmpv3MembershipReport{obj: &otg.FlowIgmpv3MembershipReport{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIgmpv3MembershipReport) msg() *otg.FlowIgmpv3MembershipReport {
	return obj.obj
}

func (obj *flowIgmpv3MembershipReport) setMsg(msg *otg.FlowIgmpv3MembershipReport) FlowIgmpv3MembershipReport {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIgmpv3MembershipReport struct {
	obj *flowIgmpv3MembershipReport
}

type marshalFlowIgmpv3MembershipReport interface {
	// ToProto marshals FlowIgmpv3MembershipReport to protobuf object *otg.FlowIgmpv3MembershipReport
	ToProto() (*otg.FlowIgmpv3MembershipReport, error)
	// ToPbText marshals FlowIgmpv3MembershipReport to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIgmpv3MembershipReport to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIgmpv3MembershipReport to JSON text
	ToJson() (string, error)
}

type unMarshalflowIgmpv3MembershipReport struct {
	obj *flowIgmpv3MembershipReport
}

type unMarshalFlowIgmpv3MembershipReport interface {
	// FromProto unmarshals FlowIgmpv3MembershipReport from protobuf object *otg.FlowIgmpv3MembershipReport
	FromProto(msg *otg.FlowIgmpv3MembershipReport) (FlowIgmpv3MembershipReport, error)
	// FromPbText unmarshals FlowIgmpv3MembershipReport from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIgmpv3MembershipReport from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIgmpv3MembershipReport from JSON text
	FromJson(value string) error
}

func (obj *flowIgmpv3MembershipReport) Marshal() marshalFlowIgmpv3MembershipReport {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIgmpv3MembershipReport{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIgmpv3MembershipReport) Unmarshal() unMarshalFlowIgmpv3MembershipReport {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIgmpv3MembershipReport{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIgmpv3MembershipReport) ToProto() (*otg.FlowIgmpv3MembershipReport, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIgmpv3MembershipReport) FromProto(msg *otg.FlowIgmpv3MembershipReport) (FlowIgmpv3MembershipReport, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIgmpv3MembershipReport) ToPbText() (string, error) {
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

func (m *unMarshalflowIgmpv3MembershipReport) FromPbText(value string) error {
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

func (m *marshalflowIgmpv3MembershipReport) ToYaml() (string, error) {
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

func (m *unMarshalflowIgmpv3MembershipReport) FromYaml(value string) error {
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

func (m *marshalflowIgmpv3MembershipReport) ToJson() (string, error) {
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

func (m *unMarshalflowIgmpv3MembershipReport) FromJson(value string) error {
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

func (obj *flowIgmpv3MembershipReport) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIgmpv3MembershipReport) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIgmpv3MembershipReport) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIgmpv3MembershipReport) Clone() (FlowIgmpv3MembershipReport, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIgmpv3MembershipReport()
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

func (obj *flowIgmpv3MembershipReport) setNil() {
	obj.typeHolder = nil
	obj.reservedHolder = nil
	obj.checksumHolder = nil
	obj.flagsHolder = nil
	obj.numberOfGroupRecordsHolder = nil
	obj.groupRecordsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIgmpv3MembershipReport is iGMPv3 Version 3 Membership Report message (RFC 9776 Section 4.2).
type FlowIgmpv3MembershipReport interface {
	Validation
	// msg marshals FlowIgmpv3MembershipReport to protobuf object *otg.FlowIgmpv3MembershipReport
	// and doesn't set defaults
	msg() *otg.FlowIgmpv3MembershipReport
	// setMsg unmarshals FlowIgmpv3MembershipReport from protobuf object *otg.FlowIgmpv3MembershipReport
	// and doesn't set defaults
	setMsg(*otg.FlowIgmpv3MembershipReport) FlowIgmpv3MembershipReport
	// provides marshal interface
	Marshal() marshalFlowIgmpv3MembershipReport
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIgmpv3MembershipReport
	// validate validates FlowIgmpv3MembershipReport
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIgmpv3MembershipReport, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowIgmpv3MembershipReportType, set in FlowIgmpv3MembershipReport.
	// PatternFlowIgmpv3MembershipReportType is iGMP message type, 34 (0x22) for a Version 3 Membership Report.
	Type() PatternFlowIgmpv3MembershipReportType
	// SetType assigns PatternFlowIgmpv3MembershipReportType provided by user to FlowIgmpv3MembershipReport.
	// PatternFlowIgmpv3MembershipReportType is iGMP message type, 34 (0x22) for a Version 3 Membership Report.
	SetType(value PatternFlowIgmpv3MembershipReportType) FlowIgmpv3MembershipReport
	// HasType checks if Type has been set in FlowIgmpv3MembershipReport
	HasType() bool
	// Reserved returns PatternFlowIgmpv3MembershipReportReserved, set in FlowIgmpv3MembershipReport.
	// PatternFlowIgmpv3MembershipReportReserved is reserved field, set to zero by the sender and ignored by the receiver.
	Reserved() PatternFlowIgmpv3MembershipReportReserved
	// SetReserved assigns PatternFlowIgmpv3MembershipReportReserved provided by user to FlowIgmpv3MembershipReport.
	// PatternFlowIgmpv3MembershipReportReserved is reserved field, set to zero by the sender and ignored by the receiver.
	SetReserved(value PatternFlowIgmpv3MembershipReportReserved) FlowIgmpv3MembershipReport
	// HasReserved checks if Reserved has been set in FlowIgmpv3MembershipReport
	HasReserved() bool
	// Checksum returns PatternFlowIgmpv3MembershipReportChecksum, set in FlowIgmpv3MembershipReport.
	// PatternFlowIgmpv3MembershipReportChecksum is checksum of the whole IGMP message, a 16-bit one's complement sum with no pseudo-header.
	Checksum() PatternFlowIgmpv3MembershipReportChecksum
	// SetChecksum assigns PatternFlowIgmpv3MembershipReportChecksum provided by user to FlowIgmpv3MembershipReport.
	// PatternFlowIgmpv3MembershipReportChecksum is checksum of the whole IGMP message, a 16-bit one's complement sum with no pseudo-header.
	SetChecksum(value PatternFlowIgmpv3MembershipReportChecksum) FlowIgmpv3MembershipReport
	// HasChecksum checks if Checksum has been set in FlowIgmpv3MembershipReport
	HasChecksum() bool
	// Flags returns PatternFlowIgmpv3MembershipReportFlags, set in FlowIgmpv3MembershipReport.
	// PatternFlowIgmpv3MembershipReportFlags is flags field of the report, 16 bits that are managed by IANA and are set to zero by a conforming sender.
	Flags() PatternFlowIgmpv3MembershipReportFlags
	// SetFlags assigns PatternFlowIgmpv3MembershipReportFlags provided by user to FlowIgmpv3MembershipReport.
	// PatternFlowIgmpv3MembershipReportFlags is flags field of the report, 16 bits that are managed by IANA and are set to zero by a conforming sender.
	SetFlags(value PatternFlowIgmpv3MembershipReportFlags) FlowIgmpv3MembershipReport
	// HasFlags checks if Flags has been set in FlowIgmpv3MembershipReport
	HasFlags() bool
	// NumberOfGroupRecords returns PatternFlowIgmpv3MembershipReportNumberOfGroupRecords, set in FlowIgmpv3MembershipReport.
	// PatternFlowIgmpv3MembershipReportNumberOfGroupRecords is number of group records. When auto, it is set to the number of entries in group_records.
	NumberOfGroupRecords() PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	// SetNumberOfGroupRecords assigns PatternFlowIgmpv3MembershipReportNumberOfGroupRecords provided by user to FlowIgmpv3MembershipReport.
	// PatternFlowIgmpv3MembershipReportNumberOfGroupRecords is number of group records. When auto, it is set to the number of entries in group_records.
	SetNumberOfGroupRecords(value PatternFlowIgmpv3MembershipReportNumberOfGroupRecords) FlowIgmpv3MembershipReport
	// HasNumberOfGroupRecords checks if NumberOfGroupRecords has been set in FlowIgmpv3MembershipReport
	HasNumberOfGroupRecords() bool
	// GroupRecords returns FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIterIter, set in FlowIgmpv3MembershipReport
	GroupRecords() FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter
	setNil()
}

// description is TBD
// Type returns a PatternFlowIgmpv3MembershipReportType
func (obj *flowIgmpv3MembershipReport) Type() PatternFlowIgmpv3MembershipReportType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowIgmpv3MembershipReportType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowIgmpv3MembershipReportType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowIgmpv3MembershipReportType
func (obj *flowIgmpv3MembershipReport) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowIgmpv3MembershipReportType value in the FlowIgmpv3MembershipReport object
func (obj *flowIgmpv3MembershipReport) SetType(value PatternFlowIgmpv3MembershipReportType) FlowIgmpv3MembershipReport {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Reserved returns a PatternFlowIgmpv3MembershipReportReserved
func (obj *flowIgmpv3MembershipReport) Reserved() PatternFlowIgmpv3MembershipReportReserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowIgmpv3MembershipReportReserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowIgmpv3MembershipReportReserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowIgmpv3MembershipReportReserved
func (obj *flowIgmpv3MembershipReport) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowIgmpv3MembershipReportReserved value in the FlowIgmpv3MembershipReport object
func (obj *flowIgmpv3MembershipReport) SetReserved(value PatternFlowIgmpv3MembershipReportReserved) FlowIgmpv3MembershipReport {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

// description is TBD
// Checksum returns a PatternFlowIgmpv3MembershipReportChecksum
func (obj *flowIgmpv3MembershipReport) Checksum() PatternFlowIgmpv3MembershipReportChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = NewPatternFlowIgmpv3MembershipReportChecksum().msg()
	}
	if obj.checksumHolder == nil {
		obj.checksumHolder = &patternFlowIgmpv3MembershipReportChecksum{obj: obj.obj.Checksum}
	}
	return obj.checksumHolder
}

// description is TBD
// Checksum returns a PatternFlowIgmpv3MembershipReportChecksum
func (obj *flowIgmpv3MembershipReport) HasChecksum() bool {
	return obj.obj.Checksum != nil
}

// description is TBD
// SetChecksum sets the PatternFlowIgmpv3MembershipReportChecksum value in the FlowIgmpv3MembershipReport object
func (obj *flowIgmpv3MembershipReport) SetChecksum(value PatternFlowIgmpv3MembershipReportChecksum) FlowIgmpv3MembershipReport {

	obj.checksumHolder = nil
	obj.obj.Checksum = value.msg()

	return obj
}

// description is TBD
// Flags returns a PatternFlowIgmpv3MembershipReportFlags
func (obj *flowIgmpv3MembershipReport) Flags() PatternFlowIgmpv3MembershipReportFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewPatternFlowIgmpv3MembershipReportFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &patternFlowIgmpv3MembershipReportFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// description is TBD
// Flags returns a PatternFlowIgmpv3MembershipReportFlags
func (obj *flowIgmpv3MembershipReport) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the PatternFlowIgmpv3MembershipReportFlags value in the FlowIgmpv3MembershipReport object
func (obj *flowIgmpv3MembershipReport) SetFlags(value PatternFlowIgmpv3MembershipReportFlags) FlowIgmpv3MembershipReport {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// description is TBD
// NumberOfGroupRecords returns a PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
func (obj *flowIgmpv3MembershipReport) NumberOfGroupRecords() PatternFlowIgmpv3MembershipReportNumberOfGroupRecords {
	if obj.obj.NumberOfGroupRecords == nil {
		obj.obj.NumberOfGroupRecords = NewPatternFlowIgmpv3MembershipReportNumberOfGroupRecords().msg()
	}
	if obj.numberOfGroupRecordsHolder == nil {
		obj.numberOfGroupRecordsHolder = &patternFlowIgmpv3MembershipReportNumberOfGroupRecords{obj: obj.obj.NumberOfGroupRecords}
	}
	return obj.numberOfGroupRecordsHolder
}

// description is TBD
// NumberOfGroupRecords returns a PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
func (obj *flowIgmpv3MembershipReport) HasNumberOfGroupRecords() bool {
	return obj.obj.NumberOfGroupRecords != nil
}

// description is TBD
// SetNumberOfGroupRecords sets the PatternFlowIgmpv3MembershipReportNumberOfGroupRecords value in the FlowIgmpv3MembershipReport object
func (obj *flowIgmpv3MembershipReport) SetNumberOfGroupRecords(value PatternFlowIgmpv3MembershipReportNumberOfGroupRecords) FlowIgmpv3MembershipReport {

	obj.numberOfGroupRecordsHolder = nil
	obj.obj.NumberOfGroupRecords = value.msg()

	return obj
}

// List of group records, each one reports the membership state of one multicast group.
// GroupRecords returns a []FlowIgmpv3GroupRecord
func (obj *flowIgmpv3MembershipReport) GroupRecords() FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter {
	if len(obj.obj.GroupRecords) == 0 {
		obj.obj.GroupRecords = []*otg.FlowIgmpv3GroupRecord{}
	}
	if obj.groupRecordsHolder == nil {
		obj.groupRecordsHolder = newFlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter(&obj.obj.GroupRecords).setMsg(obj)
	}
	return obj.groupRecordsHolder
}

type flowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter struct {
	obj                        *flowIgmpv3MembershipReport
	flowIgmpv3GroupRecordSlice []FlowIgmpv3GroupRecord
	fieldPtr                   *[]*otg.FlowIgmpv3GroupRecord
}

func newFlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter(ptr *[]*otg.FlowIgmpv3GroupRecord) FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter {
	return &flowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter{fieldPtr: ptr}
}

type FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter interface {
	setMsg(*flowIgmpv3MembershipReport) FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter
	Items() []FlowIgmpv3GroupRecord
	Add() FlowIgmpv3GroupRecord
	Append(items ...FlowIgmpv3GroupRecord) FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter
	Set(index int, newObj FlowIgmpv3GroupRecord) FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter
	Clear() FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter
	clearHolderSlice() FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter
	appendHolderSlice(item FlowIgmpv3GroupRecord) FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter
}

func (obj *flowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter) setMsg(msg *flowIgmpv3MembershipReport) FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowIgmpv3GroupRecord{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter) Items() []FlowIgmpv3GroupRecord {
	return obj.flowIgmpv3GroupRecordSlice
}

func (obj *flowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter) Add() FlowIgmpv3GroupRecord {
	newObj := &otg.FlowIgmpv3GroupRecord{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowIgmpv3GroupRecord{obj: newObj}
	newLibObj.setDefault()
	obj.flowIgmpv3GroupRecordSlice = append(obj.flowIgmpv3GroupRecordSlice, newLibObj)
	return newLibObj
}

func (obj *flowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter) Append(items ...FlowIgmpv3GroupRecord) FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowIgmpv3GroupRecordSlice = append(obj.flowIgmpv3GroupRecordSlice, item)
	}
	return obj
}

func (obj *flowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter) Set(index int, newObj FlowIgmpv3GroupRecord) FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowIgmpv3GroupRecordSlice[index] = newObj
	return obj
}
func (obj *flowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter) Clear() FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowIgmpv3GroupRecord{}
		obj.flowIgmpv3GroupRecordSlice = []FlowIgmpv3GroupRecord{}
	}
	return obj
}
func (obj *flowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter) clearHolderSlice() FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter {
	if len(obj.flowIgmpv3GroupRecordSlice) > 0 {
		obj.flowIgmpv3GroupRecordSlice = []FlowIgmpv3GroupRecord{}
	}
	return obj
}
func (obj *flowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter) appendHolderSlice(item FlowIgmpv3GroupRecord) FlowIgmpv3MembershipReportFlowIgmpv3GroupRecordIter {
	obj.flowIgmpv3GroupRecordSlice = append(obj.flowIgmpv3GroupRecordSlice, item)
	return obj
}

func (obj *flowIgmpv3MembershipReport) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(vObj, set_default)
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(vObj, set_default)
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.NumberOfGroupRecords != nil {

		obj.NumberOfGroupRecords().validateObj(vObj, set_default)
	}

	if len(obj.obj.GroupRecords) != 0 {

		if set_default {
			obj.GroupRecords().clearHolderSlice()
			for _, item := range obj.obj.GroupRecords {
				obj.GroupRecords().appendHolderSlice(&flowIgmpv3GroupRecord{obj: item})
			}
		}
		for _, item := range obj.GroupRecords().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowIgmpv3MembershipReport) setDefault() {

}
