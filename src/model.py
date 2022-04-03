from dataclasses import dataclass
from datetime import date
from typing import Optional


@dataclass(frozen=True)
class OrderLine:
    orderid: str
    sku: str  # stock-keeping init
    quantity: int


class Batch:
    def __init__(self, ref: str, sku: str, quantity: int, eta: Optional[date]):
        self.reference: str = ref
        self.sku: str = sku
        self.available_quantity: int = quantity
        self.eta: date = eta

    def allocate(self, line: OrderLine):
        self.available_quantity -= line.quantity
