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

// ***** PatternFlowArpProtocolLengthMetricTag *****
type patternFlowArpProtocolLengthMetricTag struct {
	validation
	obj          *otg.PatternFlowArpProtocolLengthMetricTag
	marshaller   marshalPatternFlowArpProtocolLengthMetricTag
	unMarshaller unMarshalPatternFlowArpProtocolLengthMetricTag
}

func NewPatternFlowArpProtocolLengthMetricTag() PatternFlowArpProtocolLengthMetricTag {
	obj := patternFlowArpProtocolLengthMetricTag{obj: &otg.PatternFlowArpProtocolLengthMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpProtocolLengthMetricTag) msg() *otg.PatternFlowArpProtocolLengthMetricTag {
	return obj.obj
}

func (obj *patternFlowArpProtocolLengthMetricTag) setMsg(msg *otg.PatternFlowArpProtocolLengthMetricTag) PatternFlowArpProtocolLengthMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpProtocolLengthMetricTag struct {
	obj *patternFlowArpProtocolLengthMetricTag
}

type marshalPatternFlowArpProtocolLengthMetricTag interface {
	// ToProto marshals PatternFlowArpProtocolLengthMetricTag to protobuf object *otg.PatternFlowArpProtocolLengthMetricTag
	ToProto() (*otg.PatternFlowArpProtocolLengthMetricTag, error)
	// ToPbText marshals PatternFlowArpProtocolLengthMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpProtocolLengthMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpProtocolLengthMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpProtocolLengthMetricTag struct {
	obj *patternFlowArpProtocolLengthMetricTag
}

type unMarshalPatternFlowArpProtocolLengthMetricTag interface {
	// FromProto unmarshals PatternFlowArpProtocolLengthMetricTag from protobuf object *otg.PatternFlowArpProtocolLengthMetricTag
	FromProto(msg *otg.PatternFlowArpProtocolLengthMetricTag) (PatternFlowArpProtocolLengthMetricTag, error)
	// FromPbText unmarshals PatternFlowArpProtocolLengthMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpProtocolLengthMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpProtocolLengthMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpProtocolLengthMetricTag) Marshal() marshalPatternFlowArpProtocolLengthMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpProtocolLengthMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpProtocolLengthMetricTag) Unmarshal() unMarshalPatternFlowArpProtocolLengthMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpProtocolLengthMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpProtocolLengthMetricTag) ToProto() (*otg.PatternFlowArpProtocolLengthMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpProtocolLengthMetricTag) FromProto(msg *otg.PatternFlowArpProtocolLengthMetricTag) (PatternFlowArpProtocolLengthMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpProtocolLengthMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolLengthMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpProtocolLengthMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolLengthMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpProtocolLengthMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolLengthMetricTag) FromJson(value string) error {
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

func (obj *patternFlowArpProtocolLengthMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpProtocolLengthMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpProtocolLengthMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpProtocolLengthMetricTag) Clone() (PatternFlowArpProtocolLengthMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpProtocolLengthMetricTag()
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

// PatternFlowArpProtocolLengthMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowArpProtocolLengthMetricTag interface {
	Validation
	// msg marshals PatternFlowArpProtocolLengthMetricTag to protobuf object *otg.PatternFlowArpProtocolLengthMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowArpProtocolLengthMetricTag
	// setMsg unmarshals PatternFlowArpProtocolLengthMetricTag from protobuf object *otg.PatternFlowArpProtocolLengthMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpProtocolLengthMetricTag) PatternFlowArpProtocolLengthMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowArpProtocolLengthMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpProtocolLengthMetricTag
	// validate validates PatternFlowArpProtocolLengthMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpProtocolLengthMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowArpProtocolLengthMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowArpProtocolLengthMetricTag
	SetName(value string) PatternFlowArpProtocolLengthMetricTag
	// Offset returns uint32, set in PatternFlowArpProtocolLengthMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowArpProtocolLengthMetricTag
	SetOffset(value uint32) PatternFlowArpProtocolLengthMetricTag
	// HasOffset checks if Offset has been set in PatternFlowArpProtocolLengthMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowArpProtocolLengthMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowArpProtocolLengthMetricTag
	SetLength(value uint32) PatternFlowArpProtocolLengthMetricTag
	// HasLength checks if Length has been set in PatternFlowArpProtocolLengthMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowArpProtocolLengthMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowArpProtocolLengthMetricTag object
func (obj *patternFlowArpProtocolLengthMetricTag) SetName(value string) PatternFlowArpProtocolLengthMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpProtocolLengthMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpProtocolLengthMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowArpProtocolLengthMetricTag object
func (obj *patternFlowArpProtocolLengthMetricTag) SetOffset(value uint32) PatternFlowArpProtocolLengthMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpProtocolLengthMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpProtocolLengthMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowArpProtocolLengthMetricTag object
func (obj *patternFlowArpProtocolLengthMetricTag) SetLength(value uint32) PatternFlowArpProtocolLengthMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowArpProtocolLengthMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowArpProtocolLengthMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowArpProtocolLengthMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolLengthMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowArpProtocolLengthMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowArpProtocolLengthMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
