package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpDataOffsetMetricTag *****
type patternFlowTcpDataOffsetMetricTag struct {
	validation
	obj          *otg.PatternFlowTcpDataOffsetMetricTag
	marshaller   marshalPatternFlowTcpDataOffsetMetricTag
	unMarshaller unMarshalPatternFlowTcpDataOffsetMetricTag
}

func NewPatternFlowTcpDataOffsetMetricTag() PatternFlowTcpDataOffsetMetricTag {
	obj := patternFlowTcpDataOffsetMetricTag{obj: &otg.PatternFlowTcpDataOffsetMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpDataOffsetMetricTag) msg() *otg.PatternFlowTcpDataOffsetMetricTag {
	return obj.obj
}

func (obj *patternFlowTcpDataOffsetMetricTag) setMsg(msg *otg.PatternFlowTcpDataOffsetMetricTag) PatternFlowTcpDataOffsetMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpDataOffsetMetricTag struct {
	obj *patternFlowTcpDataOffsetMetricTag
}

type marshalPatternFlowTcpDataOffsetMetricTag interface {
	// ToProto marshals PatternFlowTcpDataOffsetMetricTag to protobuf object *otg.PatternFlowTcpDataOffsetMetricTag
	ToProto() (*otg.PatternFlowTcpDataOffsetMetricTag, error)
	// ToPbText marshals PatternFlowTcpDataOffsetMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpDataOffsetMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpDataOffsetMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpDataOffsetMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpDataOffsetMetricTag struct {
	obj *patternFlowTcpDataOffsetMetricTag
}

type unMarshalPatternFlowTcpDataOffsetMetricTag interface {
	// FromProto unmarshals PatternFlowTcpDataOffsetMetricTag from protobuf object *otg.PatternFlowTcpDataOffsetMetricTag
	FromProto(msg *otg.PatternFlowTcpDataOffsetMetricTag) (PatternFlowTcpDataOffsetMetricTag, error)
	// FromPbText unmarshals PatternFlowTcpDataOffsetMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpDataOffsetMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpDataOffsetMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpDataOffsetMetricTag) Marshal() marshalPatternFlowTcpDataOffsetMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpDataOffsetMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpDataOffsetMetricTag) Unmarshal() unMarshalPatternFlowTcpDataOffsetMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpDataOffsetMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpDataOffsetMetricTag) ToProto() (*otg.PatternFlowTcpDataOffsetMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpDataOffsetMetricTag) FromProto(msg *otg.PatternFlowTcpDataOffsetMetricTag) (PatternFlowTcpDataOffsetMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpDataOffsetMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpDataOffsetMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpDataOffsetMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpDataOffsetMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpDataOffsetMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowTcpDataOffsetMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpDataOffsetMetricTag) FromJson(value string) error {
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

func (obj *patternFlowTcpDataOffsetMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpDataOffsetMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpDataOffsetMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpDataOffsetMetricTag) Clone() (PatternFlowTcpDataOffsetMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpDataOffsetMetricTag()
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

// PatternFlowTcpDataOffsetMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowTcpDataOffsetMetricTag interface {
	Validation
	// msg marshals PatternFlowTcpDataOffsetMetricTag to protobuf object *otg.PatternFlowTcpDataOffsetMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpDataOffsetMetricTag
	// setMsg unmarshals PatternFlowTcpDataOffsetMetricTag from protobuf object *otg.PatternFlowTcpDataOffsetMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpDataOffsetMetricTag) PatternFlowTcpDataOffsetMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowTcpDataOffsetMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpDataOffsetMetricTag
	// validate validates PatternFlowTcpDataOffsetMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpDataOffsetMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowTcpDataOffsetMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowTcpDataOffsetMetricTag
	SetName(value string) PatternFlowTcpDataOffsetMetricTag
	// Offset returns uint32, set in PatternFlowTcpDataOffsetMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowTcpDataOffsetMetricTag
	SetOffset(value uint32) PatternFlowTcpDataOffsetMetricTag
	// HasOffset checks if Offset has been set in PatternFlowTcpDataOffsetMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowTcpDataOffsetMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowTcpDataOffsetMetricTag
	SetLength(value uint32) PatternFlowTcpDataOffsetMetricTag
	// HasLength checks if Length has been set in PatternFlowTcpDataOffsetMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowTcpDataOffsetMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowTcpDataOffsetMetricTag object
func (obj *patternFlowTcpDataOffsetMetricTag) SetName(value string) PatternFlowTcpDataOffsetMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpDataOffsetMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpDataOffsetMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowTcpDataOffsetMetricTag object
func (obj *patternFlowTcpDataOffsetMetricTag) SetOffset(value uint32) PatternFlowTcpDataOffsetMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpDataOffsetMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpDataOffsetMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowTcpDataOffsetMetricTag object
func (obj *patternFlowTcpDataOffsetMetricTag) SetLength(value uint32) PatternFlowTcpDataOffsetMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowTcpDataOffsetMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowTcpDataOffsetMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpDataOffsetMetricTag.Offset <= 3 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 4 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowTcpDataOffsetMetricTag.Length <= 4 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowTcpDataOffsetMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(4)
	}

}
