package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2OpaqueLsaAsla *****
type ospfv2OpaqueLsaAsla struct {
	validation
	obj          *otg.Ospfv2OpaqueLsaAsla
	marshaller   marshalOspfv2OpaqueLsaAsla
	unMarshaller unMarshalOspfv2OpaqueLsaAsla
}

func NewOspfv2OpaqueLsaAsla() Ospfv2OpaqueLsaAsla {
	obj := ospfv2OpaqueLsaAsla{obj: &otg.Ospfv2OpaqueLsaAsla{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2OpaqueLsaAsla) msg() *otg.Ospfv2OpaqueLsaAsla {
	return obj.obj
}

func (obj *ospfv2OpaqueLsaAsla) setMsg(msg *otg.Ospfv2OpaqueLsaAsla) Ospfv2OpaqueLsaAsla {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2OpaqueLsaAsla struct {
	obj *ospfv2OpaqueLsaAsla
}

type marshalOspfv2OpaqueLsaAsla interface {
	// ToProto marshals Ospfv2OpaqueLsaAsla to protobuf object *otg.Ospfv2OpaqueLsaAsla
	ToProto() (*otg.Ospfv2OpaqueLsaAsla, error)
	// ToPbText marshals Ospfv2OpaqueLsaAsla to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2OpaqueLsaAsla to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2OpaqueLsaAsla to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2OpaqueLsaAsla struct {
	obj *ospfv2OpaqueLsaAsla
}

type unMarshalOspfv2OpaqueLsaAsla interface {
	// FromProto unmarshals Ospfv2OpaqueLsaAsla from protobuf object *otg.Ospfv2OpaqueLsaAsla
	FromProto(msg *otg.Ospfv2OpaqueLsaAsla) (Ospfv2OpaqueLsaAsla, error)
	// FromPbText unmarshals Ospfv2OpaqueLsaAsla from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2OpaqueLsaAsla from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2OpaqueLsaAsla from JSON text
	FromJson(value string) error
}

func (obj *ospfv2OpaqueLsaAsla) Marshal() marshalOspfv2OpaqueLsaAsla {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2OpaqueLsaAsla{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2OpaqueLsaAsla) Unmarshal() unMarshalOspfv2OpaqueLsaAsla {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2OpaqueLsaAsla{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2OpaqueLsaAsla) ToProto() (*otg.Ospfv2OpaqueLsaAsla, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2OpaqueLsaAsla) FromProto(msg *otg.Ospfv2OpaqueLsaAsla) (Ospfv2OpaqueLsaAsla, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2OpaqueLsaAsla) ToPbText() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaAsla) FromPbText(value string) error {
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

func (m *marshalospfv2OpaqueLsaAsla) ToYaml() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaAsla) FromYaml(value string) error {
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

func (m *marshalospfv2OpaqueLsaAsla) ToJson() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaAsla) FromJson(value string) error {
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

func (obj *ospfv2OpaqueLsaAsla) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaAsla) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaAsla) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2OpaqueLsaAsla) Clone() (Ospfv2OpaqueLsaAsla, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2OpaqueLsaAsla()
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

// Ospfv2OpaqueLsaAsla is an Application-Specific Link Attribute carried within an ASLA sub-TLV (RFC 9492).
type Ospfv2OpaqueLsaAsla interface {
	Validation
	// msg marshals Ospfv2OpaqueLsaAsla to protobuf object *otg.Ospfv2OpaqueLsaAsla
	// and doesn't set defaults
	msg() *otg.Ospfv2OpaqueLsaAsla
	// setMsg unmarshals Ospfv2OpaqueLsaAsla from protobuf object *otg.Ospfv2OpaqueLsaAsla
	// and doesn't set defaults
	setMsg(*otg.Ospfv2OpaqueLsaAsla) Ospfv2OpaqueLsaAsla
	// provides marshal interface
	Marshal() marshalOspfv2OpaqueLsaAsla
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2OpaqueLsaAsla
	// validate validates Ospfv2OpaqueLsaAsla
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2OpaqueLsaAsla, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns Ospfv2OpaqueLsaAslaTypeEnum, set in Ospfv2OpaqueLsaAsla
	Type() Ospfv2OpaqueLsaAslaTypeEnum
	// SetType assigns Ospfv2OpaqueLsaAslaTypeEnum provided by user to Ospfv2OpaqueLsaAsla
	SetType(value Ospfv2OpaqueLsaAslaTypeEnum) Ospfv2OpaqueLsaAsla
	// HasType checks if Type has been set in Ospfv2OpaqueLsaAsla
	HasType() bool
	// Value returns string, set in Ospfv2OpaqueLsaAsla.
	Value() string
	// SetValue assigns string provided by user to Ospfv2OpaqueLsaAsla
	SetValue(value string) Ospfv2OpaqueLsaAsla
	// HasValue checks if Value has been set in Ospfv2OpaqueLsaAsla
	HasValue() bool
}

type Ospfv2OpaqueLsaAslaTypeEnum string

// Enum of Type on Ospfv2OpaqueLsaAsla
var Ospfv2OpaqueLsaAslaType = struct {
	RSVP_TE                            Ospfv2OpaqueLsaAslaTypeEnum
	SEGMENT_ROUTING_POLICY             Ospfv2OpaqueLsaAslaTypeEnum
	LOOP_FREE_ALTERNATE                Ospfv2OpaqueLsaAslaTypeEnum
	FLEXIBLE_ALGORITHM                 Ospfv2OpaqueLsaAslaTypeEnum
	SHARED_RISK_LINK_GROUP             Ospfv2OpaqueLsaAslaTypeEnum
	UNIDIRECTIONAL_LINK_DELAY          Ospfv2OpaqueLsaAslaTypeEnum
	MIN_MAX_UNIDIRECTIONAL_LINK_DELAY  Ospfv2OpaqueLsaAslaTypeEnum
	UNIDIRECTIONAL_DELAY_VARIATION     Ospfv2OpaqueLsaAslaTypeEnum
	UNIDIRECTIONAL_LINK_LOSS           Ospfv2OpaqueLsaAslaTypeEnum
	UNIDIRECTIONAL_RESIDUAL_BANDWIDTH  Ospfv2OpaqueLsaAslaTypeEnum
	UNIDIRECTIONAL_AVAILABLE_BANDWIDTH Ospfv2OpaqueLsaAslaTypeEnum
	UNIDIRECTIONAL_UTILIZED_BANDWIDTH  Ospfv2OpaqueLsaAslaTypeEnum
	ADMINISTRATIVE_GROUP               Ospfv2OpaqueLsaAslaTypeEnum
	EXTENDED_ADMINISTRATIVE_GROUP      Ospfv2OpaqueLsaAslaTypeEnum
	TE_METRIC                          Ospfv2OpaqueLsaAslaTypeEnum
	MAXIMUM_LINK_BANDWIDTH             Ospfv2OpaqueLsaAslaTypeEnum
	GENERIC_METRIC                     Ospfv2OpaqueLsaAslaTypeEnum
}{
	RSVP_TE:                            Ospfv2OpaqueLsaAslaTypeEnum("rsvp_te"),
	SEGMENT_ROUTING_POLICY:             Ospfv2OpaqueLsaAslaTypeEnum("segment_routing_policy"),
	LOOP_FREE_ALTERNATE:                Ospfv2OpaqueLsaAslaTypeEnum("loop_free_alternate"),
	FLEXIBLE_ALGORITHM:                 Ospfv2OpaqueLsaAslaTypeEnum("flexible_algorithm"),
	SHARED_RISK_LINK_GROUP:             Ospfv2OpaqueLsaAslaTypeEnum("shared_risk_link_group"),
	UNIDIRECTIONAL_LINK_DELAY:          Ospfv2OpaqueLsaAslaTypeEnum("unidirectional_link_delay"),
	MIN_MAX_UNIDIRECTIONAL_LINK_DELAY:  Ospfv2OpaqueLsaAslaTypeEnum("min_max_unidirectional_link_delay"),
	UNIDIRECTIONAL_DELAY_VARIATION:     Ospfv2OpaqueLsaAslaTypeEnum("unidirectional_delay_variation"),
	UNIDIRECTIONAL_LINK_LOSS:           Ospfv2OpaqueLsaAslaTypeEnum("unidirectional_link_loss"),
	UNIDIRECTIONAL_RESIDUAL_BANDWIDTH:  Ospfv2OpaqueLsaAslaTypeEnum("unidirectional_residual_bandwidth"),
	UNIDIRECTIONAL_AVAILABLE_BANDWIDTH: Ospfv2OpaqueLsaAslaTypeEnum("unidirectional_available_bandwidth"),
	UNIDIRECTIONAL_UTILIZED_BANDWIDTH:  Ospfv2OpaqueLsaAslaTypeEnum("unidirectional_utilized_bandwidth"),
	ADMINISTRATIVE_GROUP:               Ospfv2OpaqueLsaAslaTypeEnum("administrative_group"),
	EXTENDED_ADMINISTRATIVE_GROUP:      Ospfv2OpaqueLsaAslaTypeEnum("extended_administrative_group"),
	TE_METRIC:                          Ospfv2OpaqueLsaAslaTypeEnum("te_metric"),
	MAXIMUM_LINK_BANDWIDTH:             Ospfv2OpaqueLsaAslaTypeEnum("maximum_link_bandwidth"),
	GENERIC_METRIC:                     Ospfv2OpaqueLsaAslaTypeEnum("generic_metric"),
}

func (obj *ospfv2OpaqueLsaAsla) Type() Ospfv2OpaqueLsaAslaTypeEnum {
	return Ospfv2OpaqueLsaAslaTypeEnum(obj.obj.Type.Enum().String())
}

// The Application-Specific Link Attribute Type field: either a Link Attribute
// Application Identifier bit (RFC 9479), or the type of a nested link-attribute
// sub-TLV reusing the OSPFv2 Extended Link TLV Sub-TLV type space (RFC 9492).
// Type returns a string
func (obj *ospfv2OpaqueLsaAsla) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *ospfv2OpaqueLsaAsla) SetType(value Ospfv2OpaqueLsaAslaTypeEnum) Ospfv2OpaqueLsaAsla {
	intValue, ok := otg.Ospfv2OpaqueLsaAsla_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2OpaqueLsaAslaTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2OpaqueLsaAsla_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// The Application-Specific Link Attribute Value field, returned as a lowercase hexadecimal string.
// Value returns a string
func (obj *ospfv2OpaqueLsaAsla) Value() string {

	return *obj.obj.Value

}

// The Application-Specific Link Attribute Value field, returned as a lowercase hexadecimal string.
// Value returns a string
func (obj *ospfv2OpaqueLsaAsla) HasValue() bool {
	return obj.obj.Value != nil
}

// The Application-Specific Link Attribute Value field, returned as a lowercase hexadecimal string.
// SetValue sets the string value in the Ospfv2OpaqueLsaAsla object
func (obj *ospfv2OpaqueLsaAsla) SetValue(value string) Ospfv2OpaqueLsaAsla {

	obj.obj.Value = &value
	return obj
}

func (obj *ospfv2OpaqueLsaAsla) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv2OpaqueLsaAsla) setDefault() {

}
