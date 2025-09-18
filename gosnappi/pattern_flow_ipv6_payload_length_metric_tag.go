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

// ***** PatternFlowIpv6PayloadLengthMetricTag *****
type patternFlowIpv6PayloadLengthMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv6PayloadLengthMetricTag
	marshaller   marshalPatternFlowIpv6PayloadLengthMetricTag
	unMarshaller unMarshalPatternFlowIpv6PayloadLengthMetricTag
}

func NewPatternFlowIpv6PayloadLengthMetricTag() PatternFlowIpv6PayloadLengthMetricTag {
	obj := patternFlowIpv6PayloadLengthMetricTag{obj: &otg.PatternFlowIpv6PayloadLengthMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6PayloadLengthMetricTag) msg() *otg.PatternFlowIpv6PayloadLengthMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv6PayloadLengthMetricTag) setMsg(msg *otg.PatternFlowIpv6PayloadLengthMetricTag) PatternFlowIpv6PayloadLengthMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6PayloadLengthMetricTag struct {
	obj *patternFlowIpv6PayloadLengthMetricTag
}

type marshalPatternFlowIpv6PayloadLengthMetricTag interface {
	// ToProto marshals PatternFlowIpv6PayloadLengthMetricTag to protobuf object *otg.PatternFlowIpv6PayloadLengthMetricTag
	ToProto() (*otg.PatternFlowIpv6PayloadLengthMetricTag, error)
	// ToPbText marshals PatternFlowIpv6PayloadLengthMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6PayloadLengthMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6PayloadLengthMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6PayloadLengthMetricTag struct {
	obj *patternFlowIpv6PayloadLengthMetricTag
}

type unMarshalPatternFlowIpv6PayloadLengthMetricTag interface {
	// FromProto unmarshals PatternFlowIpv6PayloadLengthMetricTag from protobuf object *otg.PatternFlowIpv6PayloadLengthMetricTag
	FromProto(msg *otg.PatternFlowIpv6PayloadLengthMetricTag) (PatternFlowIpv6PayloadLengthMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv6PayloadLengthMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6PayloadLengthMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6PayloadLengthMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6PayloadLengthMetricTag) Marshal() marshalPatternFlowIpv6PayloadLengthMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6PayloadLengthMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6PayloadLengthMetricTag) Unmarshal() unMarshalPatternFlowIpv6PayloadLengthMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6PayloadLengthMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6PayloadLengthMetricTag) ToProto() (*otg.PatternFlowIpv6PayloadLengthMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6PayloadLengthMetricTag) FromProto(msg *otg.PatternFlowIpv6PayloadLengthMetricTag) (PatternFlowIpv6PayloadLengthMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6PayloadLengthMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6PayloadLengthMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6PayloadLengthMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6PayloadLengthMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6PayloadLengthMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6PayloadLengthMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv6PayloadLengthMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6PayloadLengthMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6PayloadLengthMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6PayloadLengthMetricTag) Clone() (PatternFlowIpv6PayloadLengthMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6PayloadLengthMetricTag()
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

// PatternFlowIpv6PayloadLengthMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv6PayloadLengthMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv6PayloadLengthMetricTag to protobuf object *otg.PatternFlowIpv6PayloadLengthMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6PayloadLengthMetricTag
	// setMsg unmarshals PatternFlowIpv6PayloadLengthMetricTag from protobuf object *otg.PatternFlowIpv6PayloadLengthMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6PayloadLengthMetricTag) PatternFlowIpv6PayloadLengthMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6PayloadLengthMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6PayloadLengthMetricTag
	// validate validates PatternFlowIpv6PayloadLengthMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6PayloadLengthMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv6PayloadLengthMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv6PayloadLengthMetricTag
	SetName(value string) PatternFlowIpv6PayloadLengthMetricTag
	// Offset returns uint32, set in PatternFlowIpv6PayloadLengthMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv6PayloadLengthMetricTag
	SetOffset(value uint32) PatternFlowIpv6PayloadLengthMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv6PayloadLengthMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv6PayloadLengthMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv6PayloadLengthMetricTag
	SetLength(value uint32) PatternFlowIpv6PayloadLengthMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv6PayloadLengthMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv6PayloadLengthMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv6PayloadLengthMetricTag object
func (obj *patternFlowIpv6PayloadLengthMetricTag) SetName(value string) PatternFlowIpv6PayloadLengthMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6PayloadLengthMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6PayloadLengthMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv6PayloadLengthMetricTag object
func (obj *patternFlowIpv6PayloadLengthMetricTag) SetOffset(value uint32) PatternFlowIpv6PayloadLengthMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6PayloadLengthMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6PayloadLengthMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv6PayloadLengthMetricTag object
func (obj *patternFlowIpv6PayloadLengthMetricTag) SetLength(value uint32) PatternFlowIpv6PayloadLengthMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv6PayloadLengthMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv6PayloadLengthMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowIpv6PayloadLengthMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6PayloadLengthMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv6PayloadLengthMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv6PayloadLengthMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
