# PTP Operator
## Sponsors
- Joseph Richard <joseph.richard@redhat.com>
- Aneesh Puttur <aputtur@redhat.com>
- Kevin Clarke <kevin.clarke@intel.com>
- Eoghan Russell <eoghan.russell@intel.com>
## Background
The Precision Time Protocol (PTP) is a protocol used to synchronize clocks in a network. When used in conjunction with hardware support, PTP is capable of sub-microsecond accuracy, which is far better than is normally obtainable with NTP. PTP support is divided between the kernel and user space. The Linux kernel now includes support for PTP clocks, which are provided by network drivers. The actual implementation of the protocol is known as linuxptp, a PTPv2 implementation according to the IEEE standard 1588 for Linux.
## Existing project
Current projects:
* https://github.com/openshift/ptp-operator
* https://github.com/openshift/linuxptp-daemon 

Will live upstream as:
* https://github.com/k8snetworkplumbingwg/ptp-operator
* https://github.com/k8snetworkplumbingwg/linuxptp-daemon
## Upstreaming Rationale
Lack of existing upstream capability for PTP deployment on k8s.  This project is something that will be beneficial to a larger community, as PTP is important in lots of use cases.
Better capability for hardware manufacturers to deliver device-specific functionality in a way that will be consumed by customers.  This is especially important as modern hardware often requires device-specific configuration to perform optimally.
## Upstreaming Requirements
* Running on vanilla k8s
* Remove RHEL-specific docker files
* Remove openshift import paths


## Timeline
* Proposal submitted upstream
	* Aim for 20/4/2023
* Proposal approved upstream
  * Aim for 4/5/2023
* Create upstream project
  * Aim for before 11/5/2023
* Create daemon project
* Pull in recent openshift operator code
* Docker paths updated
* Operator images building from upstream
* First release

## Other considerations after upstreaming
* More hardware-configurability within plugins
  * Current support is for very-basic static configuration of e810.  Want to improve this as well as add support for other vendors down the road
* Consider using api to pass config
  * Currently just passing config from operator to daemon through config file.  May want to improve this in the future.
* Add CI support
  * Create CI tests for the upstream repo to run on new code submissions. This would be done with github actions and would probably use kind (https://kind.sigs.k8s.io) as the test environment.



## Maintainers
* Joseph Richard <joseph.richard@redhat.com>
* Aneesh Puttur <aputtur@redhat.com>
* Kevin Clarke <kevin.clarke@intel.com>
* Eoghan Russell <eoghan.russell@intel.com>
