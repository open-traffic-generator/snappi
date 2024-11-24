package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LinkStateTE *****
type linkStateTE struct {
	validation
	obj                      *otg.LinkStateTE
	marshaller               marshalLinkStateTE
	unMarshaller             unMarshalLinkStateTE
	priorityBandwidthsHolder LinkStatepriorityBandwidths
}

func NewLinkStateTE() LinkStateTE {
	obj := linkStateTE{obj: &otg.LinkStateTE{}}
	obj.setDefault()
	return &obj
}

func (obj *linkStateTE) msg() *otg.LinkStateTE {
	return obj.obj
}

func (obj *linkStateTE) setMsg(msg *otg.LinkStateTE) LinkStateTE {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallinkStateTE struct {
	obj *linkStateTE
}

type marshalLinkStateTE interface {
	// ToProto marshals LinkStateTE to protobuf object *otg.LinkStateTE
	ToProto() (*otg.LinkStateTE, error)
	// ToPbText marshals LinkStateTE to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LinkStateTE to YAML text
	ToYaml() (string, error)
	// ToJson marshals LinkStateTE to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LinkStateTE to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallinkStateTE struct {
	obj *linkStateTE
}

type unMarshalLinkStateTE interface {
	// FromProto unmarshals LinkStateTE from protobuf object *otg.LinkStateTE
	FromProto(msg *otg.LinkStateTE) (LinkStateTE, error)
	// FromPbText unmarshals LinkStateTE from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LinkStateTE from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LinkStateTE from JSON text
	FromJson(value string) error
}

func (obj *linkStateTE) Marshal() marshalLinkStateTE {
	if obj.marshaller == nil {
		obj.marshaller = &marshallinkStateTE{obj: obj}
	}
	return obj.marshaller
}

func (obj *linkStateTE) Unmarshal() unMarshalLinkStateTE {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallinkStateTE{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallinkStateTE) ToProto() (*otg.LinkStateTE, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallinkStateTE) FromProto(msg *otg.LinkStateTE) (LinkStateTE, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallinkStateTE) ToPbText() (string, error) {
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

func (m *unMarshallinkStateTE) FromPbText(value string) error {
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

func (m *marshallinkStateTE) ToYaml() (string, error) {
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

func (m *unMarshallinkStateTE) FromYaml(value string) error {
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

func (m *marshallinkStateTE) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshallinkStateTE) ToJson() (string, error) {
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

func (m *unMarshallinkStateTE) FromJson(value string) error {
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

func (obj *linkStateTE) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *linkStateTE) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *linkStateTE) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *linkStateTE) Clone() (LinkStateTE, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLinkStateTE()
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

func (obj *linkStateTE) setNil() {
	obj.priorityBandwidthsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LinkStateTE is a container for Traffic Engineering properties on a interface.
type LinkStateTE interface {
	Validation
	// msg marshals LinkStateTE to protobuf object *otg.LinkStateTE
	// and doesn't set defaults
	msg() *otg.LinkStateTE
	// setMsg unmarshals LinkStateTE from protobuf object *otg.LinkStateTE
	// and doesn't set defaults
	setMsg(*otg.LinkStateTE) LinkStateTE
	// provides marshal interface
	Marshal() marshalLinkStateTE
	// provides unmarshal interface
	Unmarshal() unMarshalLinkStateTE
	// validate validates LinkStateTE
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LinkStateTE, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AdministrativeGroup returns string, set in LinkStateTE.
	AdministrativeGroup() string
	// SetAdministrativeGroup assigns string provided by user to LinkStateTE
	SetAdministrativeGroup(value string) LinkStateTE
	// HasAdministrativeGroup checks if AdministrativeGroup has been set in LinkStateTE
	HasAdministrativeGroup() bool
	// MetricLevel returns uint32, set in LinkStateTE.
	MetricLevel() uint32
	// SetMetricLevel assigns uint32 provided by user to LinkStateTE
	SetMetricLevel(value uint32) LinkStateTE
	// HasMetricLevel checks if MetricLevel has been set in LinkStateTE
	HasMetricLevel() bool
	// MaxBandwith returns uint32, set in LinkStateTE.
	MaxBandwith() uint32
	// SetMaxBandwith assigns uint32 provided by user to LinkStateTE
	SetMaxBandwith(value uint32) LinkStateTE
	// HasMaxBandwith checks if MaxBandwith has been set in LinkStateTE
	HasMaxBandwith() bool
	// MaxReservableBandwidth returns uint32, set in LinkStateTE.
	MaxReservableBandwidth() uint32
	// SetMaxReservableBandwidth assigns uint32 provided by user to LinkStateTE
	SetMaxReservableBandwidth(value uint32) LinkStateTE
	// HasMaxReservableBandwidth checks if MaxReservableBandwidth has been set in LinkStateTE
	HasMaxReservableBandwidth() bool
	// PriorityBandwidths returns LinkStatepriorityBandwidths, set in LinkStateTE.
	// LinkStatepriorityBandwidths is specifies the amount of bandwidth that can be reserved with a setup priority of 0
	// through 7, arranged in increasing order with priority 0 having highest priority.
	// In ISIS, this is sent in sub-TLV (11) of Extended IS Reachability TLV.
	PriorityBandwidths() LinkStatepriorityBandwidths
	// SetPriorityBandwidths assigns LinkStatepriorityBandwidths provided by user to LinkStateTE.
	// LinkStatepriorityBandwidths is specifies the amount of bandwidth that can be reserved with a setup priority of 0
	// through 7, arranged in increasing order with priority 0 having highest priority.
	// In ISIS, this is sent in sub-TLV (11) of Extended IS Reachability TLV.
	SetPriorityBandwidths(value LinkStatepriorityBandwidths) LinkStateTE
	// HasPriorityBandwidths checks if PriorityBandwidths has been set in LinkStateTE
	HasPriorityBandwidths() bool
	setNil()
}

// The Administrative group sub-TLV (sub-TLV 3). It is a 4-octet
// user-defined bit mask used to assign administrative group numbers
// to the interface, for use in assigning colors and resource classes.
// Each set bit corresponds to a single administrative group for this
// interface. The settings translate into Group numbers, which range
// from 0 to 31 (integers).
// AdministrativeGroup returns a string
func (obj *linkStateTE) AdministrativeGroup() string {

	return *obj.obj.AdministrativeGroup

}

// The Administrative group sub-TLV (sub-TLV 3). It is a 4-octet
// user-defined bit mask used to assign administrative group numbers
// to the interface, for use in assigning colors and resource classes.
// Each set bit corresponds to a single administrative group for this
// interface. The settings translate into Group numbers, which range
// from 0 to 31 (integers).
// AdministrativeGroup returns a string
func (obj *linkStateTE) HasAdministrativeGroup() bool {
	return obj.obj.AdministrativeGroup != nil
}

// The Administrative group sub-TLV (sub-TLV 3). It is a 4-octet
// user-defined bit mask used to assign administrative group numbers
// to the interface, for use in assigning colors and resource classes.
// Each set bit corresponds to a single administrative group for this
// interface. The settings translate into Group numbers, which range
// from 0 to 31 (integers).
// SetAdministrativeGroup sets the string value in the LinkStateTE object
func (obj *linkStateTE) SetAdministrativeGroup(value string) LinkStateTE {

	obj.obj.AdministrativeGroup = &value
	return obj
}

// The user-assigned link metric for Traffic Engineering.
// MetricLevel returns a uint32
func (obj *linkStateTE) MetricLevel() uint32 {

	return *obj.obj.MetricLevel

}

// The user-assigned link metric for Traffic Engineering.
// MetricLevel returns a uint32
func (obj *linkStateTE) HasMetricLevel() bool {
	return obj.obj.MetricLevel != nil
}

// The user-assigned link metric for Traffic Engineering.
// SetMetricLevel sets the uint32 value in the LinkStateTE object
func (obj *linkStateTE) SetMetricLevel(value uint32) LinkStateTE {

	obj.obj.MetricLevel = &value
	return obj
}

// The maximum link bandwidth (sub-TLV 9) in bytes/sec allowed for this
// link for a direction.
// MaxBandwith returns a uint32
func (obj *linkStateTE) MaxBandwith() uint32 {

	return *obj.obj.MaxBandwith

}

// The maximum link bandwidth (sub-TLV 9) in bytes/sec allowed for this
// link for a direction.
// MaxBandwith returns a uint32
func (obj *linkStateTE) HasMaxBandwith() bool {
	return obj.obj.MaxBandwith != nil
}

// The maximum link bandwidth (sub-TLV 9) in bytes/sec allowed for this
// link for a direction.
// SetMaxBandwith sets the uint32 value in the LinkStateTE object
func (obj *linkStateTE) SetMaxBandwith(value uint32) LinkStateTE {

	obj.obj.MaxBandwith = &value
	return obj
}

// The maximum link bandwidth (sub-TLV 10) in bytes/sec allowed for this
// link in a direction.
// MaxReservableBandwidth returns a uint32
func (obj *linkStateTE) MaxReservableBandwidth() uint32 {

	return *obj.obj.MaxReservableBandwidth

}

// The maximum link bandwidth (sub-TLV 10) in bytes/sec allowed for this
// link in a direction.
// MaxReservableBandwidth returns a uint32
func (obj *linkStateTE) HasMaxReservableBandwidth() bool {
	return obj.obj.MaxReservableBandwidth != nil
}

// The maximum link bandwidth (sub-TLV 10) in bytes/sec allowed for this
// link in a direction.
// SetMaxReservableBandwidth sets the uint32 value in the LinkStateTE object
func (obj *linkStateTE) SetMaxReservableBandwidth(value uint32) LinkStateTE {

	obj.obj.MaxReservableBandwidth = &value
	return obj
}

// Configuration of bandwidths of priority 0 through priority 7.
// PriorityBandwidths returns a LinkStatepriorityBandwidths
func (obj *linkStateTE) PriorityBandwidths() LinkStatepriorityBandwidths {
	if obj.obj.PriorityBandwidths == nil {
		obj.obj.PriorityBandwidths = NewLinkStatepriorityBandwidths().msg()
	}
	if obj.priorityBandwidthsHolder == nil {
		obj.priorityBandwidthsHolder = &linkStatepriorityBandwidths{obj: obj.obj.PriorityBandwidths}
	}
	return obj.priorityBandwidthsHolder
}

// Configuration of bandwidths of priority 0 through priority 7.
// PriorityBandwidths returns a LinkStatepriorityBandwidths
func (obj *linkStateTE) HasPriorityBandwidths() bool {
	return obj.obj.PriorityBandwidths != nil
}

// Configuration of bandwidths of priority 0 through priority 7.
// SetPriorityBandwidths sets the LinkStatepriorityBandwidths value in the LinkStateTE object
func (obj *linkStateTE) SetPriorityBandwidths(value LinkStatepriorityBandwidths) LinkStateTE {

	obj.priorityBandwidthsHolder = nil
	obj.obj.PriorityBandwidths = value.msg()

	return obj
}

func (obj *linkStateTE) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.AdministrativeGroup != nil {

		err := obj.validateHex(obj.AdministrativeGroup())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LinkStateTE.AdministrativeGroup"))
		}

	}

	if obj.obj.PriorityBandwidths != nil {

		obj.PriorityBandwidths().validateObj(vObj, set_default)
	}

}

func (obj *linkStateTE) setDefault() {
	if obj.obj.AdministrativeGroup == nil {
		obj.SetAdministrativeGroup("00000000")
	}
	if obj.obj.MetricLevel == nil {
		obj.SetMetricLevel(0)
	}
	if obj.obj.MaxBandwith == nil {
		obj.SetMaxBandwith(125000000)
	}
	if obj.obj.MaxReservableBandwidth == nil {
		obj.SetMaxReservableBandwidth(125000000)
	}

}
