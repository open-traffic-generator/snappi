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

// ***** PatternFlowIgmpv1VersionMetricTag *****
type patternFlowIgmpv1VersionMetricTag struct {
	validation
	obj          *otg.PatternFlowIgmpv1VersionMetricTag
	marshaller   marshalPatternFlowIgmpv1VersionMetricTag
	unMarshaller unMarshalPatternFlowIgmpv1VersionMetricTag
}

func NewPatternFlowIgmpv1VersionMetricTag() PatternFlowIgmpv1VersionMetricTag {
	obj := patternFlowIgmpv1VersionMetricTag{obj: &otg.PatternFlowIgmpv1VersionMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv1VersionMetricTag) msg() *otg.PatternFlowIgmpv1VersionMetricTag {
	return obj.obj
}

func (obj *patternFlowIgmpv1VersionMetricTag) setMsg(msg *otg.PatternFlowIgmpv1VersionMetricTag) PatternFlowIgmpv1VersionMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv1VersionMetricTag struct {
	obj *patternFlowIgmpv1VersionMetricTag
}

type marshalPatternFlowIgmpv1VersionMetricTag interface {
	// ToProto marshals PatternFlowIgmpv1VersionMetricTag to protobuf object *otg.PatternFlowIgmpv1VersionMetricTag
	ToProto() (*otg.PatternFlowIgmpv1VersionMetricTag, error)
	// ToPbText marshals PatternFlowIgmpv1VersionMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv1VersionMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv1VersionMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv1VersionMetricTag struct {
	obj *patternFlowIgmpv1VersionMetricTag
}

type unMarshalPatternFlowIgmpv1VersionMetricTag interface {
	// FromProto unmarshals PatternFlowIgmpv1VersionMetricTag from protobuf object *otg.PatternFlowIgmpv1VersionMetricTag
	FromProto(msg *otg.PatternFlowIgmpv1VersionMetricTag) (PatternFlowIgmpv1VersionMetricTag, error)
	// FromPbText unmarshals PatternFlowIgmpv1VersionMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv1VersionMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv1VersionMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv1VersionMetricTag) Marshal() marshalPatternFlowIgmpv1VersionMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv1VersionMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv1VersionMetricTag) Unmarshal() unMarshalPatternFlowIgmpv1VersionMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv1VersionMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv1VersionMetricTag) ToProto() (*otg.PatternFlowIgmpv1VersionMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv1VersionMetricTag) FromProto(msg *otg.PatternFlowIgmpv1VersionMetricTag) (PatternFlowIgmpv1VersionMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv1VersionMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1VersionMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv1VersionMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1VersionMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv1VersionMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1VersionMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIgmpv1VersionMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1VersionMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1VersionMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv1VersionMetricTag) Clone() (PatternFlowIgmpv1VersionMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv1VersionMetricTag()
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

// PatternFlowIgmpv1VersionMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIgmpv1VersionMetricTag interface {
	Validation
	// msg marshals PatternFlowIgmpv1VersionMetricTag to protobuf object *otg.PatternFlowIgmpv1VersionMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv1VersionMetricTag
	// setMsg unmarshals PatternFlowIgmpv1VersionMetricTag from protobuf object *otg.PatternFlowIgmpv1VersionMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv1VersionMetricTag) PatternFlowIgmpv1VersionMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv1VersionMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv1VersionMetricTag
	// validate validates PatternFlowIgmpv1VersionMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv1VersionMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIgmpv1VersionMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIgmpv1VersionMetricTag
	SetName(value string) PatternFlowIgmpv1VersionMetricTag
	// Offset returns uint32, set in PatternFlowIgmpv1VersionMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIgmpv1VersionMetricTag
	SetOffset(value uint32) PatternFlowIgmpv1VersionMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIgmpv1VersionMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIgmpv1VersionMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIgmpv1VersionMetricTag
	SetLength(value uint32) PatternFlowIgmpv1VersionMetricTag
	// HasLength checks if Length has been set in PatternFlowIgmpv1VersionMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIgmpv1VersionMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIgmpv1VersionMetricTag object
func (obj *patternFlowIgmpv1VersionMetricTag) SetName(value string) PatternFlowIgmpv1VersionMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIgmpv1VersionMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIgmpv1VersionMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIgmpv1VersionMetricTag object
func (obj *patternFlowIgmpv1VersionMetricTag) SetOffset(value uint32) PatternFlowIgmpv1VersionMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIgmpv1VersionMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIgmpv1VersionMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIgmpv1VersionMetricTag object
func (obj *patternFlowIgmpv1VersionMetricTag) SetLength(value uint32) PatternFlowIgmpv1VersionMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIgmpv1VersionMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIgmpv1VersionMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowIgmpv1VersionMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1VersionMetricTag.Offset <= 3 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 4 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIgmpv1VersionMetricTag.Length <= 4 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIgmpv1VersionMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(4)
	}

}
