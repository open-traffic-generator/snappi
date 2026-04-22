package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRv6SidStructure *****
type isisSRv6SidStructure struct {
	validation
	obj          *otg.IsisSRv6SidStructure
	marshaller   marshalIsisSRv6SidStructure
	unMarshaller unMarshalIsisSRv6SidStructure
}

func NewIsisSRv6SidStructure() IsisSRv6SidStructure {
	obj := isisSRv6SidStructure{obj: &otg.IsisSRv6SidStructure{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRv6SidStructure) msg() *otg.IsisSRv6SidStructure {
	return obj.obj
}

func (obj *isisSRv6SidStructure) setMsg(msg *otg.IsisSRv6SidStructure) IsisSRv6SidStructure {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRv6SidStructure struct {
	obj *isisSRv6SidStructure
}

type marshalIsisSRv6SidStructure interface {
	// ToProto marshals IsisSRv6SidStructure to protobuf object *otg.IsisSRv6SidStructure
	ToProto() (*otg.IsisSRv6SidStructure, error)
	// ToPbText marshals IsisSRv6SidStructure to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRv6SidStructure to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRv6SidStructure to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRv6SidStructure struct {
	obj *isisSRv6SidStructure
}

type unMarshalIsisSRv6SidStructure interface {
	// FromProto unmarshals IsisSRv6SidStructure from protobuf object *otg.IsisSRv6SidStructure
	FromProto(msg *otg.IsisSRv6SidStructure) (IsisSRv6SidStructure, error)
	// FromPbText unmarshals IsisSRv6SidStructure from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRv6SidStructure from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRv6SidStructure from JSON text
	FromJson(value string) error
}

func (obj *isisSRv6SidStructure) Marshal() marshalIsisSRv6SidStructure {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRv6SidStructure{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRv6SidStructure) Unmarshal() unMarshalIsisSRv6SidStructure {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRv6SidStructure{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRv6SidStructure) ToProto() (*otg.IsisSRv6SidStructure, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRv6SidStructure) FromProto(msg *otg.IsisSRv6SidStructure) (IsisSRv6SidStructure, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRv6SidStructure) ToPbText() (string, error) {
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

func (m *unMarshalisisSRv6SidStructure) FromPbText(value string) error {
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

func (m *marshalisisSRv6SidStructure) ToYaml() (string, error) {
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

func (m *unMarshalisisSRv6SidStructure) FromYaml(value string) error {
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

func (m *marshalisisSRv6SidStructure) ToJson() (string, error) {
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

func (m *unMarshalisisSRv6SidStructure) FromJson(value string) error {
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

func (obj *isisSRv6SidStructure) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRv6SidStructure) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRv6SidStructure) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRv6SidStructure) Clone() (IsisSRv6SidStructure, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRv6SidStructure()
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

// IsisSRv6SidStructure is sRv6 SID Structure Sub-Sub-TLV (type 1), carried within SRv6 SID Sub-TLVs (End SID type 5, End.X SID type 43/44). Describes the internal bit-field decomposition of the SRv6 SID value so that receiving routers can interpret each field independently. The four length fields (lb_length + ln_length + function_length + argument_length) MUST NOT exceed 128 bits. Required when advertising Micro-SID (uSID) SIDs to describe the compressed encoding. Example for common uSID F3216 format:
// lb_length=32, ln_length=16, function_length=16, argument_length=0
// Reference: RFC 9352 Section 9, RFC 9800.
type IsisSRv6SidStructure interface {
	Validation
	// msg marshals IsisSRv6SidStructure to protobuf object *otg.IsisSRv6SidStructure
	// and doesn't set defaults
	msg() *otg.IsisSRv6SidStructure
	// setMsg unmarshals IsisSRv6SidStructure from protobuf object *otg.IsisSRv6SidStructure
	// and doesn't set defaults
	setMsg(*otg.IsisSRv6SidStructure) IsisSRv6SidStructure
	// provides marshal interface
	Marshal() marshalIsisSRv6SidStructure
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRv6SidStructure
	// validate validates IsisSRv6SidStructure
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRv6SidStructure, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LocatorBlockLength returns uint32, set in IsisSRv6SidStructure.
	LocatorBlockLength() uint32
	// SetLocatorBlockLength assigns uint32 provided by user to IsisSRv6SidStructure
	SetLocatorBlockLength(value uint32) IsisSRv6SidStructure
	// HasLocatorBlockLength checks if LocatorBlockLength has been set in IsisSRv6SidStructure
	HasLocatorBlockLength() bool
	// LocatorNodeLength returns uint32, set in IsisSRv6SidStructure.
	LocatorNodeLength() uint32
	// SetLocatorNodeLength assigns uint32 provided by user to IsisSRv6SidStructure
	SetLocatorNodeLength(value uint32) IsisSRv6SidStructure
	// HasLocatorNodeLength checks if LocatorNodeLength has been set in IsisSRv6SidStructure
	HasLocatorNodeLength() bool
	// FunctionLength returns uint32, set in IsisSRv6SidStructure.
	FunctionLength() uint32
	// SetFunctionLength assigns uint32 provided by user to IsisSRv6SidStructure
	SetFunctionLength(value uint32) IsisSRv6SidStructure
	// HasFunctionLength checks if FunctionLength has been set in IsisSRv6SidStructure
	HasFunctionLength() bool
	// ArgumentLength returns uint32, set in IsisSRv6SidStructure.
	ArgumentLength() uint32
	// SetArgumentLength assigns uint32 provided by user to IsisSRv6SidStructure
	SetArgumentLength(value uint32) IsisSRv6SidStructure
	// HasArgumentLength checks if ArgumentLength has been set in IsisSRv6SidStructure
	HasArgumentLength() bool
}

// Number of bits in the Locator Block (LB) field of the SID. The Locator Block is the common IPv6 prefix shared across all SIDs within the same SRv6 domain block (e.g., 32 bits in the uSID F3216 format). Minimum value 1.
// LocatorBlockLength returns a uint32
func (obj *isisSRv6SidStructure) LocatorBlockLength() uint32 {

	return *obj.obj.LocatorBlockLength

}

// Number of bits in the Locator Block (LB) field of the SID. The Locator Block is the common IPv6 prefix shared across all SIDs within the same SRv6 domain block (e.g., 32 bits in the uSID F3216 format). Minimum value 1.
// LocatorBlockLength returns a uint32
func (obj *isisSRv6SidStructure) HasLocatorBlockLength() bool {
	return obj.obj.LocatorBlockLength != nil
}

// Number of bits in the Locator Block (LB) field of the SID. The Locator Block is the common IPv6 prefix shared across all SIDs within the same SRv6 domain block (e.g., 32 bits in the uSID F3216 format). Minimum value 1.
// SetLocatorBlockLength sets the uint32 value in the IsisSRv6SidStructure object
func (obj *isisSRv6SidStructure) SetLocatorBlockLength(value uint32) IsisSRv6SidStructure {

	obj.obj.LocatorBlockLength = &value
	return obj
}

// Number of bits in the Locator Node (LN) field of the SID. Identifies the specific node within the Locator Block (e.g., 16 bits in the uSID F3216 format, providing up to 65534 addressable nodes per domain block).
// LocatorNodeLength returns a uint32
func (obj *isisSRv6SidStructure) LocatorNodeLength() uint32 {

	return *obj.obj.LocatorNodeLength

}

// Number of bits in the Locator Node (LN) field of the SID. Identifies the specific node within the Locator Block (e.g., 16 bits in the uSID F3216 format, providing up to 65534 addressable nodes per domain block).
// LocatorNodeLength returns a uint32
func (obj *isisSRv6SidStructure) HasLocatorNodeLength() bool {
	return obj.obj.LocatorNodeLength != nil
}

// Number of bits in the Locator Node (LN) field of the SID. Identifies the specific node within the Locator Block (e.g., 16 bits in the uSID F3216 format, providing up to 65534 addressable nodes per domain block).
// SetLocatorNodeLength sets the uint32 value in the IsisSRv6SidStructure object
func (obj *isisSRv6SidStructure) SetLocatorNodeLength(value uint32) IsisSRv6SidStructure {

	obj.obj.LocatorNodeLength = &value
	return obj
}

// Number of bits in the Function field of the SID. Identifies the specific endpoint behavior instantiated on this node (e.g., 16 bits in the uSID F3216 format, supporting up to 65534 distinct functions per node).
// FunctionLength returns a uint32
func (obj *isisSRv6SidStructure) FunctionLength() uint32 {

	return *obj.obj.FunctionLength

}

// Number of bits in the Function field of the SID. Identifies the specific endpoint behavior instantiated on this node (e.g., 16 bits in the uSID F3216 format, supporting up to 65534 distinct functions per node).
// FunctionLength returns a uint32
func (obj *isisSRv6SidStructure) HasFunctionLength() bool {
	return obj.obj.FunctionLength != nil
}

// Number of bits in the Function field of the SID. Identifies the specific endpoint behavior instantiated on this node (e.g., 16 bits in the uSID F3216 format, supporting up to 65534 distinct functions per node).
// SetFunctionLength sets the uint32 value in the IsisSRv6SidStructure object
func (obj *isisSRv6SidStructure) SetFunctionLength(value uint32) IsisSRv6SidStructure {

	obj.obj.FunctionLength = &value
	return obj
}

// Number of bits in the Argument field of the SID. Carries additional per-behavior information such as VPN table identifiers or compressed next-SID encodings for Micro-SID stacking. Set to 0 for standard SIDs that carry no argument. When non-zero in a uSID context, enables packing multiple micro-segment identifiers within a single 128-bit SID carrier (RFC 9800).
// ArgumentLength returns a uint32
func (obj *isisSRv6SidStructure) ArgumentLength() uint32 {

	return *obj.obj.ArgumentLength

}

// Number of bits in the Argument field of the SID. Carries additional per-behavior information such as VPN table identifiers or compressed next-SID encodings for Micro-SID stacking. Set to 0 for standard SIDs that carry no argument. When non-zero in a uSID context, enables packing multiple micro-segment identifiers within a single 128-bit SID carrier (RFC 9800).
// ArgumentLength returns a uint32
func (obj *isisSRv6SidStructure) HasArgumentLength() bool {
	return obj.obj.ArgumentLength != nil
}

// Number of bits in the Argument field of the SID. Carries additional per-behavior information such as VPN table identifiers or compressed next-SID encodings for Micro-SID stacking. Set to 0 for standard SIDs that carry no argument. When non-zero in a uSID context, enables packing multiple micro-segment identifiers within a single 128-bit SID carrier (RFC 9800).
// SetArgumentLength sets the uint32 value in the IsisSRv6SidStructure object
func (obj *isisSRv6SidStructure) SetArgumentLength(value uint32) IsisSRv6SidStructure {

	obj.obj.ArgumentLength = &value
	return obj
}

func (obj *isisSRv6SidStructure) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LocatorBlockLength != nil {

		if *obj.obj.LocatorBlockLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6SidStructure.LocatorBlockLength <= 128 but Got %d", *obj.obj.LocatorBlockLength))
		}

	}

	if obj.obj.LocatorNodeLength != nil {

		if *obj.obj.LocatorNodeLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6SidStructure.LocatorNodeLength <= 128 but Got %d", *obj.obj.LocatorNodeLength))
		}

	}

	if obj.obj.FunctionLength != nil {

		if *obj.obj.FunctionLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6SidStructure.FunctionLength <= 128 but Got %d", *obj.obj.FunctionLength))
		}

	}

	if obj.obj.ArgumentLength != nil {

		if *obj.obj.ArgumentLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6SidStructure.ArgumentLength <= 128 but Got %d", *obj.obj.ArgumentLength))
		}

	}

}

func (obj *isisSRv6SidStructure) setDefault() {
	if obj.obj.LocatorBlockLength == nil {
		obj.SetLocatorBlockLength(32)
	}
	if obj.obj.LocatorNodeLength == nil {
		obj.SetLocatorNodeLength(16)
	}
	if obj.obj.FunctionLength == nil {
		obj.SetFunctionLength(16)
	}
	if obj.obj.ArgumentLength == nil {
		obj.SetArgumentLength(0)
	}

}
