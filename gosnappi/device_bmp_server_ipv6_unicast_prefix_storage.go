package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBmpServerIpv6UnicastPrefixStorage *****
type deviceBmpServerIpv6UnicastPrefixStorage struct {
	validation
	obj           *otg.DeviceBmpServerIpv6UnicastPrefixStorage
	marshaller    marshalDeviceBmpServerIpv6UnicastPrefixStorage
	unMarshaller  unMarshalDeviceBmpServerIpv6UnicastPrefixStorage
	discardHolder DeviceBmpServerIpv6UnicastPrefixDiscard
	storeHolder   DeviceBmpServerIpv6UnicastPrefixStore
}

func NewDeviceBmpServerIpv6UnicastPrefixStorage() DeviceBmpServerIpv6UnicastPrefixStorage {
	obj := deviceBmpServerIpv6UnicastPrefixStorage{obj: &otg.DeviceBmpServerIpv6UnicastPrefixStorage{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBmpServerIpv6UnicastPrefixStorage) msg() *otg.DeviceBmpServerIpv6UnicastPrefixStorage {
	return obj.obj
}

func (obj *deviceBmpServerIpv6UnicastPrefixStorage) setMsg(msg *otg.DeviceBmpServerIpv6UnicastPrefixStorage) DeviceBmpServerIpv6UnicastPrefixStorage {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBmpServerIpv6UnicastPrefixStorage struct {
	obj *deviceBmpServerIpv6UnicastPrefixStorage
}

type marshalDeviceBmpServerIpv6UnicastPrefixStorage interface {
	// ToProto marshals DeviceBmpServerIpv6UnicastPrefixStorage to protobuf object *otg.DeviceBmpServerIpv6UnicastPrefixStorage
	ToProto() (*otg.DeviceBmpServerIpv6UnicastPrefixStorage, error)
	// ToPbText marshals DeviceBmpServerIpv6UnicastPrefixStorage to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBmpServerIpv6UnicastPrefixStorage to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBmpServerIpv6UnicastPrefixStorage to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBmpServerIpv6UnicastPrefixStorage struct {
	obj *deviceBmpServerIpv6UnicastPrefixStorage
}

type unMarshalDeviceBmpServerIpv6UnicastPrefixStorage interface {
	// FromProto unmarshals DeviceBmpServerIpv6UnicastPrefixStorage from protobuf object *otg.DeviceBmpServerIpv6UnicastPrefixStorage
	FromProto(msg *otg.DeviceBmpServerIpv6UnicastPrefixStorage) (DeviceBmpServerIpv6UnicastPrefixStorage, error)
	// FromPbText unmarshals DeviceBmpServerIpv6UnicastPrefixStorage from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBmpServerIpv6UnicastPrefixStorage from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBmpServerIpv6UnicastPrefixStorage from JSON text
	FromJson(value string) error
}

func (obj *deviceBmpServerIpv6UnicastPrefixStorage) Marshal() marshalDeviceBmpServerIpv6UnicastPrefixStorage {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBmpServerIpv6UnicastPrefixStorage{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBmpServerIpv6UnicastPrefixStorage) Unmarshal() unMarshalDeviceBmpServerIpv6UnicastPrefixStorage {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBmpServerIpv6UnicastPrefixStorage{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBmpServerIpv6UnicastPrefixStorage) ToProto() (*otg.DeviceBmpServerIpv6UnicastPrefixStorage, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBmpServerIpv6UnicastPrefixStorage) FromProto(msg *otg.DeviceBmpServerIpv6UnicastPrefixStorage) (DeviceBmpServerIpv6UnicastPrefixStorage, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBmpServerIpv6UnicastPrefixStorage) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv6UnicastPrefixStorage) FromPbText(value string) error {
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

func (m *marshaldeviceBmpServerIpv6UnicastPrefixStorage) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv6UnicastPrefixStorage) FromYaml(value string) error {
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

func (m *marshaldeviceBmpServerIpv6UnicastPrefixStorage) ToJson() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv6UnicastPrefixStorage) FromJson(value string) error {
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

func (obj *deviceBmpServerIpv6UnicastPrefixStorage) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBmpServerIpv6UnicastPrefixStorage) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBmpServerIpv6UnicastPrefixStorage) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBmpServerIpv6UnicastPrefixStorage) Clone() (DeviceBmpServerIpv6UnicastPrefixStorage, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBmpServerIpv6UnicastPrefixStorage()
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

func (obj *deviceBmpServerIpv6UnicastPrefixStorage) setNil() {
	obj.discardHolder = nil
	obj.storeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceBmpServerIpv6UnicastPrefixStorage is optional object containing information about whether IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv6 unicast prefixes are not stored by the BMP server.
type DeviceBmpServerIpv6UnicastPrefixStorage interface {
	Validation
	// msg marshals DeviceBmpServerIpv6UnicastPrefixStorage to protobuf object *otg.DeviceBmpServerIpv6UnicastPrefixStorage
	// and doesn't set defaults
	msg() *otg.DeviceBmpServerIpv6UnicastPrefixStorage
	// setMsg unmarshals DeviceBmpServerIpv6UnicastPrefixStorage from protobuf object *otg.DeviceBmpServerIpv6UnicastPrefixStorage
	// and doesn't set defaults
	setMsg(*otg.DeviceBmpServerIpv6UnicastPrefixStorage) DeviceBmpServerIpv6UnicastPrefixStorage
	// provides marshal interface
	Marshal() marshalDeviceBmpServerIpv6UnicastPrefixStorage
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBmpServerIpv6UnicastPrefixStorage
	// validate validates DeviceBmpServerIpv6UnicastPrefixStorage
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBmpServerIpv6UnicastPrefixStorage, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns DeviceBmpServerIpv6UnicastPrefixStorageChoiceEnum, set in DeviceBmpServerIpv6UnicastPrefixStorage
	Choice() DeviceBmpServerIpv6UnicastPrefixStorageChoiceEnum
	// setChoice assigns DeviceBmpServerIpv6UnicastPrefixStorageChoiceEnum provided by user to DeviceBmpServerIpv6UnicastPrefixStorage
	setChoice(value DeviceBmpServerIpv6UnicastPrefixStorageChoiceEnum) DeviceBmpServerIpv6UnicastPrefixStorage
	// HasChoice checks if Choice has been set in DeviceBmpServerIpv6UnicastPrefixStorage
	HasChoice() bool
	// Discard returns DeviceBmpServerIpv6UnicastPrefixDiscard, set in DeviceBmpServerIpv6UnicastPrefixStorage.
	// DeviceBmpServerIpv6UnicastPrefixDiscard is the exception list can be used to specify exceptions to the specification to discard all IPv6 unicast prefixes. It is expected that when required, there would normally be a limited number of exceptions.
	Discard() DeviceBmpServerIpv6UnicastPrefixDiscard
	// SetDiscard assigns DeviceBmpServerIpv6UnicastPrefixDiscard provided by user to DeviceBmpServerIpv6UnicastPrefixStorage.
	// DeviceBmpServerIpv6UnicastPrefixDiscard is the exception list can be used to specify exceptions to the specification to discard all IPv6 unicast prefixes. It is expected that when required, there would normally be a limited number of exceptions.
	SetDiscard(value DeviceBmpServerIpv6UnicastPrefixDiscard) DeviceBmpServerIpv6UnicastPrefixStorage
	// HasDiscard checks if Discard has been set in DeviceBmpServerIpv6UnicastPrefixStorage
	HasDiscard() bool
	// Store returns DeviceBmpServerIpv6UnicastPrefixStore, set in DeviceBmpServerIpv6UnicastPrefixStorage.
	// DeviceBmpServerIpv6UnicastPrefixStore is the exception list can be used to specify exceptions to the specification to store all IPv6 unicast prefixes. It is expected that when required, there would normally be a limited number of exceptions.
	Store() DeviceBmpServerIpv6UnicastPrefixStore
	// SetStore assigns DeviceBmpServerIpv6UnicastPrefixStore provided by user to DeviceBmpServerIpv6UnicastPrefixStorage.
	// DeviceBmpServerIpv6UnicastPrefixStore is the exception list can be used to specify exceptions to the specification to store all IPv6 unicast prefixes. It is expected that when required, there would normally be a limited number of exceptions.
	SetStore(value DeviceBmpServerIpv6UnicastPrefixStore) DeviceBmpServerIpv6UnicastPrefixStorage
	// HasStore checks if Store has been set in DeviceBmpServerIpv6UnicastPrefixStorage
	HasStore() bool
	setNil()
}

type DeviceBmpServerIpv6UnicastPrefixStorageChoiceEnum string

// Enum of Choice on DeviceBmpServerIpv6UnicastPrefixStorage
var DeviceBmpServerIpv6UnicastPrefixStorageChoice = struct {
	DISCARD DeviceBmpServerIpv6UnicastPrefixStorageChoiceEnum
	STORE   DeviceBmpServerIpv6UnicastPrefixStorageChoiceEnum
}{
	DISCARD: DeviceBmpServerIpv6UnicastPrefixStorageChoiceEnum("discard"),
	STORE:   DeviceBmpServerIpv6UnicastPrefixStorageChoiceEnum("store"),
}

func (obj *deviceBmpServerIpv6UnicastPrefixStorage) Choice() DeviceBmpServerIpv6UnicastPrefixStorageChoiceEnum {
	return DeviceBmpServerIpv6UnicastPrefixStorageChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *deviceBmpServerIpv6UnicastPrefixStorage) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *deviceBmpServerIpv6UnicastPrefixStorage) setChoice(value DeviceBmpServerIpv6UnicastPrefixStorageChoiceEnum) DeviceBmpServerIpv6UnicastPrefixStorage {
	intValue, ok := otg.DeviceBmpServerIpv6UnicastPrefixStorage_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on DeviceBmpServerIpv6UnicastPrefixStorageChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.DeviceBmpServerIpv6UnicastPrefixStorage_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Store = nil
	obj.storeHolder = nil
	obj.obj.Discard = nil
	obj.discardHolder = nil

	if value == DeviceBmpServerIpv6UnicastPrefixStorageChoice.DISCARD {
		obj.obj.Discard = NewDeviceBmpServerIpv6UnicastPrefixDiscard().msg()
	}

	if value == DeviceBmpServerIpv6UnicastPrefixStorageChoice.STORE {
		obj.obj.Store = NewDeviceBmpServerIpv6UnicastPrefixStore().msg()
	}

	return obj
}

// If this option is chosen, IPv6 unicast prefixes learned from the BMP client will not be stored by the BMP server. Exception list can be used to create exceptions to allow only some subset of the learned IPv6 unicast prefixes to be stored while the remaining learned prefixes are discarded. This can be useful when only some prefixes need to be verified amongst a large number of learned prefixes.
// Discard returns a DeviceBmpServerIpv6UnicastPrefixDiscard
func (obj *deviceBmpServerIpv6UnicastPrefixStorage) Discard() DeviceBmpServerIpv6UnicastPrefixDiscard {
	if obj.obj.Discard == nil {
		obj.setChoice(DeviceBmpServerIpv6UnicastPrefixStorageChoice.DISCARD)
	}
	if obj.discardHolder == nil {
		obj.discardHolder = &deviceBmpServerIpv6UnicastPrefixDiscard{obj: obj.obj.Discard}
	}
	return obj.discardHolder
}

// If this option is chosen, IPv6 unicast prefixes learned from the BMP client will not be stored by the BMP server. Exception list can be used to create exceptions to allow only some subset of the learned IPv6 unicast prefixes to be stored while the remaining learned prefixes are discarded. This can be useful when only some prefixes need to be verified amongst a large number of learned prefixes.
// Discard returns a DeviceBmpServerIpv6UnicastPrefixDiscard
func (obj *deviceBmpServerIpv6UnicastPrefixStorage) HasDiscard() bool {
	return obj.obj.Discard != nil
}

// If this option is chosen, IPv6 unicast prefixes learned from the BMP client will not be stored by the BMP server. Exception list can be used to create exceptions to allow only some subset of the learned IPv6 unicast prefixes to be stored while the remaining learned prefixes are discarded. This can be useful when only some prefixes need to be verified amongst a large number of learned prefixes.
// SetDiscard sets the DeviceBmpServerIpv6UnicastPrefixDiscard value in the DeviceBmpServerIpv6UnicastPrefixStorage object
func (obj *deviceBmpServerIpv6UnicastPrefixStorage) SetDiscard(value DeviceBmpServerIpv6UnicastPrefixDiscard) DeviceBmpServerIpv6UnicastPrefixStorage {
	obj.setChoice(DeviceBmpServerIpv6UnicastPrefixStorageChoice.DISCARD)
	obj.discardHolder = nil
	obj.obj.Discard = value.msg()

	return obj
}

// If this option is chosen, IPv6 unicast prefixes learned from the BMP client will be stored by the BMP server. Exception list can be used to create exceptions to store all the learned Ipv6 unicast prefixes except some subset of the learned IPv6 unicast prefixes. This can be useful when only some prefixes need to be verified amongst a large number of learned prefixes.
// Store returns a DeviceBmpServerIpv6UnicastPrefixStore
func (obj *deviceBmpServerIpv6UnicastPrefixStorage) Store() DeviceBmpServerIpv6UnicastPrefixStore {
	if obj.obj.Store == nil {
		obj.setChoice(DeviceBmpServerIpv6UnicastPrefixStorageChoice.STORE)
	}
	if obj.storeHolder == nil {
		obj.storeHolder = &deviceBmpServerIpv6UnicastPrefixStore{obj: obj.obj.Store}
	}
	return obj.storeHolder
}

// If this option is chosen, IPv6 unicast prefixes learned from the BMP client will be stored by the BMP server. Exception list can be used to create exceptions to store all the learned Ipv6 unicast prefixes except some subset of the learned IPv6 unicast prefixes. This can be useful when only some prefixes need to be verified amongst a large number of learned prefixes.
// Store returns a DeviceBmpServerIpv6UnicastPrefixStore
func (obj *deviceBmpServerIpv6UnicastPrefixStorage) HasStore() bool {
	return obj.obj.Store != nil
}

// If this option is chosen, IPv6 unicast prefixes learned from the BMP client will be stored by the BMP server. Exception list can be used to create exceptions to store all the learned Ipv6 unicast prefixes except some subset of the learned IPv6 unicast prefixes. This can be useful when only some prefixes need to be verified amongst a large number of learned prefixes.
// SetStore sets the DeviceBmpServerIpv6UnicastPrefixStore value in the DeviceBmpServerIpv6UnicastPrefixStorage object
func (obj *deviceBmpServerIpv6UnicastPrefixStorage) SetStore(value DeviceBmpServerIpv6UnicastPrefixStore) DeviceBmpServerIpv6UnicastPrefixStorage {
	obj.setChoice(DeviceBmpServerIpv6UnicastPrefixStorageChoice.STORE)
	obj.storeHolder = nil
	obj.obj.Store = value.msg()

	return obj
}

func (obj *deviceBmpServerIpv6UnicastPrefixStorage) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Discard != nil {

		obj.Discard().validateObj(vObj, set_default)
	}

	if obj.obj.Store != nil {

		obj.Store().validateObj(vObj, set_default)
	}

}

func (obj *deviceBmpServerIpv6UnicastPrefixStorage) setDefault() {
	var choices_set int = 0
	var choice DeviceBmpServerIpv6UnicastPrefixStorageChoiceEnum

	if obj.obj.Discard != nil {
		choices_set += 1
		choice = DeviceBmpServerIpv6UnicastPrefixStorageChoice.DISCARD
	}

	if obj.obj.Store != nil {
		choices_set += 1
		choice = DeviceBmpServerIpv6UnicastPrefixStorageChoice.STORE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(DeviceBmpServerIpv6UnicastPrefixStorageChoice.DISCARD)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in DeviceBmpServerIpv6UnicastPrefixStorage")
			}
		} else {
			intVal := otg.DeviceBmpServerIpv6UnicastPrefixStorage_Choice_Enum_value[string(choice)]
			enumValue := otg.DeviceBmpServerIpv6UnicastPrefixStorage_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
