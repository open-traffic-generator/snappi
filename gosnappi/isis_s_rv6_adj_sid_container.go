package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRv6AdjSidContainer *****
type isisSRv6AdjSidContainer struct {
	validation
	obj               *otg.IsisSRv6AdjSidContainer
	marshaller        marshalIsisSRv6AdjSidContainer
	unMarshaller      unMarshalIsisSRv6AdjSidContainer
	sidsHolder        IsisSRv6AdjSidContainerIsisSRv6AdjSidIter
	srv6LinkMsdHolder IsisSRv6Msd
}

func NewIsisSRv6AdjSidContainer() IsisSRv6AdjSidContainer {
	obj := isisSRv6AdjSidContainer{obj: &otg.IsisSRv6AdjSidContainer{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRv6AdjSidContainer) msg() *otg.IsisSRv6AdjSidContainer {
	return obj.obj
}

func (obj *isisSRv6AdjSidContainer) setMsg(msg *otg.IsisSRv6AdjSidContainer) IsisSRv6AdjSidContainer {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRv6AdjSidContainer struct {
	obj *isisSRv6AdjSidContainer
}

type marshalIsisSRv6AdjSidContainer interface {
	// ToProto marshals IsisSRv6AdjSidContainer to protobuf object *otg.IsisSRv6AdjSidContainer
	ToProto() (*otg.IsisSRv6AdjSidContainer, error)
	// ToPbText marshals IsisSRv6AdjSidContainer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRv6AdjSidContainer to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRv6AdjSidContainer to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRv6AdjSidContainer struct {
	obj *isisSRv6AdjSidContainer
}

type unMarshalIsisSRv6AdjSidContainer interface {
	// FromProto unmarshals IsisSRv6AdjSidContainer from protobuf object *otg.IsisSRv6AdjSidContainer
	FromProto(msg *otg.IsisSRv6AdjSidContainer) (IsisSRv6AdjSidContainer, error)
	// FromPbText unmarshals IsisSRv6AdjSidContainer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRv6AdjSidContainer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRv6AdjSidContainer from JSON text
	FromJson(value string) error
}

func (obj *isisSRv6AdjSidContainer) Marshal() marshalIsisSRv6AdjSidContainer {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRv6AdjSidContainer{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRv6AdjSidContainer) Unmarshal() unMarshalIsisSRv6AdjSidContainer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRv6AdjSidContainer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRv6AdjSidContainer) ToProto() (*otg.IsisSRv6AdjSidContainer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRv6AdjSidContainer) FromProto(msg *otg.IsisSRv6AdjSidContainer) (IsisSRv6AdjSidContainer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRv6AdjSidContainer) ToPbText() (string, error) {
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

func (m *unMarshalisisSRv6AdjSidContainer) FromPbText(value string) error {
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

func (m *marshalisisSRv6AdjSidContainer) ToYaml() (string, error) {
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

func (m *unMarshalisisSRv6AdjSidContainer) FromYaml(value string) error {
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

func (m *marshalisisSRv6AdjSidContainer) ToJson() (string, error) {
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

func (m *unMarshalisisSRv6AdjSidContainer) FromJson(value string) error {
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

func (obj *isisSRv6AdjSidContainer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRv6AdjSidContainer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRv6AdjSidContainer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRv6AdjSidContainer) Clone() (IsisSRv6AdjSidContainer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRv6AdjSidContainer()
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

func (obj *isisSRv6AdjSidContainer) setNil() {
	obj.sidsHolder = nil
	obj.srv6LinkMsdHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisSRv6AdjSidContainer is container for SRv6 adjacency-related sub-TLVs advertised on an IS-IS interface. Groups the list of End.X SID sub-TLVs (RFC 9352 Sections 8.1-8.2) with the per-link Maximum SID Depth (MSD) sub-TLVs (RFC 9352 Section 6) that are advertised together in the Extended IS Reachability TLV for this interface.
type IsisSRv6AdjSidContainer interface {
	Validation
	// msg marshals IsisSRv6AdjSidContainer to protobuf object *otg.IsisSRv6AdjSidContainer
	// and doesn't set defaults
	msg() *otg.IsisSRv6AdjSidContainer
	// setMsg unmarshals IsisSRv6AdjSidContainer from protobuf object *otg.IsisSRv6AdjSidContainer
	// and doesn't set defaults
	setMsg(*otg.IsisSRv6AdjSidContainer) IsisSRv6AdjSidContainer
	// provides marshal interface
	Marshal() marshalIsisSRv6AdjSidContainer
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRv6AdjSidContainer
	// validate validates IsisSRv6AdjSidContainer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRv6AdjSidContainer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Sids returns IsisSRv6AdjSidContainerIsisSRv6AdjSidIterIter, set in IsisSRv6AdjSidContainer
	Sids() IsisSRv6AdjSidContainerIsisSRv6AdjSidIter
	// Srv6LinkMsd returns IsisSRv6Msd, set in IsisSRv6AdjSidContainer.
	// IsisSRv6Msd is sRv6 Maximum SID Depth (MSD) sub-TLV values, used for both node-level advertisement (carried in IS-IS Router CAPABILITY TLV 242 via Node MSD sub-TLV type 23, RFC 9352 Section 6) and link-level advertisement (carried in Extended IS Reachability TLV via Link MSD sub-TLV type 15, RFC 9352 Section 6). When a property is present, the corresponding MSD sub-TLV is advertised with the configured value. Omit a property to suppress advertisement of that MSD type. Reference: RFC 9352 Section 6.
	Srv6LinkMsd() IsisSRv6Msd
	// SetSrv6LinkMsd assigns IsisSRv6Msd provided by user to IsisSRv6AdjSidContainer.
	// IsisSRv6Msd is sRv6 Maximum SID Depth (MSD) sub-TLV values, used for both node-level advertisement (carried in IS-IS Router CAPABILITY TLV 242 via Node MSD sub-TLV type 23, RFC 9352 Section 6) and link-level advertisement (carried in Extended IS Reachability TLV via Link MSD sub-TLV type 15, RFC 9352 Section 6). When a property is present, the corresponding MSD sub-TLV is advertised with the configured value. Omit a property to suppress advertisement of that MSD type. Reference: RFC 9352 Section 6.
	SetSrv6LinkMsd(value IsisSRv6Msd) IsisSRv6AdjSidContainer
	// HasSrv6LinkMsd checks if Srv6LinkMsd has been set in IsisSRv6AdjSidContainer
	HasSrv6LinkMsd() bool
	setNil()
}

// List of SRv6 Adjacency SID sub-TLVs (End.X SID) for this interface. Point-to-point interfaces advertise End.X SID sub-TLV (sub-TLV type 43); broadcast interfaces advertise LAN End.X SID sub-TLV (sub-TLV type 44). Each entry binds a 128-bit SRv6 SID to this specific outgoing adjacency and advertises the associated endpoint behavior. Reference: RFC 9352 Sections 8.1-8.2.
// Sids returns a []IsisSRv6AdjSid
func (obj *isisSRv6AdjSidContainer) Sids() IsisSRv6AdjSidContainerIsisSRv6AdjSidIter {
	if len(obj.obj.Sids) == 0 {
		obj.obj.Sids = []*otg.IsisSRv6AdjSid{}
	}
	if obj.sidsHolder == nil {
		obj.sidsHolder = newIsisSRv6AdjSidContainerIsisSRv6AdjSidIter(&obj.obj.Sids).setMsg(obj)
	}
	return obj.sidsHolder
}

type isisSRv6AdjSidContainerIsisSRv6AdjSidIter struct {
	obj                 *isisSRv6AdjSidContainer
	isisSRv6AdjSidSlice []IsisSRv6AdjSid
	fieldPtr            *[]*otg.IsisSRv6AdjSid
}

func newIsisSRv6AdjSidContainerIsisSRv6AdjSidIter(ptr *[]*otg.IsisSRv6AdjSid) IsisSRv6AdjSidContainerIsisSRv6AdjSidIter {
	return &isisSRv6AdjSidContainerIsisSRv6AdjSidIter{fieldPtr: ptr}
}

type IsisSRv6AdjSidContainerIsisSRv6AdjSidIter interface {
	setMsg(*isisSRv6AdjSidContainer) IsisSRv6AdjSidContainerIsisSRv6AdjSidIter
	Items() []IsisSRv6AdjSid
	Add() IsisSRv6AdjSid
	Append(items ...IsisSRv6AdjSid) IsisSRv6AdjSidContainerIsisSRv6AdjSidIter
	Set(index int, newObj IsisSRv6AdjSid) IsisSRv6AdjSidContainerIsisSRv6AdjSidIter
	Clear() IsisSRv6AdjSidContainerIsisSRv6AdjSidIter
	clearHolderSlice() IsisSRv6AdjSidContainerIsisSRv6AdjSidIter
	appendHolderSlice(item IsisSRv6AdjSid) IsisSRv6AdjSidContainerIsisSRv6AdjSidIter
}

func (obj *isisSRv6AdjSidContainerIsisSRv6AdjSidIter) setMsg(msg *isisSRv6AdjSidContainer) IsisSRv6AdjSidContainerIsisSRv6AdjSidIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisSRv6AdjSid{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisSRv6AdjSidContainerIsisSRv6AdjSidIter) Items() []IsisSRv6AdjSid {
	return obj.isisSRv6AdjSidSlice
}

func (obj *isisSRv6AdjSidContainerIsisSRv6AdjSidIter) Add() IsisSRv6AdjSid {
	newObj := &otg.IsisSRv6AdjSid{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisSRv6AdjSid{obj: newObj}
	newLibObj.setDefault()
	obj.isisSRv6AdjSidSlice = append(obj.isisSRv6AdjSidSlice, newLibObj)
	return newLibObj
}

func (obj *isisSRv6AdjSidContainerIsisSRv6AdjSidIter) Append(items ...IsisSRv6AdjSid) IsisSRv6AdjSidContainerIsisSRv6AdjSidIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisSRv6AdjSidSlice = append(obj.isisSRv6AdjSidSlice, item)
	}
	return obj
}

func (obj *isisSRv6AdjSidContainerIsisSRv6AdjSidIter) Set(index int, newObj IsisSRv6AdjSid) IsisSRv6AdjSidContainerIsisSRv6AdjSidIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisSRv6AdjSidSlice[index] = newObj
	return obj
}
func (obj *isisSRv6AdjSidContainerIsisSRv6AdjSidIter) Clear() IsisSRv6AdjSidContainerIsisSRv6AdjSidIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisSRv6AdjSid{}
		obj.isisSRv6AdjSidSlice = []IsisSRv6AdjSid{}
	}
	return obj
}
func (obj *isisSRv6AdjSidContainerIsisSRv6AdjSidIter) clearHolderSlice() IsisSRv6AdjSidContainerIsisSRv6AdjSidIter {
	if len(obj.isisSRv6AdjSidSlice) > 0 {
		obj.isisSRv6AdjSidSlice = []IsisSRv6AdjSid{}
	}
	return obj
}
func (obj *isisSRv6AdjSidContainerIsisSRv6AdjSidIter) appendHolderSlice(item IsisSRv6AdjSid) IsisSRv6AdjSidContainerIsisSRv6AdjSidIter {
	obj.isisSRv6AdjSidSlice = append(obj.isisSRv6AdjSidSlice, item)
	return obj
}

// When present, advertises per-link SRv6 Maximum SID Depth (MSD) sub-TLVs within the IS-IS Extended IS Reachability TLV for this interface. Each populated child object causes the corresponding link MSD sub-TLV to be advertised; omit a child object to suppress that MSD type. Reference: RFC 9352 Section 6.
// Srv6LinkMsd returns a IsisSRv6Msd
func (obj *isisSRv6AdjSidContainer) Srv6LinkMsd() IsisSRv6Msd {
	if obj.obj.Srv6LinkMsd == nil {
		obj.obj.Srv6LinkMsd = NewIsisSRv6Msd().msg()
	}
	if obj.srv6LinkMsdHolder == nil {
		obj.srv6LinkMsdHolder = &isisSRv6Msd{obj: obj.obj.Srv6LinkMsd}
	}
	return obj.srv6LinkMsdHolder
}

// When present, advertises per-link SRv6 Maximum SID Depth (MSD) sub-TLVs within the IS-IS Extended IS Reachability TLV for this interface. Each populated child object causes the corresponding link MSD sub-TLV to be advertised; omit a child object to suppress that MSD type. Reference: RFC 9352 Section 6.
// Srv6LinkMsd returns a IsisSRv6Msd
func (obj *isisSRv6AdjSidContainer) HasSrv6LinkMsd() bool {
	return obj.obj.Srv6LinkMsd != nil
}

// When present, advertises per-link SRv6 Maximum SID Depth (MSD) sub-TLVs within the IS-IS Extended IS Reachability TLV for this interface. Each populated child object causes the corresponding link MSD sub-TLV to be advertised; omit a child object to suppress that MSD type. Reference: RFC 9352 Section 6.
// SetSrv6LinkMsd sets the IsisSRv6Msd value in the IsisSRv6AdjSidContainer object
func (obj *isisSRv6AdjSidContainer) SetSrv6LinkMsd(value IsisSRv6Msd) IsisSRv6AdjSidContainer {

	obj.srv6LinkMsdHolder = nil
	obj.obj.Srv6LinkMsd = value.msg()

	return obj
}

func (obj *isisSRv6AdjSidContainer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Sids) != 0 {

		if set_default {
			obj.Sids().clearHolderSlice()
			for _, item := range obj.obj.Sids {
				obj.Sids().appendHolderSlice(&isisSRv6AdjSid{obj: item})
			}
		}
		for _, item := range obj.Sids().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Srv6LinkMsd != nil {

		obj.Srv6LinkMsd().validateObj(vObj, set_default)
	}

}

func (obj *isisSRv6AdjSidContainer) setDefault() {

}
