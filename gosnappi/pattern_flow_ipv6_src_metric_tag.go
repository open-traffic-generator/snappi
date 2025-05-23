package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SrcMetricTag *****
type patternFlowIpv6SrcMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv6SrcMetricTag
	marshaller   marshalPatternFlowIpv6SrcMetricTag
	unMarshaller unMarshalPatternFlowIpv6SrcMetricTag
}

func NewPatternFlowIpv6SrcMetricTag() PatternFlowIpv6SrcMetricTag {
	obj := patternFlowIpv6SrcMetricTag{obj: &otg.PatternFlowIpv6SrcMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SrcMetricTag) msg() *otg.PatternFlowIpv6SrcMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv6SrcMetricTag) setMsg(msg *otg.PatternFlowIpv6SrcMetricTag) PatternFlowIpv6SrcMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SrcMetricTag struct {
	obj *patternFlowIpv6SrcMetricTag
}

type marshalPatternFlowIpv6SrcMetricTag interface {
	// ToProto marshals PatternFlowIpv6SrcMetricTag to protobuf object *otg.PatternFlowIpv6SrcMetricTag
	ToProto() (*otg.PatternFlowIpv6SrcMetricTag, error)
	// ToPbText marshals PatternFlowIpv6SrcMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SrcMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SrcMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv6SrcMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv6SrcMetricTag struct {
	obj *patternFlowIpv6SrcMetricTag
}

type unMarshalPatternFlowIpv6SrcMetricTag interface {
	// FromProto unmarshals PatternFlowIpv6SrcMetricTag from protobuf object *otg.PatternFlowIpv6SrcMetricTag
	FromProto(msg *otg.PatternFlowIpv6SrcMetricTag) (PatternFlowIpv6SrcMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv6SrcMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SrcMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SrcMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SrcMetricTag) Marshal() marshalPatternFlowIpv6SrcMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SrcMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SrcMetricTag) Unmarshal() unMarshalPatternFlowIpv6SrcMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SrcMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SrcMetricTag) ToProto() (*otg.PatternFlowIpv6SrcMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SrcMetricTag) FromProto(msg *otg.PatternFlowIpv6SrcMetricTag) (PatternFlowIpv6SrcMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SrcMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SrcMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SrcMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SrcMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SrcMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv6SrcMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SrcMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv6SrcMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SrcMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SrcMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SrcMetricTag) Clone() (PatternFlowIpv6SrcMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SrcMetricTag()
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

// PatternFlowIpv6SrcMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv6SrcMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv6SrcMetricTag to protobuf object *otg.PatternFlowIpv6SrcMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SrcMetricTag
	// setMsg unmarshals PatternFlowIpv6SrcMetricTag from protobuf object *otg.PatternFlowIpv6SrcMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SrcMetricTag) PatternFlowIpv6SrcMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SrcMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SrcMetricTag
	// validate validates PatternFlowIpv6SrcMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SrcMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv6SrcMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv6SrcMetricTag
	SetName(value string) PatternFlowIpv6SrcMetricTag
	// Offset returns uint32, set in PatternFlowIpv6SrcMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv6SrcMetricTag
	SetOffset(value uint32) PatternFlowIpv6SrcMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv6SrcMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv6SrcMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv6SrcMetricTag
	SetLength(value uint32) PatternFlowIpv6SrcMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv6SrcMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv6SrcMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv6SrcMetricTag object
func (obj *patternFlowIpv6SrcMetricTag) SetName(value string) PatternFlowIpv6SrcMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6SrcMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6SrcMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv6SrcMetricTag object
func (obj *patternFlowIpv6SrcMetricTag) SetOffset(value uint32) PatternFlowIpv6SrcMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6SrcMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6SrcMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv6SrcMetricTag object
func (obj *patternFlowIpv6SrcMetricTag) SetLength(value uint32) PatternFlowIpv6SrcMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv6SrcMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv6SrcMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 127 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SrcMetricTag.Offset <= 127 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv6SrcMetricTag.Length <= 128 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv6SrcMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(128)
	}

}
