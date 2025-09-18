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

// ***** PatternFlowTcpEcnEchoMetricTag *****
type patternFlowTcpEcnEchoMetricTag struct {
	validation
	obj          *otg.PatternFlowTcpEcnEchoMetricTag
	marshaller   marshalPatternFlowTcpEcnEchoMetricTag
	unMarshaller unMarshalPatternFlowTcpEcnEchoMetricTag
}

func NewPatternFlowTcpEcnEchoMetricTag() PatternFlowTcpEcnEchoMetricTag {
	obj := patternFlowTcpEcnEchoMetricTag{obj: &otg.PatternFlowTcpEcnEchoMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpEcnEchoMetricTag) msg() *otg.PatternFlowTcpEcnEchoMetricTag {
	return obj.obj
}

func (obj *patternFlowTcpEcnEchoMetricTag) setMsg(msg *otg.PatternFlowTcpEcnEchoMetricTag) PatternFlowTcpEcnEchoMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpEcnEchoMetricTag struct {
	obj *patternFlowTcpEcnEchoMetricTag
}

type marshalPatternFlowTcpEcnEchoMetricTag interface {
	// ToProto marshals PatternFlowTcpEcnEchoMetricTag to protobuf object *otg.PatternFlowTcpEcnEchoMetricTag
	ToProto() (*otg.PatternFlowTcpEcnEchoMetricTag, error)
	// ToPbText marshals PatternFlowTcpEcnEchoMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpEcnEchoMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpEcnEchoMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpEcnEchoMetricTag struct {
	obj *patternFlowTcpEcnEchoMetricTag
}

type unMarshalPatternFlowTcpEcnEchoMetricTag interface {
	// FromProto unmarshals PatternFlowTcpEcnEchoMetricTag from protobuf object *otg.PatternFlowTcpEcnEchoMetricTag
	FromProto(msg *otg.PatternFlowTcpEcnEchoMetricTag) (PatternFlowTcpEcnEchoMetricTag, error)
	// FromPbText unmarshals PatternFlowTcpEcnEchoMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpEcnEchoMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpEcnEchoMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpEcnEchoMetricTag) Marshal() marshalPatternFlowTcpEcnEchoMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpEcnEchoMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpEcnEchoMetricTag) Unmarshal() unMarshalPatternFlowTcpEcnEchoMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpEcnEchoMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpEcnEchoMetricTag) ToProto() (*otg.PatternFlowTcpEcnEchoMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpEcnEchoMetricTag) FromProto(msg *otg.PatternFlowTcpEcnEchoMetricTag) (PatternFlowTcpEcnEchoMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpEcnEchoMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnEchoMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpEcnEchoMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnEchoMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpEcnEchoMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnEchoMetricTag) FromJson(value string) error {
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

func (obj *patternFlowTcpEcnEchoMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnEchoMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnEchoMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpEcnEchoMetricTag) Clone() (PatternFlowTcpEcnEchoMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpEcnEchoMetricTag()
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

// PatternFlowTcpEcnEchoMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowTcpEcnEchoMetricTag interface {
	Validation
	// msg marshals PatternFlowTcpEcnEchoMetricTag to protobuf object *otg.PatternFlowTcpEcnEchoMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpEcnEchoMetricTag
	// setMsg unmarshals PatternFlowTcpEcnEchoMetricTag from protobuf object *otg.PatternFlowTcpEcnEchoMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpEcnEchoMetricTag) PatternFlowTcpEcnEchoMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowTcpEcnEchoMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpEcnEchoMetricTag
	// validate validates PatternFlowTcpEcnEchoMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpEcnEchoMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowTcpEcnEchoMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowTcpEcnEchoMetricTag
	SetName(value string) PatternFlowTcpEcnEchoMetricTag
	// Offset returns uint32, set in PatternFlowTcpEcnEchoMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowTcpEcnEchoMetricTag
	SetOffset(value uint32) PatternFlowTcpEcnEchoMetricTag
	// HasOffset checks if Offset has been set in PatternFlowTcpEcnEchoMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowTcpEcnEchoMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowTcpEcnEchoMetricTag
	SetLength(value uint32) PatternFlowTcpEcnEchoMetricTag
	// HasLength checks if Length has been set in PatternFlowTcpEcnEchoMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowTcpEcnEchoMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowTcpEcnEchoMetricTag object
func (obj *patternFlowTcpEcnEchoMetricTag) SetName(value string) PatternFlowTcpEcnEchoMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpEcnEchoMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpEcnEchoMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowTcpEcnEchoMetricTag object
func (obj *patternFlowTcpEcnEchoMetricTag) SetOffset(value uint32) PatternFlowTcpEcnEchoMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpEcnEchoMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpEcnEchoMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowTcpEcnEchoMetricTag object
func (obj *patternFlowTcpEcnEchoMetricTag) SetLength(value uint32) PatternFlowTcpEcnEchoMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowTcpEcnEchoMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowTcpEcnEchoMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowTcpEcnEchoMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnEchoMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowTcpEcnEchoMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowTcpEcnEchoMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
