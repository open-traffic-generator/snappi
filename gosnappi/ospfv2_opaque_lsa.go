package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2OpaqueLsa *****
type ospfv2OpaqueLsa struct {
	validation
	obj               *otg.Ospfv2OpaqueLsa
	marshaller        marshalOspfv2OpaqueLsa
	unMarshaller      unMarshalOspfv2OpaqueLsa
	commonAttrsHolder Ospfv2CommonAttrs
}

func NewOspfv2OpaqueLsa() Ospfv2OpaqueLsa {
	obj := ospfv2OpaqueLsa{obj: &otg.Ospfv2OpaqueLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2OpaqueLsa) msg() *otg.Ospfv2OpaqueLsa {
	return obj.obj
}

func (obj *ospfv2OpaqueLsa) setMsg(msg *otg.Ospfv2OpaqueLsa) Ospfv2OpaqueLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2OpaqueLsa struct {
	obj *ospfv2OpaqueLsa
}

type marshalOspfv2OpaqueLsa interface {
	// ToProto marshals Ospfv2OpaqueLsa to protobuf object *otg.Ospfv2OpaqueLsa
	ToProto() (*otg.Ospfv2OpaqueLsa, error)
	// ToPbText marshals Ospfv2OpaqueLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2OpaqueLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2OpaqueLsa to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2OpaqueLsa struct {
	obj *ospfv2OpaqueLsa
}

type unMarshalOspfv2OpaqueLsa interface {
	// FromProto unmarshals Ospfv2OpaqueLsa from protobuf object *otg.Ospfv2OpaqueLsa
	FromProto(msg *otg.Ospfv2OpaqueLsa) (Ospfv2OpaqueLsa, error)
	// FromPbText unmarshals Ospfv2OpaqueLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2OpaqueLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2OpaqueLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv2OpaqueLsa) Marshal() marshalOspfv2OpaqueLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2OpaqueLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2OpaqueLsa) Unmarshal() unMarshalOspfv2OpaqueLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2OpaqueLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2OpaqueLsa) ToProto() (*otg.Ospfv2OpaqueLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2OpaqueLsa) FromProto(msg *otg.Ospfv2OpaqueLsa) (Ospfv2OpaqueLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2OpaqueLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsa) FromPbText(value string) error {
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

func (m *marshalospfv2OpaqueLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsa) FromYaml(value string) error {
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

func (m *marshalospfv2OpaqueLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsa) FromJson(value string) error {
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

func (obj *ospfv2OpaqueLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2OpaqueLsa) Clone() (Ospfv2OpaqueLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2OpaqueLsa()
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

func (obj *ospfv2OpaqueLsa) setNil() {
	obj.commonAttrsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2OpaqueLsa is contents of OSPFv2 NSSA-LSA - Type 7.
type Ospfv2OpaqueLsa interface {
	Validation
	// msg marshals Ospfv2OpaqueLsa to protobuf object *otg.Ospfv2OpaqueLsa
	// and doesn't set defaults
	msg() *otg.Ospfv2OpaqueLsa
	// setMsg unmarshals Ospfv2OpaqueLsa from protobuf object *otg.Ospfv2OpaqueLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv2OpaqueLsa) Ospfv2OpaqueLsa
	// provides marshal interface
	Marshal() marshalOspfv2OpaqueLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2OpaqueLsa
	// validate validates Ospfv2OpaqueLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2OpaqueLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CommonAttrs returns Ospfv2CommonAttrs, set in Ospfv2OpaqueLsa.
	// Ospfv2CommonAttrs is tBD
	CommonAttrs() Ospfv2CommonAttrs
	// SetCommonAttrs assigns Ospfv2CommonAttrs provided by user to Ospfv2OpaqueLsa.
	// Ospfv2CommonAttrs is tBD
	SetCommonAttrs(value Ospfv2CommonAttrs) Ospfv2OpaqueLsa
	// OpaqueLsaType returns Ospfv2OpaqueLsaOpaqueLsaTypeEnum, set in Ospfv2OpaqueLsa
	OpaqueLsaType() Ospfv2OpaqueLsaOpaqueLsaTypeEnum
	// SetOpaqueLsaType assigns Ospfv2OpaqueLsaOpaqueLsaTypeEnum provided by user to Ospfv2OpaqueLsa
	SetOpaqueLsaType(value Ospfv2OpaqueLsaOpaqueLsaTypeEnum) Ospfv2OpaqueLsa
	// HasOpaqueLsaType checks if OpaqueLsaType has been set in Ospfv2OpaqueLsa
	HasOpaqueLsaType() bool
	setNil()
}

// Common LSA attributes.
// CommonAttrs returns a Ospfv2CommonAttrs
func (obj *ospfv2OpaqueLsa) CommonAttrs() Ospfv2CommonAttrs {
	if obj.obj.CommonAttrs == nil {
		obj.obj.CommonAttrs = NewOspfv2CommonAttrs().msg()
	}
	if obj.commonAttrsHolder == nil {
		obj.commonAttrsHolder = &ospfv2CommonAttrs{obj: obj.obj.CommonAttrs}
	}
	return obj.commonAttrsHolder
}

// Common LSA attributes.
// SetCommonAttrs sets the Ospfv2CommonAttrs value in the Ospfv2OpaqueLsa object
func (obj *ospfv2OpaqueLsa) SetCommonAttrs(value Ospfv2CommonAttrs) Ospfv2OpaqueLsa {

	obj.commonAttrsHolder = nil
	obj.obj.CommonAttrs = value.msg()

	return obj
}

type Ospfv2OpaqueLsaOpaqueLsaTypeEnum string

// Enum of OpaqueLsaType on Ospfv2OpaqueLsa
var Ospfv2OpaqueLsaOpaqueLsaType = struct {
	OPAQUE_LOCAL  Ospfv2OpaqueLsaOpaqueLsaTypeEnum
	OPAQUE_AREA   Ospfv2OpaqueLsaOpaqueLsaTypeEnum
	OPAQUE_DOMAIN Ospfv2OpaqueLsaOpaqueLsaTypeEnum
}{
	OPAQUE_LOCAL:  Ospfv2OpaqueLsaOpaqueLsaTypeEnum("opaque_local"),
	OPAQUE_AREA:   Ospfv2OpaqueLsaOpaqueLsaTypeEnum("opaque_area"),
	OPAQUE_DOMAIN: Ospfv2OpaqueLsaOpaqueLsaTypeEnum("opaque_domain"),
}

func (obj *ospfv2OpaqueLsa) OpaqueLsaType() Ospfv2OpaqueLsaOpaqueLsaTypeEnum {
	return Ospfv2OpaqueLsaOpaqueLsaTypeEnum(obj.obj.OpaqueLsaType.Enum().String())
}

// The type of Opaque TE LSAs. The LSA type.
// OpaqueLsaType returns a string
func (obj *ospfv2OpaqueLsa) HasOpaqueLsaType() bool {
	return obj.obj.OpaqueLsaType != nil
}

func (obj *ospfv2OpaqueLsa) SetOpaqueLsaType(value Ospfv2OpaqueLsaOpaqueLsaTypeEnum) Ospfv2OpaqueLsa {
	intValue, ok := otg.Ospfv2OpaqueLsa_OpaqueLsaType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2OpaqueLsaOpaqueLsaTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2OpaqueLsa_OpaqueLsaType_Enum(intValue)
	obj.obj.OpaqueLsaType = &enumValue

	return obj
}

func (obj *ospfv2OpaqueLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// CommonAttrs is required
	if obj.obj.CommonAttrs == nil {
		vObj.validationErrors = append(vObj.validationErrors, "CommonAttrs is required field on interface Ospfv2OpaqueLsa")
	}

	if obj.obj.CommonAttrs != nil {

		obj.CommonAttrs().validateObj(vObj, set_default)
	}

}

func (obj *ospfv2OpaqueLsa) setDefault() {

}
