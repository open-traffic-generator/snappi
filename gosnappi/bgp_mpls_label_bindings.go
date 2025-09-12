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
	obj            *otg.BgpMplsLabelBindings
	marshaller     marshalBgpMplsLabelBindings
	unMarshaller   unMarshalBgpMplsLabelBindings
	labelsHolder   BgpMplsLabelBindingsRouteMplsLabelValueIter
	srLabelsHolder BgpMplsLabelBindingsRouteMplsLabelIndexIter
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
	obj.labelsHolder = nil
	obj.srLabelsHolder = nil
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
	// Labels returns BgpMplsLabelBindingsRouteMplsLabelValueIterIter, set in BgpMplsLabelBindings
	Labels() BgpMplsLabelBindingsRouteMplsLabelValueIter
	// SrLabels returns BgpMplsLabelBindingsRouteMplsLabelIndexIterIter, set in BgpMplsLabelBindings
	SrLabels() BgpMplsLabelBindingsRouteMplsLabelIndexIter
	setNil()
}

type BgpMplsLabelBindingsChoiceEnum string

// Enum of Choice on BgpMplsLabelBindings
var BgpMplsLabelBindingsChoice = struct {
	LABELS    BgpMplsLabelBindingsChoiceEnum
	SR_LABELS BgpMplsLabelBindingsChoiceEnum
}{
	LABELS:    BgpMplsLabelBindingsChoiceEnum("labels"),
	SR_LABELS: BgpMplsLabelBindingsChoiceEnum("sr_labels"),
}

func (obj *bgpMplsLabelBindings) Choice() BgpMplsLabelBindingsChoiceEnum {
	return BgpMplsLabelBindingsChoiceEnum(obj.obj.Choice.Enum().String())
}

// Bind one or more MPLS Labels or SR indices to Address Prefixes.
// To advertise multiple labels for an address prefix, the BGP capabilities device.bgp.capabliti.ipv4_mpls_multi/bgp.capabliti.ipv6_mpls_multi etc. to be configured first.
// - labels: One or more Label Values are to be configured here.
// - sr_labels: One or more SR indices are to be configured here. At least one SRGB range has been configured in device.bgp.mpls_srgb_ranges.
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
	obj.obj.SrLabels = nil
	obj.srLabelsHolder = nil
	obj.obj.Labels = nil
	obj.labelsHolder = nil

	if value == BgpMplsLabelBindingsChoice.LABELS {
		obj.obj.Labels = []*otg.RouteMplsLabelValue{}
	}

	if value == BgpMplsLabelBindingsChoice.SR_LABELS {
		obj.obj.SrLabels = []*otg.RouteMplsLabelIndex{}
	}

	return obj
}

// description is TBD
// Labels returns a []RouteMplsLabelValue
func (obj *bgpMplsLabelBindings) Labels() BgpMplsLabelBindingsRouteMplsLabelValueIter {
	if len(obj.obj.Labels) == 0 {
		obj.setChoice(BgpMplsLabelBindingsChoice.LABELS)
	}
	if obj.labelsHolder == nil {
		obj.labelsHolder = newBgpMplsLabelBindingsRouteMplsLabelValueIter(&obj.obj.Labels).setMsg(obj)
	}
	return obj.labelsHolder
}

type bgpMplsLabelBindingsRouteMplsLabelValueIter struct {
	obj                      *bgpMplsLabelBindings
	routeMplsLabelValueSlice []RouteMplsLabelValue
	fieldPtr                 *[]*otg.RouteMplsLabelValue
}

func newBgpMplsLabelBindingsRouteMplsLabelValueIter(ptr *[]*otg.RouteMplsLabelValue) BgpMplsLabelBindingsRouteMplsLabelValueIter {
	return &bgpMplsLabelBindingsRouteMplsLabelValueIter{fieldPtr: ptr}
}

type BgpMplsLabelBindingsRouteMplsLabelValueIter interface {
	setMsg(*bgpMplsLabelBindings) BgpMplsLabelBindingsRouteMplsLabelValueIter
	Items() []RouteMplsLabelValue
	Add() RouteMplsLabelValue
	Append(items ...RouteMplsLabelValue) BgpMplsLabelBindingsRouteMplsLabelValueIter
	Set(index int, newObj RouteMplsLabelValue) BgpMplsLabelBindingsRouteMplsLabelValueIter
	Clear() BgpMplsLabelBindingsRouteMplsLabelValueIter
	clearHolderSlice() BgpMplsLabelBindingsRouteMplsLabelValueIter
	appendHolderSlice(item RouteMplsLabelValue) BgpMplsLabelBindingsRouteMplsLabelValueIter
}

func (obj *bgpMplsLabelBindingsRouteMplsLabelValueIter) setMsg(msg *bgpMplsLabelBindings) BgpMplsLabelBindingsRouteMplsLabelValueIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&routeMplsLabelValue{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpMplsLabelBindingsRouteMplsLabelValueIter) Items() []RouteMplsLabelValue {
	return obj.routeMplsLabelValueSlice
}

func (obj *bgpMplsLabelBindingsRouteMplsLabelValueIter) Add() RouteMplsLabelValue {
	newObj := &otg.RouteMplsLabelValue{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &routeMplsLabelValue{obj: newObj}
	newLibObj.setDefault()
	obj.routeMplsLabelValueSlice = append(obj.routeMplsLabelValueSlice, newLibObj)
	return newLibObj
}

func (obj *bgpMplsLabelBindingsRouteMplsLabelValueIter) Append(items ...RouteMplsLabelValue) BgpMplsLabelBindingsRouteMplsLabelValueIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.routeMplsLabelValueSlice = append(obj.routeMplsLabelValueSlice, item)
	}
	return obj
}

func (obj *bgpMplsLabelBindingsRouteMplsLabelValueIter) Set(index int, newObj RouteMplsLabelValue) BgpMplsLabelBindingsRouteMplsLabelValueIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.routeMplsLabelValueSlice[index] = newObj
	return obj
}
func (obj *bgpMplsLabelBindingsRouteMplsLabelValueIter) Clear() BgpMplsLabelBindingsRouteMplsLabelValueIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RouteMplsLabelValue{}
		obj.routeMplsLabelValueSlice = []RouteMplsLabelValue{}
	}
	return obj
}
func (obj *bgpMplsLabelBindingsRouteMplsLabelValueIter) clearHolderSlice() BgpMplsLabelBindingsRouteMplsLabelValueIter {
	if len(obj.routeMplsLabelValueSlice) > 0 {
		obj.routeMplsLabelValueSlice = []RouteMplsLabelValue{}
	}
	return obj
}
func (obj *bgpMplsLabelBindingsRouteMplsLabelValueIter) appendHolderSlice(item RouteMplsLabelValue) BgpMplsLabelBindingsRouteMplsLabelValueIter {
	obj.routeMplsLabelValueSlice = append(obj.routeMplsLabelValueSlice, item)
	return obj
}

// description is TBD
// SrLabels returns a []RouteMplsLabelIndex
func (obj *bgpMplsLabelBindings) SrLabels() BgpMplsLabelBindingsRouteMplsLabelIndexIter {
	if len(obj.obj.SrLabels) == 0 {
		obj.setChoice(BgpMplsLabelBindingsChoice.SR_LABELS)
	}
	if obj.srLabelsHolder == nil {
		obj.srLabelsHolder = newBgpMplsLabelBindingsRouteMplsLabelIndexIter(&obj.obj.SrLabels).setMsg(obj)
	}
	return obj.srLabelsHolder
}

type bgpMplsLabelBindingsRouteMplsLabelIndexIter struct {
	obj                      *bgpMplsLabelBindings
	routeMplsLabelIndexSlice []RouteMplsLabelIndex
	fieldPtr                 *[]*otg.RouteMplsLabelIndex
}

func newBgpMplsLabelBindingsRouteMplsLabelIndexIter(ptr *[]*otg.RouteMplsLabelIndex) BgpMplsLabelBindingsRouteMplsLabelIndexIter {
	return &bgpMplsLabelBindingsRouteMplsLabelIndexIter{fieldPtr: ptr}
}

type BgpMplsLabelBindingsRouteMplsLabelIndexIter interface {
	setMsg(*bgpMplsLabelBindings) BgpMplsLabelBindingsRouteMplsLabelIndexIter
	Items() []RouteMplsLabelIndex
	Add() RouteMplsLabelIndex
	Append(items ...RouteMplsLabelIndex) BgpMplsLabelBindingsRouteMplsLabelIndexIter
	Set(index int, newObj RouteMplsLabelIndex) BgpMplsLabelBindingsRouteMplsLabelIndexIter
	Clear() BgpMplsLabelBindingsRouteMplsLabelIndexIter
	clearHolderSlice() BgpMplsLabelBindingsRouteMplsLabelIndexIter
	appendHolderSlice(item RouteMplsLabelIndex) BgpMplsLabelBindingsRouteMplsLabelIndexIter
}

func (obj *bgpMplsLabelBindingsRouteMplsLabelIndexIter) setMsg(msg *bgpMplsLabelBindings) BgpMplsLabelBindingsRouteMplsLabelIndexIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&routeMplsLabelIndex{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpMplsLabelBindingsRouteMplsLabelIndexIter) Items() []RouteMplsLabelIndex {
	return obj.routeMplsLabelIndexSlice
}

func (obj *bgpMplsLabelBindingsRouteMplsLabelIndexIter) Add() RouteMplsLabelIndex {
	newObj := &otg.RouteMplsLabelIndex{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &routeMplsLabelIndex{obj: newObj}
	newLibObj.setDefault()
	obj.routeMplsLabelIndexSlice = append(obj.routeMplsLabelIndexSlice, newLibObj)
	return newLibObj
}

func (obj *bgpMplsLabelBindingsRouteMplsLabelIndexIter) Append(items ...RouteMplsLabelIndex) BgpMplsLabelBindingsRouteMplsLabelIndexIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.routeMplsLabelIndexSlice = append(obj.routeMplsLabelIndexSlice, item)
	}
	return obj
}

func (obj *bgpMplsLabelBindingsRouteMplsLabelIndexIter) Set(index int, newObj RouteMplsLabelIndex) BgpMplsLabelBindingsRouteMplsLabelIndexIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.routeMplsLabelIndexSlice[index] = newObj
	return obj
}
func (obj *bgpMplsLabelBindingsRouteMplsLabelIndexIter) Clear() BgpMplsLabelBindingsRouteMplsLabelIndexIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RouteMplsLabelIndex{}
		obj.routeMplsLabelIndexSlice = []RouteMplsLabelIndex{}
	}
	return obj
}
func (obj *bgpMplsLabelBindingsRouteMplsLabelIndexIter) clearHolderSlice() BgpMplsLabelBindingsRouteMplsLabelIndexIter {
	if len(obj.routeMplsLabelIndexSlice) > 0 {
		obj.routeMplsLabelIndexSlice = []RouteMplsLabelIndex{}
	}
	return obj
}
func (obj *bgpMplsLabelBindingsRouteMplsLabelIndexIter) appendHolderSlice(item RouteMplsLabelIndex) BgpMplsLabelBindingsRouteMplsLabelIndexIter {
	obj.routeMplsLabelIndexSlice = append(obj.routeMplsLabelIndexSlice, item)
	return obj
}

func (obj *bgpMplsLabelBindings) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Labels) != 0 {

		if set_default {
			obj.Labels().clearHolderSlice()
			for _, item := range obj.obj.Labels {
				obj.Labels().appendHolderSlice(&routeMplsLabelValue{obj: item})
			}
		}
		for _, item := range obj.Labels().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.SrLabels) != 0 {

		if set_default {
			obj.SrLabels().clearHolderSlice()
			for _, item := range obj.obj.SrLabels {
				obj.SrLabels().appendHolderSlice(&routeMplsLabelIndex{obj: item})
			}
		}
		for _, item := range obj.SrLabels().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpMplsLabelBindings) setDefault() {
	var choices_set int = 0
	var choice BgpMplsLabelBindingsChoiceEnum

	if len(obj.obj.Labels) > 0 {
		choices_set += 1
		choice = BgpMplsLabelBindingsChoice.LABELS
	}

	if len(obj.obj.SrLabels) > 0 {
		choices_set += 1
		choice = BgpMplsLabelBindingsChoice.SR_LABELS
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpMplsLabelBindingsChoice.LABELS)

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
