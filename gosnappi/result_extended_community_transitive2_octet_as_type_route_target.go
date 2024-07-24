package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget *****
type resultExtendedCommunityTransitive2OctetAsTypeRouteTarget struct {
	validation
	obj          *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	marshaller   marshalResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	unMarshaller unMarshalResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
}

func NewResultExtendedCommunityTransitive2OctetAsTypeRouteTarget() ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget {
	obj := resultExtendedCommunityTransitive2OctetAsTypeRouteTarget{obj: &otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) msg() *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget {
	return obj.obj
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) setMsg(msg *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget) ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityTransitive2OctetAsTypeRouteTarget struct {
	obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget
}

type marshalResultExtendedCommunityTransitive2OctetAsTypeRouteTarget interface {
	// ToProto marshals ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget to protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	ToProto() (*otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget, error)
	// ToPbText marshals ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget to JSON text
	ToJson() (string, error)
}

type unMarshalresultExtendedCommunityTransitive2OctetAsTypeRouteTarget struct {
	obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget
}

type unMarshalResultExtendedCommunityTransitive2OctetAsTypeRouteTarget interface {
	// FromProto unmarshals ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget from protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	FromProto(msg *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget) (ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget, error)
	// FromPbText unmarshals ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) Marshal() marshalResultExtendedCommunityTransitive2OctetAsTypeRouteTarget {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityTransitive2OctetAsTypeRouteTarget{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) Unmarshal() unMarshalResultExtendedCommunityTransitive2OctetAsTypeRouteTarget {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityTransitive2OctetAsTypeRouteTarget{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityTransitive2OctetAsTypeRouteTarget) ToProto() (*otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsTypeRouteTarget) FromProto(msg *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget) (ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityTransitive2OctetAsTypeRouteTarget) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsTypeRouteTarget) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityTransitive2OctetAsTypeRouteTarget) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsTypeRouteTarget) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityTransitive2OctetAsTypeRouteTarget) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsTypeRouteTarget) FromJson(value string) error {
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

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) Clone() (ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityTransitive2OctetAsTypeRouteTarget()
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

// ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP Update message.  It is sent with sub-type as 0x02.
type ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget interface {
	Validation
	// msg marshals ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget to protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// setMsg unmarshals ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget from protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget) ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// validate validates ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Global2ByteAs returns uint32, set in ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget.
	Global2ByteAs() uint32
	// SetGlobal2ByteAs assigns uint32 provided by user to ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	SetGlobal2ByteAs(value uint32) ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// HasGlobal2ByteAs checks if Global2ByteAs has been set in ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	HasGlobal2ByteAs() bool
	// Local4ByteAdmin returns uint32, set in ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget.
	Local4ByteAdmin() uint32
	// SetLocal4ByteAdmin assigns uint32 provided by user to ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	SetLocal4ByteAdmin(value uint32) ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// HasLocal4ByteAdmin checks if Local4ByteAdmin has been set in ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	HasLocal4ByteAdmin() bool
}

// The two octet IANA assigned AS value assigned to the Autonomous System.
// Global2ByteAs returns a uint32
func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) Global2ByteAs() uint32 {

	return *obj.obj.Global_2ByteAs

}

// The two octet IANA assigned AS value assigned to the Autonomous System.
// Global2ByteAs returns a uint32
func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) HasGlobal2ByteAs() bool {
	return obj.obj.Global_2ByteAs != nil
}

// The two octet IANA assigned AS value assigned to the Autonomous System.
// SetGlobal2ByteAs sets the uint32 value in the ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget object
func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) SetGlobal2ByteAs(value uint32) ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget {

	obj.obj.Global_2ByteAs = &value
	return obj
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local4ByteAdmin returns a uint32
func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) Local4ByteAdmin() uint32 {

	return *obj.obj.Local_4ByteAdmin

}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local4ByteAdmin returns a uint32
func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) HasLocal4ByteAdmin() bool {
	return obj.obj.Local_4ByteAdmin != nil
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// SetLocal4ByteAdmin sets the uint32 value in the ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget object
func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) SetLocal4ByteAdmin(value uint32) ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget {

	obj.obj.Local_4ByteAdmin = &value
	return obj
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Global_2ByteAs != nil {

		if *obj.obj.Global_2ByteAs > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget.Global_2ByteAs <= 65535 but Got %d", *obj.obj.Global_2ByteAs))
		}

	}

}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteTarget) setDefault() {

}
