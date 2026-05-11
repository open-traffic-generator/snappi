package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRv6Msd *****
type isisSRv6Msd struct {
	validation
	obj                *otg.IsisSRv6Msd
	marshaller         marshalIsisSRv6Msd
	unMarshaller       unMarshalIsisSRv6Msd
	maxSlHolder        IsisSRv6MsdValue
	maxEndPopSrhHolder IsisSRv6MsdValue
	maxHEncapsHolder   IsisSRv6MsdValue
	maxEndDSrhHolder   IsisSRv6MsdValue
	maxTInsertHolder   IsisSRv6MsdValue
	maxTEncapsHolder   IsisSRv6MsdValue
}

func NewIsisSRv6Msd() IsisSRv6Msd {
	obj := isisSRv6Msd{obj: &otg.IsisSRv6Msd{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRv6Msd) msg() *otg.IsisSRv6Msd {
	return obj.obj
}

func (obj *isisSRv6Msd) setMsg(msg *otg.IsisSRv6Msd) IsisSRv6Msd {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRv6Msd struct {
	obj *isisSRv6Msd
}

type marshalIsisSRv6Msd interface {
	// ToProto marshals IsisSRv6Msd to protobuf object *otg.IsisSRv6Msd
	ToProto() (*otg.IsisSRv6Msd, error)
	// ToPbText marshals IsisSRv6Msd to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRv6Msd to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRv6Msd to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRv6Msd struct {
	obj *isisSRv6Msd
}

type unMarshalIsisSRv6Msd interface {
	// FromProto unmarshals IsisSRv6Msd from protobuf object *otg.IsisSRv6Msd
	FromProto(msg *otg.IsisSRv6Msd) (IsisSRv6Msd, error)
	// FromPbText unmarshals IsisSRv6Msd from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRv6Msd from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRv6Msd from JSON text
	FromJson(value string) error
}

func (obj *isisSRv6Msd) Marshal() marshalIsisSRv6Msd {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRv6Msd{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRv6Msd) Unmarshal() unMarshalIsisSRv6Msd {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRv6Msd{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRv6Msd) ToProto() (*otg.IsisSRv6Msd, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRv6Msd) FromProto(msg *otg.IsisSRv6Msd) (IsisSRv6Msd, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRv6Msd) ToPbText() (string, error) {
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

func (m *unMarshalisisSRv6Msd) FromPbText(value string) error {
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

func (m *marshalisisSRv6Msd) ToYaml() (string, error) {
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

func (m *unMarshalisisSRv6Msd) FromYaml(value string) error {
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

func (m *marshalisisSRv6Msd) ToJson() (string, error) {
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

func (m *unMarshalisisSRv6Msd) FromJson(value string) error {
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

func (obj *isisSRv6Msd) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRv6Msd) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRv6Msd) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRv6Msd) Clone() (IsisSRv6Msd, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRv6Msd()
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

func (obj *isisSRv6Msd) setNil() {
	obj.maxSlHolder = nil
	obj.maxEndPopSrhHolder = nil
	obj.maxHEncapsHolder = nil
	obj.maxEndDSrhHolder = nil
	obj.maxTInsertHolder = nil
	obj.maxTEncapsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisSRv6Msd is sRv6 Maximum SID Depth (MSD) sub-TLV values, used for both node-level advertisement (carried in IS-IS Router CAPABILITY TLV 242 via Node MSD sub-TLV type 23, RFC 9352 Section 6) and link-level advertisement (carried in Extended IS Reachability TLV via Link MSD sub-TLV type 15, RFC 9352 Section 6). When a property is present, the corresponding MSD sub-TLV is advertised with the configured value. Omit a property to suppress advertisement of that MSD type. Reference: RFC 9352 Section 6.
type IsisSRv6Msd interface {
	Validation
	// msg marshals IsisSRv6Msd to protobuf object *otg.IsisSRv6Msd
	// and doesn't set defaults
	msg() *otg.IsisSRv6Msd
	// setMsg unmarshals IsisSRv6Msd from protobuf object *otg.IsisSRv6Msd
	// and doesn't set defaults
	setMsg(*otg.IsisSRv6Msd) IsisSRv6Msd
	// provides marshal interface
	Marshal() marshalIsisSRv6Msd
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRv6Msd
	// validate validates IsisSRv6Msd
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRv6Msd, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MaxSl returns IsisSRv6MsdValue, set in IsisSRv6Msd.
	// IsisSRv6MsdValue is a single MSD (Maximum SID Depth) value entry (RFC 9352 Section 6). When this object is present, the parent MSD type is advertised with the configured value. Omit the object to suppress advertisement of that MSD type.
	MaxSl() IsisSRv6MsdValue
	// SetMaxSl assigns IsisSRv6MsdValue provided by user to IsisSRv6Msd.
	// IsisSRv6MsdValue is a single MSD (Maximum SID Depth) value entry (RFC 9352 Section 6). When this object is present, the parent MSD type is advertised with the configured value. Omit the object to suppress advertisement of that MSD type.
	SetMaxSl(value IsisSRv6MsdValue) IsisSRv6Msd
	// HasMaxSl checks if MaxSl has been set in IsisSRv6Msd
	HasMaxSl() bool
	// MaxEndPopSrh returns IsisSRv6MsdValue, set in IsisSRv6Msd.
	// IsisSRv6MsdValue is a single MSD (Maximum SID Depth) value entry (RFC 9352 Section 6). When this object is present, the parent MSD type is advertised with the configured value. Omit the object to suppress advertisement of that MSD type.
	MaxEndPopSrh() IsisSRv6MsdValue
	// SetMaxEndPopSrh assigns IsisSRv6MsdValue provided by user to IsisSRv6Msd.
	// IsisSRv6MsdValue is a single MSD (Maximum SID Depth) value entry (RFC 9352 Section 6). When this object is present, the parent MSD type is advertised with the configured value. Omit the object to suppress advertisement of that MSD type.
	SetMaxEndPopSrh(value IsisSRv6MsdValue) IsisSRv6Msd
	// HasMaxEndPopSrh checks if MaxEndPopSrh has been set in IsisSRv6Msd
	HasMaxEndPopSrh() bool
	// MaxHEncaps returns IsisSRv6MsdValue, set in IsisSRv6Msd.
	// IsisSRv6MsdValue is a single MSD (Maximum SID Depth) value entry (RFC 9352 Section 6). When this object is present, the parent MSD type is advertised with the configured value. Omit the object to suppress advertisement of that MSD type.
	MaxHEncaps() IsisSRv6MsdValue
	// SetMaxHEncaps assigns IsisSRv6MsdValue provided by user to IsisSRv6Msd.
	// IsisSRv6MsdValue is a single MSD (Maximum SID Depth) value entry (RFC 9352 Section 6). When this object is present, the parent MSD type is advertised with the configured value. Omit the object to suppress advertisement of that MSD type.
	SetMaxHEncaps(value IsisSRv6MsdValue) IsisSRv6Msd
	// HasMaxHEncaps checks if MaxHEncaps has been set in IsisSRv6Msd
	HasMaxHEncaps() bool
	// MaxEndDSrh returns IsisSRv6MsdValue, set in IsisSRv6Msd.
	// IsisSRv6MsdValue is a single MSD (Maximum SID Depth) value entry (RFC 9352 Section 6). When this object is present, the parent MSD type is advertised with the configured value. Omit the object to suppress advertisement of that MSD type.
	MaxEndDSrh() IsisSRv6MsdValue
	// SetMaxEndDSrh assigns IsisSRv6MsdValue provided by user to IsisSRv6Msd.
	// IsisSRv6MsdValue is a single MSD (Maximum SID Depth) value entry (RFC 9352 Section 6). When this object is present, the parent MSD type is advertised with the configured value. Omit the object to suppress advertisement of that MSD type.
	SetMaxEndDSrh(value IsisSRv6MsdValue) IsisSRv6Msd
	// HasMaxEndDSrh checks if MaxEndDSrh has been set in IsisSRv6Msd
	HasMaxEndDSrh() bool
	// MaxTInsert returns IsisSRv6MsdValue, set in IsisSRv6Msd.
	// IsisSRv6MsdValue is a single MSD (Maximum SID Depth) value entry (RFC 9352 Section 6). When this object is present, the parent MSD type is advertised with the configured value. Omit the object to suppress advertisement of that MSD type.
	MaxTInsert() IsisSRv6MsdValue
	// SetMaxTInsert assigns IsisSRv6MsdValue provided by user to IsisSRv6Msd.
	// IsisSRv6MsdValue is a single MSD (Maximum SID Depth) value entry (RFC 9352 Section 6). When this object is present, the parent MSD type is advertised with the configured value. Omit the object to suppress advertisement of that MSD type.
	SetMaxTInsert(value IsisSRv6MsdValue) IsisSRv6Msd
	// HasMaxTInsert checks if MaxTInsert has been set in IsisSRv6Msd
	HasMaxTInsert() bool
	// MaxTEncaps returns IsisSRv6MsdValue, set in IsisSRv6Msd.
	// IsisSRv6MsdValue is a single MSD (Maximum SID Depth) value entry (RFC 9352 Section 6). When this object is present, the parent MSD type is advertised with the configured value. Omit the object to suppress advertisement of that MSD type.
	MaxTEncaps() IsisSRv6MsdValue
	// SetMaxTEncaps assigns IsisSRv6MsdValue provided by user to IsisSRv6Msd.
	// IsisSRv6MsdValue is a single MSD (Maximum SID Depth) value entry (RFC 9352 Section 6). When this object is present, the parent MSD type is advertised with the configured value. Omit the object to suppress advertisement of that MSD type.
	SetMaxTEncaps(value IsisSRv6MsdValue) IsisSRv6Msd
	// HasMaxTEncaps checks if MaxTEncaps has been set in IsisSRv6Msd
	HasMaxTEncaps() bool
	setNil()
}

// When present, advertises Maximum Segments Left (Max-SL, MSD type 41, RFC 9352 Section 6) - the maximum value of the Segments Left field in an SRH that this node can correctly process.
// MaxSl returns a IsisSRv6MsdValue
func (obj *isisSRv6Msd) MaxSl() IsisSRv6MsdValue {
	if obj.obj.MaxSl == nil {
		obj.obj.MaxSl = NewIsisSRv6MsdValue().msg()
	}
	if obj.maxSlHolder == nil {
		obj.maxSlHolder = &isisSRv6MsdValue{obj: obj.obj.MaxSl}
	}
	return obj.maxSlHolder
}

// When present, advertises Maximum Segments Left (Max-SL, MSD type 41, RFC 9352 Section 6) - the maximum value of the Segments Left field in an SRH that this node can correctly process.
// MaxSl returns a IsisSRv6MsdValue
func (obj *isisSRv6Msd) HasMaxSl() bool {
	return obj.obj.MaxSl != nil
}

// When present, advertises Maximum Segments Left (Max-SL, MSD type 41, RFC 9352 Section 6) - the maximum value of the Segments Left field in an SRH that this node can correctly process.
// SetMaxSl sets the IsisSRv6MsdValue value in the IsisSRv6Msd object
func (obj *isisSRv6Msd) SetMaxSl(value IsisSRv6MsdValue) IsisSRv6Msd {

	obj.maxSlHolder = nil
	obj.obj.MaxSl = value.msg()

	return obj
}

// When present, advertises Maximum End Pop SRH (Max-End-Pop, MSD type 42, RFC 9352 Section 6) - the maximum size of the SRH this node can remove when processing an End.Pop operation.
// MaxEndPopSrh returns a IsisSRv6MsdValue
func (obj *isisSRv6Msd) MaxEndPopSrh() IsisSRv6MsdValue {
	if obj.obj.MaxEndPopSrh == nil {
		obj.obj.MaxEndPopSrh = NewIsisSRv6MsdValue().msg()
	}
	if obj.maxEndPopSrhHolder == nil {
		obj.maxEndPopSrhHolder = &isisSRv6MsdValue{obj: obj.obj.MaxEndPopSrh}
	}
	return obj.maxEndPopSrhHolder
}

// When present, advertises Maximum End Pop SRH (Max-End-Pop, MSD type 42, RFC 9352 Section 6) - the maximum size of the SRH this node can remove when processing an End.Pop operation.
// MaxEndPopSrh returns a IsisSRv6MsdValue
func (obj *isisSRv6Msd) HasMaxEndPopSrh() bool {
	return obj.obj.MaxEndPopSrh != nil
}

// When present, advertises Maximum End Pop SRH (Max-End-Pop, MSD type 42, RFC 9352 Section 6) - the maximum size of the SRH this node can remove when processing an End.Pop operation.
// SetMaxEndPopSrh sets the IsisSRv6MsdValue value in the IsisSRv6Msd object
func (obj *isisSRv6Msd) SetMaxEndPopSrh(value IsisSRv6MsdValue) IsisSRv6Msd {

	obj.maxEndPopSrhHolder = nil
	obj.obj.MaxEndPopSrh = value.msg()

	return obj
}

// When present, advertises Maximum H.Encaps (Max-H.Encaps, MSD type 44, RFC 9352 Section 6) - the maximum number of SRv6 SID encapsulations this node can perform as a headend.
// MaxHEncaps returns a IsisSRv6MsdValue
func (obj *isisSRv6Msd) MaxHEncaps() IsisSRv6MsdValue {
	if obj.obj.MaxHEncaps == nil {
		obj.obj.MaxHEncaps = NewIsisSRv6MsdValue().msg()
	}
	if obj.maxHEncapsHolder == nil {
		obj.maxHEncapsHolder = &isisSRv6MsdValue{obj: obj.obj.MaxHEncaps}
	}
	return obj.maxHEncapsHolder
}

// When present, advertises Maximum H.Encaps (Max-H.Encaps, MSD type 44, RFC 9352 Section 6) - the maximum number of SRv6 SID encapsulations this node can perform as a headend.
// MaxHEncaps returns a IsisSRv6MsdValue
func (obj *isisSRv6Msd) HasMaxHEncaps() bool {
	return obj.obj.MaxHEncaps != nil
}

// When present, advertises Maximum H.Encaps (Max-H.Encaps, MSD type 44, RFC 9352 Section 6) - the maximum number of SRv6 SID encapsulations this node can perform as a headend.
// SetMaxHEncaps sets the IsisSRv6MsdValue value in the IsisSRv6Msd object
func (obj *isisSRv6Msd) SetMaxHEncaps(value IsisSRv6MsdValue) IsisSRv6Msd {

	obj.maxHEncapsHolder = nil
	obj.obj.MaxHEncaps = value.msg()

	return obj
}

// When present, advertises Maximum End.D SRH (Max-End-D, MSD type 45, RFC 9352 Section 6) - the maximum size of the SRH this node can process when performing an End.D decapsulation function.
// MaxEndDSrh returns a IsisSRv6MsdValue
func (obj *isisSRv6Msd) MaxEndDSrh() IsisSRv6MsdValue {
	if obj.obj.MaxEndDSrh == nil {
		obj.obj.MaxEndDSrh = NewIsisSRv6MsdValue().msg()
	}
	if obj.maxEndDSrhHolder == nil {
		obj.maxEndDSrhHolder = &isisSRv6MsdValue{obj: obj.obj.MaxEndDSrh}
	}
	return obj.maxEndDSrhHolder
}

// When present, advertises Maximum End.D SRH (Max-End-D, MSD type 45, RFC 9352 Section 6) - the maximum size of the SRH this node can process when performing an End.D decapsulation function.
// MaxEndDSrh returns a IsisSRv6MsdValue
func (obj *isisSRv6Msd) HasMaxEndDSrh() bool {
	return obj.obj.MaxEndDSrh != nil
}

// When present, advertises Maximum End.D SRH (Max-End-D, MSD type 45, RFC 9352 Section 6) - the maximum size of the SRH this node can process when performing an End.D decapsulation function.
// SetMaxEndDSrh sets the IsisSRv6MsdValue value in the IsisSRv6Msd object
func (obj *isisSRv6Msd) SetMaxEndDSrh(value IsisSRv6MsdValue) IsisSRv6Msd {

	obj.maxEndDSrhHolder = nil
	obj.obj.MaxEndDSrh = value.msg()

	return obj
}

// When present, advertises Maximum T.Insert (Max-T.Insert, MSD type 46, RFC 9352 Section 6) - the maximum size of the SRH this node can insert using a T.Insert operation.
// MaxTInsert returns a IsisSRv6MsdValue
func (obj *isisSRv6Msd) MaxTInsert() IsisSRv6MsdValue {
	if obj.obj.MaxTInsert == nil {
		obj.obj.MaxTInsert = NewIsisSRv6MsdValue().msg()
	}
	if obj.maxTInsertHolder == nil {
		obj.maxTInsertHolder = &isisSRv6MsdValue{obj: obj.obj.MaxTInsert}
	}
	return obj.maxTInsertHolder
}

// When present, advertises Maximum T.Insert (Max-T.Insert, MSD type 46, RFC 9352 Section 6) - the maximum size of the SRH this node can insert using a T.Insert operation.
// MaxTInsert returns a IsisSRv6MsdValue
func (obj *isisSRv6Msd) HasMaxTInsert() bool {
	return obj.obj.MaxTInsert != nil
}

// When present, advertises Maximum T.Insert (Max-T.Insert, MSD type 46, RFC 9352 Section 6) - the maximum size of the SRH this node can insert using a T.Insert operation.
// SetMaxTInsert sets the IsisSRv6MsdValue value in the IsisSRv6Msd object
func (obj *isisSRv6Msd) SetMaxTInsert(value IsisSRv6MsdValue) IsisSRv6Msd {

	obj.maxTInsertHolder = nil
	obj.obj.MaxTInsert = value.msg()

	return obj
}

// When present, advertises Maximum T.Encaps (Max-T.Encaps, MSD type 47, RFC 9352 Section 6) - the maximum size of the SRH this node can insert using a T.Encaps operation.
// MaxTEncaps returns a IsisSRv6MsdValue
func (obj *isisSRv6Msd) MaxTEncaps() IsisSRv6MsdValue {
	if obj.obj.MaxTEncaps == nil {
		obj.obj.MaxTEncaps = NewIsisSRv6MsdValue().msg()
	}
	if obj.maxTEncapsHolder == nil {
		obj.maxTEncapsHolder = &isisSRv6MsdValue{obj: obj.obj.MaxTEncaps}
	}
	return obj.maxTEncapsHolder
}

// When present, advertises Maximum T.Encaps (Max-T.Encaps, MSD type 47, RFC 9352 Section 6) - the maximum size of the SRH this node can insert using a T.Encaps operation.
// MaxTEncaps returns a IsisSRv6MsdValue
func (obj *isisSRv6Msd) HasMaxTEncaps() bool {
	return obj.obj.MaxTEncaps != nil
}

// When present, advertises Maximum T.Encaps (Max-T.Encaps, MSD type 47, RFC 9352 Section 6) - the maximum size of the SRH this node can insert using a T.Encaps operation.
// SetMaxTEncaps sets the IsisSRv6MsdValue value in the IsisSRv6Msd object
func (obj *isisSRv6Msd) SetMaxTEncaps(value IsisSRv6MsdValue) IsisSRv6Msd {

	obj.maxTEncapsHolder = nil
	obj.obj.MaxTEncaps = value.msg()

	return obj
}

func (obj *isisSRv6Msd) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MaxSl != nil {

		obj.MaxSl().validateObj(vObj, set_default)
	}

	if obj.obj.MaxEndPopSrh != nil {

		obj.MaxEndPopSrh().validateObj(vObj, set_default)
	}

	if obj.obj.MaxHEncaps != nil {

		obj.MaxHEncaps().validateObj(vObj, set_default)
	}

	if obj.obj.MaxEndDSrh != nil {

		obj.MaxEndDSrh().validateObj(vObj, set_default)
	}

	if obj.obj.MaxTInsert != nil {

		obj.MaxTInsert().validateObj(vObj, set_default)
	}

	if obj.obj.MaxTEncaps != nil {

		obj.MaxTEncaps().validateObj(vObj, set_default)
	}

}

func (obj *isisSRv6Msd) setDefault() {

}
