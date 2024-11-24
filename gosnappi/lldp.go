package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Lldp *****
type lldp struct {
	validation
	obj              *otg.Lldp
	marshaller       marshalLldp
	unMarshaller     unMarshalLldp
	connectionHolder LldpConnection
	chassisIdHolder  LldpChassisId
	portIdHolder     LldpPortId
	systemNameHolder LldpSystemName
	orgInfosHolder   LldpLldpOrgInfoIter
}

func NewLldp() Lldp {
	obj := lldp{obj: &otg.Lldp{}}
	obj.setDefault()
	return &obj
}

func (obj *lldp) msg() *otg.Lldp {
	return obj.obj
}

func (obj *lldp) setMsg(msg *otg.Lldp) Lldp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallldp struct {
	obj *lldp
}

type marshalLldp interface {
	// ToProto marshals Lldp to protobuf object *otg.Lldp
	ToProto() (*otg.Lldp, error)
	// ToPbText marshals Lldp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Lldp to YAML text
	ToYaml() (string, error)
	// ToJson marshals Lldp to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Lldp to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallldp struct {
	obj *lldp
}

type unMarshalLldp interface {
	// FromProto unmarshals Lldp from protobuf object *otg.Lldp
	FromProto(msg *otg.Lldp) (Lldp, error)
	// FromPbText unmarshals Lldp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Lldp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Lldp from JSON text
	FromJson(value string) error
}

func (obj *lldp) Marshal() marshalLldp {
	if obj.marshaller == nil {
		obj.marshaller = &marshallldp{obj: obj}
	}
	return obj.marshaller
}

func (obj *lldp) Unmarshal() unMarshalLldp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallldp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallldp) ToProto() (*otg.Lldp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallldp) FromProto(msg *otg.Lldp) (Lldp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallldp) ToPbText() (string, error) {
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

func (m *unMarshallldp) FromPbText(value string) error {
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

func (m *marshallldp) ToYaml() (string, error) {
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

func (m *unMarshallldp) FromYaml(value string) error {
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

func (m *marshallldp) ToJsonRaw() (string, error) {
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

func (m *marshallldp) ToJson() (string, error) {
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

func (m *unMarshallldp) FromJson(value string) error {
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

func (obj *lldp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lldp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lldp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lldp) Clone() (Lldp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLldp()
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

func (obj *lldp) setNil() {
	obj.connectionHolder = nil
	obj.chassisIdHolder = nil
	obj.portIdHolder = nil
	obj.systemNameHolder = nil
	obj.orgInfosHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Lldp is configuration of LLDP protocol IEEE Ref: https://www.ieee802.org/1/files/public/docs2002/lldp-protocol-00.pdf
type Lldp interface {
	Validation
	// msg marshals Lldp to protobuf object *otg.Lldp
	// and doesn't set defaults
	msg() *otg.Lldp
	// setMsg unmarshals Lldp from protobuf object *otg.Lldp
	// and doesn't set defaults
	setMsg(*otg.Lldp) Lldp
	// provides marshal interface
	Marshal() marshalLldp
	// provides unmarshal interface
	Unmarshal() unMarshalLldp
	// validate validates Lldp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Lldp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Connection returns LldpConnection, set in Lldp.
	// LldpConnection is lLDP connection to a test port. In future if more connection options arise  LLDP connection object will be enhanced.
	Connection() LldpConnection
	// SetConnection assigns LldpConnection provided by user to Lldp.
	// LldpConnection is lLDP connection to a test port. In future if more connection options arise  LLDP connection object will be enhanced.
	SetConnection(value LldpConnection) Lldp
	// ChassisId returns LldpChassisId, set in Lldp.
	// LldpChassisId is the Chassis ID is a mandatory TLV which identifies the chassis component of the endpoint identifier associated  with the transmitting LLDP agent. This field identifies the format and source of the chassis identifier string.  It is based on the enumerator defined by the LldpChassisIdSubtype object from IEEE 802.1AB MIB.
	ChassisId() LldpChassisId
	// SetChassisId assigns LldpChassisId provided by user to Lldp.
	// LldpChassisId is the Chassis ID is a mandatory TLV which identifies the chassis component of the endpoint identifier associated  with the transmitting LLDP agent. This field identifies the format and source of the chassis identifier string.  It is based on the enumerator defined by the LldpChassisIdSubtype object from IEEE 802.1AB MIB.
	SetChassisId(value LldpChassisId) Lldp
	// HasChassisId checks if ChassisId has been set in Lldp
	HasChassisId() bool
	// PortId returns LldpPortId, set in Lldp.
	// LldpPortId is the Port ID is a mandatory TLV which identifies the port component of the endpoint identifier associated with  the transmitting LLDP agent.This field identifies the format and source of the port identifier string. It is  based on the enumerator defined by the PtopoPortIdType object from RFC2922.
	PortId() LldpPortId
	// SetPortId assigns LldpPortId provided by user to Lldp.
	// LldpPortId is the Port ID is a mandatory TLV which identifies the port component of the endpoint identifier associated with  the transmitting LLDP agent.This field identifies the format and source of the port identifier string. It is  based on the enumerator defined by the PtopoPortIdType object from RFC2922.
	SetPortId(value LldpPortId) Lldp
	// HasPortId checks if PortId has been set in Lldp
	HasPortId() bool
	// SystemName returns LldpSystemName, set in Lldp.
	// LldpSystemName is the system Name configured in the System Name TLV.
	SystemName() LldpSystemName
	// SetSystemName assigns LldpSystemName provided by user to Lldp.
	// LldpSystemName is the system Name configured in the System Name TLV.
	SetSystemName(value LldpSystemName) Lldp
	// HasSystemName checks if SystemName has been set in Lldp
	HasSystemName() bool
	// HoldTime returns uint32, set in Lldp.
	HoldTime() uint32
	// SetHoldTime assigns uint32 provided by user to Lldp
	SetHoldTime(value uint32) Lldp
	// HasHoldTime checks if HoldTime has been set in Lldp
	HasHoldTime() bool
	// AdvertisementInterval returns uint32, set in Lldp.
	AdvertisementInterval() uint32
	// SetAdvertisementInterval assigns uint32 provided by user to Lldp
	SetAdvertisementInterval(value uint32) Lldp
	// HasAdvertisementInterval checks if AdvertisementInterval has been set in Lldp
	HasAdvertisementInterval() bool
	// Name returns string, set in Lldp.
	Name() string
	// SetName assigns string provided by user to Lldp
	SetName(value string) Lldp
	// OrgInfos returns LldpLldpOrgInfoIterIter, set in Lldp
	OrgInfos() LldpLldpOrgInfoIter
	setNil()
}

// The unique name of the object on which LLDP is running.
// Connection returns a LldpConnection
func (obj *lldp) Connection() LldpConnection {
	if obj.obj.Connection == nil {
		obj.obj.Connection = NewLldpConnection().msg()
	}
	if obj.connectionHolder == nil {
		obj.connectionHolder = &lldpConnection{obj: obj.obj.Connection}
	}
	return obj.connectionHolder
}

// The unique name of the object on which LLDP is running.
// SetConnection sets the LldpConnection value in the Lldp object
func (obj *lldp) SetConnection(value LldpConnection) Lldp {

	obj.connectionHolder = nil
	obj.obj.Connection = value.msg()

	return obj
}

// The Chassis ID is a mandatory TLV which identifies the chassis component of the endpoint identifier associated  with the transmitting LLDP agent. If mac address is specified it should be in colon seperated mac address format.
// ChassisId returns a LldpChassisId
func (obj *lldp) ChassisId() LldpChassisId {
	if obj.obj.ChassisId == nil {
		obj.obj.ChassisId = NewLldpChassisId().msg()
	}
	if obj.chassisIdHolder == nil {
		obj.chassisIdHolder = &lldpChassisId{obj: obj.obj.ChassisId}
	}
	return obj.chassisIdHolder
}

// The Chassis ID is a mandatory TLV which identifies the chassis component of the endpoint identifier associated  with the transmitting LLDP agent. If mac address is specified it should be in colon seperated mac address format.
// ChassisId returns a LldpChassisId
func (obj *lldp) HasChassisId() bool {
	return obj.obj.ChassisId != nil
}

// The Chassis ID is a mandatory TLV which identifies the chassis component of the endpoint identifier associated  with the transmitting LLDP agent. If mac address is specified it should be in colon seperated mac address format.
// SetChassisId sets the LldpChassisId value in the Lldp object
func (obj *lldp) SetChassisId(value LldpChassisId) Lldp {

	obj.chassisIdHolder = nil
	obj.obj.ChassisId = value.msg()

	return obj
}

// The Port ID is a mandatory TLV which identifies the port component of the endpoint identifier associated with  the transmitting LLDP agent. If the specified port is an IEEE 802.3 Repeater port, then this TLV is optional.
// PortId returns a LldpPortId
func (obj *lldp) PortId() LldpPortId {
	if obj.obj.PortId == nil {
		obj.obj.PortId = NewLldpPortId().msg()
	}
	if obj.portIdHolder == nil {
		obj.portIdHolder = &lldpPortId{obj: obj.obj.PortId}
	}
	return obj.portIdHolder
}

// The Port ID is a mandatory TLV which identifies the port component of the endpoint identifier associated with  the transmitting LLDP agent. If the specified port is an IEEE 802.3 Repeater port, then this TLV is optional.
// PortId returns a LldpPortId
func (obj *lldp) HasPortId() bool {
	return obj.obj.PortId != nil
}

// The Port ID is a mandatory TLV which identifies the port component of the endpoint identifier associated with  the transmitting LLDP agent. If the specified port is an IEEE 802.3 Repeater port, then this TLV is optional.
// SetPortId sets the LldpPortId value in the Lldp object
func (obj *lldp) SetPortId(value LldpPortId) Lldp {

	obj.portIdHolder = nil
	obj.obj.PortId = value.msg()

	return obj
}

// The system name field shall contain an alpha-numeric string that indicates the system's administratively assigned  name. The system name should be the system's fully qualified domain name. If implementations support IETF RFC  3418, the sysName object should be used for this field.
// SystemName returns a LldpSystemName
func (obj *lldp) SystemName() LldpSystemName {
	if obj.obj.SystemName == nil {
		obj.obj.SystemName = NewLldpSystemName().msg()
	}
	if obj.systemNameHolder == nil {
		obj.systemNameHolder = &lldpSystemName{obj: obj.obj.SystemName}
	}
	return obj.systemNameHolder
}

// The system name field shall contain an alpha-numeric string that indicates the system's administratively assigned  name. The system name should be the system's fully qualified domain name. If implementations support IETF RFC  3418, the sysName object should be used for this field.
// SystemName returns a LldpSystemName
func (obj *lldp) HasSystemName() bool {
	return obj.obj.SystemName != nil
}

// The system name field shall contain an alpha-numeric string that indicates the system's administratively assigned  name. The system name should be the system's fully qualified domain name. If implementations support IETF RFC  3418, the sysName object should be used for this field.
// SetSystemName sets the LldpSystemName value in the Lldp object
func (obj *lldp) SetSystemName(value LldpSystemName) Lldp {

	obj.systemNameHolder = nil
	obj.obj.SystemName = value.msg()

	return obj
}

// Specifies the amount of time in seconds a receiving device should maintain LLDP information sent  by the device before discarding it.
// HoldTime returns a uint32
func (obj *lldp) HoldTime() uint32 {

	return *obj.obj.HoldTime

}

// Specifies the amount of time in seconds a receiving device should maintain LLDP information sent  by the device before discarding it.
// HoldTime returns a uint32
func (obj *lldp) HasHoldTime() bool {
	return obj.obj.HoldTime != nil
}

// Specifies the amount of time in seconds a receiving device should maintain LLDP information sent  by the device before discarding it.
// SetHoldTime sets the uint32 value in the Lldp object
func (obj *lldp) SetHoldTime(value uint32) Lldp {

	obj.obj.HoldTime = &value
	return obj
}

// Set the transmission frequency of LLDP updates in seconds.
// AdvertisementInterval returns a uint32
func (obj *lldp) AdvertisementInterval() uint32 {

	return *obj.obj.AdvertisementInterval

}

// Set the transmission frequency of LLDP updates in seconds.
// AdvertisementInterval returns a uint32
func (obj *lldp) HasAdvertisementInterval() bool {
	return obj.obj.AdvertisementInterval != nil
}

// Set the transmission frequency of LLDP updates in seconds.
// SetAdvertisementInterval sets the uint32 value in the Lldp object
func (obj *lldp) SetAdvertisementInterval(value uint32) Lldp {

	obj.obj.AdvertisementInterval = &value
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *lldp) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Lldp object
func (obj *lldp) SetName(value string) Lldp {

	obj.obj.Name = &value
	return obj
}

// The Organization Information is used to define the organization specific TLVs. The organization specific TLV is defined in IEEE 802.1AB-2016 specification. This category is provided to allow different organizations, such as  IEEE 802.1, IEEE 802.3, IETF, as well as individual software and equipment vendors, to define TLVs that advertise  information to remote entities attached to the same media.
// OrgInfos returns a []LldpOrgInfo
func (obj *lldp) OrgInfos() LldpLldpOrgInfoIter {
	if len(obj.obj.OrgInfos) == 0 {
		obj.obj.OrgInfos = []*otg.LldpOrgInfo{}
	}
	if obj.orgInfosHolder == nil {
		obj.orgInfosHolder = newLldpLldpOrgInfoIter(&obj.obj.OrgInfos).setMsg(obj)
	}
	return obj.orgInfosHolder
}

type lldpLldpOrgInfoIter struct {
	obj              *lldp
	lldpOrgInfoSlice []LldpOrgInfo
	fieldPtr         *[]*otg.LldpOrgInfo
}

func newLldpLldpOrgInfoIter(ptr *[]*otg.LldpOrgInfo) LldpLldpOrgInfoIter {
	return &lldpLldpOrgInfoIter{fieldPtr: ptr}
}

type LldpLldpOrgInfoIter interface {
	setMsg(*lldp) LldpLldpOrgInfoIter
	Items() []LldpOrgInfo
	Add() LldpOrgInfo
	Append(items ...LldpOrgInfo) LldpLldpOrgInfoIter
	Set(index int, newObj LldpOrgInfo) LldpLldpOrgInfoIter
	Clear() LldpLldpOrgInfoIter
	clearHolderSlice() LldpLldpOrgInfoIter
	appendHolderSlice(item LldpOrgInfo) LldpLldpOrgInfoIter
}

func (obj *lldpLldpOrgInfoIter) setMsg(msg *lldp) LldpLldpOrgInfoIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&lldpOrgInfo{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *lldpLldpOrgInfoIter) Items() []LldpOrgInfo {
	return obj.lldpOrgInfoSlice
}

func (obj *lldpLldpOrgInfoIter) Add() LldpOrgInfo {
	newObj := &otg.LldpOrgInfo{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &lldpOrgInfo{obj: newObj}
	newLibObj.setDefault()
	obj.lldpOrgInfoSlice = append(obj.lldpOrgInfoSlice, newLibObj)
	return newLibObj
}

func (obj *lldpLldpOrgInfoIter) Append(items ...LldpOrgInfo) LldpLldpOrgInfoIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.lldpOrgInfoSlice = append(obj.lldpOrgInfoSlice, item)
	}
	return obj
}

func (obj *lldpLldpOrgInfoIter) Set(index int, newObj LldpOrgInfo) LldpLldpOrgInfoIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.lldpOrgInfoSlice[index] = newObj
	return obj
}
func (obj *lldpLldpOrgInfoIter) Clear() LldpLldpOrgInfoIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.LldpOrgInfo{}
		obj.lldpOrgInfoSlice = []LldpOrgInfo{}
	}
	return obj
}
func (obj *lldpLldpOrgInfoIter) clearHolderSlice() LldpLldpOrgInfoIter {
	if len(obj.lldpOrgInfoSlice) > 0 {
		obj.lldpOrgInfoSlice = []LldpOrgInfo{}
	}
	return obj
}
func (obj *lldpLldpOrgInfoIter) appendHolderSlice(item LldpOrgInfo) LldpLldpOrgInfoIter {
	obj.lldpOrgInfoSlice = append(obj.lldpOrgInfoSlice, item)
	return obj
}

func (obj *lldp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Connection is required
	if obj.obj.Connection == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Connection is required field on interface Lldp")
	}

	if obj.obj.Connection != nil {

		obj.Connection().validateObj(vObj, set_default)
	}

	if obj.obj.ChassisId != nil {

		obj.ChassisId().validateObj(vObj, set_default)
	}

	if obj.obj.PortId != nil {

		obj.PortId().validateObj(vObj, set_default)
	}

	if obj.obj.SystemName != nil {

		obj.SystemName().validateObj(vObj, set_default)
	}

	if obj.obj.HoldTime != nil {

		if *obj.obj.HoldTime < 10 || *obj.obj.HoldTime > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("10 <= Lldp.HoldTime <= 65535 but Got %d", *obj.obj.HoldTime))
		}

	}

	if obj.obj.AdvertisementInterval != nil {

		if *obj.obj.AdvertisementInterval < 5 || *obj.obj.AdvertisementInterval > 65534 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("5 <= Lldp.AdvertisementInterval <= 65534 but Got %d", *obj.obj.AdvertisementInterval))
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Lldp")
	}

	if len(obj.obj.OrgInfos) != 0 {

		if set_default {
			obj.OrgInfos().clearHolderSlice()
			for _, item := range obj.obj.OrgInfos {
				obj.OrgInfos().appendHolderSlice(&lldpOrgInfo{obj: item})
			}
		}
		for _, item := range obj.OrgInfos().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *lldp) setDefault() {
	if obj.obj.HoldTime == nil {
		obj.SetHoldTime(120)
	}
	if obj.obj.AdvertisementInterval == nil {
		obj.SetAdvertisementInterval(30)
	}

}
