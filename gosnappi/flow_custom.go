package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCustom *****
type flowCustom struct {
	validation
	obj              *otg.FlowCustom
	marshaller       marshalFlowCustom
	unMarshaller     unMarshalFlowCustom
	metricTagsHolder FlowCustomFlowCustomMetricTagIter
}

func NewFlowCustom() FlowCustom {
	obj := flowCustom{obj: &otg.FlowCustom{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCustom) msg() *otg.FlowCustom {
	return obj.obj
}

func (obj *flowCustom) setMsg(msg *otg.FlowCustom) FlowCustom {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCustom struct {
	obj *flowCustom
}

type marshalFlowCustom interface {
	// ToProto marshals FlowCustom to protobuf object *otg.FlowCustom
	ToProto() (*otg.FlowCustom, error)
	// ToPbText marshals FlowCustom to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCustom to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCustom to JSON text
	ToJson() (string, error)
}

type unMarshalflowCustom struct {
	obj *flowCustom
}

type unMarshalFlowCustom interface {
	// FromProto unmarshals FlowCustom from protobuf object *otg.FlowCustom
	FromProto(msg *otg.FlowCustom) (FlowCustom, error)
	// FromPbText unmarshals FlowCustom from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCustom from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCustom from JSON text
	FromJson(value string) error
}

func (obj *flowCustom) Marshal() marshalFlowCustom {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCustom{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCustom) Unmarshal() unMarshalFlowCustom {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCustom{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCustom) ToProto() (*otg.FlowCustom, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCustom) FromProto(msg *otg.FlowCustom) (FlowCustom, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCustom) ToPbText() (string, error) {
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

func (m *unMarshalflowCustom) FromPbText(value string) error {
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

func (m *marshalflowCustom) ToYaml() (string, error) {
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

func (m *unMarshalflowCustom) FromYaml(value string) error {
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

func (m *marshalflowCustom) ToJson() (string, error) {
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

func (m *unMarshalflowCustom) FromJson(value string) error {
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

func (obj *flowCustom) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCustom) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCustom) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCustom) Clone() (FlowCustom, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCustom()
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

func (obj *flowCustom) setNil() {
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCustom is custom packet header
type FlowCustom interface {
	Validation
	// msg marshals FlowCustom to protobuf object *otg.FlowCustom
	// and doesn't set defaults
	msg() *otg.FlowCustom
	// setMsg unmarshals FlowCustom from protobuf object *otg.FlowCustom
	// and doesn't set defaults
	setMsg(*otg.FlowCustom) FlowCustom
	// provides marshal interface
	Marshal() marshalFlowCustom
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCustom
	// validate validates FlowCustom
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCustom, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Bytes returns string, set in FlowCustom.
	Bytes() string
	// SetBytes assigns string provided by user to FlowCustom
	SetBytes(value string) FlowCustom
	// MetricTags returns FlowCustomFlowCustomMetricTagIterIter, set in FlowCustom
	MetricTags() FlowCustomFlowCustomMetricTagIter
	setNil()
}

// A custom packet header defined as a string of hex bytes. The string MUST contain sequence of valid hex bytes. Spaces or colons can be part of the bytes but will be discarded. This packet header can be used in multiple places in the packet.
// Bytes returns a string
func (obj *flowCustom) Bytes() string {

	return *obj.obj.Bytes

}

// A custom packet header defined as a string of hex bytes. The string MUST contain sequence of valid hex bytes. Spaces or colons can be part of the bytes but will be discarded. This packet header can be used in multiple places in the packet.
// SetBytes sets the string value in the FlowCustom object
func (obj *flowCustom) SetBytes(value string) FlowCustom {

	obj.obj.Bytes = &value
	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits
// in a corresponding header field for metrics per each applicable value.
// These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []FlowCustomMetricTag
func (obj *flowCustom) MetricTags() FlowCustomFlowCustomMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.FlowCustomMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newFlowCustomFlowCustomMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type flowCustomFlowCustomMetricTagIter struct {
	obj                      *flowCustom
	flowCustomMetricTagSlice []FlowCustomMetricTag
	fieldPtr                 *[]*otg.FlowCustomMetricTag
}

func newFlowCustomFlowCustomMetricTagIter(ptr *[]*otg.FlowCustomMetricTag) FlowCustomFlowCustomMetricTagIter {
	return &flowCustomFlowCustomMetricTagIter{fieldPtr: ptr}
}

type FlowCustomFlowCustomMetricTagIter interface {
	setMsg(*flowCustom) FlowCustomFlowCustomMetricTagIter
	Items() []FlowCustomMetricTag
	Add() FlowCustomMetricTag
	Append(items ...FlowCustomMetricTag) FlowCustomFlowCustomMetricTagIter
	Set(index int, newObj FlowCustomMetricTag) FlowCustomFlowCustomMetricTagIter
	Clear() FlowCustomFlowCustomMetricTagIter
	clearHolderSlice() FlowCustomFlowCustomMetricTagIter
	appendHolderSlice(item FlowCustomMetricTag) FlowCustomFlowCustomMetricTagIter
}

func (obj *flowCustomFlowCustomMetricTagIter) setMsg(msg *flowCustom) FlowCustomFlowCustomMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowCustomMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowCustomFlowCustomMetricTagIter) Items() []FlowCustomMetricTag {
	return obj.flowCustomMetricTagSlice
}

func (obj *flowCustomFlowCustomMetricTagIter) Add() FlowCustomMetricTag {
	newObj := &otg.FlowCustomMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowCustomMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.flowCustomMetricTagSlice = append(obj.flowCustomMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *flowCustomFlowCustomMetricTagIter) Append(items ...FlowCustomMetricTag) FlowCustomFlowCustomMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowCustomMetricTagSlice = append(obj.flowCustomMetricTagSlice, item)
	}
	return obj
}

func (obj *flowCustomFlowCustomMetricTagIter) Set(index int, newObj FlowCustomMetricTag) FlowCustomFlowCustomMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowCustomMetricTagSlice[index] = newObj
	return obj
}
func (obj *flowCustomFlowCustomMetricTagIter) Clear() FlowCustomFlowCustomMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowCustomMetricTag{}
		obj.flowCustomMetricTagSlice = []FlowCustomMetricTag{}
	}
	return obj
}
func (obj *flowCustomFlowCustomMetricTagIter) clearHolderSlice() FlowCustomFlowCustomMetricTagIter {
	if len(obj.flowCustomMetricTagSlice) > 0 {
		obj.flowCustomMetricTagSlice = []FlowCustomMetricTag{}
	}
	return obj
}
func (obj *flowCustomFlowCustomMetricTagIter) appendHolderSlice(item FlowCustomMetricTag) FlowCustomFlowCustomMetricTagIter {
	obj.flowCustomMetricTagSlice = append(obj.flowCustomMetricTagSlice, item)
	return obj
}

func (obj *flowCustom) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Bytes is required
	if obj.obj.Bytes == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Bytes is required field on interface FlowCustom")
	}
	if obj.obj.Bytes != nil {

		if !regexp.MustCompile(`^[A-Fa-f0-9: ]+$`).MatchString(*obj.obj.Bytes) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"FlowCustom.Bytes should adhere to this regex pattern '%s', but Got %s", `^[A-Fa-f0-9: ]+$`, *obj.obj.Bytes))
		}

	}

	if len(obj.obj.MetricTags) != 0 {

		if set_default {
			obj.MetricTags().clearHolderSlice()
			for _, item := range obj.obj.MetricTags {
				obj.MetricTags().appendHolderSlice(&flowCustomMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowCustom) setDefault() {

}
