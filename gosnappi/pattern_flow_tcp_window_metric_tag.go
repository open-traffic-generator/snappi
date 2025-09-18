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

// ***** PatternFlowTcpWindowMetricTag *****
type patternFlowTcpWindowMetricTag struct {
	validation
	obj          *otg.PatternFlowTcpWindowMetricTag
	marshaller   marshalPatternFlowTcpWindowMetricTag
	unMarshaller unMarshalPatternFlowTcpWindowMetricTag
}

func NewPatternFlowTcpWindowMetricTag() PatternFlowTcpWindowMetricTag {
	obj := patternFlowTcpWindowMetricTag{obj: &otg.PatternFlowTcpWindowMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpWindowMetricTag) msg() *otg.PatternFlowTcpWindowMetricTag {
	return obj.obj
}

func (obj *patternFlowTcpWindowMetricTag) setMsg(msg *otg.PatternFlowTcpWindowMetricTag) PatternFlowTcpWindowMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpWindowMetricTag struct {
	obj *patternFlowTcpWindowMetricTag
}

type marshalPatternFlowTcpWindowMetricTag interface {
	// ToProto marshals PatternFlowTcpWindowMetricTag to protobuf object *otg.PatternFlowTcpWindowMetricTag
	ToProto() (*otg.PatternFlowTcpWindowMetricTag, error)
	// ToPbText marshals PatternFlowTcpWindowMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpWindowMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpWindowMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpWindowMetricTag struct {
	obj *patternFlowTcpWindowMetricTag
}

type unMarshalPatternFlowTcpWindowMetricTag interface {
	// FromProto unmarshals PatternFlowTcpWindowMetricTag from protobuf object *otg.PatternFlowTcpWindowMetricTag
	FromProto(msg *otg.PatternFlowTcpWindowMetricTag) (PatternFlowTcpWindowMetricTag, error)
	// FromPbText unmarshals PatternFlowTcpWindowMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpWindowMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpWindowMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpWindowMetricTag) Marshal() marshalPatternFlowTcpWindowMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpWindowMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpWindowMetricTag) Unmarshal() unMarshalPatternFlowTcpWindowMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpWindowMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpWindowMetricTag) ToProto() (*otg.PatternFlowTcpWindowMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpWindowMetricTag) FromProto(msg *otg.PatternFlowTcpWindowMetricTag) (PatternFlowTcpWindowMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpWindowMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpWindowMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpWindowMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpWindowMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpWindowMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpWindowMetricTag) FromJson(value string) error {
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

func (obj *patternFlowTcpWindowMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpWindowMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpWindowMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpWindowMetricTag) Clone() (PatternFlowTcpWindowMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpWindowMetricTag()
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

// PatternFlowTcpWindowMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowTcpWindowMetricTag interface {
	Validation
	// msg marshals PatternFlowTcpWindowMetricTag to protobuf object *otg.PatternFlowTcpWindowMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpWindowMetricTag
	// setMsg unmarshals PatternFlowTcpWindowMetricTag from protobuf object *otg.PatternFlowTcpWindowMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpWindowMetricTag) PatternFlowTcpWindowMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowTcpWindowMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpWindowMetricTag
	// validate validates PatternFlowTcpWindowMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpWindowMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowTcpWindowMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowTcpWindowMetricTag
	SetName(value string) PatternFlowTcpWindowMetricTag
	// Offset returns uint32, set in PatternFlowTcpWindowMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowTcpWindowMetricTag
	SetOffset(value uint32) PatternFlowTcpWindowMetricTag
	// HasOffset checks if Offset has been set in PatternFlowTcpWindowMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowTcpWindowMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowTcpWindowMetricTag
	SetLength(value uint32) PatternFlowTcpWindowMetricTag
	// HasLength checks if Length has been set in PatternFlowTcpWindowMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowTcpWindowMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowTcpWindowMetricTag object
func (obj *patternFlowTcpWindowMetricTag) SetName(value string) PatternFlowTcpWindowMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpWindowMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpWindowMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowTcpWindowMetricTag object
func (obj *patternFlowTcpWindowMetricTag) SetOffset(value uint32) PatternFlowTcpWindowMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpWindowMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpWindowMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowTcpWindowMetricTag object
func (obj *patternFlowTcpWindowMetricTag) SetLength(value uint32) PatternFlowTcpWindowMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowTcpWindowMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowTcpWindowMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowTcpWindowMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpWindowMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowTcpWindowMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowTcpWindowMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
