package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv1GroupAddressMetricTag *****
type patternFlowIgmpv1GroupAddressMetricTag struct {
	validation
	obj          *otg.PatternFlowIgmpv1GroupAddressMetricTag
	marshaller   marshalPatternFlowIgmpv1GroupAddressMetricTag
	unMarshaller unMarshalPatternFlowIgmpv1GroupAddressMetricTag
}

func NewPatternFlowIgmpv1GroupAddressMetricTag() PatternFlowIgmpv1GroupAddressMetricTag {
	obj := patternFlowIgmpv1GroupAddressMetricTag{obj: &otg.PatternFlowIgmpv1GroupAddressMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv1GroupAddressMetricTag) msg() *otg.PatternFlowIgmpv1GroupAddressMetricTag {
	return obj.obj
}

func (obj *patternFlowIgmpv1GroupAddressMetricTag) setMsg(msg *otg.PatternFlowIgmpv1GroupAddressMetricTag) PatternFlowIgmpv1GroupAddressMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv1GroupAddressMetricTag struct {
	obj *patternFlowIgmpv1GroupAddressMetricTag
}

type marshalPatternFlowIgmpv1GroupAddressMetricTag interface {
	// ToProto marshals PatternFlowIgmpv1GroupAddressMetricTag to protobuf object *otg.PatternFlowIgmpv1GroupAddressMetricTag
	ToProto() (*otg.PatternFlowIgmpv1GroupAddressMetricTag, error)
	// ToPbText marshals PatternFlowIgmpv1GroupAddressMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv1GroupAddressMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv1GroupAddressMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIgmpv1GroupAddressMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIgmpv1GroupAddressMetricTag struct {
	obj *patternFlowIgmpv1GroupAddressMetricTag
}

type unMarshalPatternFlowIgmpv1GroupAddressMetricTag interface {
	// FromProto unmarshals PatternFlowIgmpv1GroupAddressMetricTag from protobuf object *otg.PatternFlowIgmpv1GroupAddressMetricTag
	FromProto(msg *otg.PatternFlowIgmpv1GroupAddressMetricTag) (PatternFlowIgmpv1GroupAddressMetricTag, error)
	// FromPbText unmarshals PatternFlowIgmpv1GroupAddressMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv1GroupAddressMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv1GroupAddressMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv1GroupAddressMetricTag) Marshal() marshalPatternFlowIgmpv1GroupAddressMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv1GroupAddressMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv1GroupAddressMetricTag) Unmarshal() unMarshalPatternFlowIgmpv1GroupAddressMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv1GroupAddressMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv1GroupAddressMetricTag) ToProto() (*otg.PatternFlowIgmpv1GroupAddressMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv1GroupAddressMetricTag) FromProto(msg *otg.PatternFlowIgmpv1GroupAddressMetricTag) (PatternFlowIgmpv1GroupAddressMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv1GroupAddressMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1GroupAddressMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv1GroupAddressMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1GroupAddressMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv1GroupAddressMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIgmpv1GroupAddressMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1GroupAddressMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIgmpv1GroupAddressMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1GroupAddressMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1GroupAddressMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv1GroupAddressMetricTag) Clone() (PatternFlowIgmpv1GroupAddressMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv1GroupAddressMetricTag()
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

// PatternFlowIgmpv1GroupAddressMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIgmpv1GroupAddressMetricTag interface {
	Validation
	// msg marshals PatternFlowIgmpv1GroupAddressMetricTag to protobuf object *otg.PatternFlowIgmpv1GroupAddressMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv1GroupAddressMetricTag
	// setMsg unmarshals PatternFlowIgmpv1GroupAddressMetricTag from protobuf object *otg.PatternFlowIgmpv1GroupAddressMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv1GroupAddressMetricTag) PatternFlowIgmpv1GroupAddressMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv1GroupAddressMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv1GroupAddressMetricTag
	// validate validates PatternFlowIgmpv1GroupAddressMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv1GroupAddressMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIgmpv1GroupAddressMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIgmpv1GroupAddressMetricTag
	SetName(value string) PatternFlowIgmpv1GroupAddressMetricTag
	// Offset returns uint32, set in PatternFlowIgmpv1GroupAddressMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIgmpv1GroupAddressMetricTag
	SetOffset(value uint32) PatternFlowIgmpv1GroupAddressMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIgmpv1GroupAddressMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIgmpv1GroupAddressMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIgmpv1GroupAddressMetricTag
	SetLength(value uint32) PatternFlowIgmpv1GroupAddressMetricTag
	// HasLength checks if Length has been set in PatternFlowIgmpv1GroupAddressMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIgmpv1GroupAddressMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIgmpv1GroupAddressMetricTag object
func (obj *patternFlowIgmpv1GroupAddressMetricTag) SetName(value string) PatternFlowIgmpv1GroupAddressMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIgmpv1GroupAddressMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIgmpv1GroupAddressMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIgmpv1GroupAddressMetricTag object
func (obj *patternFlowIgmpv1GroupAddressMetricTag) SetOffset(value uint32) PatternFlowIgmpv1GroupAddressMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIgmpv1GroupAddressMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIgmpv1GroupAddressMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIgmpv1GroupAddressMetricTag object
func (obj *patternFlowIgmpv1GroupAddressMetricTag) SetLength(value uint32) PatternFlowIgmpv1GroupAddressMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIgmpv1GroupAddressMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIgmpv1GroupAddressMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1GroupAddressMetricTag.Offset <= 31 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIgmpv1GroupAddressMetricTag.Length <= 32 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIgmpv1GroupAddressMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(32)
	}

}
