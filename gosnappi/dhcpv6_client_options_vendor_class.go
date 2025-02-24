package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsVendorClass *****
type dhcpv6ClientOptionsVendorClass struct {
	validation
	obj                          *otg.Dhcpv6ClientOptionsVendorClass
	marshaller                   marshalDhcpv6ClientOptionsVendorClass
	unMarshaller                 unMarshalDhcpv6ClientOptionsVendorClass
	associatedDhcpMessagesHolder Dhcpv6ClientOptionsIncludedMessages
}

func NewDhcpv6ClientOptionsVendorClass() Dhcpv6ClientOptionsVendorClass {
	obj := dhcpv6ClientOptionsVendorClass{obj: &otg.Dhcpv6ClientOptionsVendorClass{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsVendorClass) msg() *otg.Dhcpv6ClientOptionsVendorClass {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsVendorClass) setMsg(msg *otg.Dhcpv6ClientOptionsVendorClass) Dhcpv6ClientOptionsVendorClass {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsVendorClass struct {
	obj *dhcpv6ClientOptionsVendorClass
}

type marshalDhcpv6ClientOptionsVendorClass interface {
	// ToProto marshals Dhcpv6ClientOptionsVendorClass to protobuf object *otg.Dhcpv6ClientOptionsVendorClass
	ToProto() (*otg.Dhcpv6ClientOptionsVendorClass, error)
	// ToPbText marshals Dhcpv6ClientOptionsVendorClass to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsVendorClass to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsVendorClass to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ClientOptionsVendorClass to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ClientOptionsVendorClass struct {
	obj *dhcpv6ClientOptionsVendorClass
}

type unMarshalDhcpv6ClientOptionsVendorClass interface {
	// FromProto unmarshals Dhcpv6ClientOptionsVendorClass from protobuf object *otg.Dhcpv6ClientOptionsVendorClass
	FromProto(msg *otg.Dhcpv6ClientOptionsVendorClass) (Dhcpv6ClientOptionsVendorClass, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsVendorClass from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsVendorClass from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsVendorClass from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsVendorClass) Marshal() marshalDhcpv6ClientOptionsVendorClass {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsVendorClass{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsVendorClass) Unmarshal() unMarshalDhcpv6ClientOptionsVendorClass {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsVendorClass{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsVendorClass) ToProto() (*otg.Dhcpv6ClientOptionsVendorClass, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsVendorClass) FromProto(msg *otg.Dhcpv6ClientOptionsVendorClass) (Dhcpv6ClientOptionsVendorClass, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsVendorClass) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsVendorClass) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsVendorClass) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsVendorClass) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsVendorClass) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ClientOptionsVendorClass) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsVendorClass) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsVendorClass) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsVendorClass) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsVendorClass) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsVendorClass) Clone() (Dhcpv6ClientOptionsVendorClass, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsVendorClass()
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

func (obj *dhcpv6ClientOptionsVendorClass) setNil() {
	obj.associatedDhcpMessagesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ClientOptionsVendorClass is this option is used by a client to identify the vendor that manufactured the hardware on which the client is running. The option code is 16.
type Dhcpv6ClientOptionsVendorClass interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsVendorClass to protobuf object *otg.Dhcpv6ClientOptionsVendorClass
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsVendorClass
	// setMsg unmarshals Dhcpv6ClientOptionsVendorClass from protobuf object *otg.Dhcpv6ClientOptionsVendorClass
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsVendorClass) Dhcpv6ClientOptionsVendorClass
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsVendorClass
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsVendorClass
	// validate validates Dhcpv6ClientOptionsVendorClass
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsVendorClass, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EnterpriseNumber returns uint32, set in Dhcpv6ClientOptionsVendorClass.
	EnterpriseNumber() uint32
	// SetEnterpriseNumber assigns uint32 provided by user to Dhcpv6ClientOptionsVendorClass
	SetEnterpriseNumber(value uint32) Dhcpv6ClientOptionsVendorClass
	// ClassData returns []string, set in Dhcpv6ClientOptionsVendorClass.
	ClassData() []string
	// SetClassData assigns []string provided by user to Dhcpv6ClientOptionsVendorClass
	SetClassData(value []string) Dhcpv6ClientOptionsVendorClass
	// AssociatedDhcpMessages returns Dhcpv6ClientOptionsIncludedMessages, set in Dhcpv6ClientOptionsVendorClass.
	// Dhcpv6ClientOptionsIncludedMessages is the dhcpv6 client messages where the option will be included. If all is selected the selected option will be added in the all the Dhcpv6 client messages, else based on the selection in particular Dhcpv6 client messages the option will be included.
	AssociatedDhcpMessages() Dhcpv6ClientOptionsIncludedMessages
	// SetAssociatedDhcpMessages assigns Dhcpv6ClientOptionsIncludedMessages provided by user to Dhcpv6ClientOptionsVendorClass.
	// Dhcpv6ClientOptionsIncludedMessages is the dhcpv6 client messages where the option will be included. If all is selected the selected option will be added in the all the Dhcpv6 client messages, else based on the selection in particular Dhcpv6 client messages the option will be included.
	SetAssociatedDhcpMessages(value Dhcpv6ClientOptionsIncludedMessages) Dhcpv6ClientOptionsVendorClass
	setNil()
}

// The vendor's registered Enterprise Number as registered with IANA.
// EnterpriseNumber returns a uint32
func (obj *dhcpv6ClientOptionsVendorClass) EnterpriseNumber() uint32 {

	return *obj.obj.EnterpriseNumber

}

// The vendor's registered Enterprise Number as registered with IANA.
// SetEnterpriseNumber sets the uint32 value in the Dhcpv6ClientOptionsVendorClass object
func (obj *dhcpv6ClientOptionsVendorClass) SetEnterpriseNumber(value uint32) Dhcpv6ClientOptionsVendorClass {

	obj.obj.EnterpriseNumber = &value
	return obj
}

// The opaque data representing the hardware configuration of the host on which the client is running. Examples of  class data instances might include the version of the operating system the client is running or the amount of memory  installed on the client.
// ClassData returns a []string
func (obj *dhcpv6ClientOptionsVendorClass) ClassData() []string {
	if obj.obj.ClassData == nil {
		obj.obj.ClassData = make([]string, 0)
	}
	return obj.obj.ClassData
}

// The opaque data representing the hardware configuration of the host on which the client is running. Examples of  class data instances might include the version of the operating system the client is running or the amount of memory  installed on the client.
// SetClassData sets the []string value in the Dhcpv6ClientOptionsVendorClass object
func (obj *dhcpv6ClientOptionsVendorClass) SetClassData(value []string) Dhcpv6ClientOptionsVendorClass {

	if obj.obj.ClassData == nil {
		obj.obj.ClassData = make([]string, 0)
	}
	obj.obj.ClassData = value

	return obj
}

// The dhcpv6 client messages where this option is included.
// AssociatedDhcpMessages returns a Dhcpv6ClientOptionsIncludedMessages
func (obj *dhcpv6ClientOptionsVendorClass) AssociatedDhcpMessages() Dhcpv6ClientOptionsIncludedMessages {
	if obj.obj.AssociatedDhcpMessages == nil {
		obj.obj.AssociatedDhcpMessages = NewDhcpv6ClientOptionsIncludedMessages().msg()
	}
	if obj.associatedDhcpMessagesHolder == nil {
		obj.associatedDhcpMessagesHolder = &dhcpv6ClientOptionsIncludedMessages{obj: obj.obj.AssociatedDhcpMessages}
	}
	return obj.associatedDhcpMessagesHolder
}

// The dhcpv6 client messages where this option is included.
// SetAssociatedDhcpMessages sets the Dhcpv6ClientOptionsIncludedMessages value in the Dhcpv6ClientOptionsVendorClass object
func (obj *dhcpv6ClientOptionsVendorClass) SetAssociatedDhcpMessages(value Dhcpv6ClientOptionsIncludedMessages) Dhcpv6ClientOptionsVendorClass {

	obj.associatedDhcpMessagesHolder = nil
	obj.obj.AssociatedDhcpMessages = value.msg()

	return obj
}

func (obj *dhcpv6ClientOptionsVendorClass) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// EnterpriseNumber is required
	if obj.obj.EnterpriseNumber == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EnterpriseNumber is required field on interface Dhcpv6ClientOptionsVendorClass")
	}
	if obj.obj.EnterpriseNumber != nil {

		if *obj.obj.EnterpriseNumber > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6ClientOptionsVendorClass.EnterpriseNumber <= 4294967295 but Got %d", *obj.obj.EnterpriseNumber))
		}

	}

	// AssociatedDhcpMessages is required
	if obj.obj.AssociatedDhcpMessages == nil {
		vObj.validationErrors = append(vObj.validationErrors, "AssociatedDhcpMessages is required field on interface Dhcpv6ClientOptionsVendorClass")
	}

	if obj.obj.AssociatedDhcpMessages != nil {

		obj.AssociatedDhcpMessages().validateObj(vObj, set_default)
	}

}

func (obj *dhcpv6ClientOptionsVendorClass) setDefault() {

}
