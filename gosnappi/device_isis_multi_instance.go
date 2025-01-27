package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceIsisMultiInstance *****
type deviceIsisMultiInstance struct {
	validation
	obj          *otg.DeviceIsisMultiInstance
	marshaller   marshalDeviceIsisMultiInstance
	unMarshaller unMarshalDeviceIsisMultiInstance
}

func NewDeviceIsisMultiInstance() DeviceIsisMultiInstance {
	obj := deviceIsisMultiInstance{obj: &otg.DeviceIsisMultiInstance{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceIsisMultiInstance) msg() *otg.DeviceIsisMultiInstance {
	return obj.obj
}

func (obj *deviceIsisMultiInstance) setMsg(msg *otg.DeviceIsisMultiInstance) DeviceIsisMultiInstance {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceIsisMultiInstance struct {
	obj *deviceIsisMultiInstance
}

type marshalDeviceIsisMultiInstance interface {
	// ToProto marshals DeviceIsisMultiInstance to protobuf object *otg.DeviceIsisMultiInstance
	ToProto() (*otg.DeviceIsisMultiInstance, error)
	// ToPbText marshals DeviceIsisMultiInstance to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceIsisMultiInstance to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceIsisMultiInstance to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals DeviceIsisMultiInstance to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldeviceIsisMultiInstance struct {
	obj *deviceIsisMultiInstance
}

type unMarshalDeviceIsisMultiInstance interface {
	// FromProto unmarshals DeviceIsisMultiInstance from protobuf object *otg.DeviceIsisMultiInstance
	FromProto(msg *otg.DeviceIsisMultiInstance) (DeviceIsisMultiInstance, error)
	// FromPbText unmarshals DeviceIsisMultiInstance from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceIsisMultiInstance from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceIsisMultiInstance from JSON text
	FromJson(value string) error
}

func (obj *deviceIsisMultiInstance) Marshal() marshalDeviceIsisMultiInstance {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceIsisMultiInstance{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceIsisMultiInstance) Unmarshal() unMarshalDeviceIsisMultiInstance {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceIsisMultiInstance{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceIsisMultiInstance) ToProto() (*otg.DeviceIsisMultiInstance, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceIsisMultiInstance) FromProto(msg *otg.DeviceIsisMultiInstance) (DeviceIsisMultiInstance, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceIsisMultiInstance) ToPbText() (string, error) {
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

func (m *unMarshaldeviceIsisMultiInstance) FromPbText(value string) error {
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

func (m *marshaldeviceIsisMultiInstance) ToYaml() (string, error) {
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

func (m *unMarshaldeviceIsisMultiInstance) FromYaml(value string) error {
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

func (m *marshaldeviceIsisMultiInstance) ToJsonRaw() (string, error) {
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

func (m *marshaldeviceIsisMultiInstance) ToJson() (string, error) {
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

func (m *unMarshaldeviceIsisMultiInstance) FromJson(value string) error {
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

func (obj *deviceIsisMultiInstance) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceIsisMultiInstance) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceIsisMultiInstance) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceIsisMultiInstance) Clone() (DeviceIsisMultiInstance, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceIsisMultiInstance()
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

// DeviceIsisMultiInstance is this container properties of an Multi-Instance-capable router (MI-RTR).
type DeviceIsisMultiInstance interface {
	Validation
	// msg marshals DeviceIsisMultiInstance to protobuf object *otg.DeviceIsisMultiInstance
	// and doesn't set defaults
	msg() *otg.DeviceIsisMultiInstance
	// setMsg unmarshals DeviceIsisMultiInstance from protobuf object *otg.DeviceIsisMultiInstance
	// and doesn't set defaults
	setMsg(*otg.DeviceIsisMultiInstance) DeviceIsisMultiInstance
	// provides marshal interface
	Marshal() marshalDeviceIsisMultiInstance
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceIsisMultiInstance
	// validate validates DeviceIsisMultiInstance
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceIsisMultiInstance, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Iid returns uint32, set in DeviceIsisMultiInstance.
	Iid() uint32
	// SetIid assigns uint32 provided by user to DeviceIsisMultiInstance
	SetIid(value uint32) DeviceIsisMultiInstance
	// HasIid checks if Iid has been set in DeviceIsisMultiInstance
	HasIid() bool
	// Itids returns []uint32, set in DeviceIsisMultiInstance.
	Itids() []uint32
	// SetItids assigns []uint32 provided by user to DeviceIsisMultiInstance
	SetItids(value []uint32) DeviceIsisMultiInstance
}

// Instance Identifier (IID) TLV will associate a PDU with an ISIS instance  by using a unique 16-bit number and including one or more  Instance-Specific Topology Identifiers (ITIDs).
// Iid returns a uint32
func (obj *deviceIsisMultiInstance) Iid() uint32 {

	return *obj.obj.Iid

}

// Instance Identifier (IID) TLV will associate a PDU with an ISIS instance  by using a unique 16-bit number and including one or more  Instance-Specific Topology Identifiers (ITIDs).
// Iid returns a uint32
func (obj *deviceIsisMultiInstance) HasIid() bool {
	return obj.obj.Iid != nil
}

// Instance Identifier (IID) TLV will associate a PDU with an ISIS instance  by using a unique 16-bit number and including one or more  Instance-Specific Topology Identifiers (ITIDs).
// SetIid sets the uint32 value in the DeviceIsisMultiInstance object
func (obj *deviceIsisMultiInstance) SetIid(value uint32) DeviceIsisMultiInstance {

	obj.obj.Iid = &value
	return obj
}

// This contains one or more ITIDs that will be advertised in IID TLV.
// Itids returns a []uint32
func (obj *deviceIsisMultiInstance) Itids() []uint32 {
	if obj.obj.Itids == nil {
		obj.obj.Itids = make([]uint32, 0)
	}
	return obj.obj.Itids
}

// This contains one or more ITIDs that will be advertised in IID TLV.
// SetItids sets the []uint32 value in the DeviceIsisMultiInstance object
func (obj *deviceIsisMultiInstance) SetItids(value []uint32) DeviceIsisMultiInstance {

	if obj.obj.Itids == nil {
		obj.obj.Itids = make([]uint32, 0)
	}
	obj.obj.Itids = value

	return obj
}

func (obj *deviceIsisMultiInstance) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Iid != nil {

		if *obj.obj.Iid > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= DeviceIsisMultiInstance.Iid <= 65535 but Got %d", *obj.obj.Iid))
		}

	}

	if obj.obj.Itids != nil {

		for _, item := range obj.obj.Itids {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= DeviceIsisMultiInstance.Itids <= 65535 but Got %d", item))
			}

		}

	}

}

func (obj *deviceIsisMultiInstance) setDefault() {
	if obj.obj.Iid == nil {
		obj.SetIid(1)
	}

}
