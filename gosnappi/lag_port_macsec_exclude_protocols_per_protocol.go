package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortMacsecExcludeProtocolsPerProtocol *****
type lagPortMacsecExcludeProtocolsPerProtocol struct {
	validation
	obj          *otg.LagPortMacsecExcludeProtocolsPerProtocol
	marshaller   marshalLagPortMacsecExcludeProtocolsPerProtocol
	unMarshaller unMarshalLagPortMacsecExcludeProtocolsPerProtocol
}

func NewLagPortMacsecExcludeProtocolsPerProtocol() LagPortMacsecExcludeProtocolsPerProtocol {
	obj := lagPortMacsecExcludeProtocolsPerProtocol{obj: &otg.LagPortMacsecExcludeProtocolsPerProtocol{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortMacsecExcludeProtocolsPerProtocol) msg() *otg.LagPortMacsecExcludeProtocolsPerProtocol {
	return obj.obj
}

func (obj *lagPortMacsecExcludeProtocolsPerProtocol) setMsg(msg *otg.LagPortMacsecExcludeProtocolsPerProtocol) LagPortMacsecExcludeProtocolsPerProtocol {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortMacsecExcludeProtocolsPerProtocol struct {
	obj *lagPortMacsecExcludeProtocolsPerProtocol
}

type marshalLagPortMacsecExcludeProtocolsPerProtocol interface {
	// ToProto marshals LagPortMacsecExcludeProtocolsPerProtocol to protobuf object *otg.LagPortMacsecExcludeProtocolsPerProtocol
	ToProto() (*otg.LagPortMacsecExcludeProtocolsPerProtocol, error)
	// ToPbText marshals LagPortMacsecExcludeProtocolsPerProtocol to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortMacsecExcludeProtocolsPerProtocol to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortMacsecExcludeProtocolsPerProtocol to JSON text
	ToJson() (string, error)
}

type unMarshallagPortMacsecExcludeProtocolsPerProtocol struct {
	obj *lagPortMacsecExcludeProtocolsPerProtocol
}

type unMarshalLagPortMacsecExcludeProtocolsPerProtocol interface {
	// FromProto unmarshals LagPortMacsecExcludeProtocolsPerProtocol from protobuf object *otg.LagPortMacsecExcludeProtocolsPerProtocol
	FromProto(msg *otg.LagPortMacsecExcludeProtocolsPerProtocol) (LagPortMacsecExcludeProtocolsPerProtocol, error)
	// FromPbText unmarshals LagPortMacsecExcludeProtocolsPerProtocol from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortMacsecExcludeProtocolsPerProtocol from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortMacsecExcludeProtocolsPerProtocol from JSON text
	FromJson(value string) error
}

func (obj *lagPortMacsecExcludeProtocolsPerProtocol) Marshal() marshalLagPortMacsecExcludeProtocolsPerProtocol {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortMacsecExcludeProtocolsPerProtocol{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortMacsecExcludeProtocolsPerProtocol) Unmarshal() unMarshalLagPortMacsecExcludeProtocolsPerProtocol {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortMacsecExcludeProtocolsPerProtocol{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortMacsecExcludeProtocolsPerProtocol) ToProto() (*otg.LagPortMacsecExcludeProtocolsPerProtocol, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortMacsecExcludeProtocolsPerProtocol) FromProto(msg *otg.LagPortMacsecExcludeProtocolsPerProtocol) (LagPortMacsecExcludeProtocolsPerProtocol, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortMacsecExcludeProtocolsPerProtocol) ToPbText() (string, error) {
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

func (m *unMarshallagPortMacsecExcludeProtocolsPerProtocol) FromPbText(value string) error {
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

func (m *marshallagPortMacsecExcludeProtocolsPerProtocol) ToYaml() (string, error) {
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

func (m *unMarshallagPortMacsecExcludeProtocolsPerProtocol) FromYaml(value string) error {
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

func (m *marshallagPortMacsecExcludeProtocolsPerProtocol) ToJson() (string, error) {
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

func (m *unMarshallagPortMacsecExcludeProtocolsPerProtocol) FromJson(value string) error {
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

func (obj *lagPortMacsecExcludeProtocolsPerProtocol) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortMacsecExcludeProtocolsPerProtocol) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortMacsecExcludeProtocolsPerProtocol) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortMacsecExcludeProtocolsPerProtocol) Clone() (LagPortMacsecExcludeProtocolsPerProtocol, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortMacsecExcludeProtocolsPerProtocol()
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

// LagPortMacsecExcludeProtocolsPerProtocol is per protocol choice to exclude from MACsec encapsulation at Tx.
type LagPortMacsecExcludeProtocolsPerProtocol interface {
	Validation
	// msg marshals LagPortMacsecExcludeProtocolsPerProtocol to protobuf object *otg.LagPortMacsecExcludeProtocolsPerProtocol
	// and doesn't set defaults
	msg() *otg.LagPortMacsecExcludeProtocolsPerProtocol
	// setMsg unmarshals LagPortMacsecExcludeProtocolsPerProtocol from protobuf object *otg.LagPortMacsecExcludeProtocolsPerProtocol
	// and doesn't set defaults
	setMsg(*otg.LagPortMacsecExcludeProtocolsPerProtocol) LagPortMacsecExcludeProtocolsPerProtocol
	// provides marshal interface
	Marshal() marshalLagPortMacsecExcludeProtocolsPerProtocol
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortMacsecExcludeProtocolsPerProtocol
	// validate validates LagPortMacsecExcludeProtocolsPerProtocol
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortMacsecExcludeProtocolsPerProtocol, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Lacp returns bool, set in LagPortMacsecExcludeProtocolsPerProtocol.
	Lacp() bool
	// SetLacp assigns bool provided by user to LagPortMacsecExcludeProtocolsPerProtocol
	SetLacp(value bool) LagPortMacsecExcludeProtocolsPerProtocol
	// HasLacp checks if Lacp has been set in LagPortMacsecExcludeProtocolsPerProtocol
	HasLacp() bool
}

// Determines whether LACP should be excluded from MACsec encapsulation or not.
// Lacp returns a bool
func (obj *lagPortMacsecExcludeProtocolsPerProtocol) Lacp() bool {

	return *obj.obj.Lacp

}

// Determines whether LACP should be excluded from MACsec encapsulation or not.
// Lacp returns a bool
func (obj *lagPortMacsecExcludeProtocolsPerProtocol) HasLacp() bool {
	return obj.obj.Lacp != nil
}

// Determines whether LACP should be excluded from MACsec encapsulation or not.
// SetLacp sets the bool value in the LagPortMacsecExcludeProtocolsPerProtocol object
func (obj *lagPortMacsecExcludeProtocolsPerProtocol) SetLacp(value bool) LagPortMacsecExcludeProtocolsPerProtocol {

	obj.obj.Lacp = &value
	return obj
}

func (obj *lagPortMacsecExcludeProtocolsPerProtocol) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lagPortMacsecExcludeProtocolsPerProtocol) setDefault() {
	if obj.obj.Lacp == nil {
		obj.SetLacp(true)
	}

}
