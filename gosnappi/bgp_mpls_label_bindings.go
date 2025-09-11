package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpMplsLabelBindings *****
type bgpMplsLabelBindings struct {
	validation
	obj                          *otg.BgpMplsLabelBindings
	marshaller                   marshalBgpMplsLabelBindings
	unMarshaller                 unMarshalBgpMplsLabelBindings
	singleLabelValueHolder       BgpMplsLabelBindingsSingleLabelValue
	multipleLabelValuesHolder    BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter
	singleSrLabelIndexHolder     BgpMplsLabelBindingsSingleLabelIndex
	multipleSrLabelIndicesHolder BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter
}

func NewBgpMplsLabelBindings() BgpMplsLabelBindings {
	obj := bgpMplsLabelBindings{obj: &otg.BgpMplsLabelBindings{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpMplsLabelBindings) msg() *otg.BgpMplsLabelBindings {
	return obj.obj
}

func (obj *bgpMplsLabelBindings) setMsg(msg *otg.BgpMplsLabelBindings) BgpMplsLabelBindings {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpMplsLabelBindings struct {
	obj *bgpMplsLabelBindings
}

type marshalBgpMplsLabelBindings interface {
	// ToProto marshals BgpMplsLabelBindings to protobuf object *otg.BgpMplsLabelBindings
	ToProto() (*otg.BgpMplsLabelBindings, error)
	// ToPbText marshals BgpMplsLabelBindings to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpMplsLabelBindings to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpMplsLabelBindings to JSON text
	ToJson() (string, error)
}

type unMarshalbgpMplsLabelBindings struct {
	obj *bgpMplsLabelBindings
}

type unMarshalBgpMplsLabelBindings interface {
	// FromProto unmarshals BgpMplsLabelBindings from protobuf object *otg.BgpMplsLabelBindings
	FromProto(msg *otg.BgpMplsLabelBindings) (BgpMplsLabelBindings, error)
	// FromPbText unmarshals BgpMplsLabelBindings from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpMplsLabelBindings from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpMplsLabelBindings from JSON text
	FromJson(value string) error
}

func (obj *bgpMplsLabelBindings) Marshal() marshalBgpMplsLabelBindings {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpMplsLabelBindings{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpMplsLabelBindings) Unmarshal() unMarshalBgpMplsLabelBindings {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpMplsLabelBindings{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpMplsLabelBindings) ToProto() (*otg.BgpMplsLabelBindings, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpMplsLabelBindings) FromProto(msg *otg.BgpMplsLabelBindings) (BgpMplsLabelBindings, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpMplsLabelBindings) ToPbText() (string, error) {
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

func (m *unMarshalbgpMplsLabelBindings) FromPbText(value string) error {
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

func (m *marshalbgpMplsLabelBindings) ToYaml() (string, error) {
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

func (m *unMarshalbgpMplsLabelBindings) FromYaml(value string) error {
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

func (m *marshalbgpMplsLabelBindings) ToJson() (string, error) {
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

func (m *unMarshalbgpMplsLabelBindings) FromJson(value string) error {
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

func (obj *bgpMplsLabelBindings) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpMplsLabelBindings) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpMplsLabelBindings) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpMplsLabelBindings) Clone() (BgpMplsLabelBindings, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpMplsLabelBindings()
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

func (obj *bgpMplsLabelBindings) setNil() {
	obj.singleLabelValueHolder = nil
	obj.multipleLabelValuesHolder = nil
	obj.singleSrLabelIndexHolder = nil
	obj.multipleSrLabelIndicesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpMplsLabelBindings is bGP may be used to advertise that a particular node (N) has bound a particular MPLS label, or a particular sequence of MPLS labels,
// to a particular address prefix.
// This is done by sending a Multiprotocol BGP UPDATE message with with an MP_REACH_NLRI attribute.
// The Network Address of Next Hop field of that attribute contains an IP address of node N.
// References: https://datatracker.ietf.org/doc/html/rfc3107
// & https://datatracker.ietf.org/doc/html/rfc8277.
type BgpMplsLabelBindings interface {
	Validation
	// msg marshals BgpMplsLabelBindings to protobuf object *otg.BgpMplsLabelBindings
	// and doesn't set defaults
	msg() *otg.BgpMplsLabelBindings
	// setMsg unmarshals BgpMplsLabelBindings from protobuf object *otg.BgpMplsLabelBindings
	// and doesn't set defaults
	setMsg(*otg.BgpMplsLabelBindings) BgpMplsLabelBindings
	// provides marshal interface
	Marshal() marshalBgpMplsLabelBindings
	// provides unmarshal interface
	Unmarshal() unMarshalBgpMplsLabelBindings
	// validate validates BgpMplsLabelBindings
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpMplsLabelBindings, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpMplsLabelBindingsChoiceEnum, set in BgpMplsLabelBindings
	Choice() BgpMplsLabelBindingsChoiceEnum
	// setChoice assigns BgpMplsLabelBindingsChoiceEnum provided by user to BgpMplsLabelBindings
	setChoice(value BgpMplsLabelBindingsChoiceEnum) BgpMplsLabelBindings
	// HasChoice checks if Choice has been set in BgpMplsLabelBindings
	HasChoice() bool
	// getter for MutipleLabelValues to set choice.
	MutipleLabelValues()
	// SingleLabelValue returns BgpMplsLabelBindingsSingleLabelValue, set in BgpMplsLabelBindings.
	// BgpMplsLabelBindingsSingleLabelValue is container for Single MPLS Value.
	SingleLabelValue() BgpMplsLabelBindingsSingleLabelValue
	// SetSingleLabelValue assigns BgpMplsLabelBindingsSingleLabelValue provided by user to BgpMplsLabelBindings.
	// BgpMplsLabelBindingsSingleLabelValue is container for Single MPLS Value.
	SetSingleLabelValue(value BgpMplsLabelBindingsSingleLabelValue) BgpMplsLabelBindings
	// HasSingleLabelValue checks if SingleLabelValue has been set in BgpMplsLabelBindings
	HasSingleLabelValue() bool
	// MultipleLabelValues returns BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIterIter, set in BgpMplsLabelBindings
	MultipleLabelValues() BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter
	// SingleSrLabelIndex returns BgpMplsLabelBindingsSingleLabelIndex, set in BgpMplsLabelBindings.
	// BgpMplsLabelBindingsSingleLabelIndex is container for Single MPLS SR Index.
	SingleSrLabelIndex() BgpMplsLabelBindingsSingleLabelIndex
	// SetSingleSrLabelIndex assigns BgpMplsLabelBindingsSingleLabelIndex provided by user to BgpMplsLabelBindings.
	// BgpMplsLabelBindingsSingleLabelIndex is container for Single MPLS SR Index.
	SetSingleSrLabelIndex(value BgpMplsLabelBindingsSingleLabelIndex) BgpMplsLabelBindings
	// HasSingleSrLabelIndex checks if SingleSrLabelIndex has been set in BgpMplsLabelBindings
	HasSingleSrLabelIndex() bool
	// MultipleSrLabelIndices returns BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIterIter, set in BgpMplsLabelBindings
	MultipleSrLabelIndices() BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter
	setNil()
}

type BgpMplsLabelBindingsChoiceEnum string

// Enum of Choice on BgpMplsLabelBindings
var BgpMplsLabelBindingsChoice = struct {
	SINGLE_LABEL_VALUE        BgpMplsLabelBindingsChoiceEnum
	MUTIPLE_LABEL_VALUES      BgpMplsLabelBindingsChoiceEnum
	SINGLE_SR_LABEL_INDEX     BgpMplsLabelBindingsChoiceEnum
	MULTIPLE_SR_LABEL_INDICES BgpMplsLabelBindingsChoiceEnum
}{
	SINGLE_LABEL_VALUE:        BgpMplsLabelBindingsChoiceEnum("single_label_value"),
	MUTIPLE_LABEL_VALUES:      BgpMplsLabelBindingsChoiceEnum("mutiple_label_values"),
	SINGLE_SR_LABEL_INDEX:     BgpMplsLabelBindingsChoiceEnum("single_sr_label_index"),
	MULTIPLE_SR_LABEL_INDICES: BgpMplsLabelBindingsChoiceEnum("multiple_sr_label_indices"),
}

func (obj *bgpMplsLabelBindings) Choice() BgpMplsLabelBindingsChoiceEnum {
	return BgpMplsLabelBindingsChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for MutipleLabelValues to set choice
func (obj *bgpMplsLabelBindings) MutipleLabelValues() {
	obj.setChoice(BgpMplsLabelBindingsChoice.MUTIPLE_LABEL_VALUES)
}

// Bind one or more MPLS Labels to Address Prefixes.
// - single_label: In this case,  mpls_lebel_capability under BGP capability is chosen as single_label. Single Label Value is to be configure here.
// - mutiple_label_values: The mpls_lebel_capability under BGP capability is chosen as multipl_labels. Multiple Label Values are to be configure here.
// - single_sr_label_index: The mpls_lebel_capability under BGP capability is chosen as single_label. At least one SRGB range has been configured in mpls_srgb_ranges.
// Single Label index is to be configure here.
// - multiple_sr_label_indices: The mpls_lebel_capability under BGP capability is chosen as multipl_labels. At least one SRGB range has been configured in mpls_srgb_ranges.
// Multiple Label indices are to be configure here.
// Choice returns a string
func (obj *bgpMplsLabelBindings) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpMplsLabelBindings) setChoice(value BgpMplsLabelBindingsChoiceEnum) BgpMplsLabelBindings {
	intValue, ok := otg.BgpMplsLabelBindings_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpMplsLabelBindingsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpMplsLabelBindings_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.MultipleSrLabelIndices = nil
	obj.multipleSrLabelIndicesHolder = nil
	obj.obj.SingleSrLabelIndex = nil
	obj.singleSrLabelIndexHolder = nil
	obj.obj.SingleLabelValue = nil
	obj.singleLabelValueHolder = nil

	if value == BgpMplsLabelBindingsChoice.SINGLE_LABEL_VALUE {
		obj.obj.SingleLabelValue = NewBgpMplsLabelBindingsSingleLabelValue().msg()
	}

	if value == BgpMplsLabelBindingsChoice.SINGLE_SR_LABEL_INDEX {
		obj.obj.SingleSrLabelIndex = NewBgpMplsLabelBindingsSingleLabelIndex().msg()
	}

	if value == BgpMplsLabelBindingsChoice.MULTIPLE_SR_LABEL_INDICES {
		obj.obj.MultipleSrLabelIndices = []*otg.BgpMplsLabelBindingsSingleLabelIndex{}
	}

	return obj
}

// Group of MPLS Label value bind to a prefix
// SingleLabelValue returns a BgpMplsLabelBindingsSingleLabelValue
func (obj *bgpMplsLabelBindings) SingleLabelValue() BgpMplsLabelBindingsSingleLabelValue {
	if obj.obj.SingleLabelValue == nil {
		obj.setChoice(BgpMplsLabelBindingsChoice.SINGLE_LABEL_VALUE)
	}
	if obj.singleLabelValueHolder == nil {
		obj.singleLabelValueHolder = &bgpMplsLabelBindingsSingleLabelValue{obj: obj.obj.SingleLabelValue}
	}
	return obj.singleLabelValueHolder
}

// Group of MPLS Label value bind to a prefix
// SingleLabelValue returns a BgpMplsLabelBindingsSingleLabelValue
func (obj *bgpMplsLabelBindings) HasSingleLabelValue() bool {
	return obj.obj.SingleLabelValue != nil
}

// Group of MPLS Label value bind to a prefix
// SetSingleLabelValue sets the BgpMplsLabelBindingsSingleLabelValue value in the BgpMplsLabelBindings object
func (obj *bgpMplsLabelBindings) SetSingleLabelValue(value BgpMplsLabelBindingsSingleLabelValue) BgpMplsLabelBindings {
	obj.setChoice(BgpMplsLabelBindingsChoice.SINGLE_LABEL_VALUE)
	obj.singleLabelValueHolder = nil
	obj.obj.SingleLabelValue = value.msg()

	return obj
}

// A list of group of MPLS Label values bind to a prefix
// MultipleLabelValues returns a []BgpMplsLabelBindingsSingleLabelValue
func (obj *bgpMplsLabelBindings) MultipleLabelValues() BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter {
	if len(obj.obj.MultipleLabelValues) == 0 {
		obj.obj.MultipleLabelValues = []*otg.BgpMplsLabelBindingsSingleLabelValue{}
	}
	if obj.multipleLabelValuesHolder == nil {
		obj.multipleLabelValuesHolder = newBgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter(&obj.obj.MultipleLabelValues).setMsg(obj)
	}
	return obj.multipleLabelValuesHolder
}

type bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter struct {
	obj                                       *bgpMplsLabelBindings
	bgpMplsLabelBindingsSingleLabelValueSlice []BgpMplsLabelBindingsSingleLabelValue
	fieldPtr                                  *[]*otg.BgpMplsLabelBindingsSingleLabelValue
}

func newBgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter(ptr *[]*otg.BgpMplsLabelBindingsSingleLabelValue) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter {
	return &bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter{fieldPtr: ptr}
}

type BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter interface {
	setMsg(*bgpMplsLabelBindings) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter
	Items() []BgpMplsLabelBindingsSingleLabelValue
	Add() BgpMplsLabelBindingsSingleLabelValue
	Append(items ...BgpMplsLabelBindingsSingleLabelValue) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter
	Set(index int, newObj BgpMplsLabelBindingsSingleLabelValue) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter
	Clear() BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter
	clearHolderSlice() BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter
	appendHolderSlice(item BgpMplsLabelBindingsSingleLabelValue) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter
}

func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter) setMsg(msg *bgpMplsLabelBindings) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpMplsLabelBindingsSingleLabelValue{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter) Items() []BgpMplsLabelBindingsSingleLabelValue {
	return obj.bgpMplsLabelBindingsSingleLabelValueSlice
}

func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter) Add() BgpMplsLabelBindingsSingleLabelValue {
	newObj := &otg.BgpMplsLabelBindingsSingleLabelValue{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpMplsLabelBindingsSingleLabelValue{obj: newObj}
	newLibObj.setDefault()
	obj.bgpMplsLabelBindingsSingleLabelValueSlice = append(obj.bgpMplsLabelBindingsSingleLabelValueSlice, newLibObj)
	return newLibObj
}

func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter) Append(items ...BgpMplsLabelBindingsSingleLabelValue) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpMplsLabelBindingsSingleLabelValueSlice = append(obj.bgpMplsLabelBindingsSingleLabelValueSlice, item)
	}
	return obj
}

func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter) Set(index int, newObj BgpMplsLabelBindingsSingleLabelValue) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpMplsLabelBindingsSingleLabelValueSlice[index] = newObj
	return obj
}
func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter) Clear() BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpMplsLabelBindingsSingleLabelValue{}
		obj.bgpMplsLabelBindingsSingleLabelValueSlice = []BgpMplsLabelBindingsSingleLabelValue{}
	}
	return obj
}
func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter) clearHolderSlice() BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter {
	if len(obj.bgpMplsLabelBindingsSingleLabelValueSlice) > 0 {
		obj.bgpMplsLabelBindingsSingleLabelValueSlice = []BgpMplsLabelBindingsSingleLabelValue{}
	}
	return obj
}
func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter) appendHolderSlice(item BgpMplsLabelBindingsSingleLabelValue) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelValueIter {
	obj.bgpMplsLabelBindingsSingleLabelValueSlice = append(obj.bgpMplsLabelBindingsSingleLabelValueSlice, item)
	return obj
}

// Group of MPLS SR Label Index bind to a prefix
// SingleSrLabelIndex returns a BgpMplsLabelBindingsSingleLabelIndex
func (obj *bgpMplsLabelBindings) SingleSrLabelIndex() BgpMplsLabelBindingsSingleLabelIndex {
	if obj.obj.SingleSrLabelIndex == nil {
		obj.setChoice(BgpMplsLabelBindingsChoice.SINGLE_SR_LABEL_INDEX)
	}
	if obj.singleSrLabelIndexHolder == nil {
		obj.singleSrLabelIndexHolder = &bgpMplsLabelBindingsSingleLabelIndex{obj: obj.obj.SingleSrLabelIndex}
	}
	return obj.singleSrLabelIndexHolder
}

// Group of MPLS SR Label Index bind to a prefix
// SingleSrLabelIndex returns a BgpMplsLabelBindingsSingleLabelIndex
func (obj *bgpMplsLabelBindings) HasSingleSrLabelIndex() bool {
	return obj.obj.SingleSrLabelIndex != nil
}

// Group of MPLS SR Label Index bind to a prefix
// SetSingleSrLabelIndex sets the BgpMplsLabelBindingsSingleLabelIndex value in the BgpMplsLabelBindings object
func (obj *bgpMplsLabelBindings) SetSingleSrLabelIndex(value BgpMplsLabelBindingsSingleLabelIndex) BgpMplsLabelBindings {
	obj.setChoice(BgpMplsLabelBindingsChoice.SINGLE_SR_LABEL_INDEX)
	obj.singleSrLabelIndexHolder = nil
	obj.obj.SingleSrLabelIndex = value.msg()

	return obj
}

// A list of group of MPLS SR Label indices bind to a prefix
// MultipleSrLabelIndices returns a []BgpMplsLabelBindingsSingleLabelIndex
func (obj *bgpMplsLabelBindings) MultipleSrLabelIndices() BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter {
	if len(obj.obj.MultipleSrLabelIndices) == 0 {
		obj.setChoice(BgpMplsLabelBindingsChoice.MULTIPLE_SR_LABEL_INDICES)
	}
	if obj.multipleSrLabelIndicesHolder == nil {
		obj.multipleSrLabelIndicesHolder = newBgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter(&obj.obj.MultipleSrLabelIndices).setMsg(obj)
	}
	return obj.multipleSrLabelIndicesHolder
}

type bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter struct {
	obj                                       *bgpMplsLabelBindings
	bgpMplsLabelBindingsSingleLabelIndexSlice []BgpMplsLabelBindingsSingleLabelIndex
	fieldPtr                                  *[]*otg.BgpMplsLabelBindingsSingleLabelIndex
}

func newBgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter(ptr *[]*otg.BgpMplsLabelBindingsSingleLabelIndex) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter {
	return &bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter{fieldPtr: ptr}
}

type BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter interface {
	setMsg(*bgpMplsLabelBindings) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter
	Items() []BgpMplsLabelBindingsSingleLabelIndex
	Add() BgpMplsLabelBindingsSingleLabelIndex
	Append(items ...BgpMplsLabelBindingsSingleLabelIndex) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter
	Set(index int, newObj BgpMplsLabelBindingsSingleLabelIndex) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter
	Clear() BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter
	clearHolderSlice() BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter
	appendHolderSlice(item BgpMplsLabelBindingsSingleLabelIndex) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter
}

func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter) setMsg(msg *bgpMplsLabelBindings) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpMplsLabelBindingsSingleLabelIndex{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter) Items() []BgpMplsLabelBindingsSingleLabelIndex {
	return obj.bgpMplsLabelBindingsSingleLabelIndexSlice
}

func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter) Add() BgpMplsLabelBindingsSingleLabelIndex {
	newObj := &otg.BgpMplsLabelBindingsSingleLabelIndex{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpMplsLabelBindingsSingleLabelIndex{obj: newObj}
	newLibObj.setDefault()
	obj.bgpMplsLabelBindingsSingleLabelIndexSlice = append(obj.bgpMplsLabelBindingsSingleLabelIndexSlice, newLibObj)
	return newLibObj
}

func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter) Append(items ...BgpMplsLabelBindingsSingleLabelIndex) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpMplsLabelBindingsSingleLabelIndexSlice = append(obj.bgpMplsLabelBindingsSingleLabelIndexSlice, item)
	}
	return obj
}

func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter) Set(index int, newObj BgpMplsLabelBindingsSingleLabelIndex) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpMplsLabelBindingsSingleLabelIndexSlice[index] = newObj
	return obj
}
func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter) Clear() BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpMplsLabelBindingsSingleLabelIndex{}
		obj.bgpMplsLabelBindingsSingleLabelIndexSlice = []BgpMplsLabelBindingsSingleLabelIndex{}
	}
	return obj
}
func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter) clearHolderSlice() BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter {
	if len(obj.bgpMplsLabelBindingsSingleLabelIndexSlice) > 0 {
		obj.bgpMplsLabelBindingsSingleLabelIndexSlice = []BgpMplsLabelBindingsSingleLabelIndex{}
	}
	return obj
}
func (obj *bgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter) appendHolderSlice(item BgpMplsLabelBindingsSingleLabelIndex) BgpMplsLabelBindingsBgpMplsLabelBindingsSingleLabelIndexIter {
	obj.bgpMplsLabelBindingsSingleLabelIndexSlice = append(obj.bgpMplsLabelBindingsSingleLabelIndexSlice, item)
	return obj
}

func (obj *bgpMplsLabelBindings) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SingleLabelValue != nil {

		obj.SingleLabelValue().validateObj(vObj, set_default)
	}

	if len(obj.obj.MultipleLabelValues) != 0 {

		if set_default {
			obj.MultipleLabelValues().clearHolderSlice()
			for _, item := range obj.obj.MultipleLabelValues {
				obj.MultipleLabelValues().appendHolderSlice(&bgpMplsLabelBindingsSingleLabelValue{obj: item})
			}
		}
		for _, item := range obj.MultipleLabelValues().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.SingleSrLabelIndex != nil {

		obj.SingleSrLabelIndex().validateObj(vObj, set_default)
	}

	if len(obj.obj.MultipleSrLabelIndices) != 0 {

		if set_default {
			obj.MultipleSrLabelIndices().clearHolderSlice()
			for _, item := range obj.obj.MultipleSrLabelIndices {
				obj.MultipleSrLabelIndices().appendHolderSlice(&bgpMplsLabelBindingsSingleLabelIndex{obj: item})
			}
		}
		for _, item := range obj.MultipleSrLabelIndices().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpMplsLabelBindings) setDefault() {
	var choices_set int = 0
	var choice BgpMplsLabelBindingsChoiceEnum

	if obj.obj.SingleLabelValue != nil {
		choices_set += 1
		choice = BgpMplsLabelBindingsChoice.SINGLE_LABEL_VALUE
	}

	if obj.obj.SingleSrLabelIndex != nil {
		choices_set += 1
		choice = BgpMplsLabelBindingsChoice.SINGLE_SR_LABEL_INDEX
	}

	if len(obj.obj.MultipleSrLabelIndices) > 0 {
		choices_set += 1
		choice = BgpMplsLabelBindingsChoice.MULTIPLE_SR_LABEL_INDICES
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpMplsLabelBindingsChoice.SINGLE_LABEL)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpMplsLabelBindings")
			}
		} else {
			intVal := otg.BgpMplsLabelBindings_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpMplsLabelBindings_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
