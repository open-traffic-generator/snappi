package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspSRv6SidStructure *****
type isisLspSRv6SidStructure struct {
	validation
	obj          *otg.IsisLspSRv6SidStructure
	marshaller   marshalIsisLspSRv6SidStructure
	unMarshaller unMarshalIsisLspSRv6SidStructure
}

func NewIsisLspSRv6SidStructure() IsisLspSRv6SidStructure {
	obj := isisLspSRv6SidStructure{obj: &otg.IsisLspSRv6SidStructure{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspSRv6SidStructure) msg() *otg.IsisLspSRv6SidStructure {
	return obj.obj
}

func (obj *isisLspSRv6SidStructure) setMsg(msg *otg.IsisLspSRv6SidStructure) IsisLspSRv6SidStructure {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspSRv6SidStructure struct {
	obj *isisLspSRv6SidStructure
}

type marshalIsisLspSRv6SidStructure interface {
	// ToProto marshals IsisLspSRv6SidStructure to protobuf object *otg.IsisLspSRv6SidStructure
	ToProto() (*otg.IsisLspSRv6SidStructure, error)
	// ToPbText marshals IsisLspSRv6SidStructure to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspSRv6SidStructure to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspSRv6SidStructure to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspSRv6SidStructure struct {
	obj *isisLspSRv6SidStructure
}

type unMarshalIsisLspSRv6SidStructure interface {
	// FromProto unmarshals IsisLspSRv6SidStructure from protobuf object *otg.IsisLspSRv6SidStructure
	FromProto(msg *otg.IsisLspSRv6SidStructure) (IsisLspSRv6SidStructure, error)
	// FromPbText unmarshals IsisLspSRv6SidStructure from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspSRv6SidStructure from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspSRv6SidStructure from JSON text
	FromJson(value string) error
}

func (obj *isisLspSRv6SidStructure) Marshal() marshalIsisLspSRv6SidStructure {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspSRv6SidStructure{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspSRv6SidStructure) Unmarshal() unMarshalIsisLspSRv6SidStructure {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspSRv6SidStructure{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspSRv6SidStructure) ToProto() (*otg.IsisLspSRv6SidStructure, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspSRv6SidStructure) FromProto(msg *otg.IsisLspSRv6SidStructure) (IsisLspSRv6SidStructure, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspSRv6SidStructure) ToPbText() (string, error) {
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

func (m *unMarshalisisLspSRv6SidStructure) FromPbText(value string) error {
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

func (m *marshalisisLspSRv6SidStructure) ToYaml() (string, error) {
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

func (m *unMarshalisisLspSRv6SidStructure) FromYaml(value string) error {
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

func (m *marshalisisLspSRv6SidStructure) ToJson() (string, error) {
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

func (m *unMarshalisisLspSRv6SidStructure) FromJson(value string) error {
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

func (obj *isisLspSRv6SidStructure) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspSRv6SidStructure) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspSRv6SidStructure) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspSRv6SidStructure) Clone() (IsisLspSRv6SidStructure, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspSRv6SidStructure()
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

// IsisLspSRv6SidStructure is sRv6 SID Structure Sub-Sub-TLV (type 1) learned from a received LSP. Describes the internal bit-field decomposition of an SRv6 SID, enabling routers to independently interpret each component of the 128-bit SID value. The sum of all four length fields must not exceed 128 bits. Reference: RFC 9352 Section 9.
type IsisLspSRv6SidStructure interface {
	Validation
	// msg marshals IsisLspSRv6SidStructure to protobuf object *otg.IsisLspSRv6SidStructure
	// and doesn't set defaults
	msg() *otg.IsisLspSRv6SidStructure
	// setMsg unmarshals IsisLspSRv6SidStructure from protobuf object *otg.IsisLspSRv6SidStructure
	// and doesn't set defaults
	setMsg(*otg.IsisLspSRv6SidStructure) IsisLspSRv6SidStructure
	// provides marshal interface
	Marshal() marshalIsisLspSRv6SidStructure
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspSRv6SidStructure
	// validate validates IsisLspSRv6SidStructure
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspSRv6SidStructure, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LocatorBlockLength returns uint32, set in IsisLspSRv6SidStructure.
	LocatorBlockLength() uint32
	// SetLocatorBlockLength assigns uint32 provided by user to IsisLspSRv6SidStructure
	SetLocatorBlockLength(value uint32) IsisLspSRv6SidStructure
	// HasLocatorBlockLength checks if LocatorBlockLength has been set in IsisLspSRv6SidStructure
	HasLocatorBlockLength() bool
	// LocatorNodeLength returns uint32, set in IsisLspSRv6SidStructure.
	LocatorNodeLength() uint32
	// SetLocatorNodeLength assigns uint32 provided by user to IsisLspSRv6SidStructure
	SetLocatorNodeLength(value uint32) IsisLspSRv6SidStructure
	// HasLocatorNodeLength checks if LocatorNodeLength has been set in IsisLspSRv6SidStructure
	HasLocatorNodeLength() bool
	// FunctionLength returns uint32, set in IsisLspSRv6SidStructure.
	FunctionLength() uint32
	// SetFunctionLength assigns uint32 provided by user to IsisLspSRv6SidStructure
	SetFunctionLength(value uint32) IsisLspSRv6SidStructure
	// HasFunctionLength checks if FunctionLength has been set in IsisLspSRv6SidStructure
	HasFunctionLength() bool
	// ArgumentLength returns uint32, set in IsisLspSRv6SidStructure.
	ArgumentLength() uint32
	// SetArgumentLength assigns uint32 provided by user to IsisLspSRv6SidStructure
	SetArgumentLength(value uint32) IsisLspSRv6SidStructure
	// HasArgumentLength checks if ArgumentLength has been set in IsisLspSRv6SidStructure
	HasArgumentLength() bool
}

// Number of bits in the Locator Block (LB) field of the SID. The common IPv6 prefix shared by all SIDs in the same SRv6 domain block (e.g., 32 bits in the uSID F3216 format).
// LocatorBlockLength returns a uint32
func (obj *isisLspSRv6SidStructure) LocatorBlockLength() uint32 {

	return *obj.obj.LocatorBlockLength

}

// Number of bits in the Locator Block (LB) field of the SID. The common IPv6 prefix shared by all SIDs in the same SRv6 domain block (e.g., 32 bits in the uSID F3216 format).
// LocatorBlockLength returns a uint32
func (obj *isisLspSRv6SidStructure) HasLocatorBlockLength() bool {
	return obj.obj.LocatorBlockLength != nil
}

// Number of bits in the Locator Block (LB) field of the SID. The common IPv6 prefix shared by all SIDs in the same SRv6 domain block (e.g., 32 bits in the uSID F3216 format).
// SetLocatorBlockLength sets the uint32 value in the IsisLspSRv6SidStructure object
func (obj *isisLspSRv6SidStructure) SetLocatorBlockLength(value uint32) IsisLspSRv6SidStructure {

	obj.obj.LocatorBlockLength = &value
	return obj
}

// Number of bits in the Locator Node (LN) field of the SID. Identifies the specific node within the Locator Block (e.g., 16 bits in the uSID F3216 format).
// LocatorNodeLength returns a uint32
func (obj *isisLspSRv6SidStructure) LocatorNodeLength() uint32 {

	return *obj.obj.LocatorNodeLength

}

// Number of bits in the Locator Node (LN) field of the SID. Identifies the specific node within the Locator Block (e.g., 16 bits in the uSID F3216 format).
// LocatorNodeLength returns a uint32
func (obj *isisLspSRv6SidStructure) HasLocatorNodeLength() bool {
	return obj.obj.LocatorNodeLength != nil
}

// Number of bits in the Locator Node (LN) field of the SID. Identifies the specific node within the Locator Block (e.g., 16 bits in the uSID F3216 format).
// SetLocatorNodeLength sets the uint32 value in the IsisLspSRv6SidStructure object
func (obj *isisLspSRv6SidStructure) SetLocatorNodeLength(value uint32) IsisLspSRv6SidStructure {

	obj.obj.LocatorNodeLength = &value
	return obj
}

// Number of bits in the Function field of the SID. Identifies the specific endpoint behavior on this node (e.g., 16 bits in the uSID F3216 format).
// FunctionLength returns a uint32
func (obj *isisLspSRv6SidStructure) FunctionLength() uint32 {

	return *obj.obj.FunctionLength

}

// Number of bits in the Function field of the SID. Identifies the specific endpoint behavior on this node (e.g., 16 bits in the uSID F3216 format).
// FunctionLength returns a uint32
func (obj *isisLspSRv6SidStructure) HasFunctionLength() bool {
	return obj.obj.FunctionLength != nil
}

// Number of bits in the Function field of the SID. Identifies the specific endpoint behavior on this node (e.g., 16 bits in the uSID F3216 format).
// SetFunctionLength sets the uint32 value in the IsisLspSRv6SidStructure object
func (obj *isisLspSRv6SidStructure) SetFunctionLength(value uint32) IsisLspSRv6SidStructure {

	obj.obj.FunctionLength = &value
	return obj
}

// Number of bits in the Argument field of the SID. Used for VPN table IDs or compressed next-SID encodings in Micro-SID stacking. 0 for standard SIDs without argument (RFC 9800).
// ArgumentLength returns a uint32
func (obj *isisLspSRv6SidStructure) ArgumentLength() uint32 {

	return *obj.obj.ArgumentLength

}

// Number of bits in the Argument field of the SID. Used for VPN table IDs or compressed next-SID encodings in Micro-SID stacking. 0 for standard SIDs without argument (RFC 9800).
// ArgumentLength returns a uint32
func (obj *isisLspSRv6SidStructure) HasArgumentLength() bool {
	return obj.obj.ArgumentLength != nil
}

// Number of bits in the Argument field of the SID. Used for VPN table IDs or compressed next-SID encodings in Micro-SID stacking. 0 for standard SIDs without argument (RFC 9800).
// SetArgumentLength sets the uint32 value in the IsisLspSRv6SidStructure object
func (obj *isisLspSRv6SidStructure) SetArgumentLength(value uint32) IsisLspSRv6SidStructure {

	obj.obj.ArgumentLength = &value
	return obj
}

func (obj *isisLspSRv6SidStructure) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LocatorBlockLength != nil {

		if *obj.obj.LocatorBlockLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspSRv6SidStructure.LocatorBlockLength <= 128 but Got %d", *obj.obj.LocatorBlockLength))
		}

	}

	if obj.obj.LocatorNodeLength != nil {

		if *obj.obj.LocatorNodeLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspSRv6SidStructure.LocatorNodeLength <= 128 but Got %d", *obj.obj.LocatorNodeLength))
		}

	}

	if obj.obj.FunctionLength != nil {

		if *obj.obj.FunctionLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspSRv6SidStructure.FunctionLength <= 128 but Got %d", *obj.obj.FunctionLength))
		}

	}

	if obj.obj.ArgumentLength != nil {

		if *obj.obj.ArgumentLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspSRv6SidStructure.ArgumentLength <= 128 but Got %d", *obj.obj.ArgumentLength))
		}

	}

}

func (obj *isisLspSRv6SidStructure) setDefault() {

}
