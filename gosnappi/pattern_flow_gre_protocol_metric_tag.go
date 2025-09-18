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

// ***** PatternFlowGreProtocolMetricTag *****
type patternFlowGreProtocolMetricTag struct {
	validation
	obj          *otg.PatternFlowGreProtocolMetricTag
	marshaller   marshalPatternFlowGreProtocolMetricTag
	unMarshaller unMarshalPatternFlowGreProtocolMetricTag
}

func NewPatternFlowGreProtocolMetricTag() PatternFlowGreProtocolMetricTag {
	obj := patternFlowGreProtocolMetricTag{obj: &otg.PatternFlowGreProtocolMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGreProtocolMetricTag) msg() *otg.PatternFlowGreProtocolMetricTag {
	return obj.obj
}

func (obj *patternFlowGreProtocolMetricTag) setMsg(msg *otg.PatternFlowGreProtocolMetricTag) PatternFlowGreProtocolMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGreProtocolMetricTag struct {
	obj *patternFlowGreProtocolMetricTag
}

type marshalPatternFlowGreProtocolMetricTag interface {
	// ToProto marshals PatternFlowGreProtocolMetricTag to protobuf object *otg.PatternFlowGreProtocolMetricTag
	ToProto() (*otg.PatternFlowGreProtocolMetricTag, error)
	// ToPbText marshals PatternFlowGreProtocolMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGreProtocolMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGreProtocolMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGreProtocolMetricTag struct {
	obj *patternFlowGreProtocolMetricTag
}

type unMarshalPatternFlowGreProtocolMetricTag interface {
	// FromProto unmarshals PatternFlowGreProtocolMetricTag from protobuf object *otg.PatternFlowGreProtocolMetricTag
	FromProto(msg *otg.PatternFlowGreProtocolMetricTag) (PatternFlowGreProtocolMetricTag, error)
	// FromPbText unmarshals PatternFlowGreProtocolMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGreProtocolMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGreProtocolMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGreProtocolMetricTag) Marshal() marshalPatternFlowGreProtocolMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGreProtocolMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGreProtocolMetricTag) Unmarshal() unMarshalPatternFlowGreProtocolMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGreProtocolMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGreProtocolMetricTag) ToProto() (*otg.PatternFlowGreProtocolMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGreProtocolMetricTag) FromProto(msg *otg.PatternFlowGreProtocolMetricTag) (PatternFlowGreProtocolMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGreProtocolMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGreProtocolMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGreProtocolMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGreProtocolMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGreProtocolMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGreProtocolMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGreProtocolMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGreProtocolMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGreProtocolMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGreProtocolMetricTag) Clone() (PatternFlowGreProtocolMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGreProtocolMetricTag()
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

// PatternFlowGreProtocolMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGreProtocolMetricTag interface {
	Validation
	// msg marshals PatternFlowGreProtocolMetricTag to protobuf object *otg.PatternFlowGreProtocolMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGreProtocolMetricTag
	// setMsg unmarshals PatternFlowGreProtocolMetricTag from protobuf object *otg.PatternFlowGreProtocolMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGreProtocolMetricTag) PatternFlowGreProtocolMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGreProtocolMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGreProtocolMetricTag
	// validate validates PatternFlowGreProtocolMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGreProtocolMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGreProtocolMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGreProtocolMetricTag
	SetName(value string) PatternFlowGreProtocolMetricTag
	// Offset returns uint32, set in PatternFlowGreProtocolMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGreProtocolMetricTag
	SetOffset(value uint32) PatternFlowGreProtocolMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGreProtocolMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGreProtocolMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGreProtocolMetricTag
	SetLength(value uint32) PatternFlowGreProtocolMetricTag
	// HasLength checks if Length has been set in PatternFlowGreProtocolMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGreProtocolMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGreProtocolMetricTag object
func (obj *patternFlowGreProtocolMetricTag) SetName(value string) PatternFlowGreProtocolMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGreProtocolMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGreProtocolMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGreProtocolMetricTag object
func (obj *patternFlowGreProtocolMetricTag) SetOffset(value uint32) PatternFlowGreProtocolMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGreProtocolMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGreProtocolMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGreProtocolMetricTag object
func (obj *patternFlowGreProtocolMetricTag) SetLength(value uint32) PatternFlowGreProtocolMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGreProtocolMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGreProtocolMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowGreProtocolMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreProtocolMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGreProtocolMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGreProtocolMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
