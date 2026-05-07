package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6Routing *****
type flowIpv6Routing struct {
	validation
	obj                      *otg.FlowIpv6Routing
	marshaller               marshalFlowIpv6Routing
	unMarshaller             unMarshalFlowIpv6Routing
	segmentRoutingHolder     FlowIpv6SegmentRouting
	segmentRoutingUsidHolder FlowIpv6SegmentRoutingUsid
}

func NewFlowIpv6Routing() FlowIpv6Routing {
	obj := flowIpv6Routing{obj: &otg.FlowIpv6Routing{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6Routing) msg() *otg.FlowIpv6Routing {
	return obj.obj
}

func (obj *flowIpv6Routing) setMsg(msg *otg.FlowIpv6Routing) FlowIpv6Routing {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6Routing struct {
	obj *flowIpv6Routing
}

type marshalFlowIpv6Routing interface {
	// ToProto marshals FlowIpv6Routing to protobuf object *otg.FlowIpv6Routing
	ToProto() (*otg.FlowIpv6Routing, error)
	// ToPbText marshals FlowIpv6Routing to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6Routing to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6Routing to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6Routing struct {
	obj *flowIpv6Routing
}

type unMarshalFlowIpv6Routing interface {
	// FromProto unmarshals FlowIpv6Routing from protobuf object *otg.FlowIpv6Routing
	FromProto(msg *otg.FlowIpv6Routing) (FlowIpv6Routing, error)
	// FromPbText unmarshals FlowIpv6Routing from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6Routing from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6Routing from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6Routing) Marshal() marshalFlowIpv6Routing {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6Routing{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6Routing) Unmarshal() unMarshalFlowIpv6Routing {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6Routing{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6Routing) ToProto() (*otg.FlowIpv6Routing, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6Routing) FromProto(msg *otg.FlowIpv6Routing) (FlowIpv6Routing, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6Routing) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6Routing) FromPbText(value string) error {
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

func (m *marshalflowIpv6Routing) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6Routing) FromYaml(value string) error {
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

func (m *marshalflowIpv6Routing) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6Routing) FromJson(value string) error {
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

func (obj *flowIpv6Routing) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6Routing) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6Routing) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6Routing) Clone() (FlowIpv6Routing, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6Routing()
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

func (obj *flowIpv6Routing) setNil() {
	obj.segmentRoutingHolder = nil
	obj.segmentRoutingUsidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6Routing is iPv6 routing packet headers.
type FlowIpv6Routing interface {
	Validation
	// msg marshals FlowIpv6Routing to protobuf object *otg.FlowIpv6Routing
	// and doesn't set defaults
	msg() *otg.FlowIpv6Routing
	// setMsg unmarshals FlowIpv6Routing from protobuf object *otg.FlowIpv6Routing
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6Routing) FlowIpv6Routing
	// provides marshal interface
	Marshal() marshalFlowIpv6Routing
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6Routing
	// validate validates FlowIpv6Routing
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6Routing, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowIpv6RoutingChoiceEnum, set in FlowIpv6Routing
	Choice() FlowIpv6RoutingChoiceEnum
	// setChoice assigns FlowIpv6RoutingChoiceEnum provided by user to FlowIpv6Routing
	setChoice(value FlowIpv6RoutingChoiceEnum) FlowIpv6Routing
	// HasChoice checks if Choice has been set in FlowIpv6Routing
	HasChoice() bool
	// SegmentRouting returns FlowIpv6SegmentRouting, set in FlowIpv6Routing.
	// FlowIpv6SegmentRouting is iPv6 Segment Routing Header (SRH) with full 128-bit Segment IDs
	// (RFC 8754 Section 2). Each segment list entry carries a complete SRv6 SID
	// (locator + function + argument) as a 128-bit IPv6 address.
	// The Routing Type field is always 4 (RFC 8754 Section 2.1).
	//
	// Packet stack example - 2-SID path (H.Encaps, Node1 then Node2 then receiver):
	//
	// ```
	// +--------------------------------------------------+
	// | Ethernet                                         |
	// +--------------------------------------------------+
	// | Outer IPv6                                       |
	// |   next_header : 43 (Routing Extension Header)   |
	// |   dst         : fc00:0:1::  (Segment[1], active)|
	// +--------------------------------------------------+
	// | SRH (Routing Type 4)                            |
	// |   segments_left : 1  (auto = num segments - 1)  |
	// |   last_entry    : 1  (auto = num segments - 1)  |
	// |   flags, tag    : 0                             |
	// |   Segment[0]    : fc00:0:2::  (last hop)        |
	// |   Segment[1]    : fc00:0:1::  (first hop)       |
	// +--------------------------------------------------+
	// | Inner payload (IPv4, IPv6, etc.)                 |
	// +--------------------------------------------------+
	// ```
	//
	// The outer IPv6 dst is set to Segment[segments_left], which is the first
	// SID to visit. At each hop, segments_left is decremented and dst is
	// updated to the next SID. The segment list is encoded in reverse path
	// order: Segment[0] is the last hop, Segment[n-1] is the first hop.
	SegmentRouting() FlowIpv6SegmentRouting
	// SetSegmentRouting assigns FlowIpv6SegmentRouting provided by user to FlowIpv6Routing.
	// FlowIpv6SegmentRouting is iPv6 Segment Routing Header (SRH) with full 128-bit Segment IDs
	// (RFC 8754 Section 2). Each segment list entry carries a complete SRv6 SID
	// (locator + function + argument) as a 128-bit IPv6 address.
	// The Routing Type field is always 4 (RFC 8754 Section 2.1).
	//
	// Packet stack example - 2-SID path (H.Encaps, Node1 then Node2 then receiver):
	//
	// ```
	// +--------------------------------------------------+
	// | Ethernet                                         |
	// +--------------------------------------------------+
	// | Outer IPv6                                       |
	// |   next_header : 43 (Routing Extension Header)   |
	// |   dst         : fc00:0:1::  (Segment[1], active)|
	// +--------------------------------------------------+
	// | SRH (Routing Type 4)                            |
	// |   segments_left : 1  (auto = num segments - 1)  |
	// |   last_entry    : 1  (auto = num segments - 1)  |
	// |   flags, tag    : 0                             |
	// |   Segment[0]    : fc00:0:2::  (last hop)        |
	// |   Segment[1]    : fc00:0:1::  (first hop)       |
	// +--------------------------------------------------+
	// | Inner payload (IPv4, IPv6, etc.)                 |
	// +--------------------------------------------------+
	// ```
	//
	// The outer IPv6 dst is set to Segment[segments_left], which is the first
	// SID to visit. At each hop, segments_left is decremented and dst is
	// updated to the next SID. The segment list is encoded in reverse path
	// order: Segment[0] is the last hop, Segment[n-1] is the first hop.
	SetSegmentRouting(value FlowIpv6SegmentRouting) FlowIpv6Routing
	// HasSegmentRouting checks if SegmentRouting has been set in FlowIpv6Routing
	HasSegmentRouting() bool
	// SegmentRoutingUsid returns FlowIpv6SegmentRoutingUsid, set in FlowIpv6Routing.
	// FlowIpv6SegmentRoutingUsid is iPv6 Segment Routing Header (SRH) with uSID containers (RFC 9800 Section 4).
	// Each segment list entry is a pre-computed 128-bit uSID container value
	// supplied as a plain IPv6 address. The container packs multiple Micro-SIDs
	// but appears on the wire as a single 128-bit address - no decomposition
	// fields (locator block length, node length, function length) are present
	// in the packet. The user pre-computes the container value and supplies it
	// directly, allowing raw traffic generation independent of any control
	// plane configuration on the emulated device.
	// The Routing Type field is always 4 (RFC 8754 Section 2.1).
	//
	// Packet stack example - 1 container (F3216, Node2 + Node3 packed into one container):
	//
	// Container value assembly (done by user, outside the schema):
	// lb=fc00::/32 | node2-fn1 (0002:0001) | node3-fn1 (0003:0001) | zeros
	// => fc00:0:2:1:3:1::
	//
	// ```
	// +--------------------------------------------------+
	// | Ethernet                                         |
	// +--------------------------------------------------+
	// | Outer IPv6                                       |
	// |   next_header : 43 (Routing Extension Header)   |
	// |   dst         : fc00:0:2:1:3:1:: (container)    |
	// +--------------------------------------------------+
	// | SRH (Routing Type 4)                            |
	// |   segments_left : 0  (auto)                     |
	// |   last_entry    : 0  (auto)                     |
	// |   flags, tag    : 0                             |
	// |   Segment[0]    : fc00:0:2:1:3:1::              |
	// +--------------------------------------------------+
	// | Inner payload (IPv4, IPv6, etc.)                 |
	// +--------------------------------------------------+
	// ```
	//
	// At Node2: reads first uSID (node2-fn1), shifts container left,
	// outer dst becomes fc00:0:3:1::, forwards.
	// At Node3: reads next uSID (node3-fn1), container exhausted,
	// strips SRH, delivers payload.
	//
	// Packet stack example - 2 containers (3-hop path: Node2 + Node3, then Node4):
	//
	// ```
	// +--------------------------------------------------+
	// | Ethernet                                         |
	// +--------------------------------------------------+
	// | Outer IPv6                                       |
	// |   dst : fc00:0:2:1:3:1:: (Segment[1], active)   |
	// +--------------------------------------------------+
	// | SRH (Routing Type 4)                            |
	// |   segments_left : 1  (auto)                     |
	// |   last_entry    : 1  (auto)                     |
	// |   flags, tag    : 0                             |
	// |   Segment[0]    : fc00:0:4:1::      (last hop)  |
	// |   Segment[1]    : fc00:0:2:1:3:1::  (first)     |
	// +--------------------------------------------------+
	// | Inner payload (IPv4, IPv6, etc.)                 |
	// +--------------------------------------------------+
	// ```
	//
	// The segment list is encoded in reverse path order: Segment[0] is the
	// last container (last hop), Segment[n-1] is the first container to visit.
	SegmentRoutingUsid() FlowIpv6SegmentRoutingUsid
	// SetSegmentRoutingUsid assigns FlowIpv6SegmentRoutingUsid provided by user to FlowIpv6Routing.
	// FlowIpv6SegmentRoutingUsid is iPv6 Segment Routing Header (SRH) with uSID containers (RFC 9800 Section 4).
	// Each segment list entry is a pre-computed 128-bit uSID container value
	// supplied as a plain IPv6 address. The container packs multiple Micro-SIDs
	// but appears on the wire as a single 128-bit address - no decomposition
	// fields (locator block length, node length, function length) are present
	// in the packet. The user pre-computes the container value and supplies it
	// directly, allowing raw traffic generation independent of any control
	// plane configuration on the emulated device.
	// The Routing Type field is always 4 (RFC 8754 Section 2.1).
	//
	// Packet stack example - 1 container (F3216, Node2 + Node3 packed into one container):
	//
	// Container value assembly (done by user, outside the schema):
	// lb=fc00::/32 | node2-fn1 (0002:0001) | node3-fn1 (0003:0001) | zeros
	// => fc00:0:2:1:3:1::
	//
	// ```
	// +--------------------------------------------------+
	// | Ethernet                                         |
	// +--------------------------------------------------+
	// | Outer IPv6                                       |
	// |   next_header : 43 (Routing Extension Header)   |
	// |   dst         : fc00:0:2:1:3:1:: (container)    |
	// +--------------------------------------------------+
	// | SRH (Routing Type 4)                            |
	// |   segments_left : 0  (auto)                     |
	// |   last_entry    : 0  (auto)                     |
	// |   flags, tag    : 0                             |
	// |   Segment[0]    : fc00:0:2:1:3:1::              |
	// +--------------------------------------------------+
	// | Inner payload (IPv4, IPv6, etc.)                 |
	// +--------------------------------------------------+
	// ```
	//
	// At Node2: reads first uSID (node2-fn1), shifts container left,
	// outer dst becomes fc00:0:3:1::, forwards.
	// At Node3: reads next uSID (node3-fn1), container exhausted,
	// strips SRH, delivers payload.
	//
	// Packet stack example - 2 containers (3-hop path: Node2 + Node3, then Node4):
	//
	// ```
	// +--------------------------------------------------+
	// | Ethernet                                         |
	// +--------------------------------------------------+
	// | Outer IPv6                                       |
	// |   dst : fc00:0:2:1:3:1:: (Segment[1], active)   |
	// +--------------------------------------------------+
	// | SRH (Routing Type 4)                            |
	// |   segments_left : 1  (auto)                     |
	// |   last_entry    : 1  (auto)                     |
	// |   flags, tag    : 0                             |
	// |   Segment[0]    : fc00:0:4:1::      (last hop)  |
	// |   Segment[1]    : fc00:0:2:1:3:1::  (first)     |
	// +--------------------------------------------------+
	// | Inner payload (IPv4, IPv6, etc.)                 |
	// +--------------------------------------------------+
	// ```
	//
	// The segment list is encoded in reverse path order: Segment[0] is the
	// last container (last hop), Segment[n-1] is the first container to visit.
	SetSegmentRoutingUsid(value FlowIpv6SegmentRoutingUsid) FlowIpv6Routing
	// HasSegmentRoutingUsid checks if SegmentRoutingUsid has been set in FlowIpv6Routing
	HasSegmentRoutingUsid() bool
	setNil()
}

type FlowIpv6RoutingChoiceEnum string

// Enum of Choice on FlowIpv6Routing
var FlowIpv6RoutingChoice = struct {
	SEGMENT_ROUTING      FlowIpv6RoutingChoiceEnum
	SEGMENT_ROUTING_USID FlowIpv6RoutingChoiceEnum
}{
	SEGMENT_ROUTING:      FlowIpv6RoutingChoiceEnum("segment_routing"),
	SEGMENT_ROUTING_USID: FlowIpv6RoutingChoiceEnum("segment_routing_usid"),
}

func (obj *flowIpv6Routing) Choice() FlowIpv6RoutingChoiceEnum {
	return FlowIpv6RoutingChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowIpv6Routing) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowIpv6Routing) setChoice(value FlowIpv6RoutingChoiceEnum) FlowIpv6Routing {
	intValue, ok := otg.FlowIpv6Routing_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowIpv6RoutingChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowIpv6Routing_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.SegmentRoutingUsid = nil
	obj.segmentRoutingUsidHolder = nil
	obj.obj.SegmentRouting = nil
	obj.segmentRoutingHolder = nil

	if value == FlowIpv6RoutingChoice.SEGMENT_ROUTING {
		obj.obj.SegmentRouting = NewFlowIpv6SegmentRouting().msg()
	}

	if value == FlowIpv6RoutingChoice.SEGMENT_ROUTING_USID {
		obj.obj.SegmentRoutingUsid = NewFlowIpv6SegmentRoutingUsid().msg()
	}

	return obj
}

// description is TBD
// SegmentRouting returns a FlowIpv6SegmentRouting
func (obj *flowIpv6Routing) SegmentRouting() FlowIpv6SegmentRouting {
	if obj.obj.SegmentRouting == nil {
		obj.setChoice(FlowIpv6RoutingChoice.SEGMENT_ROUTING)
	}
	if obj.segmentRoutingHolder == nil {
		obj.segmentRoutingHolder = &flowIpv6SegmentRouting{obj: obj.obj.SegmentRouting}
	}
	return obj.segmentRoutingHolder
}

// description is TBD
// SegmentRouting returns a FlowIpv6SegmentRouting
func (obj *flowIpv6Routing) HasSegmentRouting() bool {
	return obj.obj.SegmentRouting != nil
}

// description is TBD
// SetSegmentRouting sets the FlowIpv6SegmentRouting value in the FlowIpv6Routing object
func (obj *flowIpv6Routing) SetSegmentRouting(value FlowIpv6SegmentRouting) FlowIpv6Routing {
	obj.setChoice(FlowIpv6RoutingChoice.SEGMENT_ROUTING)
	obj.segmentRoutingHolder = nil
	obj.obj.SegmentRouting = value.msg()

	return obj
}

// description is TBD
// SegmentRoutingUsid returns a FlowIpv6SegmentRoutingUsid
func (obj *flowIpv6Routing) SegmentRoutingUsid() FlowIpv6SegmentRoutingUsid {
	if obj.obj.SegmentRoutingUsid == nil {
		obj.setChoice(FlowIpv6RoutingChoice.SEGMENT_ROUTING_USID)
	}
	if obj.segmentRoutingUsidHolder == nil {
		obj.segmentRoutingUsidHolder = &flowIpv6SegmentRoutingUsid{obj: obj.obj.SegmentRoutingUsid}
	}
	return obj.segmentRoutingUsidHolder
}

// description is TBD
// SegmentRoutingUsid returns a FlowIpv6SegmentRoutingUsid
func (obj *flowIpv6Routing) HasSegmentRoutingUsid() bool {
	return obj.obj.SegmentRoutingUsid != nil
}

// description is TBD
// SetSegmentRoutingUsid sets the FlowIpv6SegmentRoutingUsid value in the FlowIpv6Routing object
func (obj *flowIpv6Routing) SetSegmentRoutingUsid(value FlowIpv6SegmentRoutingUsid) FlowIpv6Routing {
	obj.setChoice(FlowIpv6RoutingChoice.SEGMENT_ROUTING_USID)
	obj.segmentRoutingUsidHolder = nil
	obj.obj.SegmentRoutingUsid = value.msg()

	return obj
}

func (obj *flowIpv6Routing) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SegmentRouting != nil {

		obj.SegmentRouting().validateObj(vObj, set_default)
	}

	if obj.obj.SegmentRoutingUsid != nil {

		obj.SegmentRoutingUsid().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6Routing) setDefault() {
	var choices_set int = 0
	var choice FlowIpv6RoutingChoiceEnum

	if obj.obj.SegmentRouting != nil {
		choices_set += 1
		choice = FlowIpv6RoutingChoice.SEGMENT_ROUTING
	}

	if obj.obj.SegmentRoutingUsid != nil {
		choices_set += 1
		choice = FlowIpv6RoutingChoice.SEGMENT_ROUTING_USID
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowIpv6RoutingChoice.SEGMENT_ROUTING)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowIpv6Routing")
			}
		} else {
			intVal := otg.FlowIpv6Routing_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowIpv6Routing_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
