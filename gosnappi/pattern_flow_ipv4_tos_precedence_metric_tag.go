package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosPrecedenceMetricTag *****
type patternFlowIpv4TosPrecedenceMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4TosPrecedenceMetricTag
	marshaller   marshalPatternFlowIpv4TosPrecedenceMetricTag
	unMarshaller unMarshalPatternFlowIpv4TosPrecedenceMetricTag
}

func NewPatternFlowIpv4TosPrecedenceMetricTag() PatternFlowIpv4TosPrecedenceMetricTag {
	obj := patternFlowIpv4TosPrecedenceMetricTag{obj: &otg.PatternFlowIpv4TosPrecedenceMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosPrecedenceMetricTag) msg() *otg.PatternFlowIpv4TosPrecedenceMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4TosPrecedenceMetricTag) setMsg(msg *otg.PatternFlowIpv4TosPrecedenceMetricTag) PatternFlowIpv4TosPrecedenceMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosPrecedenceMetricTag struct {
	obj *patternFlowIpv4TosPrecedenceMetricTag
}

type marshalPatternFlowIpv4TosPrecedenceMetricTag interface {
	// ToProto marshals PatternFlowIpv4TosPrecedenceMetricTag to protobuf object *otg.PatternFlowIpv4TosPrecedenceMetricTag
	ToProto() (*otg.PatternFlowIpv4TosPrecedenceMetricTag, error)
	// ToPbText marshals PatternFlowIpv4TosPrecedenceMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosPrecedenceMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosPrecedenceMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4TosPrecedenceMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4TosPrecedenceMetricTag struct {
	obj *patternFlowIpv4TosPrecedenceMetricTag
}

type unMarshalPatternFlowIpv4TosPrecedenceMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4TosPrecedenceMetricTag from protobuf object *otg.PatternFlowIpv4TosPrecedenceMetricTag
	FromProto(msg *otg.PatternFlowIpv4TosPrecedenceMetricTag) (PatternFlowIpv4TosPrecedenceMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4TosPrecedenceMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosPrecedenceMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosPrecedenceMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosPrecedenceMetricTag) Marshal() marshalPatternFlowIpv4TosPrecedenceMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosPrecedenceMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosPrecedenceMetricTag) Unmarshal() unMarshalPatternFlowIpv4TosPrecedenceMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosPrecedenceMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosPrecedenceMetricTag) ToProto() (*otg.PatternFlowIpv4TosPrecedenceMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosPrecedenceMetricTag) FromProto(msg *otg.PatternFlowIpv4TosPrecedenceMetricTag) (PatternFlowIpv4TosPrecedenceMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosPrecedenceMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosPrecedenceMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosPrecedenceMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosPrecedenceMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosPrecedenceMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4TosPrecedenceMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosPrecedenceMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosPrecedenceMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosPrecedenceMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosPrecedenceMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosPrecedenceMetricTag) Clone() (PatternFlowIpv4TosPrecedenceMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosPrecedenceMetricTag()
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

// PatternFlowIpv4TosPrecedenceMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4TosPrecedenceMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4TosPrecedenceMetricTag to protobuf object *otg.PatternFlowIpv4TosPrecedenceMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosPrecedenceMetricTag
	// setMsg unmarshals PatternFlowIpv4TosPrecedenceMetricTag from protobuf object *otg.PatternFlowIpv4TosPrecedenceMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosPrecedenceMetricTag) PatternFlowIpv4TosPrecedenceMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosPrecedenceMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosPrecedenceMetricTag
	// validate validates PatternFlowIpv4TosPrecedenceMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosPrecedenceMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4TosPrecedenceMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4TosPrecedenceMetricTag
	SetName(value string) PatternFlowIpv4TosPrecedenceMetricTag
	// Offset returns uint32, set in PatternFlowIpv4TosPrecedenceMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4TosPrecedenceMetricTag
	SetOffset(value uint32) PatternFlowIpv4TosPrecedenceMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4TosPrecedenceMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4TosPrecedenceMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4TosPrecedenceMetricTag
	SetLength(value uint32) PatternFlowIpv4TosPrecedenceMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4TosPrecedenceMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4TosPrecedenceMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4TosPrecedenceMetricTag object
func (obj *patternFlowIpv4TosPrecedenceMetricTag) SetName(value string) PatternFlowIpv4TosPrecedenceMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TosPrecedenceMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TosPrecedenceMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4TosPrecedenceMetricTag object
func (obj *patternFlowIpv4TosPrecedenceMetricTag) SetOffset(value uint32) PatternFlowIpv4TosPrecedenceMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TosPrecedenceMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TosPrecedenceMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4TosPrecedenceMetricTag object
func (obj *patternFlowIpv4TosPrecedenceMetricTag) SetLength(value uint32) PatternFlowIpv4TosPrecedenceMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4TosPrecedenceMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4TosPrecedenceMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosPrecedenceMetricTag.Offset <= 2 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4TosPrecedenceMetricTag.Length <= 3 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4TosPrecedenceMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(3)
	}

}
