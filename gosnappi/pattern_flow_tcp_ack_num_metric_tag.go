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

// ***** PatternFlowTcpAckNumMetricTag *****
type patternFlowTcpAckNumMetricTag struct {
	validation
	obj          *otg.PatternFlowTcpAckNumMetricTag
	marshaller   marshalPatternFlowTcpAckNumMetricTag
	unMarshaller unMarshalPatternFlowTcpAckNumMetricTag
}

func NewPatternFlowTcpAckNumMetricTag() PatternFlowTcpAckNumMetricTag {
	obj := patternFlowTcpAckNumMetricTag{obj: &otg.PatternFlowTcpAckNumMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpAckNumMetricTag) msg() *otg.PatternFlowTcpAckNumMetricTag {
	return obj.obj
}

func (obj *patternFlowTcpAckNumMetricTag) setMsg(msg *otg.PatternFlowTcpAckNumMetricTag) PatternFlowTcpAckNumMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpAckNumMetricTag struct {
	obj *patternFlowTcpAckNumMetricTag
}

type marshalPatternFlowTcpAckNumMetricTag interface {
	// ToProto marshals PatternFlowTcpAckNumMetricTag to protobuf object *otg.PatternFlowTcpAckNumMetricTag
	ToProto() (*otg.PatternFlowTcpAckNumMetricTag, error)
	// ToPbText marshals PatternFlowTcpAckNumMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpAckNumMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpAckNumMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpAckNumMetricTag struct {
	obj *patternFlowTcpAckNumMetricTag
}

type unMarshalPatternFlowTcpAckNumMetricTag interface {
	// FromProto unmarshals PatternFlowTcpAckNumMetricTag from protobuf object *otg.PatternFlowTcpAckNumMetricTag
	FromProto(msg *otg.PatternFlowTcpAckNumMetricTag) (PatternFlowTcpAckNumMetricTag, error)
	// FromPbText unmarshals PatternFlowTcpAckNumMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpAckNumMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpAckNumMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpAckNumMetricTag) Marshal() marshalPatternFlowTcpAckNumMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpAckNumMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpAckNumMetricTag) Unmarshal() unMarshalPatternFlowTcpAckNumMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpAckNumMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpAckNumMetricTag) ToProto() (*otg.PatternFlowTcpAckNumMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpAckNumMetricTag) FromProto(msg *otg.PatternFlowTcpAckNumMetricTag) (PatternFlowTcpAckNumMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpAckNumMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpAckNumMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpAckNumMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpAckNumMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpAckNumMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpAckNumMetricTag) FromJson(value string) error {
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

func (obj *patternFlowTcpAckNumMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpAckNumMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpAckNumMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpAckNumMetricTag) Clone() (PatternFlowTcpAckNumMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpAckNumMetricTag()
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

// PatternFlowTcpAckNumMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowTcpAckNumMetricTag interface {
	Validation
	// msg marshals PatternFlowTcpAckNumMetricTag to protobuf object *otg.PatternFlowTcpAckNumMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpAckNumMetricTag
	// setMsg unmarshals PatternFlowTcpAckNumMetricTag from protobuf object *otg.PatternFlowTcpAckNumMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpAckNumMetricTag) PatternFlowTcpAckNumMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowTcpAckNumMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpAckNumMetricTag
	// validate validates PatternFlowTcpAckNumMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpAckNumMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowTcpAckNumMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowTcpAckNumMetricTag
	SetName(value string) PatternFlowTcpAckNumMetricTag
	// Offset returns uint32, set in PatternFlowTcpAckNumMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowTcpAckNumMetricTag
	SetOffset(value uint32) PatternFlowTcpAckNumMetricTag
	// HasOffset checks if Offset has been set in PatternFlowTcpAckNumMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowTcpAckNumMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowTcpAckNumMetricTag
	SetLength(value uint32) PatternFlowTcpAckNumMetricTag
	// HasLength checks if Length has been set in PatternFlowTcpAckNumMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowTcpAckNumMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowTcpAckNumMetricTag object
func (obj *patternFlowTcpAckNumMetricTag) SetName(value string) PatternFlowTcpAckNumMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpAckNumMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpAckNumMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowTcpAckNumMetricTag object
func (obj *patternFlowTcpAckNumMetricTag) SetOffset(value uint32) PatternFlowTcpAckNumMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpAckNumMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpAckNumMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowTcpAckNumMetricTag object
func (obj *patternFlowTcpAckNumMetricTag) SetLength(value uint32) PatternFlowTcpAckNumMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowTcpAckNumMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowTcpAckNumMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowTcpAckNumMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpAckNumMetricTag.Offset <= 31 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowTcpAckNumMetricTag.Length <= 32 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowTcpAckNumMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(32)
	}

}
