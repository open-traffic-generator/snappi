package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpCtlFinMetricTag *****
type patternFlowTcpCtlFinMetricTag struct {
	validation
	obj          *otg.PatternFlowTcpCtlFinMetricTag
	marshaller   marshalPatternFlowTcpCtlFinMetricTag
	unMarshaller unMarshalPatternFlowTcpCtlFinMetricTag
}

func NewPatternFlowTcpCtlFinMetricTag() PatternFlowTcpCtlFinMetricTag {
	obj := patternFlowTcpCtlFinMetricTag{obj: &otg.PatternFlowTcpCtlFinMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlFinMetricTag) msg() *otg.PatternFlowTcpCtlFinMetricTag {
	return obj.obj
}

func (obj *patternFlowTcpCtlFinMetricTag) setMsg(msg *otg.PatternFlowTcpCtlFinMetricTag) PatternFlowTcpCtlFinMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlFinMetricTag struct {
	obj *patternFlowTcpCtlFinMetricTag
}

type marshalPatternFlowTcpCtlFinMetricTag interface {
	// ToProto marshals PatternFlowTcpCtlFinMetricTag to protobuf object *otg.PatternFlowTcpCtlFinMetricTag
	ToProto() (*otg.PatternFlowTcpCtlFinMetricTag, error)
	// ToPbText marshals PatternFlowTcpCtlFinMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlFinMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlFinMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpCtlFinMetricTag struct {
	obj *patternFlowTcpCtlFinMetricTag
}

type unMarshalPatternFlowTcpCtlFinMetricTag interface {
	// FromProto unmarshals PatternFlowTcpCtlFinMetricTag from protobuf object *otg.PatternFlowTcpCtlFinMetricTag
	FromProto(msg *otg.PatternFlowTcpCtlFinMetricTag) (PatternFlowTcpCtlFinMetricTag, error)
	// FromPbText unmarshals PatternFlowTcpCtlFinMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlFinMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlFinMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlFinMetricTag) Marshal() marshalPatternFlowTcpCtlFinMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlFinMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlFinMetricTag) Unmarshal() unMarshalPatternFlowTcpCtlFinMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlFinMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlFinMetricTag) ToProto() (*otg.PatternFlowTcpCtlFinMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlFinMetricTag) FromProto(msg *otg.PatternFlowTcpCtlFinMetricTag) (PatternFlowTcpCtlFinMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlFinMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlFinMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlFinMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlFinMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlFinMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlFinMetricTag) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlFinMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlFinMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlFinMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlFinMetricTag) Clone() (PatternFlowTcpCtlFinMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlFinMetricTag()
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

// PatternFlowTcpCtlFinMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowTcpCtlFinMetricTag interface {
	Validation
	// msg marshals PatternFlowTcpCtlFinMetricTag to protobuf object *otg.PatternFlowTcpCtlFinMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlFinMetricTag
	// setMsg unmarshals PatternFlowTcpCtlFinMetricTag from protobuf object *otg.PatternFlowTcpCtlFinMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlFinMetricTag) PatternFlowTcpCtlFinMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlFinMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlFinMetricTag
	// validate validates PatternFlowTcpCtlFinMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlFinMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowTcpCtlFinMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowTcpCtlFinMetricTag
	SetName(value string) PatternFlowTcpCtlFinMetricTag
	// Offset returns uint32, set in PatternFlowTcpCtlFinMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowTcpCtlFinMetricTag
	SetOffset(value uint32) PatternFlowTcpCtlFinMetricTag
	// HasOffset checks if Offset has been set in PatternFlowTcpCtlFinMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowTcpCtlFinMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowTcpCtlFinMetricTag
	SetLength(value uint32) PatternFlowTcpCtlFinMetricTag
	// HasLength checks if Length has been set in PatternFlowTcpCtlFinMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowTcpCtlFinMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowTcpCtlFinMetricTag object
func (obj *patternFlowTcpCtlFinMetricTag) SetName(value string) PatternFlowTcpCtlFinMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpCtlFinMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpCtlFinMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowTcpCtlFinMetricTag object
func (obj *patternFlowTcpCtlFinMetricTag) SetOffset(value uint32) PatternFlowTcpCtlFinMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpCtlFinMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpCtlFinMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowTcpCtlFinMetricTag object
func (obj *patternFlowTcpCtlFinMetricTag) SetLength(value uint32) PatternFlowTcpCtlFinMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowTcpCtlFinMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowTcpCtlFinMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlFinMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowTcpCtlFinMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowTcpCtlFinMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
