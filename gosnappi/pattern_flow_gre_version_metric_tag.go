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

// ***** PatternFlowGreVersionMetricTag *****
type patternFlowGreVersionMetricTag struct {
	validation
	obj          *otg.PatternFlowGreVersionMetricTag
	marshaller   marshalPatternFlowGreVersionMetricTag
	unMarshaller unMarshalPatternFlowGreVersionMetricTag
}

func NewPatternFlowGreVersionMetricTag() PatternFlowGreVersionMetricTag {
	obj := patternFlowGreVersionMetricTag{obj: &otg.PatternFlowGreVersionMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGreVersionMetricTag) msg() *otg.PatternFlowGreVersionMetricTag {
	return obj.obj
}

func (obj *patternFlowGreVersionMetricTag) setMsg(msg *otg.PatternFlowGreVersionMetricTag) PatternFlowGreVersionMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGreVersionMetricTag struct {
	obj *patternFlowGreVersionMetricTag
}

type marshalPatternFlowGreVersionMetricTag interface {
	// ToProto marshals PatternFlowGreVersionMetricTag to protobuf object *otg.PatternFlowGreVersionMetricTag
	ToProto() (*otg.PatternFlowGreVersionMetricTag, error)
	// ToPbText marshals PatternFlowGreVersionMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGreVersionMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGreVersionMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGreVersionMetricTag struct {
	obj *patternFlowGreVersionMetricTag
}

type unMarshalPatternFlowGreVersionMetricTag interface {
	// FromProto unmarshals PatternFlowGreVersionMetricTag from protobuf object *otg.PatternFlowGreVersionMetricTag
	FromProto(msg *otg.PatternFlowGreVersionMetricTag) (PatternFlowGreVersionMetricTag, error)
	// FromPbText unmarshals PatternFlowGreVersionMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGreVersionMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGreVersionMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGreVersionMetricTag) Marshal() marshalPatternFlowGreVersionMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGreVersionMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGreVersionMetricTag) Unmarshal() unMarshalPatternFlowGreVersionMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGreVersionMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGreVersionMetricTag) ToProto() (*otg.PatternFlowGreVersionMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGreVersionMetricTag) FromProto(msg *otg.PatternFlowGreVersionMetricTag) (PatternFlowGreVersionMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGreVersionMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGreVersionMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGreVersionMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGreVersionMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGreVersionMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGreVersionMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGreVersionMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGreVersionMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGreVersionMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGreVersionMetricTag) Clone() (PatternFlowGreVersionMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGreVersionMetricTag()
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

// PatternFlowGreVersionMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGreVersionMetricTag interface {
	Validation
	// msg marshals PatternFlowGreVersionMetricTag to protobuf object *otg.PatternFlowGreVersionMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGreVersionMetricTag
	// setMsg unmarshals PatternFlowGreVersionMetricTag from protobuf object *otg.PatternFlowGreVersionMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGreVersionMetricTag) PatternFlowGreVersionMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGreVersionMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGreVersionMetricTag
	// validate validates PatternFlowGreVersionMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGreVersionMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGreVersionMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGreVersionMetricTag
	SetName(value string) PatternFlowGreVersionMetricTag
	// Offset returns uint32, set in PatternFlowGreVersionMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGreVersionMetricTag
	SetOffset(value uint32) PatternFlowGreVersionMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGreVersionMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGreVersionMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGreVersionMetricTag
	SetLength(value uint32) PatternFlowGreVersionMetricTag
	// HasLength checks if Length has been set in PatternFlowGreVersionMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGreVersionMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGreVersionMetricTag object
func (obj *patternFlowGreVersionMetricTag) SetName(value string) PatternFlowGreVersionMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGreVersionMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGreVersionMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGreVersionMetricTag object
func (obj *patternFlowGreVersionMetricTag) SetOffset(value uint32) PatternFlowGreVersionMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGreVersionMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGreVersionMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGreVersionMetricTag object
func (obj *patternFlowGreVersionMetricTag) SetLength(value uint32) PatternFlowGreVersionMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGreVersionMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGreVersionMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowGreVersionMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreVersionMetricTag.Offset <= 2 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGreVersionMetricTag.Length <= 3 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGreVersionMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(3)
	}

}
