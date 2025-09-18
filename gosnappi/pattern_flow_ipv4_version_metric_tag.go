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

// ***** PatternFlowIpv4VersionMetricTag *****
type patternFlowIpv4VersionMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4VersionMetricTag
	marshaller   marshalPatternFlowIpv4VersionMetricTag
	unMarshaller unMarshalPatternFlowIpv4VersionMetricTag
}

func NewPatternFlowIpv4VersionMetricTag() PatternFlowIpv4VersionMetricTag {
	obj := patternFlowIpv4VersionMetricTag{obj: &otg.PatternFlowIpv4VersionMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4VersionMetricTag) msg() *otg.PatternFlowIpv4VersionMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4VersionMetricTag) setMsg(msg *otg.PatternFlowIpv4VersionMetricTag) PatternFlowIpv4VersionMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4VersionMetricTag struct {
	obj *patternFlowIpv4VersionMetricTag
}

type marshalPatternFlowIpv4VersionMetricTag interface {
	// ToProto marshals PatternFlowIpv4VersionMetricTag to protobuf object *otg.PatternFlowIpv4VersionMetricTag
	ToProto() (*otg.PatternFlowIpv4VersionMetricTag, error)
	// ToPbText marshals PatternFlowIpv4VersionMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4VersionMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4VersionMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4VersionMetricTag struct {
	obj *patternFlowIpv4VersionMetricTag
}

type unMarshalPatternFlowIpv4VersionMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4VersionMetricTag from protobuf object *otg.PatternFlowIpv4VersionMetricTag
	FromProto(msg *otg.PatternFlowIpv4VersionMetricTag) (PatternFlowIpv4VersionMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4VersionMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4VersionMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4VersionMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4VersionMetricTag) Marshal() marshalPatternFlowIpv4VersionMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4VersionMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4VersionMetricTag) Unmarshal() unMarshalPatternFlowIpv4VersionMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4VersionMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4VersionMetricTag) ToProto() (*otg.PatternFlowIpv4VersionMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4VersionMetricTag) FromProto(msg *otg.PatternFlowIpv4VersionMetricTag) (PatternFlowIpv4VersionMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4VersionMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4VersionMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4VersionMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4VersionMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4VersionMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4VersionMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4VersionMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4VersionMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4VersionMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4VersionMetricTag) Clone() (PatternFlowIpv4VersionMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4VersionMetricTag()
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

// PatternFlowIpv4VersionMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4VersionMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4VersionMetricTag to protobuf object *otg.PatternFlowIpv4VersionMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4VersionMetricTag
	// setMsg unmarshals PatternFlowIpv4VersionMetricTag from protobuf object *otg.PatternFlowIpv4VersionMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4VersionMetricTag) PatternFlowIpv4VersionMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4VersionMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4VersionMetricTag
	// validate validates PatternFlowIpv4VersionMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4VersionMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4VersionMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4VersionMetricTag
	SetName(value string) PatternFlowIpv4VersionMetricTag
	// Offset returns uint32, set in PatternFlowIpv4VersionMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4VersionMetricTag
	SetOffset(value uint32) PatternFlowIpv4VersionMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4VersionMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4VersionMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4VersionMetricTag
	SetLength(value uint32) PatternFlowIpv4VersionMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4VersionMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4VersionMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4VersionMetricTag object
func (obj *patternFlowIpv4VersionMetricTag) SetName(value string) PatternFlowIpv4VersionMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4VersionMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4VersionMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4VersionMetricTag object
func (obj *patternFlowIpv4VersionMetricTag) SetOffset(value uint32) PatternFlowIpv4VersionMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4VersionMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4VersionMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4VersionMetricTag object
func (obj *patternFlowIpv4VersionMetricTag) SetLength(value uint32) PatternFlowIpv4VersionMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4VersionMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4VersionMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowIpv4VersionMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4VersionMetricTag.Offset <= 3 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 4 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4VersionMetricTag.Length <= 4 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4VersionMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(4)
	}

}
