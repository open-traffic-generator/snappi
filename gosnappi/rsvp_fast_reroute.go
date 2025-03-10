package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpFastReroute *****
type rsvpFastReroute struct {
	validation
	obj          *otg.RsvpFastReroute
	marshaller   marshalRsvpFastReroute
	unMarshaller unMarshalRsvpFastReroute
}

func NewRsvpFastReroute() RsvpFastReroute {
	obj := rsvpFastReroute{obj: &otg.RsvpFastReroute{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpFastReroute) msg() *otg.RsvpFastReroute {
	return obj.obj
}

func (obj *rsvpFastReroute) setMsg(msg *otg.RsvpFastReroute) RsvpFastReroute {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpFastReroute struct {
	obj *rsvpFastReroute
}

type marshalRsvpFastReroute interface {
	// ToProto marshals RsvpFastReroute to protobuf object *otg.RsvpFastReroute
	ToProto() (*otg.RsvpFastReroute, error)
	// ToPbText marshals RsvpFastReroute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpFastReroute to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpFastReroute to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals RsvpFastReroute to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrsvpFastReroute struct {
	obj *rsvpFastReroute
}

type unMarshalRsvpFastReroute interface {
	// FromProto unmarshals RsvpFastReroute from protobuf object *otg.RsvpFastReroute
	FromProto(msg *otg.RsvpFastReroute) (RsvpFastReroute, error)
	// FromPbText unmarshals RsvpFastReroute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpFastReroute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpFastReroute from JSON text
	FromJson(value string) error
}

func (obj *rsvpFastReroute) Marshal() marshalRsvpFastReroute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpFastReroute{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpFastReroute) Unmarshal() unMarshalRsvpFastReroute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpFastReroute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpFastReroute) ToProto() (*otg.RsvpFastReroute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpFastReroute) FromProto(msg *otg.RsvpFastReroute) (RsvpFastReroute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpFastReroute) ToPbText() (string, error) {
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

func (m *unMarshalrsvpFastReroute) FromPbText(value string) error {
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

func (m *marshalrsvpFastReroute) ToYaml() (string, error) {
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

func (m *unMarshalrsvpFastReroute) FromYaml(value string) error {
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

func (m *marshalrsvpFastReroute) ToJsonRaw() (string, error) {
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

func (m *marshalrsvpFastReroute) ToJson() (string, error) {
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

func (m *unMarshalrsvpFastReroute) FromJson(value string) error {
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

func (obj *rsvpFastReroute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpFastReroute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpFastReroute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpFastReroute) Clone() (RsvpFastReroute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpFastReroute()
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

// RsvpFastReroute is configuration for the optional RSVP-TE FAST_REROUTE object included in Path Messages as defined in RFC4090.
type RsvpFastReroute interface {
	Validation
	// msg marshals RsvpFastReroute to protobuf object *otg.RsvpFastReroute
	// and doesn't set defaults
	msg() *otg.RsvpFastReroute
	// setMsg unmarshals RsvpFastReroute from protobuf object *otg.RsvpFastReroute
	// and doesn't set defaults
	setMsg(*otg.RsvpFastReroute) RsvpFastReroute
	// provides marshal interface
	Marshal() marshalRsvpFastReroute
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpFastReroute
	// validate validates RsvpFastReroute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpFastReroute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SetupPriority returns uint32, set in RsvpFastReroute.
	SetupPriority() uint32
	// SetSetupPriority assigns uint32 provided by user to RsvpFastReroute
	SetSetupPriority(value uint32) RsvpFastReroute
	// HasSetupPriority checks if SetupPriority has been set in RsvpFastReroute
	HasSetupPriority() bool
	// HoldingPriority returns uint32, set in RsvpFastReroute.
	HoldingPriority() uint32
	// SetHoldingPriority assigns uint32 provided by user to RsvpFastReroute
	SetHoldingPriority(value uint32) RsvpFastReroute
	// HasHoldingPriority checks if HoldingPriority has been set in RsvpFastReroute
	HasHoldingPriority() bool
	// HopLimit returns uint32, set in RsvpFastReroute.
	HopLimit() uint32
	// SetHopLimit assigns uint32 provided by user to RsvpFastReroute
	SetHopLimit(value uint32) RsvpFastReroute
	// HasHopLimit checks if HopLimit has been set in RsvpFastReroute
	HasHopLimit() bool
	// Bandwidth returns float32, set in RsvpFastReroute.
	Bandwidth() float32
	// SetBandwidth assigns float32 provided by user to RsvpFastReroute
	SetBandwidth(value float32) RsvpFastReroute
	// HasBandwidth checks if Bandwidth has been set in RsvpFastReroute
	HasBandwidth() bool
	// ExcludeAny returns string, set in RsvpFastReroute.
	ExcludeAny() string
	// SetExcludeAny assigns string provided by user to RsvpFastReroute
	SetExcludeAny(value string) RsvpFastReroute
	// HasExcludeAny checks if ExcludeAny has been set in RsvpFastReroute
	HasExcludeAny() bool
	// IncludeAny returns string, set in RsvpFastReroute.
	IncludeAny() string
	// SetIncludeAny assigns string provided by user to RsvpFastReroute
	SetIncludeAny(value string) RsvpFastReroute
	// HasIncludeAny checks if IncludeAny has been set in RsvpFastReroute
	HasIncludeAny() bool
	// IncludeAll returns string, set in RsvpFastReroute.
	IncludeAll() string
	// SetIncludeAll assigns string provided by user to RsvpFastReroute
	SetIncludeAll(value string) RsvpFastReroute
	// HasIncludeAll checks if IncludeAll has been set in RsvpFastReroute
	HasIncludeAll() bool
	// OneToOneBackupDesired returns bool, set in RsvpFastReroute.
	OneToOneBackupDesired() bool
	// SetOneToOneBackupDesired assigns bool provided by user to RsvpFastReroute
	SetOneToOneBackupDesired(value bool) RsvpFastReroute
	// HasOneToOneBackupDesired checks if OneToOneBackupDesired has been set in RsvpFastReroute
	HasOneToOneBackupDesired() bool
	// FacilityBackupDesired returns bool, set in RsvpFastReroute.
	FacilityBackupDesired() bool
	// SetFacilityBackupDesired assigns bool provided by user to RsvpFastReroute
	SetFacilityBackupDesired(value bool) RsvpFastReroute
	// HasFacilityBackupDesired checks if FacilityBackupDesired has been set in RsvpFastReroute
	HasFacilityBackupDesired() bool
}

// Specifies the value of the Setup Priority field. This controls whether the backup LSP should pre-empt existing LSP that is setup with certain Holding Priority. While setting up a backup LSP, preemption of existing LSP can happen  if resource limitation is encountered (e.g bandwidth availability).
// SetupPriority returns a uint32
func (obj *rsvpFastReroute) SetupPriority() uint32 {

	return *obj.obj.SetupPriority

}

// Specifies the value of the Setup Priority field. This controls whether the backup LSP should pre-empt existing LSP that is setup with certain Holding Priority. While setting up a backup LSP, preemption of existing LSP can happen  if resource limitation is encountered (e.g bandwidth availability).
// SetupPriority returns a uint32
func (obj *rsvpFastReroute) HasSetupPriority() bool {
	return obj.obj.SetupPriority != nil
}

// Specifies the value of the Setup Priority field. This controls whether the backup LSP should pre-empt existing LSP that is setup with certain Holding Priority. While setting up a backup LSP, preemption of existing LSP can happen  if resource limitation is encountered (e.g bandwidth availability).
// SetSetupPriority sets the uint32 value in the RsvpFastReroute object
func (obj *rsvpFastReroute) SetSetupPriority(value uint32) RsvpFastReroute {

	obj.obj.SetupPriority = &value
	return obj
}

// Specifies the value of the Holding Priority field. This controls whether a new LSP being created with certain Setup Priority should pre-empt this LSP set up with this Holding Priority. While setting up a new LSP, preemption of existing LSP can happen  if resource limitation is encountered (e.g bandwidth availability).
// HoldingPriority returns a uint32
func (obj *rsvpFastReroute) HoldingPriority() uint32 {

	return *obj.obj.HoldingPriority

}

// Specifies the value of the Holding Priority field. This controls whether a new LSP being created with certain Setup Priority should pre-empt this LSP set up with this Holding Priority. While setting up a new LSP, preemption of existing LSP can happen  if resource limitation is encountered (e.g bandwidth availability).
// HoldingPriority returns a uint32
func (obj *rsvpFastReroute) HasHoldingPriority() bool {
	return obj.obj.HoldingPriority != nil
}

// Specifies the value of the Holding Priority field. This controls whether a new LSP being created with certain Setup Priority should pre-empt this LSP set up with this Holding Priority. While setting up a new LSP, preemption of existing LSP can happen  if resource limitation is encountered (e.g bandwidth availability).
// SetHoldingPriority sets the uint32 value in the RsvpFastReroute object
func (obj *rsvpFastReroute) SetHoldingPriority(value uint32) RsvpFastReroute {

	obj.obj.HoldingPriority = &value
	return obj
}

// Specifies the value of the Hop Limit field. This controls the maximum number of hops the LSP should traverse to reach the  LSP end-point.
// HopLimit returns a uint32
func (obj *rsvpFastReroute) HopLimit() uint32 {

	return *obj.obj.HopLimit

}

// Specifies the value of the Hop Limit field. This controls the maximum number of hops the LSP should traverse to reach the  LSP end-point.
// HopLimit returns a uint32
func (obj *rsvpFastReroute) HasHopLimit() bool {
	return obj.obj.HopLimit != nil
}

// Specifies the value of the Hop Limit field. This controls the maximum number of hops the LSP should traverse to reach the  LSP end-point.
// SetHopLimit sets the uint32 value in the RsvpFastReroute object
func (obj *rsvpFastReroute) SetHopLimit(value uint32) RsvpFastReroute {

	obj.obj.HopLimit = &value
	return obj
}

// Specifies the value of the Bandwidth field as a 32-bit IEEE floating point integer, in bytes per second, as desired for the LSP.
// Bandwidth returns a float32
func (obj *rsvpFastReroute) Bandwidth() float32 {

	return *obj.obj.Bandwidth

}

// Specifies the value of the Bandwidth field as a 32-bit IEEE floating point integer, in bytes per second, as desired for the LSP.
// Bandwidth returns a float32
func (obj *rsvpFastReroute) HasBandwidth() bool {
	return obj.obj.Bandwidth != nil
}

// Specifies the value of the Bandwidth field as a 32-bit IEEE floating point integer, in bytes per second, as desired for the LSP.
// SetBandwidth sets the float32 value in the RsvpFastReroute object
func (obj *rsvpFastReroute) SetBandwidth(value float32) RsvpFastReroute {

	obj.obj.Bandwidth = &value
	return obj
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link unacceptable. A null set (all bits set to zero) doesn't render the link unacceptable. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// ExcludeAny returns a string
func (obj *rsvpFastReroute) ExcludeAny() string {

	return *obj.obj.ExcludeAny

}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link unacceptable. A null set (all bits set to zero) doesn't render the link unacceptable. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// ExcludeAny returns a string
func (obj *rsvpFastReroute) HasExcludeAny() bool {
	return obj.obj.ExcludeAny != nil
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link unacceptable. A null set (all bits set to zero) doesn't render the link unacceptable. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// SetExcludeAny sets the string value in the RsvpFastReroute object
func (obj *rsvpFastReroute) SetExcludeAny(value string) RsvpFastReroute {

	obj.obj.ExcludeAny = &value
	return obj
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// IncludeAny returns a string
func (obj *rsvpFastReroute) IncludeAny() string {

	return *obj.obj.IncludeAny

}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// IncludeAny returns a string
func (obj *rsvpFastReroute) HasIncludeAny() bool {
	return obj.obj.IncludeAny != nil
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// SetIncludeAny sets the string value in the RsvpFastReroute object
func (obj *rsvpFastReroute) SetIncludeAny(value string) RsvpFastReroute {

	obj.obj.IncludeAny = &value
	return obj
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel all of which must be present for a link to be acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// IncludeAll returns a string
func (obj *rsvpFastReroute) IncludeAll() string {

	return *obj.obj.IncludeAll

}

// A 32-bit vector representing a set of attribute filters associated with a tunnel all of which must be present for a link to be acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// IncludeAll returns a string
func (obj *rsvpFastReroute) HasIncludeAll() bool {
	return obj.obj.IncludeAll != nil
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel all of which must be present for a link to be acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// SetIncludeAll sets the string value in the RsvpFastReroute object
func (obj *rsvpFastReroute) SetIncludeAll(value string) RsvpFastReroute {

	obj.obj.IncludeAll = &value
	return obj
}

// Requests protection via the one-to-one backup method.
// OneToOneBackupDesired returns a bool
func (obj *rsvpFastReroute) OneToOneBackupDesired() bool {

	return *obj.obj.OneToOneBackupDesired

}

// Requests protection via the one-to-one backup method.
// OneToOneBackupDesired returns a bool
func (obj *rsvpFastReroute) HasOneToOneBackupDesired() bool {
	return obj.obj.OneToOneBackupDesired != nil
}

// Requests protection via the one-to-one backup method.
// SetOneToOneBackupDesired sets the bool value in the RsvpFastReroute object
func (obj *rsvpFastReroute) SetOneToOneBackupDesired(value bool) RsvpFastReroute {

	obj.obj.OneToOneBackupDesired = &value
	return obj
}

// Requests protection via the facility backup method.
// FacilityBackupDesired returns a bool
func (obj *rsvpFastReroute) FacilityBackupDesired() bool {

	return *obj.obj.FacilityBackupDesired

}

// Requests protection via the facility backup method.
// FacilityBackupDesired returns a bool
func (obj *rsvpFastReroute) HasFacilityBackupDesired() bool {
	return obj.obj.FacilityBackupDesired != nil
}

// Requests protection via the facility backup method.
// SetFacilityBackupDesired sets the bool value in the RsvpFastReroute object
func (obj *rsvpFastReroute) SetFacilityBackupDesired(value bool) RsvpFastReroute {

	obj.obj.FacilityBackupDesired = &value
	return obj
}

func (obj *rsvpFastReroute) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SetupPriority != nil {

		if *obj.obj.SetupPriority > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpFastReroute.SetupPriority <= 7 but Got %d", *obj.obj.SetupPriority))
		}

	}

	if obj.obj.HoldingPriority != nil {

		if *obj.obj.HoldingPriority > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpFastReroute.HoldingPriority <= 7 but Got %d", *obj.obj.HoldingPriority))
		}

	}

	if obj.obj.HopLimit != nil {

		if *obj.obj.HopLimit > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpFastReroute.HopLimit <= 255 but Got %d", *obj.obj.HopLimit))
		}

	}

	if obj.obj.ExcludeAny != nil {

		if len(*obj.obj.ExcludeAny) < 0 || len(*obj.obj.ExcludeAny) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of RsvpFastReroute.ExcludeAny <= 8 but Got %d",
					len(*obj.obj.ExcludeAny)))
		}

	}

	if obj.obj.IncludeAny != nil {

		if len(*obj.obj.IncludeAny) < 0 || len(*obj.obj.IncludeAny) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of RsvpFastReroute.IncludeAny <= 8 but Got %d",
					len(*obj.obj.IncludeAny)))
		}

	}

	if obj.obj.IncludeAll != nil {

		if len(*obj.obj.IncludeAll) < 0 || len(*obj.obj.IncludeAll) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of RsvpFastReroute.IncludeAll <= 8 but Got %d",
					len(*obj.obj.IncludeAll)))
		}

	}

}

func (obj *rsvpFastReroute) setDefault() {
	if obj.obj.SetupPriority == nil {
		obj.SetSetupPriority(7)
	}
	if obj.obj.HoldingPriority == nil {
		obj.SetHoldingPriority(7)
	}
	if obj.obj.HopLimit == nil {
		obj.SetHopLimit(3)
	}
	if obj.obj.Bandwidth == nil {
		obj.SetBandwidth(0)
	}
	if obj.obj.ExcludeAny == nil {
		obj.SetExcludeAny("0")
	}
	if obj.obj.IncludeAny == nil {
		obj.SetIncludeAny("0")
	}
	if obj.obj.IncludeAll == nil {
		obj.SetIncludeAll("0")
	}
	if obj.obj.OneToOneBackupDesired == nil {
		obj.SetOneToOneBackupDesired(false)
	}
	if obj.obj.FacilityBackupDesired == nil {
		obj.SetFacilityBackupDesired(false)
	}

}
