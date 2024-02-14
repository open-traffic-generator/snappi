package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathExplicitRouteType1ASNumber *****
type flowRSVPPathExplicitRouteType1ASNumber struct {
	validation
	obj          *otg.FlowRSVPPathExplicitRouteType1ASNumber
	marshaller   marshalFlowRSVPPathExplicitRouteType1ASNumber
	unMarshaller unMarshalFlowRSVPPathExplicitRouteType1ASNumber
	lBitHolder   PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	lengthHolder FlowRSVPExplicitRouteLength
}

func NewFlowRSVPPathExplicitRouteType1ASNumber() FlowRSVPPathExplicitRouteType1ASNumber {
	obj := flowRSVPPathExplicitRouteType1ASNumber{obj: &otg.FlowRSVPPathExplicitRouteType1ASNumber{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathExplicitRouteType1ASNumber) msg() *otg.FlowRSVPPathExplicitRouteType1ASNumber {
	return obj.obj
}

func (obj *flowRSVPPathExplicitRouteType1ASNumber) setMsg(msg *otg.FlowRSVPPathExplicitRouteType1ASNumber) FlowRSVPPathExplicitRouteType1ASNumber {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathExplicitRouteType1ASNumber struct {
	obj *flowRSVPPathExplicitRouteType1ASNumber
}

type marshalFlowRSVPPathExplicitRouteType1ASNumber interface {
	// ToProto marshals FlowRSVPPathExplicitRouteType1ASNumber to protobuf object *otg.FlowRSVPPathExplicitRouteType1ASNumber
	ToProto() (*otg.FlowRSVPPathExplicitRouteType1ASNumber, error)
	// ToPbText marshals FlowRSVPPathExplicitRouteType1ASNumber to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathExplicitRouteType1ASNumber to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathExplicitRouteType1ASNumber to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathExplicitRouteType1ASNumber struct {
	obj *flowRSVPPathExplicitRouteType1ASNumber
}

type unMarshalFlowRSVPPathExplicitRouteType1ASNumber interface {
	// FromProto unmarshals FlowRSVPPathExplicitRouteType1ASNumber from protobuf object *otg.FlowRSVPPathExplicitRouteType1ASNumber
	FromProto(msg *otg.FlowRSVPPathExplicitRouteType1ASNumber) (FlowRSVPPathExplicitRouteType1ASNumber, error)
	// FromPbText unmarshals FlowRSVPPathExplicitRouteType1ASNumber from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathExplicitRouteType1ASNumber from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathExplicitRouteType1ASNumber from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathExplicitRouteType1ASNumber) Marshal() marshalFlowRSVPPathExplicitRouteType1ASNumber {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathExplicitRouteType1ASNumber{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathExplicitRouteType1ASNumber) Unmarshal() unMarshalFlowRSVPPathExplicitRouteType1ASNumber {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathExplicitRouteType1ASNumber{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathExplicitRouteType1ASNumber) ToProto() (*otg.FlowRSVPPathExplicitRouteType1ASNumber, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathExplicitRouteType1ASNumber) FromProto(msg *otg.FlowRSVPPathExplicitRouteType1ASNumber) (FlowRSVPPathExplicitRouteType1ASNumber, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathExplicitRouteType1ASNumber) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathExplicitRouteType1ASNumber) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalflowRSVPPathExplicitRouteType1ASNumber) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathExplicitRouteType1ASNumber) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalflowRSVPPathExplicitRouteType1ASNumber) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathExplicitRouteType1ASNumber) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *flowRSVPPathExplicitRouteType1ASNumber) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathExplicitRouteType1ASNumber) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathExplicitRouteType1ASNumber) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathExplicitRouteType1ASNumber) Clone() (FlowRSVPPathExplicitRouteType1ASNumber, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathExplicitRouteType1ASNumber()
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

func (obj *flowRSVPPathExplicitRouteType1ASNumber) setNil() {
	obj.lBitHolder = nil
	obj.lengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathExplicitRouteType1ASNumber is class = EXPLICIT_ROUTE, Type1 ROUTE_RECORD C-Type = 1 Subobject: Autonomous system number, C-Type: 32
type FlowRSVPPathExplicitRouteType1ASNumber interface {
	Validation
	// msg marshals FlowRSVPPathExplicitRouteType1ASNumber to protobuf object *otg.FlowRSVPPathExplicitRouteType1ASNumber
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathExplicitRouteType1ASNumber
	// setMsg unmarshals FlowRSVPPathExplicitRouteType1ASNumber from protobuf object *otg.FlowRSVPPathExplicitRouteType1ASNumber
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathExplicitRouteType1ASNumber) FlowRSVPPathExplicitRouteType1ASNumber
	// provides marshal interface
	Marshal() marshalFlowRSVPPathExplicitRouteType1ASNumber
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathExplicitRouteType1ASNumber
	// validate validates FlowRSVPPathExplicitRouteType1ASNumber
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathExplicitRouteType1ASNumber, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LBit returns PatternFlowRSVPPathExplicitRouteType1ASNumberLBit, set in FlowRSVPPathExplicitRouteType1ASNumber.
	// PatternFlowRSVPPathExplicitRouteType1ASNumberLBit is the L bit is an attribute of the subobject. The L bit is set if the subobject represents a loose hop in the explicit route. If the bit is not set, the subobject represents a strict hop in the explicit route.
	LBit() PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	// SetLBit assigns PatternFlowRSVPPathExplicitRouteType1ASNumberLBit provided by user to FlowRSVPPathExplicitRouteType1ASNumber.
	// PatternFlowRSVPPathExplicitRouteType1ASNumberLBit is the L bit is an attribute of the subobject. The L bit is set if the subobject represents a loose hop in the explicit route. If the bit is not set, the subobject represents a strict hop in the explicit route.
	SetLBit(value PatternFlowRSVPPathExplicitRouteType1ASNumberLBit) FlowRSVPPathExplicitRouteType1ASNumber
	// HasLBit checks if LBit has been set in FlowRSVPPathExplicitRouteType1ASNumber
	HasLBit() bool
	// Length returns FlowRSVPExplicitRouteLength, set in FlowRSVPPathExplicitRouteType1ASNumber.
	// FlowRSVPExplicitRouteLength is description is TBD
	Length() FlowRSVPExplicitRouteLength
	// SetLength assigns FlowRSVPExplicitRouteLength provided by user to FlowRSVPPathExplicitRouteType1ASNumber.
	// FlowRSVPExplicitRouteLength is description is TBD
	SetLength(value FlowRSVPExplicitRouteLength) FlowRSVPPathExplicitRouteType1ASNumber
	// HasLength checks if Length has been set in FlowRSVPPathExplicitRouteType1ASNumber
	HasLength() bool
	// AsNumber returns uint32, set in FlowRSVPPathExplicitRouteType1ASNumber.
	AsNumber() uint32
	// SetAsNumber assigns uint32 provided by user to FlowRSVPPathExplicitRouteType1ASNumber
	SetAsNumber(value uint32) FlowRSVPPathExplicitRouteType1ASNumber
	// HasAsNumber checks if AsNumber has been set in FlowRSVPPathExplicitRouteType1ASNumber
	HasAsNumber() bool
	setNil()
}

// description is TBD
// LBit returns a PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
func (obj *flowRSVPPathExplicitRouteType1ASNumber) LBit() PatternFlowRSVPPathExplicitRouteType1ASNumberLBit {
	if obj.obj.LBit == nil {
		obj.obj.LBit = NewPatternFlowRSVPPathExplicitRouteType1ASNumberLBit().msg()
	}
	if obj.lBitHolder == nil {
		obj.lBitHolder = &patternFlowRSVPPathExplicitRouteType1ASNumberLBit{obj: obj.obj.LBit}
	}
	return obj.lBitHolder
}

// description is TBD
// LBit returns a PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
func (obj *flowRSVPPathExplicitRouteType1ASNumber) HasLBit() bool {
	return obj.obj.LBit != nil
}

// description is TBD
// SetLBit sets the PatternFlowRSVPPathExplicitRouteType1ASNumberLBit value in the FlowRSVPPathExplicitRouteType1ASNumber object
func (obj *flowRSVPPathExplicitRouteType1ASNumber) SetLBit(value PatternFlowRSVPPathExplicitRouteType1ASNumberLBit) FlowRSVPPathExplicitRouteType1ASNumber {

	obj.lBitHolder = nil
	obj.obj.LBit = value.msg()

	return obj
}

// The Length contains the total length of the subobject in bytes,including L, Type and Length fields.   The Length MUST be atleast 4, and MUST be a multiple of 4.
// Length returns a FlowRSVPExplicitRouteLength
func (obj *flowRSVPPathExplicitRouteType1ASNumber) Length() FlowRSVPExplicitRouteLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewFlowRSVPExplicitRouteLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &flowRSVPExplicitRouteLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// The Length contains the total length of the subobject in bytes,including L, Type and Length fields.   The Length MUST be atleast 4, and MUST be a multiple of 4.
// Length returns a FlowRSVPExplicitRouteLength
func (obj *flowRSVPPathExplicitRouteType1ASNumber) HasLength() bool {
	return obj.obj.Length != nil
}

// The Length contains the total length of the subobject in bytes,including L, Type and Length fields.   The Length MUST be atleast 4, and MUST be a multiple of 4.
// SetLength sets the FlowRSVPExplicitRouteLength value in the FlowRSVPPathExplicitRouteType1ASNumber object
func (obj *flowRSVPPathExplicitRouteType1ASNumber) SetLength(value FlowRSVPExplicitRouteLength) FlowRSVPPathExplicitRouteType1ASNumber {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// Autonomous System number to be set in the ERO sub-object that this LSP should traverse through. This field is applicable only if the value of 'type' is set to 'as_number'.
// AsNumber returns a uint32
func (obj *flowRSVPPathExplicitRouteType1ASNumber) AsNumber() uint32 {

	return *obj.obj.AsNumber

}

// Autonomous System number to be set in the ERO sub-object that this LSP should traverse through. This field is applicable only if the value of 'type' is set to 'as_number'.
// AsNumber returns a uint32
func (obj *flowRSVPPathExplicitRouteType1ASNumber) HasAsNumber() bool {
	return obj.obj.AsNumber != nil
}

// Autonomous System number to be set in the ERO sub-object that this LSP should traverse through. This field is applicable only if the value of 'type' is set to 'as_number'.
// SetAsNumber sets the uint32 value in the FlowRSVPPathExplicitRouteType1ASNumber object
func (obj *flowRSVPPathExplicitRouteType1ASNumber) SetAsNumber(value uint32) FlowRSVPPathExplicitRouteType1ASNumber {

	obj.obj.AsNumber = &value
	return obj
}

func (obj *flowRSVPPathExplicitRouteType1ASNumber) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LBit != nil {

		obj.LBit().validateObj(vObj, set_default)
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.AsNumber != nil {

		if *obj.obj.AsNumber > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRSVPPathExplicitRouteType1ASNumber.AsNumber <= 65535 but Got %d", *obj.obj.AsNumber))
		}

	}

}

func (obj *flowRSVPPathExplicitRouteType1ASNumber) setDefault() {
	if obj.obj.AsNumber == nil {
		obj.SetAsNumber(0)
	}

}
