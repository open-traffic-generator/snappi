package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpSrcPortMetricTag *****
type patternFlowTcpSrcPortMetricTag struct {
	validation
	obj          *otg.PatternFlowTcpSrcPortMetricTag
	marshaller   marshalPatternFlowTcpSrcPortMetricTag
	unMarshaller unMarshalPatternFlowTcpSrcPortMetricTag
}

func NewPatternFlowTcpSrcPortMetricTag() PatternFlowTcpSrcPortMetricTag {
	obj := patternFlowTcpSrcPortMetricTag{obj: &otg.PatternFlowTcpSrcPortMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpSrcPortMetricTag) msg() *otg.PatternFlowTcpSrcPortMetricTag {
	return obj.obj
}

func (obj *patternFlowTcpSrcPortMetricTag) setMsg(msg *otg.PatternFlowTcpSrcPortMetricTag) PatternFlowTcpSrcPortMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpSrcPortMetricTag struct {
	obj *patternFlowTcpSrcPortMetricTag
}

type marshalPatternFlowTcpSrcPortMetricTag interface {
	// ToProto marshals PatternFlowTcpSrcPortMetricTag to protobuf object *otg.PatternFlowTcpSrcPortMetricTag
	ToProto() (*otg.PatternFlowTcpSrcPortMetricTag, error)
	// ToPbText marshals PatternFlowTcpSrcPortMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpSrcPortMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpSrcPortMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpSrcPortMetricTag struct {
	obj *patternFlowTcpSrcPortMetricTag
}

type unMarshalPatternFlowTcpSrcPortMetricTag interface {
	// FromProto unmarshals PatternFlowTcpSrcPortMetricTag from protobuf object *otg.PatternFlowTcpSrcPortMetricTag
	FromProto(msg *otg.PatternFlowTcpSrcPortMetricTag) (PatternFlowTcpSrcPortMetricTag, error)
	// FromPbText unmarshals PatternFlowTcpSrcPortMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpSrcPortMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpSrcPortMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpSrcPortMetricTag) Marshal() marshalPatternFlowTcpSrcPortMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpSrcPortMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpSrcPortMetricTag) Unmarshal() unMarshalPatternFlowTcpSrcPortMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpSrcPortMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpSrcPortMetricTag) ToProto() (*otg.PatternFlowTcpSrcPortMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpSrcPortMetricTag) FromProto(msg *otg.PatternFlowTcpSrcPortMetricTag) (PatternFlowTcpSrcPortMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpSrcPortMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpSrcPortMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpSrcPortMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpSrcPortMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpSrcPortMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpSrcPortMetricTag) FromJson(value string) error {
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

func (obj *patternFlowTcpSrcPortMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpSrcPortMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpSrcPortMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpSrcPortMetricTag) Clone() (PatternFlowTcpSrcPortMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpSrcPortMetricTag()
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

// PatternFlowTcpSrcPortMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowTcpSrcPortMetricTag interface {
	Validation
	// msg marshals PatternFlowTcpSrcPortMetricTag to protobuf object *otg.PatternFlowTcpSrcPortMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpSrcPortMetricTag
	// setMsg unmarshals PatternFlowTcpSrcPortMetricTag from protobuf object *otg.PatternFlowTcpSrcPortMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpSrcPortMetricTag) PatternFlowTcpSrcPortMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowTcpSrcPortMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpSrcPortMetricTag
	// validate validates PatternFlowTcpSrcPortMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpSrcPortMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowTcpSrcPortMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowTcpSrcPortMetricTag
	SetName(value string) PatternFlowTcpSrcPortMetricTag
	// Offset returns uint32, set in PatternFlowTcpSrcPortMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowTcpSrcPortMetricTag
	SetOffset(value uint32) PatternFlowTcpSrcPortMetricTag
	// HasOffset checks if Offset has been set in PatternFlowTcpSrcPortMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowTcpSrcPortMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowTcpSrcPortMetricTag
	SetLength(value uint32) PatternFlowTcpSrcPortMetricTag
	// HasLength checks if Length has been set in PatternFlowTcpSrcPortMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowTcpSrcPortMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowTcpSrcPortMetricTag object
func (obj *patternFlowTcpSrcPortMetricTag) SetName(value string) PatternFlowTcpSrcPortMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpSrcPortMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpSrcPortMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowTcpSrcPortMetricTag object
func (obj *patternFlowTcpSrcPortMetricTag) SetOffset(value uint32) PatternFlowTcpSrcPortMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpSrcPortMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpSrcPortMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowTcpSrcPortMetricTag object
func (obj *patternFlowTcpSrcPortMetricTag) SetLength(value uint32) PatternFlowTcpSrcPortMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowTcpSrcPortMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowTcpSrcPortMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpSrcPortMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowTcpSrcPortMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowTcpSrcPortMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
