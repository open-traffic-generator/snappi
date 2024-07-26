package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpDstPortMetricTag *****
type patternFlowTcpDstPortMetricTag struct {
	validation
	obj          *otg.PatternFlowTcpDstPortMetricTag
	marshaller   marshalPatternFlowTcpDstPortMetricTag
	unMarshaller unMarshalPatternFlowTcpDstPortMetricTag
}

func NewPatternFlowTcpDstPortMetricTag() PatternFlowTcpDstPortMetricTag {
	obj := patternFlowTcpDstPortMetricTag{obj: &otg.PatternFlowTcpDstPortMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpDstPortMetricTag) msg() *otg.PatternFlowTcpDstPortMetricTag {
	return obj.obj
}

func (obj *patternFlowTcpDstPortMetricTag) setMsg(msg *otg.PatternFlowTcpDstPortMetricTag) PatternFlowTcpDstPortMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpDstPortMetricTag struct {
	obj *patternFlowTcpDstPortMetricTag
}

type marshalPatternFlowTcpDstPortMetricTag interface {
	// ToProto marshals PatternFlowTcpDstPortMetricTag to protobuf object *otg.PatternFlowTcpDstPortMetricTag
	ToProto() (*otg.PatternFlowTcpDstPortMetricTag, error)
	// ToPbText marshals PatternFlowTcpDstPortMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpDstPortMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpDstPortMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpDstPortMetricTag struct {
	obj *patternFlowTcpDstPortMetricTag
}

type unMarshalPatternFlowTcpDstPortMetricTag interface {
	// FromProto unmarshals PatternFlowTcpDstPortMetricTag from protobuf object *otg.PatternFlowTcpDstPortMetricTag
	FromProto(msg *otg.PatternFlowTcpDstPortMetricTag) (PatternFlowTcpDstPortMetricTag, error)
	// FromPbText unmarshals PatternFlowTcpDstPortMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpDstPortMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpDstPortMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpDstPortMetricTag) Marshal() marshalPatternFlowTcpDstPortMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpDstPortMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpDstPortMetricTag) Unmarshal() unMarshalPatternFlowTcpDstPortMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpDstPortMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpDstPortMetricTag) ToProto() (*otg.PatternFlowTcpDstPortMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpDstPortMetricTag) FromProto(msg *otg.PatternFlowTcpDstPortMetricTag) (PatternFlowTcpDstPortMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpDstPortMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpDstPortMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpDstPortMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpDstPortMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpDstPortMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpDstPortMetricTag) FromJson(value string) error {
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

func (obj *patternFlowTcpDstPortMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpDstPortMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpDstPortMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpDstPortMetricTag) Clone() (PatternFlowTcpDstPortMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpDstPortMetricTag()
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

// PatternFlowTcpDstPortMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowTcpDstPortMetricTag interface {
	Validation
	// msg marshals PatternFlowTcpDstPortMetricTag to protobuf object *otg.PatternFlowTcpDstPortMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpDstPortMetricTag
	// setMsg unmarshals PatternFlowTcpDstPortMetricTag from protobuf object *otg.PatternFlowTcpDstPortMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpDstPortMetricTag) PatternFlowTcpDstPortMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowTcpDstPortMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpDstPortMetricTag
	// validate validates PatternFlowTcpDstPortMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpDstPortMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowTcpDstPortMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowTcpDstPortMetricTag
	SetName(value string) PatternFlowTcpDstPortMetricTag
	// Offset returns uint32, set in PatternFlowTcpDstPortMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowTcpDstPortMetricTag
	SetOffset(value uint32) PatternFlowTcpDstPortMetricTag
	// HasOffset checks if Offset has been set in PatternFlowTcpDstPortMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowTcpDstPortMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowTcpDstPortMetricTag
	SetLength(value uint32) PatternFlowTcpDstPortMetricTag
	// HasLength checks if Length has been set in PatternFlowTcpDstPortMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowTcpDstPortMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowTcpDstPortMetricTag object
func (obj *patternFlowTcpDstPortMetricTag) SetName(value string) PatternFlowTcpDstPortMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpDstPortMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpDstPortMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowTcpDstPortMetricTag object
func (obj *patternFlowTcpDstPortMetricTag) SetOffset(value uint32) PatternFlowTcpDstPortMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpDstPortMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpDstPortMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowTcpDstPortMetricTag object
func (obj *patternFlowTcpDstPortMetricTag) SetLength(value uint32) PatternFlowTcpDstPortMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowTcpDstPortMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowTcpDstPortMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpDstPortMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowTcpDstPortMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowTcpDstPortMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
