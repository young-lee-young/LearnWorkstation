from scapy.all import *

ans, unans = srp(Ether(dst="ff:ff:ff:ff:ff:ff") / ARP(pdst="192.168.203.0/24"), timeout=10)

ans.summary(lambda s, r: r.sprintf("%Ether.src% %ARP.psrc%"))
