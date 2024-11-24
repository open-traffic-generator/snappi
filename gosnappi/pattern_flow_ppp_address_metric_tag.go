package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPppAddressMetricTag *****
type patternFlowPppAddressMetricTag struct {
	validation
	obj          *otg.PatternFlowPppAddressMetricTag
	marshaller   marshalPatternFlowPppAddressMetricTag
	unMarshaller unMarshalPatternFlowPppAddressMetricTag
}

func NewPatternFlowPppAddressMetricTag() PatternFlowPppAddressMetricTag {
	obj := patternFlowPppAddressMetricTag{obj: &otg.PatternFlowPppAddressMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPppAddressMetricTag) msg() *otg.PatternFlowPppAddressMetricTag {
	return obj.obj
}

func (obj *patternFlowPppAddressMetricTag) setMsg(msg *otg.PatternFlowPppAddressMetricTag) PatternFlowPppAddressMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPppAddressMetricTag struct {
	obj *patternFlowPppAddressMetricTag
}

type marshalPatternFlowPppAddressMetricTag interface {
	// ToProto marshals PatternFlowPppAddressMetricTag to protobuf object *otg.PatternFlowPppAddressMetricTag
	ToProto() (*otg.PatternFlowPppAddressMetricTag, error)
	// ToPbText marshals PatternFlowPppAddressMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPppAddressMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPppAddressMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowPppAddressMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowPppAddressMetricTag struct {
	obj *patternFlowPppAddressMetricTag
}

type unMarshalPatternFlowPppAddressMetricTag interface {
	// FromProto unmarshals PatternFlowPppAddressMetricTag from protobuf object *otg.PatternFlowPppAddressMetricTag
	FromProto(msg *otg.PatternFlowPppAddressMetricTag) (PatternFlowPppAddressMetricTag, error)
	// FromPbText unmarshals PatternFlowPppAddressMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPppAddressMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPppAddressMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPppAddressMetricTag) Marshal() marshalPatternFlowPppAddressMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPppAddressMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPppAddressMetricTag) Unmarshal() unMarshalPatternFlowPppAddressMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPppAddressMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPppAddressMetricTag) ToProto() (*otg.PatternFlowPppAddressMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPppAddressMetricTag) FromProto(msg *otg.PatternFlowPppAddressMetricTag) (PatternFlowPppAddressMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPppAddressMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPppAddressMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowPppAddressMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPppAddressMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowPppAddressMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowPppAddressMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPppAddressMetricTag) FromJson(value string) error {
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

func (obj *patternFlowPppAddressMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPppAddressMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPppAddressMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPppAddressMetricTag) Clone() (PatternFlowPppAddressMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPppAddressMetricTag()
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

// PatternFlowPppAddressMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowPppAddressMetricTag interface {
	Validation
	// msg marshals PatternFlowPppAddressMetricTag to protobuf object *otg.PatternFlowPppAddressMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowPppAddressMetricTag
	// setMsg unmarshals PatternFlowPppAddressMetricTag from protobuf object *otg.PatternFlowPppAddressMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPppAddressMetricTag) PatternFlowPppAddressMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowPppAddressMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPppAddressMetricTag
	// validate validates PatternFlowPppAddressMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPppAddressMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowPppAddressMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowPppAddressMetricTag
	SetName(value string) PatternFlowPppAddressMetricTag
	// Offset returns uint32, set in PatternFlowPppAddressMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowPppAddressMetricTag
	SetOffset(value uint32) PatternFlowPppAddressMetricTag
	// HasOffset checks if Offset has been set in PatternFlowPppAddressMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowPppAddressMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowPppAddressMetricTag
	SetLength(value uint32) PatternFlowPppAddressMetricTag
	// HasLength checks if Length has been set in PatternFlowPppAddressMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowPppAddressMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowPppAddressMetricTag object
func (obj *patternFlowPppAddressMetricTag) SetName(value string) PatternFlowPppAddressMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPppAddressMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPppAddressMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowPppAddressMetricTag object
func (obj *patternFlowPppAddressMetricTag) SetOffset(value uint32) PatternFlowPppAddressMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPppAddressMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPppAddressMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowPppAddressMetricTag object
func (obj *patternFlowPppAddressMetricTag) SetLength(value uint32) PatternFlowPppAddressMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowPppAddressMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowPppAddressMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppAddressMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowPppAddressMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowPppAddressMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
