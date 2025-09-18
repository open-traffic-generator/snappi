package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DhcpServerV4Pool *****
type dhcpServerV4Pool struct {
	validation
	obj           *otg.DhcpServerV4Pool
	marshaller    marshalDhcpServerV4Pool
	unMarshaller  unMarshalDhcpServerV4Pool
	optionsHolder DhcpServerV4PoolOption
}

func NewDhcpServerV4Pool() DhcpServerV4Pool {
	obj := dhcpServerV4Pool{obj: &otg.DhcpServerV4Pool{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpServerV4Pool) msg() *otg.DhcpServerV4Pool {
	return obj.obj
}

func (obj *dhcpServerV4Pool) setMsg(msg *otg.DhcpServerV4Pool) DhcpServerV4Pool {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpServerV4Pool struct {
	obj *dhcpServerV4Pool
}

type marshalDhcpServerV4Pool interface {
	// ToProto marshals DhcpServerV4Pool to protobuf object *otg.DhcpServerV4Pool
	ToProto() (*otg.DhcpServerV4Pool, error)
	// ToPbText marshals DhcpServerV4Pool to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DhcpServerV4Pool to YAML text
	ToYaml() (string, error)
	// ToJson marshals DhcpServerV4Pool to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpServerV4Pool struct {
	obj *dhcpServerV4Pool
}

type unMarshalDhcpServerV4Pool interface {
	// FromProto unmarshals DhcpServerV4Pool from protobuf object *otg.DhcpServerV4Pool
	FromProto(msg *otg.DhcpServerV4Pool) (DhcpServerV4Pool, error)
	// FromPbText unmarshals DhcpServerV4Pool from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DhcpServerV4Pool from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DhcpServerV4Pool from JSON text
	FromJson(value string) error
}

func (obj *dhcpServerV4Pool) Marshal() marshalDhcpServerV4Pool {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpServerV4Pool{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpServerV4Pool) Unmarshal() unMarshalDhcpServerV4Pool {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpServerV4Pool{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpServerV4Pool) ToProto() (*otg.DhcpServerV4Pool, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpServerV4Pool) FromProto(msg *otg.DhcpServerV4Pool) (DhcpServerV4Pool, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpServerV4Pool) ToPbText() (string, error) {
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

func (m *unMarshaldhcpServerV4Pool) FromPbText(value string) error {
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

func (m *marshaldhcpServerV4Pool) ToYaml() (string, error) {
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

func (m *unMarshaldhcpServerV4Pool) FromYaml(value string) error {
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

func (m *marshaldhcpServerV4Pool) ToJson() (string, error) {
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

func (m *unMarshaldhcpServerV4Pool) FromJson(value string) error {
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

func (obj *dhcpServerV4Pool) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpServerV4Pool) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpServerV4Pool) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpServerV4Pool) Clone() (DhcpServerV4Pool, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpServerV4Pool()
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

func (obj *dhcpServerV4Pool) setNil() {
	obj.optionsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DhcpServerV4Pool is configuration for DHCPv4 address pool for a lease.
type DhcpServerV4Pool interface {
	Validation
	// msg marshals DhcpServerV4Pool to protobuf object *otg.DhcpServerV4Pool
	// and doesn't set defaults
	msg() *otg.DhcpServerV4Pool
	// setMsg unmarshals DhcpServerV4Pool from protobuf object *otg.DhcpServerV4Pool
	// and doesn't set defaults
	setMsg(*otg.DhcpServerV4Pool) DhcpServerV4Pool
	// provides marshal interface
	Marshal() marshalDhcpServerV4Pool
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpServerV4Pool
	// validate validates DhcpServerV4Pool
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DhcpServerV4Pool, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in DhcpServerV4Pool.
	Name() string
	// SetName assigns string provided by user to DhcpServerV4Pool
	SetName(value string) DhcpServerV4Pool
	// HasName checks if Name has been set in DhcpServerV4Pool
	HasName() bool
	// LeaseTime returns uint32, set in DhcpServerV4Pool.
	LeaseTime() uint32
	// SetLeaseTime assigns uint32 provided by user to DhcpServerV4Pool
	SetLeaseTime(value uint32) DhcpServerV4Pool
	// HasLeaseTime checks if LeaseTime has been set in DhcpServerV4Pool
	HasLeaseTime() bool
	// StartAddress returns string, set in DhcpServerV4Pool.
	StartAddress() string
	// SetStartAddress assigns string provided by user to DhcpServerV4Pool
	SetStartAddress(value string) DhcpServerV4Pool
	// PrefixLength returns uint32, set in DhcpServerV4Pool.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to DhcpServerV4Pool
	SetPrefixLength(value uint32) DhcpServerV4Pool
	// HasPrefixLength checks if PrefixLength has been set in DhcpServerV4Pool
	HasPrefixLength() bool
	// Count returns uint32, set in DhcpServerV4Pool.
	Count() uint32
	// SetCount assigns uint32 provided by user to DhcpServerV4Pool
	SetCount(value uint32) DhcpServerV4Pool
	// HasCount checks if Count has been set in DhcpServerV4Pool
	HasCount() bool
	// Step returns uint32, set in DhcpServerV4Pool.
	Step() uint32
	// SetStep assigns uint32 provided by user to DhcpServerV4Pool
	SetStep(value uint32) DhcpServerV4Pool
	// HasStep checks if Step has been set in DhcpServerV4Pool
	HasStep() bool
	// Options returns DhcpServerV4PoolOption, set in DhcpServerV4Pool.
	// DhcpServerV4PoolOption is optional configuration for DHCPv4 address pool for the lease.
	Options() DhcpServerV4PoolOption
	// SetOptions assigns DhcpServerV4PoolOption provided by user to DhcpServerV4Pool.
	// DhcpServerV4PoolOption is optional configuration for DHCPv4 address pool for the lease.
	SetOptions(value DhcpServerV4PoolOption) DhcpServerV4Pool
	// HasOptions checks if Options has been set in DhcpServerV4Pool
	HasOptions() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *dhcpServerV4Pool) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *dhcpServerV4Pool) HasName() bool {
	return obj.obj.Name != nil
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DhcpServerV4Pool object
func (obj *dhcpServerV4Pool) SetName(value string) DhcpServerV4Pool {

	obj.obj.Name = &value
	return obj
}

// The duration of time in seconds that is assigned to a lease.
// LeaseTime returns a uint32
func (obj *dhcpServerV4Pool) LeaseTime() uint32 {

	return *obj.obj.LeaseTime

}

// The duration of time in seconds that is assigned to a lease.
// LeaseTime returns a uint32
func (obj *dhcpServerV4Pool) HasLeaseTime() bool {
	return obj.obj.LeaseTime != nil
}

// The duration of time in seconds that is assigned to a lease.
// SetLeaseTime sets the uint32 value in the DhcpServerV4Pool object
func (obj *dhcpServerV4Pool) SetLeaseTime(value uint32) DhcpServerV4Pool {

	obj.obj.LeaseTime = &value
	return obj
}

// The IPv4 address of the first lease pool.
// StartAddress returns a string
func (obj *dhcpServerV4Pool) StartAddress() string {

	return *obj.obj.StartAddress

}

// The IPv4 address of the first lease pool.
// SetStartAddress sets the string value in the DhcpServerV4Pool object
func (obj *dhcpServerV4Pool) SetStartAddress(value string) DhcpServerV4Pool {

	obj.obj.StartAddress = &value
	return obj
}

// The IPv4 network prefix length to be applied to the address.
// PrefixLength returns a uint32
func (obj *dhcpServerV4Pool) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The IPv4 network prefix length to be applied to the address.
// PrefixLength returns a uint32
func (obj *dhcpServerV4Pool) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The IPv4 network prefix length to be applied to the address.
// SetPrefixLength sets the uint32 value in the DhcpServerV4Pool object
func (obj *dhcpServerV4Pool) SetPrefixLength(value uint32) DhcpServerV4Pool {

	obj.obj.PrefixLength = &value
	return obj
}

// The total number of addresses in the pool.
// Count returns a uint32
func (obj *dhcpServerV4Pool) Count() uint32 {

	return *obj.obj.Count

}

// The total number of addresses in the pool.
// Count returns a uint32
func (obj *dhcpServerV4Pool) HasCount() bool {
	return obj.obj.Count != nil
}

// The total number of addresses in the pool.
// SetCount sets the uint32 value in the DhcpServerV4Pool object
func (obj *dhcpServerV4Pool) SetCount(value uint32) DhcpServerV4Pool {

	obj.obj.Count = &value
	return obj
}

// The increment value for the lease address within the lease pool. The value is incremented according to the prefix_length and step.
// Step returns a uint32
func (obj *dhcpServerV4Pool) Step() uint32 {

	return *obj.obj.Step

}

// The increment value for the lease address within the lease pool. The value is incremented according to the prefix_length and step.
// Step returns a uint32
func (obj *dhcpServerV4Pool) HasStep() bool {
	return obj.obj.Step != nil
}

// The increment value for the lease address within the lease pool. The value is incremented according to the prefix_length and step.
// SetStep sets the uint32 value in the DhcpServerV4Pool object
func (obj *dhcpServerV4Pool) SetStep(value uint32) DhcpServerV4Pool {

	obj.obj.Step = &value
	return obj
}

// Optional configuration for DHCPv4 address pool for the lease.
// Options returns a DhcpServerV4PoolOption
func (obj *dhcpServerV4Pool) Options() DhcpServerV4PoolOption {
	if obj.obj.Options == nil {
		obj.obj.Options = NewDhcpServerV4PoolOption().msg()
	}
	if obj.optionsHolder == nil {
		obj.optionsHolder = &dhcpServerV4PoolOption{obj: obj.obj.Options}
	}
	return obj.optionsHolder
}

// Optional configuration for DHCPv4 address pool for the lease.
// Options returns a DhcpServerV4PoolOption
func (obj *dhcpServerV4Pool) HasOptions() bool {
	return obj.obj.Options != nil
}

// Optional configuration for DHCPv4 address pool for the lease.
// SetOptions sets the DhcpServerV4PoolOption value in the DhcpServerV4Pool object
func (obj *dhcpServerV4Pool) SetOptions(value DhcpServerV4PoolOption) DhcpServerV4Pool {

	obj.optionsHolder = nil
	obj.obj.Options = value.msg()

	return obj
}

func (obj *dhcpServerV4Pool) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"DhcpServerV4Pool.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.LeaseTime != nil {

		if *obj.obj.LeaseTime < 10 || *obj.obj.LeaseTime > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("10 <= DhcpServerV4Pool.LeaseTime <= 4294967295 but Got %d", *obj.obj.LeaseTime))
		}

	}

	// StartAddress is required
	if obj.obj.StartAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "StartAddress is required field on interface DhcpServerV4Pool")
	}
	if obj.obj.StartAddress != nil {

		err := obj.validateIpv4(obj.StartAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DhcpServerV4Pool.StartAddress"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= DhcpServerV4Pool.PrefixLength <= 32 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count < 1 || *obj.obj.Count > 1000000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DhcpServerV4Pool.Count <= 1000000 but Got %d", *obj.obj.Count))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step < 1 || *obj.obj.Step > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DhcpServerV4Pool.Step <= 4294967295 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Options != nil {

		obj.Options().validateObj(vObj, set_default)
	}

}

func (obj *dhcpServerV4Pool) setDefault() {
	if obj.obj.LeaseTime == nil {
		obj.SetLeaseTime(86400)
	}
	if obj.obj.PrefixLength == nil {
		obj.SetPrefixLength(24)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}

}
