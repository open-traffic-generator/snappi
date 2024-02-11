package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagProtocol *****
type lagProtocol struct {
	validation
	obj          *otg.LagProtocol
	marshaller   marshalLagProtocol
	unMarshaller unMarshalLagProtocol
	lacpHolder   LagProtocolLacp
	staticHolder LagProtocolStatic
}

func NewLagProtocol() LagProtocol {
	obj := lagProtocol{obj: &otg.LagProtocol{}}
	obj.setDefault()
	return &obj
}

func (obj *lagProtocol) msg() *otg.LagProtocol {
	return obj.obj
}

func (obj *lagProtocol) setMsg(msg *otg.LagProtocol) LagProtocol {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagProtocol struct {
	obj *lagProtocol
}

type marshalLagProtocol interface {
	// ToProto marshals LagProtocol to protobuf object *otg.LagProtocol
	ToProto() (*otg.LagProtocol, error)
	// ToPbText marshals LagProtocol to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagProtocol to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagProtocol to JSON text
	ToJson() (string, error)
}

type unMarshallagProtocol struct {
	obj *lagProtocol
}

type unMarshalLagProtocol interface {
	// FromProto unmarshals LagProtocol from protobuf object *otg.LagProtocol
	FromProto(msg *otg.LagProtocol) (LagProtocol, error)
	// FromPbText unmarshals LagProtocol from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagProtocol from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagProtocol from JSON text
	FromJson(value string) error
}

func (obj *lagProtocol) Marshal() marshalLagProtocol {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagProtocol{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagProtocol) Unmarshal() unMarshalLagProtocol {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagProtocol{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagProtocol) ToProto() (*otg.LagProtocol, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagProtocol) FromProto(msg *otg.LagProtocol) (LagProtocol, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagProtocol) ToPbText() (string, error) {
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

func (m *unMarshallagProtocol) FromPbText(value string) error {
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

func (m *marshallagProtocol) ToYaml() (string, error) {
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

func (m *unMarshallagProtocol) FromYaml(value string) error {
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

func (m *marshallagProtocol) ToJson() (string, error) {
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

func (m *unMarshallagProtocol) FromJson(value string) error {
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

func (obj *lagProtocol) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagProtocol) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagProtocol) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagProtocol) Clone() (LagProtocol, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagProtocol()
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

func (obj *lagProtocol) setNil() {
	obj.lacpHolder = nil
	obj.staticHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LagProtocol is description is TBD
type LagProtocol interface {
	Validation
	// msg marshals LagProtocol to protobuf object *otg.LagProtocol
	// and doesn't set defaults
	msg() *otg.LagProtocol
	// setMsg unmarshals LagProtocol from protobuf object *otg.LagProtocol
	// and doesn't set defaults
	setMsg(*otg.LagProtocol) LagProtocol
	// provides marshal interface
	Marshal() marshalLagProtocol
	// provides unmarshal interface
	Unmarshal() unMarshalLagProtocol
	// validate validates LagProtocol
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagProtocol, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns LagProtocolChoiceEnum, set in LagProtocol
	Choice() LagProtocolChoiceEnum
	// setChoice assigns LagProtocolChoiceEnum provided by user to LagProtocol
	setChoice(value LagProtocolChoiceEnum) LagProtocol
	// HasChoice checks if Choice has been set in LagProtocol
	HasChoice() bool
	// Lacp returns LagProtocolLacp, set in LagProtocol.
	// LagProtocolLacp is the container for link aggregation control protocol settings of a LAG (ports group).
	Lacp() LagProtocolLacp
	// SetLacp assigns LagProtocolLacp provided by user to LagProtocol.
	// LagProtocolLacp is the container for link aggregation control protocol settings of a LAG (ports group).
	SetLacp(value LagProtocolLacp) LagProtocol
	// HasLacp checks if Lacp has been set in LagProtocol
	HasLacp() bool
	// Static returns LagProtocolStatic, set in LagProtocol.
	// LagProtocolStatic is the container for static link aggregation protocol settings.
	Static() LagProtocolStatic
	// SetStatic assigns LagProtocolStatic provided by user to LagProtocol.
	// LagProtocolStatic is the container for static link aggregation protocol settings.
	SetStatic(value LagProtocolStatic) LagProtocol
	// HasStatic checks if Static has been set in LagProtocol
	HasStatic() bool
	setNil()
}

type LagProtocolChoiceEnum string

// Enum of Choice on LagProtocol
var LagProtocolChoice = struct {
	LACP   LagProtocolChoiceEnum
	STATIC LagProtocolChoiceEnum
}{
	LACP:   LagProtocolChoiceEnum("lacp"),
	STATIC: LagProtocolChoiceEnum("static"),
}

func (obj *lagProtocol) Choice() LagProtocolChoiceEnum {
	return LagProtocolChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of controlling protocol for the LAG (ports group).
// Choice returns a string
func (obj *lagProtocol) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *lagProtocol) setChoice(value LagProtocolChoiceEnum) LagProtocol {
	intValue, ok := otg.LagProtocol_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LagProtocolChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.LagProtocol_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Static = nil
	obj.staticHolder = nil
	obj.obj.Lacp = nil
	obj.lacpHolder = nil

	if value == LagProtocolChoice.LACP {
		obj.obj.Lacp = NewLagProtocolLacp().msg()
	}

	if value == LagProtocolChoice.STATIC {
		obj.obj.Static = NewLagProtocolStatic().msg()
	}

	return obj
}

// description is TBD
// Lacp returns a LagProtocolLacp
func (obj *lagProtocol) Lacp() LagProtocolLacp {
	if obj.obj.Lacp == nil {
		obj.setChoice(LagProtocolChoice.LACP)
	}
	if obj.lacpHolder == nil {
		obj.lacpHolder = &lagProtocolLacp{obj: obj.obj.Lacp}
	}
	return obj.lacpHolder
}

// description is TBD
// Lacp returns a LagProtocolLacp
func (obj *lagProtocol) HasLacp() bool {
	return obj.obj.Lacp != nil
}

// description is TBD
// SetLacp sets the LagProtocolLacp value in the LagProtocol object
func (obj *lagProtocol) SetLacp(value LagProtocolLacp) LagProtocol {
	obj.setChoice(LagProtocolChoice.LACP)
	obj.lacpHolder = nil
	obj.obj.Lacp = value.msg()

	return obj
}

// description is TBD
// Static returns a LagProtocolStatic
func (obj *lagProtocol) Static() LagProtocolStatic {
	if obj.obj.Static == nil {
		obj.setChoice(LagProtocolChoice.STATIC)
	}
	if obj.staticHolder == nil {
		obj.staticHolder = &lagProtocolStatic{obj: obj.obj.Static}
	}
	return obj.staticHolder
}

// description is TBD
// Static returns a LagProtocolStatic
func (obj *lagProtocol) HasStatic() bool {
	return obj.obj.Static != nil
}

// description is TBD
// SetStatic sets the LagProtocolStatic value in the LagProtocol object
func (obj *lagProtocol) SetStatic(value LagProtocolStatic) LagProtocol {
	obj.setChoice(LagProtocolChoice.STATIC)
	obj.staticHolder = nil
	obj.obj.Static = value.msg()

	return obj
}

func (obj *lagProtocol) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Lacp != nil {

		obj.Lacp().validateObj(vObj, set_default)
	}

	if obj.obj.Static != nil {

		obj.Static().validateObj(vObj, set_default)
	}

}

func (obj *lagProtocol) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(LagProtocolChoice.LACP)

	}

}
