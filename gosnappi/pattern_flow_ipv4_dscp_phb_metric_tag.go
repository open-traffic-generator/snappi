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

// ***** PatternFlowIpv4DscpPhbMetricTag *****
type patternFlowIpv4DscpPhbMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4DscpPhbMetricTag
	marshaller   marshalPatternFlowIpv4DscpPhbMetricTag
	unMarshaller unMarshalPatternFlowIpv4DscpPhbMetricTag
}

func NewPatternFlowIpv4DscpPhbMetricTag() PatternFlowIpv4DscpPhbMetricTag {
	obj := patternFlowIpv4DscpPhbMetricTag{obj: &otg.PatternFlowIpv4DscpPhbMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4DscpPhbMetricTag) msg() *otg.PatternFlowIpv4DscpPhbMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4DscpPhbMetricTag) setMsg(msg *otg.PatternFlowIpv4DscpPhbMetricTag) PatternFlowIpv4DscpPhbMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4DscpPhbMetricTag struct {
	obj *patternFlowIpv4DscpPhbMetricTag
}

type marshalPatternFlowIpv4DscpPhbMetricTag interface {
	// ToProto marshals PatternFlowIpv4DscpPhbMetricTag to protobuf object *otg.PatternFlowIpv4DscpPhbMetricTag
	ToProto() (*otg.PatternFlowIpv4DscpPhbMetricTag, error)
	// ToPbText marshals PatternFlowIpv4DscpPhbMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4DscpPhbMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4DscpPhbMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4DscpPhbMetricTag struct {
	obj *patternFlowIpv4DscpPhbMetricTag
}

type unMarshalPatternFlowIpv4DscpPhbMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4DscpPhbMetricTag from protobuf object *otg.PatternFlowIpv4DscpPhbMetricTag
	FromProto(msg *otg.PatternFlowIpv4DscpPhbMetricTag) (PatternFlowIpv4DscpPhbMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4DscpPhbMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4DscpPhbMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4DscpPhbMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4DscpPhbMetricTag) Marshal() marshalPatternFlowIpv4DscpPhbMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4DscpPhbMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4DscpPhbMetricTag) Unmarshal() unMarshalPatternFlowIpv4DscpPhbMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4DscpPhbMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4DscpPhbMetricTag) ToProto() (*otg.PatternFlowIpv4DscpPhbMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4DscpPhbMetricTag) FromProto(msg *otg.PatternFlowIpv4DscpPhbMetricTag) (PatternFlowIpv4DscpPhbMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4DscpPhbMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpPhbMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4DscpPhbMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpPhbMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4DscpPhbMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpPhbMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4DscpPhbMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DscpPhbMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DscpPhbMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4DscpPhbMetricTag) Clone() (PatternFlowIpv4DscpPhbMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4DscpPhbMetricTag()
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

// PatternFlowIpv4DscpPhbMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4DscpPhbMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4DscpPhbMetricTag to protobuf object *otg.PatternFlowIpv4DscpPhbMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4DscpPhbMetricTag
	// setMsg unmarshals PatternFlowIpv4DscpPhbMetricTag from protobuf object *otg.PatternFlowIpv4DscpPhbMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4DscpPhbMetricTag) PatternFlowIpv4DscpPhbMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4DscpPhbMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4DscpPhbMetricTag
	// validate validates PatternFlowIpv4DscpPhbMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4DscpPhbMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4DscpPhbMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4DscpPhbMetricTag
	SetName(value string) PatternFlowIpv4DscpPhbMetricTag
	// Offset returns uint32, set in PatternFlowIpv4DscpPhbMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4DscpPhbMetricTag
	SetOffset(value uint32) PatternFlowIpv4DscpPhbMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4DscpPhbMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4DscpPhbMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4DscpPhbMetricTag
	SetLength(value uint32) PatternFlowIpv4DscpPhbMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4DscpPhbMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4DscpPhbMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4DscpPhbMetricTag object
func (obj *patternFlowIpv4DscpPhbMetricTag) SetName(value string) PatternFlowIpv4DscpPhbMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4DscpPhbMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4DscpPhbMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4DscpPhbMetricTag object
func (obj *patternFlowIpv4DscpPhbMetricTag) SetOffset(value uint32) PatternFlowIpv4DscpPhbMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4DscpPhbMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4DscpPhbMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4DscpPhbMetricTag object
func (obj *patternFlowIpv4DscpPhbMetricTag) SetLength(value uint32) PatternFlowIpv4DscpPhbMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4DscpPhbMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4DscpPhbMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowIpv4DscpPhbMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 5 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpPhbMetricTag.Offset <= 5 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 6 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4DscpPhbMetricTag.Length <= 6 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4DscpPhbMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(6)
	}

}
