from dataclasses import dataclass
from typing import List

@dataclass
class OltInterfaceVaniOnuPresentOnThisOlt:
    onu_fiber_distance: int
    detected_serial_number: str
    detected_registration_id: str

@dataclass
class OltInterfaceVani:
    onu_id: int
    management_tcont_alloc_id: int
    management_gemport_id: int
    onu_presence_state: str
    onu_presentt: OltInterfaceVaniOnuPresentOnThisOlt

@dataclass
class OltInterface:
    name: str
    type: str
    admin_state: str
    oper_state: str
    last_changed: str
    if_index: str
    lower_layer_if: str
    vani: OltInterfaceVani

@dataclass
class ComponentState:
    state_last_changed: str
    admin_state: str
    oper_state: str
    standby_state: str

@dataclass
class OltComponent:
    name: str
    type: str
    parent: str
    model: str
    state: ComponentState

@dataclass
class PonInfo(OltComponent):
    onuinfo: list[OltInterface]

    @classmethod
    def from_component(
        cls,
        component: OltComponent,
        onuinfo: List[OltInterface]
    ) -> "PonInfo":
        return cls(
            name=component.name,
            type=component.type,
            parent=component.parent,
            model=component.model,
            state=component.state,
            onuinfo=onuinfo
        )
@dataclass
class LtInfo(OltComponent):
    poninfo: list[PonInfo]

    @classmethod
    def from_component(
        cls,
        component: OltComponent,
        poninfo: List[PonInfo]
    ) -> "LtInfo":
        return cls(
            name=component.name,
            type=component.type,
            parent=component.parent,
            model=component.model,
            state=component.state,
            poninfo=poninfo
        )
@dataclass
class NtInfo(OltComponent):
    ltinfo: list[LtInfo]

    @classmethod
    def from_component(
        cls,
        component: OltComponent,
        ltinfo: List[LtInfo]
    ) -> "NtInfo":
        return cls(
            name=component.name,
            type=component.type,
            parent=component.parent,
            model=component.model,
            state=component.state,
            ltinfo=ltinfo
        )
