package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpResourceAffinities *****
type rsvpResourceAffinities struct {
	validation
	obj          *otg.RsvpResourceAffinities
	marshaller   marshalRsvpResourceAffinities
	unMarshaller unMarshalRsvpResourceAffinities
}

func NewRsvpResourceAffinities() RsvpResourceAffinities {
	obj := rsvpResourceAffinities{obj: &otg.RsvpResourceAffinities{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpResourceAffinities) msg() *otg.RsvpResourceAffinities {
	return obj.obj
}

func (obj *rsvpResourceAffinities) setMsg(msg *otg.RsvpResourceAffinities) RsvpResourceAffinities {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpResourceAffinities struct {
	obj *rsvpResourceAffinities
}

type marshalRsvpResourceAffinities interface {
	// ToProto marshals RsvpResourceAffinities to protobuf object *otg.RsvpResourceAffinities
	ToProto() (*otg.RsvpResourceAffinities, error)
	// ToPbText marshals RsvpResourceAffinities to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpResourceAffinities to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpResourceAffinities to JSON text
	ToJson() (string, error)
}

type unMarshalrsvpResourceAffinities struct {
	obj *rsvpResourceAffinities
}

type unMarshalRsvpResourceAffinities interface {
	// FromProto unmarshals RsvpResourceAffinities from protobuf object *otg.RsvpResourceAffinities
	FromProto(msg *otg.RsvpResourceAffinities) (RsvpResourceAffinities, error)
	// FromPbText unmarshals RsvpResourceAffinities from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpResourceAffinities from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpResourceAffinities from JSON text
	FromJson(value string) error
}

func (obj *rsvpResourceAffinities) Marshal() marshalRsvpResourceAffinities {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpResourceAffinities{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpResourceAffinities) Unmarshal() unMarshalRsvpResourceAffinities {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpResourceAffinities{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpResourceAffinities) ToProto() (*otg.RsvpResourceAffinities, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpResourceAffinities) FromProto(msg *otg.RsvpResourceAffinities) (RsvpResourceAffinities, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpResourceAffinities) ToPbText() (string, error) {
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

func (m *unMarshalrsvpResourceAffinities) FromPbText(value string) error {
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

func (m *marshalrsvpResourceAffinities) ToYaml() (string, error) {
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

func (m *unMarshalrsvpResourceAffinities) FromYaml(value string) error {
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

func (m *marshalrsvpResourceAffinities) ToJson() (string, error) {
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

func (m *unMarshalrsvpResourceAffinities) FromJson(value string) error {
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

func (obj *rsvpResourceAffinities) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpResourceAffinities) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpResourceAffinities) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpResourceAffinities) Clone() (RsvpResourceAffinities, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpResourceAffinities()
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

// RsvpResourceAffinities is this is an optional object. If included, the extended SESSION_ATTRIBUTE object is sent in the Path message containing
// the additional fields included in this object. This contains a set of three bitmaps using which further constraints can be
// set on the path calculated for the LSP based on the Admin Group settings in the IGP (e.g ISIS or OSPF interface).
type RsvpResourceAffinities interface {
	Validation
	// msg marshals RsvpResourceAffinities to protobuf object *otg.RsvpResourceAffinities
	// and doesn't set defaults
	msg() *otg.RsvpResourceAffinities
	// setMsg unmarshals RsvpResourceAffinities from protobuf object *otg.RsvpResourceAffinities
	// and doesn't set defaults
	setMsg(*otg.RsvpResourceAffinities) RsvpResourceAffinities
	// provides marshal interface
	Marshal() marshalRsvpResourceAffinities
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpResourceAffinities
	// validate validates RsvpResourceAffinities
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpResourceAffinities, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ExcludeAny returns string, set in RsvpResourceAffinities.
	ExcludeAny() string
	// SetExcludeAny assigns string provided by user to RsvpResourceAffinities
	SetExcludeAny(value string) RsvpResourceAffinities
	// HasExcludeAny checks if ExcludeAny has been set in RsvpResourceAffinities
	HasExcludeAny() bool
	// IncludeAny returns string, set in RsvpResourceAffinities.
	IncludeAny() string
	// SetIncludeAny assigns string provided by user to RsvpResourceAffinities
	SetIncludeAny(value string) RsvpResourceAffinities
	// HasIncludeAny checks if IncludeAny has been set in RsvpResourceAffinities
	HasIncludeAny() bool
	// IncludeAll returns string, set in RsvpResourceAffinities.
	IncludeAll() string
	// SetIncludeAll assigns string provided by user to RsvpResourceAffinities
	SetIncludeAll(value string) RsvpResourceAffinities
	// HasIncludeAll checks if IncludeAll has been set in RsvpResourceAffinities
	HasIncludeAll() bool
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link unacceptable.  A null set (all bits set to zero) doesn't render the link unacceptable.  The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// ExcludeAny returns a string
func (obj *rsvpResourceAffinities) ExcludeAny() string {

	return *obj.obj.ExcludeAny

}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link unacceptable.  A null set (all bits set to zero) doesn't render the link unacceptable.  The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// ExcludeAny returns a string
func (obj *rsvpResourceAffinities) HasExcludeAny() bool {
	return obj.obj.ExcludeAny != nil
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link unacceptable.  A null set (all bits set to zero) doesn't render the link unacceptable.  The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// SetExcludeAny sets the string value in the RsvpResourceAffinities object
func (obj *rsvpResourceAffinities) SetExcludeAny(value string) RsvpResourceAffinities {

	obj.obj.ExcludeAny = &value
	return obj
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// IncludeAny returns a string
func (obj *rsvpResourceAffinities) IncludeAny() string {

	return *obj.obj.IncludeAny

}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// IncludeAny returns a string
func (obj *rsvpResourceAffinities) HasIncludeAny() bool {
	return obj.obj.IncludeAny != nil
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// SetIncludeAny sets the string value in the RsvpResourceAffinities object
func (obj *rsvpResourceAffinities) SetIncludeAny(value string) RsvpResourceAffinities {

	obj.obj.IncludeAny = &value
	return obj
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel all of which must be present for a link to be acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// IncludeAll returns a string
func (obj *rsvpResourceAffinities) IncludeAll() string {

	return *obj.obj.IncludeAll

}

// A 32-bit vector representing a set of attribute filters associated with a tunnel all of which must be present for a link to be acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// IncludeAll returns a string
func (obj *rsvpResourceAffinities) HasIncludeAll() bool {
	return obj.obj.IncludeAll != nil
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel all of which must be present for a link to be acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// SetIncludeAll sets the string value in the RsvpResourceAffinities object
func (obj *rsvpResourceAffinities) SetIncludeAll(value string) RsvpResourceAffinities {

	obj.obj.IncludeAll = &value
	return obj
}

func (obj *rsvpResourceAffinities) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ExcludeAny != nil {

		if len(*obj.obj.ExcludeAny) < 0 || len(*obj.obj.ExcludeAny) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of RsvpResourceAffinities.ExcludeAny <= 8 but Got %d",
					len(*obj.obj.ExcludeAny)))
		}

	}

	if obj.obj.IncludeAny != nil {

		if len(*obj.obj.IncludeAny) < 0 || len(*obj.obj.IncludeAny) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of RsvpResourceAffinities.IncludeAny <= 8 but Got %d",
					len(*obj.obj.IncludeAny)))
		}

	}

	if obj.obj.IncludeAll != nil {

		if len(*obj.obj.IncludeAll) < 0 || len(*obj.obj.IncludeAll) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of RsvpResourceAffinities.IncludeAll <= 8 but Got %d",
					len(*obj.obj.IncludeAll)))
		}

	}

}

func (obj *rsvpResourceAffinities) setDefault() {
	if obj.obj.ExcludeAny == nil {
		obj.SetExcludeAny("0")
	}
	if obj.obj.IncludeAny == nil {
		obj.SetIncludeAny("0")
	}
	if obj.obj.IncludeAll == nil {
		obj.SetIncludeAll("0")
	}

}
