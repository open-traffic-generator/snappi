package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetCbfcVcMappingVlanPcpDei *****
type ultraEthernetCbfcVcMappingVlanPcpDei struct {
	validation
	obj          *otg.UltraEthernetCbfcVcMappingVlanPcpDei
	marshaller   marshalUltraEthernetCbfcVcMappingVlanPcpDei
	unMarshaller unMarshalUltraEthernetCbfcVcMappingVlanPcpDei
}

func NewUltraEthernetCbfcVcMappingVlanPcpDei() UltraEthernetCbfcVcMappingVlanPcpDei {
	obj := ultraEthernetCbfcVcMappingVlanPcpDei{obj: &otg.UltraEthernetCbfcVcMappingVlanPcpDei{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetCbfcVcMappingVlanPcpDei) msg() *otg.UltraEthernetCbfcVcMappingVlanPcpDei {
	return obj.obj
}

func (obj *ultraEthernetCbfcVcMappingVlanPcpDei) setMsg(msg *otg.UltraEthernetCbfcVcMappingVlanPcpDei) UltraEthernetCbfcVcMappingVlanPcpDei {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetCbfcVcMappingVlanPcpDei struct {
	obj *ultraEthernetCbfcVcMappingVlanPcpDei
}

type marshalUltraEthernetCbfcVcMappingVlanPcpDei interface {
	// ToProto marshals UltraEthernetCbfcVcMappingVlanPcpDei to protobuf object *otg.UltraEthernetCbfcVcMappingVlanPcpDei
	ToProto() (*otg.UltraEthernetCbfcVcMappingVlanPcpDei, error)
	// ToPbText marshals UltraEthernetCbfcVcMappingVlanPcpDei to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetCbfcVcMappingVlanPcpDei to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetCbfcVcMappingVlanPcpDei to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetCbfcVcMappingVlanPcpDei struct {
	obj *ultraEthernetCbfcVcMappingVlanPcpDei
}

type unMarshalUltraEthernetCbfcVcMappingVlanPcpDei interface {
	// FromProto unmarshals UltraEthernetCbfcVcMappingVlanPcpDei from protobuf object *otg.UltraEthernetCbfcVcMappingVlanPcpDei
	FromProto(msg *otg.UltraEthernetCbfcVcMappingVlanPcpDei) (UltraEthernetCbfcVcMappingVlanPcpDei, error)
	// FromPbText unmarshals UltraEthernetCbfcVcMappingVlanPcpDei from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetCbfcVcMappingVlanPcpDei from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetCbfcVcMappingVlanPcpDei from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetCbfcVcMappingVlanPcpDei) Marshal() marshalUltraEthernetCbfcVcMappingVlanPcpDei {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetCbfcVcMappingVlanPcpDei{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetCbfcVcMappingVlanPcpDei) Unmarshal() unMarshalUltraEthernetCbfcVcMappingVlanPcpDei {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetCbfcVcMappingVlanPcpDei{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetCbfcVcMappingVlanPcpDei) ToProto() (*otg.UltraEthernetCbfcVcMappingVlanPcpDei, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetCbfcVcMappingVlanPcpDei) FromProto(msg *otg.UltraEthernetCbfcVcMappingVlanPcpDei) (UltraEthernetCbfcVcMappingVlanPcpDei, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetCbfcVcMappingVlanPcpDei) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcMappingVlanPcpDei) FromPbText(value string) error {
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

func (m *marshalultraEthernetCbfcVcMappingVlanPcpDei) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcMappingVlanPcpDei) FromYaml(value string) error {
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

func (m *marshalultraEthernetCbfcVcMappingVlanPcpDei) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcMappingVlanPcpDei) FromJson(value string) error {
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

func (obj *ultraEthernetCbfcVcMappingVlanPcpDei) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcVcMappingVlanPcpDei) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcVcMappingVlanPcpDei) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetCbfcVcMappingVlanPcpDei) Clone() (UltraEthernetCbfcVcMappingVlanPcpDei, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetCbfcVcMappingVlanPcpDei()
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

// UltraEthernetCbfcVcMappingVlanPcpDei is vLAN PCP/DEI values that classify packets into this virtual channel.
type UltraEthernetCbfcVcMappingVlanPcpDei interface {
	Validation
	// msg marshals UltraEthernetCbfcVcMappingVlanPcpDei to protobuf object *otg.UltraEthernetCbfcVcMappingVlanPcpDei
	// and doesn't set defaults
	msg() *otg.UltraEthernetCbfcVcMappingVlanPcpDei
	// setMsg unmarshals UltraEthernetCbfcVcMappingVlanPcpDei from protobuf object *otg.UltraEthernetCbfcVcMappingVlanPcpDei
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetCbfcVcMappingVlanPcpDei) UltraEthernetCbfcVcMappingVlanPcpDei
	// provides marshal interface
	Marshal() marshalUltraEthernetCbfcVcMappingVlanPcpDei
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetCbfcVcMappingVlanPcpDei
	// validate validates UltraEthernetCbfcVcMappingVlanPcpDei
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetCbfcVcMappingVlanPcpDei, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PcpDei0 returns []uint32, set in UltraEthernetCbfcVcMappingVlanPcpDei.
	PcpDei0() []uint32
	// SetPcpDei0 assigns []uint32 provided by user to UltraEthernetCbfcVcMappingVlanPcpDei
	SetPcpDei0(value []uint32) UltraEthernetCbfcVcMappingVlanPcpDei
	// PcpDei1 returns []uint32, set in UltraEthernetCbfcVcMappingVlanPcpDei.
	PcpDei1() []uint32
	// SetPcpDei1 assigns []uint32 provided by user to UltraEthernetCbfcVcMappingVlanPcpDei
	SetPcpDei1(value []uint32) UltraEthernetCbfcVcMappingVlanPcpDei
}

// The VLAN PCP values, with DEI = 0, mapped to this VC. Valid values are 0 to 7.
// PcpDei0 returns a []uint32
func (obj *ultraEthernetCbfcVcMappingVlanPcpDei) PcpDei0() []uint32 {
	if obj.obj.PcpDei0 == nil {
		obj.obj.PcpDei0 = make([]uint32, 0)
	}
	return obj.obj.PcpDei0
}

// The VLAN PCP values, with DEI = 0, mapped to this VC. Valid values are 0 to 7.
// SetPcpDei0 sets the []uint32 value in the UltraEthernetCbfcVcMappingVlanPcpDei object
func (obj *ultraEthernetCbfcVcMappingVlanPcpDei) SetPcpDei0(value []uint32) UltraEthernetCbfcVcMappingVlanPcpDei {

	if obj.obj.PcpDei0 == nil {
		obj.obj.PcpDei0 = make([]uint32, 0)
	}
	obj.obj.PcpDei0 = value

	return obj
}

// The VLAN PCP values, with DEI = 1, mapped to this VC. Valid values are 0 to 7.
// PcpDei1 returns a []uint32
func (obj *ultraEthernetCbfcVcMappingVlanPcpDei) PcpDei1() []uint32 {
	if obj.obj.PcpDei1 == nil {
		obj.obj.PcpDei1 = make([]uint32, 0)
	}
	return obj.obj.PcpDei1
}

// The VLAN PCP values, with DEI = 1, mapped to this VC. Valid values are 0 to 7.
// SetPcpDei1 sets the []uint32 value in the UltraEthernetCbfcVcMappingVlanPcpDei object
func (obj *ultraEthernetCbfcVcMappingVlanPcpDei) SetPcpDei1(value []uint32) UltraEthernetCbfcVcMappingVlanPcpDei {

	if obj.obj.PcpDei1 == nil {
		obj.obj.PcpDei1 = make([]uint32, 0)
	}
	obj.obj.PcpDei1 = value

	return obj
}

func (obj *ultraEthernetCbfcVcMappingVlanPcpDei) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PcpDei0 != nil {

		for _, item := range obj.obj.PcpDei0 {
			if item > 7 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= UltraEthernetCbfcVcMappingVlanPcpDei.PcpDei0 <= 7 but Got %d", item))
			}

		}

	}

	if obj.obj.PcpDei1 != nil {

		for _, item := range obj.obj.PcpDei1 {
			if item > 7 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= UltraEthernetCbfcVcMappingVlanPcpDei.PcpDei1 <= 7 but Got %d", item))
			}

		}

	}

}

func (obj *ultraEthernetCbfcVcMappingVlanPcpDei) setDefault() {

}
