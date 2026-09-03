class RpcGenerator:
    def __init__(self):
        pass

    def rpc_filter_hardware_state_with_class(self, class_name: str):
        rpc = f"""
            <filter xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" type="subtree">
                <hardware-state xmlns="urn:ietf:params:xml:ns:yang:ietf-hardware">
                    <component>
                        <class xmlns:nokia-hwi="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-hardware-identities">
                            {class_name}
                        </class>
                    </component>
                </hardware-state >
            </filter>
        """
        return rpc

    def rpc_filter_hardware_with_class(self, class_name: str):
        rpc = f"""
            <filter xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" type="subtree">
                <hardware xmlns="urn:ietf:params:xml:ns:yang:ietf-hardware">
                    <component>
                        <class xmlns:nokia-hwi="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-hardware-identities">
                            {class_name}
                        </class>
                    </component>
                </hardware>
            </filter>
        """
        return rpc
    def edit_device_name(self):
        rpc = f"""
        <config xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
            <system xmlns="urn:ietf:params:xml:ns:yang:ietf-system">
                <hostname>minghe-device</hostname>
            </system>
        </config>
        """
        return rpc
    def get_interfaces_info(self):
        rpc = f"""
            <filter xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" type="subtree">
                <interfaces xmlns="urn:ietf:params:xml:ns:yang:ietf-interfaces">
                    <interface>
                        <name>cp1_ont1</name>
                    </interface>
                </interfaces>
            </filter>
        """
        return rpc

    def rpc_filter_interfaces_with_type(self, type_name: str):
        rpc = f"""
            <filter xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" type="subtree">
                <interfaces xmlns="urn:ietf:params:xml:ns:yang:ietf-interfaces">
                    <type xmlns:bbf-xponift="urn:bbf:yang:bbf-xpon-if-type">
                        {type_name}
                    </type>
                </interfaces>
            </filter>
        """
        return rpc

    def rpc_filter_interfaces_state_with_type(self, type_name: str):
        rpc = f"""
            <filter xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" type="subtree">
                <interfaces-state xmlns="urn:ietf:params:xml:ns:yang:ietf-interfaces">
                    <interface>
                        <type xmlns:bbf-xponift="urn:bbf:yang:bbf-xpon-if-type">
                            {type_name}
                        </type>
                    </interface>
                </interfaces-state>
            </filter>
        """
        return rpc
    def get_hardware_component_list(self):
        rpc = f"""
            <filter xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" type="subtree">
                <hardware xmlns="urn:ietf:params:xml:ns:yang:ietf-hardware">
                    <component>
                    </component>
                </hardware>
            </filter>
        """
        return rpc
    def get_hardware_state_component_list(self):
        rpc = f"""
            <filter xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" type="subtree">
                <hardware-state xmlns="urn:ietf:params:xml:ns:yang:ietf-hardware">
                    <component>
                    </component>
                </hardware-state >
            </filter>
        """
        return rpc
    def get_online_nt_list(self):
        return self.rpc_filter_hardware_state_with_class("nokia-hwi:nt")
    def get_online_lt_list(self):
        return self.rpc_filter_hardware_state_with_class("nokia-hwi:lt")
    def get_config_nt_list(self):
        return self.rpc_filter_hardware_with_class("nokia-hwi:nt")
    def get_config_lt_list(self):
        return self.rpc_filter_hardware_with_class("nokia-hwi:lt")
    def get_pon_port_list(self):
        return self.rpc_filter_hardware_state_with_class("bbf-hwt:transceiver")
    def get_online_onu_list(self):
        return self.rpc_filter_interfaces_state_with_type("bbf-xponift:v-ani")