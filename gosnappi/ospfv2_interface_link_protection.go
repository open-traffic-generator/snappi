package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2InterfaceLinkProtection *****
type ospfv2InterfaceLinkProtection struct {
	validation
	obj          *otg.Ospfv2InterfaceLinkProtection
	marshaller   marshalOspfv2InterfaceLinkProtection
	unMarshaller unMarshalOspfv2InterfaceLinkProtection
}

func NewOspfv2InterfaceLinkProtection() Ospfv2InterfaceLinkProtection {
	obj := ospfv2InterfaceLinkProtection{obj: &otg.Ospfv2InterfaceLinkProtection{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2InterfaceLinkProtection) msg() *otg.Ospfv2InterfaceLinkProtection {
	return obj.obj
}

func (obj *ospfv2InterfaceLinkProtection) setMsg(msg *otg.Ospfv2InterfaceLinkProtection) Ospfv2InterfaceLinkProtection {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2InterfaceLinkProtection struct {
	obj *ospfv2InterfaceLinkProtection
}

type marshalOspfv2InterfaceLinkProtection interface {
	// ToProto marshals Ospfv2InterfaceLinkProtection to protobuf object *otg.Ospfv2InterfaceLinkProtection
	ToProto() (*otg.Ospfv2InterfaceLinkProtection, error)
	// ToPbText marshals Ospfv2InterfaceLinkProtection to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2InterfaceLinkProtection to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2InterfaceLinkProtection to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv2InterfaceLinkProtection to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv2InterfaceLinkProtection struct {
	obj *ospfv2InterfaceLinkProtection
}

type unMarshalOspfv2InterfaceLinkProtection interface {
	// FromProto unmarshals Ospfv2InterfaceLinkProtection from protobuf object *otg.Ospfv2InterfaceLinkProtection
	FromProto(msg *otg.Ospfv2InterfaceLinkProtection) (Ospfv2InterfaceLinkProtection, error)
	// FromPbText unmarshals Ospfv2InterfaceLinkProtection from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2InterfaceLinkProtection from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2InterfaceLinkProtection from JSON text
	FromJson(value string) error
}

func (obj *ospfv2InterfaceLinkProtection) Marshal() marshalOspfv2InterfaceLinkProtection {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2InterfaceLinkProtection{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2InterfaceLinkProtection) Unmarshal() unMarshalOspfv2InterfaceLinkProtection {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2InterfaceLinkProtection{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2InterfaceLinkProtection) ToProto() (*otg.Ospfv2InterfaceLinkProtection, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2InterfaceLinkProtection) FromProto(msg *otg.Ospfv2InterfaceLinkProtection) (Ospfv2InterfaceLinkProtection, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2InterfaceLinkProtection) ToPbText() (string, error) {
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

func (m *unMarshalospfv2InterfaceLinkProtection) FromPbText(value string) error {
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

func (m *marshalospfv2InterfaceLinkProtection) ToYaml() (string, error) {
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

func (m *unMarshalospfv2InterfaceLinkProtection) FromYaml(value string) error {
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

func (m *marshalospfv2InterfaceLinkProtection) ToJsonRaw() (string, error) {
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

func (m *marshalospfv2InterfaceLinkProtection) ToJson() (string, error) {
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

func (m *unMarshalospfv2InterfaceLinkProtection) FromJson(value string) error {
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

func (obj *ospfv2InterfaceLinkProtection) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2InterfaceLinkProtection) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2InterfaceLinkProtection) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2InterfaceLinkProtection) Clone() (Ospfv2InterfaceLinkProtection, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2InterfaceLinkProtection()
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

// Ospfv2InterfaceLinkProtection is optional container for the link protection sub TLV (type 20).
type Ospfv2InterfaceLinkProtection interface {
	Validation
	// msg marshals Ospfv2InterfaceLinkProtection to protobuf object *otg.Ospfv2InterfaceLinkProtection
	// and doesn't set defaults
	msg() *otg.Ospfv2InterfaceLinkProtection
	// setMsg unmarshals Ospfv2InterfaceLinkProtection from protobuf object *otg.Ospfv2InterfaceLinkProtection
	// and doesn't set defaults
	setMsg(*otg.Ospfv2InterfaceLinkProtection) Ospfv2InterfaceLinkProtection
	// provides marshal interface
	Marshal() marshalOspfv2InterfaceLinkProtection
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2InterfaceLinkProtection
	// validate validates Ospfv2InterfaceLinkProtection
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2InterfaceLinkProtection, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ExtraTraffic returns bool, set in Ospfv2InterfaceLinkProtection.
	ExtraTraffic() bool
	// SetExtraTraffic assigns bool provided by user to Ospfv2InterfaceLinkProtection
	SetExtraTraffic(value bool) Ospfv2InterfaceLinkProtection
	// HasExtraTraffic checks if ExtraTraffic has been set in Ospfv2InterfaceLinkProtection
	HasExtraTraffic() bool
	// Unprotected returns bool, set in Ospfv2InterfaceLinkProtection.
	Unprotected() bool
	// SetUnprotected assigns bool provided by user to Ospfv2InterfaceLinkProtection
	SetUnprotected(value bool) Ospfv2InterfaceLinkProtection
	// HasUnprotected checks if Unprotected has been set in Ospfv2InterfaceLinkProtection
	HasUnprotected() bool
	// Shared returns bool, set in Ospfv2InterfaceLinkProtection.
	Shared() bool
	// SetShared assigns bool provided by user to Ospfv2InterfaceLinkProtection
	SetShared(value bool) Ospfv2InterfaceLinkProtection
	// HasShared checks if Shared has been set in Ospfv2InterfaceLinkProtection
	HasShared() bool
	// Dedicated1To1 returns bool, set in Ospfv2InterfaceLinkProtection.
	Dedicated1To1() bool
	// SetDedicated1To1 assigns bool provided by user to Ospfv2InterfaceLinkProtection
	SetDedicated1To1(value bool) Ospfv2InterfaceLinkProtection
	// HasDedicated1To1 checks if Dedicated1To1 has been set in Ospfv2InterfaceLinkProtection
	HasDedicated1To1() bool
	// Dedicated1Plus1 returns bool, set in Ospfv2InterfaceLinkProtection.
	Dedicated1Plus1() bool
	// SetDedicated1Plus1 assigns bool provided by user to Ospfv2InterfaceLinkProtection
	SetDedicated1Plus1(value bool) Ospfv2InterfaceLinkProtection
	// HasDedicated1Plus1 checks if Dedicated1Plus1 has been set in Ospfv2InterfaceLinkProtection
	HasDedicated1Plus1() bool
	// Enhanced returns bool, set in Ospfv2InterfaceLinkProtection.
	Enhanced() bool
	// SetEnhanced assigns bool provided by user to Ospfv2InterfaceLinkProtection
	SetEnhanced(value bool) Ospfv2InterfaceLinkProtection
	// HasEnhanced checks if Enhanced has been set in Ospfv2InterfaceLinkProtection
	HasEnhanced() bool
	// Reserved40 returns bool, set in Ospfv2InterfaceLinkProtection.
	Reserved40() bool
	// SetReserved40 assigns bool provided by user to Ospfv2InterfaceLinkProtection
	SetReserved40(value bool) Ospfv2InterfaceLinkProtection
	// HasReserved40 checks if Reserved40 has been set in Ospfv2InterfaceLinkProtection
	HasReserved40() bool
	// Reserved80 returns bool, set in Ospfv2InterfaceLinkProtection.
	Reserved80() bool
	// SetReserved80 assigns bool provided by user to Ospfv2InterfaceLinkProtection
	SetReserved80(value bool) Ospfv2InterfaceLinkProtection
	// HasReserved80 checks if Reserved80 has been set in Ospfv2InterfaceLinkProtection
	HasReserved80() bool
}

// Enable this to protect other link or links. LSAs on a link of this type are lost
// if any of the links fail.
// ExtraTraffic returns a bool
func (obj *ospfv2InterfaceLinkProtection) ExtraTraffic() bool {

	return *obj.obj.ExtraTraffic

}

// Enable this to protect other link or links. LSAs on a link of this type are lost
// if any of the links fail.
// ExtraTraffic returns a bool
func (obj *ospfv2InterfaceLinkProtection) HasExtraTraffic() bool {
	return obj.obj.ExtraTraffic != nil
}

// Enable this to protect other link or links. LSAs on a link of this type are lost
// if any of the links fail.
// SetExtraTraffic sets the bool value in the Ospfv2InterfaceLinkProtection object
func (obj *ospfv2InterfaceLinkProtection) SetExtraTraffic(value bool) Ospfv2InterfaceLinkProtection {

	obj.obj.ExtraTraffic = &value
	return obj
}

// Enabling this signifies that there is no other link protecting this
// link. LSAs on a link of this type are lost if the link fails.
// Unprotected returns a bool
func (obj *ospfv2InterfaceLinkProtection) Unprotected() bool {

	return *obj.obj.Unprotected

}

// Enabling this signifies that there is no other link protecting this
// link. LSAs on a link of this type are lost if the link fails.
// Unprotected returns a bool
func (obj *ospfv2InterfaceLinkProtection) HasUnprotected() bool {
	return obj.obj.Unprotected != nil
}

// Enabling this signifies that there is no other link protecting this
// link. LSAs on a link of this type are lost if the link fails.
// SetUnprotected sets the bool value in the Ospfv2InterfaceLinkProtection object
func (obj *ospfv2InterfaceLinkProtection) SetUnprotected(value bool) Ospfv2InterfaceLinkProtection {

	obj.obj.Unprotected = &value
	return obj
}

// Enable this to share the Extra Traffic links between one or more
// links of type Shared.There are one or more disjoint links of type
// Extra Traffic that are protecting this link.
// Shared returns a bool
func (obj *ospfv2InterfaceLinkProtection) Shared() bool {

	return *obj.obj.Shared

}

// Enable this to share the Extra Traffic links between one or more
// links of type Shared.There are one or more disjoint links of type
// Extra Traffic that are protecting this link.
// Shared returns a bool
func (obj *ospfv2InterfaceLinkProtection) HasShared() bool {
	return obj.obj.Shared != nil
}

// Enable this to share the Extra Traffic links between one or more
// links of type Shared.There are one or more disjoint links of type
// Extra Traffic that are protecting this link.
// SetShared sets the bool value in the Ospfv2InterfaceLinkProtection object
func (obj *ospfv2InterfaceLinkProtection) SetShared(value bool) Ospfv2InterfaceLinkProtection {

	obj.obj.Shared = &value
	return obj
}

// Enabling this signifies that there is one dedicated disjoint link
// of type Extra Traffic that is protecting this link.
// Dedicated1To1 returns a bool
func (obj *ospfv2InterfaceLinkProtection) Dedicated1To1() bool {

	return *obj.obj.Dedicated_1To_1

}

// Enabling this signifies that there is one dedicated disjoint link
// of type Extra Traffic that is protecting this link.
// Dedicated1To1 returns a bool
func (obj *ospfv2InterfaceLinkProtection) HasDedicated1To1() bool {
	return obj.obj.Dedicated_1To_1 != nil
}

// Enabling this signifies that there is one dedicated disjoint link
// of type Extra Traffic that is protecting this link.
// SetDedicated1To1 sets the bool value in the Ospfv2InterfaceLinkProtection object
func (obj *ospfv2InterfaceLinkProtection) SetDedicated1To1(value bool) Ospfv2InterfaceLinkProtection {

	obj.obj.Dedicated_1To_1 = &value
	return obj
}

// Enabling this signifies that a dedicated disjoint link is protecting
// this link. However, the protecting link is not advertised in the
// link state database and is therefore not available for the routing
// of LSAs.
// Dedicated1Plus1 returns a bool
func (obj *ospfv2InterfaceLinkProtection) Dedicated1Plus1() bool {

	return *obj.obj.Dedicated_1Plus_1

}

// Enabling this signifies that a dedicated disjoint link is protecting
// this link. However, the protecting link is not advertised in the
// link state database and is therefore not available for the routing
// of LSAs.
// Dedicated1Plus1 returns a bool
func (obj *ospfv2InterfaceLinkProtection) HasDedicated1Plus1() bool {
	return obj.obj.Dedicated_1Plus_1 != nil
}

// Enabling this signifies that a dedicated disjoint link is protecting
// this link. However, the protecting link is not advertised in the
// link state database and is therefore not available for the routing
// of LSAs.
// SetDedicated1Plus1 sets the bool value in the Ospfv2InterfaceLinkProtection object
func (obj *ospfv2InterfaceLinkProtection) SetDedicated1Plus1(value bool) Ospfv2InterfaceLinkProtection {

	obj.obj.Dedicated_1Plus_1 = &value
	return obj
}

// Enabling this signifies that a protection scheme that is more
// reliable than Dedicated 1+1.
// Enhanced returns a bool
func (obj *ospfv2InterfaceLinkProtection) Enhanced() bool {

	return *obj.obj.Enhanced

}

// Enabling this signifies that a protection scheme that is more
// reliable than Dedicated 1+1.
// Enhanced returns a bool
func (obj *ospfv2InterfaceLinkProtection) HasEnhanced() bool {
	return obj.obj.Enhanced != nil
}

// Enabling this signifies that a protection scheme that is more
// reliable than Dedicated 1+1.
// SetEnhanced sets the bool value in the Ospfv2InterfaceLinkProtection object
func (obj *ospfv2InterfaceLinkProtection) SetEnhanced(value bool) Ospfv2InterfaceLinkProtection {

	obj.obj.Enhanced = &value
	return obj
}

// This is a Protection Scheme with value 0x40.
// Reserved40 returns a bool
func (obj *ospfv2InterfaceLinkProtection) Reserved40() bool {

	return *obj.obj.Reserved_40

}

// This is a Protection Scheme with value 0x40.
// Reserved40 returns a bool
func (obj *ospfv2InterfaceLinkProtection) HasReserved40() bool {
	return obj.obj.Reserved_40 != nil
}

// This is a Protection Scheme with value 0x40.
// SetReserved40 sets the bool value in the Ospfv2InterfaceLinkProtection object
func (obj *ospfv2InterfaceLinkProtection) SetReserved40(value bool) Ospfv2InterfaceLinkProtection {

	obj.obj.Reserved_40 = &value
	return obj
}

// This is a Protection Scheme with value 0x80.
// Reserved80 returns a bool
func (obj *ospfv2InterfaceLinkProtection) Reserved80() bool {

	return *obj.obj.Reserved_80

}

// This is a Protection Scheme with value 0x80.
// Reserved80 returns a bool
func (obj *ospfv2InterfaceLinkProtection) HasReserved80() bool {
	return obj.obj.Reserved_80 != nil
}

// This is a Protection Scheme with value 0x80.
// SetReserved80 sets the bool value in the Ospfv2InterfaceLinkProtection object
func (obj *ospfv2InterfaceLinkProtection) SetReserved80(value bool) Ospfv2InterfaceLinkProtection {

	obj.obj.Reserved_80 = &value
	return obj
}

func (obj *ospfv2InterfaceLinkProtection) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv2InterfaceLinkProtection) setDefault() {
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
