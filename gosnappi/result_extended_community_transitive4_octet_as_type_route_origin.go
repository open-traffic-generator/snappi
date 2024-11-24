package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin *****
type resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin struct {
	validation
	obj          *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	marshaller   marshalResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	unMarshaller unMarshalResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
}

func NewResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin() ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin {
	obj := resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin{obj: &otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) msg() *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin {
	return obj.obj
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) setMsg(msg *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityTransitive4OctetAsTypeRouteOrigin struct {
	obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
}

type marshalResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin interface {
	// ToProto marshals ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin to protobuf object *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	ToProto() (*otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin, error)
	// ToPbText marshals ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalresultExtendedCommunityTransitive4OctetAsTypeRouteOrigin struct {
	obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
}

type unMarshalResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin interface {
	// FromProto unmarshals ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin from protobuf object *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	FromProto(msg *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) (ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin, error)
	// FromPbText unmarshals ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) Marshal() marshalResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityTransitive4OctetAsTypeRouteOrigin{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) Unmarshal() unMarshalResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityTransitive4OctetAsTypeRouteOrigin{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) ToProto() (*otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) FromProto(msg *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) (ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) ToJsonRaw() (string, error) {
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

func (m *marshalresultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) FromJson(value string) error {
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

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) Clone() (ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin()
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

// ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP. It is sent with sub-type as 0x03.
type ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin interface {
	Validation
	// msg marshals ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin to protobuf object *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// setMsg unmarshals ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin from protobuf object *otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// validate validates ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Global4ByteAs returns uint32, set in ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin.
	Global4ByteAs() uint32
	// SetGlobal4ByteAs assigns uint32 provided by user to ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	SetGlobal4ByteAs(value uint32) ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// HasGlobal4ByteAs checks if Global4ByteAs has been set in ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	HasGlobal4ByteAs() bool
	// Local2ByteAdmin returns uint32, set in ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin.
	Local2ByteAdmin() uint32
	// SetLocal2ByteAdmin assigns uint32 provided by user to ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	SetLocal2ByteAdmin(value uint32) ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// HasLocal2ByteAdmin checks if Local2ByteAdmin has been set in ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	HasLocal2ByteAdmin() bool
}

// The four octet IANA assigned AS value assigned to the Autonomous System.
// Global4ByteAs returns a uint32
func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) Global4ByteAs() uint32 {

	return *obj.obj.Global_4ByteAs

}

// The four octet IANA assigned AS value assigned to the Autonomous System.
// Global4ByteAs returns a uint32
func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) HasGlobal4ByteAs() bool {
	return obj.obj.Global_4ByteAs != nil
}

// The four octet IANA assigned AS value assigned to the Autonomous System.
// SetGlobal4ByteAs sets the uint32 value in the ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin object
func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) SetGlobal4ByteAs(value uint32) ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin {

	obj.obj.Global_4ByteAs = &value
	return obj
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) Local2ByteAdmin() uint32 {

	return *obj.obj.Local_2ByteAdmin

}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) HasLocal2ByteAdmin() bool {
	return obj.obj.Local_2ByteAdmin != nil
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// SetLocal2ByteAdmin sets the uint32 value in the ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin object
func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) SetLocal2ByteAdmin(value uint32) ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin {

	obj.obj.Local_2ByteAdmin = &value
	return obj
}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Local_2ByteAdmin != nil {

		if *obj.obj.Local_2ByteAdmin > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin.Local_2ByteAdmin <= 65535 but Got %d", *obj.obj.Local_2ByteAdmin))
		}

	}

}

func (obj *resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) setDefault() {

}
