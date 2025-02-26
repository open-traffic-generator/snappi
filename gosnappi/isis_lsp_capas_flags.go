package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspCapasFlags *****
type isisLspCapasFlags struct {
	validation
	obj          *otg.IsisLspCapasFlags
	marshaller   marshalIsisLspCapasFlags
	unMarshaller unMarshalIsisLspCapasFlags
}

func NewIsisLspCapasFlags() IsisLspCapasFlags {
	obj := isisLspCapasFlags{obj: &otg.IsisLspCapasFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspCapasFlags) msg() *otg.IsisLspCapasFlags {
	return obj.obj
}

func (obj *isisLspCapasFlags) setMsg(msg *otg.IsisLspCapasFlags) IsisLspCapasFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspCapasFlags struct {
	obj *isisLspCapasFlags
}

type marshalIsisLspCapasFlags interface {
	// ToProto marshals IsisLspCapasFlags to protobuf object *otg.IsisLspCapasFlags
	ToProto() (*otg.IsisLspCapasFlags, error)
	// ToPbText marshals IsisLspCapasFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspCapasFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspCapasFlags to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspCapasFlags struct {
	obj *isisLspCapasFlags
}

type unMarshalIsisLspCapasFlags interface {
	// FromProto unmarshals IsisLspCapasFlags from protobuf object *otg.IsisLspCapasFlags
	FromProto(msg *otg.IsisLspCapasFlags) (IsisLspCapasFlags, error)
	// FromPbText unmarshals IsisLspCapasFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspCapasFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspCapasFlags from JSON text
	FromJson(value string) error
}

func (obj *isisLspCapasFlags) Marshal() marshalIsisLspCapasFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspCapasFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspCapasFlags) Unmarshal() unMarshalIsisLspCapasFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspCapasFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspCapasFlags) ToProto() (*otg.IsisLspCapasFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspCapasFlags) FromProto(msg *otg.IsisLspCapasFlags) (IsisLspCapasFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspCapasFlags) ToPbText() (string, error) {
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

func (m *unMarshalisisLspCapasFlags) FromPbText(value string) error {
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

func (m *marshalisisLspCapasFlags) ToYaml() (string, error) {
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

func (m *unMarshalisisLspCapasFlags) FromYaml(value string) error {
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

func (m *marshalisisLspCapasFlags) ToJson() (string, error) {
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

func (m *unMarshalisisLspCapasFlags) FromJson(value string) error {
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

func (obj *isisLspCapasFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspCapasFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspCapasFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspCapasFlags) Clone() (IsisLspCapasFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspCapasFlags()
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

// IsisLspCapasFlags is container for the configuration of IS-IS SR-CAPABILITY flags.
type IsisLspCapasFlags interface {
	Validation
	// msg marshals IsisLspCapasFlags to protobuf object *otg.IsisLspCapasFlags
	// and doesn't set defaults
	msg() *otg.IsisLspCapasFlags
	// setMsg unmarshals IsisLspCapasFlags from protobuf object *otg.IsisLspCapasFlags
	// and doesn't set defaults
	setMsg(*otg.IsisLspCapasFlags) IsisLspCapasFlags
	// provides marshal interface
	Marshal() marshalIsisLspCapasFlags
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspCapasFlags
	// validate validates IsisLspCapasFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspCapasFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Mpls returns bool, set in IsisLspCapasFlags.
	Ipv4Mpls() bool
	// SetIpv4Mpls assigns bool provided by user to IsisLspCapasFlags
	SetIpv4Mpls(value bool) IsisLspCapasFlags
	// HasIpv4Mpls checks if Ipv4Mpls has been set in IsisLspCapasFlags
	HasIpv4Mpls() bool
	// Ipv6Mpls returns bool, set in IsisLspCapasFlags.
	Ipv6Mpls() bool
	// SetIpv6Mpls assigns bool provided by user to IsisLspCapasFlags
	SetIpv6Mpls(value bool) IsisLspCapasFlags
	// HasIpv6Mpls checks if Ipv6Mpls has been set in IsisLspCapasFlags
	HasIpv6Mpls() bool
}

// I-Flag for the MPLS IPv4 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv4 packets on all interfaces.
// Ipv4Mpls returns a bool
func (obj *isisLspCapasFlags) Ipv4Mpls() bool {

	return *obj.obj.Ipv4Mpls

}

// I-Flag for the MPLS IPv4 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv4 packets on all interfaces.
// Ipv4Mpls returns a bool
func (obj *isisLspCapasFlags) HasIpv4Mpls() bool {
	return obj.obj.Ipv4Mpls != nil
}

// I-Flag for the MPLS IPv4 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv4 packets on all interfaces.
// SetIpv4Mpls sets the bool value in the IsisLspCapasFlags object
func (obj *isisLspCapasFlags) SetIpv4Mpls(value bool) IsisLspCapasFlags {

	obj.obj.Ipv4Mpls = &value
	return obj
}

// V-Flag for the MPLS IPv6 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv6 packets on all interfaces.
// Ipv6Mpls returns a bool
func (obj *isisLspCapasFlags) Ipv6Mpls() bool {

	return *obj.obj.Ipv6Mpls

}

// V-Flag for the MPLS IPv6 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv6 packets on all interfaces.
// Ipv6Mpls returns a bool
func (obj *isisLspCapasFlags) HasIpv6Mpls() bool {
	return obj.obj.Ipv6Mpls != nil
}

// V-Flag for the MPLS IPv6 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv6 packets on all interfaces.
// SetIpv6Mpls sets the bool value in the IsisLspCapasFlags object
func (obj *isisLspCapasFlags) SetIpv6Mpls(value bool) IsisLspCapasFlags {

	obj.obj.Ipv6Mpls = &value
	return obj
}

func (obj *isisLspCapasFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisLspCapasFlags) setDefault() {

}
