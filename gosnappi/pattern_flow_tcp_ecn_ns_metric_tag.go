package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpEcnNsMetricTag *****
type patternFlowTcpEcnNsMetricTag struct {
	validation
	obj          *otg.PatternFlowTcpEcnNsMetricTag
	marshaller   marshalPatternFlowTcpEcnNsMetricTag
	unMarshaller unMarshalPatternFlowTcpEcnNsMetricTag
}

func NewPatternFlowTcpEcnNsMetricTag() PatternFlowTcpEcnNsMetricTag {
	obj := patternFlowTcpEcnNsMetricTag{obj: &otg.PatternFlowTcpEcnNsMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpEcnNsMetricTag) msg() *otg.PatternFlowTcpEcnNsMetricTag {
	return obj.obj
}

func (obj *patternFlowTcpEcnNsMetricTag) setMsg(msg *otg.PatternFlowTcpEcnNsMetricTag) PatternFlowTcpEcnNsMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpEcnNsMetricTag struct {
	obj *patternFlowTcpEcnNsMetricTag
}

type marshalPatternFlowTcpEcnNsMetricTag interface {
	// ToProto marshals PatternFlowTcpEcnNsMetricTag to protobuf object *otg.PatternFlowTcpEcnNsMetricTag
	ToProto() (*otg.PatternFlowTcpEcnNsMetricTag, error)
	// ToPbText marshals PatternFlowTcpEcnNsMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpEcnNsMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpEcnNsMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpEcnNsMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpEcnNsMetricTag struct {
	obj *patternFlowTcpEcnNsMetricTag
}

type unMarshalPatternFlowTcpEcnNsMetricTag interface {
	// FromProto unmarshals PatternFlowTcpEcnNsMetricTag from protobuf object *otg.PatternFlowTcpEcnNsMetricTag
	FromProto(msg *otg.PatternFlowTcpEcnNsMetricTag) (PatternFlowTcpEcnNsMetricTag, error)
	// FromPbText unmarshals PatternFlowTcpEcnNsMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpEcnNsMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpEcnNsMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpEcnNsMetricTag) Marshal() marshalPatternFlowTcpEcnNsMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpEcnNsMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpEcnNsMetricTag) Unmarshal() unMarshalPatternFlowTcpEcnNsMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpEcnNsMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpEcnNsMetricTag) ToProto() (*otg.PatternFlowTcpEcnNsMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpEcnNsMetricTag) FromProto(msg *otg.PatternFlowTcpEcnNsMetricTag) (PatternFlowTcpEcnNsMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpEcnNsMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnNsMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpEcnNsMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnNsMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpEcnNsMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowTcpEcnNsMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnNsMetricTag) FromJson(value string) error {
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

func (obj *patternFlowTcpEcnNsMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnNsMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnNsMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpEcnNsMetricTag) Clone() (PatternFlowTcpEcnNsMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpEcnNsMetricTag()
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

// PatternFlowTcpEcnNsMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowTcpEcnNsMetricTag interface {
	Validation
	// msg marshals PatternFlowTcpEcnNsMetricTag to protobuf object *otg.PatternFlowTcpEcnNsMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpEcnNsMetricTag
	// setMsg unmarshals PatternFlowTcpEcnNsMetricTag from protobuf object *otg.PatternFlowTcpEcnNsMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpEcnNsMetricTag) PatternFlowTcpEcnNsMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowTcpEcnNsMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpEcnNsMetricTag
	// validate validates PatternFlowTcpEcnNsMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpEcnNsMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowTcpEcnNsMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowTcpEcnNsMetricTag
	SetName(value string) PatternFlowTcpEcnNsMetricTag
	// Offset returns uint32, set in PatternFlowTcpEcnNsMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowTcpEcnNsMetricTag
	SetOffset(value uint32) PatternFlowTcpEcnNsMetricTag
	// HasOffset checks if Offset has been set in PatternFlowTcpEcnNsMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowTcpEcnNsMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowTcpEcnNsMetricTag
	SetLength(value uint32) PatternFlowTcpEcnNsMetricTag
	// HasLength checks if Length has been set in PatternFlowTcpEcnNsMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowTcpEcnNsMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowTcpEcnNsMetricTag object
func (obj *patternFlowTcpEcnNsMetricTag) SetName(value string) PatternFlowTcpEcnNsMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpEcnNsMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpEcnNsMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowTcpEcnNsMetricTag object
func (obj *patternFlowTcpEcnNsMetricTag) SetOffset(value uint32) PatternFlowTcpEcnNsMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpEcnNsMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpEcnNsMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowTcpEcnNsMetricTag object
func (obj *patternFlowTcpEcnNsMetricTag) SetLength(value uint32) PatternFlowTcpEcnNsMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowTcpEcnNsMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowTcpEcnNsMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnNsMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowTcpEcnNsMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowTcpEcnNsMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
