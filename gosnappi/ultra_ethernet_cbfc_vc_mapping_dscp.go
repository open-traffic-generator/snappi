package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetCbfcVcMappingDscp *****
type ultraEthernetCbfcVcMappingDscp struct {
	validation
	obj          *otg.UltraEthernetCbfcVcMappingDscp
	marshaller   marshalUltraEthernetCbfcVcMappingDscp
	unMarshaller unMarshalUltraEthernetCbfcVcMappingDscp
}

func NewUltraEthernetCbfcVcMappingDscp() UltraEthernetCbfcVcMappingDscp {
	obj := ultraEthernetCbfcVcMappingDscp{obj: &otg.UltraEthernetCbfcVcMappingDscp{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetCbfcVcMappingDscp) msg() *otg.UltraEthernetCbfcVcMappingDscp {
	return obj.obj
}

func (obj *ultraEthernetCbfcVcMappingDscp) setMsg(msg *otg.UltraEthernetCbfcVcMappingDscp) UltraEthernetCbfcVcMappingDscp {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetCbfcVcMappingDscp struct {
	obj *ultraEthernetCbfcVcMappingDscp
}

type marshalUltraEthernetCbfcVcMappingDscp interface {
	// ToProto marshals UltraEthernetCbfcVcMappingDscp to protobuf object *otg.UltraEthernetCbfcVcMappingDscp
	ToProto() (*otg.UltraEthernetCbfcVcMappingDscp, error)
	// ToPbText marshals UltraEthernetCbfcVcMappingDscp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetCbfcVcMappingDscp to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetCbfcVcMappingDscp to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetCbfcVcMappingDscp struct {
	obj *ultraEthernetCbfcVcMappingDscp
}

type unMarshalUltraEthernetCbfcVcMappingDscp interface {
	// FromProto unmarshals UltraEthernetCbfcVcMappingDscp from protobuf object *otg.UltraEthernetCbfcVcMappingDscp
	FromProto(msg *otg.UltraEthernetCbfcVcMappingDscp) (UltraEthernetCbfcVcMappingDscp, error)
	// FromPbText unmarshals UltraEthernetCbfcVcMappingDscp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetCbfcVcMappingDscp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetCbfcVcMappingDscp from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetCbfcVcMappingDscp) Marshal() marshalUltraEthernetCbfcVcMappingDscp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetCbfcVcMappingDscp{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetCbfcVcMappingDscp) Unmarshal() unMarshalUltraEthernetCbfcVcMappingDscp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetCbfcVcMappingDscp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetCbfcVcMappingDscp) ToProto() (*otg.UltraEthernetCbfcVcMappingDscp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetCbfcVcMappingDscp) FromProto(msg *otg.UltraEthernetCbfcVcMappingDscp) (UltraEthernetCbfcVcMappingDscp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetCbfcVcMappingDscp) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcMappingDscp) FromPbText(value string) error {
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

func (m *marshalultraEthernetCbfcVcMappingDscp) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcMappingDscp) FromYaml(value string) error {
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

func (m *marshalultraEthernetCbfcVcMappingDscp) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcMappingDscp) FromJson(value string) error {
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

func (obj *ultraEthernetCbfcVcMappingDscp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcVcMappingDscp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcVcMappingDscp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetCbfcVcMappingDscp) Clone() (UltraEthernetCbfcVcMappingDscp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetCbfcVcMappingDscp()
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

// UltraEthernetCbfcVcMappingDscp is iPv4/IPv6 DSCP values that classify packets into this virtual channel.
type UltraEthernetCbfcVcMappingDscp interface {
	Validation
	// msg marshals UltraEthernetCbfcVcMappingDscp to protobuf object *otg.UltraEthernetCbfcVcMappingDscp
	// and doesn't set defaults
	msg() *otg.UltraEthernetCbfcVcMappingDscp
	// setMsg unmarshals UltraEthernetCbfcVcMappingDscp from protobuf object *otg.UltraEthernetCbfcVcMappingDscp
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetCbfcVcMappingDscp) UltraEthernetCbfcVcMappingDscp
	// provides marshal interface
	Marshal() marshalUltraEthernetCbfcVcMappingDscp
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetCbfcVcMappingDscp
	// validate validates UltraEthernetCbfcVcMappingDscp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetCbfcVcMappingDscp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Values returns []uint32, set in UltraEthernetCbfcVcMappingDscp.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to UltraEthernetCbfcVcMappingDscp
	SetValues(value []uint32) UltraEthernetCbfcVcMappingDscp
}

// The DSCP values mapped to this VC. Valid values are 0 to 63.
// Values returns a []uint32
func (obj *ultraEthernetCbfcVcMappingDscp) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	return obj.obj.Values
}

// The DSCP values mapped to this VC. Valid values are 0 to 63.
// SetValues sets the []uint32 value in the UltraEthernetCbfcVcMappingDscp object
func (obj *ultraEthernetCbfcVcMappingDscp) SetValues(value []uint32) UltraEthernetCbfcVcMappingDscp {

	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

func (obj *ultraEthernetCbfcVcMappingDscp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 63 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= UltraEthernetCbfcVcMappingDscp.Values <= 63 but Got %d", item))
			}

		}

	}

}

func (obj *ultraEthernetCbfcVcMappingDscp) setDefault() {

}
