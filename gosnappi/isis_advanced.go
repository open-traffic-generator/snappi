package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisAdvanced *****
type isisAdvanced struct {
	validation
	obj          *otg.IsisAdvanced
	marshaller   marshalIsisAdvanced
	unMarshaller unMarshalIsisAdvanced
}

func NewIsisAdvanced() IsisAdvanced {
	obj := isisAdvanced{obj: &otg.IsisAdvanced{}}
	obj.setDefault()
	return &obj
}

func (obj *isisAdvanced) msg() *otg.IsisAdvanced {
	return obj.obj
}

func (obj *isisAdvanced) setMsg(msg *otg.IsisAdvanced) IsisAdvanced {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisAdvanced struct {
	obj *isisAdvanced
}

type marshalIsisAdvanced interface {
	// ToProto marshals IsisAdvanced to protobuf object *otg.IsisAdvanced
	ToProto() (*otg.IsisAdvanced, error)
	// ToPbText marshals IsisAdvanced to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisAdvanced to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisAdvanced to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisAdvanced to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisAdvanced struct {
	obj *isisAdvanced
}

type unMarshalIsisAdvanced interface {
	// FromProto unmarshals IsisAdvanced from protobuf object *otg.IsisAdvanced
	FromProto(msg *otg.IsisAdvanced) (IsisAdvanced, error)
	// FromPbText unmarshals IsisAdvanced from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisAdvanced from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisAdvanced from JSON text
	FromJson(value string) error
}

func (obj *isisAdvanced) Marshal() marshalIsisAdvanced {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisAdvanced{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisAdvanced) Unmarshal() unMarshalIsisAdvanced {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisAdvanced{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisAdvanced) ToProto() (*otg.IsisAdvanced, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisAdvanced) FromProto(msg *otg.IsisAdvanced) (IsisAdvanced, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisAdvanced) ToPbText() (string, error) {
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

func (m *unMarshalisisAdvanced) FromPbText(value string) error {
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

func (m *marshalisisAdvanced) ToYaml() (string, error) {
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

func (m *unMarshalisisAdvanced) FromYaml(value string) error {
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

func (m *marshalisisAdvanced) ToJsonRaw() (string, error) {
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

func (m *marshalisisAdvanced) ToJson() (string, error) {
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

func (m *unMarshalisisAdvanced) FromJson(value string) error {
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

func (obj *isisAdvanced) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisAdvanced) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisAdvanced) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisAdvanced) Clone() (IsisAdvanced, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisAdvanced()
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

// IsisAdvanced is contains ISIS router advanced properties.
type IsisAdvanced interface {
	Validation
	// msg marshals IsisAdvanced to protobuf object *otg.IsisAdvanced
	// and doesn't set defaults
	msg() *otg.IsisAdvanced
	// setMsg unmarshals IsisAdvanced from protobuf object *otg.IsisAdvanced
	// and doesn't set defaults
	setMsg(*otg.IsisAdvanced) IsisAdvanced
	// provides marshal interface
	Marshal() marshalIsisAdvanced
	// provides unmarshal interface
	Unmarshal() unMarshalIsisAdvanced
	// validate validates IsisAdvanced
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisAdvanced, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EnableHelloPadding returns bool, set in IsisAdvanced.
	EnableHelloPadding() bool
	// SetEnableHelloPadding assigns bool provided by user to IsisAdvanced
	SetEnableHelloPadding(value bool) IsisAdvanced
	// HasEnableHelloPadding checks if EnableHelloPadding has been set in IsisAdvanced
	HasEnableHelloPadding() bool
	// MaxAreaAddresses returns uint32, set in IsisAdvanced.
	MaxAreaAddresses() uint32
	// SetMaxAreaAddresses assigns uint32 provided by user to IsisAdvanced
	SetMaxAreaAddresses(value uint32) IsisAdvanced
	// HasMaxAreaAddresses checks if MaxAreaAddresses has been set in IsisAdvanced
	HasMaxAreaAddresses() bool
	// AreaAddresses returns []string, set in IsisAdvanced.
	AreaAddresses() []string
	// SetAreaAddresses assigns []string provided by user to IsisAdvanced
	SetAreaAddresses(value []string) IsisAdvanced
	// LspRefreshRate returns uint32, set in IsisAdvanced.
	LspRefreshRate() uint32
	// SetLspRefreshRate assigns uint32 provided by user to IsisAdvanced
	SetLspRefreshRate(value uint32) IsisAdvanced
	// HasLspRefreshRate checks if LspRefreshRate has been set in IsisAdvanced
	HasLspRefreshRate() bool
	// LspLifetime returns uint32, set in IsisAdvanced.
	LspLifetime() uint32
	// SetLspLifetime assigns uint32 provided by user to IsisAdvanced
	SetLspLifetime(value uint32) IsisAdvanced
	// HasLspLifetime checks if LspLifetime has been set in IsisAdvanced
	HasLspLifetime() bool
	// PsnpInterval returns uint32, set in IsisAdvanced.
	PsnpInterval() uint32
	// SetPsnpInterval assigns uint32 provided by user to IsisAdvanced
	SetPsnpInterval(value uint32) IsisAdvanced
	// HasPsnpInterval checks if PsnpInterval has been set in IsisAdvanced
	HasPsnpInterval() bool
	// CsnpInterval returns uint32, set in IsisAdvanced.
	CsnpInterval() uint32
	// SetCsnpInterval assigns uint32 provided by user to IsisAdvanced
	SetCsnpInterval(value uint32) IsisAdvanced
	// HasCsnpInterval checks if CsnpInterval has been set in IsisAdvanced
	HasCsnpInterval() bool
	// MaxLspSize returns uint32, set in IsisAdvanced.
	MaxLspSize() uint32
	// SetMaxLspSize assigns uint32 provided by user to IsisAdvanced
	SetMaxLspSize(value uint32) IsisAdvanced
	// HasMaxLspSize checks if MaxLspSize has been set in IsisAdvanced
	HasMaxLspSize() bool
	// LspMgroupMinTransInterval returns uint32, set in IsisAdvanced.
	LspMgroupMinTransInterval() uint32
	// SetLspMgroupMinTransInterval assigns uint32 provided by user to IsisAdvanced
	SetLspMgroupMinTransInterval(value uint32) IsisAdvanced
	// HasLspMgroupMinTransInterval checks if LspMgroupMinTransInterval has been set in IsisAdvanced
	HasLspMgroupMinTransInterval() bool
	// EnableAttachedBit returns bool, set in IsisAdvanced.
	EnableAttachedBit() bool
	// SetEnableAttachedBit assigns bool provided by user to IsisAdvanced
	SetEnableAttachedBit(value bool) IsisAdvanced
	// HasEnableAttachedBit checks if EnableAttachedBit has been set in IsisAdvanced
	HasEnableAttachedBit() bool
}

// It enables padding of Hello message to MTU size.
// EnableHelloPadding returns a bool
func (obj *isisAdvanced) EnableHelloPadding() bool {

	return *obj.obj.EnableHelloPadding

}

// It enables padding of Hello message to MTU size.
// EnableHelloPadding returns a bool
func (obj *isisAdvanced) HasEnableHelloPadding() bool {
	return obj.obj.EnableHelloPadding != nil
}

// It enables padding of Hello message to MTU size.
// SetEnableHelloPadding sets the bool value in the IsisAdvanced object
func (obj *isisAdvanced) SetEnableHelloPadding(value bool) IsisAdvanced {

	obj.obj.EnableHelloPadding = &value
	return obj
}

// The Number of Area Addresses permitted, with a valid range from 0 to 254.  A zero indicates a maximum of 3 addresses.
// MaxAreaAddresses returns a uint32
func (obj *isisAdvanced) MaxAreaAddresses() uint32 {

	return *obj.obj.MaxAreaAddresses

}

// The Number of Area Addresses permitted, with a valid range from 0 to 254.  A zero indicates a maximum of 3 addresses.
// MaxAreaAddresses returns a uint32
func (obj *isisAdvanced) HasMaxAreaAddresses() bool {
	return obj.obj.MaxAreaAddresses != nil
}

// The Number of Area Addresses permitted, with a valid range from 0 to 254.  A zero indicates a maximum of 3 addresses.
// SetMaxAreaAddresses sets the uint32 value in the IsisAdvanced object
func (obj *isisAdvanced) SetMaxAreaAddresses(value uint32) IsisAdvanced {

	obj.obj.MaxAreaAddresses = &value
	return obj
}

// Its combination of the ISP and HO-DSP.Usually all nodes within an area have  the same area address. If no area addresses are configured, a default area of "490001" will be advertised.
// AreaAddresses returns a []string
func (obj *isisAdvanced) AreaAddresses() []string {
	if obj.obj.AreaAddresses == nil {
		obj.obj.AreaAddresses = make([]string, 0)
	}
	return obj.obj.AreaAddresses
}

// Its combination of the ISP and HO-DSP.Usually all nodes within an area have  the same area address. If no area addresses are configured, a default area of "490001" will be advertised.
// SetAreaAddresses sets the []string value in the IsisAdvanced object
func (obj *isisAdvanced) SetAreaAddresses(value []string) IsisAdvanced {

	if obj.obj.AreaAddresses == nil {
		obj.obj.AreaAddresses = make([]string, 0)
	}
	obj.obj.AreaAddresses = value

	return obj
}

// The rate at which LSPs are re-sent in seconds.
// LspRefreshRate returns a uint32
func (obj *isisAdvanced) LspRefreshRate() uint32 {

	return *obj.obj.LspRefreshRate

}

// The rate at which LSPs are re-sent in seconds.
// LspRefreshRate returns a uint32
func (obj *isisAdvanced) HasLspRefreshRate() bool {
	return obj.obj.LspRefreshRate != nil
}

// The rate at which LSPs are re-sent in seconds.
// SetLspRefreshRate sets the uint32 value in the IsisAdvanced object
func (obj *isisAdvanced) SetLspRefreshRate(value uint32) IsisAdvanced {

	obj.obj.LspRefreshRate = &value
	return obj
}

// The MaxAge for retaining a learned LSP on this router in seconds.
// LspLifetime returns a uint32
func (obj *isisAdvanced) LspLifetime() uint32 {

	return *obj.obj.LspLifetime

}

// The MaxAge for retaining a learned LSP on this router in seconds.
// LspLifetime returns a uint32
func (obj *isisAdvanced) HasLspLifetime() bool {
	return obj.obj.LspLifetime != nil
}

// The MaxAge for retaining a learned LSP on this router in seconds.
// SetLspLifetime sets the uint32 value in the IsisAdvanced object
func (obj *isisAdvanced) SetLspLifetime(value uint32) IsisAdvanced {

	obj.obj.LspLifetime = &value
	return obj
}

// The number of milliseconds between transmissions of Partial Sequence Number PDU.
// PsnpInterval returns a uint32
func (obj *isisAdvanced) PsnpInterval() uint32 {

	return *obj.obj.PsnpInterval

}

// The number of milliseconds between transmissions of Partial Sequence Number PDU.
// PsnpInterval returns a uint32
func (obj *isisAdvanced) HasPsnpInterval() bool {
	return obj.obj.PsnpInterval != nil
}

// The number of milliseconds between transmissions of Partial Sequence Number PDU.
// SetPsnpInterval sets the uint32 value in the IsisAdvanced object
func (obj *isisAdvanced) SetPsnpInterval(value uint32) IsisAdvanced {

	obj.obj.PsnpInterval = &value
	return obj
}

// The number of milliseconds between transmissions of Partial Sequence Number PDU.
// CsnpInterval returns a uint32
func (obj *isisAdvanced) CsnpInterval() uint32 {

	return *obj.obj.CsnpInterval

}

// The number of milliseconds between transmissions of Partial Sequence Number PDU.
// CsnpInterval returns a uint32
func (obj *isisAdvanced) HasCsnpInterval() bool {
	return obj.obj.CsnpInterval != nil
}

// The number of milliseconds between transmissions of Partial Sequence Number PDU.
// SetCsnpInterval sets the uint32 value in the IsisAdvanced object
func (obj *isisAdvanced) SetCsnpInterval(value uint32) IsisAdvanced {

	obj.obj.CsnpInterval = &value
	return obj
}

// The maximum size in bytes of any LSP that can be transmitted over a link of equal or less  than maximum MTU size.
// MaxLspSize returns a uint32
func (obj *isisAdvanced) MaxLspSize() uint32 {

	return *obj.obj.MaxLspSize

}

// The maximum size in bytes of any LSP that can be transmitted over a link of equal or less  than maximum MTU size.
// MaxLspSize returns a uint32
func (obj *isisAdvanced) HasMaxLspSize() bool {
	return obj.obj.MaxLspSize != nil
}

// The maximum size in bytes of any LSP that can be transmitted over a link of equal or less  than maximum MTU size.
// SetMaxLspSize sets the uint32 value in the IsisAdvanced object
func (obj *isisAdvanced) SetMaxLspSize(value uint32) IsisAdvanced {

	obj.obj.MaxLspSize = &value
	return obj
}

// The number of seconds between transmissions of LSPs/MGROUP-PDUs.
// LspMgroupMinTransInterval returns a uint32
func (obj *isisAdvanced) LspMgroupMinTransInterval() uint32 {

	return *obj.obj.LspMgroupMinTransInterval

}

// The number of seconds between transmissions of LSPs/MGROUP-PDUs.
// LspMgroupMinTransInterval returns a uint32
func (obj *isisAdvanced) HasLspMgroupMinTransInterval() bool {
	return obj.obj.LspMgroupMinTransInterval != nil
}

// The number of seconds between transmissions of LSPs/MGROUP-PDUs.
// SetLspMgroupMinTransInterval sets the uint32 value in the IsisAdvanced object
func (obj *isisAdvanced) SetLspMgroupMinTransInterval(value uint32) IsisAdvanced {

	obj.obj.LspMgroupMinTransInterval = &value
	return obj
}

// If the Attached bit is enabled, it indicates that the ISIS router is attached to another area  or the Level 2 backbone. The purpose of an Attached-Bit is to accomplish Inter-Area Routing.  When an L1/L2 router is connected to more than one area, it sets the Attached-bit on its L1 LSP. This can cause a default route ( 0.0.0.0/0 ) to be installed by the receiving router.
// EnableAttachedBit returns a bool
func (obj *isisAdvanced) EnableAttachedBit() bool {

	return *obj.obj.EnableAttachedBit

}

// If the Attached bit is enabled, it indicates that the ISIS router is attached to another area  or the Level 2 backbone. The purpose of an Attached-Bit is to accomplish Inter-Area Routing.  When an L1/L2 router is connected to more than one area, it sets the Attached-bit on its L1 LSP. This can cause a default route ( 0.0.0.0/0 ) to be installed by the receiving router.
// EnableAttachedBit returns a bool
func (obj *isisAdvanced) HasEnableAttachedBit() bool {
	return obj.obj.EnableAttachedBit != nil
}

// If the Attached bit is enabled, it indicates that the ISIS router is attached to another area  or the Level 2 backbone. The purpose of an Attached-Bit is to accomplish Inter-Area Routing.  When an L1/L2 router is connected to more than one area, it sets the Attached-bit on its L1 LSP. This can cause a default route ( 0.0.0.0/0 ) to be installed by the receiving router.
// SetEnableAttachedBit sets the bool value in the IsisAdvanced object
func (obj *isisAdvanced) SetEnableAttachedBit(value bool) IsisAdvanced {

	obj.obj.EnableAttachedBit = &value
	return obj
}

func (obj *isisAdvanced) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MaxAreaAddresses != nil {

		if *obj.obj.MaxAreaAddresses > 254 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisAdvanced.MaxAreaAddresses <= 254 but Got %d", *obj.obj.MaxAreaAddresses))
		}

	}

	if obj.obj.AreaAddresses != nil {

		err := obj.validateHexSlice(obj.AreaAddresses())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisAdvanced.AreaAddresses"))
		}

	}

	if obj.obj.LspRefreshRate != nil {

		if *obj.obj.LspRefreshRate > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisAdvanced.LspRefreshRate <= 65535 but Got %d", *obj.obj.LspRefreshRate))
		}

	}

	if obj.obj.LspLifetime != nil {

		if *obj.obj.LspLifetime > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisAdvanced.LspLifetime <= 65535 but Got %d", *obj.obj.LspLifetime))
		}

	}

	if obj.obj.PsnpInterval != nil {

		if *obj.obj.PsnpInterval > 60000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisAdvanced.PsnpInterval <= 60000 but Got %d", *obj.obj.PsnpInterval))
		}

	}

	if obj.obj.CsnpInterval != nil {

		if *obj.obj.CsnpInterval < 1 || *obj.obj.CsnpInterval > 65535000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= IsisAdvanced.CsnpInterval <= 65535000 but Got %d", *obj.obj.CsnpInterval))
		}

	}

	if obj.obj.MaxLspSize != nil {

		if *obj.obj.MaxLspSize < 64 || *obj.obj.MaxLspSize > 9216 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("64 <= IsisAdvanced.MaxLspSize <= 9216 but Got %d", *obj.obj.MaxLspSize))
		}

	}

	if obj.obj.LspMgroupMinTransInterval != nil {

		if *obj.obj.LspMgroupMinTransInterval < 1 || *obj.obj.LspMgroupMinTransInterval > 60000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= IsisAdvanced.LspMgroupMinTransInterval <= 60000 but Got %d", *obj.obj.LspMgroupMinTransInterval))
		}

	}

}

func (obj *isisAdvanced) setDefault() {
	if obj.obj.EnableHelloPadding == nil {
		obj.SetEnableHelloPadding(true)
	}
	if obj.obj.MaxAreaAddresses == nil {
		obj.SetMaxAreaAddresses(3)
	}
	if obj.obj.LspRefreshRate == nil {
		obj.SetLspRefreshRate(600)
	}
	if obj.obj.LspLifetime == nil {
		obj.SetLspLifetime(1200)
	}
	if obj.obj.PsnpInterval == nil {
		obj.SetPsnpInterval(2000)
	}
	if obj.obj.CsnpInterval == nil {
		obj.SetCsnpInterval(10000)
	}
	if obj.obj.MaxLspSize == nil {
		obj.SetMaxLspSize(1492)
	}
	if obj.obj.LspMgroupMinTransInterval == nil {
		obj.SetLspMgroupMinTransInterval(5000)
	}
	if obj.obj.EnableAttachedBit == nil {
		obj.SetEnableAttachedBit(true)
	}

}
