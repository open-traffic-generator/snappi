package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisBasic *****
type isisBasic struct {
	validation
	obj          *otg.IsisBasic
	marshaller   marshalIsisBasic
	unMarshaller unMarshalIsisBasic
}

func NewIsisBasic() IsisBasic {
	obj := isisBasic{obj: &otg.IsisBasic{}}
	obj.setDefault()
	return &obj
}

func (obj *isisBasic) msg() *otg.IsisBasic {
	return obj.obj
}

func (obj *isisBasic) setMsg(msg *otg.IsisBasic) IsisBasic {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisBasic struct {
	obj *isisBasic
}

type marshalIsisBasic interface {
	// ToProto marshals IsisBasic to protobuf object *otg.IsisBasic
	ToProto() (*otg.IsisBasic, error)
	// ToPbText marshals IsisBasic to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisBasic to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisBasic to JSON text
	ToJson() (string, error)
}

type unMarshalisisBasic struct {
	obj *isisBasic
}

type unMarshalIsisBasic interface {
	// FromProto unmarshals IsisBasic from protobuf object *otg.IsisBasic
	FromProto(msg *otg.IsisBasic) (IsisBasic, error)
	// FromPbText unmarshals IsisBasic from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisBasic from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisBasic from JSON text
	FromJson(value string) error
}

func (obj *isisBasic) Marshal() marshalIsisBasic {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisBasic{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisBasic) Unmarshal() unMarshalIsisBasic {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisBasic{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisBasic) ToProto() (*otg.IsisBasic, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisBasic) FromProto(msg *otg.IsisBasic) (IsisBasic, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisBasic) ToPbText() (string, error) {
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

func (m *unMarshalisisBasic) FromPbText(value string) error {
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

func (m *marshalisisBasic) ToYaml() (string, error) {
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

func (m *unMarshalisisBasic) FromYaml(value string) error {
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

func (m *marshalisisBasic) ToJson() (string, error) {
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

func (m *unMarshalisisBasic) FromJson(value string) error {
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

func (obj *isisBasic) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisBasic) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisBasic) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisBasic) Clone() (IsisBasic, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisBasic()
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

// IsisBasic is this contains ISIS router basic properties.
type IsisBasic interface {
	Validation
	// msg marshals IsisBasic to protobuf object *otg.IsisBasic
	// and doesn't set defaults
	msg() *otg.IsisBasic
	// setMsg unmarshals IsisBasic from protobuf object *otg.IsisBasic
	// and doesn't set defaults
	setMsg(*otg.IsisBasic) IsisBasic
	// provides marshal interface
	Marshal() marshalIsisBasic
	// provides unmarshal interface
	Unmarshal() unMarshalIsisBasic
	// validate validates IsisBasic
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisBasic, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4TeRouterId returns string, set in IsisBasic.
	Ipv4TeRouterId() string
	// SetIpv4TeRouterId assigns string provided by user to IsisBasic
	SetIpv4TeRouterId(value string) IsisBasic
	// HasIpv4TeRouterId checks if Ipv4TeRouterId has been set in IsisBasic
	HasIpv4TeRouterId() bool
	// Hostname returns string, set in IsisBasic.
	Hostname() string
	// SetHostname assigns string provided by user to IsisBasic
	SetHostname(value string) IsisBasic
	// HasHostname checks if Hostname has been set in IsisBasic
	HasHostname() bool
	// EnableWideMetric returns bool, set in IsisBasic.
	EnableWideMetric() bool
	// SetEnableWideMetric assigns bool provided by user to IsisBasic
	SetEnableWideMetric(value bool) IsisBasic
	// HasEnableWideMetric checks if EnableWideMetric has been set in IsisBasic
	HasEnableWideMetric() bool
	// LearnedLspFilter returns bool, set in IsisBasic.
	LearnedLspFilter() bool
	// SetLearnedLspFilter assigns bool provided by user to IsisBasic
	SetLearnedLspFilter(value bool) IsisBasic
	// HasLearnedLspFilter checks if LearnedLspFilter has been set in IsisBasic
	HasLearnedLspFilter() bool
}

// IPv4 Traffic Engineering(TE) router id. This address should be configured as an IPv4 Loopback address in 'ipv4_loopbacks' in the Device.
// Ipv4TeRouterId returns a string
func (obj *isisBasic) Ipv4TeRouterId() string {

	return *obj.obj.Ipv4TeRouterId

}

// IPv4 Traffic Engineering(TE) router id. This address should be configured as an IPv4 Loopback address in 'ipv4_loopbacks' in the Device.
// Ipv4TeRouterId returns a string
func (obj *isisBasic) HasIpv4TeRouterId() bool {
	return obj.obj.Ipv4TeRouterId != nil
}

// IPv4 Traffic Engineering(TE) router id. This address should be configured as an IPv4 Loopback address in 'ipv4_loopbacks' in the Device.
// SetIpv4TeRouterId sets the string value in the IsisBasic object
func (obj *isisBasic) SetIpv4TeRouterId(value string) IsisBasic {

	obj.obj.Ipv4TeRouterId = &value
	return obj
}

// Host name for the router. The host name is transmitted in all the packets sent from the router.
// Hostname returns a string
func (obj *isisBasic) Hostname() string {

	return *obj.obj.Hostname

}

// Host name for the router. The host name is transmitted in all the packets sent from the router.
// Hostname returns a string
func (obj *isisBasic) HasHostname() bool {
	return obj.obj.Hostname != nil
}

// Host name for the router. The host name is transmitted in all the packets sent from the router.
// SetHostname sets the string value in the IsisBasic object
func (obj *isisBasic) SetHostname(value string) IsisBasic {

	obj.obj.Hostname = &value
	return obj
}

// When set to true, it allows sending of more detailed metric information  for the routes using 32-bit wide values using TLV 135 IP reachability and  more detailed reachability information for IS reachability by using TLV 22.  The detailed usage is described in RFC3784.
// EnableWideMetric returns a bool
func (obj *isisBasic) EnableWideMetric() bool {

	return *obj.obj.EnableWideMetric

}

// When set to true, it allows sending of more detailed metric information  for the routes using 32-bit wide values using TLV 135 IP reachability and  more detailed reachability information for IS reachability by using TLV 22.  The detailed usage is described in RFC3784.
// EnableWideMetric returns a bool
func (obj *isisBasic) HasEnableWideMetric() bool {
	return obj.obj.EnableWideMetric != nil
}

// When set to true, it allows sending of more detailed metric information  for the routes using 32-bit wide values using TLV 135 IP reachability and  more detailed reachability information for IS reachability by using TLV 22.  The detailed usage is described in RFC3784.
// SetEnableWideMetric sets the bool value in the IsisBasic object
func (obj *isisBasic) SetEnableWideMetric(value bool) IsisBasic {

	obj.obj.EnableWideMetric = &value
	return obj
}

// Configuration for controlling storage of ISIS learned LSPs are received from the neighbors.
// LearnedLspFilter returns a bool
func (obj *isisBasic) LearnedLspFilter() bool {

	return *obj.obj.LearnedLspFilter

}

// Configuration for controlling storage of ISIS learned LSPs are received from the neighbors.
// LearnedLspFilter returns a bool
func (obj *isisBasic) HasLearnedLspFilter() bool {
	return obj.obj.LearnedLspFilter != nil
}

// Configuration for controlling storage of ISIS learned LSPs are received from the neighbors.
// SetLearnedLspFilter sets the bool value in the IsisBasic object
func (obj *isisBasic) SetLearnedLspFilter(value bool) IsisBasic {

	obj.obj.LearnedLspFilter = &value
	return obj
}

func (obj *isisBasic) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4TeRouterId != nil {

		err := obj.validateIpv4(obj.Ipv4TeRouterId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisBasic.Ipv4TeRouterId"))
		}

	}

}

func (obj *isisBasic) setDefault() {
	if obj.obj.EnableWideMetric == nil {
		obj.SetEnableWideMetric(true)
	}
	if obj.obj.LearnedLspFilter == nil {
		obj.SetLearnedLspFilter(false)
	}

}
