package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4OptionsTimestampFormat *****
type flowIpv4OptionsTimestampFormat struct {
	validation
	obj                                    *otg.FlowIpv4OptionsTimestampFormat
	marshaller                             marshalFlowIpv4OptionsTimestampFormat
	unMarshaller                           unMarshalFlowIpv4OptionsTimestampFormat
	timestampsHolder                       FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter
	addressAndTimestampsHolder             FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter
	prespecifiedAddressAndTimestampsHolder FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter
}

func NewFlowIpv4OptionsTimestampFormat() FlowIpv4OptionsTimestampFormat {
	obj := flowIpv4OptionsTimestampFormat{obj: &otg.FlowIpv4OptionsTimestampFormat{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4OptionsTimestampFormat) msg() *otg.FlowIpv4OptionsTimestampFormat {
	return obj.obj
}

func (obj *flowIpv4OptionsTimestampFormat) setMsg(msg *otg.FlowIpv4OptionsTimestampFormat) FlowIpv4OptionsTimestampFormat {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4OptionsTimestampFormat struct {
	obj *flowIpv4OptionsTimestampFormat
}

type marshalFlowIpv4OptionsTimestampFormat interface {
	// ToProto marshals FlowIpv4OptionsTimestampFormat to protobuf object *otg.FlowIpv4OptionsTimestampFormat
	ToProto() (*otg.FlowIpv4OptionsTimestampFormat, error)
	// ToPbText marshals FlowIpv4OptionsTimestampFormat to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4OptionsTimestampFormat to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4OptionsTimestampFormat to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv4OptionsTimestampFormat struct {
	obj *flowIpv4OptionsTimestampFormat
}

type unMarshalFlowIpv4OptionsTimestampFormat interface {
	// FromProto unmarshals FlowIpv4OptionsTimestampFormat from protobuf object *otg.FlowIpv4OptionsTimestampFormat
	FromProto(msg *otg.FlowIpv4OptionsTimestampFormat) (FlowIpv4OptionsTimestampFormat, error)
	// FromPbText unmarshals FlowIpv4OptionsTimestampFormat from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4OptionsTimestampFormat from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4OptionsTimestampFormat from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4OptionsTimestampFormat) Marshal() marshalFlowIpv4OptionsTimestampFormat {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4OptionsTimestampFormat{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4OptionsTimestampFormat) Unmarshal() unMarshalFlowIpv4OptionsTimestampFormat {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4OptionsTimestampFormat{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4OptionsTimestampFormat) ToProto() (*otg.FlowIpv4OptionsTimestampFormat, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4OptionsTimestampFormat) FromProto(msg *otg.FlowIpv4OptionsTimestampFormat) (FlowIpv4OptionsTimestampFormat, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4OptionsTimestampFormat) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampFormat) FromPbText(value string) error {
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

func (m *marshalflowIpv4OptionsTimestampFormat) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampFormat) FromYaml(value string) error {
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

func (m *marshalflowIpv4OptionsTimestampFormat) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampFormat) FromJson(value string) error {
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

func (obj *flowIpv4OptionsTimestampFormat) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsTimestampFormat) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsTimestampFormat) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4OptionsTimestampFormat) Clone() (FlowIpv4OptionsTimestampFormat, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4OptionsTimestampFormat()
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

func (obj *flowIpv4OptionsTimestampFormat) setNil() {
	obj.timestampsHolder = nil
	obj.addressAndTimestampsHolder = nil
	obj.prespecifiedAddressAndTimestampsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv4OptionsTimestampFormat is the format field defines the structural layout of the options data area and the specific recording logic routers must follow when appending timestamps and IP addresses.
type FlowIpv4OptionsTimestampFormat interface {
	Validation
	// msg marshals FlowIpv4OptionsTimestampFormat to protobuf object *otg.FlowIpv4OptionsTimestampFormat
	// and doesn't set defaults
	msg() *otg.FlowIpv4OptionsTimestampFormat
	// setMsg unmarshals FlowIpv4OptionsTimestampFormat from protobuf object *otg.FlowIpv4OptionsTimestampFormat
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4OptionsTimestampFormat) FlowIpv4OptionsTimestampFormat
	// provides marshal interface
	Marshal() marshalFlowIpv4OptionsTimestampFormat
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4OptionsTimestampFormat
	// validate validates FlowIpv4OptionsTimestampFormat
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4OptionsTimestampFormat, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowIpv4OptionsTimestampFormatChoiceEnum, set in FlowIpv4OptionsTimestampFormat
	Choice() FlowIpv4OptionsTimestampFormatChoiceEnum
	// setChoice assigns FlowIpv4OptionsTimestampFormatChoiceEnum provided by user to FlowIpv4OptionsTimestampFormat
	setChoice(value FlowIpv4OptionsTimestampFormatChoiceEnum) FlowIpv4OptionsTimestampFormat
	// HasChoice checks if Choice has been set in FlowIpv4OptionsTimestampFormat
	HasChoice() bool
	// Timestamps returns FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIterIter, set in FlowIpv4OptionsTimestampFormat
	Timestamps() FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter
	// AddressAndTimestamps returns FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIterIter, set in FlowIpv4OptionsTimestampFormat
	AddressAndTimestamps() FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter
	// PrespecifiedAddressAndTimestamps returns FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIterIter, set in FlowIpv4OptionsTimestampFormat
	PrespecifiedAddressAndTimestamps() FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter
	setNil()
}

type FlowIpv4OptionsTimestampFormatChoiceEnum string

// Enum of Choice on FlowIpv4OptionsTimestampFormat
var FlowIpv4OptionsTimestampFormatChoice = struct {
	TIMESTAMPS                          FlowIpv4OptionsTimestampFormatChoiceEnum
	ADDRESS_AND_TIMESTAMPS              FlowIpv4OptionsTimestampFormatChoiceEnum
	PRESPECIFIED_ADDRESS_AND_TIMESTAMPS FlowIpv4OptionsTimestampFormatChoiceEnum
}{
	TIMESTAMPS:                          FlowIpv4OptionsTimestampFormatChoiceEnum("timestamps"),
	ADDRESS_AND_TIMESTAMPS:              FlowIpv4OptionsTimestampFormatChoiceEnum("address_and_timestamps"),
	PRESPECIFIED_ADDRESS_AND_TIMESTAMPS: FlowIpv4OptionsTimestampFormatChoiceEnum("prespecified_address_and_timestamps"),
}

func (obj *flowIpv4OptionsTimestampFormat) Choice() FlowIpv4OptionsTimestampFormatChoiceEnum {
	return FlowIpv4OptionsTimestampFormatChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowIpv4OptionsTimestampFormat) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowIpv4OptionsTimestampFormat) setChoice(value FlowIpv4OptionsTimestampFormatChoiceEnum) FlowIpv4OptionsTimestampFormat {
	intValue, ok := otg.FlowIpv4OptionsTimestampFormat_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowIpv4OptionsTimestampFormatChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowIpv4OptionsTimestampFormat_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.PrespecifiedAddressAndTimestamps = nil
	obj.prespecifiedAddressAndTimestampsHolder = nil
	obj.obj.AddressAndTimestamps = nil
	obj.addressAndTimestampsHolder = nil
	obj.obj.Timestamps = nil
	obj.timestampsHolder = nil

	if value == FlowIpv4OptionsTimestampFormatChoice.TIMESTAMPS {
		obj.obj.Timestamps = []*otg.FlowIpv4OptionsTimestampFormatTimestamp{}
	}

	if value == FlowIpv4OptionsTimestampFormatChoice.ADDRESS_AND_TIMESTAMPS {
		obj.obj.AddressAndTimestamps = []*otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp{}
	}

	if value == FlowIpv4OptionsTimestampFormatChoice.PRESPECIFIED_ADDRESS_AND_TIMESTAMPS {
		obj.obj.PrespecifiedAddressAndTimestamps = []*otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp{}
	}

	return obj
}

// A recording mode where routers append only their 32-bit universal timestamps sequentially into the options data area.
// Timestamps returns a []FlowIpv4OptionsTimestampFormatTimestamp
func (obj *flowIpv4OptionsTimestampFormat) Timestamps() FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter {
	if len(obj.obj.Timestamps) == 0 {
		obj.setChoice(FlowIpv4OptionsTimestampFormatChoice.TIMESTAMPS)
	}
	if obj.timestampsHolder == nil {
		obj.timestampsHolder = newFlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter(&obj.obj.Timestamps).setMsg(obj)
	}
	return obj.timestampsHolder
}

type flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter struct {
	obj                                          *flowIpv4OptionsTimestampFormat
	flowIpv4OptionsTimestampFormatTimestampSlice []FlowIpv4OptionsTimestampFormatTimestamp
	fieldPtr                                     *[]*otg.FlowIpv4OptionsTimestampFormatTimestamp
}

func newFlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter(ptr *[]*otg.FlowIpv4OptionsTimestampFormatTimestamp) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter {
	return &flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter{fieldPtr: ptr}
}

type FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter interface {
	setMsg(*flowIpv4OptionsTimestampFormat) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter
	Items() []FlowIpv4OptionsTimestampFormatTimestamp
	Add() FlowIpv4OptionsTimestampFormatTimestamp
	Append(items ...FlowIpv4OptionsTimestampFormatTimestamp) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter
	Set(index int, newObj FlowIpv4OptionsTimestampFormatTimestamp) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter
	Clear() FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter
	clearHolderSlice() FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter
	appendHolderSlice(item FlowIpv4OptionsTimestampFormatTimestamp) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter
}

func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter) setMsg(msg *flowIpv4OptionsTimestampFormat) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowIpv4OptionsTimestampFormatTimestamp{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter) Items() []FlowIpv4OptionsTimestampFormatTimestamp {
	return obj.flowIpv4OptionsTimestampFormatTimestampSlice
}

func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter) Add() FlowIpv4OptionsTimestampFormatTimestamp {
	newObj := &otg.FlowIpv4OptionsTimestampFormatTimestamp{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowIpv4OptionsTimestampFormatTimestamp{obj: newObj}
	newLibObj.setDefault()
	obj.flowIpv4OptionsTimestampFormatTimestampSlice = append(obj.flowIpv4OptionsTimestampFormatTimestampSlice, newLibObj)
	return newLibObj
}

func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter) Append(items ...FlowIpv4OptionsTimestampFormatTimestamp) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowIpv4OptionsTimestampFormatTimestampSlice = append(obj.flowIpv4OptionsTimestampFormatTimestampSlice, item)
	}
	return obj
}

func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter) Set(index int, newObj FlowIpv4OptionsTimestampFormatTimestamp) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowIpv4OptionsTimestampFormatTimestampSlice[index] = newObj
	return obj
}
func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter) Clear() FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowIpv4OptionsTimestampFormatTimestamp{}
		obj.flowIpv4OptionsTimestampFormatTimestampSlice = []FlowIpv4OptionsTimestampFormatTimestamp{}
	}
	return obj
}
func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter) clearHolderSlice() FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter {
	if len(obj.flowIpv4OptionsTimestampFormatTimestampSlice) > 0 {
		obj.flowIpv4OptionsTimestampFormatTimestampSlice = []FlowIpv4OptionsTimestampFormatTimestamp{}
	}
	return obj
}
func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter) appendHolderSlice(item FlowIpv4OptionsTimestampFormatTimestamp) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatTimestampIter {
	obj.flowIpv4OptionsTimestampFormatTimestampSlice = append(obj.flowIpv4OptionsTimestampFormatTimestampSlice, item)
	return obj
}

// A recording mode where each router appends its 32-bit IPv4 address followed immediately by its 32-bit universal timestamp into the options data area.
// AddressAndTimestamps returns a []FlowIpv4OptionsTimestampFormatAddressAndTimestamp
func (obj *flowIpv4OptionsTimestampFormat) AddressAndTimestamps() FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter {
	if len(obj.obj.AddressAndTimestamps) == 0 {
		obj.setChoice(FlowIpv4OptionsTimestampFormatChoice.ADDRESS_AND_TIMESTAMPS)
	}
	if obj.addressAndTimestampsHolder == nil {
		obj.addressAndTimestampsHolder = newFlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter(&obj.obj.AddressAndTimestamps).setMsg(obj)
	}
	return obj.addressAndTimestampsHolder
}

type flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter struct {
	obj                                                    *flowIpv4OptionsTimestampFormat
	flowIpv4OptionsTimestampFormatAddressAndTimestampSlice []FlowIpv4OptionsTimestampFormatAddressAndTimestamp
	fieldPtr                                               *[]*otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp
}

func newFlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter(ptr *[]*otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter {
	return &flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter{fieldPtr: ptr}
}

type FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter interface {
	setMsg(*flowIpv4OptionsTimestampFormat) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter
	Items() []FlowIpv4OptionsTimestampFormatAddressAndTimestamp
	Add() FlowIpv4OptionsTimestampFormatAddressAndTimestamp
	Append(items ...FlowIpv4OptionsTimestampFormatAddressAndTimestamp) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter
	Set(index int, newObj FlowIpv4OptionsTimestampFormatAddressAndTimestamp) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter
	Clear() FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter
	clearHolderSlice() FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter
	appendHolderSlice(item FlowIpv4OptionsTimestampFormatAddressAndTimestamp) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter
}

func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter) setMsg(msg *flowIpv4OptionsTimestampFormat) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowIpv4OptionsTimestampFormatAddressAndTimestamp{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter) Items() []FlowIpv4OptionsTimestampFormatAddressAndTimestamp {
	return obj.flowIpv4OptionsTimestampFormatAddressAndTimestampSlice
}

func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter) Add() FlowIpv4OptionsTimestampFormatAddressAndTimestamp {
	newObj := &otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowIpv4OptionsTimestampFormatAddressAndTimestamp{obj: newObj}
	newLibObj.setDefault()
	obj.flowIpv4OptionsTimestampFormatAddressAndTimestampSlice = append(obj.flowIpv4OptionsTimestampFormatAddressAndTimestampSlice, newLibObj)
	return newLibObj
}

func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter) Append(items ...FlowIpv4OptionsTimestampFormatAddressAndTimestamp) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowIpv4OptionsTimestampFormatAddressAndTimestampSlice = append(obj.flowIpv4OptionsTimestampFormatAddressAndTimestampSlice, item)
	}
	return obj
}

func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter) Set(index int, newObj FlowIpv4OptionsTimestampFormatAddressAndTimestamp) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowIpv4OptionsTimestampFormatAddressAndTimestampSlice[index] = newObj
	return obj
}
func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter) Clear() FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp{}
		obj.flowIpv4OptionsTimestampFormatAddressAndTimestampSlice = []FlowIpv4OptionsTimestampFormatAddressAndTimestamp{}
	}
	return obj
}
func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter) clearHolderSlice() FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter {
	if len(obj.flowIpv4OptionsTimestampFormatAddressAndTimestampSlice) > 0 {
		obj.flowIpv4OptionsTimestampFormatAddressAndTimestampSlice = []FlowIpv4OptionsTimestampFormatAddressAndTimestamp{}
	}
	return obj
}
func (obj *flowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter) appendHolderSlice(item FlowIpv4OptionsTimestampFormatAddressAndTimestamp) FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter {
	obj.flowIpv4OptionsTimestampFormatAddressAndTimestampSlice = append(obj.flowIpv4OptionsTimestampFormatAddressAndTimestampSlice, item)
	return obj
}

// A selective recording mode where only routers whose IPv4 addresses match the pre-listed entries in the header record their 32-bit universal timestamps into the options data area.
// PrespecifiedAddressAndTimestamps returns a []FlowIpv4OptionsTimestampFormatAddressAndTimestamp
func (obj *flowIpv4OptionsTimestampFormat) PrespecifiedAddressAndTimestamps() FlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter {
	if len(obj.obj.PrespecifiedAddressAndTimestamps) == 0 {
		obj.setChoice(FlowIpv4OptionsTimestampFormatChoice.PRESPECIFIED_ADDRESS_AND_TIMESTAMPS)
	}
	if obj.prespecifiedAddressAndTimestampsHolder == nil {
		obj.prespecifiedAddressAndTimestampsHolder = newFlowIpv4OptionsTimestampFormatFlowIpv4OptionsTimestampFormatAddressAndTimestampIter(&obj.obj.PrespecifiedAddressAndTimestamps).setMsg(obj)
	}
	return obj.prespecifiedAddressAndTimestampsHolder
}

func (obj *flowIpv4OptionsTimestampFormat) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Timestamps) != 0 {

		if set_default {
			obj.Timestamps().clearHolderSlice()
			for _, item := range obj.obj.Timestamps {
				obj.Timestamps().appendHolderSlice(&flowIpv4OptionsTimestampFormatTimestamp{obj: item})
			}
		}
		for _, item := range obj.Timestamps().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.AddressAndTimestamps) != 0 {

		if set_default {
			obj.AddressAndTimestamps().clearHolderSlice()
			for _, item := range obj.obj.AddressAndTimestamps {
				obj.AddressAndTimestamps().appendHolderSlice(&flowIpv4OptionsTimestampFormatAddressAndTimestamp{obj: item})
			}
		}
		for _, item := range obj.AddressAndTimestamps().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.PrespecifiedAddressAndTimestamps) != 0 {

		if set_default {
			obj.PrespecifiedAddressAndTimestamps().clearHolderSlice()
			for _, item := range obj.obj.PrespecifiedAddressAndTimestamps {
				obj.PrespecifiedAddressAndTimestamps().appendHolderSlice(&flowIpv4OptionsTimestampFormatAddressAndTimestamp{obj: item})
			}
		}
		for _, item := range obj.PrespecifiedAddressAndTimestamps().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowIpv4OptionsTimestampFormat) setDefault() {
	var choices_set int = 0
	var choice FlowIpv4OptionsTimestampFormatChoiceEnum

	if len(obj.obj.Timestamps) > 0 {
		choices_set += 1
		choice = FlowIpv4OptionsTimestampFormatChoice.TIMESTAMPS
	}

	if len(obj.obj.AddressAndTimestamps) > 0 {
		choices_set += 1
		choice = FlowIpv4OptionsTimestampFormatChoice.ADDRESS_AND_TIMESTAMPS
	}

	if len(obj.obj.PrespecifiedAddressAndTimestamps) > 0 {
		choices_set += 1
		choice = FlowIpv4OptionsTimestampFormatChoice.PRESPECIFIED_ADDRESS_AND_TIMESTAMPS
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowIpv4OptionsTimestampFormat")
			}
		} else {
			intVal := otg.FlowIpv4OptionsTimestampFormat_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowIpv4OptionsTimestampFormat_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
