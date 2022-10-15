package gosnappi

import "fmt"

func (obj *config) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Ports) != 0 {

		if setDefault {
			obj.Ports().clearHolderSlice()
			for _, item := range obj.obj.Ports {
				newObj := newPort(obj.validator)
				newObj.self().obj = item
				obj.Ports().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ports().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Lags) != 0 {

		if setDefault {
			obj.Lags().clearHolderSlice()
			for _, item := range obj.obj.Lags {
				newObj := newLag(obj.validator)
				newObj.self().obj = item
				obj.Lags().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Lags().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Layer1) != 0 {

		if setDefault {
			obj.Layer1().clearHolderSlice()
			for _, item := range obj.obj.Layer1 {
				newObj := newLayer1(obj.validator)
				newObj.self().obj = item
				obj.Layer1().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Layer1().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Captures) != 0 {

		if setDefault {
			obj.Captures().clearHolderSlice()
			for _, item := range obj.obj.Captures {
				newObj := newCapture(obj.validator)
				newObj.self().obj = item
				obj.Captures().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Captures().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Devices) != 0 {

		if setDefault {
			obj.Devices().clearHolderSlice()
			for _, item := range obj.obj.Devices {
				newObj := newDevice(obj.validator)
				newObj.self().obj = item
				obj.Devices().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Devices().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Flows) != 0 {

		if setDefault {
			obj.Flows().clearHolderSlice()
			for _, item := range obj.obj.Flows {
				newObj := newFlow(obj.validator)
				newObj.self().obj = item
				obj.Flows().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Flows().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.Events != nil {

		obj.Events().validateObj(setDefault)
	}

	if obj.obj.Options != nil {

		obj.Options().validateObj(setDefault)
	}

}

func (obj *config) checkUnique() {

}

func (obj *config) checkConstraint() {

}

func (obj *transmitState) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// State is required
	if obj.obj.State.Number() == 0 {
		obj.errors = append(obj.errors, "State is required field on interface TransmitState")
	}
}

func (obj *transmitState) checkUnique() {

}

func (obj *transmitState) checkConstraint() {

	flownamesCons := []string{
		"flow.Name",
	}

	for _, v := range obj.FlowNames() {
		if !obj.validateConstraint(flownamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid flow.Name type", v),
			)
		}
	}

}

func (obj *linkState) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// State is required
	if obj.obj.State.Number() == 0 {
		obj.errors = append(obj.errors, "State is required field on interface LinkState")
	}
}

func (obj *linkState) checkUnique() {

}

func (obj *linkState) checkConstraint() {

	portnamesCons := []string{
		"port.Name",
	}

	for _, v := range obj.PortNames() {
		if !obj.validateConstraint(portnamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid port.Name type", v),
			)
		}
	}

}

func (obj *captureState) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// State is required
	if obj.obj.State.Number() == 0 {
		obj.errors = append(obj.errors, "State is required field on interface CaptureState")
	}
}

func (obj *captureState) checkUnique() {

}

func (obj *captureState) checkConstraint() {

	portnamesCons := []string{
		"port.Name",
	}

	for _, v := range obj.PortNames() {
		if !obj.validateConstraint(portnamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid port.Name type", v),
			)
		}
	}

}

func (obj *flowsUpdate) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// PropertyNames is required
	if obj.obj.PropertyNames == nil {
		obj.errors = append(obj.errors, "PropertyNames is required field on interface FlowsUpdate")
	}

	if len(obj.obj.Flows) != 0 {

		if setDefault {
			obj.Flows().clearHolderSlice()
			for _, item := range obj.obj.Flows {
				newObj := newFlow(obj.validator)
				newObj.self().obj = item
				obj.Flows().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Flows().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *flowsUpdate) checkUnique() {

}

func (obj *flowsUpdate) checkConstraint() {

}

func (obj *routeState) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// State is required
	if obj.obj.State.Number() == 0 {
		obj.errors = append(obj.errors, "State is required field on interface RouteState")
	}
}

func (obj *routeState) checkUnique() {

}

func (obj *routeState) checkConstraint() {

	namesCons := []string{
		"bgpV4RouteRange.Name", "bgpV6RouteRange.Name", "isisV4RouteRange.Name", "isisV6RouteRange.Name",
	}

	for _, v := range obj.Names() {
		if !obj.validateConstraint(namesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid bgpV4RouteRange.Name|bgpV6RouteRange.Name|isisV4RouteRange.Name|isisV6RouteRange.Name type", v),
			)
		}
	}

}

func (obj *pingRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Endpoints) != 0 {

		if setDefault {
			obj.Endpoints().clearHolderSlice()
			for _, item := range obj.obj.Endpoints {
				newObj := newPing(obj.validator)
				newObj.self().obj = item
				obj.Endpoints().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Endpoints().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *pingRequest) checkUnique() {

}

func (obj *pingRequest) checkConstraint() {

}

func (obj *protocolState) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State.Number() == 0 {
		obj.errors = append(obj.errors, "State is required field on interface ProtocolState")
	}
}

func (obj *protocolState) checkUnique() {

}

func (obj *protocolState) checkConstraint() {

}

func (obj *deviceState) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.LacpMemberState != nil {

		obj.LacpMemberState().validateObj(setDefault)
	}

}

func (obj *deviceState) checkUnique() {

}

func (obj *deviceState) checkConstraint() {

}

func (obj *metricsRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Port != nil {

		obj.Port().validateObj(setDefault)
	}

	if obj.obj.Flow != nil {

		obj.Flow().validateObj(setDefault)
	}

	if obj.obj.Bgpv4 != nil {

		obj.Bgpv4().validateObj(setDefault)
	}

	if obj.obj.Bgpv6 != nil {

		obj.Bgpv6().validateObj(setDefault)
	}

	if obj.obj.Isis != nil {

		obj.Isis().validateObj(setDefault)
	}

	if obj.obj.Lag != nil {

		obj.Lag().validateObj(setDefault)
	}

	if obj.obj.LacpLagMember != nil {

		obj.LacpLagMember().validateObj(setDefault)
	}

}

func (obj *metricsRequest) checkUnique() {

}

func (obj *metricsRequest) checkConstraint() {

}

func (obj *statesRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Ipv4Neighbors != nil {

		obj.Ipv4Neighbors().validateObj(setDefault)
	}

	if obj.obj.Ipv6Neighbors != nil {

		obj.Ipv6Neighbors().validateObj(setDefault)
	}

	if obj.obj.BgpPrefixes != nil {

		obj.BgpPrefixes().validateObj(setDefault)
	}

	if obj.obj.IsisLsps != nil {

		obj.IsisLsps().validateObj(setDefault)
	}

}

func (obj *statesRequest) checkUnique() {

}

func (obj *statesRequest) checkConstraint() {

}

func (obj *captureRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// PortName is required
	if obj.obj.PortName == "" {
		obj.errors = append(obj.errors, "PortName is required field on interface CaptureRequest")
	}
}

func (obj *captureRequest) checkUnique() {

}

func (obj *captureRequest) checkConstraint() {

	portnameCons := []string{
		"port.Name",
	}

	if !obj.validateConstraint(portnameCons, obj.PortName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid port.Name type", obj.PortName()),
		)
	}

}

func (obj *setConfigResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *setConfigResponse) checkUnique() {

}

func (obj *setConfigResponse) checkConstraint() {

}

func (obj *getConfigResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *getConfigResponse) checkUnique() {

}

func (obj *getConfigResponse) checkConstraint() {

}

func (obj *setTransmitStateResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *setTransmitStateResponse) checkUnique() {

}

func (obj *setTransmitStateResponse) checkConstraint() {

}

func (obj *setLinkStateResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *setLinkStateResponse) checkUnique() {

}

func (obj *setLinkStateResponse) checkConstraint() {

}

func (obj *setCaptureStateResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *setCaptureStateResponse) checkUnique() {

}

func (obj *setCaptureStateResponse) checkConstraint() {

}

func (obj *updateFlowsResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *updateFlowsResponse) checkUnique() {

}

func (obj *updateFlowsResponse) checkConstraint() {

}

func (obj *setRouteStateResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *setRouteStateResponse) checkUnique() {

}

func (obj *setRouteStateResponse) checkConstraint() {

}

func (obj *sendPingResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *sendPingResponse) checkUnique() {

}

func (obj *sendPingResponse) checkConstraint() {

}

func (obj *setProtocolStateResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *setProtocolStateResponse) checkUnique() {

}

func (obj *setProtocolStateResponse) checkConstraint() {

}

func (obj *setDeviceStateResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *setDeviceStateResponse) checkUnique() {

}

func (obj *setDeviceStateResponse) checkConstraint() {

}

func (obj *getMetricsResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *getMetricsResponse) checkUnique() {

}

func (obj *getMetricsResponse) checkConstraint() {

}

func (obj *getStatesResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *getStatesResponse) checkUnique() {

}

func (obj *getStatesResponse) checkConstraint() {

}

func (obj *getCaptureResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *getCaptureResponse) checkUnique() {

}

func (obj *getCaptureResponse) checkConstraint() {

}

func (obj *port) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface Port")
	}
}

func (obj *port) checkUnique() {
	if !obj.isUnique("port", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *port) checkConstraint() {

}

func (obj *lag) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	if len(obj.obj.Ports) != 0 {

		if setDefault {
			obj.Ports().clearHolderSlice()
			for _, item := range obj.obj.Ports {
				newObj := newLagPort(obj.validator)
				newObj.self().obj = item
				obj.Ports().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ports().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.Protocol != nil {

		obj.Protocol().validateObj(setDefault)
	}

	if obj.obj.MinLinks != nil {
		if *obj.obj.MinLinks < 0 || *obj.obj.MinLinks > 32 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= Lag.MinLinks <= 32 but Got %d", *obj.obj.MinLinks))
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface Lag")
	}
}

func (obj *lag) checkUnique() {
	if !obj.isUnique("lag", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *lag) checkConstraint() {

}

func (obj *layer1) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// PortNames is required
	if obj.obj.PortNames == nil {
		obj.errors = append(obj.errors, "PortNames is required field on interface Layer1")
	}

	if obj.obj.Mtu != nil {
		if *obj.obj.Mtu < 64 || *obj.obj.Mtu > 9000 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("64 <= Layer1.Mtu <= 9000 but Got %d", *obj.obj.Mtu))
		}

	}

	if obj.obj.AutoNegotiation != nil {

		obj.AutoNegotiation().validateObj(setDefault)
	}

	if obj.obj.FlowControl != nil {

		obj.FlowControl().validateObj(setDefault)
	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface Layer1")
	}
}

func (obj *layer1) checkUnique() {
	if !obj.isUnique("layer1", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *layer1) checkConstraint() {

	portnamesCons := []string{
		"port.Name",
	}

	for _, v := range obj.PortNames() {
		if !obj.validateConstraint(portnamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid port.Name type", v),
			)
		}
	}

}

func (obj *capture) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// PortNames is required
	if obj.obj.PortNames == nil {
		obj.errors = append(obj.errors, "PortNames is required field on interface Capture")
	}

	if len(obj.obj.Filters) != 0 {

		if setDefault {
			obj.Filters().clearHolderSlice()
			for _, item := range obj.obj.Filters {
				newObj := newCaptureFilter(obj.validator)
				newObj.self().obj = item
				obj.Filters().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Filters().Items() {
			item.validateObj(setDefault)
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface Capture")
	}
}

func (obj *capture) checkUnique() {
	if !obj.isUnique("capture", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *capture) checkConstraint() {

	portnamesCons := []string{
		"port.Name",
	}

	for _, v := range obj.PortNames() {
		if !obj.validateConstraint(portnamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid port.Name type", v),
			)
		}
	}

}

func (obj *device) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	if len(obj.obj.Ethernets) != 0 {

		if setDefault {
			obj.Ethernets().clearHolderSlice()
			for _, item := range obj.obj.Ethernets {
				newObj := newDeviceEthernet(obj.validator)
				newObj.self().obj = item
				obj.Ethernets().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ethernets().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Ipv4Loopbacks) != 0 {

		if setDefault {
			obj.Ipv4Loopbacks().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Loopbacks {
				newObj := newDeviceIpv4Loopback(obj.validator)
				newObj.self().obj = item
				obj.Ipv4Loopbacks().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ipv4Loopbacks().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Ipv6Loopbacks) != 0 {

		if setDefault {
			obj.Ipv6Loopbacks().clearHolderSlice()
			for _, item := range obj.obj.Ipv6Loopbacks {
				newObj := newDeviceIpv6Loopback(obj.validator)
				newObj.self().obj = item
				obj.Ipv6Loopbacks().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ipv6Loopbacks().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.Isis != nil {

		obj.Isis().validateObj(setDefault)
	}

	if obj.obj.Bgp != nil {

		obj.Bgp().validateObj(setDefault)
	}

	if obj.obj.Vxlan != nil {

		obj.Vxlan().validateObj(setDefault)
	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface Device")
	}
}

func (obj *device) checkUnique() {
	if !obj.isUnique("device", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *device) checkConstraint() {

}

func (obj *flow) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	// TxRx is required
	if obj.obj.TxRx == nil {
		obj.errors = append(obj.errors, "TxRx is required field on interface Flow")
	}

	if obj.obj.TxRx != nil {

		obj.TxRx().validateObj(setDefault)
	}

	if len(obj.obj.Packet) != 0 {

		if setDefault {
			obj.Packet().clearHolderSlice()
			for _, item := range obj.obj.Packet {
				newObj := newFlowHeader(obj.validator)
				newObj.self().obj = item
				obj.Packet().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Packet().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.Size != nil {

		obj.Size().validateObj(setDefault)
	}

	if obj.obj.Rate != nil {

		obj.Rate().validateObj(setDefault)
	}

	if obj.obj.Duration != nil {

		obj.Duration().validateObj(setDefault)
	}

	if obj.obj.Metrics != nil {

		obj.Metrics().validateObj(setDefault)
	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface Flow")
	}
}

func (obj *flow) checkUnique() {
	if !obj.isUnique("flow", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *flow) checkConstraint() {

}

func (obj *event) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Link != nil {

		obj.Link().validateObj(setDefault)
	}

	if obj.obj.RxRateThreshold != nil {

		obj.RxRateThreshold().validateObj(setDefault)
	}

	if obj.obj.RouteAdvertiseWithdraw != nil {

		obj.RouteAdvertiseWithdraw().validateObj(setDefault)
	}

}

func (obj *event) checkUnique() {

}

func (obj *event) checkConstraint() {

}

func (obj *configOptions) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.PortOptions != nil {

		obj.PortOptions().validateObj(setDefault)
	}

}

func (obj *configOptions) checkUnique() {

}

func (obj *configOptions) checkConstraint() {

}

func (obj *ping) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Ipv4 != nil {

		obj.Ipv4().validateObj(setDefault)
	}

	if obj.obj.Ipv6 != nil {

		obj.Ipv6().validateObj(setDefault)
	}

}

func (obj *ping) checkUnique() {

}

func (obj *ping) checkConstraint() {

}

func (obj *lacpMemberState) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// State is required
	if obj.obj.State.Number() == 0 {
		obj.errors = append(obj.errors, "State is required field on interface LacpMemberState")
	}
}

func (obj *lacpMemberState) checkUnique() {

}

func (obj *lacpMemberState) checkConstraint() {

	lagmemberportnamesCons := []string{
		"port.Name",
	}

	for _, v := range obj.LagMemberPortNames() {
		if !obj.validateConstraint(lagmemberportnamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid port.Name type", v),
			)
		}
	}

}

func (obj *portMetricsRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

}

func (obj *portMetricsRequest) checkUnique() {

}

func (obj *portMetricsRequest) checkConstraint() {

	portnamesCons := []string{
		"port.Name",
	}

	for _, v := range obj.PortNames() {
		if !obj.validateConstraint(portnamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid port.Name type", v),
			)
		}
	}

}

func (obj *flowMetricsRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	if obj.obj.MetricGroups != nil {

		obj.MetricGroups().validateObj(setDefault)
	}

}

func (obj *flowMetricsRequest) checkUnique() {

}

func (obj *flowMetricsRequest) checkConstraint() {

	flownamesCons := []string{
		"flow.Name",
	}

	for _, v := range obj.FlowNames() {
		if !obj.validateConstraint(flownamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid flow.Name type", v),
			)
		}
	}

}

func (obj *bgpv4MetricsRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

}

func (obj *bgpv4MetricsRequest) checkUnique() {

}

func (obj *bgpv4MetricsRequest) checkConstraint() {

	peernamesCons := []string{
		"bgpV4Peer.Name",
	}

	for _, v := range obj.PeerNames() {
		if !obj.validateConstraint(peernamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid bgpV4Peer.Name type", v),
			)
		}
	}

}

func (obj *bgpv6MetricsRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

}

func (obj *bgpv6MetricsRequest) checkUnique() {

}

func (obj *bgpv6MetricsRequest) checkConstraint() {

	peernamesCons := []string{
		"bgpV6Peer.Name",
	}

	for _, v := range obj.PeerNames() {
		if !obj.validateConstraint(peernamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid bgpV6Peer.Name type", v),
			)
		}
	}

}

func (obj *isisMetricsRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

}

func (obj *isisMetricsRequest) checkUnique() {

}

func (obj *isisMetricsRequest) checkConstraint() {

	routernamesCons := []string{
		"deviceIsisRouter.Name",
	}

	for _, v := range obj.RouterNames() {
		if !obj.validateConstraint(routernamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid deviceIsisRouter.Name type", v),
			)
		}
	}

}

func (obj *lagMetricsRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

}

func (obj *lagMetricsRequest) checkUnique() {

}

func (obj *lagMetricsRequest) checkConstraint() {

	lagnamesCons := []string{
		"lag.Name",
	}

	for _, v := range obj.LagNames() {
		if !obj.validateConstraint(lagnamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid lag.Name type", v),
			)
		}
	}

}

func (obj *lacpLagMemberMetricsRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

}

func (obj *lacpLagMemberMetricsRequest) checkUnique() {

}

func (obj *lacpLagMemberMetricsRequest) checkConstraint() {

	lagnamesCons := []string{
		"lag.Name",
	}

	for _, v := range obj.LagNames() {
		if !obj.validateConstraint(lagnamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid lag.Name type", v),
			)
		}
	}

	lagmemberportnamesCons := []string{
		"port.Name",
	}

	for _, v := range obj.LagMemberPortNames() {
		if !obj.validateConstraint(lagmemberportnamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid port.Name type", v),
			)
		}
	}

}

func (obj *neighborsv4StatesRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

}

func (obj *neighborsv4StatesRequest) checkUnique() {

}

func (obj *neighborsv4StatesRequest) checkConstraint() {

	ethernetnamesCons := []string{
		"deviceEthernet.Name",
	}

	for _, v := range obj.EthernetNames() {
		if !obj.validateConstraint(ethernetnamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid deviceEthernet.Name type", v),
			)
		}
	}

}

func (obj *neighborsv6StatesRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

}

func (obj *neighborsv6StatesRequest) checkUnique() {

}

func (obj *neighborsv6StatesRequest) checkConstraint() {

	ethernetnamesCons := []string{
		"deviceEthernet.Name",
	}

	for _, v := range obj.EthernetNames() {
		if !obj.validateConstraint(ethernetnamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid deviceEthernet.Name type", v),
			)
		}
	}

}

func (obj *bgpPrefixStateRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	if len(obj.obj.Ipv4UnicastFilters) != 0 {

		if setDefault {
			obj.Ipv4UnicastFilters().clearHolderSlice()
			for _, item := range obj.obj.Ipv4UnicastFilters {
				newObj := newBgpPrefixIpv4UnicastFilter(obj.validator)
				newObj.self().obj = item
				obj.Ipv4UnicastFilters().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ipv4UnicastFilters().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Ipv6UnicastFilters) != 0 {

		if setDefault {
			obj.Ipv6UnicastFilters().clearHolderSlice()
			for _, item := range obj.obj.Ipv6UnicastFilters {
				newObj := newBgpPrefixIpv6UnicastFilter(obj.validator)
				newObj.self().obj = item
				obj.Ipv6UnicastFilters().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ipv6UnicastFilters().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *bgpPrefixStateRequest) checkUnique() {

}

func (obj *bgpPrefixStateRequest) checkConstraint() {

	bgppeernamesCons := []string{
		"bgpV4Peer.Name", "bgpV6Peer.Name",
	}

	for _, v := range obj.BgpPeerNames() {
		if !obj.validateConstraint(bgppeernamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid bgpV4Peer.Name|bgpV6Peer.Name type", v),
			)
		}
	}

}

func (obj *isisLspsStateRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

}

func (obj *isisLspsStateRequest) checkUnique() {

}

func (obj *isisLspsStateRequest) checkConstraint() {

	isisrouternamesCons := []string{
		"deviceIsisRouter.Name",
	}

	for _, v := range obj.IsisRouterNames() {
		if !obj.validateConstraint(isisrouternamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid deviceIsisRouter.Name type", v),
			)
		}
	}

}

func (obj *responseWarning) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *responseWarning) checkUnique() {

}

func (obj *responseWarning) checkConstraint() {

}

func (obj *responseError) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *responseError) checkUnique() {

}

func (obj *responseError) checkConstraint() {

}

func (obj *pingResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Responses) != 0 {

		if setDefault {
			obj.Responses().clearHolderSlice()
			for _, item := range obj.obj.Responses {
				newObj := newResponse(obj.validator)
				newObj.self().obj = item
				obj.Responses().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Responses().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *pingResponse) checkUnique() {

}

func (obj *pingResponse) checkConstraint() {

}

func (obj *metricsResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.PortMetrics) != 0 {

		if setDefault {
			obj.PortMetrics().clearHolderSlice()
			for _, item := range obj.obj.PortMetrics {
				newObj := newPortMetric(obj.validator)
				newObj.self().obj = item
				obj.PortMetrics().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.PortMetrics().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.FlowMetrics) != 0 {

		if setDefault {
			obj.FlowMetrics().clearHolderSlice()
			for _, item := range obj.obj.FlowMetrics {
				newObj := newFlowMetric(obj.validator)
				newObj.self().obj = item
				obj.FlowMetrics().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.FlowMetrics().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Bgpv4Metrics) != 0 {

		if setDefault {
			obj.Bgpv4Metrics().clearHolderSlice()
			for _, item := range obj.obj.Bgpv4Metrics {
				newObj := newBgpv4Metric(obj.validator)
				newObj.self().obj = item
				obj.Bgpv4Metrics().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Bgpv4Metrics().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Bgpv6Metrics) != 0 {

		if setDefault {
			obj.Bgpv6Metrics().clearHolderSlice()
			for _, item := range obj.obj.Bgpv6Metrics {
				newObj := newBgpv6Metric(obj.validator)
				newObj.self().obj = item
				obj.Bgpv6Metrics().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Bgpv6Metrics().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.IsisMetrics) != 0 {

		if setDefault {
			obj.IsisMetrics().clearHolderSlice()
			for _, item := range obj.obj.IsisMetrics {
				newObj := newIsisMetric(obj.validator)
				newObj.self().obj = item
				obj.IsisMetrics().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.IsisMetrics().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.LagMetrics) != 0 {

		if setDefault {
			obj.LagMetrics().clearHolderSlice()
			for _, item := range obj.obj.LagMetrics {
				newObj := newLagMetric(obj.validator)
				newObj.self().obj = item
				obj.LagMetrics().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.LagMetrics().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.LacpLagMemberMetrics) != 0 {

		if setDefault {
			obj.LacpLagMemberMetrics().clearHolderSlice()
			for _, item := range obj.obj.LacpLagMemberMetrics {
				newObj := newLacpLagMemberMetric(obj.validator)
				newObj.self().obj = item
				obj.LacpLagMemberMetrics().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.LacpLagMemberMetrics().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *metricsResponse) checkUnique() {

}

func (obj *metricsResponse) checkConstraint() {

}

func (obj *statesResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Ipv4Neighbors) != 0 {

		if setDefault {
			obj.Ipv4Neighbors().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Neighbors {
				newObj := newNeighborsv4State(obj.validator)
				newObj.self().obj = item
				obj.Ipv4Neighbors().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ipv4Neighbors().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Ipv6Neighbors) != 0 {

		if setDefault {
			obj.Ipv6Neighbors().clearHolderSlice()
			for _, item := range obj.obj.Ipv6Neighbors {
				newObj := newNeighborsv6State(obj.validator)
				newObj.self().obj = item
				obj.Ipv6Neighbors().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ipv6Neighbors().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.BgpPrefixes) != 0 {

		if setDefault {
			obj.BgpPrefixes().clearHolderSlice()
			for _, item := range obj.obj.BgpPrefixes {
				newObj := newBgpPrefixesState(obj.validator)
				newObj.self().obj = item
				obj.BgpPrefixes().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.BgpPrefixes().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.IsisLsps) != 0 {

		if setDefault {
			obj.IsisLsps().clearHolderSlice()
			for _, item := range obj.obj.IsisLsps {
				newObj := newIsisLspsState(obj.validator)
				newObj.self().obj = item
				obj.IsisLsps().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.IsisLsps().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *statesResponse) checkUnique() {

}

func (obj *statesResponse) checkConstraint() {

}

func (obj *lagPort) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// PortName is required
	if obj.obj.PortName == "" {
		obj.errors = append(obj.errors, "PortName is required field on interface LagPort")
	}

	if obj.obj.Lacp != nil {

		obj.Lacp().validateObj(setDefault)
	}

	// Ethernet is required
	if obj.obj.Ethernet == nil {
		obj.errors = append(obj.errors, "Ethernet is required field on interface LagPort")
	}

	if obj.obj.Ethernet != nil {

		obj.Ethernet().validateObj(setDefault)
	}

}

func (obj *lagPort) checkUnique() {

}

func (obj *lagPort) checkConstraint() {

	portnameCons := []string{
		"port.Name",
	}

	if !obj.validateConstraint(portnameCons, obj.PortName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid port.Name type", obj.PortName()),
		)
	}

}

func (obj *lagProtocol) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Lacp != nil {

		obj.Lacp().validateObj(setDefault)
	}

	if obj.obj.Static != nil {

		obj.Static().validateObj(setDefault)
	}

}

func (obj *lagProtocol) checkUnique() {

}

func (obj *lagProtocol) checkConstraint() {

}

func (obj *layer1AutoNegotiation) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *layer1AutoNegotiation) checkUnique() {

}

func (obj *layer1AutoNegotiation) checkConstraint() {

}

func (obj *layer1FlowControl) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.DirectedAddress != nil {

		err := obj.validateMac(obj.DirectedAddress())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on Layer1FlowControl.DirectedAddress"))
		}

	}

	if obj.obj.Ieee_802_1Qbb != nil {

		obj.Ieee8021Qbb().validateObj(setDefault)
	}

	if obj.obj.Ieee_802_3X != nil {

		obj.Ieee8023X().validateObj(setDefault)
	}

}

func (obj *layer1FlowControl) checkUnique() {

}

func (obj *layer1FlowControl) checkConstraint() {

}

func (obj *captureFilter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		obj.Custom().validateObj(setDefault)
	}

	if obj.obj.Ethernet != nil {

		obj.Ethernet().validateObj(setDefault)
	}

	if obj.obj.Vlan != nil {

		obj.Vlan().validateObj(setDefault)
	}

	if obj.obj.Ipv4 != nil {

		obj.Ipv4().validateObj(setDefault)
	}

	if obj.obj.Ipv6 != nil {

		obj.Ipv6().validateObj(setDefault)
	}

}

func (obj *captureFilter) checkUnique() {

}

func (obj *captureFilter) checkConstraint() {

}

func (obj *deviceEthernet) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// PortName is deprecated
	if obj.obj.PortName != nil {
		obj.deprecated(`PortName: property port_name is deprecated from the snappi version x.x.x.
please use connection.port_name instead of port_name`)
	}

	if obj.obj.Connection != nil {

		obj.Connection().validateObj(setDefault)
	}

	if len(obj.obj.Ipv4Addresses) != 0 {

		if setDefault {
			obj.Ipv4Addresses().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Addresses {
				newObj := newDeviceIpv4(obj.validator)
				newObj.self().obj = item
				obj.Ipv4Addresses().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ipv4Addresses().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Ipv6Addresses) != 0 {

		if setDefault {
			obj.Ipv6Addresses().clearHolderSlice()
			for _, item := range obj.obj.Ipv6Addresses {
				newObj := newDeviceIpv6(obj.validator)
				newObj.self().obj = item
				obj.Ipv6Addresses().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ipv6Addresses().Items() {
			item.validateObj(setDefault)
		}

	}

	// Mac is required
	if obj.obj.Mac == "" {
		obj.errors = append(obj.errors, "Mac is required field on interface DeviceEthernet")
	}
	if obj.obj.Mac != "" {

		err := obj.validateMac(obj.Mac())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on DeviceEthernet.Mac"))
		}

	}

	if obj.obj.Mtu != nil {
		if *obj.obj.Mtu < 0 || *obj.obj.Mtu > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= DeviceEthernet.Mtu <= 65535 but Got %d", *obj.obj.Mtu))
		}

	}

	if len(obj.obj.Vlans) != 0 {

		if setDefault {
			obj.Vlans().clearHolderSlice()
			for _, item := range obj.obj.Vlans {
				newObj := newDeviceVlan(obj.validator)
				newObj.self().obj = item
				obj.Vlans().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Vlans().Items() {
			item.validateObj(setDefault)
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface DeviceEthernet")
	}
}

func (obj *deviceEthernet) checkUnique() {
	if !obj.isUnique("deviceEthernet", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *deviceEthernet) checkConstraint() {

	portnameCons := []string{
		"port.Name", "lag.Name",
	}

	if !obj.validateConstraint(portnameCons, obj.PortName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid port.Name|lag.Name type", obj.PortName()),
		)
	}

}

func (obj *deviceIpv4Loopback) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// EthName is required
	if obj.obj.EthName == "" {
		obj.errors = append(obj.errors, "EthName is required field on interface DeviceIpv4Loopback")
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv4(obj.Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv4Loopback.Address"))
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface DeviceIpv4Loopback")
	}
}

func (obj *deviceIpv4Loopback) checkUnique() {
	if !obj.isUnique("deviceIpv4Loopback", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *deviceIpv4Loopback) checkConstraint() {

	ethnameCons := []string{
		"deviceEthernet.Name",
	}

	if !obj.validateConstraint(ethnameCons, obj.EthName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid deviceEthernet.Name type", obj.EthName()),
		)
	}

}

func (obj *deviceIpv6Loopback) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// EthName is required
	if obj.obj.EthName == "" {
		obj.errors = append(obj.errors, "EthName is required field on interface DeviceIpv6Loopback")
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv6(obj.Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv6Loopback.Address"))
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface DeviceIpv6Loopback")
	}
}

func (obj *deviceIpv6Loopback) checkUnique() {
	if !obj.isUnique("deviceIpv6Loopback", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *deviceIpv6Loopback) checkConstraint() {

	ethnameCons := []string{
		"deviceEthernet.Name",
	}

	if !obj.validateConstraint(ethnameCons, obj.EthName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid deviceEthernet.Name type", obj.EthName()),
		)
	}

}

func (obj *deviceIsisRouter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	if obj.obj.Instance != nil {

		obj.Instance().validateObj(setDefault)
	}

	// SystemId is required
	if obj.obj.SystemId == "" {
		obj.errors = append(obj.errors, "SystemId is required field on interface DeviceIsisRouter")
	}
	if obj.obj.SystemId != "" {

		err := obj.validateHex(obj.SystemId())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIsisRouter.SystemId"))
		}

	}

	if len(obj.obj.Interfaces) != 0 {

		if setDefault {
			obj.Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Interfaces {
				newObj := newIsisInterface(obj.validator)
				newObj.self().obj = item
				obj.Interfaces().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Interfaces().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.Basic != nil {

		obj.Basic().validateObj(setDefault)
	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(setDefault)
	}

	if obj.obj.RouterAuth != nil {

		obj.RouterAuth().validateObj(setDefault)
	}

	if len(obj.obj.V4Routes) != 0 {

		if setDefault {
			obj.V4Routes().clearHolderSlice()
			for _, item := range obj.obj.V4Routes {
				newObj := newIsisV4RouteRange(obj.validator)
				newObj.self().obj = item
				obj.V4Routes().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.V4Routes().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.V6Routes) != 0 {

		if setDefault {
			obj.V6Routes().clearHolderSlice()
			for _, item := range obj.obj.V6Routes {
				newObj := newIsisV6RouteRange(obj.validator)
				newObj.self().obj = item
				obj.V6Routes().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.V6Routes().Items() {
			item.validateObj(setDefault)
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface DeviceIsisRouter")
	}
}

func (obj *deviceIsisRouter) checkUnique() {
	if !obj.isUnique("deviceIsisRouter", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *deviceIsisRouter) checkConstraint() {

}

func (obj *deviceBgpRouter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// RouterId is required
	if obj.obj.RouterId == "" {
		obj.errors = append(obj.errors, "RouterId is required field on interface DeviceBgpRouter")
	}
	if obj.obj.RouterId != "" {

		err := obj.validateIpv4(obj.RouterId())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on DeviceBgpRouter.RouterId"))
		}

	}

	if len(obj.obj.Ipv4Interfaces) != 0 {

		if setDefault {
			obj.Ipv4Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Interfaces {
				newObj := newBgpV4Interface(obj.validator)
				newObj.self().obj = item
				obj.Ipv4Interfaces().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ipv4Interfaces().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Ipv6Interfaces) != 0 {

		if setDefault {
			obj.Ipv6Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Ipv6Interfaces {
				newObj := newBgpV6Interface(obj.validator)
				newObj.self().obj = item
				obj.Ipv6Interfaces().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ipv6Interfaces().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *deviceBgpRouter) checkUnique() {

}

func (obj *deviceBgpRouter) checkConstraint() {

}

func (obj *deviceVxlan) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.V4Tunnels) != 0 {

		if setDefault {
			obj.V4Tunnels().clearHolderSlice()
			for _, item := range obj.obj.V4Tunnels {
				newObj := newVxlanV4Tunnel(obj.validator)
				newObj.self().obj = item
				obj.V4Tunnels().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.V4Tunnels().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.V6Tunnels) != 0 {

		if setDefault {
			obj.V6Tunnels().clearHolderSlice()
			for _, item := range obj.obj.V6Tunnels {
				newObj := newVxlanV6Tunnel(obj.validator)
				newObj.self().obj = item
				obj.V6Tunnels().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.V6Tunnels().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *deviceVxlan) checkUnique() {

}

func (obj *deviceVxlan) checkConstraint() {

}

func (obj *flowTxRx) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Port != nil {

		obj.Port().validateObj(setDefault)
	}

	if obj.obj.Device != nil {

		obj.Device().validateObj(setDefault)
	}

}

func (obj *flowTxRx) checkUnique() {

}

func (obj *flowTxRx) checkConstraint() {

}

func (obj *flowHeader) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		obj.Custom().validateObj(setDefault)
	}

	if obj.obj.Ethernet != nil {

		obj.Ethernet().validateObj(setDefault)
	}

	if obj.obj.Vlan != nil {

		obj.Vlan().validateObj(setDefault)
	}

	if obj.obj.Vxlan != nil {

		obj.Vxlan().validateObj(setDefault)
	}

	if obj.obj.Ipv4 != nil {

		obj.Ipv4().validateObj(setDefault)
	}

	if obj.obj.Ipv6 != nil {

		obj.Ipv6().validateObj(setDefault)
	}

	if obj.obj.Pfcpause != nil {

		obj.Pfcpause().validateObj(setDefault)
	}

	if obj.obj.Ethernetpause != nil {

		obj.Ethernetpause().validateObj(setDefault)
	}

	if obj.obj.Tcp != nil {

		obj.Tcp().validateObj(setDefault)
	}

	if obj.obj.Udp != nil {

		obj.Udp().validateObj(setDefault)
	}

	if obj.obj.Gre != nil {

		obj.Gre().validateObj(setDefault)
	}

	if obj.obj.Gtpv1 != nil {

		obj.Gtpv1().validateObj(setDefault)
	}

	if obj.obj.Gtpv2 != nil {

		obj.Gtpv2().validateObj(setDefault)
	}

	if obj.obj.Arp != nil {

		obj.Arp().validateObj(setDefault)
	}

	if obj.obj.Icmp != nil {

		obj.Icmp().validateObj(setDefault)
	}

	if obj.obj.Icmpv6 != nil {

		obj.Icmpv6().validateObj(setDefault)
	}

	if obj.obj.Ppp != nil {

		obj.Ppp().validateObj(setDefault)
	}

	if obj.obj.Igmpv1 != nil {

		obj.Igmpv1().validateObj(setDefault)
	}

	if obj.obj.Mpls != nil {

		obj.Mpls().validateObj(setDefault)
	}

}

func (obj *flowHeader) checkUnique() {

}

func (obj *flowHeader) checkConstraint() {

}

func (obj *flowSize) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Random != nil {

		obj.Random().validateObj(setDefault)
	}

}

func (obj *flowSize) checkUnique() {

}

func (obj *flowSize) checkConstraint() {

}

func (obj *flowRate) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Pps != nil {
		if *obj.obj.Pps < 1 || *obj.obj.Pps > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= FlowRate.Pps <= 9223372036854775807 but Got %d", *obj.obj.Pps))
		}

	}

	if obj.obj.Bps != nil {
		if *obj.obj.Bps < 672 || *obj.obj.Bps > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("672 <= FlowRate.Bps <= 9223372036854775807 but Got %d", *obj.obj.Bps))
		}

	}

	if obj.obj.Kbps != nil {
		if *obj.obj.Kbps < 1 || *obj.obj.Kbps > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= FlowRate.Kbps <= 9223372036854775807 but Got %d", *obj.obj.Kbps))
		}

	}

	if obj.obj.Mbps != nil {
		if *obj.obj.Mbps < 1 || *obj.obj.Mbps > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= FlowRate.Mbps <= 9223372036854775807 but Got %d", *obj.obj.Mbps))
		}

	}

	if obj.obj.Gbps != nil {
		if *obj.obj.Gbps < 1 || *obj.obj.Gbps > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= FlowRate.Gbps <= 2147483647 but Got %d", *obj.obj.Gbps))
		}

	}

	if obj.obj.Percentage != nil {
		if *obj.obj.Percentage < 0 || *obj.obj.Percentage > 100 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowRate.Percentage <= 100 but Got %f", *obj.obj.Percentage))
		}

	}

}

func (obj *flowRate) checkUnique() {

}

func (obj *flowRate) checkConstraint() {

}

func (obj *flowDuration) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.FixedPackets != nil {

		obj.FixedPackets().validateObj(setDefault)
	}

	if obj.obj.FixedSeconds != nil {

		obj.FixedSeconds().validateObj(setDefault)
	}

	if obj.obj.Burst != nil {

		obj.Burst().validateObj(setDefault)
	}

	if obj.obj.Continuous != nil {

		obj.Continuous().validateObj(setDefault)
	}

}

func (obj *flowDuration) checkUnique() {

}

func (obj *flowDuration) checkConstraint() {

}

func (obj *flowMetrics) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Latency != nil {

		obj.Latency().validateObj(setDefault)
	}

}

func (obj *flowMetrics) checkUnique() {

}

func (obj *flowMetrics) checkConstraint() {

}

func (obj *eventLink) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *eventLink) checkUnique() {

}

func (obj *eventLink) checkConstraint() {

}

func (obj *eventRxRateThreshold) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Threshold != nil {
		if *obj.obj.Threshold < 0 || *obj.obj.Threshold > 100 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= EventRxRateThreshold.Threshold <= 100 but Got %f", *obj.obj.Threshold))
		}

	}

}

func (obj *eventRxRateThreshold) checkUnique() {

}

func (obj *eventRxRateThreshold) checkConstraint() {

}

func (obj *eventRouteAdvertiseWithdraw) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *eventRouteAdvertiseWithdraw) checkUnique() {

}

func (obj *eventRouteAdvertiseWithdraw) checkConstraint() {

}

func (obj *portOptions) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *portOptions) checkUnique() {

}

func (obj *portOptions) checkConstraint() {

}

func (obj *pingIpv4) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	if obj.obj.DstIp != nil {

		err := obj.validateIpv4(obj.DstIp())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PingIpv4.DstIp"))
		}

	}

}

func (obj *pingIpv4) checkUnique() {

}

func (obj *pingIpv4) checkConstraint() {

	srcnameCons := []string{
		"deviceIpv4.Name",
	}

	if !obj.validateConstraint(srcnameCons, obj.SrcName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid deviceIpv4.Name type", obj.SrcName()),
		)
	}

}

func (obj *pingIpv6) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	if obj.obj.DstIp != nil {

		err := obj.validateIpv6(obj.DstIp())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PingIpv6.DstIp"))
		}

	}

}

func (obj *pingIpv6) checkUnique() {

}

func (obj *pingIpv6) checkConstraint() {

	srcnameCons := []string{
		"deviceIpv6.Name",
	}

	if !obj.validateConstraint(srcnameCons, obj.SrcName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid deviceIpv6.Name type", obj.SrcName()),
		)
	}

}

func (obj *flowMetricGroupRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

}

func (obj *flowMetricGroupRequest) checkUnique() {

}

func (obj *flowMetricGroupRequest) checkConstraint() {

	ingressCons := []string{
		"flow.Packet//metricGroup",
	}

	for _, v := range obj.Ingress() {
		if !obj.validateConstraint(ingressCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid flow.Packet//metricGroup type", v),
			)
		}
	}

	egressCons := []string{
		"flow.Egress//metricGroup",
	}

	for _, v := range obj.Egress() {
		if !obj.validateConstraint(egressCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid flow.Egress//metricGroup type", v),
			)
		}
	}

}

func (obj *bgpPrefixIpv4UnicastFilter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Addresses != nil {

		err := obj.validateIpv4Slice(obj.Addresses())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv4UnicastFilter.Addresses"))
		}

	}

}

func (obj *bgpPrefixIpv4UnicastFilter) checkUnique() {

}

func (obj *bgpPrefixIpv4UnicastFilter) checkConstraint() {

}

func (obj *bgpPrefixIpv6UnicastFilter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Addresses != nil {

		err := obj.validateIpv6Slice(obj.Addresses())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv6UnicastFilter.Addresses"))
		}

	}

}

func (obj *bgpPrefixIpv6UnicastFilter) checkUnique() {

}

func (obj *bgpPrefixIpv6UnicastFilter) checkConstraint() {

}

func (obj *response) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

}

func (obj *response) checkUnique() {

}

func (obj *response) checkConstraint() {

	srcnameCons := []string{
		"deviceIpv4.Name", "deviceIpv6.Name",
	}

	if !obj.validateConstraint(srcnameCons, obj.SrcName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid deviceIpv4.Name|deviceIpv6.Name type", obj.SrcName()),
		)
	}

}

func (obj *portMetric) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	if obj.obj.FramesTx != nil {
		if *obj.obj.FramesTx < 0 || *obj.obj.FramesTx > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PortMetric.FramesTx <= 9223372036854775807 but Got %d", *obj.obj.FramesTx))
		}

	}

	if obj.obj.FramesRx != nil {
		if *obj.obj.FramesRx < 0 || *obj.obj.FramesRx > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PortMetric.FramesRx <= 9223372036854775807 but Got %d", *obj.obj.FramesRx))
		}

	}

	if obj.obj.BytesTx != nil {
		if *obj.obj.BytesTx < 0 || *obj.obj.BytesTx > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PortMetric.BytesTx <= 9223372036854775807 but Got %d", *obj.obj.BytesTx))
		}

	}

	if obj.obj.BytesRx != nil {
		if *obj.obj.BytesRx < 0 || *obj.obj.BytesRx > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PortMetric.BytesRx <= 9223372036854775807 but Got %d", *obj.obj.BytesRx))
		}

	}

}

func (obj *portMetric) checkUnique() {

}

func (obj *portMetric) checkConstraint() {

	nameCons := []string{
		"port.Name",
	}

	if !obj.validateConstraint(nameCons, obj.Name()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid port.Name type", obj.Name()),
		)
	}

}

func (obj *flowMetric) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.MetricGroups) != 0 {

		if setDefault {
			obj.MetricGroups().clearHolderSlice()
			for _, item := range obj.obj.MetricGroups {
				newObj := newFlowMetricGroup(obj.validator)
				newObj.self().obj = item
				obj.MetricGroups().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.MetricGroups().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.FramesTx != nil {
		if *obj.obj.FramesTx < 0 || *obj.obj.FramesTx > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowMetric.FramesTx <= 9223372036854775807 but Got %d", *obj.obj.FramesTx))
		}

	}

	if obj.obj.FramesRx != nil {
		if *obj.obj.FramesRx < 0 || *obj.obj.FramesRx > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowMetric.FramesRx <= 9223372036854775807 but Got %d", *obj.obj.FramesRx))
		}

	}

	if obj.obj.BytesTx != nil {
		if *obj.obj.BytesTx < 0 || *obj.obj.BytesTx > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowMetric.BytesTx <= 9223372036854775807 but Got %d", *obj.obj.BytesTx))
		}

	}

	if obj.obj.BytesRx != nil {
		if *obj.obj.BytesRx < 0 || *obj.obj.BytesRx > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowMetric.BytesRx <= 9223372036854775807 but Got %d", *obj.obj.BytesRx))
		}

	}

	if obj.obj.Timestamps != nil {

		obj.Timestamps().validateObj(setDefault)
	}

	if obj.obj.Latency != nil {

		obj.Latency().validateObj(setDefault)
	}

}

func (obj *flowMetric) checkUnique() {

}

func (obj *flowMetric) checkConstraint() {

}

func (obj *bgpv4Metric) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *bgpv4Metric) checkUnique() {

}

func (obj *bgpv4Metric) checkConstraint() {

}

func (obj *bgpv6Metric) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *bgpv6Metric) checkUnique() {

}

func (obj *bgpv6Metric) checkConstraint() {

}

func (obj *isisMetric) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *isisMetric) checkUnique() {

}

func (obj *isisMetric) checkConstraint() {

}

func (obj *lagMetric) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	if obj.obj.FramesTx != nil {
		if *obj.obj.FramesTx < 0 || *obj.obj.FramesTx > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LagMetric.FramesTx <= 9223372036854775807 but Got %d", *obj.obj.FramesTx))
		}

	}

	if obj.obj.FramesRx != nil {
		if *obj.obj.FramesRx < 0 || *obj.obj.FramesRx > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LagMetric.FramesRx <= 9223372036854775807 but Got %d", *obj.obj.FramesRx))
		}

	}

	if obj.obj.BytesTx != nil {
		if *obj.obj.BytesTx < 0 || *obj.obj.BytesTx > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LagMetric.BytesTx <= 9223372036854775807 but Got %d", *obj.obj.BytesTx))
		}

	}

	if obj.obj.BytesRx != nil {
		if *obj.obj.BytesRx < 0 || *obj.obj.BytesRx > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LagMetric.BytesRx <= 9223372036854775807 but Got %d", *obj.obj.BytesRx))
		}

	}

}

func (obj *lagMetric) checkUnique() {

}

func (obj *lagMetric) checkConstraint() {

	nameCons := []string{
		"lag.Name",
	}

	if !obj.validateConstraint(nameCons, obj.Name()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid lag.Name type", obj.Name()),
		)
	}

}

func (obj *lacpLagMemberMetric) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.LacpInPkts != nil {
		if *obj.obj.LacpInPkts < 0 || *obj.obj.LacpInPkts > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LacpLagMemberMetric.LacpInPkts <= 9223372036854775807 but Got %d", *obj.obj.LacpInPkts))
		}

	}

	if obj.obj.LacpOutPkts != nil {
		if *obj.obj.LacpOutPkts < 0 || *obj.obj.LacpOutPkts > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LacpLagMemberMetric.LacpOutPkts <= 9223372036854775807 but Got %d", *obj.obj.LacpOutPkts))
		}

	}

	if obj.obj.LacpRxErrors != nil {
		if *obj.obj.LacpRxErrors < 0 || *obj.obj.LacpRxErrors > 9223372036854775807 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LacpLagMemberMetric.LacpRxErrors <= 9223372036854775807 but Got %d", *obj.obj.LacpRxErrors))
		}

	}

	if obj.obj.SystemId != nil {

		err := obj.validateMac(obj.SystemId())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on LacpLagMemberMetric.SystemId"))
		}

	}

	if obj.obj.PartnerId != nil {

		err := obj.validateMac(obj.PartnerId())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on LacpLagMemberMetric.PartnerId"))
		}

	}

}

func (obj *lacpLagMemberMetric) checkUnique() {

}

func (obj *lacpLagMemberMetric) checkConstraint() {

}

func (obj *neighborsv4State) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// EthernetName is required
	if obj.obj.EthernetName == "" {
		obj.errors = append(obj.errors, "EthernetName is required field on interface Neighborsv4State")
	}

	// Ipv4Address is required
	if obj.obj.Ipv4Address == "" {
		obj.errors = append(obj.errors, "Ipv4Address is required field on interface Neighborsv4State")
	}
	if obj.obj.Ipv4Address != "" {

		err := obj.validateIpv4(obj.Ipv4Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on Neighborsv4State.Ipv4Address"))
		}

	}

	if obj.obj.LinkLayerAddress != nil {

		err := obj.validateMac(obj.LinkLayerAddress())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on Neighborsv4State.LinkLayerAddress"))
		}

	}

}

func (obj *neighborsv4State) checkUnique() {

}

func (obj *neighborsv4State) checkConstraint() {

}

func (obj *neighborsv6State) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// EthernetName is required
	if obj.obj.EthernetName == "" {
		obj.errors = append(obj.errors, "EthernetName is required field on interface Neighborsv6State")
	}

	// Ipv6Address is required
	if obj.obj.Ipv6Address == "" {
		obj.errors = append(obj.errors, "Ipv6Address is required field on interface Neighborsv6State")
	}
	if obj.obj.Ipv6Address != "" {

		err := obj.validateIpv6(obj.Ipv6Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on Neighborsv6State.Ipv6Address"))
		}

	}

	if obj.obj.LinkLayerAddress != nil {

		err := obj.validateMac(obj.LinkLayerAddress())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on Neighborsv6State.LinkLayerAddress"))
		}

	}

}

func (obj *neighborsv6State) checkUnique() {

}

func (obj *neighborsv6State) checkConstraint() {

}

func (obj *bgpPrefixesState) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Ipv4UnicastPrefixes) != 0 {

		if setDefault {
			obj.Ipv4UnicastPrefixes().clearHolderSlice()
			for _, item := range obj.obj.Ipv4UnicastPrefixes {
				newObj := newBgpPrefixIpv4UnicastState(obj.validator)
				newObj.self().obj = item
				obj.Ipv4UnicastPrefixes().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ipv4UnicastPrefixes().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Ipv6UnicastPrefixes) != 0 {

		if setDefault {
			obj.Ipv6UnicastPrefixes().clearHolderSlice()
			for _, item := range obj.obj.Ipv6UnicastPrefixes {
				newObj := newBgpPrefixIpv6UnicastState(obj.validator)
				newObj.self().obj = item
				obj.Ipv6UnicastPrefixes().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ipv6UnicastPrefixes().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *bgpPrefixesState) checkUnique() {

}

func (obj *bgpPrefixesState) checkConstraint() {

}

func (obj *isisLspsState) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Lsps) != 0 {

		if setDefault {
			obj.Lsps().clearHolderSlice()
			for _, item := range obj.obj.Lsps {
				newObj := newIsisLspState(obj.validator)
				newObj.self().obj = item
				obj.Lsps().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Lsps().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *isisLspsState) checkUnique() {

}

func (obj *isisLspsState) checkConstraint() {

}

func (obj *lagPortLacp) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.ActorPortNumber != nil {
		if *obj.obj.ActorPortNumber < 0 || *obj.obj.ActorPortNumber > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LagPortLacp.ActorPortNumber <= 65535 but Got %d", *obj.obj.ActorPortNumber))
		}

	}

	if obj.obj.ActorPortPriority != nil {
		if *obj.obj.ActorPortPriority < 0 || *obj.obj.ActorPortPriority > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LagPortLacp.ActorPortPriority <= 65535 but Got %d", *obj.obj.ActorPortPriority))
		}

	}

	if obj.obj.LacpduPeriodicTimeInterval != nil {
		if *obj.obj.LacpduPeriodicTimeInterval < 0 || *obj.obj.LacpduPeriodicTimeInterval > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LagPortLacp.LacpduPeriodicTimeInterval <= 65535 but Got %d", *obj.obj.LacpduPeriodicTimeInterval))
		}

	}

	if obj.obj.LacpduTimeout != nil {
		if *obj.obj.LacpduTimeout < 0 || *obj.obj.LacpduTimeout > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LagPortLacp.LacpduTimeout <= 65535 but Got %d", *obj.obj.LacpduTimeout))
		}

	}

}

func (obj *lagPortLacp) checkUnique() {

}

func (obj *lagPortLacp) checkConstraint() {

}

func (obj *deviceEthernetBase) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	// Mac is required
	if obj.obj.Mac == "" {
		obj.errors = append(obj.errors, "Mac is required field on interface DeviceEthernetBase")
	}
	if obj.obj.Mac != "" {

		err := obj.validateMac(obj.Mac())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on DeviceEthernetBase.Mac"))
		}

	}

	if obj.obj.Mtu != nil {
		if *obj.obj.Mtu < 0 || *obj.obj.Mtu > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= DeviceEthernetBase.Mtu <= 65535 but Got %d", *obj.obj.Mtu))
		}

	}

	if len(obj.obj.Vlans) != 0 {

		if setDefault {
			obj.Vlans().clearHolderSlice()
			for _, item := range obj.obj.Vlans {
				newObj := newDeviceVlan(obj.validator)
				newObj.self().obj = item
				obj.Vlans().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Vlans().Items() {
			item.validateObj(setDefault)
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface DeviceEthernetBase")
	}
}

func (obj *deviceEthernetBase) checkUnique() {
	if !obj.isUnique("deviceEthernetBase", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *deviceEthernetBase) checkConstraint() {

}

func (obj *lagProtocolLacp) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.ActorSystemId != nil {

		err := obj.validateMac(obj.ActorSystemId())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on LagProtocolLacp.ActorSystemId"))
		}

	}

	if obj.obj.ActorSystemPriority != nil {
		if *obj.obj.ActorSystemPriority < 0 || *obj.obj.ActorSystemPriority > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LagProtocolLacp.ActorSystemPriority <= 65535 but Got %d", *obj.obj.ActorSystemPriority))
		}

	}

	if obj.obj.ActorKey != nil {
		if *obj.obj.ActorKey < 0 || *obj.obj.ActorKey > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LagProtocolLacp.ActorKey <= 65535 but Got %d", *obj.obj.ActorKey))
		}

	}

}

func (obj *lagProtocolLacp) checkUnique() {

}

func (obj *lagProtocolLacp) checkConstraint() {

}

func (obj *lagProtocolStatic) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.LagId != nil {
		if *obj.obj.LagId < 0 || *obj.obj.LagId > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LagProtocolStatic.LagId <= 65535 but Got %d", *obj.obj.LagId))
		}

	}

}

func (obj *lagProtocolStatic) checkUnique() {

}

func (obj *lagProtocolStatic) checkConstraint() {

}

func (obj *layer1Ieee8021Qbb) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *layer1Ieee8021Qbb) checkUnique() {

}

func (obj *layer1Ieee8021Qbb) checkConstraint() {

}

func (obj *layer1Ieee8023X) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *layer1Ieee8023X) checkUnique() {

}

func (obj *layer1Ieee8023X) checkConstraint() {

}

func (obj *captureCustom) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateHex(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on CaptureCustom.Value"))
		}

	}

	if obj.obj.Mask != nil {

		err := obj.validateHex(obj.Mask())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on CaptureCustom.Mask"))
		}

	}

}

func (obj *captureCustom) checkUnique() {

}

func (obj *captureCustom) checkConstraint() {

}

func (obj *captureEthernet) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Src != nil {

		obj.Src().validateObj(setDefault)
	}

	if obj.obj.Dst != nil {

		obj.Dst().validateObj(setDefault)
	}

	if obj.obj.EtherType != nil {

		obj.EtherType().validateObj(setDefault)
	}

	if obj.obj.PfcQueue != nil {

		obj.PfcQueue().validateObj(setDefault)
	}

}

func (obj *captureEthernet) checkUnique() {

}

func (obj *captureEthernet) checkConstraint() {

}

func (obj *captureVlan) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Priority != nil {

		obj.Priority().validateObj(setDefault)
	}

	if obj.obj.Cfi != nil {

		obj.Cfi().validateObj(setDefault)
	}

	if obj.obj.Id != nil {

		obj.Id().validateObj(setDefault)
	}

	if obj.obj.Protocol != nil {

		obj.Protocol().validateObj(setDefault)
	}

}

func (obj *captureVlan) checkUnique() {

}

func (obj *captureVlan) checkConstraint() {

}

func (obj *captureIpv4) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(setDefault)
	}

	if obj.obj.HeaderLength != nil {

		obj.HeaderLength().validateObj(setDefault)
	}

	if obj.obj.Priority != nil {

		obj.Priority().validateObj(setDefault)
	}

	if obj.obj.TotalLength != nil {

		obj.TotalLength().validateObj(setDefault)
	}

	if obj.obj.Identification != nil {

		obj.Identification().validateObj(setDefault)
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(setDefault)
	}

	if obj.obj.DontFragment != nil {

		obj.DontFragment().validateObj(setDefault)
	}

	if obj.obj.MoreFragments != nil {

		obj.MoreFragments().validateObj(setDefault)
	}

	if obj.obj.FragmentOffset != nil {

		obj.FragmentOffset().validateObj(setDefault)
	}

	if obj.obj.TimeToLive != nil {

		obj.TimeToLive().validateObj(setDefault)
	}

	if obj.obj.Protocol != nil {

		obj.Protocol().validateObj(setDefault)
	}

	if obj.obj.HeaderChecksum != nil {

		obj.HeaderChecksum().validateObj(setDefault)
	}

	if obj.obj.Src != nil {

		obj.Src().validateObj(setDefault)
	}

	if obj.obj.Dst != nil {

		obj.Dst().validateObj(setDefault)
	}

}

func (obj *captureIpv4) checkUnique() {

}

func (obj *captureIpv4) checkConstraint() {

}

func (obj *captureIpv6) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(setDefault)
	}

	if obj.obj.TrafficClass != nil {

		obj.TrafficClass().validateObj(setDefault)
	}

	if obj.obj.FlowLabel != nil {

		obj.FlowLabel().validateObj(setDefault)
	}

	if obj.obj.PayloadLength != nil {

		obj.PayloadLength().validateObj(setDefault)
	}

	if obj.obj.NextHeader != nil {

		obj.NextHeader().validateObj(setDefault)
	}

	if obj.obj.HopLimit != nil {

		obj.HopLimit().validateObj(setDefault)
	}

	if obj.obj.Src != nil {

		obj.Src().validateObj(setDefault)
	}

	if obj.obj.Dst != nil {

		obj.Dst().validateObj(setDefault)
	}

}

func (obj *captureIpv6) checkUnique() {

}

func (obj *captureIpv6) checkConstraint() {

}

func (obj *ethernetConnection) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

}

func (obj *ethernetConnection) checkUnique() {

}

func (obj *ethernetConnection) checkConstraint() {

	portnameCons := []string{
		"port.Name",
	}

	if !obj.validateConstraint(portnameCons, obj.PortName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid port.Name type", obj.PortName()),
		)
	}

	lagnameCons := []string{
		"lag.Name",
	}

	if !obj.validateConstraint(lagnameCons, obj.LagName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid lag.Name type", obj.LagName()),
		)
	}

	vxlannameCons := []string{
		"vxlanV4Tunnel.Name", "vxlanV6Tunnel.Name",
	}

	if !obj.validateConstraint(vxlannameCons, obj.VxlanName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid vxlanV4Tunnel.Name|vxlanV6Tunnel.Name type", obj.VxlanName()),
		)
	}

}

func (obj *deviceIpv4) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	// Gateway is required
	if obj.obj.Gateway == "" {
		obj.errors = append(obj.errors, "Gateway is required field on interface DeviceIpv4")
	}
	if obj.obj.Gateway != "" {

		err := obj.validateIpv4(obj.Gateway())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv4.Gateway"))
		}

	}

	if obj.obj.GatewayMac != nil {

		obj.GatewayMac().validateObj(setDefault)
	}

	// Address is required
	if obj.obj.Address == "" {
		obj.errors = append(obj.errors, "Address is required field on interface DeviceIpv4")
	}
	if obj.obj.Address != "" {

		err := obj.validateIpv4(obj.Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv4.Address"))
		}

	}

	if obj.obj.Prefix != nil {
		if *obj.obj.Prefix < 1 || *obj.obj.Prefix > 32 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= DeviceIpv4.Prefix <= 32 but Got %d", *obj.obj.Prefix))
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface DeviceIpv4")
	}
}

func (obj *deviceIpv4) checkUnique() {
	if !obj.isUnique("deviceIpv4", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *deviceIpv4) checkConstraint() {

}

func (obj *deviceIpv6) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	// Gateway is required
	if obj.obj.Gateway == "" {
		obj.errors = append(obj.errors, "Gateway is required field on interface DeviceIpv6")
	}
	if obj.obj.Gateway != "" {

		err := obj.validateIpv6(obj.Gateway())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv6.Gateway"))
		}

	}

	if obj.obj.GatewayMac != nil {

		obj.GatewayMac().validateObj(setDefault)
	}

	// Address is required
	if obj.obj.Address == "" {
		obj.errors = append(obj.errors, "Address is required field on interface DeviceIpv6")
	}
	if obj.obj.Address != "" {

		err := obj.validateIpv6(obj.Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv6.Address"))
		}

	}

	if obj.obj.Prefix != nil {
		if *obj.obj.Prefix < 1 || *obj.obj.Prefix > 128 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= DeviceIpv6.Prefix <= 128 but Got %d", *obj.obj.Prefix))
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface DeviceIpv6")
	}
}

func (obj *deviceIpv6) checkUnique() {
	if !obj.isUnique("deviceIpv6", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *deviceIpv6) checkConstraint() {

}

func (obj *deviceVlan) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	if obj.obj.Priority != nil {
		if *obj.obj.Priority < 0 || *obj.obj.Priority > 3 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= DeviceVlan.Priority <= 3 but Got %d", *obj.obj.Priority))
		}

	}

	if obj.obj.Id != nil {
		if *obj.obj.Id < 0 || *obj.obj.Id > 4095 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= DeviceVlan.Id <= 4095 but Got %d", *obj.obj.Id))
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface DeviceVlan")
	}
}

func (obj *deviceVlan) checkUnique() {
	if !obj.isUnique("deviceVlan", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *deviceVlan) checkConstraint() {

}

func (obj *deviceIsisMultiInstance) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Iid != nil {
		if *obj.obj.Iid < 0 || *obj.obj.Iid > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= DeviceIsisMultiInstance.Iid <= 65535 but Got %d", *obj.obj.Iid))
		}

	}

}

func (obj *deviceIsisMultiInstance) checkUnique() {

}

func (obj *deviceIsisMultiInstance) checkConstraint() {

}

func (obj *isisInterface) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// EthName is required
	if obj.obj.EthName == "" {
		obj.errors = append(obj.errors, "EthName is required field on interface IsisInterface")
	}

	if obj.obj.L1Settings != nil {

		obj.L1Settings().validateObj(setDefault)
	}

	if obj.obj.L2Settings != nil {

		obj.L2Settings().validateObj(setDefault)
	}

	if len(obj.obj.MultiTopologyIds) != 0 {

		if setDefault {
			obj.MultiTopologyIds().clearHolderSlice()
			for _, item := range obj.obj.MultiTopologyIds {
				newObj := newIsisMT(obj.validator)
				newObj.self().obj = item
				obj.MultiTopologyIds().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.MultiTopologyIds().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.TrafficEngineering) != 0 {

		if setDefault {
			obj.TrafficEngineering().clearHolderSlice()
			for _, item := range obj.obj.TrafficEngineering {
				newObj := newLinkStateTE(obj.validator)
				newObj.self().obj = item
				obj.TrafficEngineering().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.TrafficEngineering().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.Authentication != nil {

		obj.Authentication().validateObj(setDefault)
	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(setDefault)
	}

	if obj.obj.LinkProtection != nil {

		obj.LinkProtection().validateObj(setDefault)
	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface IsisInterface")
	}
}

func (obj *isisInterface) checkUnique() {
	if !obj.isUnique("isisInterface", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *isisInterface) checkConstraint() {

	ethnameCons := []string{
		"deviceEthernet.Name",
	}

	if !obj.validateConstraint(ethnameCons, obj.EthName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid deviceEthernet.Name type", obj.EthName()),
		)
	}

}

func (obj *isisBasic) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Ipv4TeRouterId != nil {

		err := obj.validateIpv4(obj.Ipv4TeRouterId())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on IsisBasic.Ipv4TeRouterId"))
		}

	}

}

func (obj *isisBasic) checkUnique() {

}

func (obj *isisBasic) checkConstraint() {

}

func (obj *isisAdvanced) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.MaxAreaAddresses != nil {
		if *obj.obj.MaxAreaAddresses < 0 || *obj.obj.MaxAreaAddresses > 254 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= IsisAdvanced.MaxAreaAddresses <= 254 but Got %d", *obj.obj.MaxAreaAddresses))
		}

	}

	if obj.obj.AreaAddresses != nil {

		err := obj.validateHexSlice(obj.AreaAddresses())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on IsisAdvanced.AreaAddresses"))
		}

	}

	if obj.obj.LspRefreshRate != nil {
		if *obj.obj.LspRefreshRate < 1 || *obj.obj.LspRefreshRate > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= IsisAdvanced.LspRefreshRate <= 65535 but Got %d", *obj.obj.LspRefreshRate))
		}

	}

	if obj.obj.LspLifetime != nil {
		if *obj.obj.LspLifetime < 1 || *obj.obj.LspLifetime > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= IsisAdvanced.LspLifetime <= 65535 but Got %d", *obj.obj.LspLifetime))
		}

	}

	if obj.obj.PsnpInterval != nil {
		if *obj.obj.PsnpInterval < 1 || *obj.obj.PsnpInterval > 60000 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= IsisAdvanced.PsnpInterval <= 60000 but Got %d", *obj.obj.PsnpInterval))
		}

	}

	if obj.obj.CsnpInterval != nil {
		if *obj.obj.CsnpInterval < 1 || *obj.obj.CsnpInterval > 65535000 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= IsisAdvanced.CsnpInterval <= 65535000 but Got %d", *obj.obj.CsnpInterval))
		}

	}

	if obj.obj.MaxLspSize != nil {
		if *obj.obj.MaxLspSize < 64 || *obj.obj.MaxLspSize > 9216 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("64 <= IsisAdvanced.MaxLspSize <= 9216 but Got %d", *obj.obj.MaxLspSize))
		}

	}

	if obj.obj.LspMgroupMinTransInterval != nil {
		if *obj.obj.LspMgroupMinTransInterval < 1 || *obj.obj.LspMgroupMinTransInterval > 60000 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= IsisAdvanced.LspMgroupMinTransInterval <= 60000 but Got %d", *obj.obj.LspMgroupMinTransInterval))
		}

	}

}

func (obj *isisAdvanced) checkUnique() {

}

func (obj *isisAdvanced) checkConstraint() {

}

func (obj *isisAuthentication) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.AreaAuth != nil {

		obj.AreaAuth().validateObj(setDefault)
	}

	if obj.obj.DomainAuth != nil {

		obj.DomainAuth().validateObj(setDefault)
	}

}

func (obj *isisAuthentication) checkUnique() {

}

func (obj *isisAuthentication) checkConstraint() {

}

func (obj *isisV4RouteRange) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	if len(obj.obj.Addresses) != 0 {

		if setDefault {
			obj.Addresses().clearHolderSlice()
			for _, item := range obj.obj.Addresses {
				newObj := newV4RouteAddress(obj.validator)
				newObj.self().obj = item
				obj.Addresses().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Addresses().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.LinkMetric != nil {
		if *obj.obj.LinkMetric < 0 || *obj.obj.LinkMetric > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= IsisV4RouteRange.LinkMetric <= 16777215 but Got %d", *obj.obj.LinkMetric))
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface IsisV4RouteRange")
	}
}

func (obj *isisV4RouteRange) checkUnique() {
	if !obj.isUnique("isisV4RouteRange", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *isisV4RouteRange) checkConstraint() {

}

func (obj *isisV6RouteRange) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	if len(obj.obj.Addresses) != 0 {

		if setDefault {
			obj.Addresses().clearHolderSlice()
			for _, item := range obj.obj.Addresses {
				newObj := newV6RouteAddress(obj.validator)
				newObj.self().obj = item
				obj.Addresses().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Addresses().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.LinkMetric != nil {
		if *obj.obj.LinkMetric < 0 || *obj.obj.LinkMetric > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= IsisV6RouteRange.LinkMetric <= 16777215 but Got %d", *obj.obj.LinkMetric))
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface IsisV6RouteRange")
	}
}

func (obj *isisV6RouteRange) checkUnique() {
	if !obj.isUnique("isisV6RouteRange", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *isisV6RouteRange) checkConstraint() {

}

func (obj *bgpV4Interface) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// Ipv4Name is required
	if obj.obj.Ipv4Name == "" {
		obj.errors = append(obj.errors, "Ipv4Name is required field on interface BgpV4Interface")
	}

	if len(obj.obj.Peers) != 0 {

		if setDefault {
			obj.Peers().clearHolderSlice()
			for _, item := range obj.obj.Peers {
				newObj := newBgpV4Peer(obj.validator)
				newObj.self().obj = item
				obj.Peers().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Peers().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *bgpV4Interface) checkUnique() {

}

func (obj *bgpV4Interface) checkConstraint() {

	ipv4nameCons := []string{
		"deviceIpv4.Name", "deviceIpv4Loopback.Name",
	}

	if !obj.validateConstraint(ipv4nameCons, obj.Ipv4Name()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid deviceIpv4.Name|deviceIpv4Loopback.Name type", obj.Ipv4Name()),
		)
	}

}

func (obj *bgpV6Interface) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// Ipv6Name is required
	if obj.obj.Ipv6Name == "" {
		obj.errors = append(obj.errors, "Ipv6Name is required field on interface BgpV6Interface")
	}

	if len(obj.obj.Peers) != 0 {

		if setDefault {
			obj.Peers().clearHolderSlice()
			for _, item := range obj.obj.Peers {
				newObj := newBgpV6Peer(obj.validator)
				newObj.self().obj = item
				obj.Peers().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Peers().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *bgpV6Interface) checkUnique() {

}

func (obj *bgpV6Interface) checkConstraint() {

	ipv6nameCons := []string{
		"deviceIpv6.Name", "deviceIpv6Loopback.Name",
	}

	if !obj.validateConstraint(ipv6nameCons, obj.Ipv6Name()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid deviceIpv6.Name|deviceIpv6Loopback.Name type", obj.Ipv6Name()),
		)
	}

}

func (obj *vxlanV4Tunnel) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// SourceInterface is required
	if obj.obj.SourceInterface == "" {
		obj.errors = append(obj.errors, "SourceInterface is required field on interface VxlanV4Tunnel")
	}

	if obj.obj.DestinationIpMode != nil {

		obj.DestinationIpMode().validateObj(setDefault)
	}

	if obj.obj.Vni != 0 {
		if obj.obj.Vni < 1 || obj.obj.Vni > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= VxlanV4Tunnel.Vni <= 16777215 but Got %d", obj.obj.Vni))
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface VxlanV4Tunnel")
	}
}

func (obj *vxlanV4Tunnel) checkUnique() {
	if !obj.isUnique("vxlanV4Tunnel", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *vxlanV4Tunnel) checkConstraint() {

	sourceinterfaceCons := []string{
		"deviceIpv4.Name", "deviceIpv4Loopback.Name",
	}

	if !obj.validateConstraint(sourceinterfaceCons, obj.SourceInterface()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid deviceIpv4.Name|deviceIpv4Loopback.Name type", obj.SourceInterface()),
		)
	}

}

func (obj *vxlanV6Tunnel) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// SourceInterface is required
	if obj.obj.SourceInterface == "" {
		obj.errors = append(obj.errors, "SourceInterface is required field on interface VxlanV6Tunnel")
	}

	if obj.obj.DestinationIpMode != nil {

		obj.DestinationIpMode().validateObj(setDefault)
	}

	if obj.obj.Vni != 0 {
		if obj.obj.Vni < 1 || obj.obj.Vni > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= VxlanV6Tunnel.Vni <= 16777215 but Got %d", obj.obj.Vni))
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface VxlanV6Tunnel")
	}
}

func (obj *vxlanV6Tunnel) checkUnique() {
	if !obj.isUnique("vxlanV6Tunnel", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *vxlanV6Tunnel) checkConstraint() {

	sourceinterfaceCons := []string{
		"deviceIpv6.Name", "deviceIpv6Loopback.Name",
	}

	if !obj.validateConstraint(sourceinterfaceCons, obj.SourceInterface()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid deviceIpv6.Name|deviceIpv6Loopback.Name type", obj.SourceInterface()),
		)
	}

}

func (obj *flowPort) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// TxName is required
	if obj.obj.TxName == "" {
		obj.errors = append(obj.errors, "TxName is required field on interface FlowPort")
	}
}

func (obj *flowPort) checkUnique() {

}

func (obj *flowPort) checkConstraint() {

	txnameCons := []string{
		"port.Name", "lag.Name",
	}

	if !obj.validateConstraint(txnameCons, obj.TxName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid port.Name|lag.Name type", obj.TxName()),
		)
	}

	rxnameCons := []string{
		"port.Name", "lag.Name",
	}

	if !obj.validateConstraint(rxnameCons, obj.RxName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid port.Name|lag.Name type", obj.RxName()),
		)
	}

}

func (obj *flowRouter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"config"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// TxNames is required
	if obj.obj.TxNames == nil {
		obj.errors = append(obj.errors, "TxNames is required field on interface FlowRouter")
	}

	// RxNames is required
	if obj.obj.RxNames == nil {
		obj.errors = append(obj.errors, "RxNames is required field on interface FlowRouter")
	}
}

func (obj *flowRouter) checkUnique() {

}

func (obj *flowRouter) checkConstraint() {

	txnamesCons := []string{
		"deviceEthernet.Name", "deviceIpv4.Name", "deviceIpv6.Name", "bgpV4RouteRange.Name", "bgpV6RouteRange.Name", "bgpCMacIpRange.Name",
	}

	for _, v := range obj.TxNames() {
		if !obj.validateConstraint(txnamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid deviceEthernet.Name|deviceIpv4.Name|deviceIpv6.Name|bgpV4RouteRange.Name|bgpV6RouteRange.Name|bgpCMacIpRange.Name type", v),
			)
		}
	}

	rxnamesCons := []string{
		"deviceEthernet.Name", "deviceIpv4.Name", "deviceIpv6.Name", "bgpV4RouteRange.Name", "bgpV6RouteRange.Name", "bgpCMacIpRange.Name",
	}

	for _, v := range obj.RxNames() {
		if !obj.validateConstraint(rxnamesCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid deviceEthernet.Name|deviceIpv4.Name|deviceIpv6.Name|bgpV4RouteRange.Name|bgpV6RouteRange.Name|bgpCMacIpRange.Name type", v),
			)
		}
	}

}

func (obj *flowCustom) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// Bytes is required
	if obj.obj.Bytes == "" {
		obj.errors = append(obj.errors, "Bytes is required field on interface FlowCustom")
	}
}

func (obj *flowCustom) checkUnique() {

}

func (obj *flowCustom) checkConstraint() {

}

func (obj *flowEthernet) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Dst != nil {

		obj.Dst().validateObj(setDefault)
	}

	if obj.obj.Src != nil {

		obj.Src().validateObj(setDefault)
	}

	if obj.obj.EtherType != nil {

		obj.EtherType().validateObj(setDefault)
	}

	if obj.obj.PfcQueue != nil {

		obj.PfcQueue().validateObj(setDefault)
	}

}

func (obj *flowEthernet) checkUnique() {

}

func (obj *flowEthernet) checkConstraint() {

}

func (obj *flowVlan) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Priority != nil {

		obj.Priority().validateObj(setDefault)
	}

	if obj.obj.Cfi != nil {

		obj.Cfi().validateObj(setDefault)
	}

	if obj.obj.Id != nil {

		obj.Id().validateObj(setDefault)
	}

	if obj.obj.Tpid != nil {

		obj.Tpid().validateObj(setDefault)
	}

}

func (obj *flowVlan) checkUnique() {

}

func (obj *flowVlan) checkConstraint() {

}

func (obj *flowVxlan) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(setDefault)
	}

	if obj.obj.Reserved0 != nil {

		obj.Reserved0().validateObj(setDefault)
	}

	if obj.obj.Vni != nil {

		obj.Vni().validateObj(setDefault)
	}

	if obj.obj.Reserved1 != nil {

		obj.Reserved1().validateObj(setDefault)
	}

}

func (obj *flowVxlan) checkUnique() {

}

func (obj *flowVxlan) checkConstraint() {

}

func (obj *flowIpv4) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(setDefault)
	}

	if obj.obj.HeaderLength != nil {

		obj.HeaderLength().validateObj(setDefault)
	}

	if obj.obj.Priority != nil {

		obj.Priority().validateObj(setDefault)
	}

	if obj.obj.TotalLength != nil {

		obj.TotalLength().validateObj(setDefault)
	}

	if obj.obj.Identification != nil {

		obj.Identification().validateObj(setDefault)
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(setDefault)
	}

	if obj.obj.DontFragment != nil {

		obj.DontFragment().validateObj(setDefault)
	}

	if obj.obj.MoreFragments != nil {

		obj.MoreFragments().validateObj(setDefault)
	}

	if obj.obj.FragmentOffset != nil {

		obj.FragmentOffset().validateObj(setDefault)
	}

	if obj.obj.TimeToLive != nil {

		obj.TimeToLive().validateObj(setDefault)
	}

	if obj.obj.Protocol != nil {

		obj.Protocol().validateObj(setDefault)
	}

	if obj.obj.HeaderChecksum != nil {

		obj.HeaderChecksum().validateObj(setDefault)
	}

	if obj.obj.Src != nil {

		obj.Src().validateObj(setDefault)
	}

	if obj.obj.Dst != nil {

		obj.Dst().validateObj(setDefault)
	}

}

func (obj *flowIpv4) checkUnique() {

}

func (obj *flowIpv4) checkConstraint() {

}

func (obj *flowIpv6) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(setDefault)
	}

	if obj.obj.TrafficClass != nil {

		obj.TrafficClass().validateObj(setDefault)
	}

	if obj.obj.FlowLabel != nil {

		obj.FlowLabel().validateObj(setDefault)
	}

	if obj.obj.PayloadLength != nil {

		obj.PayloadLength().validateObj(setDefault)
	}

	if obj.obj.NextHeader != nil {

		obj.NextHeader().validateObj(setDefault)
	}

	if obj.obj.HopLimit != nil {

		obj.HopLimit().validateObj(setDefault)
	}

	if obj.obj.Src != nil {

		obj.Src().validateObj(setDefault)
	}

	if obj.obj.Dst != nil {

		obj.Dst().validateObj(setDefault)
	}

}

func (obj *flowIpv6) checkUnique() {

}

func (obj *flowIpv6) checkConstraint() {

}

func (obj *flowPfcPause) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Dst != nil {

		obj.Dst().validateObj(setDefault)
	}

	if obj.obj.Src != nil {

		obj.Src().validateObj(setDefault)
	}

	if obj.obj.EtherType != nil {

		obj.EtherType().validateObj(setDefault)
	}

	if obj.obj.ControlOpCode != nil {

		obj.ControlOpCode().validateObj(setDefault)
	}

	if obj.obj.ClassEnableVector != nil {

		obj.ClassEnableVector().validateObj(setDefault)
	}

	if obj.obj.PauseClass_0 != nil {

		obj.PauseClass0().validateObj(setDefault)
	}

	if obj.obj.PauseClass_1 != nil {

		obj.PauseClass1().validateObj(setDefault)
	}

	if obj.obj.PauseClass_2 != nil {

		obj.PauseClass2().validateObj(setDefault)
	}

	if obj.obj.PauseClass_3 != nil {

		obj.PauseClass3().validateObj(setDefault)
	}

	if obj.obj.PauseClass_4 != nil {

		obj.PauseClass4().validateObj(setDefault)
	}

	if obj.obj.PauseClass_5 != nil {

		obj.PauseClass5().validateObj(setDefault)
	}

	if obj.obj.PauseClass_6 != nil {

		obj.PauseClass6().validateObj(setDefault)
	}

	if obj.obj.PauseClass_7 != nil {

		obj.PauseClass7().validateObj(setDefault)
	}

}

func (obj *flowPfcPause) checkUnique() {

}

func (obj *flowPfcPause) checkConstraint() {

}

func (obj *flowEthernetPause) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Dst != nil {

		obj.Dst().validateObj(setDefault)
	}

	if obj.obj.Src != nil {

		obj.Src().validateObj(setDefault)
	}

	if obj.obj.EtherType != nil {

		obj.EtherType().validateObj(setDefault)
	}

	if obj.obj.ControlOpCode != nil {

		obj.ControlOpCode().validateObj(setDefault)
	}

	if obj.obj.Time != nil {

		obj.Time().validateObj(setDefault)
	}

}

func (obj *flowEthernetPause) checkUnique() {

}

func (obj *flowEthernetPause) checkConstraint() {

}

func (obj *flowTcp) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.SrcPort != nil {

		obj.SrcPort().validateObj(setDefault)
	}

	if obj.obj.DstPort != nil {

		obj.DstPort().validateObj(setDefault)
	}

	if obj.obj.SeqNum != nil {

		obj.SeqNum().validateObj(setDefault)
	}

	if obj.obj.AckNum != nil {

		obj.AckNum().validateObj(setDefault)
	}

	if obj.obj.DataOffset != nil {

		obj.DataOffset().validateObj(setDefault)
	}

	if obj.obj.EcnNs != nil {

		obj.EcnNs().validateObj(setDefault)
	}

	if obj.obj.EcnCwr != nil {

		obj.EcnCwr().validateObj(setDefault)
	}

	if obj.obj.EcnEcho != nil {

		obj.EcnEcho().validateObj(setDefault)
	}

	if obj.obj.CtlUrg != nil {

		obj.CtlUrg().validateObj(setDefault)
	}

	if obj.obj.CtlAck != nil {

		obj.CtlAck().validateObj(setDefault)
	}

	if obj.obj.CtlPsh != nil {

		obj.CtlPsh().validateObj(setDefault)
	}

	if obj.obj.CtlRst != nil {

		obj.CtlRst().validateObj(setDefault)
	}

	if obj.obj.CtlSyn != nil {

		obj.CtlSyn().validateObj(setDefault)
	}

	if obj.obj.CtlFin != nil {

		obj.CtlFin().validateObj(setDefault)
	}

	if obj.obj.Window != nil {

		obj.Window().validateObj(setDefault)
	}

}

func (obj *flowTcp) checkUnique() {

}

func (obj *flowTcp) checkConstraint() {

}

func (obj *flowUdp) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.SrcPort != nil {

		obj.SrcPort().validateObj(setDefault)
	}

	if obj.obj.DstPort != nil {

		obj.DstPort().validateObj(setDefault)
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(setDefault)
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(setDefault)
	}

}

func (obj *flowUdp) checkUnique() {

}

func (obj *flowUdp) checkConstraint() {

}

func (obj *flowGre) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.ChecksumPresent != nil {

		obj.ChecksumPresent().validateObj(setDefault)
	}

	if obj.obj.Reserved0 != nil {

		obj.Reserved0().validateObj(setDefault)
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(setDefault)
	}

	if obj.obj.Protocol != nil {

		obj.Protocol().validateObj(setDefault)
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(setDefault)
	}

	if obj.obj.Reserved1 != nil {

		obj.Reserved1().validateObj(setDefault)
	}

}

func (obj *flowGre) checkUnique() {

}

func (obj *flowGre) checkConstraint() {

}

func (obj *flowGtpv1) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(setDefault)
	}

	if obj.obj.ProtocolType != nil {

		obj.ProtocolType().validateObj(setDefault)
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(setDefault)
	}

	if obj.obj.EFlag != nil {

		obj.EFlag().validateObj(setDefault)
	}

	if obj.obj.SFlag != nil {

		obj.SFlag().validateObj(setDefault)
	}

	if obj.obj.PnFlag != nil {

		obj.PnFlag().validateObj(setDefault)
	}

	if obj.obj.MessageType != nil {

		obj.MessageType().validateObj(setDefault)
	}

	if obj.obj.MessageLength != nil {

		obj.MessageLength().validateObj(setDefault)
	}

	if obj.obj.Teid != nil {

		obj.Teid().validateObj(setDefault)
	}

	if obj.obj.SquenceNumber != nil {

		obj.SquenceNumber().validateObj(setDefault)
	}

	if obj.obj.NPduNumber != nil {

		obj.NPduNumber().validateObj(setDefault)
	}

	if obj.obj.NextExtensionHeaderType != nil {

		obj.NextExtensionHeaderType().validateObj(setDefault)
	}

	if len(obj.obj.ExtensionHeaders) != 0 {

		if setDefault {
			obj.ExtensionHeaders().clearHolderSlice()
			for _, item := range obj.obj.ExtensionHeaders {
				newObj := newFlowGtpExtension(obj.validator)
				newObj.self().obj = item
				obj.ExtensionHeaders().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.ExtensionHeaders().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *flowGtpv1) checkUnique() {

}

func (obj *flowGtpv1) checkConstraint() {

}

func (obj *flowGtpv2) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(setDefault)
	}

	if obj.obj.PiggybackingFlag != nil {

		obj.PiggybackingFlag().validateObj(setDefault)
	}

	if obj.obj.TeidFlag != nil {

		obj.TeidFlag().validateObj(setDefault)
	}

	if obj.obj.Spare1 != nil {

		obj.Spare1().validateObj(setDefault)
	}

	if obj.obj.MessageType != nil {

		obj.MessageType().validateObj(setDefault)
	}

	if obj.obj.MessageLength != nil {

		obj.MessageLength().validateObj(setDefault)
	}

	if obj.obj.Teid != nil {

		obj.Teid().validateObj(setDefault)
	}

	if obj.obj.SequenceNumber != nil {

		obj.SequenceNumber().validateObj(setDefault)
	}

	if obj.obj.Spare2 != nil {

		obj.Spare2().validateObj(setDefault)
	}

}

func (obj *flowGtpv2) checkUnique() {

}

func (obj *flowGtpv2) checkConstraint() {

}

func (obj *flowArp) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.HardwareType != nil {

		obj.HardwareType().validateObj(setDefault)
	}

	if obj.obj.ProtocolType != nil {

		obj.ProtocolType().validateObj(setDefault)
	}

	if obj.obj.HardwareLength != nil {

		obj.HardwareLength().validateObj(setDefault)
	}

	if obj.obj.ProtocolLength != nil {

		obj.ProtocolLength().validateObj(setDefault)
	}

	if obj.obj.Operation != nil {

		obj.Operation().validateObj(setDefault)
	}

	if obj.obj.SenderHardwareAddr != nil {

		obj.SenderHardwareAddr().validateObj(setDefault)
	}

	if obj.obj.SenderProtocolAddr != nil {

		obj.SenderProtocolAddr().validateObj(setDefault)
	}

	if obj.obj.TargetHardwareAddr != nil {

		obj.TargetHardwareAddr().validateObj(setDefault)
	}

	if obj.obj.TargetProtocolAddr != nil {

		obj.TargetProtocolAddr().validateObj(setDefault)
	}

}

func (obj *flowArp) checkUnique() {

}

func (obj *flowArp) checkConstraint() {

}

func (obj *flowIcmp) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Echo != nil {

		obj.Echo().validateObj(setDefault)
	}

}

func (obj *flowIcmp) checkUnique() {

}

func (obj *flowIcmp) checkConstraint() {

}

func (obj *flowIcmpv6) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Echo != nil {

		obj.Echo().validateObj(setDefault)
	}

}

func (obj *flowIcmpv6) checkUnique() {

}

func (obj *flowIcmpv6) checkConstraint() {

}

func (obj *flowPpp) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		obj.Address().validateObj(setDefault)
	}

	if obj.obj.Control != nil {

		obj.Control().validateObj(setDefault)
	}

	if obj.obj.ProtocolType != nil {

		obj.ProtocolType().validateObj(setDefault)
	}

}

func (obj *flowPpp) checkUnique() {

}

func (obj *flowPpp) checkConstraint() {

}

func (obj *flowIgmpv1) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(setDefault)
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(setDefault)
	}

	if obj.obj.Unused != nil {

		obj.Unused().validateObj(setDefault)
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(setDefault)
	}

	if obj.obj.GroupAddress != nil {

		obj.GroupAddress().validateObj(setDefault)
	}

}

func (obj *flowIgmpv1) checkUnique() {

}

func (obj *flowIgmpv1) checkConstraint() {

}

func (obj *flowMpls) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Label != nil {

		obj.Label().validateObj(setDefault)
	}

	if obj.obj.TrafficClass != nil {

		obj.TrafficClass().validateObj(setDefault)
	}

	if obj.obj.BottomOfStack != nil {

		obj.BottomOfStack().validateObj(setDefault)
	}

	if obj.obj.TimeToLive != nil {

		obj.TimeToLive().validateObj(setDefault)
	}

}

func (obj *flowMpls) checkUnique() {

}

func (obj *flowMpls) checkConstraint() {

}

func (obj *flowSizeIncrement) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 1 || *obj.obj.Start > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= FlowSizeIncrement.Start <= 2147483647 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.End != nil {
		if *obj.obj.End < 64 || *obj.obj.End > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("64 <= FlowSizeIncrement.End <= 2147483647 but Got %d", *obj.obj.End))
		}

	}

}

func (obj *flowSizeIncrement) checkUnique() {

}

func (obj *flowSizeIncrement) checkConstraint() {

}

func (obj *flowSizeRandom) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *flowSizeRandom) checkUnique() {

}

func (obj *flowSizeRandom) checkConstraint() {

}

func (obj *flowFixedPackets) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Packets != nil {
		if *obj.obj.Packets < 1 || *obj.obj.Packets > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= FlowFixedPackets.Packets <= 2147483647 but Got %d", *obj.obj.Packets))
		}

	}

	if obj.obj.Gap != nil {
		if *obj.obj.Gap < 0 || *obj.obj.Gap > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowFixedPackets.Gap <= 2147483647 but Got %d", *obj.obj.Gap))
		}

	}

	if obj.obj.Delay != nil {

		obj.Delay().validateObj(setDefault)
	}

}

func (obj *flowFixedPackets) checkUnique() {

}

func (obj *flowFixedPackets) checkConstraint() {

}

func (obj *flowFixedSeconds) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Seconds != nil {
		if *obj.obj.Seconds < 0 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowFixedSeconds.Seconds <= max(float32) but Got %f", *obj.obj.Seconds))
		}

	}

	if obj.obj.Gap != nil {
		if *obj.obj.Gap < 0 || *obj.obj.Gap > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowFixedSeconds.Gap <= 2147483647 but Got %d", *obj.obj.Gap))
		}

	}

	if obj.obj.Delay != nil {

		obj.Delay().validateObj(setDefault)
	}

}

func (obj *flowFixedSeconds) checkUnique() {

}

func (obj *flowFixedSeconds) checkConstraint() {

}

func (obj *flowBurst) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Bursts != nil {
		if *obj.obj.Bursts < 0 || *obj.obj.Bursts > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowBurst.Bursts <= 2147483647 but Got %d", *obj.obj.Bursts))
		}

	}

	if obj.obj.Packets != nil {
		if *obj.obj.Packets < 1 || *obj.obj.Packets > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= FlowBurst.Packets <= 2147483647 but Got %d", *obj.obj.Packets))
		}

	}

	if obj.obj.Gap != nil {
		if *obj.obj.Gap < 0 || *obj.obj.Gap > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowBurst.Gap <= 2147483647 but Got %d", *obj.obj.Gap))
		}

	}

	if obj.obj.InterBurstGap != nil {

		obj.InterBurstGap().validateObj(setDefault)
	}

}

func (obj *flowBurst) checkUnique() {

}

func (obj *flowBurst) checkConstraint() {

}

func (obj *flowContinuous) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Gap != nil {
		if *obj.obj.Gap < 0 || *obj.obj.Gap > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowContinuous.Gap <= 2147483647 but Got %d", *obj.obj.Gap))
		}

	}

	if obj.obj.Delay != nil {

		obj.Delay().validateObj(setDefault)
	}

}

func (obj *flowContinuous) checkUnique() {

}

func (obj *flowContinuous) checkConstraint() {

}

func (obj *flowLatencyMetrics) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *flowLatencyMetrics) checkUnique() {

}

func (obj *flowLatencyMetrics) checkConstraint() {

}

func (obj *flowMetricGroup) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *flowMetricGroup) checkUnique() {

}

func (obj *flowMetricGroup) checkConstraint() {

}

func (obj *metricTimestamp) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *metricTimestamp) checkUnique() {

}

func (obj *metricTimestamp) checkConstraint() {

}

func (obj *metricLatency) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *metricLatency) checkUnique() {

}

func (obj *metricLatency) checkConstraint() {

}

func (obj *bgpPrefixIpv4UnicastState) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Ipv4NextHop != nil {

		err := obj.validateIpv4(obj.Ipv4NextHop())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv4UnicastState.Ipv4NextHop"))
		}

	}

	if obj.obj.Ipv6NextHop != nil {

		err := obj.validateIpv6(obj.Ipv6NextHop())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv4UnicastState.Ipv6NextHop"))
		}

	}

	if len(obj.obj.Communities) != 0 {

		if setDefault {
			obj.Communities().clearHolderSlice()
			for _, item := range obj.obj.Communities {
				newObj := newResultBgpCommunity(obj.validator)
				newObj.self().obj = item
				obj.Communities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Communities().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(setDefault)
	}

}

func (obj *bgpPrefixIpv4UnicastState) checkUnique() {

}

func (obj *bgpPrefixIpv4UnicastState) checkConstraint() {

}

func (obj *bgpPrefixIpv6UnicastState) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Ipv4NextHop != nil {

		err := obj.validateIpv4(obj.Ipv4NextHop())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv6UnicastState.Ipv4NextHop"))
		}

	}

	if obj.obj.Ipv6NextHop != nil {

		err := obj.validateIpv6(obj.Ipv6NextHop())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv6UnicastState.Ipv6NextHop"))
		}

	}

	if len(obj.obj.Communities) != 0 {

		if setDefault {
			obj.Communities().clearHolderSlice()
			for _, item := range obj.obj.Communities {
				newObj := newResultBgpCommunity(obj.validator)
				newObj.self().obj = item
				obj.Communities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Communities().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(setDefault)
	}

}

func (obj *bgpPrefixIpv6UnicastState) checkUnique() {

}

func (obj *bgpPrefixIpv6UnicastState) checkConstraint() {

}

func (obj *isisLspState) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// LspId is required
	if obj.obj.LspId == "" {
		obj.errors = append(obj.errors, "LspId is required field on interface IsisLspState")
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(setDefault)
	}

	if obj.obj.Tlvs != nil {

		obj.Tlvs().validateObj(setDefault)
	}

}

func (obj *isisLspState) checkUnique() {

}

func (obj *isisLspState) checkConstraint() {

}

func (obj *captureField) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateHex(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on CaptureField.Value"))
		}

	}

	if obj.obj.Mask != nil {

		err := obj.validateHex(obj.Mask())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on CaptureField.Mask"))
		}

	}

}

func (obj *captureField) checkUnique() {

}

func (obj *captureField) checkConstraint() {

}

func (obj *deviceIpv4GatewayMAC) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Auto != nil {

		err := obj.validateMac(obj.Auto())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv4GatewayMAC.Auto"))
		}

	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv4GatewayMAC.Value"))
		}

	}

}

func (obj *deviceIpv4GatewayMAC) checkUnique() {

}

func (obj *deviceIpv4GatewayMAC) checkConstraint() {

}

func (obj *deviceIpv6GatewayMAC) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Auto != nil {

		err := obj.validateMac(obj.Auto())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv6GatewayMAC.Auto"))
		}

	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv6GatewayMAC.Value"))
		}

	}

}

func (obj *deviceIpv6GatewayMAC) checkUnique() {

}

func (obj *deviceIpv6GatewayMAC) checkConstraint() {

}

func (obj *isisInterfaceLevel) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *isisInterfaceLevel) checkUnique() {

}

func (obj *isisInterfaceLevel) checkConstraint() {

}

func (obj *isisMT) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.MtId != nil {
		if *obj.obj.MtId < 0 || *obj.obj.MtId > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= IsisMT.MtId <= 65535 but Got %d", *obj.obj.MtId))
		}

	}

	if obj.obj.LinkMetric != nil {
		if *obj.obj.LinkMetric < 0 || *obj.obj.LinkMetric > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= IsisMT.LinkMetric <= 16777215 but Got %d", *obj.obj.LinkMetric))
		}

	}

}

func (obj *isisMT) checkUnique() {

}

func (obj *isisMT) checkConstraint() {

}

func (obj *linkStateTE) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.AdministrativeGroup != nil {

		err := obj.validateHex(obj.AdministrativeGroup())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on LinkStateTE.AdministrativeGroup"))
		}

	}

	if obj.obj.MetricLevel != nil {
		if *obj.obj.MetricLevel < 0 || *obj.obj.MetricLevel > 4261412864 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LinkStateTE.MetricLevel <= 4261412864 but Got %d", *obj.obj.MetricLevel))
		}

	}

	if obj.obj.MaxBandwith != nil {
		if *obj.obj.MaxBandwith < 0 || *obj.obj.MaxBandwith > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LinkStateTE.MaxBandwith <= 4294967295 but Got %d", *obj.obj.MaxBandwith))
		}

	}

	if obj.obj.MaxReservableBandwidth != nil {
		if *obj.obj.MaxReservableBandwidth < 0 || *obj.obj.MaxReservableBandwidth > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LinkStateTE.MaxReservableBandwidth <= 4294967295 but Got %d", *obj.obj.MaxReservableBandwidth))
		}

	}

	if obj.obj.PriorityBandwidths != nil {

		obj.PriorityBandwidths().validateObj(setDefault)
	}

}

func (obj *linkStateTE) checkUnique() {

}

func (obj *linkStateTE) checkConstraint() {

}

func (obj *isisInterfaceAuthentication) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// AuthType is required
	if obj.obj.AuthType.Number() == 0 {
		obj.errors = append(obj.errors, "AuthType is required field on interface IsisInterfaceAuthentication")
	}

	if obj.obj.Md5 != nil {
		if len(*obj.obj.Md5) < 0 || len(*obj.obj.Md5) > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf(
					"0 <= length of IsisInterfaceAuthentication.Md5 <= 255 but Got %d",
					len(*obj.obj.Md5)))
		}

	}

	if obj.obj.Password != nil {
		if len(*obj.obj.Password) < 0 || len(*obj.obj.Password) > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf(
					"0 <= length of IsisInterfaceAuthentication.Password <= 255 but Got %d",
					len(*obj.obj.Password)))
		}

	}

}

func (obj *isisInterfaceAuthentication) checkUnique() {

}

func (obj *isisInterfaceAuthentication) checkConstraint() {

}

func (obj *isisInterfaceAdvanced) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *isisInterfaceAdvanced) checkUnique() {

}

func (obj *isisInterfaceAdvanced) checkConstraint() {

}

func (obj *isisInterfaceLinkProtection) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *isisInterfaceLinkProtection) checkUnique() {

}

func (obj *isisInterfaceLinkProtection) checkConstraint() {

}

func (obj *isisAuthenticationBase) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// AuthType is required
	if obj.obj.AuthType.Number() == 0 {
		obj.errors = append(obj.errors, "AuthType is required field on interface IsisAuthenticationBase")
	}

	if obj.obj.Md5 != nil {
		if len(*obj.obj.Md5) < 0 || len(*obj.obj.Md5) > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf(
					"0 <= length of IsisAuthenticationBase.Md5 <= 255 but Got %d",
					len(*obj.obj.Md5)))
		}

	}

	if obj.obj.Password != nil {
		if len(*obj.obj.Password) < 0 || len(*obj.obj.Password) > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf(
					"0 <= length of IsisAuthenticationBase.Password <= 255 but Got %d",
					len(*obj.obj.Password)))
		}

	}

}

func (obj *isisAuthenticationBase) checkUnique() {

}

func (obj *isisAuthenticationBase) checkConstraint() {

}

func (obj *v4RouteAddress) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// Address is required
	if obj.obj.Address == "" {
		obj.errors = append(obj.errors, "Address is required field on interface V4RouteAddress")
	}
	if obj.obj.Address != "" {

		err := obj.validateIpv4(obj.Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on V4RouteAddress.Address"))
		}

	}

	if obj.obj.Prefix != nil {
		if *obj.obj.Prefix < 0 || *obj.obj.Prefix > 32 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= V4RouteAddress.Prefix <= 32 but Got %d", *obj.obj.Prefix))
		}

	}

	if obj.obj.Count != nil {
		if *obj.obj.Count < 1 || *obj.obj.Count > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= V4RouteAddress.Count <= 2147483647 but Got %d", *obj.obj.Count))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 1 || *obj.obj.Step > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= V4RouteAddress.Step <= 2147483647 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *v4RouteAddress) checkUnique() {

}

func (obj *v4RouteAddress) checkConstraint() {

}

func (obj *v6RouteAddress) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// Address is required
	if obj.obj.Address == "" {
		obj.errors = append(obj.errors, "Address is required field on interface V6RouteAddress")
	}
	if obj.obj.Address != "" {

		err := obj.validateIpv6(obj.Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on V6RouteAddress.Address"))
		}

	}

	if obj.obj.Prefix != nil {
		if *obj.obj.Prefix < 0 || *obj.obj.Prefix > 128 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= V6RouteAddress.Prefix <= 128 but Got %d", *obj.obj.Prefix))
		}

	}

	if obj.obj.Count != nil {
		if *obj.obj.Count < 1 || *obj.obj.Count > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= V6RouteAddress.Count <= 2147483647 but Got %d", *obj.obj.Count))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 1 || *obj.obj.Step > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= V6RouteAddress.Step <= 2147483647 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *v6RouteAddress) checkUnique() {

}

func (obj *v6RouteAddress) checkConstraint() {

}

func (obj *bgpV4Peer) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	// PeerAddress is required
	if obj.obj.PeerAddress == "" {
		obj.errors = append(obj.errors, "PeerAddress is required field on interface BgpV4Peer")
	}
	if obj.obj.PeerAddress != "" {

		err := obj.validateIpv4(obj.PeerAddress())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpV4Peer.PeerAddress"))
		}

	}

	if len(obj.obj.EvpnEthernetSegments) != 0 {

		if setDefault {
			obj.EvpnEthernetSegments().clearHolderSlice()
			for _, item := range obj.obj.EvpnEthernetSegments {
				newObj := newBgpV4EthernetSegment(obj.validator)
				newObj.self().obj = item
				obj.EvpnEthernetSegments().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.EvpnEthernetSegments().Items() {
			item.validateObj(setDefault)
		}

	}

	// AsType is required
	if obj.obj.AsType.Number() == 0 {
		obj.errors = append(obj.errors, "AsType is required field on interface BgpV4Peer")
	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(setDefault)
	}

	if obj.obj.Capability != nil {

		obj.Capability().validateObj(setDefault)
	}

	if obj.obj.LearnedInformationFilter != nil {

		obj.LearnedInformationFilter().validateObj(setDefault)
	}

	if len(obj.obj.V4Routes) != 0 {

		if setDefault {
			obj.V4Routes().clearHolderSlice()
			for _, item := range obj.obj.V4Routes {
				newObj := newBgpV4RouteRange(obj.validator)
				newObj.self().obj = item
				obj.V4Routes().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.V4Routes().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.V6Routes) != 0 {

		if setDefault {
			obj.V6Routes().clearHolderSlice()
			for _, item := range obj.obj.V6Routes {
				newObj := newBgpV6RouteRange(obj.validator)
				newObj.self().obj = item
				obj.V6Routes().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.V6Routes().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.V4SrtePolicies) != 0 {

		if setDefault {
			obj.V4SrtePolicies().clearHolderSlice()
			for _, item := range obj.obj.V4SrtePolicies {
				newObj := newBgpSrteV4Policy(obj.validator)
				newObj.self().obj = item
				obj.V4SrtePolicies().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.V4SrtePolicies().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.V6SrtePolicies) != 0 {

		if setDefault {
			obj.V6SrtePolicies().clearHolderSlice()
			for _, item := range obj.obj.V6SrtePolicies {
				newObj := newBgpSrteV6Policy(obj.validator)
				newObj.self().obj = item
				obj.V6SrtePolicies().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.V6SrtePolicies().Items() {
			item.validateObj(setDefault)
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface BgpV4Peer")
	}
}

func (obj *bgpV4Peer) checkUnique() {
	if !obj.isUnique("bgpV4Peer", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *bgpV4Peer) checkConstraint() {

}

func (obj *bgpV6Peer) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	// PeerAddress is required
	if obj.obj.PeerAddress == "" {
		obj.errors = append(obj.errors, "PeerAddress is required field on interface BgpV6Peer")
	}
	if obj.obj.PeerAddress != "" {

		err := obj.validateIpv6(obj.PeerAddress())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpV6Peer.PeerAddress"))
		}

	}

	if obj.obj.SegmentRouting != nil {

		obj.SegmentRouting().validateObj(setDefault)
	}

	if len(obj.obj.EvpnEthernetSegments) != 0 {

		if setDefault {
			obj.EvpnEthernetSegments().clearHolderSlice()
			for _, item := range obj.obj.EvpnEthernetSegments {
				newObj := newBgpV6EthernetSegment(obj.validator)
				newObj.self().obj = item
				obj.EvpnEthernetSegments().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.EvpnEthernetSegments().Items() {
			item.validateObj(setDefault)
		}

	}

	// AsType is required
	if obj.obj.AsType.Number() == 0 {
		obj.errors = append(obj.errors, "AsType is required field on interface BgpV6Peer")
	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(setDefault)
	}

	if obj.obj.Capability != nil {

		obj.Capability().validateObj(setDefault)
	}

	if obj.obj.LearnedInformationFilter != nil {

		obj.LearnedInformationFilter().validateObj(setDefault)
	}

	if len(obj.obj.V4Routes) != 0 {

		if setDefault {
			obj.V4Routes().clearHolderSlice()
			for _, item := range obj.obj.V4Routes {
				newObj := newBgpV4RouteRange(obj.validator)
				newObj.self().obj = item
				obj.V4Routes().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.V4Routes().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.V6Routes) != 0 {

		if setDefault {
			obj.V6Routes().clearHolderSlice()
			for _, item := range obj.obj.V6Routes {
				newObj := newBgpV6RouteRange(obj.validator)
				newObj.self().obj = item
				obj.V6Routes().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.V6Routes().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.V4SrtePolicies) != 0 {

		if setDefault {
			obj.V4SrtePolicies().clearHolderSlice()
			for _, item := range obj.obj.V4SrtePolicies {
				newObj := newBgpSrteV4Policy(obj.validator)
				newObj.self().obj = item
				obj.V4SrtePolicies().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.V4SrtePolicies().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.V6SrtePolicies) != 0 {

		if setDefault {
			obj.V6SrtePolicies().clearHolderSlice()
			for _, item := range obj.obj.V6SrtePolicies {
				newObj := newBgpSrteV6Policy(obj.validator)
				newObj.self().obj = item
				obj.V6SrtePolicies().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.V6SrtePolicies().Items() {
			item.validateObj(setDefault)
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface BgpV6Peer")
	}
}

func (obj *bgpV6Peer) checkUnique() {
	if !obj.isUnique("bgpV6Peer", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *bgpV6Peer) checkConstraint() {

}

func (obj *vxlanV4TunnelDestinationIPMode) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Unicast != nil {

		obj.Unicast().validateObj(setDefault)
	}

	if obj.obj.Multicast != nil {

		obj.Multicast().validateObj(setDefault)
	}

}

func (obj *vxlanV4TunnelDestinationIPMode) checkUnique() {

}

func (obj *vxlanV4TunnelDestinationIPMode) checkConstraint() {

}

func (obj *vxlanV6TunnelDestinationIPMode) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Unicast != nil {

		obj.Unicast().validateObj(setDefault)
	}

	if obj.obj.Multicast != nil {

		obj.Multicast().validateObj(setDefault)
	}

}

func (obj *vxlanV6TunnelDestinationIPMode) checkUnique() {

}

func (obj *vxlanV6TunnelDestinationIPMode) checkConstraint() {

}

func (obj *patternFlowEthernetDst) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetDst.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetDst.Values"))
		}

	}

	if obj.obj.Auto != nil {

		err := obj.validateMac(obj.Auto())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetDst.Auto"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowEthernetDst) checkUnique() {

}

func (obj *patternFlowEthernetDst) checkConstraint() {

}

func (obj *patternFlowEthernetSrc) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetSrc.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetSrc.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowEthernetSrc) checkUnique() {

}

func (obj *patternFlowEthernetSrc) checkConstraint() {

}

func (obj *patternFlowEthernetEtherType) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetEtherType.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowEthernetEtherType.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {
		if *obj.obj.Auto < 0 || *obj.obj.Auto > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetEtherType.Auto <= 65535 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowEthernetEtherType) checkUnique() {

}

func (obj *patternFlowEthernetEtherType) checkConstraint() {

}

func (obj *patternFlowEthernetPfcQueue) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetPfcQueue.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 7 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowEthernetPfcQueue.Values <= 7 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowEthernetPfcQueue) checkUnique() {

}

func (obj *patternFlowEthernetPfcQueue) checkConstraint() {

}

func (obj *patternFlowVlanPriority) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVlanPriority.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 7 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowVlanPriority.Values <= 7 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowVlanPriority) checkUnique() {

}

func (obj *patternFlowVlanPriority) checkConstraint() {

}

func (obj *patternFlowVlanCfi) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVlanCfi.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowVlanCfi.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowVlanCfi) checkUnique() {

}

func (obj *patternFlowVlanCfi) checkConstraint() {

}

func (obj *patternFlowVlanId) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 4095 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVlanId.Value <= 4095 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 4095 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowVlanId.Values <= 4095 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowVlanId) checkUnique() {

}

func (obj *patternFlowVlanId) checkConstraint() {

}

func (obj *patternFlowVlanTpid) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVlanTpid.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowVlanTpid.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowVlanTpid) checkUnique() {

}

func (obj *patternFlowVlanTpid) checkConstraint() {

}

func (obj *patternFlowVxlanFlags) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVxlanFlags.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowVxlanFlags.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowVxlanFlags) checkUnique() {

}

func (obj *patternFlowVxlanFlags) checkConstraint() {

}

func (obj *patternFlowVxlanReserved0) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved0.Value <= 16777215 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 16777215 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowVxlanReserved0.Values <= 16777215 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowVxlanReserved0) checkUnique() {

}

func (obj *patternFlowVxlanReserved0) checkConstraint() {

}

func (obj *patternFlowVxlanVni) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVxlanVni.Value <= 16777215 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 16777215 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowVxlanVni.Values <= 16777215 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {
		if *obj.obj.Auto < 0 || *obj.obj.Auto > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVxlanVni.Auto <= 16777215 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowVxlanVni) checkUnique() {

}

func (obj *patternFlowVxlanVni) checkConstraint() {

}

func (obj *patternFlowVxlanReserved1) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved1.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowVxlanReserved1.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowVxlanReserved1) checkUnique() {

}

func (obj *patternFlowVxlanReserved1) checkConstraint() {

}

func (obj *patternFlowIpv4Version) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4Version.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 15 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4Version.Values <= 15 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4Version) checkUnique() {

}

func (obj *patternFlowIpv4Version) checkConstraint() {

}

func (obj *patternFlowIpv4HeaderLength) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4HeaderLength.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 15 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4HeaderLength.Values <= 15 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {
		if *obj.obj.Auto < 0 || *obj.obj.Auto > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4HeaderLength.Auto <= 15 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4HeaderLength) checkUnique() {

}

func (obj *patternFlowIpv4HeaderLength) checkConstraint() {

}

func (obj *flowIpv4Priority) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Raw != nil {

		obj.Raw().validateObj(setDefault)
	}

	if obj.obj.Tos != nil {

		obj.Tos().validateObj(setDefault)
	}

	if obj.obj.Dscp != nil {

		obj.Dscp().validateObj(setDefault)
	}

}

func (obj *flowIpv4Priority) checkUnique() {

}

func (obj *flowIpv4Priority) checkConstraint() {

}

func (obj *patternFlowIpv4TotalLength) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TotalLength.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4TotalLength.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {
		if *obj.obj.Auto < 0 || *obj.obj.Auto > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TotalLength.Auto <= 65535 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4TotalLength) checkUnique() {

}

func (obj *patternFlowIpv4TotalLength) checkConstraint() {

}

func (obj *patternFlowIpv4Identification) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4Identification.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4Identification.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4Identification) checkUnique() {

}

func (obj *patternFlowIpv4Identification) checkConstraint() {

}

func (obj *patternFlowIpv4Reserved) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4Reserved.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4Reserved.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4Reserved) checkUnique() {

}

func (obj *patternFlowIpv4Reserved) checkConstraint() {

}

func (obj *patternFlowIpv4DontFragment) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4DontFragment.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4DontFragment.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4DontFragment) checkUnique() {

}

func (obj *patternFlowIpv4DontFragment) checkConstraint() {

}

func (obj *patternFlowIpv4MoreFragments) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4MoreFragments.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4MoreFragments.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4MoreFragments) checkUnique() {

}

func (obj *patternFlowIpv4MoreFragments) checkConstraint() {

}

func (obj *patternFlowIpv4FragmentOffset) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 31 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4FragmentOffset.Value <= 31 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 31 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4FragmentOffset.Values <= 31 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4FragmentOffset) checkUnique() {

}

func (obj *patternFlowIpv4FragmentOffset) checkConstraint() {

}

func (obj *patternFlowIpv4TimeToLive) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TimeToLive.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4TimeToLive.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4TimeToLive) checkUnique() {

}

func (obj *patternFlowIpv4TimeToLive) checkConstraint() {

}

func (obj *patternFlowIpv4Protocol) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4Protocol.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4Protocol.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {
		if *obj.obj.Auto < 0 || *obj.obj.Auto > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4Protocol.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4Protocol) checkUnique() {

}

func (obj *patternFlowIpv4Protocol) checkConstraint() {

}

func (obj *patternFlowIpv4HeaderChecksum) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {
		if *obj.obj.Custom < 0 || *obj.obj.Custom > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4HeaderChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowIpv4HeaderChecksum) checkUnique() {

}

func (obj *patternFlowIpv4HeaderChecksum) checkConstraint() {

}

func (obj *patternFlowIpv4Src) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4Src.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4Src.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4Src) checkUnique() {

}

func (obj *patternFlowIpv4Src) checkConstraint() {

}

func (obj *patternFlowIpv4Dst) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4Dst.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4Dst.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4Dst) checkUnique() {

}

func (obj *patternFlowIpv4Dst) checkConstraint() {

}

func (obj *patternFlowIpv6Version) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6Version.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 15 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv6Version.Values <= 15 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv6Version) checkUnique() {

}

func (obj *patternFlowIpv6Version) checkConstraint() {

}

func (obj *patternFlowIpv6TrafficClass) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6TrafficClass.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv6TrafficClass.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv6TrafficClass) checkUnique() {

}

func (obj *patternFlowIpv6TrafficClass) checkConstraint() {

}

func (obj *patternFlowIpv6FlowLabel) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1048575 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6FlowLabel.Value <= 1048575 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1048575 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv6FlowLabel.Values <= 1048575 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv6FlowLabel) checkUnique() {

}

func (obj *patternFlowIpv6FlowLabel) checkConstraint() {

}

func (obj *patternFlowIpv6PayloadLength) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6PayloadLength.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv6PayloadLength.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {
		if *obj.obj.Auto < 0 || *obj.obj.Auto > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6PayloadLength.Auto <= 65535 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv6PayloadLength) checkUnique() {

}

func (obj *patternFlowIpv6PayloadLength) checkConstraint() {

}

func (obj *patternFlowIpv6NextHeader) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6NextHeader.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv6NextHeader.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {
		if *obj.obj.Auto < 0 || *obj.obj.Auto > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6NextHeader.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv6NextHeader) checkUnique() {

}

func (obj *patternFlowIpv6NextHeader) checkConstraint() {

}

func (obj *patternFlowIpv6HopLimit) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6HopLimit.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv6HopLimit.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv6HopLimit) checkUnique() {

}

func (obj *patternFlowIpv6HopLimit) checkConstraint() {

}

func (obj *patternFlowIpv6Src) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6Src.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6Src.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv6Src) checkUnique() {

}

func (obj *patternFlowIpv6Src) checkConstraint() {

}

func (obj *patternFlowIpv6Dst) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6Dst.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6Dst.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv6Dst) checkUnique() {

}

func (obj *patternFlowIpv6Dst) checkConstraint() {

}

func (obj *patternFlowPfcPauseDst) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseDst.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseDst.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPfcPauseDst) checkUnique() {

}

func (obj *patternFlowPfcPauseDst) checkConstraint() {

}

func (obj *patternFlowPfcPauseSrc) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseSrc.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseSrc.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPfcPauseSrc) checkUnique() {

}

func (obj *patternFlowPfcPauseSrc) checkConstraint() {

}

func (obj *patternFlowPfcPauseEtherType) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseEtherType.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowPfcPauseEtherType.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPfcPauseEtherType) checkUnique() {

}

func (obj *patternFlowPfcPauseEtherType) checkConstraint() {

}

func (obj *patternFlowPfcPauseControlOpCode) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseControlOpCode.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowPfcPauseControlOpCode.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPfcPauseControlOpCode) checkUnique() {

}

func (obj *patternFlowPfcPauseControlOpCode) checkConstraint() {

}

func (obj *patternFlowPfcPauseClassEnableVector) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseClassEnableVector.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowPfcPauseClassEnableVector.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPfcPauseClassEnableVector) checkUnique() {

}

func (obj *patternFlowPfcPauseClassEnableVector) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass0) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass0.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass0.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPfcPausePauseClass0) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass0) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass1) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass1.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass1.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPfcPausePauseClass1) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass1) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass2) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass2.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass2.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPfcPausePauseClass2) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass2) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass3) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass3.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass3.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPfcPausePauseClass3) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass3) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass4) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass4.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass4.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPfcPausePauseClass4) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass4) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass5) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass5.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass5.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPfcPausePauseClass5) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass5) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass6) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass6.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass6.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPfcPausePauseClass6) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass6) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass7) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass7.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass7.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPfcPausePauseClass7) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass7) checkConstraint() {

}

func (obj *patternFlowEthernetPauseDst) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseDst.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseDst.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowEthernetPauseDst) checkUnique() {

}

func (obj *patternFlowEthernetPauseDst) checkConstraint() {

}

func (obj *patternFlowEthernetPauseSrc) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseSrc.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseSrc.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowEthernetPauseSrc) checkUnique() {

}

func (obj *patternFlowEthernetPauseSrc) checkConstraint() {

}

func (obj *patternFlowEthernetPauseEtherType) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseEtherType.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowEthernetPauseEtherType.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowEthernetPauseEtherType) checkUnique() {

}

func (obj *patternFlowEthernetPauseEtherType) checkConstraint() {

}

func (obj *patternFlowEthernetPauseControlOpCode) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseControlOpCode.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowEthernetPauseControlOpCode.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowEthernetPauseControlOpCode) checkUnique() {

}

func (obj *patternFlowEthernetPauseControlOpCode) checkConstraint() {

}

func (obj *patternFlowEthernetPauseTime) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseTime.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowEthernetPauseTime.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowEthernetPauseTime) checkUnique() {

}

func (obj *patternFlowEthernetPauseTime) checkConstraint() {

}

func (obj *patternFlowTcpSrcPort) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpSrcPort.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowTcpSrcPort.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowTcpSrcPort) checkUnique() {

}

func (obj *patternFlowTcpSrcPort) checkConstraint() {

}

func (obj *patternFlowTcpDstPort) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpDstPort.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowTcpDstPort.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowTcpDstPort) checkUnique() {

}

func (obj *patternFlowTcpDstPort) checkConstraint() {

}

func (obj *patternFlowTcpSeqNum) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpSeqNum.Value <= 4294967295 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 4294967295 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowTcpSeqNum.Values <= 4294967295 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowTcpSeqNum) checkUnique() {

}

func (obj *patternFlowTcpSeqNum) checkConstraint() {

}

func (obj *patternFlowTcpAckNum) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpAckNum.Value <= 4294967295 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 4294967295 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowTcpAckNum.Values <= 4294967295 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowTcpAckNum) checkUnique() {

}

func (obj *patternFlowTcpAckNum) checkConstraint() {

}

func (obj *patternFlowTcpDataOffset) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpDataOffset.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 15 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowTcpDataOffset.Values <= 15 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowTcpDataOffset) checkUnique() {

}

func (obj *patternFlowTcpDataOffset) checkConstraint() {

}

func (obj *patternFlowTcpEcnNs) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnNs.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowTcpEcnNs.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowTcpEcnNs) checkUnique() {

}

func (obj *patternFlowTcpEcnNs) checkConstraint() {

}

func (obj *patternFlowTcpEcnCwr) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnCwr.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowTcpEcnCwr.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowTcpEcnCwr) checkUnique() {

}

func (obj *patternFlowTcpEcnCwr) checkConstraint() {

}

func (obj *patternFlowTcpEcnEcho) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnEcho.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowTcpEcnEcho.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowTcpEcnEcho) checkUnique() {

}

func (obj *patternFlowTcpEcnEcho) checkConstraint() {

}

func (obj *patternFlowTcpCtlUrg) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlUrg.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowTcpCtlUrg.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowTcpCtlUrg) checkUnique() {

}

func (obj *patternFlowTcpCtlUrg) checkConstraint() {

}

func (obj *patternFlowTcpCtlAck) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlAck.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowTcpCtlAck.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowTcpCtlAck) checkUnique() {

}

func (obj *patternFlowTcpCtlAck) checkConstraint() {

}

func (obj *patternFlowTcpCtlPsh) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlPsh.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowTcpCtlPsh.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowTcpCtlPsh) checkUnique() {

}

func (obj *patternFlowTcpCtlPsh) checkConstraint() {

}

func (obj *patternFlowTcpCtlRst) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlRst.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowTcpCtlRst.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowTcpCtlRst) checkUnique() {

}

func (obj *patternFlowTcpCtlRst) checkConstraint() {

}

func (obj *patternFlowTcpCtlSyn) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlSyn.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowTcpCtlSyn.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowTcpCtlSyn) checkUnique() {

}

func (obj *patternFlowTcpCtlSyn) checkConstraint() {

}

func (obj *patternFlowTcpCtlFin) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlFin.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowTcpCtlFin.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowTcpCtlFin) checkUnique() {

}

func (obj *patternFlowTcpCtlFin) checkConstraint() {

}

func (obj *patternFlowTcpWindow) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpWindow.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowTcpWindow.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowTcpWindow) checkUnique() {

}

func (obj *patternFlowTcpWindow) checkConstraint() {

}

func (obj *patternFlowUdpSrcPort) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowUdpSrcPort.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowUdpSrcPort.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowUdpSrcPort) checkUnique() {

}

func (obj *patternFlowUdpSrcPort) checkConstraint() {

}

func (obj *patternFlowUdpDstPort) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowUdpDstPort.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowUdpDstPort.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowUdpDstPort) checkUnique() {

}

func (obj *patternFlowUdpDstPort) checkConstraint() {

}

func (obj *patternFlowUdpLength) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowUdpLength.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowUdpLength.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowUdpLength) checkUnique() {

}

func (obj *patternFlowUdpLength) checkConstraint() {

}

func (obj *patternFlowUdpChecksum) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {
		if *obj.obj.Custom < 0 || *obj.obj.Custom > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowUdpChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowUdpChecksum) checkUnique() {

}

func (obj *patternFlowUdpChecksum) checkConstraint() {

}

func (obj *patternFlowGreChecksumPresent) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreChecksumPresent.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGreChecksumPresent.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGreChecksumPresent) checkUnique() {

}

func (obj *patternFlowGreChecksumPresent) checkConstraint() {

}

func (obj *patternFlowGreReserved0) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 4095 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreReserved0.Value <= 4095 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 4095 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGreReserved0.Values <= 4095 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGreReserved0) checkUnique() {

}

func (obj *patternFlowGreReserved0) checkConstraint() {

}

func (obj *patternFlowGreVersion) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreVersion.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 7 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGreVersion.Values <= 7 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGreVersion) checkUnique() {

}

func (obj *patternFlowGreVersion) checkConstraint() {

}

func (obj *patternFlowGreProtocol) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreProtocol.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGreProtocol.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGreProtocol) checkUnique() {

}

func (obj *patternFlowGreProtocol) checkConstraint() {

}

func (obj *patternFlowGreChecksum) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {
		if *obj.obj.Custom < 0 || *obj.obj.Custom > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowGreChecksum) checkUnique() {

}

func (obj *patternFlowGreChecksum) checkConstraint() {

}

func (obj *patternFlowGreReserved1) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreReserved1.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGreReserved1.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGreReserved1) checkUnique() {

}

func (obj *patternFlowGreReserved1) checkConstraint() {

}

func (obj *patternFlowGtpv1Version) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1Version.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 7 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv1Version.Values <= 7 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv1Version) checkUnique() {

}

func (obj *patternFlowGtpv1Version) checkConstraint() {

}

func (obj *patternFlowGtpv1ProtocolType) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1ProtocolType.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv1ProtocolType.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv1ProtocolType) checkUnique() {

}

func (obj *patternFlowGtpv1ProtocolType) checkConstraint() {

}

func (obj *patternFlowGtpv1Reserved) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1Reserved.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv1Reserved.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv1Reserved) checkUnique() {

}

func (obj *patternFlowGtpv1Reserved) checkConstraint() {

}

func (obj *patternFlowGtpv1EFlag) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1EFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv1EFlag.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv1EFlag) checkUnique() {

}

func (obj *patternFlowGtpv1EFlag) checkConstraint() {

}

func (obj *patternFlowGtpv1SFlag) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv1SFlag.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv1SFlag) checkUnique() {

}

func (obj *patternFlowGtpv1SFlag) checkConstraint() {

}

func (obj *patternFlowGtpv1PnFlag) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1PnFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv1PnFlag.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv1PnFlag) checkUnique() {

}

func (obj *patternFlowGtpv1PnFlag) checkConstraint() {

}

func (obj *patternFlowGtpv1MessageType) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv1MessageType.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv1MessageType) checkUnique() {

}

func (obj *patternFlowGtpv1MessageType) checkConstraint() {

}

func (obj *patternFlowGtpv1MessageLength) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageLength.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv1MessageLength.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv1MessageLength) checkUnique() {

}

func (obj *patternFlowGtpv1MessageLength) checkConstraint() {

}

func (obj *patternFlowGtpv1Teid) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1Teid.Value <= 4294967295 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 4294967295 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv1Teid.Values <= 4294967295 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv1Teid) checkUnique() {

}

func (obj *patternFlowGtpv1Teid) checkConstraint() {

}

func (obj *patternFlowGtpv1SquenceNumber) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SquenceNumber.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv1SquenceNumber.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv1SquenceNumber) checkUnique() {

}

func (obj *patternFlowGtpv1SquenceNumber) checkConstraint() {

}

func (obj *patternFlowGtpv1NPduNumber) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NPduNumber.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv1NPduNumber.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv1NPduNumber) checkUnique() {

}

func (obj *patternFlowGtpv1NPduNumber) checkConstraint() {

}

func (obj *patternFlowGtpv1NextExtensionHeaderType) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NextExtensionHeaderType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv1NextExtensionHeaderType.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv1NextExtensionHeaderType) checkUnique() {

}

func (obj *patternFlowGtpv1NextExtensionHeaderType) checkConstraint() {

}

func (obj *flowGtpExtension) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.ExtensionLength != nil {

		obj.ExtensionLength().validateObj(setDefault)
	}

	if obj.obj.Contents != nil {

		obj.Contents().validateObj(setDefault)
	}

	if obj.obj.NextExtensionHeader != nil {

		obj.NextExtensionHeader().validateObj(setDefault)
	}

}

func (obj *flowGtpExtension) checkUnique() {

}

func (obj *flowGtpExtension) checkConstraint() {

}

func (obj *patternFlowGtpv2Version) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2Version.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 7 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv2Version.Values <= 7 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv2Version) checkUnique() {

}

func (obj *patternFlowGtpv2Version) checkConstraint() {

}

func (obj *patternFlowGtpv2PiggybackingFlag) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2PiggybackingFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv2PiggybackingFlag.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv2PiggybackingFlag) checkUnique() {

}

func (obj *patternFlowGtpv2PiggybackingFlag) checkConstraint() {

}

func (obj *patternFlowGtpv2TeidFlag) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2TeidFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv2TeidFlag.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv2TeidFlag) checkUnique() {

}

func (obj *patternFlowGtpv2TeidFlag) checkConstraint() {

}

func (obj *patternFlowGtpv2Spare1) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2Spare1.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 7 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv2Spare1.Values <= 7 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv2Spare1) checkUnique() {

}

func (obj *patternFlowGtpv2Spare1) checkConstraint() {

}

func (obj *patternFlowGtpv2MessageType) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv2MessageType.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv2MessageType) checkUnique() {

}

func (obj *patternFlowGtpv2MessageType) checkConstraint() {

}

func (obj *patternFlowGtpv2MessageLength) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageLength.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv2MessageLength.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv2MessageLength) checkUnique() {

}

func (obj *patternFlowGtpv2MessageLength) checkConstraint() {

}

func (obj *patternFlowGtpv2Teid) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2Teid.Value <= 4294967295 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 4294967295 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv2Teid.Values <= 4294967295 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv2Teid) checkUnique() {

}

func (obj *patternFlowGtpv2Teid) checkConstraint() {

}

func (obj *patternFlowGtpv2SequenceNumber) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2SequenceNumber.Value <= 16777215 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 16777215 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv2SequenceNumber.Values <= 16777215 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv2SequenceNumber) checkUnique() {

}

func (obj *patternFlowGtpv2SequenceNumber) checkConstraint() {

}

func (obj *patternFlowGtpv2Spare2) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2Spare2.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpv2Spare2.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpv2Spare2) checkUnique() {

}

func (obj *patternFlowGtpv2Spare2) checkConstraint() {

}

func (obj *patternFlowArpHardwareType) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareType.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowArpHardwareType.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowArpHardwareType) checkUnique() {

}

func (obj *patternFlowArpHardwareType) checkConstraint() {

}

func (obj *patternFlowArpProtocolType) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolType.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowArpProtocolType.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowArpProtocolType) checkUnique() {

}

func (obj *patternFlowArpProtocolType) checkConstraint() {

}

func (obj *patternFlowArpHardwareLength) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowArpHardwareLength.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowArpHardwareLength) checkUnique() {

}

func (obj *patternFlowArpHardwareLength) checkConstraint() {

}

func (obj *patternFlowArpProtocolLength) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowArpProtocolLength.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowArpProtocolLength) checkUnique() {

}

func (obj *patternFlowArpProtocolLength) checkConstraint() {

}

func (obj *patternFlowArpOperation) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowArpOperation.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowArpOperation.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowArpOperation) checkUnique() {

}

func (obj *patternFlowArpOperation) checkConstraint() {

}

func (obj *patternFlowArpSenderHardwareAddr) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderHardwareAddr.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderHardwareAddr.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowArpSenderHardwareAddr) checkUnique() {

}

func (obj *patternFlowArpSenderHardwareAddr) checkConstraint() {

}

func (obj *patternFlowArpSenderProtocolAddr) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderProtocolAddr.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderProtocolAddr.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowArpSenderProtocolAddr) checkUnique() {

}

func (obj *patternFlowArpSenderProtocolAddr) checkConstraint() {

}

func (obj *patternFlowArpTargetHardwareAddr) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetHardwareAddr.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetHardwareAddr.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowArpTargetHardwareAddr) checkUnique() {

}

func (obj *patternFlowArpTargetHardwareAddr) checkConstraint() {

}

func (obj *patternFlowArpTargetProtocolAddr) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetProtocolAddr.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetProtocolAddr.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowArpTargetProtocolAddr) checkUnique() {

}

func (obj *patternFlowArpTargetProtocolAddr) checkConstraint() {

}

func (obj *flowIcmpEcho) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(setDefault)
	}

	if obj.obj.Code != nil {

		obj.Code().validateObj(setDefault)
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(setDefault)
	}

	if obj.obj.Identifier != nil {

		obj.Identifier().validateObj(setDefault)
	}

	if obj.obj.SequenceNumber != nil {

		obj.SequenceNumber().validateObj(setDefault)
	}

}

func (obj *flowIcmpEcho) checkUnique() {

}

func (obj *flowIcmpEcho) checkConstraint() {

}

func (obj *flowIcmpv6Echo) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(setDefault)
	}

	if obj.obj.Code != nil {

		obj.Code().validateObj(setDefault)
	}

	if obj.obj.Identifier != nil {

		obj.Identifier().validateObj(setDefault)
	}

	if obj.obj.SequenceNumber != nil {

		obj.SequenceNumber().validateObj(setDefault)
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(setDefault)
	}

}

func (obj *flowIcmpv6Echo) checkUnique() {

}

func (obj *flowIcmpv6Echo) checkConstraint() {

}

func (obj *patternFlowPppAddress) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPppAddress.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowPppAddress.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPppAddress) checkUnique() {

}

func (obj *patternFlowPppAddress) checkConstraint() {

}

func (obj *patternFlowPppControl) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPppControl.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowPppControl.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPppControl) checkUnique() {

}

func (obj *patternFlowPppControl) checkConstraint() {

}

func (obj *patternFlowPppProtocolType) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPppProtocolType.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowPppProtocolType.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {
		if *obj.obj.Auto < 0 || *obj.obj.Auto > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPppProtocolType.Auto <= 65535 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowPppProtocolType) checkUnique() {

}

func (obj *patternFlowPppProtocolType) checkConstraint() {

}

func (obj *patternFlowIgmpv1Version) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1Version.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 15 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIgmpv1Version.Values <= 15 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIgmpv1Version) checkUnique() {

}

func (obj *patternFlowIgmpv1Version) checkConstraint() {

}

func (obj *patternFlowIgmpv1Type) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1Type.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 15 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIgmpv1Type.Values <= 15 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIgmpv1Type) checkUnique() {

}

func (obj *patternFlowIgmpv1Type) checkConstraint() {

}

func (obj *patternFlowIgmpv1Unused) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1Unused.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIgmpv1Unused.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIgmpv1Unused) checkUnique() {

}

func (obj *patternFlowIgmpv1Unused) checkConstraint() {

}

func (obj *patternFlowIgmpv1Checksum) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {
		if *obj.obj.Custom < 0 || *obj.obj.Custom > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1Checksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowIgmpv1Checksum) checkUnique() {

}

func (obj *patternFlowIgmpv1Checksum) checkConstraint() {

}

func (obj *patternFlowIgmpv1GroupAddress) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv1GroupAddress.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv1GroupAddress.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIgmpv1GroupAddress) checkUnique() {

}

func (obj *patternFlowIgmpv1GroupAddress) checkConstraint() {

}

func (obj *patternFlowMplsLabel) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1048575 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowMplsLabel.Value <= 1048575 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1048575 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowMplsLabel.Values <= 1048575 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {
		if *obj.obj.Auto < 0 || *obj.obj.Auto > 1048575 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowMplsLabel.Auto <= 1048575 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowMplsLabel) checkUnique() {

}

func (obj *patternFlowMplsLabel) checkConstraint() {

}

func (obj *patternFlowMplsTrafficClass) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowMplsTrafficClass.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 7 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowMplsTrafficClass.Values <= 7 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowMplsTrafficClass) checkUnique() {

}

func (obj *patternFlowMplsTrafficClass) checkConstraint() {

}

func (obj *patternFlowMplsBottomOfStack) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowMplsBottomOfStack.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowMplsBottomOfStack.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {
		if *obj.obj.Auto < 0 || *obj.obj.Auto > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowMplsBottomOfStack.Auto <= 1 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowMplsBottomOfStack) checkUnique() {

}

func (obj *patternFlowMplsBottomOfStack) checkConstraint() {

}

func (obj *patternFlowMplsTimeToLive) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowMplsTimeToLive.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowMplsTimeToLive.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowMplsTimeToLive) checkUnique() {

}

func (obj *patternFlowMplsTimeToLive) checkConstraint() {

}

func (obj *flowDelay) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Bytes != nil {
		if *obj.obj.Bytes < 0 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowDelay.Bytes <= max(float32) but Got %f", *obj.obj.Bytes))
		}

	}

	if obj.obj.Nanoseconds != nil {
		if *obj.obj.Nanoseconds < 0 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowDelay.Nanoseconds <= max(float32) but Got %f", *obj.obj.Nanoseconds))
		}

	}

	if obj.obj.Microseconds != nil {
		if *obj.obj.Microseconds < 0 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowDelay.Microseconds <= max(float32) but Got %f", *obj.obj.Microseconds))
		}

	}

}

func (obj *flowDelay) checkUnique() {

}

func (obj *flowDelay) checkConstraint() {

}

func (obj *flowDurationInterBurstGap) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Bytes != nil {
		if *obj.obj.Bytes < 0 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowDurationInterBurstGap.Bytes <= max(float64) but Got %f", *obj.obj.Bytes))
		}

	}

	if obj.obj.Nanoseconds != nil {
		if *obj.obj.Nanoseconds < 0 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowDurationInterBurstGap.Nanoseconds <= max(float64) but Got %f", *obj.obj.Nanoseconds))
		}

	}

	if obj.obj.Microseconds != nil {
		if *obj.obj.Microseconds < 0 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= FlowDurationInterBurstGap.Microseconds <= max(float64) but Got %f", *obj.obj.Microseconds))
		}

	}

}

func (obj *flowDurationInterBurstGap) checkUnique() {

}

func (obj *flowDurationInterBurstGap) checkConstraint() {

}

func (obj *resultBgpCommunity) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.AsNumber != nil {
		if *obj.obj.AsNumber < 0 || *obj.obj.AsNumber > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= ResultBgpCommunity.AsNumber <= 65535 but Got %d", *obj.obj.AsNumber))
		}

	}

	if obj.obj.AsCustom != nil {
		if *obj.obj.AsCustom < 0 || *obj.obj.AsCustom > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= ResultBgpCommunity.AsCustom <= 65535 but Got %d", *obj.obj.AsCustom))
		}

	}

}

func (obj *resultBgpCommunity) checkUnique() {

}

func (obj *resultBgpCommunity) checkConstraint() {

}

func (obj *resultBgpAsPath) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Segments) != 0 {

		if setDefault {
			obj.Segments().clearHolderSlice()
			for _, item := range obj.obj.Segments {
				newObj := newResultBgpAsPathSegment(obj.validator)
				newObj.self().obj = item
				obj.Segments().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Segments().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *resultBgpAsPath) checkUnique() {

}

func (obj *resultBgpAsPath) checkConstraint() {

}

func (obj *isisLspFlags) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *isisLspFlags) checkUnique() {

}

func (obj *isisLspFlags) checkConstraint() {

}

func (obj *isisLspTlvs) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.HostnameTlvs) != 0 {

		if setDefault {
			obj.HostnameTlvs().clearHolderSlice()
			for _, item := range obj.obj.HostnameTlvs {
				newObj := newIsisLspHostname(obj.validator)
				newObj.self().obj = item
				obj.HostnameTlvs().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.HostnameTlvs().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.IsReachabilityTlvs) != 0 {

		if setDefault {
			obj.IsReachabilityTlvs().clearHolderSlice()
			for _, item := range obj.obj.IsReachabilityTlvs {
				newObj := newIsisLspIsReachabilityTlv(obj.validator)
				newObj.self().obj = item
				obj.IsReachabilityTlvs().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.IsReachabilityTlvs().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.ExtendedIsReachabilityTlvs) != 0 {

		if setDefault {
			obj.ExtendedIsReachabilityTlvs().clearHolderSlice()
			for _, item := range obj.obj.ExtendedIsReachabilityTlvs {
				newObj := newIsisLspExtendedIsReachabilityTlv(obj.validator)
				newObj.self().obj = item
				obj.ExtendedIsReachabilityTlvs().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.ExtendedIsReachabilityTlvs().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Ipv4InternalReachabilityTlvs) != 0 {

		if setDefault {
			obj.Ipv4InternalReachabilityTlvs().clearHolderSlice()
			for _, item := range obj.obj.Ipv4InternalReachabilityTlvs {
				newObj := newIsisLspIpv4InternalReachabilityTlv(obj.validator)
				newObj.self().obj = item
				obj.Ipv4InternalReachabilityTlvs().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ipv4InternalReachabilityTlvs().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Ipv4ExternalReachabilityTlvs) != 0 {

		if setDefault {
			obj.Ipv4ExternalReachabilityTlvs().clearHolderSlice()
			for _, item := range obj.obj.Ipv4ExternalReachabilityTlvs {
				newObj := newIsisLspIpv4ExternalReachabilityTlv(obj.validator)
				newObj.self().obj = item
				obj.Ipv4ExternalReachabilityTlvs().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ipv4ExternalReachabilityTlvs().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.ExtendedIpv4ReachabilityTlvs) != 0 {

		if setDefault {
			obj.ExtendedIpv4ReachabilityTlvs().clearHolderSlice()
			for _, item := range obj.obj.ExtendedIpv4ReachabilityTlvs {
				newObj := newIsisLspExtendedIpv4ReachabilityTlv(obj.validator)
				newObj.self().obj = item
				obj.ExtendedIpv4ReachabilityTlvs().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.ExtendedIpv4ReachabilityTlvs().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Ipv6ReachabilityTlvs) != 0 {

		if setDefault {
			obj.Ipv6ReachabilityTlvs().clearHolderSlice()
			for _, item := range obj.obj.Ipv6ReachabilityTlvs {
				newObj := newIsisLspIpv6ReachabilityTlv(obj.validator)
				newObj.self().obj = item
				obj.Ipv6ReachabilityTlvs().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ipv6ReachabilityTlvs().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *isisLspTlvs) checkUnique() {

}

func (obj *isisLspTlvs) checkConstraint() {

}

func (obj *linkStatepriorityBandwidths) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Pb0 != nil {
		if *obj.obj.Pb0 < 0 || *obj.obj.Pb0 > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LinkStatepriorityBandwidths.Pb0 <= 4294967295 but Got %d", *obj.obj.Pb0))
		}

	}

	if obj.obj.Pb1 != nil {
		if *obj.obj.Pb1 < 0 || *obj.obj.Pb1 > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LinkStatepriorityBandwidths.Pb1 <= 4294967295 but Got %d", *obj.obj.Pb1))
		}

	}

	if obj.obj.Pb2 != nil {
		if *obj.obj.Pb2 < 0 || *obj.obj.Pb2 > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LinkStatepriorityBandwidths.Pb2 <= 4294967295 but Got %d", *obj.obj.Pb2))
		}

	}

	if obj.obj.Pb3 != nil {
		if *obj.obj.Pb3 < 0 || *obj.obj.Pb3 > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LinkStatepriorityBandwidths.Pb3 <= 4294967295 but Got %d", *obj.obj.Pb3))
		}

	}

	if obj.obj.Pb4 != nil {
		if *obj.obj.Pb4 < 0 || *obj.obj.Pb4 > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LinkStatepriorityBandwidths.Pb4 <= 4294967295 but Got %d", *obj.obj.Pb4))
		}

	}

	if obj.obj.Pb5 != nil {
		if *obj.obj.Pb5 < 0 || *obj.obj.Pb5 > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LinkStatepriorityBandwidths.Pb5 <= 4294967295 but Got %d", *obj.obj.Pb5))
		}

	}

	if obj.obj.Pb6 != nil {
		if *obj.obj.Pb6 < 0 || *obj.obj.Pb6 > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LinkStatepriorityBandwidths.Pb6 <= 4294967295 but Got %d", *obj.obj.Pb6))
		}

	}

	if obj.obj.Pb7 != nil {
		if *obj.obj.Pb7 < 0 || *obj.obj.Pb7 > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= LinkStatepriorityBandwidths.Pb7 <= 4294967295 but Got %d", *obj.obj.Pb7))
		}

	}

}

func (obj *linkStatepriorityBandwidths) checkUnique() {

}

func (obj *linkStatepriorityBandwidths) checkConstraint() {

}

func (obj *bgpV4EthernetSegment) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.DfElection != nil {

		obj.DfElection().validateObj(setDefault)
	}

	if len(obj.obj.Evis) != 0 {

		if setDefault {
			obj.Evis().clearHolderSlice()
			for _, item := range obj.obj.Evis {
				newObj := newBgpV4EvpnEvis(obj.validator)
				newObj.self().obj = item
				obj.Evis().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Evis().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.Esi != nil {

		err := obj.validateHex(obj.Esi())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpV4EthernetSegment.Esi"))
		}

	}

	if obj.obj.EsiLabel != nil {
		if *obj.obj.EsiLabel < 0 || *obj.obj.EsiLabel > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpV4EthernetSegment.EsiLabel <= 16777215 but Got %d", *obj.obj.EsiLabel))
		}

	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(setDefault)
	}

	if len(obj.obj.Communities) != 0 {

		if setDefault {
			obj.Communities().clearHolderSlice()
			for _, item := range obj.obj.Communities {
				newObj := newBgpCommunity(obj.validator)
				newObj.self().obj = item
				obj.Communities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Communities().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.ExtCommunities) != 0 {

		if setDefault {
			obj.ExtCommunities().clearHolderSlice()
			for _, item := range obj.obj.ExtCommunities {
				newObj := newBgpExtCommunity(obj.validator)
				newObj.self().obj = item
				obj.ExtCommunities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.ExtCommunities().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(setDefault)
	}

}

func (obj *bgpV4EthernetSegment) checkUnique() {

}

func (obj *bgpV4EthernetSegment) checkConstraint() {

}

func (obj *bgpAdvanced) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *bgpAdvanced) checkUnique() {

}

func (obj *bgpAdvanced) checkConstraint() {

}

func (obj *bgpCapability) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *bgpCapability) checkUnique() {

}

func (obj *bgpCapability) checkConstraint() {

}

func (obj *bgpLearnedInformationFilter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *bgpLearnedInformationFilter) checkUnique() {

}

func (obj *bgpLearnedInformationFilter) checkConstraint() {

}

func (obj *bgpV4RouteRange) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	if len(obj.obj.Addresses) != 0 {

		if setDefault {
			obj.Addresses().clearHolderSlice()
			for _, item := range obj.obj.Addresses {
				newObj := newV4RouteAddress(obj.validator)
				newObj.self().obj = item
				obj.Addresses().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Addresses().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.NextHopIpv4Address != nil {

		err := obj.validateIpv4(obj.NextHopIpv4Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpV4RouteRange.NextHopIpv4Address"))
		}

	}

	if obj.obj.NextHopIpv6Address != nil {

		err := obj.validateIpv6(obj.NextHopIpv6Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpV4RouteRange.NextHopIpv6Address"))
		}

	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(setDefault)
	}

	if len(obj.obj.Communities) != 0 {

		if setDefault {
			obj.Communities().clearHolderSlice()
			for _, item := range obj.obj.Communities {
				newObj := newBgpCommunity(obj.validator)
				newObj.self().obj = item
				obj.Communities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Communities().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(setDefault)
	}

	if obj.obj.AddPath != nil {

		obj.AddPath().validateObj(setDefault)
	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface BgpV4RouteRange")
	}
}

func (obj *bgpV4RouteRange) checkUnique() {
	if !obj.isUnique("bgpV4RouteRange", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *bgpV4RouteRange) checkConstraint() {

}

func (obj *bgpV6RouteRange) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	if len(obj.obj.Addresses) != 0 {

		if setDefault {
			obj.Addresses().clearHolderSlice()
			for _, item := range obj.obj.Addresses {
				newObj := newV6RouteAddress(obj.validator)
				newObj.self().obj = item
				obj.Addresses().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Addresses().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.NextHopIpv4Address != nil {

		err := obj.validateIpv4(obj.NextHopIpv4Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpV6RouteRange.NextHopIpv4Address"))
		}

	}

	if obj.obj.NextHopIpv6Address != nil {

		err := obj.validateIpv6(obj.NextHopIpv6Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpV6RouteRange.NextHopIpv6Address"))
		}

	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(setDefault)
	}

	if len(obj.obj.Communities) != 0 {

		if setDefault {
			obj.Communities().clearHolderSlice()
			for _, item := range obj.obj.Communities {
				newObj := newBgpCommunity(obj.validator)
				newObj.self().obj = item
				obj.Communities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Communities().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(setDefault)
	}

	if obj.obj.AddPath != nil {

		obj.AddPath().validateObj(setDefault)
	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface BgpV6RouteRange")
	}
}

func (obj *bgpV6RouteRange) checkUnique() {
	if !obj.isUnique("bgpV6RouteRange", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *bgpV6RouteRange) checkConstraint() {

}

func (obj *bgpSrteV4Policy) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	if obj.obj.Distinguisher != nil {
		if *obj.obj.Distinguisher < 0 || *obj.obj.Distinguisher > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteV4Policy.Distinguisher <= 4294967295 but Got %d", *obj.obj.Distinguisher))
		}

	}

	if obj.obj.Color != nil {
		if *obj.obj.Color < 0 || *obj.obj.Color > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteV4Policy.Color <= 4294967295 but Got %d", *obj.obj.Color))
		}

	}

	// Ipv4Endpoint is required
	if obj.obj.Ipv4Endpoint == "" {
		obj.errors = append(obj.errors, "Ipv4Endpoint is required field on interface BgpSrteV4Policy")
	}
	if obj.obj.Ipv4Endpoint != "" {

		err := obj.validateIpv4(obj.Ipv4Endpoint())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteV4Policy.Ipv4Endpoint"))
		}

	}

	if obj.obj.NextHopIpv4Address != nil {

		err := obj.validateIpv4(obj.NextHopIpv4Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteV4Policy.NextHopIpv4Address"))
		}

	}

	if obj.obj.NextHopIpv6Address != nil {

		err := obj.validateIpv6(obj.NextHopIpv6Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteV4Policy.NextHopIpv6Address"))
		}

	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(setDefault)
	}

	if obj.obj.AddPath != nil {

		obj.AddPath().validateObj(setDefault)
	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(setDefault)
	}

	if len(obj.obj.Communities) != 0 {

		if setDefault {
			obj.Communities().clearHolderSlice()
			for _, item := range obj.obj.Communities {
				newObj := newBgpCommunity(obj.validator)
				newObj.self().obj = item
				obj.Communities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Communities().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.ExtCommunities) != 0 {

		if setDefault {
			obj.ExtCommunities().clearHolderSlice()
			for _, item := range obj.obj.ExtCommunities {
				newObj := newBgpExtCommunity(obj.validator)
				newObj.self().obj = item
				obj.ExtCommunities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.ExtCommunities().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.TunnelTlvs) != 0 {

		if setDefault {
			obj.TunnelTlvs().clearHolderSlice()
			for _, item := range obj.obj.TunnelTlvs {
				newObj := newBgpSrteV4TunnelTlv(obj.validator)
				newObj.self().obj = item
				obj.TunnelTlvs().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.TunnelTlvs().Items() {
			item.validateObj(setDefault)
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface BgpSrteV4Policy")
	}
}

func (obj *bgpSrteV4Policy) checkUnique() {
	if !obj.isUnique("bgpSrteV4Policy", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *bgpSrteV4Policy) checkConstraint() {

}

func (obj *bgpSrteV6Policy) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	if obj.obj.Distinguisher != nil {
		if *obj.obj.Distinguisher < 0 || *obj.obj.Distinguisher > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteV6Policy.Distinguisher <= 4294967295 but Got %d", *obj.obj.Distinguisher))
		}

	}

	if obj.obj.Color != nil {
		if *obj.obj.Color < 0 || *obj.obj.Color > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteV6Policy.Color <= 4294967295 but Got %d", *obj.obj.Color))
		}

	}

	// Ipv6Endpoint is required
	if obj.obj.Ipv6Endpoint == "" {
		obj.errors = append(obj.errors, "Ipv6Endpoint is required field on interface BgpSrteV6Policy")
	}
	if obj.obj.Ipv6Endpoint != "" {

		err := obj.validateIpv6(obj.Ipv6Endpoint())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteV6Policy.Ipv6Endpoint"))
		}

	}

	if obj.obj.NextHopIpv4Address != nil {

		err := obj.validateIpv4(obj.NextHopIpv4Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteV6Policy.NextHopIpv4Address"))
		}

	}

	if obj.obj.NextHopIpv6Address != nil {

		err := obj.validateIpv6(obj.NextHopIpv6Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteV6Policy.NextHopIpv6Address"))
		}

	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(setDefault)
	}

	if obj.obj.AddPath != nil {

		obj.AddPath().validateObj(setDefault)
	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(setDefault)
	}

	if len(obj.obj.Communities) != 0 {

		if setDefault {
			obj.Communities().clearHolderSlice()
			for _, item := range obj.obj.Communities {
				newObj := newBgpCommunity(obj.validator)
				newObj.self().obj = item
				obj.Communities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Communities().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Extcommunities) != 0 {

		if setDefault {
			obj.Extcommunities().clearHolderSlice()
			for _, item := range obj.obj.Extcommunities {
				newObj := newBgpExtCommunity(obj.validator)
				newObj.self().obj = item
				obj.Extcommunities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Extcommunities().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.TunnelTlvs) != 0 {

		if setDefault {
			obj.TunnelTlvs().clearHolderSlice()
			for _, item := range obj.obj.TunnelTlvs {
				newObj := newBgpSrteV6TunnelTlv(obj.validator)
				newObj.self().obj = item
				obj.TunnelTlvs().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.TunnelTlvs().Items() {
			item.validateObj(setDefault)
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface BgpSrteV6Policy")
	}
}

func (obj *bgpSrteV6Policy) checkUnique() {
	if !obj.isUnique("bgpSrteV6Policy", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *bgpSrteV6Policy) checkConstraint() {

}

func (obj *bgpV6SegmentRouting) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *bgpV6SegmentRouting) checkUnique() {

}

func (obj *bgpV6SegmentRouting) checkConstraint() {

}

func (obj *bgpV6EthernetSegment) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.DfElection != nil {

		obj.DfElection().validateObj(setDefault)
	}

	if len(obj.obj.Evis) != 0 {

		if setDefault {
			obj.Evis().clearHolderSlice()
			for _, item := range obj.obj.Evis {
				newObj := newBgpV6EvpnEvis(obj.validator)
				newObj.self().obj = item
				obj.Evis().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Evis().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.Esi != nil {

		err := obj.validateHex(obj.Esi())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpV6EthernetSegment.Esi"))
		}

	}

	if obj.obj.EsiLabel != nil {
		if *obj.obj.EsiLabel < 0 || *obj.obj.EsiLabel > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpV6EthernetSegment.EsiLabel <= 16777215 but Got %d", *obj.obj.EsiLabel))
		}

	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(setDefault)
	}

	if len(obj.obj.Communities) != 0 {

		if setDefault {
			obj.Communities().clearHolderSlice()
			for _, item := range obj.obj.Communities {
				newObj := newBgpCommunity(obj.validator)
				newObj.self().obj = item
				obj.Communities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Communities().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.ExtCommunities) != 0 {

		if setDefault {
			obj.ExtCommunities().clearHolderSlice()
			for _, item := range obj.obj.ExtCommunities {
				newObj := newBgpExtCommunity(obj.validator)
				newObj.self().obj = item
				obj.ExtCommunities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.ExtCommunities().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(setDefault)
	}

}

func (obj *bgpV6EthernetSegment) checkUnique() {

}

func (obj *bgpV6EthernetSegment) checkConstraint() {

}

func (obj *vxlanV4TunnelDestinationIPModeUnicast) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Vteps) != 0 {

		if setDefault {
			obj.Vteps().clearHolderSlice()
			for _, item := range obj.obj.Vteps {
				newObj := newVxlanV4TunnelDestinationIPModeUnicastVtep(obj.validator)
				newObj.self().obj = item
				obj.Vteps().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Vteps().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *vxlanV4TunnelDestinationIPModeUnicast) checkUnique() {

}

func (obj *vxlanV4TunnelDestinationIPModeUnicast) checkConstraint() {

}

func (obj *vxlanV4TunnelDestinationIPModeMulticast) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv4(obj.Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on VxlanV4TunnelDestinationIPModeMulticast.Address"))
		}

	}

}

func (obj *vxlanV4TunnelDestinationIPModeMulticast) checkUnique() {

}

func (obj *vxlanV4TunnelDestinationIPModeMulticast) checkConstraint() {

}

func (obj *vxlanV6TunnelDestinationIPModeUnicast) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Vteps) != 0 {

		if setDefault {
			obj.Vteps().clearHolderSlice()
			for _, item := range obj.obj.Vteps {
				newObj := newVxlanV6TunnelDestinationIPModeUnicastVtep(obj.validator)
				newObj.self().obj = item
				obj.Vteps().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Vteps().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *vxlanV6TunnelDestinationIPModeUnicast) checkUnique() {

}

func (obj *vxlanV6TunnelDestinationIPModeUnicast) checkConstraint() {

}

func (obj *vxlanV6TunnelDestinationIPModeMulticast) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv6(obj.Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on VxlanV6TunnelDestinationIPModeMulticast.Address"))
		}

	}

}

func (obj *vxlanV6TunnelDestinationIPModeMulticast) checkUnique() {

}

func (obj *vxlanV6TunnelDestinationIPModeMulticast) checkConstraint() {

}

func (obj *patternFlowEthernetDstCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetDstCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetDstCounter.Step"))
		}

	}

}

func (obj *patternFlowEthernetDstCounter) checkUnique() {

}

func (obj *patternFlowEthernetDstCounter) checkConstraint() {

}

func (obj *patternFlowEthernetSrcCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetSrcCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetSrcCounter.Step"))
		}

	}

}

func (obj *patternFlowEthernetSrcCounter) checkUnique() {

}

func (obj *patternFlowEthernetSrcCounter) checkConstraint() {

}

func (obj *patternFlowEthernetEtherTypeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetEtherTypeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetEtherTypeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowEthernetEtherTypeCounter) checkUnique() {

}

func (obj *patternFlowEthernetEtherTypeCounter) checkConstraint() {

}

func (obj *patternFlowEthernetPfcQueueCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetPfcQueueCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetPfcQueueCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowEthernetPfcQueueCounter) checkUnique() {

}

func (obj *patternFlowEthernetPfcQueueCounter) checkConstraint() {

}

func (obj *patternFlowVlanPriorityCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVlanPriorityCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVlanPriorityCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowVlanPriorityCounter) checkUnique() {

}

func (obj *patternFlowVlanPriorityCounter) checkConstraint() {

}

func (obj *patternFlowVlanCfiCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVlanCfiCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVlanCfiCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowVlanCfiCounter) checkUnique() {

}

func (obj *patternFlowVlanCfiCounter) checkConstraint() {

}

func (obj *patternFlowVlanIdCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 4095 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVlanIdCounter.Start <= 4095 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 4095 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVlanIdCounter.Step <= 4095 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowVlanIdCounter) checkUnique() {

}

func (obj *patternFlowVlanIdCounter) checkConstraint() {

}

func (obj *patternFlowVlanTpidCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVlanTpidCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVlanTpidCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowVlanTpidCounter) checkUnique() {

}

func (obj *patternFlowVlanTpidCounter) checkConstraint() {

}

func (obj *patternFlowVxlanFlagsCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVxlanFlagsCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVxlanFlagsCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowVxlanFlagsCounter) checkUnique() {

}

func (obj *patternFlowVxlanFlagsCounter) checkConstraint() {

}

func (obj *patternFlowVxlanReserved0Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved0Counter.Start <= 16777215 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved0Counter.Step <= 16777215 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowVxlanReserved0Counter) checkUnique() {

}

func (obj *patternFlowVxlanReserved0Counter) checkConstraint() {

}

func (obj *patternFlowVxlanVniCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVxlanVniCounter.Start <= 16777215 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVxlanVniCounter.Step <= 16777215 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowVxlanVniCounter) checkUnique() {

}

func (obj *patternFlowVxlanVniCounter) checkConstraint() {

}

func (obj *patternFlowVxlanReserved1Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved1Counter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved1Counter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowVxlanReserved1Counter) checkUnique() {

}

func (obj *patternFlowVxlanReserved1Counter) checkConstraint() {

}

func (obj *patternFlowIpv4VersionCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4VersionCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4VersionCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4VersionCounter) checkUnique() {

}

func (obj *patternFlowIpv4VersionCounter) checkConstraint() {

}

func (obj *patternFlowIpv4HeaderLengthCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4HeaderLengthCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4HeaderLengthCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4HeaderLengthCounter) checkUnique() {

}

func (obj *patternFlowIpv4HeaderLengthCounter) checkConstraint() {

}

func (obj *patternFlowIpv4PriorityRaw) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4PriorityRaw.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4PriorityRaw.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4PriorityRaw) checkUnique() {

}

func (obj *patternFlowIpv4PriorityRaw) checkConstraint() {

}

func (obj *flowIpv4Tos) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Precedence != nil {

		obj.Precedence().validateObj(setDefault)
	}

	if obj.obj.Delay != nil {

		obj.Delay().validateObj(setDefault)
	}

	if obj.obj.Throughput != nil {

		obj.Throughput().validateObj(setDefault)
	}

	if obj.obj.Reliability != nil {

		obj.Reliability().validateObj(setDefault)
	}

	if obj.obj.Monetary != nil {

		obj.Monetary().validateObj(setDefault)
	}

	if obj.obj.Unused != nil {

		obj.Unused().validateObj(setDefault)
	}

}

func (obj *flowIpv4Tos) checkUnique() {

}

func (obj *flowIpv4Tos) checkConstraint() {

}

func (obj *flowIpv4Dscp) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Phb != nil {

		obj.Phb().validateObj(setDefault)
	}

	if obj.obj.Ecn != nil {

		obj.Ecn().validateObj(setDefault)
	}

}

func (obj *flowIpv4Dscp) checkUnique() {

}

func (obj *flowIpv4Dscp) checkConstraint() {

}

func (obj *patternFlowIpv4TotalLengthCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TotalLengthCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TotalLengthCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4TotalLengthCounter) checkUnique() {

}

func (obj *patternFlowIpv4TotalLengthCounter) checkConstraint() {

}

func (obj *patternFlowIpv4IdentificationCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4IdentificationCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4IdentificationCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4IdentificationCounter) checkUnique() {

}

func (obj *patternFlowIpv4IdentificationCounter) checkConstraint() {

}

func (obj *patternFlowIpv4ReservedCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4ReservedCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4ReservedCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4ReservedCounter) checkUnique() {

}

func (obj *patternFlowIpv4ReservedCounter) checkConstraint() {

}

func (obj *patternFlowIpv4DontFragmentCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4DontFragmentCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4DontFragmentCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4DontFragmentCounter) checkUnique() {

}

func (obj *patternFlowIpv4DontFragmentCounter) checkConstraint() {

}

func (obj *patternFlowIpv4MoreFragmentsCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4MoreFragmentsCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4MoreFragmentsCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4MoreFragmentsCounter) checkUnique() {

}

func (obj *patternFlowIpv4MoreFragmentsCounter) checkConstraint() {

}

func (obj *patternFlowIpv4FragmentOffsetCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 31 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4FragmentOffsetCounter.Start <= 31 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 31 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4FragmentOffsetCounter.Step <= 31 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4FragmentOffsetCounter) checkUnique() {

}

func (obj *patternFlowIpv4FragmentOffsetCounter) checkConstraint() {

}

func (obj *patternFlowIpv4TimeToLiveCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TimeToLiveCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TimeToLiveCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4TimeToLiveCounter) checkUnique() {

}

func (obj *patternFlowIpv4TimeToLiveCounter) checkConstraint() {

}

func (obj *patternFlowIpv4ProtocolCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4ProtocolCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4ProtocolCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4ProtocolCounter) checkUnique() {

}

func (obj *patternFlowIpv4ProtocolCounter) checkConstraint() {

}

func (obj *patternFlowIpv4SrcCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4SrcCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4SrcCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv4SrcCounter) checkUnique() {

}

func (obj *patternFlowIpv4SrcCounter) checkConstraint() {

}

func (obj *patternFlowIpv4DstCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4DstCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4DstCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv4DstCounter) checkUnique() {

}

func (obj *patternFlowIpv4DstCounter) checkConstraint() {

}

func (obj *patternFlowIpv6VersionCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6VersionCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6VersionCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv6VersionCounter) checkUnique() {

}

func (obj *patternFlowIpv6VersionCounter) checkConstraint() {

}

func (obj *patternFlowIpv6TrafficClassCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6TrafficClassCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6TrafficClassCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv6TrafficClassCounter) checkUnique() {

}

func (obj *patternFlowIpv6TrafficClassCounter) checkConstraint() {

}

func (obj *patternFlowIpv6FlowLabelCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1048575 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6FlowLabelCounter.Start <= 1048575 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1048575 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6FlowLabelCounter.Step <= 1048575 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv6FlowLabelCounter) checkUnique() {

}

func (obj *patternFlowIpv6FlowLabelCounter) checkConstraint() {

}

func (obj *patternFlowIpv6PayloadLengthCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6PayloadLengthCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6PayloadLengthCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv6PayloadLengthCounter) checkUnique() {

}

func (obj *patternFlowIpv6PayloadLengthCounter) checkConstraint() {

}

func (obj *patternFlowIpv6NextHeaderCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6NextHeaderCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6NextHeaderCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv6NextHeaderCounter) checkUnique() {

}

func (obj *patternFlowIpv6NextHeaderCounter) checkConstraint() {

}

func (obj *patternFlowIpv6HopLimitCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6HopLimitCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv6HopLimitCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv6HopLimitCounter) checkUnique() {

}

func (obj *patternFlowIpv6HopLimitCounter) checkConstraint() {

}

func (obj *patternFlowIpv6SrcCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SrcCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SrcCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv6SrcCounter) checkUnique() {

}

func (obj *patternFlowIpv6SrcCounter) checkConstraint() {

}

func (obj *patternFlowIpv6DstCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6DstCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6DstCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv6DstCounter) checkUnique() {

}

func (obj *patternFlowIpv6DstCounter) checkConstraint() {

}

func (obj *patternFlowPfcPauseDstCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseDstCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseDstCounter.Step"))
		}

	}

}

func (obj *patternFlowPfcPauseDstCounter) checkUnique() {

}

func (obj *patternFlowPfcPauseDstCounter) checkConstraint() {

}

func (obj *patternFlowPfcPauseSrcCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseSrcCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseSrcCounter.Step"))
		}

	}

}

func (obj *patternFlowPfcPauseSrcCounter) checkUnique() {

}

func (obj *patternFlowPfcPauseSrcCounter) checkConstraint() {

}

func (obj *patternFlowPfcPauseEtherTypeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseEtherTypeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseEtherTypeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowPfcPauseEtherTypeCounter) checkUnique() {

}

func (obj *patternFlowPfcPauseEtherTypeCounter) checkConstraint() {

}

func (obj *patternFlowPfcPauseControlOpCodeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseControlOpCodeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseControlOpCodeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowPfcPauseControlOpCodeCounter) checkUnique() {

}

func (obj *patternFlowPfcPauseControlOpCodeCounter) checkConstraint() {

}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseClassEnableVectorCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseClassEnableVectorCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) checkUnique() {

}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass0Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass0Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass0Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass0Counter) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass0Counter) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass1Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass1Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass1Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass1Counter) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass1Counter) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass2Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass2Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass2Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass2Counter) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass2Counter) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass3Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass3Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass3Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass3Counter) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass3Counter) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass4Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass4Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass4Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass4Counter) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass4Counter) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass5Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass5Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass5Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass5Counter) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass5Counter) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass6Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass6Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass6Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass6Counter) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass6Counter) checkConstraint() {

}

func (obj *patternFlowPfcPausePauseClass7Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass7Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass7Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass7Counter) checkUnique() {

}

func (obj *patternFlowPfcPausePauseClass7Counter) checkConstraint() {

}

func (obj *patternFlowEthernetPauseDstCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseDstCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseDstCounter.Step"))
		}

	}

}

func (obj *patternFlowEthernetPauseDstCounter) checkUnique() {

}

func (obj *patternFlowEthernetPauseDstCounter) checkConstraint() {

}

func (obj *patternFlowEthernetPauseSrcCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseSrcCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseSrcCounter.Step"))
		}

	}

}

func (obj *patternFlowEthernetPauseSrcCounter) checkUnique() {

}

func (obj *patternFlowEthernetPauseSrcCounter) checkConstraint() {

}

func (obj *patternFlowEthernetPauseEtherTypeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseEtherTypeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseEtherTypeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowEthernetPauseEtherTypeCounter) checkUnique() {

}

func (obj *patternFlowEthernetPauseEtherTypeCounter) checkConstraint() {

}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseControlOpCodeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseControlOpCodeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) checkUnique() {

}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) checkConstraint() {

}

func (obj *patternFlowEthernetPauseTimeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseTimeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseTimeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowEthernetPauseTimeCounter) checkUnique() {

}

func (obj *patternFlowEthernetPauseTimeCounter) checkConstraint() {

}

func (obj *patternFlowTcpSrcPortCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpSrcPortCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpSrcPortCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowTcpSrcPortCounter) checkUnique() {

}

func (obj *patternFlowTcpSrcPortCounter) checkConstraint() {

}

func (obj *patternFlowTcpDstPortCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpDstPortCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpDstPortCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowTcpDstPortCounter) checkUnique() {

}

func (obj *patternFlowTcpDstPortCounter) checkConstraint() {

}

func (obj *patternFlowTcpSeqNumCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpSeqNumCounter.Start <= 4294967295 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpSeqNumCounter.Step <= 4294967295 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowTcpSeqNumCounter) checkUnique() {

}

func (obj *patternFlowTcpSeqNumCounter) checkConstraint() {

}

func (obj *patternFlowTcpAckNumCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpAckNumCounter.Start <= 4294967295 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpAckNumCounter.Step <= 4294967295 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowTcpAckNumCounter) checkUnique() {

}

func (obj *patternFlowTcpAckNumCounter) checkConstraint() {

}

func (obj *patternFlowTcpDataOffsetCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpDataOffsetCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpDataOffsetCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowTcpDataOffsetCounter) checkUnique() {

}

func (obj *patternFlowTcpDataOffsetCounter) checkConstraint() {

}

func (obj *patternFlowTcpEcnNsCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnNsCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnNsCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowTcpEcnNsCounter) checkUnique() {

}

func (obj *patternFlowTcpEcnNsCounter) checkConstraint() {

}

func (obj *patternFlowTcpEcnCwrCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnCwrCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnCwrCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowTcpEcnCwrCounter) checkUnique() {

}

func (obj *patternFlowTcpEcnCwrCounter) checkConstraint() {

}

func (obj *patternFlowTcpEcnEchoCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnEchoCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnEchoCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowTcpEcnEchoCounter) checkUnique() {

}

func (obj *patternFlowTcpEcnEchoCounter) checkConstraint() {

}

func (obj *patternFlowTcpCtlUrgCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlUrgCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlUrgCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowTcpCtlUrgCounter) checkUnique() {

}

func (obj *patternFlowTcpCtlUrgCounter) checkConstraint() {

}

func (obj *patternFlowTcpCtlAckCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlAckCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlAckCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowTcpCtlAckCounter) checkUnique() {

}

func (obj *patternFlowTcpCtlAckCounter) checkConstraint() {

}

func (obj *patternFlowTcpCtlPshCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlPshCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlPshCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowTcpCtlPshCounter) checkUnique() {

}

func (obj *patternFlowTcpCtlPshCounter) checkConstraint() {

}

func (obj *patternFlowTcpCtlRstCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlRstCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlRstCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowTcpCtlRstCounter) checkUnique() {

}

func (obj *patternFlowTcpCtlRstCounter) checkConstraint() {

}

func (obj *patternFlowTcpCtlSynCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlSynCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlSynCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowTcpCtlSynCounter) checkUnique() {

}

func (obj *patternFlowTcpCtlSynCounter) checkConstraint() {

}

func (obj *patternFlowTcpCtlFinCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlFinCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlFinCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowTcpCtlFinCounter) checkUnique() {

}

func (obj *patternFlowTcpCtlFinCounter) checkConstraint() {

}

func (obj *patternFlowTcpWindowCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpWindowCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowTcpWindowCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowTcpWindowCounter) checkUnique() {

}

func (obj *patternFlowTcpWindowCounter) checkConstraint() {

}

func (obj *patternFlowUdpSrcPortCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowUdpSrcPortCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowUdpSrcPortCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowUdpSrcPortCounter) checkUnique() {

}

func (obj *patternFlowUdpSrcPortCounter) checkConstraint() {

}

func (obj *patternFlowUdpDstPortCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowUdpDstPortCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowUdpDstPortCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowUdpDstPortCounter) checkUnique() {

}

func (obj *patternFlowUdpDstPortCounter) checkConstraint() {

}

func (obj *patternFlowUdpLengthCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowUdpLengthCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowUdpLengthCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowUdpLengthCounter) checkUnique() {

}

func (obj *patternFlowUdpLengthCounter) checkConstraint() {

}

func (obj *patternFlowGreChecksumPresentCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreChecksumPresentCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreChecksumPresentCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGreChecksumPresentCounter) checkUnique() {

}

func (obj *patternFlowGreChecksumPresentCounter) checkConstraint() {

}

func (obj *patternFlowGreReserved0Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 4095 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreReserved0Counter.Start <= 4095 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 4095 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreReserved0Counter.Step <= 4095 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGreReserved0Counter) checkUnique() {

}

func (obj *patternFlowGreReserved0Counter) checkConstraint() {

}

func (obj *patternFlowGreVersionCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreVersionCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreVersionCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGreVersionCounter) checkUnique() {

}

func (obj *patternFlowGreVersionCounter) checkConstraint() {

}

func (obj *patternFlowGreProtocolCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreProtocolCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreProtocolCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGreProtocolCounter) checkUnique() {

}

func (obj *patternFlowGreProtocolCounter) checkConstraint() {

}

func (obj *patternFlowGreReserved1Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreReserved1Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGreReserved1Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGreReserved1Counter) checkUnique() {

}

func (obj *patternFlowGreReserved1Counter) checkConstraint() {

}

func (obj *patternFlowGtpv1VersionCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1VersionCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1VersionCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv1VersionCounter) checkUnique() {

}

func (obj *patternFlowGtpv1VersionCounter) checkConstraint() {

}

func (obj *patternFlowGtpv1ProtocolTypeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1ProtocolTypeCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1ProtocolTypeCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv1ProtocolTypeCounter) checkUnique() {

}

func (obj *patternFlowGtpv1ProtocolTypeCounter) checkConstraint() {

}

func (obj *patternFlowGtpv1ReservedCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1ReservedCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1ReservedCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv1ReservedCounter) checkUnique() {

}

func (obj *patternFlowGtpv1ReservedCounter) checkConstraint() {

}

func (obj *patternFlowGtpv1EFlagCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1EFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1EFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv1EFlagCounter) checkUnique() {

}

func (obj *patternFlowGtpv1EFlagCounter) checkConstraint() {

}

func (obj *patternFlowGtpv1SFlagCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv1SFlagCounter) checkUnique() {

}

func (obj *patternFlowGtpv1SFlagCounter) checkConstraint() {

}

func (obj *patternFlowGtpv1PnFlagCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1PnFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1PnFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv1PnFlagCounter) checkUnique() {

}

func (obj *patternFlowGtpv1PnFlagCounter) checkConstraint() {

}

func (obj *patternFlowGtpv1MessageTypeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv1MessageTypeCounter) checkUnique() {

}

func (obj *patternFlowGtpv1MessageTypeCounter) checkConstraint() {

}

func (obj *patternFlowGtpv1MessageLengthCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageLengthCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageLengthCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv1MessageLengthCounter) checkUnique() {

}

func (obj *patternFlowGtpv1MessageLengthCounter) checkConstraint() {

}

func (obj *patternFlowGtpv1TeidCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1TeidCounter.Start <= 4294967295 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1TeidCounter.Step <= 4294967295 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv1TeidCounter) checkUnique() {

}

func (obj *patternFlowGtpv1TeidCounter) checkConstraint() {

}

func (obj *patternFlowGtpv1SquenceNumberCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SquenceNumberCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SquenceNumberCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv1SquenceNumberCounter) checkUnique() {

}

func (obj *patternFlowGtpv1SquenceNumberCounter) checkConstraint() {

}

func (obj *patternFlowGtpv1NPduNumberCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NPduNumberCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NPduNumberCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv1NPduNumberCounter) checkUnique() {

}

func (obj *patternFlowGtpv1NPduNumberCounter) checkConstraint() {

}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NextExtensionHeaderTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NextExtensionHeaderTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) checkUnique() {

}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) checkConstraint() {

}

func (obj *patternFlowGtpExtensionExtensionLength) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionExtensionLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpExtensionExtensionLength.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpExtensionExtensionLength) checkUnique() {

}

func (obj *patternFlowGtpExtensionExtensionLength) checkConstraint() {

}

func (obj *patternFlowGtpExtensionContents) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 281474976710655 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionContents.Value <= 281474976710655 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 281474976710655 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpExtensionContents.Values <= 281474976710655 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpExtensionContents) checkUnique() {

}

func (obj *patternFlowGtpExtensionContents) checkConstraint() {

}

func (obj *patternFlowGtpExtensionNextExtensionHeader) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionNextExtensionHeader.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowGtpExtensionNextExtensionHeader.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowGtpExtensionNextExtensionHeader) checkUnique() {

}

func (obj *patternFlowGtpExtensionNextExtensionHeader) checkConstraint() {

}

func (obj *patternFlowGtpv2VersionCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2VersionCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2VersionCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv2VersionCounter) checkUnique() {

}

func (obj *patternFlowGtpv2VersionCounter) checkConstraint() {

}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2PiggybackingFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2PiggybackingFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) checkUnique() {

}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) checkConstraint() {

}

func (obj *patternFlowGtpv2TeidFlagCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2TeidFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2TeidFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv2TeidFlagCounter) checkUnique() {

}

func (obj *patternFlowGtpv2TeidFlagCounter) checkConstraint() {

}

func (obj *patternFlowGtpv2Spare1Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2Spare1Counter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2Spare1Counter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv2Spare1Counter) checkUnique() {

}

func (obj *patternFlowGtpv2Spare1Counter) checkConstraint() {

}

func (obj *patternFlowGtpv2MessageTypeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv2MessageTypeCounter) checkUnique() {

}

func (obj *patternFlowGtpv2MessageTypeCounter) checkConstraint() {

}

func (obj *patternFlowGtpv2MessageLengthCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageLengthCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageLengthCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv2MessageLengthCounter) checkUnique() {

}

func (obj *patternFlowGtpv2MessageLengthCounter) checkConstraint() {

}

func (obj *patternFlowGtpv2TeidCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2TeidCounter.Start <= 4294967295 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2TeidCounter.Step <= 4294967295 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv2TeidCounter) checkUnique() {

}

func (obj *patternFlowGtpv2TeidCounter) checkConstraint() {

}

func (obj *patternFlowGtpv2SequenceNumberCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2SequenceNumberCounter.Start <= 16777215 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2SequenceNumberCounter.Step <= 16777215 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv2SequenceNumberCounter) checkUnique() {

}

func (obj *patternFlowGtpv2SequenceNumberCounter) checkConstraint() {

}

func (obj *patternFlowGtpv2Spare2Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2Spare2Counter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpv2Spare2Counter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpv2Spare2Counter) checkUnique() {

}

func (obj *patternFlowGtpv2Spare2Counter) checkConstraint() {

}

func (obj *patternFlowArpHardwareTypeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareTypeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareTypeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowArpHardwareTypeCounter) checkUnique() {

}

func (obj *patternFlowArpHardwareTypeCounter) checkConstraint() {

}

func (obj *patternFlowArpProtocolTypeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolTypeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolTypeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowArpProtocolTypeCounter) checkUnique() {

}

func (obj *patternFlowArpProtocolTypeCounter) checkConstraint() {

}

func (obj *patternFlowArpHardwareLengthCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowArpHardwareLengthCounter) checkUnique() {

}

func (obj *patternFlowArpHardwareLengthCounter) checkConstraint() {

}

func (obj *patternFlowArpProtocolLengthCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowArpProtocolLengthCounter) checkUnique() {

}

func (obj *patternFlowArpProtocolLengthCounter) checkConstraint() {

}

func (obj *patternFlowArpOperationCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowArpOperationCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowArpOperationCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowArpOperationCounter) checkUnique() {

}

func (obj *patternFlowArpOperationCounter) checkConstraint() {

}

func (obj *patternFlowArpSenderHardwareAddrCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderHardwareAddrCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderHardwareAddrCounter.Step"))
		}

	}

}

func (obj *patternFlowArpSenderHardwareAddrCounter) checkUnique() {

}

func (obj *patternFlowArpSenderHardwareAddrCounter) checkConstraint() {

}

func (obj *patternFlowArpSenderProtocolAddrCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderProtocolAddrCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderProtocolAddrCounter.Step"))
		}

	}

}

func (obj *patternFlowArpSenderProtocolAddrCounter) checkUnique() {

}

func (obj *patternFlowArpSenderProtocolAddrCounter) checkConstraint() {

}

func (obj *patternFlowArpTargetHardwareAddrCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetHardwareAddrCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetHardwareAddrCounter.Step"))
		}

	}

}

func (obj *patternFlowArpTargetHardwareAddrCounter) checkUnique() {

}

func (obj *patternFlowArpTargetHardwareAddrCounter) checkConstraint() {

}

func (obj *patternFlowArpTargetProtocolAddrCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetProtocolAddrCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetProtocolAddrCounter.Step"))
		}

	}

}

func (obj *patternFlowArpTargetProtocolAddrCounter) checkUnique() {

}

func (obj *patternFlowArpTargetProtocolAddrCounter) checkConstraint() {

}

func (obj *patternFlowIcmpEchoType) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIcmpEchoType.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIcmpEchoType) checkUnique() {

}

func (obj *patternFlowIcmpEchoType) checkConstraint() {

}

func (obj *patternFlowIcmpEchoCode) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoCode.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIcmpEchoCode.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIcmpEchoCode) checkUnique() {

}

func (obj *patternFlowIcmpEchoCode) checkConstraint() {

}

func (obj *patternFlowIcmpEchoChecksum) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {
		if *obj.obj.Custom < 0 || *obj.obj.Custom > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowIcmpEchoChecksum) checkUnique() {

}

func (obj *patternFlowIcmpEchoChecksum) checkConstraint() {

}

func (obj *patternFlowIcmpEchoIdentifier) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoIdentifier.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIcmpEchoIdentifier.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIcmpEchoIdentifier) checkUnique() {

}

func (obj *patternFlowIcmpEchoIdentifier) checkConstraint() {

}

func (obj *patternFlowIcmpEchoSequenceNumber) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoSequenceNumber.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIcmpEchoSequenceNumber.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIcmpEchoSequenceNumber) checkUnique() {

}

func (obj *patternFlowIcmpEchoSequenceNumber) checkConstraint() {

}

func (obj *patternFlowIcmpv6EchoType) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIcmpv6EchoType.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIcmpv6EchoType) checkUnique() {

}

func (obj *patternFlowIcmpv6EchoType) checkConstraint() {

}

func (obj *patternFlowIcmpv6EchoCode) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoCode.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIcmpv6EchoCode.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIcmpv6EchoCode) checkUnique() {

}

func (obj *patternFlowIcmpv6EchoCode) checkConstraint() {

}

func (obj *patternFlowIcmpv6EchoIdentifier) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoIdentifier.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIcmpv6EchoIdentifier.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIcmpv6EchoIdentifier) checkUnique() {

}

func (obj *patternFlowIcmpv6EchoIdentifier) checkConstraint() {

}

func (obj *patternFlowIcmpv6EchoSequenceNumber) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoSequenceNumber.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 65535 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIcmpv6EchoSequenceNumber.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIcmpv6EchoSequenceNumber) checkUnique() {

}

func (obj *patternFlowIcmpv6EchoSequenceNumber) checkConstraint() {

}

func (obj *patternFlowIcmpv6EchoChecksum) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {
		if *obj.obj.Custom < 0 || *obj.obj.Custom > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowIcmpv6EchoChecksum) checkUnique() {

}

func (obj *patternFlowIcmpv6EchoChecksum) checkConstraint() {

}

func (obj *patternFlowPppAddressCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPppAddressCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPppAddressCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowPppAddressCounter) checkUnique() {

}

func (obj *patternFlowPppAddressCounter) checkConstraint() {

}

func (obj *patternFlowPppControlCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPppControlCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPppControlCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowPppControlCounter) checkUnique() {

}

func (obj *patternFlowPppControlCounter) checkConstraint() {

}

func (obj *patternFlowPppProtocolTypeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPppProtocolTypeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowPppProtocolTypeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowPppProtocolTypeCounter) checkUnique() {

}

func (obj *patternFlowPppProtocolTypeCounter) checkConstraint() {

}

func (obj *patternFlowIgmpv1VersionCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1VersionCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1VersionCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIgmpv1VersionCounter) checkUnique() {

}

func (obj *patternFlowIgmpv1VersionCounter) checkConstraint() {

}

func (obj *patternFlowIgmpv1TypeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1TypeCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 15 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1TypeCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIgmpv1TypeCounter) checkUnique() {

}

func (obj *patternFlowIgmpv1TypeCounter) checkConstraint() {

}

func (obj *patternFlowIgmpv1UnusedCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1UnusedCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1UnusedCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIgmpv1UnusedCounter) checkUnique() {

}

func (obj *patternFlowIgmpv1UnusedCounter) checkConstraint() {

}

func (obj *patternFlowIgmpv1GroupAddressCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv1GroupAddressCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv1GroupAddressCounter.Step"))
		}

	}

}

func (obj *patternFlowIgmpv1GroupAddressCounter) checkUnique() {

}

func (obj *patternFlowIgmpv1GroupAddressCounter) checkConstraint() {

}

func (obj *patternFlowMplsLabelCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1048575 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowMplsLabelCounter.Start <= 1048575 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1048575 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowMplsLabelCounter.Step <= 1048575 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowMplsLabelCounter) checkUnique() {

}

func (obj *patternFlowMplsLabelCounter) checkConstraint() {

}

func (obj *patternFlowMplsTrafficClassCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowMplsTrafficClassCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowMplsTrafficClassCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowMplsTrafficClassCounter) checkUnique() {

}

func (obj *patternFlowMplsTrafficClassCounter) checkConstraint() {

}

func (obj *patternFlowMplsBottomOfStackCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowMplsBottomOfStackCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowMplsBottomOfStackCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowMplsBottomOfStackCounter) checkUnique() {

}

func (obj *patternFlowMplsBottomOfStackCounter) checkConstraint() {

}

func (obj *patternFlowMplsTimeToLiveCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowMplsTimeToLiveCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowMplsTimeToLiveCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowMplsTimeToLiveCounter) checkUnique() {

}

func (obj *patternFlowMplsTimeToLiveCounter) checkConstraint() {

}

func (obj *resultBgpAsPathSegment) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *resultBgpAsPathSegment) checkUnique() {

}

func (obj *resultBgpAsPathSegment) checkConstraint() {

}

func (obj *isisLspHostname) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *isisLspHostname) checkUnique() {

}

func (obj *isisLspHostname) checkConstraint() {

}

func (obj *isisLspIsReachabilityTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Neighbors) != 0 {

		if setDefault {
			obj.Neighbors().clearHolderSlice()
			for _, item := range obj.obj.Neighbors {
				newObj := newIsisLspneighbor(obj.validator)
				newObj.self().obj = item
				obj.Neighbors().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Neighbors().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *isisLspIsReachabilityTlv) checkUnique() {

}

func (obj *isisLspIsReachabilityTlv) checkConstraint() {

}

func (obj *isisLspExtendedIsReachabilityTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Neighbors) != 0 {

		if setDefault {
			obj.Neighbors().clearHolderSlice()
			for _, item := range obj.obj.Neighbors {
				newObj := newIsisLspneighbor(obj.validator)
				newObj.self().obj = item
				obj.Neighbors().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Neighbors().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *isisLspExtendedIsReachabilityTlv) checkUnique() {

}

func (obj *isisLspExtendedIsReachabilityTlv) checkConstraint() {

}

func (obj *isisLspIpv4InternalReachabilityTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Prefixes) != 0 {

		if setDefault {
			obj.Prefixes().clearHolderSlice()
			for _, item := range obj.obj.Prefixes {
				newObj := newIsisLspV4Prefix(obj.validator)
				newObj.self().obj = item
				obj.Prefixes().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Prefixes().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *isisLspIpv4InternalReachabilityTlv) checkUnique() {

}

func (obj *isisLspIpv4InternalReachabilityTlv) checkConstraint() {

}

func (obj *isisLspIpv4ExternalReachabilityTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Prefixes) != 0 {

		if setDefault {
			obj.Prefixes().clearHolderSlice()
			for _, item := range obj.obj.Prefixes {
				newObj := newIsisLspV4Prefix(obj.validator)
				newObj.self().obj = item
				obj.Prefixes().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Prefixes().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *isisLspIpv4ExternalReachabilityTlv) checkUnique() {

}

func (obj *isisLspIpv4ExternalReachabilityTlv) checkConstraint() {

}

func (obj *isisLspExtendedIpv4ReachabilityTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Prefixes) != 0 {

		if setDefault {
			obj.Prefixes().clearHolderSlice()
			for _, item := range obj.obj.Prefixes {
				newObj := newIsisLspExtendedV4Prefix(obj.validator)
				newObj.self().obj = item
				obj.Prefixes().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Prefixes().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *isisLspExtendedIpv4ReachabilityTlv) checkUnique() {

}

func (obj *isisLspExtendedIpv4ReachabilityTlv) checkConstraint() {

}

func (obj *isisLspIpv6ReachabilityTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Prefixes) != 0 {

		if setDefault {
			obj.Prefixes().clearHolderSlice()
			for _, item := range obj.obj.Prefixes {
				newObj := newIsisLspV6Prefix(obj.validator)
				newObj.self().obj = item
				obj.Prefixes().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Prefixes().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *isisLspIpv6ReachabilityTlv) checkUnique() {

}

func (obj *isisLspIpv6ReachabilityTlv) checkConstraint() {

}

func (obj *bgpEthernetSegmentDfElection) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.ElectionTimer != nil {
		if *obj.obj.ElectionTimer < 1 || *obj.obj.ElectionTimer > 300 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= BgpEthernetSegmentDfElection.ElectionTimer <= 300 but Got %d", *obj.obj.ElectionTimer))
		}

	}

}

func (obj *bgpEthernetSegmentDfElection) checkUnique() {

}

func (obj *bgpEthernetSegmentDfElection) checkConstraint() {

}

func (obj *bgpV4EvpnEvis) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.EviVxlan != nil {

		obj.EviVxlan().validateObj(setDefault)
	}

}

func (obj *bgpV4EvpnEvis) checkUnique() {

}

func (obj *bgpV4EvpnEvis) checkConstraint() {

}

func (obj *bgpRouteAdvanced) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *bgpRouteAdvanced) checkUnique() {

}

func (obj *bgpRouteAdvanced) checkConstraint() {

}

func (obj *bgpCommunity) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.AsNumber != nil {
		if *obj.obj.AsNumber < 0 || *obj.obj.AsNumber > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpCommunity.AsNumber <= 65535 but Got %d", *obj.obj.AsNumber))
		}

	}

	if obj.obj.AsCustom != nil {
		if *obj.obj.AsCustom < 0 || *obj.obj.AsCustom > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpCommunity.AsCustom <= 65535 but Got %d", *obj.obj.AsCustom))
		}

	}

}

func (obj *bgpCommunity) checkUnique() {

}

func (obj *bgpCommunity) checkConstraint() {

}

func (obj *bgpExtCommunity) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateHex(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpExtCommunity.Value"))
		}

	}

}

func (obj *bgpExtCommunity) checkUnique() {

}

func (obj *bgpExtCommunity) checkConstraint() {

}

func (obj *bgpAsPath) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Segments) != 0 {

		if setDefault {
			obj.Segments().clearHolderSlice()
			for _, item := range obj.obj.Segments {
				newObj := newBgpAsPathSegment(obj.validator)
				newObj.self().obj = item
				obj.Segments().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Segments().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *bgpAsPath) checkUnique() {

}

func (obj *bgpAsPath) checkConstraint() {

}

func (obj *bgpAddPath) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *bgpAddPath) checkUnique() {

}

func (obj *bgpAddPath) checkConstraint() {

}

func (obj *bgpSrteV4TunnelTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	if obj.obj.RemoteEndpointSubTlv != nil {

		obj.RemoteEndpointSubTlv().validateObj(setDefault)
	}

	if obj.obj.ColorSubTlv != nil {

		obj.ColorSubTlv().validateObj(setDefault)
	}

	if obj.obj.BindingSubTlv != nil {

		obj.BindingSubTlv().validateObj(setDefault)
	}

	if obj.obj.PreferenceSubTlv != nil {

		obj.PreferenceSubTlv().validateObj(setDefault)
	}

	if obj.obj.PolicyPrioritySubTlv != nil {

		obj.PolicyPrioritySubTlv().validateObj(setDefault)
	}

	if obj.obj.PolicyNameSubTlv != nil {

		obj.PolicyNameSubTlv().validateObj(setDefault)
	}

	if obj.obj.ExplicitNullLabelPolicySubTlv != nil {

		obj.ExplicitNullLabelPolicySubTlv().validateObj(setDefault)
	}

	if len(obj.obj.SegmentLists) != 0 {

		if setDefault {
			obj.SegmentLists().clearHolderSlice()
			for _, item := range obj.obj.SegmentLists {
				newObj := newBgpSrteSegmentList(obj.validator)
				newObj.self().obj = item
				obj.SegmentLists().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.SegmentLists().Items() {
			item.validateObj(setDefault)
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface BgpSrteV4TunnelTlv")
	}
}

func (obj *bgpSrteV4TunnelTlv) checkUnique() {
	if !obj.isUnique("bgpSrteV4TunnelTlv", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *bgpSrteV4TunnelTlv) checkConstraint() {

}

func (obj *bgpSrteV6TunnelTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	if obj.obj.RemoteEndpointSubTlv != nil {

		obj.RemoteEndpointSubTlv().validateObj(setDefault)
	}

	if obj.obj.ColorSubTlv != nil {

		obj.ColorSubTlv().validateObj(setDefault)
	}

	if obj.obj.BindingSubTlv != nil {

		obj.BindingSubTlv().validateObj(setDefault)
	}

	if obj.obj.PreferenceSubTlv != nil {

		obj.PreferenceSubTlv().validateObj(setDefault)
	}

	if obj.obj.PolicyPrioritySubTlv != nil {

		obj.PolicyPrioritySubTlv().validateObj(setDefault)
	}

	if obj.obj.PolicyNameSubTlv != nil {

		obj.PolicyNameSubTlv().validateObj(setDefault)
	}

	if obj.obj.ExplicitNullLabelPolicySubTlv != nil {

		obj.ExplicitNullLabelPolicySubTlv().validateObj(setDefault)
	}

	if len(obj.obj.SegmentLists) != 0 {

		if setDefault {
			obj.SegmentLists().clearHolderSlice()
			for _, item := range obj.obj.SegmentLists {
				newObj := newBgpSrteSegmentList(obj.validator)
				newObj.self().obj = item
				obj.SegmentLists().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.SegmentLists().Items() {
			item.validateObj(setDefault)
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface BgpSrteV6TunnelTlv")
	}
}

func (obj *bgpSrteV6TunnelTlv) checkUnique() {
	if !obj.isUnique("bgpSrteV6TunnelTlv", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *bgpSrteV6TunnelTlv) checkConstraint() {

}

func (obj *bgpV6EvpnEvis) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.EviVxlan != nil {

		obj.EviVxlan().validateObj(setDefault)
	}

}

func (obj *bgpV6EvpnEvis) checkUnique() {

}

func (obj *bgpV6EvpnEvis) checkConstraint() {

}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.RemoteVtepAddress != nil {

		err := obj.validateIpv4(obj.RemoteVtepAddress())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on VxlanV4TunnelDestinationIPModeUnicastVtep.RemoteVtepAddress"))
		}

	}

	if len(obj.obj.ArpSuppressionCache) != 0 {

		if setDefault {
			obj.ArpSuppressionCache().clearHolderSlice()
			for _, item := range obj.obj.ArpSuppressionCache {
				newObj := newVxlanTunnelDestinationIPModeUnicastArpSuppressionCache(obj.validator)
				newObj.self().obj = item
				obj.ArpSuppressionCache().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.ArpSuppressionCache().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) checkUnique() {

}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) checkConstraint() {

}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.RemoteVtepAddress != nil {

		err := obj.validateIpv6(obj.RemoteVtepAddress())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on VxlanV6TunnelDestinationIPModeUnicastVtep.RemoteVtepAddress"))
		}

	}

	if len(obj.obj.ArpSuppressionCache) != 0 {

		if setDefault {
			obj.ArpSuppressionCache().clearHolderSlice()
			for _, item := range obj.obj.ArpSuppressionCache {
				newObj := newVxlanTunnelDestinationIPModeUnicastArpSuppressionCache(obj.validator)
				newObj.self().obj = item
				obj.ArpSuppressionCache().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.ArpSuppressionCache().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) checkUnique() {

}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) checkConstraint() {

}

func (obj *patternFlowIpv4PriorityRawCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4PriorityRawCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4PriorityRawCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4PriorityRawCounter) checkUnique() {

}

func (obj *patternFlowIpv4PriorityRawCounter) checkConstraint() {

}

func (obj *patternFlowIpv4TosPrecedence) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosPrecedence.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 7 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4TosPrecedence.Values <= 7 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4TosPrecedence) checkUnique() {

}

func (obj *patternFlowIpv4TosPrecedence) checkConstraint() {

}

func (obj *patternFlowIpv4TosDelay) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosDelay.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4TosDelay.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4TosDelay) checkUnique() {

}

func (obj *patternFlowIpv4TosDelay) checkConstraint() {

}

func (obj *patternFlowIpv4TosThroughput) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosThroughput.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4TosThroughput.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4TosThroughput) checkUnique() {

}

func (obj *patternFlowIpv4TosThroughput) checkConstraint() {

}

func (obj *patternFlowIpv4TosReliability) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosReliability.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4TosReliability.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4TosReliability) checkUnique() {

}

func (obj *patternFlowIpv4TosReliability) checkConstraint() {

}

func (obj *patternFlowIpv4TosMonetary) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosMonetary.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4TosMonetary.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4TosMonetary) checkUnique() {

}

func (obj *patternFlowIpv4TosMonetary) checkConstraint() {

}

func (obj *patternFlowIpv4TosUnused) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosUnused.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 1 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4TosUnused.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4TosUnused) checkUnique() {

}

func (obj *patternFlowIpv4TosUnused) checkConstraint() {

}

func (obj *patternFlowIpv4DscpPhb) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 63 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpPhb.Value <= 63 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 63 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4DscpPhb.Values <= 63 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4DscpPhb) checkUnique() {

}

func (obj *patternFlowIpv4DscpPhb) checkConstraint() {

}

func (obj *patternFlowIpv4DscpEcn) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 3 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpEcn.Value <= 3 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 3 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternFlowIpv4DscpEcn.Values <= 3 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternFlowIpv4DscpEcn) checkUnique() {

}

func (obj *patternFlowIpv4DscpEcn) checkConstraint() {

}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionExtensionLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionExtensionLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) checkUnique() {

}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) checkConstraint() {

}

func (obj *patternFlowGtpExtensionContentsCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 281474976710655 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionContentsCounter.Start <= 281474976710655 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 281474976710655 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionContentsCounter.Step <= 281474976710655 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpExtensionContentsCounter) checkUnique() {

}

func (obj *patternFlowGtpExtensionContentsCounter) checkConstraint() {

}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionNextExtensionHeaderCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionNextExtensionHeaderCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) checkUnique() {

}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) checkConstraint() {

}

func (obj *patternFlowIcmpEchoTypeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIcmpEchoTypeCounter) checkUnique() {

}

func (obj *patternFlowIcmpEchoTypeCounter) checkConstraint() {

}

func (obj *patternFlowIcmpEchoCodeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoCodeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoCodeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIcmpEchoCodeCounter) checkUnique() {

}

func (obj *patternFlowIcmpEchoCodeCounter) checkConstraint() {

}

func (obj *patternFlowIcmpEchoIdentifierCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoIdentifierCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoIdentifierCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIcmpEchoIdentifierCounter) checkUnique() {

}

func (obj *patternFlowIcmpEchoIdentifierCounter) checkConstraint() {

}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoSequenceNumberCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoSequenceNumberCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) checkUnique() {

}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) checkConstraint() {

}

func (obj *patternFlowIcmpv6EchoTypeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIcmpv6EchoTypeCounter) checkUnique() {

}

func (obj *patternFlowIcmpv6EchoTypeCounter) checkConstraint() {

}

func (obj *patternFlowIcmpv6EchoCodeCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoCodeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoCodeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIcmpv6EchoCodeCounter) checkUnique() {

}

func (obj *patternFlowIcmpv6EchoCodeCounter) checkConstraint() {

}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoIdentifierCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoIdentifierCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) checkUnique() {

}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) checkConstraint() {

}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoSequenceNumberCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoSequenceNumberCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) checkUnique() {

}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) checkConstraint() {

}

func (obj *isisLspneighbor) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.SystemId != nil {

		err := obj.validateHex(obj.SystemId())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on IsisLspneighbor.SystemId"))
		}

	}

}

func (obj *isisLspneighbor) checkUnique() {

}

func (obj *isisLspneighbor) checkConstraint() {

}

func (obj *isisLspV4Prefix) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *isisLspV4Prefix) checkUnique() {

}

func (obj *isisLspV4Prefix) checkConstraint() {

}

func (obj *isisLspExtendedV4Prefix) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Ipv4Address != nil {

		err := obj.validateIpv4(obj.Ipv4Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on IsisLspExtendedV4Prefix.Ipv4Address"))
		}

	}

	if obj.obj.PrefixAttributes != nil {

		obj.PrefixAttributes().validateObj(setDefault)
	}

}

func (obj *isisLspExtendedV4Prefix) checkUnique() {

}

func (obj *isisLspExtendedV4Prefix) checkConstraint() {

}

func (obj *isisLspV6Prefix) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Ipv6Address != nil {

		err := obj.validateIpv6(obj.Ipv6Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on IsisLspV6Prefix.Ipv6Address"))
		}

	}

	if obj.obj.PrefixAttributes != nil {

		obj.PrefixAttributes().validateObj(setDefault)
	}

}

func (obj *isisLspV6Prefix) checkUnique() {

}

func (obj *isisLspV6Prefix) checkConstraint() {

}

func (obj *bgpV4EviVxlan) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.BroadcastDomains) != 0 {

		if setDefault {
			obj.BroadcastDomains().clearHolderSlice()
			for _, item := range obj.obj.BroadcastDomains {
				newObj := newBgpV4EviVxlanBroadcastDomain(obj.validator)
				newObj.self().obj = item
				obj.BroadcastDomains().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.BroadcastDomains().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.PmsiLabel != nil {
		if *obj.obj.PmsiLabel < 0 || *obj.obj.PmsiLabel > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpV4EviVxlan.PmsiLabel <= 16777215 but Got %d", *obj.obj.PmsiLabel))
		}

	}

	if obj.obj.AdLabel != nil {
		if *obj.obj.AdLabel < 0 || *obj.obj.AdLabel > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpV4EviVxlan.AdLabel <= 16777215 but Got %d", *obj.obj.AdLabel))
		}

	}

	if obj.obj.RouteDistinguisher != nil {

		obj.RouteDistinguisher().validateObj(setDefault)
	}

	if len(obj.obj.RouteTargetExport) != 0 {

		if setDefault {
			obj.RouteTargetExport().clearHolderSlice()
			for _, item := range obj.obj.RouteTargetExport {
				newObj := newBgpRouteTarget(obj.validator)
				newObj.self().obj = item
				obj.RouteTargetExport().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.RouteTargetExport().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.RouteTargetImport) != 0 {

		if setDefault {
			obj.RouteTargetImport().clearHolderSlice()
			for _, item := range obj.obj.RouteTargetImport {
				newObj := newBgpRouteTarget(obj.validator)
				newObj.self().obj = item
				obj.RouteTargetImport().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.RouteTargetImport().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.L3RouteTargetExport) != 0 {

		if setDefault {
			obj.L3RouteTargetExport().clearHolderSlice()
			for _, item := range obj.obj.L3RouteTargetExport {
				newObj := newBgpRouteTarget(obj.validator)
				newObj.self().obj = item
				obj.L3RouteTargetExport().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.L3RouteTargetExport().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.L3RouteTargetImport) != 0 {

		if setDefault {
			obj.L3RouteTargetImport().clearHolderSlice()
			for _, item := range obj.obj.L3RouteTargetImport {
				newObj := newBgpRouteTarget(obj.validator)
				newObj.self().obj = item
				obj.L3RouteTargetImport().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.L3RouteTargetImport().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(setDefault)
	}

	if len(obj.obj.Communities) != 0 {

		if setDefault {
			obj.Communities().clearHolderSlice()
			for _, item := range obj.obj.Communities {
				newObj := newBgpCommunity(obj.validator)
				newObj.self().obj = item
				obj.Communities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Communities().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.ExtCommunities) != 0 {

		if setDefault {
			obj.ExtCommunities().clearHolderSlice()
			for _, item := range obj.obj.ExtCommunities {
				newObj := newBgpExtCommunity(obj.validator)
				newObj.self().obj = item
				obj.ExtCommunities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.ExtCommunities().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(setDefault)
	}

}

func (obj *bgpV4EviVxlan) checkUnique() {

}

func (obj *bgpV4EviVxlan) checkConstraint() {

}

func (obj *bgpAsPathSegment) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *bgpAsPathSegment) checkUnique() {

}

func (obj *bgpAsPathSegment) checkConstraint() {

}

func (obj *bgpSrteRemoteEndpointSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.AsNumber != nil {
		if *obj.obj.AsNumber < 0 || *obj.obj.AsNumber > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteRemoteEndpointSubTlv.AsNumber <= 4294967295 but Got %d", *obj.obj.AsNumber))
		}

	}

	if obj.obj.Ipv4Address != nil {

		err := obj.validateIpv4(obj.Ipv4Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteRemoteEndpointSubTlv.Ipv4Address"))
		}

	}

	if obj.obj.Ipv6Address != nil {

		err := obj.validateIpv6(obj.Ipv6Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteRemoteEndpointSubTlv.Ipv6Address"))
		}

	}

}

func (obj *bgpSrteRemoteEndpointSubTlv) checkUnique() {

}

func (obj *bgpSrteRemoteEndpointSubTlv) checkConstraint() {

}

func (obj *bgpSrteColorSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Color != nil {

		err := obj.validateHex(obj.Color())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteColorSubTlv.Color"))
		}

	}

}

func (obj *bgpSrteColorSubTlv) checkUnique() {

}

func (obj *bgpSrteColorSubTlv) checkConstraint() {

}

func (obj *bgpSrteBindingSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Ipv6Sid != nil {

		err := obj.validateIpv6(obj.Ipv6Sid())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteBindingSubTlv.Ipv6Sid"))
		}

	}

}

func (obj *bgpSrteBindingSubTlv) checkUnique() {

}

func (obj *bgpSrteBindingSubTlv) checkConstraint() {

}

func (obj *bgpSrtePreferenceSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Preference != nil {
		if *obj.obj.Preference < 0 || *obj.obj.Preference > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrtePreferenceSubTlv.Preference <= 4294967295 but Got %d", *obj.obj.Preference))
		}

	}

}

func (obj *bgpSrtePreferenceSubTlv) checkUnique() {

}

func (obj *bgpSrtePreferenceSubTlv) checkConstraint() {

}

func (obj *bgpSrtePolicyPrioritySubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.PolicyPriority != nil {
		if *obj.obj.PolicyPriority < 0 || *obj.obj.PolicyPriority > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrtePolicyPrioritySubTlv.PolicyPriority <= 255 but Got %d", *obj.obj.PolicyPriority))
		}

	}

}

func (obj *bgpSrtePolicyPrioritySubTlv) checkUnique() {

}

func (obj *bgpSrtePolicyPrioritySubTlv) checkConstraint() {

}

func (obj *bgpSrtePolicyNameSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.PolicyName != nil {
		if len(*obj.obj.PolicyName) < 1 || len(*obj.obj.PolicyName) > 32 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf(
					"1 <= length of BgpSrtePolicyNameSubTlv.PolicyName <= 32 but Got %d",
					len(*obj.obj.PolicyName)))
		}

	}

}

func (obj *bgpSrtePolicyNameSubTlv) checkUnique() {

}

func (obj *bgpSrtePolicyNameSubTlv) checkConstraint() {

}

func (obj *bgpSrteExplicitNullLabelPolicySubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *bgpSrteExplicitNullLabelPolicySubTlv) checkUnique() {

}

func (obj *bgpSrteExplicitNullLabelPolicySubTlv) checkConstraint() {

}

func (obj *bgpSrteSegmentList) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	if obj.obj.Weight != nil {
		if *obj.obj.Weight < 0 || *obj.obj.Weight > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSegmentList.Weight <= 4294967295 but Got %d", *obj.obj.Weight))
		}

	}

	if len(obj.obj.Segments) != 0 {

		if setDefault {
			obj.Segments().clearHolderSlice()
			for _, item := range obj.obj.Segments {
				newObj := newBgpSrteSegment(obj.validator)
				newObj.self().obj = item
				obj.Segments().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Segments().Items() {
			item.validateObj(setDefault)
		}

	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface BgpSrteSegmentList")
	}
}

func (obj *bgpSrteSegmentList) checkUnique() {
	if !obj.isUnique("bgpSrteSegmentList", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *bgpSrteSegmentList) checkConstraint() {

}

func (obj *bgpV6EviVxlan) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.BroadcastDomains) != 0 {

		if setDefault {
			obj.BroadcastDomains().clearHolderSlice()
			for _, item := range obj.obj.BroadcastDomains {
				newObj := newBgpV6EviVxlanBroadcastDomain(obj.validator)
				newObj.self().obj = item
				obj.BroadcastDomains().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.BroadcastDomains().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.PmsiLabel != nil {
		if *obj.obj.PmsiLabel < 0 || *obj.obj.PmsiLabel > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpV6EviVxlan.PmsiLabel <= 16777215 but Got %d", *obj.obj.PmsiLabel))
		}

	}

	if obj.obj.AdLabel != nil {
		if *obj.obj.AdLabel < 0 || *obj.obj.AdLabel > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpV6EviVxlan.AdLabel <= 16777215 but Got %d", *obj.obj.AdLabel))
		}

	}

	if obj.obj.RouteDistinguisher != nil {

		obj.RouteDistinguisher().validateObj(setDefault)
	}

	if len(obj.obj.RouteTargetExport) != 0 {

		if setDefault {
			obj.RouteTargetExport().clearHolderSlice()
			for _, item := range obj.obj.RouteTargetExport {
				newObj := newBgpRouteTarget(obj.validator)
				newObj.self().obj = item
				obj.RouteTargetExport().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.RouteTargetExport().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.RouteTargetImport) != 0 {

		if setDefault {
			obj.RouteTargetImport().clearHolderSlice()
			for _, item := range obj.obj.RouteTargetImport {
				newObj := newBgpRouteTarget(obj.validator)
				newObj.self().obj = item
				obj.RouteTargetImport().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.RouteTargetImport().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.L3RouteTargetExport) != 0 {

		if setDefault {
			obj.L3RouteTargetExport().clearHolderSlice()
			for _, item := range obj.obj.L3RouteTargetExport {
				newObj := newBgpRouteTarget(obj.validator)
				newObj.self().obj = item
				obj.L3RouteTargetExport().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.L3RouteTargetExport().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.L3RouteTargetImport) != 0 {

		if setDefault {
			obj.L3RouteTargetImport().clearHolderSlice()
			for _, item := range obj.obj.L3RouteTargetImport {
				newObj := newBgpRouteTarget(obj.validator)
				newObj.self().obj = item
				obj.L3RouteTargetImport().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.L3RouteTargetImport().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(setDefault)
	}

	if len(obj.obj.Communities) != 0 {

		if setDefault {
			obj.Communities().clearHolderSlice()
			for _, item := range obj.obj.Communities {
				newObj := newBgpCommunity(obj.validator)
				newObj.self().obj = item
				obj.Communities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Communities().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.ExtCommunities) != 0 {

		if setDefault {
			obj.ExtCommunities().clearHolderSlice()
			for _, item := range obj.obj.ExtCommunities {
				newObj := newBgpExtCommunity(obj.validator)
				newObj.self().obj = item
				obj.ExtCommunities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.ExtCommunities().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(setDefault)
	}

}

func (obj *bgpV6EviVxlan) checkUnique() {

}

func (obj *bgpV6EviVxlan) checkConstraint() {

}

func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.RemoteVmMac != nil {

		err := obj.validateMac(obj.RemoteVmMac())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on VxlanTunnelDestinationIPModeUnicastArpSuppressionCache.RemoteVmMac"))
		}

	}

	if obj.obj.RemoteVmIpv4 != nil {

		err := obj.validateIpv4(obj.RemoteVmIpv4())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on VxlanTunnelDestinationIPModeUnicastArpSuppressionCache.RemoteVmIpv4"))
		}

	}

}

func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) checkUnique() {

}

func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) checkConstraint() {

}

func (obj *patternFlowIpv4TosPrecedenceCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosPrecedenceCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosPrecedenceCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4TosPrecedenceCounter) checkUnique() {

}

func (obj *patternFlowIpv4TosPrecedenceCounter) checkConstraint() {

}

func (obj *patternFlowIpv4TosDelayCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosDelayCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosDelayCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4TosDelayCounter) checkUnique() {

}

func (obj *patternFlowIpv4TosDelayCounter) checkConstraint() {

}

func (obj *patternFlowIpv4TosThroughputCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosThroughputCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosThroughputCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4TosThroughputCounter) checkUnique() {

}

func (obj *patternFlowIpv4TosThroughputCounter) checkConstraint() {

}

func (obj *patternFlowIpv4TosReliabilityCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosReliabilityCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosReliabilityCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4TosReliabilityCounter) checkUnique() {

}

func (obj *patternFlowIpv4TosReliabilityCounter) checkConstraint() {

}

func (obj *patternFlowIpv4TosMonetaryCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosMonetaryCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosMonetaryCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4TosMonetaryCounter) checkUnique() {

}

func (obj *patternFlowIpv4TosMonetaryCounter) checkConstraint() {

}

func (obj *patternFlowIpv4TosUnusedCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosUnusedCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 1 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosUnusedCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4TosUnusedCounter) checkUnique() {

}

func (obj *patternFlowIpv4TosUnusedCounter) checkConstraint() {

}

func (obj *patternFlowIpv4DscpPhbCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 63 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpPhbCounter.Start <= 63 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 63 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpPhbCounter.Step <= 63 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4DscpPhbCounter) checkUnique() {

}

func (obj *patternFlowIpv4DscpPhbCounter) checkConstraint() {

}

func (obj *patternFlowIpv4DscpEcnCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 3 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpEcnCounter.Start <= 3 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 3 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpEcnCounter.Step <= 3 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternFlowIpv4DscpEcnCounter) checkUnique() {

}

func (obj *patternFlowIpv4DscpEcnCounter) checkConstraint() {

}

func (obj *isisLspPrefixAttributes) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *isisLspPrefixAttributes) checkUnique() {

}

func (obj *isisLspPrefixAttributes) checkConstraint() {

}

func (obj *bgpV4EviVxlanBroadcastDomain) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.CmacIpRange) != 0 {

		if setDefault {
			obj.CmacIpRange().clearHolderSlice()
			for _, item := range obj.obj.CmacIpRange {
				newObj := newBgpCMacIpRange(obj.validator)
				newObj.self().obj = item
				obj.CmacIpRange().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.CmacIpRange().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.EthernetTagId != nil {
		if *obj.obj.EthernetTagId < 0 || *obj.obj.EthernetTagId > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpV4EviVxlanBroadcastDomain.EthernetTagId <= 4294967295 but Got %d", *obj.obj.EthernetTagId))
		}

	}

}

func (obj *bgpV4EviVxlanBroadcastDomain) checkUnique() {

}

func (obj *bgpV4EviVxlanBroadcastDomain) checkConstraint() {

}

func (obj *bgpRouteDistinguisher) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *bgpRouteDistinguisher) checkUnique() {

}

func (obj *bgpRouteDistinguisher) checkConstraint() {

}

func (obj *bgpRouteTarget) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *bgpRouteTarget) checkUnique() {

}

func (obj *bgpRouteTarget) checkConstraint() {

}

func (obj *bgpSrteSegment) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	// SegmentType is required
	if obj.obj.SegmentType.Number() == 0 {
		obj.errors = append(obj.errors, "SegmentType is required field on interface BgpSrteSegment")
	}

	if obj.obj.TypeA != nil {

		obj.TypeA().validateObj(setDefault)
	}

	if obj.obj.TypeB != nil {

		obj.TypeB().validateObj(setDefault)
	}

	if obj.obj.TypeC != nil {

		obj.TypeC().validateObj(setDefault)
	}

	if obj.obj.TypeD != nil {

		obj.TypeD().validateObj(setDefault)
	}

	if obj.obj.TypeE != nil {

		obj.TypeE().validateObj(setDefault)
	}

	if obj.obj.TypeF != nil {

		obj.TypeF().validateObj(setDefault)
	}

	if obj.obj.TypeG != nil {

		obj.TypeG().validateObj(setDefault)
	}

	if obj.obj.TypeH != nil {

		obj.TypeH().validateObj(setDefault)
	}

	if obj.obj.TypeI != nil {

		obj.TypeI().validateObj(setDefault)
	}

	if obj.obj.TypeJ != nil {

		obj.TypeJ().validateObj(setDefault)
	}

	if obj.obj.TypeK != nil {

		obj.TypeK().validateObj(setDefault)
	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface BgpSrteSegment")
	}
}

func (obj *bgpSrteSegment) checkUnique() {
	if !obj.isUnique("bgpSrteSegment", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *bgpSrteSegment) checkConstraint() {

}

func (obj *bgpV6EviVxlanBroadcastDomain) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.CmacIpRange) != 0 {

		if setDefault {
			obj.CmacIpRange().clearHolderSlice()
			for _, item := range obj.obj.CmacIpRange {
				newObj := newBgpCMacIpRange(obj.validator)
				newObj.self().obj = item
				obj.CmacIpRange().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.CmacIpRange().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.EthernetTagId != nil {
		if *obj.obj.EthernetTagId < 0 || *obj.obj.EthernetTagId > 4294967295 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpV6EviVxlanBroadcastDomain.EthernetTagId <= 4294967295 but Got %d", *obj.obj.EthernetTagId))
		}

	}

}

func (obj *bgpV6EviVxlanBroadcastDomain) checkUnique() {

}

func (obj *bgpV6EviVxlanBroadcastDomain) checkConstraint() {

}

func (obj *bgpCMacIpRange) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	if obj.obj.MacAddresses != nil {

		obj.MacAddresses().validateObj(setDefault)
	}

	if obj.obj.L2Vni != nil {
		if *obj.obj.L2Vni < 0 || *obj.obj.L2Vni > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpCMacIpRange.L2Vni <= 16777215 but Got %d", *obj.obj.L2Vni))
		}

	}

	if obj.obj.Ipv4Addresses != nil {

		obj.Ipv4Addresses().validateObj(setDefault)
	}

	if obj.obj.Ipv6Addresses != nil {

		obj.Ipv6Addresses().validateObj(setDefault)
	}

	if obj.obj.L3Vni != nil {
		if *obj.obj.L3Vni < 0 || *obj.obj.L3Vni > 16777215 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpCMacIpRange.L3Vni <= 16777215 but Got %d", *obj.obj.L3Vni))
		}

	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(setDefault)
	}

	if len(obj.obj.Communities) != 0 {

		if setDefault {
			obj.Communities().clearHolderSlice()
			for _, item := range obj.obj.Communities {
				newObj := newBgpCommunity(obj.validator)
				newObj.self().obj = item
				obj.Communities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Communities().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.ExtCommunities) != 0 {

		if setDefault {
			obj.ExtCommunities().clearHolderSlice()
			for _, item := range obj.obj.ExtCommunities {
				newObj := newBgpExtCommunity(obj.validator)
				newObj.self().obj = item
				obj.ExtCommunities().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.ExtCommunities().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(setDefault)
	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface BgpCMacIpRange")
	}
}

func (obj *bgpCMacIpRange) checkUnique() {
	if !obj.isUnique("bgpCMacIpRange", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *bgpCMacIpRange) checkConstraint() {

}

func (obj *bgpSrteSegmentATypeSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentATypeSubTlv.Flags"))
		}

	}

	if obj.obj.Label != nil {
		if *obj.obj.Label < 0 || *obj.obj.Label > 1048575 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSegmentATypeSubTlv.Label <= 1048575 but Got %d", *obj.obj.Label))
		}

	}

	if obj.obj.Tc != nil {
		if *obj.obj.Tc < 0 || *obj.obj.Tc > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSegmentATypeSubTlv.Tc <= 7 but Got %d", *obj.obj.Tc))
		}

	}

	if obj.obj.Ttl != nil {
		if *obj.obj.Ttl < 0 || *obj.obj.Ttl > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSegmentATypeSubTlv.Ttl <= 255 but Got %d", *obj.obj.Ttl))
		}

	}

}

func (obj *bgpSrteSegmentATypeSubTlv) checkUnique() {

}

func (obj *bgpSrteSegmentATypeSubTlv) checkConstraint() {

}

func (obj *bgpSrteSegmentBTypeSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentBTypeSubTlv.Flags"))
		}

	}

	// Srv6Sid is required
	if obj.obj.Srv6Sid == "" {
		obj.errors = append(obj.errors, "Srv6Sid is required field on interface BgpSrteSegmentBTypeSubTlv")
	}
	if obj.obj.Srv6Sid != "" {

		err := obj.validateIpv6(obj.Srv6Sid())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentBTypeSubTlv.Srv6Sid"))
		}

	}

	if obj.obj.Srv6SidEndpointBehavior != nil {

		obj.Srv6SidEndpointBehavior().validateObj(setDefault)
	}

}

func (obj *bgpSrteSegmentBTypeSubTlv) checkUnique() {

}

func (obj *bgpSrteSegmentBTypeSubTlv) checkConstraint() {

}

func (obj *bgpSrteSegmentCTypeSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentCTypeSubTlv.Flags"))
		}

	}

	if obj.obj.SrAlgorithm != nil {
		if *obj.obj.SrAlgorithm < 0 || *obj.obj.SrAlgorithm > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSegmentCTypeSubTlv.SrAlgorithm <= 255 but Got %d", *obj.obj.SrAlgorithm))
		}

	}

	// Ipv4NodeAddress is required
	if obj.obj.Ipv4NodeAddress == "" {
		obj.errors = append(obj.errors, "Ipv4NodeAddress is required field on interface BgpSrteSegmentCTypeSubTlv")
	}
	if obj.obj.Ipv4NodeAddress != "" {

		err := obj.validateIpv4(obj.Ipv4NodeAddress())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentCTypeSubTlv.Ipv4NodeAddress"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(setDefault)
	}

}

func (obj *bgpSrteSegmentCTypeSubTlv) checkUnique() {

}

func (obj *bgpSrteSegmentCTypeSubTlv) checkConstraint() {

}

func (obj *bgpSrteSegmentDTypeSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentDTypeSubTlv.Flags"))
		}

	}

	if obj.obj.SrAlgorithm != nil {
		if *obj.obj.SrAlgorithm < 0 || *obj.obj.SrAlgorithm > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSegmentDTypeSubTlv.SrAlgorithm <= 255 but Got %d", *obj.obj.SrAlgorithm))
		}

	}

	// Ipv6NodeAddress is required
	if obj.obj.Ipv6NodeAddress == "" {
		obj.errors = append(obj.errors, "Ipv6NodeAddress is required field on interface BgpSrteSegmentDTypeSubTlv")
	}
	if obj.obj.Ipv6NodeAddress != "" {

		err := obj.validateIpv6(obj.Ipv6NodeAddress())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentDTypeSubTlv.Ipv6NodeAddress"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(setDefault)
	}

}

func (obj *bgpSrteSegmentDTypeSubTlv) checkUnique() {

}

func (obj *bgpSrteSegmentDTypeSubTlv) checkConstraint() {

}

func (obj *bgpSrteSegmentETypeSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentETypeSubTlv.Flags"))
		}

	}

	if obj.obj.LocalInterfaceId != nil {
		if *obj.obj.LocalInterfaceId < 0 || *obj.obj.LocalInterfaceId > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSegmentETypeSubTlv.LocalInterfaceId <= 2147483647 but Got %d", *obj.obj.LocalInterfaceId))
		}

	}

	// Ipv4NodeAddress is required
	if obj.obj.Ipv4NodeAddress == "" {
		obj.errors = append(obj.errors, "Ipv4NodeAddress is required field on interface BgpSrteSegmentETypeSubTlv")
	}
	if obj.obj.Ipv4NodeAddress != "" {

		err := obj.validateIpv4(obj.Ipv4NodeAddress())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentETypeSubTlv.Ipv4NodeAddress"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(setDefault)
	}

}

func (obj *bgpSrteSegmentETypeSubTlv) checkUnique() {

}

func (obj *bgpSrteSegmentETypeSubTlv) checkConstraint() {

}

func (obj *bgpSrteSegmentFTypeSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentFTypeSubTlv.Flags"))
		}

	}

	// LocalIpv4Address is required
	if obj.obj.LocalIpv4Address == "" {
		obj.errors = append(obj.errors, "LocalIpv4Address is required field on interface BgpSrteSegmentFTypeSubTlv")
	}
	if obj.obj.LocalIpv4Address != "" {

		err := obj.validateIpv4(obj.LocalIpv4Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentFTypeSubTlv.LocalIpv4Address"))
		}

	}

	// RemoteIpv4Address is required
	if obj.obj.RemoteIpv4Address == "" {
		obj.errors = append(obj.errors, "RemoteIpv4Address is required field on interface BgpSrteSegmentFTypeSubTlv")
	}
	if obj.obj.RemoteIpv4Address != "" {

		err := obj.validateIpv4(obj.RemoteIpv4Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentFTypeSubTlv.RemoteIpv4Address"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(setDefault)
	}

}

func (obj *bgpSrteSegmentFTypeSubTlv) checkUnique() {

}

func (obj *bgpSrteSegmentFTypeSubTlv) checkConstraint() {

}

func (obj *bgpSrteSegmentGTypeSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentGTypeSubTlv.Flags"))
		}

	}

	if obj.obj.LocalInterfaceId != nil {
		if *obj.obj.LocalInterfaceId < 0 || *obj.obj.LocalInterfaceId > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSegmentGTypeSubTlv.LocalInterfaceId <= 2147483647 but Got %d", *obj.obj.LocalInterfaceId))
		}

	}

	// LocalIpv6NodeAddress is required
	if obj.obj.LocalIpv6NodeAddress == "" {
		obj.errors = append(obj.errors, "LocalIpv6NodeAddress is required field on interface BgpSrteSegmentGTypeSubTlv")
	}
	if obj.obj.LocalIpv6NodeAddress != "" {

		err := obj.validateIpv6(obj.LocalIpv6NodeAddress())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentGTypeSubTlv.LocalIpv6NodeAddress"))
		}

	}

	if obj.obj.RemoteInterfaceId != nil {
		if *obj.obj.RemoteInterfaceId < 0 || *obj.obj.RemoteInterfaceId > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSegmentGTypeSubTlv.RemoteInterfaceId <= 2147483647 but Got %d", *obj.obj.RemoteInterfaceId))
		}

	}

	// RemoteIpv6NodeAddress is required
	if obj.obj.RemoteIpv6NodeAddress == "" {
		obj.errors = append(obj.errors, "RemoteIpv6NodeAddress is required field on interface BgpSrteSegmentGTypeSubTlv")
	}
	if obj.obj.RemoteIpv6NodeAddress != "" {

		err := obj.validateIpv6(obj.RemoteIpv6NodeAddress())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentGTypeSubTlv.RemoteIpv6NodeAddress"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(setDefault)
	}

}

func (obj *bgpSrteSegmentGTypeSubTlv) checkUnique() {

}

func (obj *bgpSrteSegmentGTypeSubTlv) checkConstraint() {

}

func (obj *bgpSrteSegmentHTypeSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentHTypeSubTlv.Flags"))
		}

	}

	// LocalIpv6Address is required
	if obj.obj.LocalIpv6Address == "" {
		obj.errors = append(obj.errors, "LocalIpv6Address is required field on interface BgpSrteSegmentHTypeSubTlv")
	}
	if obj.obj.LocalIpv6Address != "" {

		err := obj.validateIpv6(obj.LocalIpv6Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentHTypeSubTlv.LocalIpv6Address"))
		}

	}

	// RemoteIpv6Address is required
	if obj.obj.RemoteIpv6Address == "" {
		obj.errors = append(obj.errors, "RemoteIpv6Address is required field on interface BgpSrteSegmentHTypeSubTlv")
	}
	if obj.obj.RemoteIpv6Address != "" {

		err := obj.validateIpv6(obj.RemoteIpv6Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentHTypeSubTlv.RemoteIpv6Address"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(setDefault)
	}

}

func (obj *bgpSrteSegmentHTypeSubTlv) checkUnique() {

}

func (obj *bgpSrteSegmentHTypeSubTlv) checkConstraint() {

}

func (obj *bgpSrteSegmentITypeSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentITypeSubTlv.Flags"))
		}

	}

	// Ipv6NodeAddress is required
	if obj.obj.Ipv6NodeAddress == "" {
		obj.errors = append(obj.errors, "Ipv6NodeAddress is required field on interface BgpSrteSegmentITypeSubTlv")
	}
	if obj.obj.Ipv6NodeAddress != "" {

		err := obj.validateIpv6(obj.Ipv6NodeAddress())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentITypeSubTlv.Ipv6NodeAddress"))
		}

	}

	if obj.obj.Srv6Sid != nil {

		err := obj.validateIpv6(obj.Srv6Sid())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentITypeSubTlv.Srv6Sid"))
		}

	}

	if obj.obj.Srv6SidEndpointBehavior != nil {

		obj.Srv6SidEndpointBehavior().validateObj(setDefault)
	}

}

func (obj *bgpSrteSegmentITypeSubTlv) checkUnique() {

}

func (obj *bgpSrteSegmentITypeSubTlv) checkConstraint() {

}

func (obj *bgpSrteSegmentJTypeSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentJTypeSubTlv.Flags"))
		}

	}

	if obj.obj.SrAlgorithm != nil {
		if *obj.obj.SrAlgorithm < 0 || *obj.obj.SrAlgorithm > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSegmentJTypeSubTlv.SrAlgorithm <= 255 but Got %d", *obj.obj.SrAlgorithm))
		}

	}

	if obj.obj.LocalInterfaceId != nil {
		if *obj.obj.LocalInterfaceId < 0 || *obj.obj.LocalInterfaceId > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSegmentJTypeSubTlv.LocalInterfaceId <= 2147483647 but Got %d", *obj.obj.LocalInterfaceId))
		}

	}

	// LocalIpv6NodeAddress is required
	if obj.obj.LocalIpv6NodeAddress == "" {
		obj.errors = append(obj.errors, "LocalIpv6NodeAddress is required field on interface BgpSrteSegmentJTypeSubTlv")
	}
	if obj.obj.LocalIpv6NodeAddress != "" {

		err := obj.validateIpv6(obj.LocalIpv6NodeAddress())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentJTypeSubTlv.LocalIpv6NodeAddress"))
		}

	}

	if obj.obj.RemoteInterfaceId != nil {
		if *obj.obj.RemoteInterfaceId < 0 || *obj.obj.RemoteInterfaceId > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSegmentJTypeSubTlv.RemoteInterfaceId <= 2147483647 but Got %d", *obj.obj.RemoteInterfaceId))
		}

	}

	// RemoteIpv6NodeAddress is required
	if obj.obj.RemoteIpv6NodeAddress == "" {
		obj.errors = append(obj.errors, "RemoteIpv6NodeAddress is required field on interface BgpSrteSegmentJTypeSubTlv")
	}
	if obj.obj.RemoteIpv6NodeAddress != "" {

		err := obj.validateIpv6(obj.RemoteIpv6NodeAddress())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentJTypeSubTlv.RemoteIpv6NodeAddress"))
		}

	}

	if obj.obj.Srv6Sid != nil {

		err := obj.validateIpv6(obj.Srv6Sid())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentJTypeSubTlv.Srv6Sid"))
		}

	}

	if obj.obj.Srv6SidEndpointBehavior != nil {

		obj.Srv6SidEndpointBehavior().validateObj(setDefault)
	}

}

func (obj *bgpSrteSegmentJTypeSubTlv) checkUnique() {

}

func (obj *bgpSrteSegmentJTypeSubTlv) checkConstraint() {

}

func (obj *bgpSrteSegmentKTypeSubTlv) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentKTypeSubTlv.Flags"))
		}

	}

	if obj.obj.SrAlgorithm != nil {
		if *obj.obj.SrAlgorithm < 0 || *obj.obj.SrAlgorithm > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSegmentKTypeSubTlv.SrAlgorithm <= 255 but Got %d", *obj.obj.SrAlgorithm))
		}

	}

	// LocalIpv6Address is required
	if obj.obj.LocalIpv6Address == "" {
		obj.errors = append(obj.errors, "LocalIpv6Address is required field on interface BgpSrteSegmentKTypeSubTlv")
	}
	if obj.obj.LocalIpv6Address != "" {

		err := obj.validateIpv6(obj.LocalIpv6Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentKTypeSubTlv.LocalIpv6Address"))
		}

	}

	// RemoteIpv6Address is required
	if obj.obj.RemoteIpv6Address == "" {
		obj.errors = append(obj.errors, "RemoteIpv6Address is required field on interface BgpSrteSegmentKTypeSubTlv")
	}
	if obj.obj.RemoteIpv6Address != "" {

		err := obj.validateIpv6(obj.RemoteIpv6Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentKTypeSubTlv.RemoteIpv6Address"))
		}

	}

	if obj.obj.Srv6Sid != nil {

		err := obj.validateIpv6(obj.Srv6Sid())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentKTypeSubTlv.Srv6Sid"))
		}

	}

	if obj.obj.Srv6SidEndpointBehavior != nil {

		obj.Srv6SidEndpointBehavior().validateObj(setDefault)
	}

}

func (obj *bgpSrteSegmentKTypeSubTlv) checkUnique() {

}

func (obj *bgpSrteSegmentKTypeSubTlv) checkConstraint() {

}

func (obj *mACRouteAddress) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// Address is required
	if obj.obj.Address == "" {
		obj.errors = append(obj.errors, "Address is required field on interface MACRouteAddress")
	}
	if obj.obj.Address != "" {

		err := obj.validateMac(obj.Address())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on MACRouteAddress.Address"))
		}

	}

	if obj.obj.Prefix != nil {
		if *obj.obj.Prefix < 0 || *obj.obj.Prefix > 48 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= MACRouteAddress.Prefix <= 48 but Got %d", *obj.obj.Prefix))
		}

	}

	if obj.obj.Count != nil {
		if *obj.obj.Count < 1 || *obj.obj.Count > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= MACRouteAddress.Count <= 2147483647 but Got %d", *obj.obj.Count))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 1 || *obj.obj.Step > 2147483647 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("1 <= MACRouteAddress.Step <= 2147483647 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *mACRouteAddress) checkUnique() {

}

func (obj *mACRouteAddress) checkConstraint() {

}

func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.LbLength != nil {
		if *obj.obj.LbLength < 0 || *obj.obj.LbLength > 128 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSRv6SIDEndpointBehaviorAndStructure.LbLength <= 128 but Got %d", *obj.obj.LbLength))
		}

	}

	if obj.obj.LnLength != nil {
		if *obj.obj.LnLength < 0 || *obj.obj.LnLength > 128 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSRv6SIDEndpointBehaviorAndStructure.LnLength <= 128 but Got %d", *obj.obj.LnLength))
		}

	}

	if obj.obj.FuncLength != nil {
		if *obj.obj.FuncLength < 0 || *obj.obj.FuncLength > 128 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSRv6SIDEndpointBehaviorAndStructure.FuncLength <= 128 but Got %d", *obj.obj.FuncLength))
		}

	}

	if obj.obj.ArgLength != nil {
		if *obj.obj.ArgLength < 0 || *obj.obj.ArgLength > 128 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSRv6SIDEndpointBehaviorAndStructure.ArgLength <= 128 but Got %d", *obj.obj.ArgLength))
		}

	}

}

func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) checkUnique() {

}

func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) checkConstraint() {

}

func (obj *bgpSrteSrMplsSid) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Label != nil {
		if *obj.obj.Label < 0 || *obj.obj.Label > 1048575 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSrMplsSid.Label <= 1048575 but Got %d", *obj.obj.Label))
		}

	}

	if obj.obj.Tc != nil {
		if *obj.obj.Tc < 0 || *obj.obj.Tc > 7 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSrMplsSid.Tc <= 7 but Got %d", *obj.obj.Tc))
		}

	}

	if obj.obj.Ttl != nil {
		if *obj.obj.Ttl < 0 || *obj.obj.Ttl > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= BgpSrteSrMplsSid.Ttl <= 255 but Got %d", *obj.obj.Ttl))
		}

	}

}

func (obj *bgpSrteSrMplsSid) checkUnique() {

}

func (obj *bgpSrteSrMplsSid) checkConstraint() {

}
