package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernet *****
type ultraEthernet struct {
	validation
	obj          *otg.UltraEthernet
	marshaller   marshalUltraEthernet
	unMarshaller unMarshalUltraEthernet
	phyHolder    UltraEthernetPhy
	llrHolder    UltraEthernetLlr
	cbfcHolder   UltraEthernetCbfc
}

func NewUltraEthernet() UltraEthernet {
	obj := ultraEthernet{obj: &otg.UltraEthernet{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernet) msg() *otg.UltraEthernet {
	return obj.obj
}

func (obj *ultraEthernet) setMsg(msg *otg.UltraEthernet) UltraEthernet {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernet struct {
	obj *ultraEthernet
}

type marshalUltraEthernet interface {
	// ToProto marshals UltraEthernet to protobuf object *otg.UltraEthernet
	ToProto() (*otg.UltraEthernet, error)
	// ToPbText marshals UltraEthernet to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernet to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernet to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernet struct {
	obj *ultraEthernet
}

type unMarshalUltraEthernet interface {
	// FromProto unmarshals UltraEthernet from protobuf object *otg.UltraEthernet
	FromProto(msg *otg.UltraEthernet) (UltraEthernet, error)
	// FromPbText unmarshals UltraEthernet from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernet from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernet from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernet) Marshal() marshalUltraEthernet {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernet{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernet) Unmarshal() unMarshalUltraEthernet {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernet{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernet) ToProto() (*otg.UltraEthernet, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernet) FromProto(msg *otg.UltraEthernet) (UltraEthernet, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernet) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernet) FromPbText(value string) error {
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

func (m *marshalultraEthernet) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernet) FromYaml(value string) error {
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

func (m *marshalultraEthernet) ToJson() (string, error) {
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

func (m *unMarshalultraEthernet) FromJson(value string) error {
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

func (obj *ultraEthernet) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernet) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernet) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernet) Clone() (UltraEthernet, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernet()
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

func (obj *ultraEthernet) setNil() {
	obj.phyHolder = nil
	obj.llrHolder = nil
	obj.cbfcHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UltraEthernet is a container for Ultra Ethernet (UE) link layer settings applied to one or
// more test ports.
//
// The Ultra Ethernet Consortium (UEC) link layer defines optional-to-implement
// features that enhance performance and reliability of the Ethernet link:
//
// - Link Layer Retry (LLR): frame based link level retransmission of lost frames.
// - Credit-based Flow Control (CBFC): per virtual channel (VC) credit based
// alternative to priority based flow control (PFC).
// - PHY Control Ordered Set (CtlOS) timing: transmit/receive spacing controls
// for the ordered sets that carry LLR and CBFC messages.
//
// These settings usually vary across a variety of test ports and most likely
// won't be portable.
type UltraEthernet interface {
	Validation
	// msg marshals UltraEthernet to protobuf object *otg.UltraEthernet
	// and doesn't set defaults
	msg() *otg.UltraEthernet
	// setMsg unmarshals UltraEthernet from protobuf object *otg.UltraEthernet
	// and doesn't set defaults
	setMsg(*otg.UltraEthernet) UltraEthernet
	// provides marshal interface
	Marshal() marshalUltraEthernet
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernet
	// validate validates UltraEthernet
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernet, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortNames returns []string, set in UltraEthernet.
	PortNames() []string
	// SetPortNames assigns []string provided by user to UltraEthernet
	SetPortNames(value []string) UltraEthernet
	// Phy returns UltraEthernetPhy, set in UltraEthernet.
	// UltraEthernetPhy is ultra Ethernet PHY layer settings that control the insertion and validation
	// of Control Ordered Sets (CtlOS). CtlOS carry the LLR (LLR_ACK, LLR_NACK,
	// LLR_INIT, LLR_INIT_ECHO) and CBFC (CF_Update) messages in the 64B/66B block
	// stream.
	//
	// Reference: UE-Specification-1.0.3 Section 6.2 (Control Ordered Sets); CtlOS
	// spacing rules Section 5.1.3.1.1.
	Phy() UltraEthernetPhy
	// SetPhy assigns UltraEthernetPhy provided by user to UltraEthernet.
	// UltraEthernetPhy is ultra Ethernet PHY layer settings that control the insertion and validation
	// of Control Ordered Sets (CtlOS). CtlOS carry the LLR (LLR_ACK, LLR_NACK,
	// LLR_INIT, LLR_INIT_ECHO) and CBFC (CF_Update) messages in the 64B/66B block
	// stream.
	//
	// Reference: UE-Specification-1.0.3 Section 6.2 (Control Ordered Sets); CtlOS
	// spacing rules Section 5.1.3.1.1.
	SetPhy(value UltraEthernetPhy) UltraEthernet
	// HasPhy checks if Phy has been set in UltraEthernet
	HasPhy() bool
	// Llr returns UltraEthernetLlr, set in UltraEthernet.
	// UltraEthernetLlr is ultra Ethernet Link Layer Retry (LLR) settings. LLR provides lossless link
	// operation by assigning a sequence number to LLR-eligible frames, holding them
	// in a replay buffer, and retransmitting them when the link partner reports a
	// loss via LLR_NACK or when a replay timer expires.
	//
	// Reference: UE-Specification-1.0.3 Section 5.1; configuration registers
	// Table 5-9, counters Table 5-13.
	Llr() UltraEthernetLlr
	// SetLlr assigns UltraEthernetLlr provided by user to UltraEthernet.
	// UltraEthernetLlr is ultra Ethernet Link Layer Retry (LLR) settings. LLR provides lossless link
	// operation by assigning a sequence number to LLR-eligible frames, holding them
	// in a replay buffer, and retransmitting them when the link partner reports a
	// loss via LLR_NACK or when a replay timer expires.
	//
	// Reference: UE-Specification-1.0.3 Section 5.1; configuration registers
	// Table 5-9, counters Table 5-13.
	SetLlr(value UltraEthernetLlr) UltraEthernet
	// HasLlr checks if Llr has been set in UltraEthernet
	HasLlr() bool
	// Cbfc returns UltraEthernetCbfc, set in UltraEthernet.
	// UltraEthernetCbfc is ultra Ethernet Credit-based Flow Control (CBFC) settings. CBFC is a per
	// virtual channel (VC) credit based flow control mechanism and is an alternative
	// to priority based flow control (PFC). The sender and receiver directions of a
	// link are configured independently.
	//
	// Reference: UE-Specification-1.0.3 Section 5.2.
	Cbfc() UltraEthernetCbfc
	// SetCbfc assigns UltraEthernetCbfc provided by user to UltraEthernet.
	// UltraEthernetCbfc is ultra Ethernet Credit-based Flow Control (CBFC) settings. CBFC is a per
	// virtual channel (VC) credit based flow control mechanism and is an alternative
	// to priority based flow control (PFC). The sender and receiver directions of a
	// link are configured independently.
	//
	// Reference: UE-Specification-1.0.3 Section 5.2.
	SetCbfc(value UltraEthernetCbfc) UltraEthernet
	// HasCbfc checks if Cbfc has been set in UltraEthernet
	HasCbfc() bool
	// Name returns string, set in UltraEthernet.
	Name() string
	// SetName assigns string provided by user to UltraEthernet
	SetName(value string) UltraEthernet
	setNil()
}

// A list of unique names of port objects that will share the Ultra Ethernet
// link layer settings.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortNames returns a []string
func (obj *ultraEthernet) PortNames() []string {
	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	return obj.obj.PortNames
}

// A list of unique names of port objects that will share the Ultra Ethernet
// link layer settings.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetPortNames sets the []string value in the UltraEthernet object
func (obj *ultraEthernet) SetPortNames(value []string) UltraEthernet {

	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	obj.obj.PortNames = value

	return obj
}

// Ultra Ethernet PHY Control Ordered Set (CtlOS) timing settings.
// Phy returns a UltraEthernetPhy
func (obj *ultraEthernet) Phy() UltraEthernetPhy {
	if obj.obj.Phy == nil {
		obj.obj.Phy = NewUltraEthernetPhy().msg()
	}
	if obj.phyHolder == nil {
		obj.phyHolder = &ultraEthernetPhy{obj: obj.obj.Phy}
	}
	return obj.phyHolder
}

// Ultra Ethernet PHY Control Ordered Set (CtlOS) timing settings.
// Phy returns a UltraEthernetPhy
func (obj *ultraEthernet) HasPhy() bool {
	return obj.obj.Phy != nil
}

// Ultra Ethernet PHY Control Ordered Set (CtlOS) timing settings.
// SetPhy sets the UltraEthernetPhy value in the UltraEthernet object
func (obj *ultraEthernet) SetPhy(value UltraEthernetPhy) UltraEthernet {

	obj.phyHolder = nil
	obj.obj.Phy = value.msg()

	return obj
}

// Ultra Ethernet Link Layer Retry (LLR) settings.
// Llr returns a UltraEthernetLlr
func (obj *ultraEthernet) Llr() UltraEthernetLlr {
	if obj.obj.Llr == nil {
		obj.obj.Llr = NewUltraEthernetLlr().msg()
	}
	if obj.llrHolder == nil {
		obj.llrHolder = &ultraEthernetLlr{obj: obj.obj.Llr}
	}
	return obj.llrHolder
}

// Ultra Ethernet Link Layer Retry (LLR) settings.
// Llr returns a UltraEthernetLlr
func (obj *ultraEthernet) HasLlr() bool {
	return obj.obj.Llr != nil
}

// Ultra Ethernet Link Layer Retry (LLR) settings.
// SetLlr sets the UltraEthernetLlr value in the UltraEthernet object
func (obj *ultraEthernet) SetLlr(value UltraEthernetLlr) UltraEthernet {

	obj.llrHolder = nil
	obj.obj.Llr = value.msg()

	return obj
}

// Ultra Ethernet Credit-based Flow Control (CBFC) settings.
// Cbfc returns a UltraEthernetCbfc
func (obj *ultraEthernet) Cbfc() UltraEthernetCbfc {
	if obj.obj.Cbfc == nil {
		obj.obj.Cbfc = NewUltraEthernetCbfc().msg()
	}
	if obj.cbfcHolder == nil {
		obj.cbfcHolder = &ultraEthernetCbfc{obj: obj.obj.Cbfc}
	}
	return obj.cbfcHolder
}

// Ultra Ethernet Credit-based Flow Control (CBFC) settings.
// Cbfc returns a UltraEthernetCbfc
func (obj *ultraEthernet) HasCbfc() bool {
	return obj.obj.Cbfc != nil
}

// Ultra Ethernet Credit-based Flow Control (CBFC) settings.
// SetCbfc sets the UltraEthernetCbfc value in the UltraEthernet object
func (obj *ultraEthernet) SetCbfc(value UltraEthernetCbfc) UltraEthernet {

	obj.cbfcHolder = nil
	obj.obj.Cbfc = value.msg()

	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *ultraEthernet) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the UltraEthernet object
func (obj *ultraEthernet) SetName(value string) UltraEthernet {

	obj.obj.Name = &value
	return obj
}

func (obj *ultraEthernet) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Phy != nil {

		obj.Phy().validateObj(vObj, set_default)
	}

	if obj.obj.Llr != nil {

		obj.Llr().validateObj(vObj, set_default)
	}

	if obj.obj.Cbfc != nil {

		obj.Cbfc().validateObj(vObj, set_default)
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface UltraEthernet")
	}
}

func (obj *ultraEthernet) setDefault() {

}
