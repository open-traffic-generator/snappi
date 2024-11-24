package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpEcnCwrMetricTag *****
type patternFlowTcpEcnCwrMetricTag struct {
	validation
	obj          *otg.PatternFlowTcpEcnCwrMetricTag
	marshaller   marshalPatternFlowTcpEcnCwrMetricTag
	unMarshaller unMarshalPatternFlowTcpEcnCwrMetricTag
}

func NewPatternFlowTcpEcnCwrMetricTag() PatternFlowTcpEcnCwrMetricTag {
	obj := patternFlowTcpEcnCwrMetricTag{obj: &otg.PatternFlowTcpEcnCwrMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpEcnCwrMetricTag) msg() *otg.PatternFlowTcpEcnCwrMetricTag {
	return obj.obj
}

func (obj *patternFlowTcpEcnCwrMetricTag) setMsg(msg *otg.PatternFlowTcpEcnCwrMetricTag) PatternFlowTcpEcnCwrMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpEcnCwrMetricTag struct {
	obj *patternFlowTcpEcnCwrMetricTag
}

type marshalPatternFlowTcpEcnCwrMetricTag interface {
	// ToProto marshals PatternFlowTcpEcnCwrMetricTag to protobuf object *otg.PatternFlowTcpEcnCwrMetricTag
	ToProto() (*otg.PatternFlowTcpEcnCwrMetricTag, error)
	// ToPbText marshals PatternFlowTcpEcnCwrMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpEcnCwrMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpEcnCwrMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpEcnCwrMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpEcnCwrMetricTag struct {
	obj *patternFlowTcpEcnCwrMetricTag
}

type unMarshalPatternFlowTcpEcnCwrMetricTag interface {
	// FromProto unmarshals PatternFlowTcpEcnCwrMetricTag from protobuf object *otg.PatternFlowTcpEcnCwrMetricTag
	FromProto(msg *otg.PatternFlowTcpEcnCwrMetricTag) (PatternFlowTcpEcnCwrMetricTag, error)
	// FromPbText unmarshals PatternFlowTcpEcnCwrMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpEcnCwrMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpEcnCwrMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpEcnCwrMetricTag) Marshal() marshalPatternFlowTcpEcnCwrMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpEcnCwrMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpEcnCwrMetricTag) Unmarshal() unMarshalPatternFlowTcpEcnCwrMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpEcnCwrMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpEcnCwrMetricTag) ToProto() (*otg.PatternFlowTcpEcnCwrMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpEcnCwrMetricTag) FromProto(msg *otg.PatternFlowTcpEcnCwrMetricTag) (PatternFlowTcpEcnCwrMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpEcnCwrMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnCwrMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpEcnCwrMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnCwrMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpEcnCwrMetricTag) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowTcpEcnCwrMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnCwrMetricTag) FromJson(value string) error {
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

func (obj *patternFlowTcpEcnCwrMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnCwrMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnCwrMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpEcnCwrMetricTag) Clone() (PatternFlowTcpEcnCwrMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpEcnCwrMetricTag()
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

// PatternFlowTcpEcnCwrMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowTcpEcnCwrMetricTag interface {
	Validation
	// msg marshals PatternFlowTcpEcnCwrMetricTag to protobuf object *otg.PatternFlowTcpEcnCwrMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpEcnCwrMetricTag
	// setMsg unmarshals PatternFlowTcpEcnCwrMetricTag from protobuf object *otg.PatternFlowTcpEcnCwrMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpEcnCwrMetricTag) PatternFlowTcpEcnCwrMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowTcpEcnCwrMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpEcnCwrMetricTag
	// validate validates PatternFlowTcpEcnCwrMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpEcnCwrMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowTcpEcnCwrMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowTcpEcnCwrMetricTag
	SetName(value string) PatternFlowTcpEcnCwrMetricTag
	// Offset returns uint32, set in PatternFlowTcpEcnCwrMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowTcpEcnCwrMetricTag
	SetOffset(value uint32) PatternFlowTcpEcnCwrMetricTag
	// HasOffset checks if Offset has been set in PatternFlowTcpEcnCwrMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowTcpEcnCwrMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowTcpEcnCwrMetricTag
	SetLength(value uint32) PatternFlowTcpEcnCwrMetricTag
	// HasLength checks if Length has been set in PatternFlowTcpEcnCwrMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowTcpEcnCwrMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowTcpEcnCwrMetricTag object
func (obj *patternFlowTcpEcnCwrMetricTag) SetName(value string) PatternFlowTcpEcnCwrMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpEcnCwrMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpEcnCwrMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowTcpEcnCwrMetricTag object
func (obj *patternFlowTcpEcnCwrMetricTag) SetOffset(value uint32) PatternFlowTcpEcnCwrMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpEcnCwrMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpEcnCwrMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowTcpEcnCwrMetricTag object
func (obj *patternFlowTcpEcnCwrMetricTag) SetLength(value uint32) PatternFlowTcpEcnCwrMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowTcpEcnCwrMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowTcpEcnCwrMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnCwrMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowTcpEcnCwrMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowTcpEcnCwrMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
