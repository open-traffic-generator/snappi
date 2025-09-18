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

// ***** PatternFlowPppControlMetricTag *****
type patternFlowPppControlMetricTag struct {
	validation
	obj          *otg.PatternFlowPppControlMetricTag
	marshaller   marshalPatternFlowPppControlMetricTag
	unMarshaller unMarshalPatternFlowPppControlMetricTag
}

func NewPatternFlowPppControlMetricTag() PatternFlowPppControlMetricTag {
	obj := patternFlowPppControlMetricTag{obj: &otg.PatternFlowPppControlMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPppControlMetricTag) msg() *otg.PatternFlowPppControlMetricTag {
	return obj.obj
}

func (obj *patternFlowPppControlMetricTag) setMsg(msg *otg.PatternFlowPppControlMetricTag) PatternFlowPppControlMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPppControlMetricTag struct {
	obj *patternFlowPppControlMetricTag
}

type marshalPatternFlowPppControlMetricTag interface {
	// ToProto marshals PatternFlowPppControlMetricTag to protobuf object *otg.PatternFlowPppControlMetricTag
	ToProto() (*otg.PatternFlowPppControlMetricTag, error)
	// ToPbText marshals PatternFlowPppControlMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPppControlMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPppControlMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPppControlMetricTag struct {
	obj *patternFlowPppControlMetricTag
}

type unMarshalPatternFlowPppControlMetricTag interface {
	// FromProto unmarshals PatternFlowPppControlMetricTag from protobuf object *otg.PatternFlowPppControlMetricTag
	FromProto(msg *otg.PatternFlowPppControlMetricTag) (PatternFlowPppControlMetricTag, error)
	// FromPbText unmarshals PatternFlowPppControlMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPppControlMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPppControlMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPppControlMetricTag) Marshal() marshalPatternFlowPppControlMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPppControlMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPppControlMetricTag) Unmarshal() unMarshalPatternFlowPppControlMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPppControlMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPppControlMetricTag) ToProto() (*otg.PatternFlowPppControlMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPppControlMetricTag) FromProto(msg *otg.PatternFlowPppControlMetricTag) (PatternFlowPppControlMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPppControlMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPppControlMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowPppControlMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPppControlMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowPppControlMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPppControlMetricTag) FromJson(value string) error {
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

func (obj *patternFlowPppControlMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPppControlMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPppControlMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPppControlMetricTag) Clone() (PatternFlowPppControlMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPppControlMetricTag()
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

// PatternFlowPppControlMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowPppControlMetricTag interface {
	Validation
	// msg marshals PatternFlowPppControlMetricTag to protobuf object *otg.PatternFlowPppControlMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowPppControlMetricTag
	// setMsg unmarshals PatternFlowPppControlMetricTag from protobuf object *otg.PatternFlowPppControlMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPppControlMetricTag) PatternFlowPppControlMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowPppControlMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPppControlMetricTag
	// validate validates PatternFlowPppControlMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPppControlMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowPppControlMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowPppControlMetricTag
	SetName(value string) PatternFlowPppControlMetricTag
	// Offset returns uint32, set in PatternFlowPppControlMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowPppControlMetricTag
	SetOffset(value uint32) PatternFlowPppControlMetricTag
	// HasOffset checks if Offset has been set in PatternFlowPppControlMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowPppControlMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowPppControlMetricTag
	SetLength(value uint32) PatternFlowPppControlMetricTag
	// HasLength checks if Length has been set in PatternFlowPppControlMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowPppControlMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowPppControlMetricTag object
func (obj *patternFlowPppControlMetricTag) SetName(value string) PatternFlowPppControlMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPppControlMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPppControlMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowPppControlMetricTag object
func (obj *patternFlowPppControlMetricTag) SetOffset(value uint32) PatternFlowPppControlMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPppControlMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPppControlMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowPppControlMetricTag object
func (obj *patternFlowPppControlMetricTag) SetLength(value uint32) PatternFlowPppControlMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowPppControlMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowPppControlMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowPppControlMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppControlMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowPppControlMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowPppControlMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
