package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget *****
type resultExtendedCommunityTransitive4OctetAsTypeRouteTarget struct {
	validation
	obj          *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	marshaller   marshalResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	unMarshaller unMarshalResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
}

func NewResultExtendedCommunityTransitive4OctetAsTypeRouteTarget() ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget {
	obj := resultExtendedCommunityTransitive4OctetAsTypeRouteTarget{obj: &otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) msg() *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget {
	return obj.obj
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) setMsg(msg *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget) ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityTransitive4OctetAsTypeRouteTarget struct {
	obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget
}

type marshalResultExtendedCommunityTransitive4OctetAsTypeRouteTarget interface {
	// ToProto marshals ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget to protobuf object *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	ToProto() (*otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget, error)
	// ToPbText marshals ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget to JSON text
	ToJson() (string, error)
}

type unMarshalresultExtendedCommunityTransitive4OctetAsTypeRouteTarget struct {
	obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget
}

type unMarshalResultExtendedCommunityTransitive4OctetAsTypeRouteTarget interface {
	// FromProto unmarshals ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget from protobuf object *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	FromProto(msg *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget) (ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget, error)
	// FromPbText unmarshals ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) Marshal() marshalResultExtendedCommunityTransitive4OctetAsTypeRouteTarget {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityTransitive4OctetAsTypeRouteTarget{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) Unmarshal() unMarshalResultExtendedCommunityTransitive4OctetAsTypeRouteTarget {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityTransitive4OctetAsTypeRouteTarget{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityTransitive4OctetAsTypeRouteTarget) ToProto() (*otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityTransitive4OctetAsTypeRouteTarget) FromProto(msg *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget) (ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityTransitive4OctetAsTypeRouteTarget) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive4OctetAsTypeRouteTarget) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityTransitive4OctetAsTypeRouteTarget) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive4OctetAsTypeRouteTarget) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityTransitive4OctetAsTypeRouteTarget) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive4OctetAsTypeRouteTarget) FromJson(value string) error {
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

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) Clone() (ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityTransitive4OctetAsTypeRouteTarget()
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

// ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP.  It is sent with sub-type as 0x02
type ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget interface {
	Validation
	// msg marshals ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget to protobuf object *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// setMsg unmarshals ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget from protobuf object *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget) ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// validate validates ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Global4ByteAs returns uint32, set in ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget.
	Global4ByteAs() uint32
	// SetGlobal4ByteAs assigns uint32 provided by user to ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	SetGlobal4ByteAs(value uint32) ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// HasGlobal4ByteAs checks if Global4ByteAs has been set in ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	HasGlobal4ByteAs() bool
	// Local2ByteAdmin returns uint32, set in ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget.
	Local2ByteAdmin() uint32
	// SetLocal2ByteAdmin assigns uint32 provided by user to ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	SetLocal2ByteAdmin(value uint32) ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// HasLocal2ByteAdmin checks if Local2ByteAdmin has been set in ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	HasLocal2ByteAdmin() bool
}

// The four octet IANA assigned AS value assigned to the Autonomous System.
// Global4ByteAs returns a uint32
func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) Global4ByteAs() uint32 {

	return *obj.obj.Global_4ByteAs

}

// The four octet IANA assigned AS value assigned to the Autonomous System.
// Global4ByteAs returns a uint32
func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) HasGlobal4ByteAs() bool {
	return obj.obj.Global_4ByteAs != nil
}

// The four octet IANA assigned AS value assigned to the Autonomous System.
// SetGlobal4ByteAs sets the uint32 value in the ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget object
func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) SetGlobal4ByteAs(value uint32) ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget {

	obj.obj.Global_4ByteAs = &value
	return obj
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) Local2ByteAdmin() uint32 {

	return *obj.obj.Local_2ByteAdmin

}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) HasLocal2ByteAdmin() bool {
	return obj.obj.Local_2ByteAdmin != nil
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// SetLocal2ByteAdmin sets the uint32 value in the ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget object
func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) SetLocal2ByteAdmin(value uint32) ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget {

	obj.obj.Local_2ByteAdmin = &value
	return obj
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Local_2ByteAdmin != nil {

		if *obj.obj.Local_2ByteAdmin > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget.Local_2ByteAdmin <= 65535 but Got %d", *obj.obj.Local_2ByteAdmin))
		}

	}

}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteTarget) setDefault() {

}
