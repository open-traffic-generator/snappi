package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisInterfaceLinkProtection *****
type isisInterfaceLinkProtection struct {
	validation
	obj          *otg.IsisInterfaceLinkProtection
	marshaller   marshalIsisInterfaceLinkProtection
	unMarshaller unMarshalIsisInterfaceLinkProtection
}

func NewIsisInterfaceLinkProtection() IsisInterfaceLinkProtection {
	obj := isisInterfaceLinkProtection{obj: &otg.IsisInterfaceLinkProtection{}}
	obj.setDefault()
	return &obj
}

func (obj *isisInterfaceLinkProtection) msg() *otg.IsisInterfaceLinkProtection {
	return obj.obj
}

func (obj *isisInterfaceLinkProtection) setMsg(msg *otg.IsisInterfaceLinkProtection) IsisInterfaceLinkProtection {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisInterfaceLinkProtection struct {
	obj *isisInterfaceLinkProtection
}

type marshalIsisInterfaceLinkProtection interface {
	// ToProto marshals IsisInterfaceLinkProtection to protobuf object *otg.IsisInterfaceLinkProtection
	ToProto() (*otg.IsisInterfaceLinkProtection, error)
	// ToPbText marshals IsisInterfaceLinkProtection to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisInterfaceLinkProtection to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisInterfaceLinkProtection to JSON text
	ToJson() (string, error)
}

type unMarshalisisInterfaceLinkProtection struct {
	obj *isisInterfaceLinkProtection
}

type unMarshalIsisInterfaceLinkProtection interface {
	// FromProto unmarshals IsisInterfaceLinkProtection from protobuf object *otg.IsisInterfaceLinkProtection
	FromProto(msg *otg.IsisInterfaceLinkProtection) (IsisInterfaceLinkProtection, error)
	// FromPbText unmarshals IsisInterfaceLinkProtection from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisInterfaceLinkProtection from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisInterfaceLinkProtection from JSON text
	FromJson(value string) error
}

func (obj *isisInterfaceLinkProtection) Marshal() marshalIsisInterfaceLinkProtection {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisInterfaceLinkProtection{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisInterfaceLinkProtection) Unmarshal() unMarshalIsisInterfaceLinkProtection {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisInterfaceLinkProtection{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisInterfaceLinkProtection) ToProto() (*otg.IsisInterfaceLinkProtection, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisInterfaceLinkProtection) FromProto(msg *otg.IsisInterfaceLinkProtection) (IsisInterfaceLinkProtection, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisInterfaceLinkProtection) ToPbText() (string, error) {
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

func (m *unMarshalisisInterfaceLinkProtection) FromPbText(value string) error {
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

func (m *marshalisisInterfaceLinkProtection) ToYaml() (string, error) {
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

func (m *unMarshalisisInterfaceLinkProtection) FromYaml(value string) error {
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

func (m *marshalisisInterfaceLinkProtection) ToJson() (string, error) {
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

func (m *unMarshalisisInterfaceLinkProtection) FromJson(value string) error {
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

func (obj *isisInterfaceLinkProtection) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisInterfaceLinkProtection) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisInterfaceLinkProtection) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisInterfaceLinkProtection) Clone() (IsisInterfaceLinkProtection, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisInterfaceLinkProtection()
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

// IsisInterfaceLinkProtection is optional container for the link protection sub TLV (type 20).
type IsisInterfaceLinkProtection interface {
	Validation
	// msg marshals IsisInterfaceLinkProtection to protobuf object *otg.IsisInterfaceLinkProtection
	// and doesn't set defaults
	msg() *otg.IsisInterfaceLinkProtection
	// setMsg unmarshals IsisInterfaceLinkProtection from protobuf object *otg.IsisInterfaceLinkProtection
	// and doesn't set defaults
	setMsg(*otg.IsisInterfaceLinkProtection) IsisInterfaceLinkProtection
	// provides marshal interface
	Marshal() marshalIsisInterfaceLinkProtection
	// provides unmarshal interface
	Unmarshal() unMarshalIsisInterfaceLinkProtection
	// validate validates IsisInterfaceLinkProtection
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisInterfaceLinkProtection, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ExtraTraffic returns bool, set in IsisInterfaceLinkProtection.
	ExtraTraffic() bool
	// SetExtraTraffic assigns bool provided by user to IsisInterfaceLinkProtection
	SetExtraTraffic(value bool) IsisInterfaceLinkProtection
	// HasExtraTraffic checks if ExtraTraffic has been set in IsisInterfaceLinkProtection
	HasExtraTraffic() bool
	// Unprotected returns bool, set in IsisInterfaceLinkProtection.
	Unprotected() bool
	// SetUnprotected assigns bool provided by user to IsisInterfaceLinkProtection
	SetUnprotected(value bool) IsisInterfaceLinkProtection
	// HasUnprotected checks if Unprotected has been set in IsisInterfaceLinkProtection
	HasUnprotected() bool
	// Shared returns bool, set in IsisInterfaceLinkProtection.
	Shared() bool
	// SetShared assigns bool provided by user to IsisInterfaceLinkProtection
	SetShared(value bool) IsisInterfaceLinkProtection
	// HasShared checks if Shared has been set in IsisInterfaceLinkProtection
	HasShared() bool
	// Dedicated1To1 returns bool, set in IsisInterfaceLinkProtection.
	Dedicated1To1() bool
	// SetDedicated1To1 assigns bool provided by user to IsisInterfaceLinkProtection
	SetDedicated1To1(value bool) IsisInterfaceLinkProtection
	// HasDedicated1To1 checks if Dedicated1To1 has been set in IsisInterfaceLinkProtection
	HasDedicated1To1() bool
	// Dedicated1Plus1 returns bool, set in IsisInterfaceLinkProtection.
	Dedicated1Plus1() bool
	// SetDedicated1Plus1 assigns bool provided by user to IsisInterfaceLinkProtection
	SetDedicated1Plus1(value bool) IsisInterfaceLinkProtection
	// HasDedicated1Plus1 checks if Dedicated1Plus1 has been set in IsisInterfaceLinkProtection
	HasDedicated1Plus1() bool
	// Enhanced returns bool, set in IsisInterfaceLinkProtection.
	Enhanced() bool
	// SetEnhanced assigns bool provided by user to IsisInterfaceLinkProtection
	SetEnhanced(value bool) IsisInterfaceLinkProtection
	// HasEnhanced checks if Enhanced has been set in IsisInterfaceLinkProtection
	HasEnhanced() bool
	// Reserved40 returns bool, set in IsisInterfaceLinkProtection.
	Reserved40() bool
	// SetReserved40 assigns bool provided by user to IsisInterfaceLinkProtection
	SetReserved40(value bool) IsisInterfaceLinkProtection
	// HasReserved40 checks if Reserved40 has been set in IsisInterfaceLinkProtection
	HasReserved40() bool
	// Reserved80 returns bool, set in IsisInterfaceLinkProtection.
	Reserved80() bool
	// SetReserved80 assigns bool provided by user to IsisInterfaceLinkProtection
	SetReserved80(value bool) IsisInterfaceLinkProtection
	// HasReserved80 checks if Reserved80 has been set in IsisInterfaceLinkProtection
	HasReserved80() bool
}

// Enable this to protect other link or links. LSPs on a link of this type are lost
// if any of the links fail.
// ExtraTraffic returns a bool
func (obj *isisInterfaceLinkProtection) ExtraTraffic() bool {

	return *obj.obj.ExtraTraffic

}

// Enable this to protect other link or links. LSPs on a link of this type are lost
// if any of the links fail.
// ExtraTraffic returns a bool
func (obj *isisInterfaceLinkProtection) HasExtraTraffic() bool {
	return obj.obj.ExtraTraffic != nil
}

// Enable this to protect other link or links. LSPs on a link of this type are lost
// if any of the links fail.
// SetExtraTraffic sets the bool value in the IsisInterfaceLinkProtection object
func (obj *isisInterfaceLinkProtection) SetExtraTraffic(value bool) IsisInterfaceLinkProtection {

	obj.obj.ExtraTraffic = &value
	return obj
}

// Enabling this signifies that there is no other link protecting this
// link. LSPs on a link of this type are lost if the link fails.
// Unprotected returns a bool
func (obj *isisInterfaceLinkProtection) Unprotected() bool {

	return *obj.obj.Unprotected

}

// Enabling this signifies that there is no other link protecting this
// link. LSPs on a link of this type are lost if the link fails.
// Unprotected returns a bool
func (obj *isisInterfaceLinkProtection) HasUnprotected() bool {
	return obj.obj.Unprotected != nil
}

// Enabling this signifies that there is no other link protecting this
// link. LSPs on a link of this type are lost if the link fails.
// SetUnprotected sets the bool value in the IsisInterfaceLinkProtection object
func (obj *isisInterfaceLinkProtection) SetUnprotected(value bool) IsisInterfaceLinkProtection {

	obj.obj.Unprotected = &value
	return obj
}

// Enable this to share the Extra Traffic links between one or more
// links of type Shared.There are one or more disjoint links of type
// Extra Traffic that are protecting this link.
// Shared returns a bool
func (obj *isisInterfaceLinkProtection) Shared() bool {

	return *obj.obj.Shared

}

// Enable this to share the Extra Traffic links between one or more
// links of type Shared.There are one or more disjoint links of type
// Extra Traffic that are protecting this link.
// Shared returns a bool
func (obj *isisInterfaceLinkProtection) HasShared() bool {
	return obj.obj.Shared != nil
}

// Enable this to share the Extra Traffic links between one or more
// links of type Shared.There are one or more disjoint links of type
// Extra Traffic that are protecting this link.
// SetShared sets the bool value in the IsisInterfaceLinkProtection object
func (obj *isisInterfaceLinkProtection) SetShared(value bool) IsisInterfaceLinkProtection {

	obj.obj.Shared = &value
	return obj
}

// Enabling this signifies that there is one dedicated disjoint link
// of type Extra Traffic that is protecting this link.
// Dedicated1To1 returns a bool
func (obj *isisInterfaceLinkProtection) Dedicated1To1() bool {

	return *obj.obj.Dedicated_1To_1

}

// Enabling this signifies that there is one dedicated disjoint link
// of type Extra Traffic that is protecting this link.
// Dedicated1To1 returns a bool
func (obj *isisInterfaceLinkProtection) HasDedicated1To1() bool {
	return obj.obj.Dedicated_1To_1 != nil
}

// Enabling this signifies that there is one dedicated disjoint link
// of type Extra Traffic that is protecting this link.
// SetDedicated1To1 sets the bool value in the IsisInterfaceLinkProtection object
func (obj *isisInterfaceLinkProtection) SetDedicated1To1(value bool) IsisInterfaceLinkProtection {

	obj.obj.Dedicated_1To_1 = &value
	return obj
}

// Enabling this signifies that a dedicated disjoint link is protecting
// this link. However, the protecting link is not advertised in the
// link state database and is therefore not available for the routing
// of LSPs.
// Dedicated1Plus1 returns a bool
func (obj *isisInterfaceLinkProtection) Dedicated1Plus1() bool {

	return *obj.obj.Dedicated_1Plus_1

}

// Enabling this signifies that a dedicated disjoint link is protecting
// this link. However, the protecting link is not advertised in the
// link state database and is therefore not available for the routing
// of LSPs.
// Dedicated1Plus1 returns a bool
func (obj *isisInterfaceLinkProtection) HasDedicated1Plus1() bool {
	return obj.obj.Dedicated_1Plus_1 != nil
}

// Enabling this signifies that a dedicated disjoint link is protecting
// this link. However, the protecting link is not advertised in the
// link state database and is therefore not available for the routing
// of LSPs.
// SetDedicated1Plus1 sets the bool value in the IsisInterfaceLinkProtection object
func (obj *isisInterfaceLinkProtection) SetDedicated1Plus1(value bool) IsisInterfaceLinkProtection {

	obj.obj.Dedicated_1Plus_1 = &value
	return obj
}

// Enabling this signifies that a protection scheme that is more
// reliable than Dedicated 1+1.
// Enhanced returns a bool
func (obj *isisInterfaceLinkProtection) Enhanced() bool {

	return *obj.obj.Enhanced

}

// Enabling this signifies that a protection scheme that is more
// reliable than Dedicated 1+1.
// Enhanced returns a bool
func (obj *isisInterfaceLinkProtection) HasEnhanced() bool {
	return obj.obj.Enhanced != nil
}

// Enabling this signifies that a protection scheme that is more
// reliable than Dedicated 1+1.
// SetEnhanced sets the bool value in the IsisInterfaceLinkProtection object
func (obj *isisInterfaceLinkProtection) SetEnhanced(value bool) IsisInterfaceLinkProtection {

	obj.obj.Enhanced = &value
	return obj
}

// This is a Protection Scheme with value 0x40.
// Reserved40 returns a bool
func (obj *isisInterfaceLinkProtection) Reserved40() bool {

	return *obj.obj.Reserved_40

}

// This is a Protection Scheme with value 0x40.
// Reserved40 returns a bool
func (obj *isisInterfaceLinkProtection) HasReserved40() bool {
	return obj.obj.Reserved_40 != nil
}

// This is a Protection Scheme with value 0x40.
// SetReserved40 sets the bool value in the IsisInterfaceLinkProtection object
func (obj *isisInterfaceLinkProtection) SetReserved40(value bool) IsisInterfaceLinkProtection {

	obj.obj.Reserved_40 = &value
	return obj
}

// This is a Protection Scheme with value 0x80.
// Reserved80 returns a bool
func (obj *isisInterfaceLinkProtection) Reserved80() bool {

	return *obj.obj.Reserved_80

}

// This is a Protection Scheme with value 0x80.
// Reserved80 returns a bool
func (obj *isisInterfaceLinkProtection) HasReserved80() bool {
	return obj.obj.Reserved_80 != nil
}

// This is a Protection Scheme with value 0x80.
// SetReserved80 sets the bool value in the IsisInterfaceLinkProtection object
func (obj *isisInterfaceLinkProtection) SetReserved80(value bool) IsisInterfaceLinkProtection {

	obj.obj.Reserved_80 = &value
	return obj
}

func (obj *isisInterfaceLinkProtection) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisInterfaceLinkProtection) setDefault() {
	if obj.obj.ExtraTraffic == nil {
		obj.SetExtraTraffic(false)
	}
	if obj.obj.Unprotected == nil {
		obj.SetUnprotected(false)
	}
	if obj.obj.Shared == nil {
		obj.SetShared(false)
	}
	if obj.obj.Dedicated_1To_1 == nil {
		obj.SetDedicated1To1(false)
	}
	if obj.obj.Dedicated_1Plus_1 == nil {
		obj.SetDedicated1Plus1(false)
	}
	if obj.obj.Enhanced == nil {
		obj.SetEnhanced(false)
	}
	if obj.obj.Reserved_40 == nil {
		obj.SetReserved40(false)
	}
	if obj.obj.Reserved_80 == nil {
		obj.SetReserved80(false)
	}

}
