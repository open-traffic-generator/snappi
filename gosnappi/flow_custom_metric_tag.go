package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCustomMetricTag *****
type flowCustomMetricTag struct {
	validation
	obj          *otg.FlowCustomMetricTag
	marshaller   marshalFlowCustomMetricTag
	unMarshaller unMarshalFlowCustomMetricTag
}

func NewFlowCustomMetricTag() FlowCustomMetricTag {
	obj := flowCustomMetricTag{obj: &otg.FlowCustomMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCustomMetricTag) msg() *otg.FlowCustomMetricTag {
	return obj.obj
}

func (obj *flowCustomMetricTag) setMsg(msg *otg.FlowCustomMetricTag) FlowCustomMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCustomMetricTag struct {
	obj *flowCustomMetricTag
}

type marshalFlowCustomMetricTag interface {
	// ToProto marshals FlowCustomMetricTag to protobuf object *otg.FlowCustomMetricTag
	ToProto() (*otg.FlowCustomMetricTag, error)
	// ToPbText marshals FlowCustomMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCustomMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCustomMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowCustomMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowCustomMetricTag struct {
	obj *flowCustomMetricTag
}

type unMarshalFlowCustomMetricTag interface {
	// FromProto unmarshals FlowCustomMetricTag from protobuf object *otg.FlowCustomMetricTag
	FromProto(msg *otg.FlowCustomMetricTag) (FlowCustomMetricTag, error)
	// FromPbText unmarshals FlowCustomMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCustomMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCustomMetricTag from JSON text
	FromJson(value string) error
}

func (obj *flowCustomMetricTag) Marshal() marshalFlowCustomMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCustomMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCustomMetricTag) Unmarshal() unMarshalFlowCustomMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCustomMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCustomMetricTag) ToProto() (*otg.FlowCustomMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCustomMetricTag) FromProto(msg *otg.FlowCustomMetricTag) (FlowCustomMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCustomMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalflowCustomMetricTag) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalflowCustomMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalflowCustomMetricTag) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalflowCustomMetricTag) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalflowCustomMetricTag) ToJson() (string, error) {
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

func (m *unMarshalflowCustomMetricTag) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *flowCustomMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCustomMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCustomMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCustomMetricTag) Clone() (FlowCustomMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCustomMetricTag()
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

// FlowCustomMetricTag is metric Tag can be used to enable tracking portion of or all bits
// in a corresponding header field for metrics per each applicable value.
// These would appear as tagged metrics in corresponding flow metrics.
type FlowCustomMetricTag interface {
	Validation
	// msg marshals FlowCustomMetricTag to protobuf object *otg.FlowCustomMetricTag
	// and doesn't set defaults
	msg() *otg.FlowCustomMetricTag
	// setMsg unmarshals FlowCustomMetricTag from protobuf object *otg.FlowCustomMetricTag
	// and doesn't set defaults
	setMsg(*otg.FlowCustomMetricTag) FlowCustomMetricTag
	// provides marshal interface
	Marshal() marshalFlowCustomMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCustomMetricTag
	// validate validates FlowCustomMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCustomMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in FlowCustomMetricTag.
	Name() string
	// SetName assigns string provided by user to FlowCustomMetricTag
	SetName(value string) FlowCustomMetricTag
	// Offset returns uint32, set in FlowCustomMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to FlowCustomMetricTag
	SetOffset(value uint32) FlowCustomMetricTag
	// HasOffset checks if Offset has been set in FlowCustomMetricTag
	HasOffset() bool
	// Length returns uint32, set in FlowCustomMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to FlowCustomMetricTag
	SetLength(value uint32) FlowCustomMetricTag
	// HasLength checks if Length has been set in FlowCustomMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable
// for configured offset and length inside corresponding header field
// Name returns a string
func (obj *flowCustomMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable
// for configured offset and length inside corresponding header field
// SetName sets the string value in the FlowCustomMetricTag object
func (obj *flowCustomMetricTag) SetName(value string) FlowCustomMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *flowCustomMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *flowCustomMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the FlowCustomMetricTag object
func (obj *flowCustomMetricTag) SetOffset(value uint32) FlowCustomMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset
// of corresponding header field
// Length returns a uint32
func (obj *flowCustomMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset
// of corresponding header field
// Length returns a uint32
func (obj *flowCustomMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset
// of corresponding header field
// SetLength sets the uint32 value in the FlowCustomMetricTag object
func (obj *flowCustomMetricTag) SetLength(value uint32) FlowCustomMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *flowCustomMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface FlowCustomMetricTag")
	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= FlowCustomMetricTag.Length <= 4294967295 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *flowCustomMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
