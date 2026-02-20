package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortMacsec *****
type lagPortMacsec struct {
	validation
	obj                    *otg.LagPortMacsec
	marshaller             marshalLagPortMacsec
	unMarshaller           unMarshalLagPortMacsec
	secureEntityHolder     SecureEntity
	excludeProtocolsHolder LagPortMacsecExcludeProtocols
}

func NewLagPortMacsec() LagPortMacsec {
	obj := lagPortMacsec{obj: &otg.LagPortMacsec{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortMacsec) msg() *otg.LagPortMacsec {
	return obj.obj
}

func (obj *lagPortMacsec) setMsg(msg *otg.LagPortMacsec) LagPortMacsec {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortMacsec struct {
	obj *lagPortMacsec
}

type marshalLagPortMacsec interface {
	// ToProto marshals LagPortMacsec to protobuf object *otg.LagPortMacsec
	ToProto() (*otg.LagPortMacsec, error)
	// ToPbText marshals LagPortMacsec to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortMacsec to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortMacsec to JSON text
	ToJson() (string, error)
}

type unMarshallagPortMacsec struct {
	obj *lagPortMacsec
}

type unMarshalLagPortMacsec interface {
	// FromProto unmarshals LagPortMacsec from protobuf object *otg.LagPortMacsec
	FromProto(msg *otg.LagPortMacsec) (LagPortMacsec, error)
	// FromPbText unmarshals LagPortMacsec from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortMacsec from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortMacsec from JSON text
	FromJson(value string) error
}

func (obj *lagPortMacsec) Marshal() marshalLagPortMacsec {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortMacsec{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortMacsec) Unmarshal() unMarshalLagPortMacsec {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortMacsec{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortMacsec) ToProto() (*otg.LagPortMacsec, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortMacsec) FromProto(msg *otg.LagPortMacsec) (LagPortMacsec, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortMacsec) ToPbText() (string, error) {
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

func (m *unMarshallagPortMacsec) FromPbText(value string) error {
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

func (m *marshallagPortMacsec) ToYaml() (string, error) {
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

func (m *unMarshallagPortMacsec) FromYaml(value string) error {
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

func (m *marshallagPortMacsec) ToJson() (string, error) {
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

func (m *unMarshallagPortMacsec) FromJson(value string) error {
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

func (obj *lagPortMacsec) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortMacsec) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortMacsec) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortMacsec) Clone() (LagPortMacsec, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortMacsec()
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

func (obj *lagPortMacsec) setNil() {
	obj.secureEntityHolder = nil
	obj.excludeProtocolsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LagPortMacsec is configuration of MACsec per LAG member port.
type LagPortMacsec interface {
	Validation
	// msg marshals LagPortMacsec to protobuf object *otg.LagPortMacsec
	// and doesn't set defaults
	msg() *otg.LagPortMacsec
	// setMsg unmarshals LagPortMacsec from protobuf object *otg.LagPortMacsec
	// and doesn't set defaults
	setMsg(*otg.LagPortMacsec) LagPortMacsec
	// provides marshal interface
	Marshal() marshalLagPortMacsec
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortMacsec
	// validate validates LagPortMacsec
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortMacsec, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SecureEntity returns SecureEntity, set in LagPortMacsec.
	// SecureEntity is configuration of a Secure Entity (SecY).
	SecureEntity() SecureEntity
	// SetSecureEntity assigns SecureEntity provided by user to LagPortMacsec.
	// SecureEntity is configuration of a Secure Entity (SecY).
	SetSecureEntity(value SecureEntity) LagPortMacsec
	// ExcludeProtocols returns LagPortMacsecExcludeProtocols, set in LagPortMacsec.
	// LagPortMacsecExcludeProtocols is protocols excluded from MACsec encapsulation at Tx.
	ExcludeProtocols() LagPortMacsecExcludeProtocols
	// SetExcludeProtocols assigns LagPortMacsecExcludeProtocols provided by user to LagPortMacsec.
	// LagPortMacsecExcludeProtocols is protocols excluded from MACsec encapsulation at Tx.
	SetExcludeProtocols(value LagPortMacsecExcludeProtocols) LagPortMacsec
	// HasExcludeProtocols checks if ExcludeProtocols has been set in LagPortMacsec
	HasExcludeProtocols() bool
	setNil()
}

// This contains the properties of Secure Entity (SecY) per LAG member port.
// SecureEntity returns a SecureEntity
func (obj *lagPortMacsec) SecureEntity() SecureEntity {
	if obj.obj.SecureEntity == nil {
		obj.obj.SecureEntity = NewSecureEntity().msg()
	}
	if obj.secureEntityHolder == nil {
		obj.secureEntityHolder = &secureEntity{obj: obj.obj.SecureEntity}
	}
	return obj.secureEntityHolder
}

// This contains the properties of Secure Entity (SecY) per LAG member port.
// SetSecureEntity sets the SecureEntity value in the LagPortMacsec object
func (obj *lagPortMacsec) SetSecureEntity(value SecureEntity) LagPortMacsec {

	obj.secureEntityHolder = nil
	obj.obj.SecureEntity = value.msg()

	return obj
}

// Protocols excluded from MACsec encapsulation.
// ExcludeProtocols returns a LagPortMacsecExcludeProtocols
func (obj *lagPortMacsec) ExcludeProtocols() LagPortMacsecExcludeProtocols {
	if obj.obj.ExcludeProtocols == nil {
		obj.obj.ExcludeProtocols = NewLagPortMacsecExcludeProtocols().msg()
	}
	if obj.excludeProtocolsHolder == nil {
		obj.excludeProtocolsHolder = &lagPortMacsecExcludeProtocols{obj: obj.obj.ExcludeProtocols}
	}
	return obj.excludeProtocolsHolder
}

// Protocols excluded from MACsec encapsulation.
// ExcludeProtocols returns a LagPortMacsecExcludeProtocols
func (obj *lagPortMacsec) HasExcludeProtocols() bool {
	return obj.obj.ExcludeProtocols != nil
}

// Protocols excluded from MACsec encapsulation.
// SetExcludeProtocols sets the LagPortMacsecExcludeProtocols value in the LagPortMacsec object
func (obj *lagPortMacsec) SetExcludeProtocols(value LagPortMacsecExcludeProtocols) LagPortMacsec {

	obj.excludeProtocolsHolder = nil
	obj.obj.ExcludeProtocols = value.msg()

	return obj
}

func (obj *lagPortMacsec) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// SecureEntity is required
	if obj.obj.SecureEntity == nil {
		vObj.validationErrors = append(vObj.validationErrors, "SecureEntity is required field on interface LagPortMacsec")
	}

	if obj.obj.SecureEntity != nil {

		obj.SecureEntity().validateObj(vObj, set_default)
	}

	if obj.obj.ExcludeProtocols != nil {

		obj.ExcludeProtocols().validateObj(vObj, set_default)
	}

}

func (obj *lagPortMacsec) setDefault() {

}
