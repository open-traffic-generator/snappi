package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVlanTpidMetricTag *****
type patternFlowVlanTpidMetricTag struct {
	validation
	obj          *otg.PatternFlowVlanTpidMetricTag
	marshaller   marshalPatternFlowVlanTpidMetricTag
	unMarshaller unMarshalPatternFlowVlanTpidMetricTag
}

func NewPatternFlowVlanTpidMetricTag() PatternFlowVlanTpidMetricTag {
	obj := patternFlowVlanTpidMetricTag{obj: &otg.PatternFlowVlanTpidMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVlanTpidMetricTag) msg() *otg.PatternFlowVlanTpidMetricTag {
	return obj.obj
}

func (obj *patternFlowVlanTpidMetricTag) setMsg(msg *otg.PatternFlowVlanTpidMetricTag) PatternFlowVlanTpidMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVlanTpidMetricTag struct {
	obj *patternFlowVlanTpidMetricTag
}

type marshalPatternFlowVlanTpidMetricTag interface {
	// ToProto marshals PatternFlowVlanTpidMetricTag to protobuf object *otg.PatternFlowVlanTpidMetricTag
	ToProto() (*otg.PatternFlowVlanTpidMetricTag, error)
	// ToPbText marshals PatternFlowVlanTpidMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVlanTpidMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVlanTpidMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowVlanTpidMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowVlanTpidMetricTag struct {
	obj *patternFlowVlanTpidMetricTag
}

type unMarshalPatternFlowVlanTpidMetricTag interface {
	// FromProto unmarshals PatternFlowVlanTpidMetricTag from protobuf object *otg.PatternFlowVlanTpidMetricTag
	FromProto(msg *otg.PatternFlowVlanTpidMetricTag) (PatternFlowVlanTpidMetricTag, error)
	// FromPbText unmarshals PatternFlowVlanTpidMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVlanTpidMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVlanTpidMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVlanTpidMetricTag) Marshal() marshalPatternFlowVlanTpidMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVlanTpidMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVlanTpidMetricTag) Unmarshal() unMarshalPatternFlowVlanTpidMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVlanTpidMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVlanTpidMetricTag) ToProto() (*otg.PatternFlowVlanTpidMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVlanTpidMetricTag) FromProto(msg *otg.PatternFlowVlanTpidMetricTag) (PatternFlowVlanTpidMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVlanTpidMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVlanTpidMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowVlanTpidMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVlanTpidMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowVlanTpidMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowVlanTpidMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVlanTpidMetricTag) FromJson(value string) error {
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

func (obj *patternFlowVlanTpidMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVlanTpidMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVlanTpidMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVlanTpidMetricTag) Clone() (PatternFlowVlanTpidMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVlanTpidMetricTag()
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

// PatternFlowVlanTpidMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowVlanTpidMetricTag interface {
	Validation
	// msg marshals PatternFlowVlanTpidMetricTag to protobuf object *otg.PatternFlowVlanTpidMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowVlanTpidMetricTag
	// setMsg unmarshals PatternFlowVlanTpidMetricTag from protobuf object *otg.PatternFlowVlanTpidMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVlanTpidMetricTag) PatternFlowVlanTpidMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowVlanTpidMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVlanTpidMetricTag
	// validate validates PatternFlowVlanTpidMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVlanTpidMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowVlanTpidMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowVlanTpidMetricTag
	SetName(value string) PatternFlowVlanTpidMetricTag
	// Offset returns uint32, set in PatternFlowVlanTpidMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowVlanTpidMetricTag
	SetOffset(value uint32) PatternFlowVlanTpidMetricTag
	// HasOffset checks if Offset has been set in PatternFlowVlanTpidMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowVlanTpidMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowVlanTpidMetricTag
	SetLength(value uint32) PatternFlowVlanTpidMetricTag
	// HasLength checks if Length has been set in PatternFlowVlanTpidMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowVlanTpidMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowVlanTpidMetricTag object
func (obj *patternFlowVlanTpidMetricTag) SetName(value string) PatternFlowVlanTpidMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVlanTpidMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVlanTpidMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowVlanTpidMetricTag object
func (obj *patternFlowVlanTpidMetricTag) SetOffset(value uint32) PatternFlowVlanTpidMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVlanTpidMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVlanTpidMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowVlanTpidMetricTag object
func (obj *patternFlowVlanTpidMetricTag) SetLength(value uint32) PatternFlowVlanTpidMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowVlanTpidMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowVlanTpidMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanTpidMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowVlanTpidMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowVlanTpidMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
