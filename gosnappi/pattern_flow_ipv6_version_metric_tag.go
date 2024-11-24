package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6VersionMetricTag *****
type patternFlowIpv6VersionMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv6VersionMetricTag
	marshaller   marshalPatternFlowIpv6VersionMetricTag
	unMarshaller unMarshalPatternFlowIpv6VersionMetricTag
}

func NewPatternFlowIpv6VersionMetricTag() PatternFlowIpv6VersionMetricTag {
	obj := patternFlowIpv6VersionMetricTag{obj: &otg.PatternFlowIpv6VersionMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6VersionMetricTag) msg() *otg.PatternFlowIpv6VersionMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv6VersionMetricTag) setMsg(msg *otg.PatternFlowIpv6VersionMetricTag) PatternFlowIpv6VersionMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6VersionMetricTag struct {
	obj *patternFlowIpv6VersionMetricTag
}

type marshalPatternFlowIpv6VersionMetricTag interface {
	// ToProto marshals PatternFlowIpv6VersionMetricTag to protobuf object *otg.PatternFlowIpv6VersionMetricTag
	ToProto() (*otg.PatternFlowIpv6VersionMetricTag, error)
	// ToPbText marshals PatternFlowIpv6VersionMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6VersionMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6VersionMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv6VersionMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv6VersionMetricTag struct {
	obj *patternFlowIpv6VersionMetricTag
}

type unMarshalPatternFlowIpv6VersionMetricTag interface {
	// FromProto unmarshals PatternFlowIpv6VersionMetricTag from protobuf object *otg.PatternFlowIpv6VersionMetricTag
	FromProto(msg *otg.PatternFlowIpv6VersionMetricTag) (PatternFlowIpv6VersionMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv6VersionMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6VersionMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6VersionMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6VersionMetricTag) Marshal() marshalPatternFlowIpv6VersionMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6VersionMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6VersionMetricTag) Unmarshal() unMarshalPatternFlowIpv6VersionMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6VersionMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6VersionMetricTag) ToProto() (*otg.PatternFlowIpv6VersionMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6VersionMetricTag) FromProto(msg *otg.PatternFlowIpv6VersionMetricTag) (PatternFlowIpv6VersionMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6VersionMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6VersionMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6VersionMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6VersionMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6VersionMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv6VersionMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6VersionMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv6VersionMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6VersionMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6VersionMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6VersionMetricTag) Clone() (PatternFlowIpv6VersionMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6VersionMetricTag()
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

// PatternFlowIpv6VersionMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv6VersionMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv6VersionMetricTag to protobuf object *otg.PatternFlowIpv6VersionMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6VersionMetricTag
	// setMsg unmarshals PatternFlowIpv6VersionMetricTag from protobuf object *otg.PatternFlowIpv6VersionMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6VersionMetricTag) PatternFlowIpv6VersionMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6VersionMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6VersionMetricTag
	// validate validates PatternFlowIpv6VersionMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6VersionMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv6VersionMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv6VersionMetricTag
	SetName(value string) PatternFlowIpv6VersionMetricTag
	// Offset returns uint32, set in PatternFlowIpv6VersionMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv6VersionMetricTag
	SetOffset(value uint32) PatternFlowIpv6VersionMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv6VersionMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv6VersionMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv6VersionMetricTag
	SetLength(value uint32) PatternFlowIpv6VersionMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv6VersionMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv6VersionMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv6VersionMetricTag object
func (obj *patternFlowIpv6VersionMetricTag) SetName(value string) PatternFlowIpv6VersionMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6VersionMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6VersionMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv6VersionMetricTag object
func (obj *patternFlowIpv6VersionMetricTag) SetOffset(value uint32) PatternFlowIpv6VersionMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6VersionMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6VersionMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv6VersionMetricTag object
func (obj *patternFlowIpv6VersionMetricTag) SetLength(value uint32) PatternFlowIpv6VersionMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv6VersionMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv6VersionMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6VersionMetricTag.Offset <= 3 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 4 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv6VersionMetricTag.Length <= 4 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv6VersionMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(4)
	}

}
