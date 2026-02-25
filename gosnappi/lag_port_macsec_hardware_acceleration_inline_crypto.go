package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortMacsecHardwareAccelerationInlineCrypto *****
type lagPortMacsecHardwareAccelerationInlineCrypto struct {
	validation
	obj          *otg.LagPortMacsecHardwareAccelerationInlineCrypto
	marshaller   marshalLagPortMacsecHardwareAccelerationInlineCrypto
	unMarshaller unMarshalLagPortMacsecHardwareAccelerationInlineCrypto
}

func NewLagPortMacsecHardwareAccelerationInlineCrypto() LagPortMacsecHardwareAccelerationInlineCrypto {
	obj := lagPortMacsecHardwareAccelerationInlineCrypto{obj: &otg.LagPortMacsecHardwareAccelerationInlineCrypto{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortMacsecHardwareAccelerationInlineCrypto) msg() *otg.LagPortMacsecHardwareAccelerationInlineCrypto {
	return obj.obj
}

func (obj *lagPortMacsecHardwareAccelerationInlineCrypto) setMsg(msg *otg.LagPortMacsecHardwareAccelerationInlineCrypto) LagPortMacsecHardwareAccelerationInlineCrypto {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortMacsecHardwareAccelerationInlineCrypto struct {
	obj *lagPortMacsecHardwareAccelerationInlineCrypto
}

type marshalLagPortMacsecHardwareAccelerationInlineCrypto interface {
	// ToProto marshals LagPortMacsecHardwareAccelerationInlineCrypto to protobuf object *otg.LagPortMacsecHardwareAccelerationInlineCrypto
	ToProto() (*otg.LagPortMacsecHardwareAccelerationInlineCrypto, error)
	// ToPbText marshals LagPortMacsecHardwareAccelerationInlineCrypto to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortMacsecHardwareAccelerationInlineCrypto to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortMacsecHardwareAccelerationInlineCrypto to JSON text
	ToJson() (string, error)
}

type unMarshallagPortMacsecHardwareAccelerationInlineCrypto struct {
	obj *lagPortMacsecHardwareAccelerationInlineCrypto
}

type unMarshalLagPortMacsecHardwareAccelerationInlineCrypto interface {
	// FromProto unmarshals LagPortMacsecHardwareAccelerationInlineCrypto from protobuf object *otg.LagPortMacsecHardwareAccelerationInlineCrypto
	FromProto(msg *otg.LagPortMacsecHardwareAccelerationInlineCrypto) (LagPortMacsecHardwareAccelerationInlineCrypto, error)
	// FromPbText unmarshals LagPortMacsecHardwareAccelerationInlineCrypto from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortMacsecHardwareAccelerationInlineCrypto from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortMacsecHardwareAccelerationInlineCrypto from JSON text
	FromJson(value string) error
}

func (obj *lagPortMacsecHardwareAccelerationInlineCrypto) Marshal() marshalLagPortMacsecHardwareAccelerationInlineCrypto {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortMacsecHardwareAccelerationInlineCrypto{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortMacsecHardwareAccelerationInlineCrypto) Unmarshal() unMarshalLagPortMacsecHardwareAccelerationInlineCrypto {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortMacsecHardwareAccelerationInlineCrypto{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortMacsecHardwareAccelerationInlineCrypto) ToProto() (*otg.LagPortMacsecHardwareAccelerationInlineCrypto, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortMacsecHardwareAccelerationInlineCrypto) FromProto(msg *otg.LagPortMacsecHardwareAccelerationInlineCrypto) (LagPortMacsecHardwareAccelerationInlineCrypto, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortMacsecHardwareAccelerationInlineCrypto) ToPbText() (string, error) {
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

func (m *unMarshallagPortMacsecHardwareAccelerationInlineCrypto) FromPbText(value string) error {
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

func (m *marshallagPortMacsecHardwareAccelerationInlineCrypto) ToYaml() (string, error) {
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

func (m *unMarshallagPortMacsecHardwareAccelerationInlineCrypto) FromYaml(value string) error {
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

func (m *marshallagPortMacsecHardwareAccelerationInlineCrypto) ToJson() (string, error) {
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

func (m *unMarshallagPortMacsecHardwareAccelerationInlineCrypto) FromJson(value string) error {
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

func (obj *lagPortMacsecHardwareAccelerationInlineCrypto) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortMacsecHardwareAccelerationInlineCrypto) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortMacsecHardwareAccelerationInlineCrypto) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortMacsecHardwareAccelerationInlineCrypto) Clone() (LagPortMacsecHardwareAccelerationInlineCrypto, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortMacsecHardwareAccelerationInlineCrypto()
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

// LagPortMacsecHardwareAccelerationInlineCrypto is inline cryto engine configuration. Encryption/ decryption are offloaded to hardware. Also dynamic fields e.g. packet number(PN) and integrity check value(ICV) are updated in MACsec header on transmit.
type LagPortMacsecHardwareAccelerationInlineCrypto interface {
	Validation
	// msg marshals LagPortMacsecHardwareAccelerationInlineCrypto to protobuf object *otg.LagPortMacsecHardwareAccelerationInlineCrypto
	// and doesn't set defaults
	msg() *otg.LagPortMacsecHardwareAccelerationInlineCrypto
	// setMsg unmarshals LagPortMacsecHardwareAccelerationInlineCrypto from protobuf object *otg.LagPortMacsecHardwareAccelerationInlineCrypto
	// and doesn't set defaults
	setMsg(*otg.LagPortMacsecHardwareAccelerationInlineCrypto) LagPortMacsecHardwareAccelerationInlineCrypto
	// provides marshal interface
	Marshal() marshalLagPortMacsecHardwareAccelerationInlineCrypto
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortMacsecHardwareAccelerationInlineCrypto
	// validate validates LagPortMacsecHardwareAccelerationInlineCrypto
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortMacsecHardwareAccelerationInlineCrypto, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RxScIdentifyingField returns LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum, set in LagPortMacsecHardwareAccelerationInlineCrypto
	RxScIdentifyingField() LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	// SetRxScIdentifyingField assigns LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum provided by user to LagPortMacsecHardwareAccelerationInlineCrypto
	SetRxScIdentifyingField(value LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum) LagPortMacsecHardwareAccelerationInlineCrypto
	// HasRxScIdentifyingField checks if RxScIdentifyingField has been set in LagPortMacsecHardwareAccelerationInlineCrypto
	HasRxScIdentifyingField() bool
}

type LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum string

// Enum of RxScIdentifyingField on LagPortMacsecHardwareAccelerationInlineCrypto
var LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingField = struct {
	SOURCE_MAC   LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	SCI_SYTEM_ID LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	SCI_PORT_ID  LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
}{
	SOURCE_MAC:   LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("source_mac"),
	SCI_SYTEM_ID: LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("sci_sytem_id"),
	SCI_PORT_ID:  LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("sci_port_id"),
}

func (obj *lagPortMacsecHardwareAccelerationInlineCrypto) RxScIdentifyingField() LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum {
	return LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum(obj.obj.RxScIdentifyingField.Enum().String())
}

// The field based on which secure channel(SC) will be identified by the receiving port. Supported fields are:- - 1) source MAC - identify SC based on source MAC field. 2) SCI system ID - identify SC based on SCI system ID field. 3) SCI port ID - identify based on SCI port ID field.
// RxScIdentifyingField returns a string
func (obj *lagPortMacsecHardwareAccelerationInlineCrypto) HasRxScIdentifyingField() bool {
	return obj.obj.RxScIdentifyingField != nil
}

func (obj *lagPortMacsecHardwareAccelerationInlineCrypto) SetRxScIdentifyingField(value LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum) LagPortMacsecHardwareAccelerationInlineCrypto {
	intValue, ok := otg.LagPortMacsecHardwareAccelerationInlineCrypto_RxScIdentifyingField_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum", string(value)))
		return obj
	}
	enumValue := otg.LagPortMacsecHardwareAccelerationInlineCrypto_RxScIdentifyingField_Enum(intValue)
	obj.obj.RxScIdentifyingField = &enumValue

	return obj
}

func (obj *lagPortMacsecHardwareAccelerationInlineCrypto) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lagPortMacsecHardwareAccelerationInlineCrypto) setDefault() {
	if obj.obj.RxScIdentifyingField == nil {
		obj.SetRxScIdentifyingField(LagPortMacsecHardwareAccelerationInlineCryptoRxScIdentifyingField.SOURCE_MAC)

	}

}
