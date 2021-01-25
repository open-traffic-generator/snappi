"""package import

The package will only require a single import.
All package functionality should be accessible from a single top level class
in the package.
"""
import snappi


"""Api class

All functionality of the package is encapsulated in the Api class.
The Api class encapsulates individual openapi.yaml#/paths objects.
If the openapi.yaml#/paths/.../operationId is present it is used as the 
name of the Api class method.
If the operationId is not present then the name of the method will be 
the path name, prefixed with method, / replaced with _, all in lowercase.

paths:
  /config
    post:
      operationId: set_config -> Api class set_config method
"""
api = snappi.api()


"""Api class component schema object factory methods

All /paths request body objects which are a $ref to a /components/schemas 
object can be accessed using a factory method on the Api class.

paths:
  /config
    post:
      requestBody:
        content:
          application/json:
            schema:
              $ref: '/components/schemas/Config' -> Api config factory method that returns a Config(SnappiObject) instance
"""
config = api.config()


"""SnappiList

All /components/schemas properties that have a type of array with items that 
are a $ref to a /components/schemas object will inherit from the SnappiList 
class.  SnappiList is a container class that supports iteration, indexing and 
length which allows for OO access.

components:
  schemas:
    Config: -> a Config(SnappiObject) class
      type: object
      properties:
        ports: -> Config ports property that returns a PortList(SnappiList) instance 
          type: array
          items: 
            $ref: '/components/schemas/Port' -> PortList port factory method 
"""
tx_port, rx_port = config.ports \
    .port(name='Tx Port', location='10.36.74.26;02;13') \
    .port(name='Rx Port', location='10.36.74.26;02;14')

flow1 = config.flows.flow(name='Port Flow')
flow1.tx_rx.port.tx_name = tx_port.name
flow1.tx_rx.port.rx_name = rx_port.name

pkt = flow1.packet.ethernet().vlan().ipv4().tcp()
eth = flow1.packet[0]
eth = pkt[0]
eth = pkt[flow1.packet.ETHERNET]
ip = pkt[flow1.packet.IPV4]

eth.src = '00:00:01:00:00:01',
eth.dst = ['00:00:02:00:00:01', '00:00:02:00:00:01']
ip.src.increment(start='1.1.1.1', step='0.0.0.1', count=10)
ip.dst.decrement(start='1.1.2.200', step='0.0.0.1', count=10)
ip.priority.dscp.phb = ip.priority.dscp.PHB_CS1

flow1.size.fixed = 128
flow1.rate.pps = 1000
flow1.duration.fixed_packets(10000)

flow2 = config.flows.flow(name='Device Flow')
flow2.tx_rx.devices.tx_names = [config.devices[0].name]
flow2.tx_rx.devices.rx_names = [config.devices[1].name]
flow2.size.increment(start=128, end=256, step=2)

api.set_config(api.config)

# start capture on all configured ports, by default if no capture names are 
# provided all captures are started which is covered in the model documentation
# to start capture on a subset of ports provide the subset of port names to the
# ports parameter
api.capture_state.ports=[]
api.capture_state.state=api.capture_state.START
api.set_capture_state(api.capture_state)


api.flow_state = api.flow_state.START
api.set_flow_state(api.flow_state)

while len(config.flows) != len([
        m for m in (re.match('^stopped$', result['state'])
                    for result in api.get_port_results(api.result_portrequest))
]):
    continue

# stop the capture on a port and receive the capture as a stream of bytes
api.capture_results.port_name = rx_port.name
pcap_bytes = api.get_capture_results(api.capture_results)

# write the pcap bytes to a local file
pcap_filename = '%s.pcap' % rx_port.name
with open(pcap_filename, 'wb') as fid:
    fid.write(pcap_bytes)

# process the pcap file
for ts, pkt in dpkt.pcap.Reader(open(pcap_filename, 'rb')):
    print(ts, pkt)
