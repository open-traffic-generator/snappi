package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBmpServerIpv4UnicastPrefixStorage *****
type deviceBmpServerIpv4UnicastPrefixStorage struct {
	validation
	obj           *otg.DeviceBmpServerIpv4UnicastPrefixStorage
	marshaller    marshalDeviceBmpServerIpv4UnicastPrefixStorage
	unMarshaller  unMarshalDeviceBmpServerIpv4UnicastPrefixStorage
	discardHolder DeviceBmpServerIpv4UnicastPrefixDiscard
	storeHolder   DeviceBmpServerIpv4UnicastPrefixStore
}

func NewDeviceBmpServerIpv4UnicastPrefixStorage() DeviceBmpServerIpv4UnicastPrefixStorage {
	obj := deviceBmpServerIpv4UnicastPrefixStorage{obj: &otg.DeviceBmpServerIpv4UnicastPrefixStorage{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBmpServerIpv4UnicastPrefixStorage) msg() *otg.DeviceBmpServerIpv4UnicastPrefixStorage {
	return obj.obj
}

func (obj *deviceBmpServerIpv4UnicastPrefixStorage) setMsg(msg *otg.DeviceBmpServerIpv4UnicastPrefixStorage) DeviceBmpServerIpv4UnicastPrefixStorage {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBmpServerIpv4UnicastPrefixStorage struct {
	obj *deviceBmpServerIpv4UnicastPrefixStorage
}

type marshalDeviceBmpServerIpv4UnicastPrefixStorage interface {
	// ToProto marshals DeviceBmpServerIpv4UnicastPrefixStorage to protobuf object *otg.DeviceBmpServerIpv4UnicastPrefixStorage
	ToProto() (*otg.DeviceBmpServerIpv4UnicastPrefixStorage, error)
	// ToPbText marshals DeviceBmpServerIpv4UnicastPrefixStorage to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBmpServerIpv4UnicastPrefixStorage to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBmpServerIpv4UnicastPrefixStorage to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBmpServerIpv4UnicastPrefixStorage struct {
	obj *deviceBmpServerIpv4UnicastPrefixStorage
}

type unMarshalDeviceBmpServerIpv4UnicastPrefixStorage interface {
	// FromProto unmarshals DeviceBmpServerIpv4UnicastPrefixStorage from protobuf object *otg.DeviceBmpServerIpv4UnicastPrefixStorage
	FromProto(msg *otg.DeviceBmpServerIpv4UnicastPrefixStorage) (DeviceBmpServerIpv4UnicastPrefixStorage, error)
	// FromPbText unmarshals DeviceBmpServerIpv4UnicastPrefixStorage from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBmpServerIpv4UnicastPrefixStorage from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBmpServerIpv4UnicastPrefixStorage from JSON text
	FromJson(value string) error
}

func (obj *deviceBmpServerIpv4UnicastPrefixStorage) Marshal() marshalDeviceBmpServerIpv4UnicastPrefixStorage {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBmpServerIpv4UnicastPrefixStorage{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBmpServerIpv4UnicastPrefixStorage) Unmarshal() unMarshalDeviceBmpServerIpv4UnicastPrefixStorage {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBmpServerIpv4UnicastPrefixStorage{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBmpServerIpv4UnicastPrefixStorage) ToProto() (*otg.DeviceBmpServerIpv4UnicastPrefixStorage, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBmpServerIpv4UnicastPrefixStorage) FromProto(msg *otg.DeviceBmpServerIpv4UnicastPrefixStorage) (DeviceBmpServerIpv4UnicastPrefixStorage, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBmpServerIpv4UnicastPrefixStorage) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv4UnicastPrefixStorage) FromPbText(value string) error {
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

func (m *marshaldeviceBmpServerIpv4UnicastPrefixStorage) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv4UnicastPrefixStorage) FromYaml(value string) error {
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

func (m *marshaldeviceBmpServerIpv4UnicastPrefixStorage) ToJson() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv4UnicastPrefixStorage) FromJson(value string) error {
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

func (obj *deviceBmpServerIpv4UnicastPrefixStorage) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBmpServerIpv4UnicastPrefixStorage) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBmpServerIpv4UnicastPrefixStorage) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBmpServerIpv4UnicastPrefixStorage) Clone() (DeviceBmpServerIpv4UnicastPrefixStorage, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBmpServerIpv4UnicastPrefixStorage()
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

func (obj *deviceBmpServerIpv4UnicastPrefixStorage) setNil() {
	obj.discardHolder = nil
	obj.storeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceBmpServerIpv4UnicastPrefixStorage is optional object containing information about whether IPv4 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv4 unicast prefixes are not stored by the BMP server.
type DeviceBmpServerIpv4UnicastPrefixStorage interface {
	Validation
	// msg marshals DeviceBmpServerIpv4UnicastPrefixStorage to protobuf object *otg.DeviceBmpServerIpv4UnicastPrefixStorage
	// and doesn't set defaults
	msg() *otg.DeviceBmpServerIpv4UnicastPrefixStorage
	// setMsg unmarshals DeviceBmpServerIpv4UnicastPrefixStorage from protobuf object *otg.DeviceBmpServerIpv4UnicastPrefixStorage
	// and doesn't set defaults
	setMsg(*otg.DeviceBmpServerIpv4UnicastPrefixStorage) DeviceBmpServerIpv4UnicastPrefixStorage
	// provides marshal interface
	Marshal() marshalDeviceBmpServerIpv4UnicastPrefixStorage
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBmpServerIpv4UnicastPrefixStorage
	// validate validates DeviceBmpServerIpv4UnicastPrefixStorage
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBmpServerIpv4UnicastPrefixStorage, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns DeviceBmpServerIpv4UnicastPrefixStorageChoiceEnum, set in DeviceBmpServerIpv4UnicastPrefixStorage
	Choice() DeviceBmpServerIpv4UnicastPrefixStorageChoiceEnum
	// setChoice assigns DeviceBmpServerIpv4UnicastPrefixStorageChoiceEnum provided by user to DeviceBmpServerIpv4UnicastPrefixStorage
	setChoice(value DeviceBmpServerIpv4UnicastPrefixStorageChoiceEnum) DeviceBmpServerIpv4UnicastPrefixStorage
	// HasChoice checks if Choice has been set in DeviceBmpServerIpv4UnicastPrefixStorage
	HasChoice() bool
	// Discard returns DeviceBmpServerIpv4UnicastPrefixDiscard, set in DeviceBmpServerIpv4UnicastPrefixStorage.
	// DeviceBmpServerIpv4UnicastPrefixDiscard is the exception list can be used to specify exceptions to the specification to discard all IPv4 unicast prefixes. It is expected that when required, there would normally be a limited number of exceptions.
	Discard() DeviceBmpServerIpv4UnicastPrefixDiscard
	// SetDiscard assigns DeviceBmpServerIpv4UnicastPrefixDiscard provided by user to DeviceBmpServerIpv4UnicastPrefixStorage.
	// DeviceBmpServerIpv4UnicastPrefixDiscard is the exception list can be used to specify exceptions to the specification to discard all IPv4 unicast prefixes. It is expected that when required, there would normally be a limited number of exceptions.
	SetDiscard(value DeviceBmpServerIpv4UnicastPrefixDiscard) DeviceBmpServerIpv4UnicastPrefixStorage
	// HasDiscard checks if Discard has been set in DeviceBmpServerIpv4UnicastPrefixStorage
	HasDiscard() bool
	// Store returns DeviceBmpServerIpv4UnicastPrefixStore, set in DeviceBmpServerIpv4UnicastPrefixStorage.
	// DeviceBmpServerIpv4UnicastPrefixStore is the exception list can be used to specify exceptions to the specification to store all IPv4 unicast prefixes. It is expected that when required, there would normally be a limited number of exceptions.
	Store() DeviceBmpServerIpv4UnicastPrefixStore
	// SetStore assigns DeviceBmpServerIpv4UnicastPrefixStore provided by user to DeviceBmpServerIpv4UnicastPrefixStorage.
	// DeviceBmpServerIpv4UnicastPrefixStore is the exception list can be used to specify exceptions to the specification to store all IPv4 unicast prefixes. It is expected that when required, there would normally be a limited number of exceptions.
	SetStore(value DeviceBmpServerIpv4UnicastPrefixStore) DeviceBmpServerIpv4UnicastPrefixStorage
	// HasStore checks if Store has been set in DeviceBmpServerIpv4UnicastPrefixStorage
	HasStore() bool
	setNil()
}

type DeviceBmpServerIpv4UnicastPrefixStorageChoiceEnum string

// Enum of Choice on DeviceBmpServerIpv4UnicastPrefixStorage
var DeviceBmpServerIpv4UnicastPrefixStorageChoice = struct {
	DISCARD DeviceBmpServerIpv4UnicastPrefixStorageChoiceEnum
	STORE   DeviceBmpServerIpv4UnicastPrefixStorageChoiceEnum
}{
	DISCARD: DeviceBmpServerIpv4UnicastPrefixStorageChoiceEnum("discard"),
	STORE:   DeviceBmpServerIpv4UnicastPrefixStorageChoiceEnum("store"),
}

func (obj *deviceBmpServerIpv4UnicastPrefixStorage) Choice() DeviceBmpServerIpv4UnicastPrefixStorageChoiceEnum {
	return DeviceBmpServerIpv4UnicastPrefixStorageChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *deviceBmpServerIpv4UnicastPrefixStorage) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *deviceBmpServerIpv4UnicastPrefixStorage) setChoice(value DeviceBmpServerIpv4UnicastPrefixStorageChoiceEnum) DeviceBmpServerIpv4UnicastPrefixStorage {
	intValue, ok := otg.DeviceBmpServerIpv4UnicastPrefixStorage_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on DeviceBmpServerIpv4UnicastPrefixStorageChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.DeviceBmpServerIpv4UnicastPrefixStorage_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Store = nil
	obj.storeHolder = nil
	obj.obj.Discard = nil
	obj.discardHolder = nil

	if value == DeviceBmpServerIpv4UnicastPrefixStorageChoice.DISCARD {
		obj.obj.Discard = NewDeviceBmpServerIpv4UnicastPrefixDiscard().msg()
	}

	if value == DeviceBmpServerIpv4UnicastPrefixStorageChoice.STORE {
		obj.obj.Store = NewDeviceBmpServerIpv4UnicastPrefixStore().msg()
	}

	return obj
}

// If this option is chosen, IPv4 unicast prefixes learned from the BMP client will not be stored by the BMP server. Exception list can be used to create exceptions to allow only some subset of the learned Ipv4 unicast prefixes to be stored while the remaining learned prefixes are discarded. This can be useful when only some prefixes need to be verified amongst a large number of learned prefixes.
// Discard returns a DeviceBmpServerIpv4UnicastPrefixDiscard
func (obj *deviceBmpServerIpv4UnicastPrefixStorage) Discard() DeviceBmpServerIpv4UnicastPrefixDiscard {
	if obj.obj.Discard == nil {
		obj.setChoice(DeviceBmpServerIpv4UnicastPrefixStorageChoice.DISCARD)
	}
	if obj.discardHolder == nil {
		obj.discardHolder = &deviceBmpServerIpv4UnicastPrefixDiscard{obj: obj.obj.Discard}
	}
	return obj.discardHolder
}

// If this option is chosen, IPv4 unicast prefixes learned from the BMP client will not be stored by the BMP server. Exception list can be used to create exceptions to allow only some subset of the learned Ipv4 unicast prefixes to be stored while the remaining learned prefixes are discarded. This can be useful when only some prefixes need to be verified amongst a large number of learned prefixes.
// Discard returns a DeviceBmpServerIpv4UnicastPrefixDiscard
func (obj *deviceBmpServerIpv4UnicastPrefixStorage) HasDiscard() bool {
	return obj.obj.Discard != nil
}

// If this option is chosen, IPv4 unicast prefixes learned from the BMP client will not be stored by the BMP server. Exception list can be used to create exceptions to allow only some subset of the learned Ipv4 unicast prefixes to be stored while the remaining learned prefixes are discarded. This can be useful when only some prefixes need to be verified amongst a large number of learned prefixes.
// SetDiscard sets the DeviceBmpServerIpv4UnicastPrefixDiscard value in the DeviceBmpServerIpv4UnicastPrefixStorage object
func (obj *deviceBmpServerIpv4UnicastPrefixStorage) SetDiscard(value DeviceBmpServerIpv4UnicastPrefixDiscard) DeviceBmpServerIpv4UnicastPrefixStorage {
	obj.setChoice(DeviceBmpServerIpv4UnicastPrefixStorageChoice.DISCARD)
	obj.discardHolder = nil
	obj.obj.Discard = value.msg()

	return obj
}

// If this option is chosen, IPv4 unicast prefixes learned from the BMP client will be stored by the BMP server. Exception list can be used to create exceptions to store all the learned Ipv4 unicast prefixes except some subset of the learned IPv4 unicast prefixes. This can be useful when only some prefixes need to be verified amongst a large number of learned prefixes.
// Store returns a DeviceBmpServerIpv4UnicastPrefixStore
func (obj *deviceBmpServerIpv4UnicastPrefixStorage) Store() DeviceBmpServerIpv4UnicastPrefixStore {
	if obj.obj.Store == nil {
		obj.setChoice(DeviceBmpServerIpv4UnicastPrefixStorageChoice.STORE)
	}
	if obj.storeHolder == nil {
		obj.storeHolder = &deviceBmpServerIpv4UnicastPrefixStore{obj: obj.obj.Store}
	}
	return obj.storeHolder
}

// If this option is chosen, IPv4 unicast prefixes learned from the BMP client will be stored by the BMP server. Exception list can be used to create exceptions to store all the learned Ipv4 unicast prefixes except some subset of the learned IPv4 unicast prefixes. This can be useful when only some prefixes need to be verified amongst a large number of learned prefixes.
// Store returns a DeviceBmpServerIpv4UnicastPrefixStore
func (obj *deviceBmpServerIpv4UnicastPrefixStorage) HasStore() bool {
	return obj.obj.Store != nil
}

// If this option is chosen, IPv4 unicast prefixes learned from the BMP client will be stored by the BMP server. Exception list can be used to create exceptions to store all the learned Ipv4 unicast prefixes except some subset of the learned IPv4 unicast prefixes. This can be useful when only some prefixes need to be verified amongst a large number of learned prefixes.
// SetStore sets the DeviceBmpServerIpv4UnicastPrefixStore value in the DeviceBmpServerIpv4UnicastPrefixStorage object
func (obj *deviceBmpServerIpv4UnicastPrefixStorage) SetStore(value DeviceBmpServerIpv4UnicastPrefixStore) DeviceBmpServerIpv4UnicastPrefixStorage {
	obj.setChoice(DeviceBmpServerIpv4UnicastPrefixStorageChoice.STORE)
	obj.storeHolder = nil
	obj.obj.Store = value.msg()

	return obj
}

func (obj *deviceBmpServerIpv4UnicastPrefixStorage) validateObj(vObj *validation, set_default bool) {
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

func (obj *deviceBmpServerIpv4UnicastPrefixStorage) setDefault() {
	var choices_set int = 0
	var choice DeviceBmpServerIpv4UnicastPrefixStorageChoiceEnum

	if obj.obj.Discard != nil {
		choices_set += 1
		choice = DeviceBmpServerIpv4UnicastPrefixStorageChoice.DISCARD
	}

	if obj.obj.Store != nil {
		choices_set += 1
		choice = DeviceBmpServerIpv4UnicastPrefixStorageChoice.STORE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(DeviceBmpServerIpv4UnicastPrefixStorageChoice.DISCARD)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in DeviceBmpServerIpv4UnicastPrefixStorage")
			}
		} else {
			intVal := otg.DeviceBmpServerIpv4UnicastPrefixStorage_Choice_Enum_value[string(choice)]
			enumValue := otg.DeviceBmpServerIpv4UnicastPrefixStorage_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
